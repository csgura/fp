package metafp

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
	"unicode"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type TaggedStruct struct {
	Package *types.Package
	Name    string
	Scope   *types.Scope
	Struct  *types.Struct
	Fields  fp.Seq[StructField]
	Info    TypeInfo
	Tags    mutable.Set[string]
}

func (r TaggedStruct) PackagedName(w ImportSet, workingPackage *types.Package) string {
	if workingPackage.Path() == r.Package.Path() {
		return r.Name
	}

	pk := w.GetImportedName(r.Package)
	return fmt.Sprintf("%s.%s", pk, r.Name)
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
				Name:     f.Name(),
				Type:     tn,
				Tag:      st.Tag(i),
				Embedded: f.Embedded(),
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

func extractTag(comment string) mutable.Set[string] {
	list := as.Seq(strings.Fields(comment))
	return seq.ToGoSet(list.Filter(as.Func2(strings.Contains).ApplyLast("@")))
}

func GetTagsOfType(p []*packages.Package, name string) mutable.Set[string] {
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
					doc := option.Map(option.Of(ts.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)
					return doc.ToSeq()
				}
				return seq.Of[string]()
			})
		})
	}).Head()

	return option.Map(comment, extractTag).OrZero()
}

func FindTaggedStruct(p []*packages.Package, tag string) fp.Seq[TaggedStruct] {

	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[TaggedStruct] {
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
					doc := option.Map(option.Of(ts.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)
					if doc.Filter(as.Func2(strings.Contains).ApplyLast(tag)).IsDefined() {
						return option.Map(LookupStruct(pk.Types, ts.Name.Name), func(v TaggedStruct) TaggedStruct {
							v.Tags = option.Map(doc, extractTag).OrZero()
							return v
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

type TypeInfo struct {
	Pkg       *types.Package
	Type      types.Type
	TypeArgs  fp.Seq[TypeInfo]
	TypeParam fp.Seq[TypeParam]
	Method    fp.Map[string, *types.Func]
}

func isSamePkg(p1 *types.Package, p2 *types.Package) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}

	return p1.Path() == p2.Path()
}

func (r TypeInfo) IsInstantiatedOf(typeParam fp.Seq[TypeParam], hasTypeParam TypeInfo) bool {

	if !isSamePkg(r.Pkg, hasTypeParam.Pkg) {
		return false
	}

	if r.Name().OrZero() != hasTypeParam.Name().OrZero() {
		return false
	}

	if r.TypeArgs.Size() != hasTypeParam.TypeArgs.Size() {
		return false
	}

	return seq.Zip(r.TypeArgs, hasTypeParam.TypeArgs).ForAll(func(t fp.Tuple2[TypeInfo, TypeInfo]) bool {
		if t.I2.IsTypeParam() {
			return ConstraintCheck(typeParam, hasTypeParam, r)
		}
		return t.I1.IsInstantiatedOf(typeParam, t.I2)
	})
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
			return fmt.Sprintf("%s.%s%s%v", r.PkgName(), name, r.TypeParam, r.TypeArgs)
		}
		return fmt.Sprintf("%s.%s", r.PkgName(), name)
	}
	if r.TypeParam.Size() > 0 {
		return fmt.Sprintf("%s%s%v", name, r.TypeParam, r.TypeArgs)
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

func (r TypeInfo) TypeParamDecl(w ImportSet, cwd *types.Package) string {
	return iterator.Map(r.TypeParam.Iterator(), func(v TypeParam) string {
		tn := w.TypeName(cwd, v.Constraint)
		return fmt.Sprintf("%s %s", v.Name, tn)
	}).MakeString(",")
}

func (r TypeInfo) TypeParamIns(w ImportSet, cwd *types.Package) string {
	return iterator.Map(r.TypeParam.Iterator(), func(v TypeParam) string {
		return v.Name
	}).MakeString(",")
}

type StructField struct {
	Name     string
	Type     TypeInfo
	Tag      string
	Embedded bool
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
	switch realtp := tpe.(type) {
	case *types.TypeParam:
		return TypeInfo{
			Type: tpe,
		}
	case *types.Named:
		args := typeArgs(realtp.TypeArgs())
		params := typeParam(realtp.TypeParams())

		method := iterator.Map(iterator.Range(0, realtp.NumMethods()), func(v int) fp.Tuple2[string, *types.Func] {
			m := realtp.Method(v)
			return as.Tuple2(m.Name(), m)
		})

		return TypeInfo{
			Pkg:       realtp.Obj().Pkg(),
			Type:      tpe,
			TypeArgs:  args,
			TypeParam: params,
			Method:    mutable.MapOf(iterator.ToGoMap(method)),
		}
	case *types.Signature:
		params := typeParam(realtp.TypeParams())

		var pk *types.Package
		if realtp.Recv() != nil {
			pk = realtp.Recv().Pkg()
		}

		return TypeInfo{
			Pkg:       pk,
			Type:      tpe,
			TypeParam: params,
		}
	case *types.Array:
		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(realtp.Elem())},
		}
	case *types.Map:

		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(realtp.Key()), typeInfo(realtp.Elem())},
		}
	case *types.Slice:
		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(realtp.Elem())},
		}
	}

	return TypeInfo{
		Type: tpe,
	}
}
