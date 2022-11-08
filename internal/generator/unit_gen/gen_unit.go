package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {

	genfp.Generate("unit", "func_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Func%d[%s any]( f func(%s) ) fp.Func%d[%s,fp.Unit] {
	return func(%s) fp.Unit {
		f(%s)
		return fp.Unit{}
	}
}
`, i, genfp.FuncTypeArgs(1, i), genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), genfp.FuncCallArgs(1, i))

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
