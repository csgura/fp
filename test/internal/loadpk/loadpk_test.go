package loadpk_test

import (
	"fmt"
	"go/types"
	"testing"

	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"golang.org/x/tools/go/packages"
)

func TestLoad(t *testing.T) {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, "github.com/csgura/fp/test/internal/loadpk")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("len pkgs = %d\n", len(pkgs))

	as.Seq(pkgs[0].Types.Scope().Names()).Foreach(func(v string) {
		fmt.Printf("name %s\n", v)
	})

	h := pkgs[0].Types.Scope().Lookup("Hello")
	fmt.Printf("Hello = %s, %T\n", h.String(), h)
	if c, ok := h.(*types.Const); ok {
		fmt.Printf("const value = %s\n", c.Val())
	}

	res := genfp.FindGenerateFromUntil(pkgs, "@ConstTest")
	for _, list := range res {
		for _, v := range list {
			fmt.Printf("File = %s\n", v.File)
			fmt.Printf("from = %d, until = %d\n", v.From, v.Until)
		}
	}
}
