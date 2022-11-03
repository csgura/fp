package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/metafp"
)

func main() {
	metafp.Generate("product", "tuple_gen.go", func(f metafp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		for i := 3; i < max.Product; i++ {

			fmt.Fprintf(f, "func Tuple%d [%s any]( %s ) fp.Tuple%d[%s] { ", i, metafp.FuncTypeArgs(1, i), metafp.FuncDeclArgs(1, i), i, metafp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, "  return fp.Tuple%d[%s] {\n", i, metafp.FuncTypeArgs(1, i))
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d: a%d,\n", j, j)
			}
			fmt.Fprintf(f, `}
		}
		`)

		}

		for i := 2; i < max.Product; i++ {

			fmt.Fprintf(f, "func TupleFromHList%d [%s any]( list %s ) fp.Tuple%d[%s] { ", i, metafp.FuncTypeArgs(1, i), metafp.ConsType(1, i, "hlist.Nil"), i, metafp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, `
				tail := TupleFromHList%d( list.Tail() )
				return Tuple%d( list.Head(), %s)
			}
			`, i-1, i, metafp.FuncCallArgs(1, i-1, "tail.I"))

		}

		for i := 2; i < max.Product; i++ {

			fmt.Fprintf(f, "func LabelledFromHList%d [%s fp.Named]( list %s ) fp.Labelled%d[%s] { ", i, metafp.FuncTypeArgs(1, i), metafp.ConsType(1, i, "hlist.Nil"), i, metafp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, `
				tail := LabelledFromHList%d( list.Tail() )
				return as.Labelled%d( list.Head(), %s)
			}
			`, i-1, i, metafp.FuncCallArgs(1, i-1, "tail.I"))

		}

		for i := 2; i < max.Func; i++ {

			fmt.Fprintf(f, "func Lift%d [%s , R any](f func(%s) R) fp.Func1[fp.Tuple%d[%s],R] { ", i, metafp.FuncTypeArgs(1, i), metafp.FuncDeclArgs(1, i), i, metafp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, `
	return func(t fp.Tuple%d[%s]) R {
					return f(t.Unapply())
				}
}

`, i, metafp.FuncTypeArgs(1, i))

		}
	})

}
