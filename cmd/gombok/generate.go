package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/metafp"
	"golang.org/x/tools/go/packages"
)

func genGenerate() {
	pack := os.Getenv("GOPACKAGE")

	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	genseq := metafp.FindTaggedCompositeVariable(pkgs, metafp.PackagedName{Package: "github.com/csgura/fp/genfp", Name: "GenerateFromUntil"}, "@fp.Generate")

	genseq.Foreach(func(cl *ast.CompositeLit) {
		gfu, err := genfp.ParseGenerateFromUntil(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s", err)
		} else {
			genfp.Generate(pack, gfu.File, func(w genfp.Writer) {
				for _, im := range gfu.Imports {
					w.GetImportedName(types.NewPackage(im.Package, im.Name))
				}

				w.Iteration(gfu.From, gfu.Until).Write(gfu.Template, map[string]any{})
			})
		}
	})
}
