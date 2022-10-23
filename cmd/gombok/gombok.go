package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"strings"
	"unicode"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/generator/common"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"

	// . "github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

type StructField struct {
	Name string
	Type TypeInfo
}

func (r StructField) Public() bool {
	return unicode.IsUpper([]rune(r.Name)[0])
}

type TypeInfo struct {
	Pkg      *types.Package
	Type     types.Type
	TypeArgs fp.Seq[TypeInfo]
}
type TaggedType struct {
	Name     string
	Package  *types.Package
	TypeSpec *ast.TypeSpec
	Struct   *types.Struct
	Fields   fp.Seq[StructField]
}

type TypeClassDerive struct {
	Package   *types.Package
	Generator *types.Package
	TypeClass types.Type
	DeriveFor types.Type
}

func findTypeClassDerive(p []*packages.Package) fp.Seq[TypeClassDerive] {
	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[TypeClassDerive] {
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

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[TypeClassDerive] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[TypeClassDerive] {
				if vs, ok := v.(*ast.ValueSpec); ok {
					doc := option.Map(option.Of(vs.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)
					if doc.Filter(as.Func2(strings.Contains).ApplyLast("@fp.Derive")).IsDefined() {

						info := &types.Info{
							Types: make(map[ast.Expr]types.TypeAndValue),
						}
						types.CheckExpr(pk.Fset, pk.Types, v.Pos(), vs.Type, info)
						ti := info.Types[vs.Type]
						if nt, ok := ti.Type.(*types.Named); ok && nt.TypeArgs().Len() == 1 {
							if tt, ok := nt.TypeArgs().At(0).(*types.Named); ok && tt.TypeArgs().Len() == 1 {
								return seq.Of(TypeClassDerive{
									Package:   pk.Types,
									Generator: nt.Obj().Pkg(),
									TypeClass: tt.Obj().Type(),
									DeriveFor: tt.TypeArgs().At(0),
								})

							}

						}
					}
				}
				return seq.Of[TypeClassDerive]()
			})
		})
	})
}
func findValueStruct(p []*packages.Package) fp.Seq[TaggedType] {

	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[TaggedType] {
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

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[TaggedType] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[TaggedType] {

				if ts, ok := v.(*ast.TypeSpec); ok {
					doc := option.Map(option.Of(ts.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)

					if doc.Filter(as.Func2(strings.Contains).ApplyLast("@fp.Value")).IsDefined() {

						l := pk.Types.Scope().Lookup(ts.Name.Name)
						if st, ok := l.Type().Underlying().(*types.Struct); ok {
							fl := iterator.Map(iterator.Range(0, st.NumFields()), func(i int) StructField {
								f := st.Field(i)
								tn := typeInfo(l.Pkg(), f.Type())
								return StructField{
									Name: f.Name(),
									Type: tn,
								}
							}).ToSeq()

							return seq.Of(TaggedType{
								Name:     ts.Name.Name,
								Package:  l.Pkg(),
								TypeSpec: ts,
								Struct:   st,
								Fields:   fl,
							})
						}
					}

				}
				return seq.Of[TaggedType]()
			})
		})
	})

}

