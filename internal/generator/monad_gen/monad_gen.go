package main

import (
	"fmt"
	"os"

	"github.com/csgura/fp/genfp"
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

	genseq := genfp.FindGenerateMonadFunctions(pkgs, "@internal.Generate")
	for file, list := range genseq {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				genfp.WriteMonadFunctions(w, gfu)
			}
		})
	}

	genseq = genfp.FindGenerateTraverseFunctions(pkgs, "@internal.Generate")
	for file, list := range genseq {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				genfp.WriteTraverseFunctions(w, gfu)
			}
		})
	}

	monadt := genfp.FindGenerateMonadTransfomers(pkgs, "@internal.Generate")
	for file, list := range monadt {

		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				fmt.Printf("target type = %s, monad type = %s, name = %s\n", gfu.TargetType, gfu.MonadType, gfu.Name)
				genfp.WriteMonadTransformers(w, gfu)
			}
		})
	}
}
