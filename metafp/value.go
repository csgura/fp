package metafp

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"
	"unicode"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/genfp/generator"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type Annotation struct {
	name   string
	params fp.Map[string, string]
}

func (r Annotation) Name() string {
	return r.name
}

func (r Annotation) Params() fp.Map[string, string] {
	return r.params
}

type TaggedStruct struct {
	Package *types.Package
	Name    string
	Scope   *types.Scope
	Struct  *types.Struct
	Fields  fp.Seq[StructField]
	Info    TypeInfo
	Tags    fp.Map[string, Annotation]
	RhsType fp.Option[TypeInfo]
}

func (r TaggedStruct) IsRecursive() bool {

	return r.Fields.Exists(func(v StructField) bool {
		return v.FieldType.HasTypeReference(mutable.EmptySet[string](), r.Info)
	})
}

func (r TaggedStruct) PackagedName(w genfp.ImportSet, workingPackage *types.Package) string {
	if workingPackage.Path() == r.Package.Path() {
		return r.Name
	}

	pk := w.GetImportedName(genfp.FromTypesPackage(r.Package))
	return fmt.Sprintf("%s.%s", pk, r.Name)
}

type TypeInfoExpr struct {
	Type TypeInfo
	Expr fp.Option[ast.Expr]
}

func (r TypeInfoExpr) String() string {
	return r.Type.String()
}

func (r TypeInfoExpr) TypeName(w genfp.ImportSet, wp genfp.WorkingPackage) string {

	if expr, ok := r.Expr.Unapply(); ok {
		_, iset := wp.EvalTypeExpr(expr)
		for _, v := range iset {
			w.AddImport(v)
		}
		return types.ExprString(expr)
	}

	return w.TypeName(wp, r.Type.Type)
}

func LookupStruct(pk *types.Package, name string) fp.Option[TaggedStruct] {
	l := pk.Scope().Lookup(name)

	if l == nil || l.Type().Underlying() == nil {
		return option.None[TaggedStruct]()
	}

	if st, ok := l.Type().Underlying().(*types.Struct); ok {
		fl := iterator.Map(iterator.Range(0, st.NumFields()), func(i int) StructField {
			f := st.Field(i)
			tn := typeInfo(f.Type())
			return StructField{
				Name:      f.Name(),
				FieldType: tn,
				Tag:       st.Tag(i),
				Embedded:  f.Embedded(),
				Pos:       f.Pos(),
			}
		}).ToSeq()

		info := typeInfo(l.Type())

		return option.Some(TaggedStruct{
			Name:    name,
			Scope:   l.Parent(),
			Package: l.Pkg(),
			Struct:  st,
			Fields:  fl,
			Info:    info,
		})
	}
	return option.None[TaggedStruct]()
}

func parseKeyValue(s string) fp.Tuple2[string, string] {
	s = strings.TrimSpace(s)
	idx := strings.Index(s, "=")
	if idx > 0 && len(s) > idx+1 {
		return as.Tuple2(strings.TrimSpace(s[:idx]), strings.TrimSpace(s[idx+1:]))
	}
	return as.Tuple2(s, "true")
}

func parseAnnotation(s string) fp.Tuple2[string, Annotation] {
	pstart := strings.Index(s, "(")
	if pstart > 0 {
		pend := strings.LastIndex(s, ")")
		if pend > pstart {
			name := strings.TrimSpace(s[:pstart])
			params := s[pstart+1 : pend]

			itr := iterator.FromSeq(strings.Split(params, ","))
			p := iterator.ToGoMap(iterator.Map(itr, parseKeyValue))
			return as.Tuple2(name, Annotation{
				name:   name,
				params: mutable.MapOf(p),
			})
		}

	}
	name := strings.TrimSpace(s)
	return as.Tuple2(name, Annotation{
		name: name,
	})

}

func extractTag(comment string) fp.Map[string, Annotation] {
	list := iterator.FromSeq(strings.Split(comment, "\n"))
	list = iterator.Map(list, strings.TrimSpace)
	list = list.Filter(as.Func2(strings.HasPrefix).ApplyLast("@"))
	ret := iterator.ToGoMap(iterator.Map(list, parseAnnotation))
	return mutable.MapOf(ret)
}