func typeInfo(pk *types.Package, tpe types.Type) TypeInfo {
	switch realtp := tpe.(type) {
	case *types.Named:
		args := fp.Seq[TypeInfo]{}

		if realtp.TypeArgs() != nil {
			args = iterator.Map(iterator.Range(0, realtp.TypeArgs().Len()), func(i int) TypeInfo {
				return typeInfo(pk, realtp.TypeArgs().At(i))
			}).ToSeq()

		}
		return TypeInfo{
			Pkg:      realtp.Obj().Pkg(),
			Type:     tpe,
			TypeArgs: args,
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

func main() {
	//	file := os.Getenv("GOFILE")
	//	line := os.Getenv("GOLINE")
	pack := os.Getenv("GOPACKAGE")

	cwd, _ := os.Getwd()

	//	fmt.Printf("cwd = %s , pack = %s file = %s, line = %s\n", try.Apply(os.Getwd()), pack, file, line)

	//packages.LoadFiles()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	// out := NewFile(pack)
	// out.Qual(path string, name string)

	common.Generate(pack, "value_generated.go", func(w common.Writer) {

		fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))
		asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		st := findValueStruct(pkgs)

		d := findTypeClassDerive(pkgs)
		d.Foreach(func(v TypeClassDerive) {
			l := v.Generator.Scope().Lookup("Option")
			fmt.Printf("lookup %s.Option = %v\n", v.Generator.Name(), l)
			fmt.Printf("derive %v for %v\n", v.TypeClass, v.DeriveFor)
		})

		st.Foreach(func(v TaggedType) {

			fmt.Fprintf(w, `
				type %sBuilder %s
			`, v.Name, v.Name)

			fmt.Fprintf(w, `
				func(r %sBuilder) Build() %s {
					return %s(r)
				}
			`, v.Name, v.Name, v.Name)

			fmt.Fprintf(w, `
				func(r %s) Builder() %sBuilder {
					return %sBuilder(r)
				}
			`, v.Name, v.Name, v.Name)

			privateFields := v.Fields.FilterNot(StructField.Public)
			privateFields.Foreach(func(f StructField) {

				uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]
				ftp := w.TypeName(pkgs[0].Types, f.Type.Type)

				fmt.Fprintf(w, `
						func (r %s) %s() %s {
							return r.%s
						}
					`, v.Name, uname, ftp, f.Name)

				fmt.Fprintf(w, `
						func (r %s) With%s(v %s) %s {
							r.%s = v
							return r
						}
					`, v.Name, uname, ftp, v.Name, f.Name)

				fmt.Fprintf(w, `
						func (r %sBuilder) %s( v %s) %sBuilder {
							r.%s = v
							return r
						}
					`, v.Name, uname, ftp, v.Name, f.Name)
			})

			fm := seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf("%s=%%v", f.Name)
			}).Iterator().MakeString(", ")

			fields := seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Iterator().MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, v.Name,
				fmtalias, v.Name, fm, fields,
			)

			tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v StructField) string {
				return w.TypeName(pkgs[0].Types, v.Type.Type)
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsTuple() fp.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, v.Name, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, StructField]) string {
				return fmt.Sprintf("r.%s = t.I%d", f.I2.Name, f.I1+1)
			}).Take(max.Product).MakeString("\n")

			fmt.Fprintf(w, `
					func (r %sBuilder) FromTuple(t fp.Tuple%d[%s] ) %sBuilder {
						%s
						return r
					}
				`, v.Name, privateFields.Size(), tp, v.Name,
				fields,
			)

			fields = seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf(`"%s" : r.%s`, f.Name, f.Name)
			}).Iterator().Take(max.Product).MakeString(",\n")

			fmt.Fprintf(w, `
					func(r %s) AsMap() map[string]any {
						return map[string]any {
							%s,
						}
					}

				`, v.Name,
				fields,
			)

			fields = iterator.Map(privateFields.Iterator(), func(f StructField) string {
				return fmt.Sprintf(`if v , ok := m["%s"].(%s); ok {
						r.%s = v
					}
						`, f.Name, w.TypeName(pkgs[0].Types, f.Type.Type),
					f.Name,
				)
			}).Take(max.Product).MakeString("\n")

			fmt.Fprintf(w, `
					func(r %sBuilder) FromMap(m map[string]any) %sBuilder {

						%s
						
						return r
					}

				`, v.Name, v.Name,
				fields,
			)

			tp = iterator.Map(privateFields.Iterator().Take(max.Product), func(v StructField) string {
				return fmt.Sprintf("fp.Tuple2[string,%s]", w.TypeName(pkgs[0].Types, v.Type.Type))
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf(`as.Tuple2("%s", r.%s)`, f.Name, f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsLabelled() fp.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, v.Name, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, StructField]) string {
				return fmt.Sprintf("r.%s = t.I%d.I2", f.I2.Name, f.I1+1)
			}).Take(max.Product).MakeString("\n")

			fmt.Fprintf(w, `
					func (r %sBuilder) FromLabelled(t fp.Tuple%d[%s] ) %sBuilder {
						%s
						return r
					}
				`, v.Name, privateFields.Size(), tp, v.Name,
				fields,
			)

		})

	})

	// seq.FromMapKeys(s.Scope.Objects).Foreach(func(k string) {
	// 	fmt.Println("key = ", k)
	// 	o := s.Scope.Objects[k]
	// 	if ts, ok := o.Decl.(*ast.TypeSpec); ok {
	// 		fmt.Printf("name = %s\n", ts.Name)
	// 		if ts.Doc != nil {
	// 			fmt.Printf("Doc = %v\n", ts.Doc.Text())
	// 		}
	// 		fmt.Printf("comment = %v\n", ts.Comment)

	// 	}
	// })

	// seq.Of(s.Decls).Foreach(func(v []ast.Decl) {

	// 	seq.Of(v...).Foreach(func(v ast.Decl) {
	// 		fmt.Println("decl = ", v)

	// 	})
	// })

	// t := pkgs[0].Types.Scope()
	// seq.Of(t.Names()...).Foreach(func(v string) {
	// 	fmt.Printf("name = %s\n", v)
	// 	c := t.Lookup(v)
	// 	c.Pos()
	// 	fmt.Println(c.String())
	// 	if st, ok := t.Lookup(v).Type().Underlying().(*types.Struct); ok {
	// 		iterator.Range(0, st.NumFields()).Foreach(func(i int) {
	// 			fmt.Printf("field name = %s : %s\n", st.Field(i).Name(), st.Field(i).Type())
	// 		})
	// 	}

	// })

}
