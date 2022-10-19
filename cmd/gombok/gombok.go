package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"golang.org/x/tools/go/packages"
)

type TaggedType struct {
	Package  *types.Package
	TypeSpec *ast.TypeSpec
	Struct   *types.Struct
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
							return seq.Of(TaggedType{
								Package:  l.Pkg(),
								TypeSpec: ts,
								Struct:   st,
							})
						}

					}

				}
				return seq.Of[TaggedType]()
			})
		})
	})

}

func typeName(pk *types.Package, tpe types.Type) string {
	ftp := tpe.String()
	if namedtp, ok := tpe.(*types.Named); ok {
		if namedtp.Obj().Pkg().Path() == pk.Path() {
			ftp = namedtp.Origin().Obj().Name()
		} else {
			if namedtp.TypeArgs() != nil {
				args := iterator.Map(iterator.Range(0, namedtp.TypeArgs().Len()), func(i int) string {
					return typeName(pk, namedtp.TypeArgs().At(i))
				}).MakeString(",")
				ftp = fmt.Sprintf("%s.%s[%s]", namedtp.Obj().Pkg().Name(), namedtp.Obj().Name(), args)
			} else {
				ftp = fmt.Sprintf("%s.%s", namedtp.Obj().Pkg().Name(), namedtp.Obj().Name())

			}
		}
	}
	return ftp
}

func main() {
	file := os.Getenv("GOFILE")
	line := os.Getenv("GOLINE")
	pack := os.Getenv("GOPACKAGE")

	cwd, _ := os.Getwd()

	fmt.Printf("cwd = %s , pack = %s file = %s, line = %s\n", try.Apply(os.Getwd()), pack, file, line)

	//packages.LoadFiles()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	st := findValueStruct(pkgs)
	st.Foreach(func(v TaggedType) {
		fmt.Println("generate value for", v.TypeSpec.Name)
		for i := 0; i < v.Struct.NumFields(); i++ {
			f := v.Struct.Field(i)
			ftp := typeName(v.Package, f.Type())
			fmt.Printf(`
				func (r %s) %s() %s {
					return r.%s
				}
			`, v.TypeSpec.Name, strings.ToUpper(f.Name()[:1])+f.Name()[1:], ftp, f.Name())

			fmt.Printf(`
				func (r %s) With%s(v %s) %s {
					r.%s = v
					return r
				}
			`, v.TypeSpec.Name, strings.ToUpper(f.Name()[:1])+f.Name()[1:], ftp, v.TypeSpec.Name, f.Name())
		}

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
