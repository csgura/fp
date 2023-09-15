package main

import (
	"fmt"
	"go/types"
	"os"

	"github.com/csgura/fp/genfp"
	"golang.org/x/tools/go/packages"
)

func genGenerate() {
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

	genseq := genfp.FindGenerateFromUntil(pkgs, "@fp.Generate")

	for file, list := range genseq {
		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range list {
				for _, im := range gfu.Imports {
					w.GetImportedName(types.NewPackage(im.Package, im.Name))
				}

				w.Iteration(gfu.From, gfu.Until).Write(gfu.Template, map[string]any{})
			}
		})
	}
}