func GetTagsOfType(p []*packages.Package, name string) fp.Map[string, Annotation] {
	comment := seq.FlatMap(p, func(pk *packages.Package) fp.Seq[string] {
		s2 := seq.FlatMap(pk.Syntax, func(v *ast.File) fp.Seq[ast.Decl] {

			return v.Decls
		})

		s3 := seq.FlatMap(s2, func(v ast.Decl) fp.Seq[*ast.GenDecl] {
			switch r := v.(type) {
			case *ast.GenDecl:
				return seq.Of(r)
			}
			return seq.Of[*ast.GenDecl]()
		})

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[string] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[string] {
				if ts, ok := v.(*ast.TypeSpec); ok && ts.Name.Name == name {
					doc := option.Map(option.Of(ts.Doc).Or(as.Supplier(gdDoc)), (*ast.CommentGroup).Text)
					return doc.ToSeq()
				}
				return seq.Of[string]()
			})
		})
	}).Head()

	return option.Map(comment, extractTag).OrZero()
}

type PackagedName struct {
	Package string
	Name    string
}

func EvalTypeExprWithImport(pk *packages.Package, typeExpr ast.Expr) (types.Type, []genfp.ImportPackage) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  map[*ast.Ident]types.Object{},
	}
	err := types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)
	if err != nil {
		fmt.Printf("check expr err = %s\n", err)
	}

	var imports []genfp.ImportPackage
	for k, v := range info.Uses {
		if pk, ok := v.(*types.PkgName); ok {
			imports = append(imports, genfp.ImportPackage{
				Package: pk.Imported().Path(),
				Name:    k.Name,
			})

		}
	}

	ti := info.Types[typeExpr]

	for _, v := range info.Uses {
		if tn, ok := v.(*types.TypeName); ok {
			if tn.Type() == ti.Type {
				fmt.Printf("type aliased\n")
			}
		}
	}

	return ti.Type, imports
}

func EvalTypeExpr(pk *packages.Package, typeExpr ast.Expr) types.Type {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)
	ti := info.Types[typeExpr]
	return ti.Type
}

func checkType(pk *packages.Package, typeExpr ast.Expr) *types.Named {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)

	ti := info.Types[typeExpr]
	if named, ok := ti.Type.(*types.Named); ok {
		return named
	}
	return nil
}

func FindTaggedCompositeVariable(p []*packages.Package, typ PackagedName, tags ...string) fp.Seq[generator.TaggedLit] {
	p = as.Seq(p).Filter(func(v *packages.Package) bool {
		return v.PkgPath == typ.Package
	})
	return generator.FindTaggedCompositeVariable(p, typ.Name, tags...)
}

