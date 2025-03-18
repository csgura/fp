package main

import (
	"fmt"
	"go/types"
	"os"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/genfp/generator"
	"golang.org/x/tools/go/packages"
)

func main() {
	pack := os.Getenv("GOPACKAGE")

	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	genmonad := generator.FindGenerateMonadFunctions(pkgs, "@internal.Generate")
	gentraverse := generator.FindGenerateTraverseFunctions(pkgs, "@internal.Generate")
	monadt := generator.FindGenerateMonadTransfomers(pkgs, "@internal.Generate")
	applicatives := generator.FindGenerateApplicatives(pkgs, "@internal.Generate")

	fileSet := map[string]bool{}
	for file := range genmonad {
		fullpath := cwd + "/" + file
		fileSet[fullpath] = true
	}
	for file := range gentraverse {
		fullpath := cwd + "/" + file
		fileSet[fullpath] = true

	}
	for file := range monadt {
		fullpath := cwd + "/" + file
		fileSet[fullpath] = true

	}

	for file := range applicatives {
		fullpath := cwd + "/" + file
		fileSet[fullpath] = true

	}

	funcList := map[string]bool{}

	for _, p := range pkgs {
		s := p.Types.Scope()
		for _, n := range s.Names() {
			o := s.Lookup(n)
			if _, ok := o.Type().(*types.Signature); ok {
				file := p.Fset.Position(o.Pos()).Filename
				if !fileSet[file] {
					funcList[o.Name()] = true
				}

			}
		}
	}

	for file, list := range monadt {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteMonadTransformers(w, gfu, funcList)
			}
		})
	}

	for file, list := range genmonad {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteMonadFunctions(w, gfu, funcList)
			}
		})
	}

	for file, list := range gentraverse {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteTraverseFunctions(w, gfu, funcList)
			}
		})
	}

	for file, list := range applicatives {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteApplicativeFunctions(w, gfu, funcList)
			}
		})
	}

}
