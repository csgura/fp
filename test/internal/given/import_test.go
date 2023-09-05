package given_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/csgura/fp/metafp"
	"golang.org/x/tools/go/packages"
)

func TestImport(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	t1Type := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("t1").Type())
	fmt.Printf("typeArgs = %v\n", t1Type.TypeArgs)
	metafp.FindTypeClassImport(pkgs).Foreach(func(v metafp.TypeClassDirective) {
		index := metafp.LoadTypeClassInstance(v.PrimitiveInstancePkg, v.TypeClass)

		res := index.Find(t1Type)
		fmt.Printf("len res = %d\n", res.Size())
		res.Foreach(func(v metafp.TypeClassInstance) {
			fmt.Println(v.Name)
		})
	})
}