func FindTaggedStruct(p []*packages.Package, tags ...string) fp.Seq[TaggedStruct] {

	tagSeq := as.Seq(tags)
	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[TaggedStruct] {
		working := genfp.NewWorkingPackage(pk.Types, pk.Fset, pk.Syntax)
		s2 := seq.FlatMap(pk.Syntax, func(v *ast.File) fp.Seq[ast.Decl] {

			return v.Decls
		})

		s3 := seq.FlatMap(s2, func(v ast.Decl) fp.Seq[*ast.GenDecl] {
			switch r := v.(type) {
			case *ast.GenDecl:
				return seq.Of(r)
			}
			return seq.Of[*ast.GenDecl]()
		})

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[TaggedStruct] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[TaggedStruct] {
				if ts, ok := v.(*ast.TypeSpec); ok {
					doc := option.Map(option.Of(ts.Doc).Or(as.Supplier(gdDoc)), (*ast.CommentGroup).Text)
					if doc.Exists(func(comment string) bool {
						return tagSeq.Exists(func(tag string) bool { return strings.Contains(comment, tag) })
					}) {

						return option.FlatMap(LookupStruct(working.Package(), ts.Name.Name), func(ret TaggedStruct) fp.Option[TaggedStruct] {
							ret.Tags = option.Map(doc, extractTag).OrZero()
							if !tagSeq.Exists(ret.Tags.Contains) {
								return option.None[TaggedStruct]()
							}
							if _, ok := ts.Type.(*ast.SelectorExpr); ok {
								info := &types.Info{
									Types: make(map[ast.Expr]types.TypeAndValue),
								}
								err := types.CheckExpr(pk.Fset, pk.Types, ts.Type.End(), ts.Type, info)
								if err != nil {
									fmt.Printf("selector err = %s\n", err)
								}
								ti := info.Types[ts.Type]
								if _, ok := ti.Type.(*types.Named); ok {
									ret.RhsType = option.Some(typeInfo(ti.Type))

								} else {
									fmt.Printf("rhsType = %s\n", ti.Type.String())
								}
							} else if _, ok := ts.Type.(*ast.IndexExpr); ok {
								info := &types.Info{
									Types: make(map[ast.Expr]types.TypeAndValue),
								}
								pos := ts.Type.End()
								if ts.TypeParams != nil {
									pos = ts.TypeParams.Closing
								}
								err := types.CheckExpr(pk.Fset, pk.Types, pos, ts.Type, info)
								if err != nil {
									fmt.Printf("index expr err = %s\n", err)
								}
								ti := info.Types[ts.Type]
								if _, ok := ti.Type.(*types.Named); ok {
									ret.RhsType = option.Some(typeInfo(ti.Type))

								}
							} else if _, ok := ts.Type.(*ast.IndexListExpr); ok {
								info := &types.Info{
									Types: make(map[ast.Expr]types.TypeAndValue),
								}
								pos := ts.Type.End()
								if ts.TypeParams != nil {
									pos = ts.TypeParams.Closing
								}
								err := types.CheckExpr(pk.Fset, pk.Types, pos, ts.Type, info)
								if err != nil {
									fmt.Printf("index list err = %s\n", err)
								}
								ti := info.Types[ts.Type]
								if _, ok := ti.Type.(*types.Named); ok {
									ret.RhsType = option.Some(typeInfo(ti.Type))

								}

							}
							// else {
							// 	//fmt.Printf("name %s , epxr = %T\n", ret.Name, ts.Type)
							// }

							return option.Some(ret)
						}).ToSeq()
					}
				}
				return seq.Of[TaggedStruct]()
			})
		})
	})

}

type TypeParam struct {
	Name       string
	Constraint types.Type
	TypeName   *types.TypeName
}

func (r TypeParam) String() string {
	return fmt.Sprintf("%s %v", r.Name, r.Constraint)
}

func (r TypeParam) IsAny() bool {
	return typeInfo(r.Constraint).IsAny()
}

type NamedTypeInfo struct {
	Package    *types.Package
	Name       string
	Info       TypeInfo
	Underlying TypeInfo
}

func (r NamedTypeInfo) PackagedName(w genfp.ImportSet, working genfp.WorkingPackage) string {
	if r.Package != nil && r.Package.Path() != working.Path() {
		pk := w.GetImportedName(genfp.FromTypesPackage(r.Package))
		return fmt.Sprintf("%s.%s", pk, r.Name)
	}
	return r.Name
}

func (r NamedTypeInfo) GenericName() string {

	if r.Package != nil {
		return fmt.Sprintf("%s.%s", r.Package.Name(), r.Name)
	}

	return r.Name
}

func iterate[T, R any](len int, getter func(idx int) T, fn func(int, T) R) []R {
	ret := []R{}
	for i := 0; i < len; i++ {
		ret = append(ret, fn(i, getter(i)))
	}
	return ret
}

