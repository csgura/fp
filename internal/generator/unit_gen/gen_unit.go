package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/metafp"
)

func main() {

	metafp.Generate("unit", "func_gen.go", func(f metafp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Func%d[%s any]( f func(%s) ) fp.Func%d[%s,fp.Unit] {
	return func(%s) fp.Unit {
		f(%s)
		return fp.Unit{}
	}
}
`, i, metafp.FuncTypeArgs(1, i), metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i), metafp.FuncDeclArgs(1, i), metafp.FuncCallArgs(1, i))

		}

	})

	// for j := 1; j <= i; j++ {
	// 	if j != 1 {
	// 		fmt.Fprintf(f, ",")
	// 	}
	// 	fmt.Fprintf(f, "a%d A%d", j, j)
	// }
	// fmt.Fprintf(f, ") R\n\n")

}
