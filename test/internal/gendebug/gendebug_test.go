package gendebug_test

import (
	"os"
	"testing"

	"github.com/csgura/fp/should"
	"golang.org/x/tools/go/packages"
)

func TestShow(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	_, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)
}