func typeId(tpe types.Type) string {
	switch realtp := tpe.(type) {
	case *types.Named:
		tpname := realtp.Origin().Obj().Name()
		nameWithPkg := tpname
		if realtp.Obj().Pkg() != nil {

			nameWithPkg = fmt.Sprintf("%s.%s", realtp.Obj().Pkg().Path(), tpname)
		}

		if realtp.TypeArgs() != nil {
			args := []string{}
			for i := 0; i < realtp.TypeArgs().Len(); i++ {
				args = append(args, typeId(realtp.TypeArgs().At(i)))
			}

			argsstr := strings.Join(args, ",")

			return fmt.Sprintf("%s[%s]", nameWithPkg, argsstr)
		} else {

			return nameWithPkg

		}

	case *types.Array:
		elemType := typeId(realtp.Elem())
		return fmt.Sprintf("[%d]%s", realtp.Len(), elemType)

	case *types.Map:
		keyType := typeId(realtp.Key())

		elemType := typeId(realtp.Elem())
		return fmt.Sprintf("map[%s]%s", keyType, elemType)
	case *types.Slice:
		elemType := typeId(realtp.Elem())
		return "[]" + elemType
	case *types.Pointer:
		elemType := typeId(realtp.Elem())
		return "*" + elemType
	case *types.Chan:
		elemType := typeId(realtp.Elem())
		switch realtp.Dir() {
		case types.RecvOnly:
			return "<-chan " + elemType
		case types.SendOnly:
			return "chan<- " + elemType
		default:
			return "chan " + elemType

		}
	case *types.Signature:
		argsstr := iterate(realtp.Params().Len(), realtp.Params().At, func(idx int, v *types.Var) string {
			return v.Name() + " " + typeId(v.Type())
		})

		resultstr := iterate(realtp.Results().Len(), realtp.Results().At, func(idx int, v *types.Var) string {
			return v.Name() + " " + typeId(v.Type())
		})

		return fmt.Sprintf("func (%s) (%s)", strings.Join(argsstr, ","), strings.Join(resultstr, ","))
	case *types.Struct:
		fields := iterate(realtp.NumFields(), realtp.Field, func(idx int, v *types.Var) string {
			if v.Embedded() {
				return fmt.Sprintf("%s %s",
					typeId(v.Type()),
					realtp.Tag(idx),
				)
			}
			return fmt.Sprintf("%s %s %s",
				v.Name(),
				typeId(v.Type()),
				realtp.Tag(idx),
			)
		})
		return fmt.Sprintf(`struct {
			%s
		}`, strings.Join(fields, "\n"))
	case *types.Interface:
		if realtp.NumMethods() == 0 {
			return "any"
		}
		embeded := iterate(realtp.NumEmbeddeds(), realtp.EmbeddedType, func(idx int, v types.Type) string {
			return typeId(realtp.EmbeddedType(idx))
		})

		fields := iterate(realtp.NumExplicitMethods(), realtp.ExplicitMethod, func(idx int, v *types.Func) string {
			m := realtp.ExplicitMethod(idx)

			return fmt.Sprintf("%s%s", m.Name(), typeId(m.Type())[4:])

		})
		return fmt.Sprintf(`interface {
			%s
			%s
		}`, strings.Join(embeded, "\n"), strings.Join(fields, "\n"))
	}

	return tpe.String()
}

type TypeInfo struct {
	ID        string
	Pkg       *types.Package
	TypeName  string
	Type      types.Type
	TypeArgs  fp.Seq[TypeInfo]
	TypeParam fp.Seq[TypeParam]
	Method    fp.Map[string, *types.Func]
}

func (r TypeInfo) PackagedName() PackagedName {
	if r.Pkg != nil {
		return PackagedName{
			Package: r.Pkg.Path(),
			Name:    r.TypeName,
		}
	}
	return PackagedName{
		Name: r.TypeName,
	}
}

func (r TypeInfo) IsSamePkg(other genfp.WorkingPackage) bool {
	return isSamePkg(other, genfp.FromTypesPackage(r.Pkg))
}

func (r TypeInfo) Fields() fp.Seq[StructField] {
	switch at := r.Type.(type) {
	case *types.Named:
		under := typeInfo(at.Underlying())
		return under.Fields()
	case *types.Struct:
		return iterate(at.NumFields(), at.Field, func(i int, f *types.Var) StructField {

			tn := typeInfo(f.Type())
			return StructField{
				Name:      f.Name(),
				FieldType: tn,
				Tag:       at.Tag(i),
				Embedded:  f.Embedded(),
				Pos:       f.Pos(),
			}
		})
	}
	return nil
}

