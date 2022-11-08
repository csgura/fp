package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {
	genfp.Generate("curried", "curried_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		for i := 2; i < max.Func; i++ {

			args := ""

			for j := 1; j <= i; j++ {
				if j != 1 {
					args = args + ","
				}
				args = args + fmt.Sprintf("A%d", j)
			}
			fmt.Fprintf(f, "func Func%d [%s, R any]( f func(%s) R ) %s { ", i, args, args, genfp.CurriedType(1, i, "R"))

			fmt.Fprintf(f, `
	return func(a1 A1) %s {
		return Func%d( func (%s) R {
			return f(%s)
		})
	}	
}	
`, genfp.CurriedType(2, i, "R"), i-1, genfp.FuncDeclArgs(2, i), genfp.FuncCallArgs(1, i))

			fmt.Fprintf(f, "func Revert%d [%s, R any]( f %s ) fp.Func%d[%s,R] { ", i, args, genfp.CurriedType(1, i, "R"), i, args)

			fmt.Fprintf(f, `
	return func(%s) R {
		return f%s
	}	
}	
`, genfp.FuncDeclArgs(1, i), genfp.CurriedCallArgs(1, i))

		}
		for i := 3; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Flip%d[%s,R any](f  %s) %s {
	return Func%d(
		func(%s, a1 A1) R {
			return f%s
		}, 
	)
}
`, i-1, genfp.FuncTypeArgs(1, i), genfp.CurriedType(1, i, "R"), genfp.CurriedType(2, i, "fp.Func1[A1, R]"),
				i, genfp.FuncDeclArgs(2, i), genfp.CurriedCallArgs(1, i),
			)

			fmt.Fprintf(f, `
			func Compose%d[%s,GA,GR any](f %s, g fp.Func1[GA,GR]) %s {
				return func(a1 A1) %s  {
					return Compose%d(f(a1), g)
				}
			}
		`, i, genfp.FuncTypeArgs(1, i), genfp.CurriedType(1, i, "GA"), genfp.CurriedType(1, i, "GR"),
				genfp.CurriedType(2, i, "GR"),
				i-1,
			)
		}

	})

}
