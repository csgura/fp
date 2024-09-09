package main

import (
	"fmt"
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

	genseq := generator.FindGenerateMonadFunctions(pkgs, "@internal.Generate")

	for file, list := range genseq {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteMonadFunctions(w, gfu)
			}
		})
	}

	genseq = generator.FindGenerateTraverseFunctions(pkgs, "@internal.Generate")
	for file, list := range genseq {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteTraverseFunctions(w, gfu)
			}
		})
	}

	monadt := generator.FindGenerateMonadTransfomers(pkgs, "@internal.Generate")
	for file, list := range monadt {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				generator.WriteMonadTransformers(w, gfu)
			}
		})
	}
}
