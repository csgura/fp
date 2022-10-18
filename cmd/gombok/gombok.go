package main

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"golang.org/x/tools/go/packages"
)

func findValueStruct(p []*packages.Package) fp.Seq[*ast.TypeSpec] {
	s := seq.FlatMap(p, func(v *packages.Package) fp.Seq[*ast.File] {
		return v.Syntax
	})

	s2 := seq.FlatMap(s, func(v *ast.File) fp.Seq[ast.Decl] {
		return v.Decls
	})

	s3 := seq.FlatMap(s2, func(v ast.Decl) fp.Seq[*ast.GenDecl] {
		switch r := v.(type) {
		case *ast.GenDecl:
			return seq.Of(r)
		}
		return seq.Of[*ast.GenDecl]()
	})

	return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[*ast.TypeSpec] {
		gdDoc := option.Of(gd.Doc)

		return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[*ast.TypeSpec] {
			if ts, ok := v.(*ast.TypeSpec); ok {
				doc := option.Map(option.Of(ts.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)

				if doc.Filter(as.Func2(strings.Contains).ApplyLast("@fp.Value")).IsDefined() {
					return seq.Of(ts)
				}

			}
			return seq.Of[*ast.TypeSpec]()
		})
	})
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

	findValueStruct(pkgs).Foreach(func(v *ast.TypeSpec) {
		fmt.Println("generate value for", v.Name)
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
