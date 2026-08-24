package main

import (
	"fmt"
	"go/token"
	"go/types"
	"os"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/genfp/generator"
	"golang.org/x/tools/go/packages"
)

func isPredefined(pkgs []*packages.Package, pos token.Pos, willGenerated map[string]bool) bool {
	for _, p := range pkgs {
		file := p.Fset.Position(pos)
		if file.Filename != "" {
			if willGenerated[file.Filename] {
				return false
			}
		}
	}
	return true
}
func listMethods(p []*packages.Package, v types.Type, fileSet map[string]bool) map[string]bool {

	ret := map[string]bool{}

	switch tp := v.(type) {
	case *types.Named:
		for m := range tp.Methods() {
			if isPredefined(p, m.Pos(), fileSet) {
				ret[m.Name()] = true
			}
		}
	}

	return ret

}

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
	genmonadmethods := generator.FindGenerateMonadMethods(pkgs, "@internal.Generate")

	gentraverse := generator.FindGenerateTraverseFunctions(pkgs, "@internal.Generate")
	monadt := generator.FindGenerateMonadTransfomers(pkgs, "@internal.Generate")
	applicatives := generator.FindGenerateApplicatives(pkgs, "@internal.Generate")

	fileSet := map[string]bool{}
	for file := range genmonad {
		fullpath := cwd + "/" + file
		fileSet[fullpath] = true
	}

	for file := range genmonadmethods {
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
	methodList := map[string]map[string]bool{}

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

	for file, list := range genmonadmethods {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generated := methodList[gfu.TargetType.Obj().Name()]
				if generated == nil {
					targetType := gfu.TargetType.Obj().Type()
					generated = listMethods(pkgs, targetType, fileSet)
				}
				generator.WriteMonadMethods(w, gfu, generated)
				methodList[gfu.TargetType.Obj().Name()] = generated
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