func (r TypeInfo) HasTypeReference(checked fp.Set[string], refType TypeInfo) bool {

	if checked.Contains(r.ID) {
		return false
	}

	checked = checked.Incl(r.ID)

	if isSamePkg(genfp.FromTypesPackage(r.Pkg), genfp.FromTypesPackage(refType.Pkg)) && r.TypeName == refType.TypeName {
		return true
	}

	if r.TypeArgs.Exists(as.Func3(TypeInfo.HasTypeReference).ApplyLast2(checked, refType)) {
		return true
	}

	switch at := r.Type.(type) {
	case *types.Named:
		under := typeInfo(at.Underlying())
		return under.HasTypeReference(checked, refType)
	case *types.Struct:
		return iterator.Range(0, at.NumFields()).Exists(func(i int) bool {
			f := at.Field(i)
			tn := typeInfo(f.Type())

			return tn.HasTypeReference(checked, refType)
		})
	}
	return false

}
func (r TypeInfo) ReplaceTypeParam(mapping fp.Map[string, TypeInfo]) fp.Tuple2[fp.Set[string], TypeInfo] {

	if r.IsTypeParam() {
		mapped := mapping.Get(r.Name().Get())
		if mapped.IsEmpty() {
			return as.Tuple2(mutable.EmptySet[string](), r)
		} else {
			return as.Tuple2(mutable.SetOf(r.Name().Get()), mapped.Get())
		}
	}

	newArgs := seq.Map(r.TypeArgs, func(v TypeInfo) fp.Tuple2[fp.Set[string], TypeInfo] {
		return v.ReplaceTypeParam(mapping)
	})

	r.TypeArgs = seq.Map(newArgs, func(v fp.Tuple2[fp.Set[string], TypeInfo]) TypeInfo {
		return v.I2
	})

	usedParam := seq.Reduce(seq.Map(newArgs, fp.Tuple2[fp.Set[string], TypeInfo].Head), monoid.MergeSet[string]())

	return as.Tuple2(usedParam, r)
}

func isSamePkg(p1 genfp.PackageId, p2 genfp.PackageId) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}

	return p1.Path() == p2.Path()
}

// Seq[Tuple2[A,B]] 같은 타입이  Seq[T any]  같은  타입의 instantiated 인지 확인하는 함수
func (r TypeInfo) IsInstantiatedOf(typeParam fp.Seq[TypeParam], genericType TypeInfo) ConstraintCheckResult {

	// package가 동일해야 함
	if !isSamePkg(genfp.FromTypesPackage(r.Pkg), genfp.FromTypesPackage(genericType.Pkg)) {
		return ConstraintCheckResult{}
	}

	// 타입 이름이 동일해야 함
	//	fmt.Printf("compare %s(%s), %s(%s)\n", r, r.TypeName, genericType, genericType.TypeName)
	if r.TypeName != genericType.TypeName {
		return ConstraintCheckResult{}
	}

	// 타입 아규먼트 개수가 동일해야 함
	if r.TypeArgs.Size() != genericType.TypeArgs.Size() {
		return ConstraintCheckResult{}
	}

	ret := ConstraintCheck(typeParam, genericType, r.TypeArgs)
	//	fmt.Printf("compare %s, %s  => %t\n", r, genericType, ret)
	return ret

	// fmt.Printf("this args = %v\n", r.TypeArgs)
	// fmt.Printf("that args = %v\n", hasTypeParam.TypeArgs)

	// fmt.Printf("that type = %v\n", hasTypeParam.Type.String())

	// return true
}

func (r TypeInfo) PkgName() string {
	if r.Pkg != nil {
		return r.Pkg.Name()
	}
	return ""
}
func (r TypeInfo) String() string {
	name := r.Name().OrZero()
	if r.Pkg != nil {
		if r.TypeParam.Size() > 0 {
			return fmt.Sprintf("%s.%s[%s](%s)", r.PkgName(), name, r.TypeParam.MakeString(","), r.TypeArgs.MakeString(","))
		}
		return fmt.Sprintf("%s.%s", r.PkgName(), name)
	}
	if r.TypeParam.Size() > 0 {
		return fmt.Sprintf("%s[%s](%s)", name, r.TypeParam.MakeString(","), r.TypeArgs.MakeString(","))
	}
	if name == "" {
		return r.Type.String()
	}
	return name

}
func (r TypeInfo) ResultType() TypeInfo {
	switch at := r.Type.(type) {

	case *types.Signature:
		if at.Results().Len() == 1 {
			rtype := at.Results().At(0)
			rtypeInfo := typeInfo(rtype.Type())
			return rtypeInfo
		}
	}

	return r
}

