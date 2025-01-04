package testmeta_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/csgura/fp/genfp/generator"
	"github.com/csgura/fp/should"
	"golang.org/x/tools/go/packages"
)

func TestFindVar(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	v := generator.FindTaggedNotInitalizedVariable(pkgs, "@test.Summon")
	should.BeTrue(t, len(v) > 0)
	fmt.Printf("name = %s, tpe = %s\n", v[0].Name, v[0].Type)
}
