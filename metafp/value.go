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

func LookupStruct(pk *types.Package, name string) fp.Option[TaggedStruct] {
	l := pk.Scope().Lookup(name)
	if st, ok := l.Type().Underlying().(*types.Struct); ok {
		fl := iterator.Map(iterator.Range(0, st.NumFields()), func(i int) StructField {
			f := st.Field(i)
			tn := typeInfo(l.Pkg(), f.Type())
			return StructField{
				Name:     f.Name(),
				Type:     tn,
				Tag:      st.Tag(i),
				Embedded: f.Embedded(),
			}
		}).ToSeq()

		info := typeInfo(l.Pkg(), l.Type())

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
}

type TypeInfo struct {
	Pkg       *types.Package
	Type      types.Type
	TypeArgs  fp.Seq[TypeInfo]
	TypeParam fp.Seq[TypeParam]
	Method    fp.Map[string, *types.Func]
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

func typeInfo(pk *types.Package, tpe types.Type) TypeInfo {
	switch realtp := tpe.(type) {
	case *types.TypeParam:
		return TypeInfo{
			Type: tpe,
		}
	case *types.Named:
		args := fp.Seq[TypeInfo]{}
		params := fp.Seq[TypeParam]{}

		if realtp.TypeArgs() != nil {
			args = iterator.Map(iterator.Range(0, realtp.TypeArgs().Len()), func(i int) TypeInfo {
				return typeInfo(pk, realtp.TypeArgs().At(i))
			}).ToSeq()

		}

		if realtp.TypeParams() != nil {
			params = iterator.Map(iterator.Range(0, realtp.TypeParams().Len()), func(i int) TypeParam {
				return TypeParam{
					Name:       realtp.TypeParams().At(i).Obj().Name(),
					Constraint: realtp.TypeParams().At(i).Constraint(),
				}
			}).ToSeq()
		}

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
	case *types.Array:
		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(pk, realtp.Elem())},
		}
	case *types.Map:

		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(pk, realtp.Key()), typeInfo(pk, realtp.Elem())},
		}
	case *types.Slice:
		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(pk, realtp.Elem())},
		}
	}

	return TypeInfo{
		Type: tpe,
	}
}