func (r TypeInfo) IsInstanceOf(tc TypeClass) bool {

	switch at := r.Type.(type) {
	case *types.Named:
		if at.Obj().Pkg().Path() == tc.Package.Path() && at.Obj().Name() == tc.Name {
			return true
		}
	}

	return false

}

func (r TypeInfo) IsAny() bool {
	switch at := r.Type.(type) {
	case *types.Alias:
		return typeInfo(at.Rhs()).IsAny()
	case *types.Interface:
		if at.NumMethods() == 0 && at.NumEmbeddeds() == 0 {
			return true
		}
	}
	return false
}

func (r TypeInfo) IsFunc() bool {
	switch r.Type.(type) {
	case *types.Signature:
		return true
	}
	return false
}

func (r TypeInfo) NumArgs() int {
	switch at := r.Type.(type) {
	case *types.Signature:
		if at.Params() == nil {
			return 0
		}
		return at.Params().Len()
	}
	return 0
}

type atLen[T any] interface {
	Len() int
	At(int) T
}

func atLenToSeq[T any, A atLen[T]](a A) fp.Seq[T] {
	return iterator.Map(iterator.Range(0, a.Len()), func(idx int) T {
		return a.At(idx)
	}).ToSeq()
}

func (r TypeInfo) FuncArgs() fp.Seq[TypeInfo] {
	switch at := r.Type.(type) {
	case *types.Signature:
		if at.Params() == nil {
			return nil
		}
		ret := atLenToSeq[*types.Var](at.Params())
		return seq.Map(ret, func(v *types.Var) TypeInfo {
			return typeInfo(v.Type())
		})
	}
	return nil
}

func (r TypeInfo) Name() fp.Option[string] {
	switch at := r.Type.(type) {
	case *types.Named:
		return option.Some(at.Obj().Name())
	case *types.Basic:
		return option.Some(at.Name())
	case *types.TypeParam:
		return option.Some(at.Obj().Name())
	case *types.Signature:
	}
	return option.None[string]()
}

func (r TypeInfo) IsNamed() bool {
	switch r.Type.(type) {
	case *types.Named:
		return true
	}
	return false
}

func (r TypeInfo) AsNamed() fp.Option[NamedTypeInfo] {
	switch at := r.Type.(type) {
	case *types.Named:
		return option.Some(NamedTypeInfo{
			Package:    r.Pkg,
			Name:       r.Name().Get(),
			Info:       r,
			Underlying: typeInfo(at.Underlying()),
		})
	}
	return option.None[NamedTypeInfo]()
}

func (r TypeInfo) IsPrintable() bool {
	switch r.Type.(type) {
	case *types.Signature:
		return false
	}
	return true
}

func (r TypeInfo) IsTypeParam() bool {
	switch r.Type.(type) {
	case *types.TypeParam:
		return true
	}
	return false
}

func (r TypeInfo) IsBasic() bool {
	switch r.Type.(type) {
	case *types.Basic:
		return true
	}
	return false
}

func (r TypeInfo) IsSlice() bool {
	switch r.Type.(type) {
	case *types.Slice:
		return true
	}
	return false
}

func (r TypeInfo) IsArray() bool {
	switch r.Type.(type) {
	case *types.Array:
		return true
	}
	return false
}

func (r TypeInfo) IsStruct() bool {
	switch r.Type.(type) {
	case *types.Struct:
		return true
	}
	return false
}

func (r TypeInfo) Underlying() TypeInfo {
	if r.IsNamed() {
		return r.AsNamed().Get().Underlying
	}

	return r
}

func (r TypeInfo) IsMap() bool {
	switch r.Type.(type) {
	case *types.Map:
		return true
	}
	return false
}

func (r TypeInfo) IsPtr() bool {
	switch r.Type.(type) {
	case *types.Pointer:
		return true
	}
	return false
}

func (r TypeInfo) IsTuple() bool {
	switch nt := r.Type.(type) {
	case *types.Named:
		if nt.Obj().Pkg().Path() == "github.com/csgura/fp" && strings.HasPrefix(nt.Obj().Name(), "Tuple") {
			return true
		}
	}
	return false
}

