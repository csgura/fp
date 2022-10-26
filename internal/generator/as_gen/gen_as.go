package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/metafp"
)

func main() {

	metafp.Generate("as", "func_gen.go", func(f metafp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Func%d[%s,R any]( f func(%s) R) fp.Func%d[%s,R] {
	return fp.Func%d[%s,R](f)
}
`, i, metafp.FuncTypeArgs(1, i), metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i))

		}

		fmt.Fprintf(f, `
func Curried2[A1, A2, R any](f fp.Func2[A1,A2,R]) fp.Func1[A1, fp.Func1[A2, R]] {
	return func(a1 A1) fp.Func1[A2, R] {
		return func(a2 A2) R {
			return f(a1, a2)
		}
	}
}		
		`)

		for i := 3; i < max.Func; i++ {

			fmt.Fprintf(f, `
func Curried%d[%s,R any]( f fp.Func%d[%s, R]) %s {
	return func(a1 A1) %s {
		return Curried%d( func(%s) R {
			return f(%s)
		} )
	}	
}
`, i, metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i), metafp.CurriedType(1, i, "R"), metafp.CurriedType(2, i, "R"), i-1, metafp.FuncDeclArgs(2, i), metafp.FuncCallArgs(1, i))

		}

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, `
func UnTupled%d[%s,R any]( f func(fp.Tuple%d[%s]) R) fp.Func%d[%s,R] {
	return func(%s) R {
		return f(Tuple%d(%s))
	}
}
`, i, metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i), metafp.FuncDeclArgs(1, i), i, metafp.FuncCallArgs(1, i))

		}

	})

	metafp.Generate("as", "tuple_gen.go", func(f metafp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		for i := 1; i < max.Product; i++ {

			fmt.Fprintf(f, "func Tuple%d [%s any]( %s ) fp.Tuple%d[%s] { ", i, metafp.FuncTypeArgs(1, i), metafp.FuncDeclArgs(1, i), i, metafp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, "  return fp.Tuple%d[%s] {\n", i, metafp.FuncTypeArgs(1, i))
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d: a%d,\n", j, j)
			}
			fmt.Fprintf(f, `}
		}
`)
		}

		fmt.Fprintf(f, `
			func HList1[A1 any](tuple fp.Tuple1[A1]) hlist.Cons[A1, hlist.Nil] {
				return hlist.Concat(tuple.Head(), hlist.Empty())
			}		
`)

		for i := 2; i < max.Product; i++ {
			fmt.Fprintf(f, "func HList%d [%s any]( tuple fp.Tuple%d[%s]) %s { ", i, metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i), metafp.ConsType(1, i, "hlist.Nil"))

			fmt.Fprintf(f, `
				return hlist.Concat( tuple.Head(), HList%d( tuple.Tail() ))
			}
			`, i-1)

		}

	})

	metafp.Generate("as", "labelled_gen.go", func(f metafp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		for i := 1; i < max.Product; i++ {

			fmt.Fprintf(f, "func Labelled%d [%s any]( %s ) fp.Labelled%d[%s] { ", i, metafp.FuncTypeArgs(1, i), metafp.FuncDeclTypeClassArgs(1, i, "fp.Field"), i, metafp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, "  return fp.Labelled%d[%s] {\n", i, metafp.FuncTypeArgs(1, i))
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d: ins%d,\n", j, j)
			}
			fmt.Fprintf(f, `}
		}
`)
		}

		fmt.Fprintf(f, `
					func HList1Labelled[A1 any](tuple fp.Labelled1[A1]) hlist.Cons[fp.Field[A1], hlist.Nil] {
						return hlist.Concat(tuple.Head(), hlist.Empty())
					}
		`)

		for i := 2; i < max.Product; i++ {
			fmt.Fprintf(f, "func HList%dLabelled [%s any]( tuple fp.Labelled%d[%s]) %s { ", i, metafp.FuncTypeArgs(1, i), i, metafp.FuncTypeArgs(1, i),
				metafp.Monad("fp.Field").ConsType(1, i, "hlist.Nil"))

			fmt.Fprintf(f, `
						return hlist.Concat( tuple.Head(), HList%dLabelled( tuple.Tail() ))
					}
					`, i-1)

		}

	})

}
