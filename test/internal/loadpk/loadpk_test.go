package loadpk_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/as"
	"golang.org/x/tools/go/packages"
)

func TestLoad(t *testing.T) {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, "github.com/csgura/fp/show")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("len pkgs = %d\n", len(pkgs))

	as.Seq(pkgs[0].Types.Scope().Names()).Foreach(func(v string) {
		fmt.Printf("name %s\n", v)
	})
}