func (r TypeInfo) IsOption() bool {
	switch nt := r.Type.(type) {
	case *types.Named:
		if nt.Obj().Pkg().Path() == "github.com/csgura/fp" && nt.Obj().Name() == "Option" {
			return true
		}
	}
	return false
}

func (r TypeInfo) IsNilable() bool {
	switch atp := r.Type.(type) {
	case *types.Pointer:
		return true
	case *types.Slice:
		return true
	case *types.Map:
		return true
	case *types.Chan:
		return true
	case *types.Signature:
		return true
	case *types.Interface:
		return true
	case *types.Basic:
		return atp.Kind() == types.String
	}
	return false
}

func (r TypeInfo) TypeParamDecl(w genfp.ImportSet, cwd genfp.WorkingPackage) string {
	if len(r.TypeParam) > 0 {
		return "[" + iterator.Map(seq.Iterator(r.TypeParam), func(v TypeParam) string {
			tn := w.TypeName(cwd, v.Constraint)
			return fmt.Sprintf("%s %s", v.Name, tn)
		}).MakeString(",") + "]"
	}
	return ""

}

func (r TypeInfo) TypeParamIns(w genfp.ImportSet, cwd genfp.WorkingPackage) string {
	if len(r.TypeParam) > 0 {
		return "[" + iterator.Map(seq.Iterator(r.TypeParam), func(v TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}
	return ""
}

func (r TypeInfo) TypeDeclStr(w genfp.ImportSet, cwd genfp.WorkingPackage) string {
	if r.TypeParam.Size() > 0 {
		return w.TypeName(cwd, r.Type) + r.TypeParamDecl(w, cwd)
	}
	return w.TypeName(cwd, r.Type)
}

func (r TypeInfo) TypeStr(w genfp.ImportSet, cwd genfp.WorkingPackage) string {
	if r.TypeParam.Size() > 0 {
		return w.TypeName(cwd, r.Type) + r.TypeParamIns(w, cwd)
	}
	return w.TypeName(cwd, r.Type)
}

func (r TypeInfo) GenericType() TypeInfo {
	switch nt := r.Type.(type) {
	case *types.Named:
		return typeInfo(nt.Obj().Type())
	}
	return r
}

type StructField struct {
	Name      string
	FieldType TypeInfo
	Tag       string
	Embedded  bool
	Pos       token.Pos
}

func (r StructField) TypeInfoExpr(wp genfp.WorkingPackage) TypeInfoExpr {
	if expr, ok := wp.FindNode(r.Pos).(*ast.Field); ok {
		return TypeInfoExpr{
			Type: r.FieldType,
			Expr: option.Some(expr.Type),
		}
	}
	return TypeInfoExpr{
		Type: r.FieldType,
		Expr: option.None[ast.Expr](),
	}
}

func (r StructField) TypeName(w genfp.ImportSet, wp genfp.WorkingPackage) string {

	if expr, ok := wp.FindNode(r.Pos).(*ast.Field); ok {
		_, iset := wp.EvalTypeExpr(expr.Type)
		for _, v := range iset {
			w.AddImport(v)
		}
		return types.ExprString(expr.Type)
	}

	return w.TypeName(wp, r.FieldType.Type)
}

func (r StructField) Public() bool {
	return !unicode.IsLower([]rune(r.Name)[0])
}

func typeArgs(args *types.TypeList) fp.Seq[TypeInfo] {
	if args == nil {
		return seq.Empty[TypeInfo]()
	}
	ret := iterator.Map(iterator.Range(0, args.Len()), func(i int) TypeInfo {
		return typeInfo(args.At(i))
	}).ToSeq()
	return ret
}
func typeParam(args *types.TypeParamList) fp.Seq[TypeParam] {
	if args == nil {
		return seq.Empty[TypeParam]()
	}
	params := iterator.Map(iterator.Range(0, args.Len()), func(i int) TypeParam {
		return TypeParam{
			Name:       args.At(i).Obj().Name(),
			Constraint: args.At(i).Constraint(),
			TypeName:   args.At(i).Obj(),
		}
	}).ToSeq()
	return params
}

func GetTypeInfo(tpe types.Type) TypeInfo {
	return typeInfo(tpe)
}

func typeInfo(tpe types.Type) TypeInfo {
	//	fmt.Printf("get info of %s\n", tpe.String())

	id := typeId(tpe)
	switch realtp := tpe.(type) {
	case *types.TypeParam:
		return TypeInfo{
			ID:       id,
			Type:     tpe,
			TypeName: realtp.Obj().Name(),
		}
	case *types.Basic:
		return TypeInfo{
			ID:       id,
			Type:     tpe,
			TypeName: realtp.Name(),
		}
	case *types.Named:
		args := typeArgs(realtp.TypeArgs())
		params := typeParam(realtp.TypeParams())

		methodMap := func() fp.Map[string, *types.Func] {
			if realtp.NumMethods() == 0 {
				if underIntf, ok := realtp.Underlying().(*types.Interface); ok {
					method := iterator.Map(iterator.Range(0, underIntf.NumMethods()), func(v int) fp.Tuple2[string, *types.Func] {
						m := underIntf.Method(v)
						return as.Tuple2(m.Name(), m)
					})
					return mutable.MapOf(iterator.ToGoMap(method))
				}
			}
			method := iterator.Map(iterator.Range(0, realtp.NumMethods()), func(v int) fp.Tuple2[string, *types.Func] {
				m := realtp.Method(v)
				return as.Tuple2(m.Name(), m)
			})
			return mutable.MapOf(iterator.ToGoMap(method))

		}()

		//TODO : type param 추가하면 난리남..
		if params.Size() > 0 && args.Size() == 0 {
			args = seq.Map(params, func(v TypeParam) TypeInfo {
				p := types.NewTypeParam(v.TypeName, v.Constraint)
				return typeInfo(p)
			})
		}
		return TypeInfo{
			ID:        id,
			Pkg:       realtp.Obj().Pkg(),
			Type:      tpe,
			TypeName:  realtp.Obj().Name(),
			TypeArgs:  args,
			TypeParam: params,
			Method:    methodMap,
		}
	case *types.Signature:
		params := typeParam(realtp.TypeParams())

		var pk *types.Package
		if realtp.Recv() != nil {
			pk = realtp.Recv().Pkg()
		}

		return TypeInfo{
			ID:        id,
			Pkg:       pk,
			Type:      tpe,
			TypeName:  "func",
			TypeParam: params,
		}
	case *types.Array:
		return TypeInfo{
			ID:       id,
			Type:     tpe,
			TypeName: "[_]",
			TypeArgs: []TypeInfo{typeInfo(realtp.Elem())},
		}
	case *types.Map:

		return TypeInfo{
			ID:       id,
			Type:     tpe,
			TypeName: "map",
			TypeArgs: []TypeInfo{typeInfo(realtp.Key()), typeInfo(realtp.Elem())},
		}
	case *types.Slice:
		elemTp := typeInfo(realtp.Elem())

		//fmt.Printf("slice elemTp = %s, istypeParam = %t\n", elemTp, elemTp.IsTypeParam())

		return TypeInfo{
			ID:       id,
			Type:     tpe,
			TypeName: "[]",
			TypeArgs: []TypeInfo{elemTp},
		}
	case *types.Pointer:
		elemTp := typeInfo(realtp.Elem())
		return TypeInfo{
			ID:       id,
			Type:     tpe,
			TypeName: "*",
			TypeArgs: []TypeInfo{elemTp},
		}
	}

	return TypeInfo{
		ID:       id,
		Type:     tpe,
		TypeName: tpe.String(),
	}
}

type findPosVisitor struct {
	pos   token.Pos
	found ast.Node
}

func (r *findPosVisitor) Visit(node ast.Node) (w ast.Visitor) {

	if node != nil && node.Pos() == r.pos {
		r.found = node
		return nil
	}

	return r
}
func FindNode(pk *packages.Package, pos token.Pos) ast.Node {
	for _, f := range pk.Syntax {
		if pos >= f.Pos() && pos <= f.End() {
			for _, d := range f.Decls {
				if pos >= d.Pos() && pos <= d.End() {

					v := &findPosVisitor{pos: pos}
					ast.Walk(v, d)
					return v.found

				}
			}
		}
	}
	return nil
}
