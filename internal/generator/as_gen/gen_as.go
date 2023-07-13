package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {

	genfp.Generate("as", "func_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
				func Func%d[%s,R any]( f func(%s) R) fp.Func%d[%s,R] {
					return fp.Func%d[%s,R](f)
				}
				`, i, genfp.FuncTypeArgs(1, i), genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, `
				func Supplier%d[%s,R any]( f func(%s) R, %s) func() R {
					return func() R {
						return f(%s)
					}
				}
				`, i, genfp.FuncTypeArgs(1, i), genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i),
				genfp.FuncCallArgs(1, i),
			)
		}

		fmt.Fprintf(f, `
func Curried2[A1, A2, R any](f func(A1,A2) R) fp.Func1[A1, fp.Func1[A2, R]] {
	return func(a1 A1) fp.Func1[A2, R] {
		return func(a2 A2) R {
			return f(a1, a2)
		}
	}
}		
		`)

		for i := 3; i < max.Func; i++ {

			fmt.Fprintf(f, `
func Curried%d[%s,R any]( f %s) %s {
	return func(a1 A1) %s {
		return Curried%d( func(%s) R {
			return f(%s)
		} )
	}	
}
`, i, genfp.FuncTypeArgs(1, i), genfp.FuncDecl("A", 1, i, "R"), genfp.CurriedType(1, i, "R"), genfp.CurriedType(2, i, "R"), i-1, genfp.FuncDeclArgs(2, i), genfp.FuncCallArgs(1, i))

		}

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, `
func UnTupled%d[%s,R any]( f func(fp.Tuple%d[%s]) R) %s {
	return func(%s) R {
		return f(Tuple%d(%s))
	}
}
`, i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i), genfp.FuncDecl("A", 1, i, "R"),
				genfp.FuncDeclArgs(1, i), i, genfp.FuncCallArgs(1, i))

		}

	})

	genfp.Generate("as", "tuple_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		for i := 1; i < max.Product; i++ {

			fmt.Fprintf(f, "func Tuple%d [%s any]( %s ) fp.Tuple%d[%s] { ", i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), i, genfp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, "  return fp.Tuple%d[%s] {\n", i, genfp.FuncTypeArgs(1, i))
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
			fmt.Fprintf(f, "func HList%d [%s any]( tuple fp.Tuple%d[%s]) %s { ", i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i), genfp.ConsType(1, i, "hlist.Nil"))

			fmt.Fprintf(f, `
				return hlist.Concat( tuple.Head(), hlist.Of%d( tuple.Tail() ))
			}
			`, i-1)

		}

	})

	genfp.Generate("as", "labelled_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		for i := 1; i < max.Product; i++ {

			fmt.Fprintf(f, "func Labelled%d [%s fp.Named]( %s ) fp.Labelled%d[%s] { ", i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), i, genfp.FuncTypeArgs(1, i))

			fmt.Fprintf(f, "  return fp.Labelled%d[%s] {\n", i, genfp.FuncTypeArgs(1, i))
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d: a%d,\n", j, j)
			}
			fmt.Fprintf(f, `}
		}
`)
		}

		fmt.Fprintf(f, `
					func HList1Labelled[A1 fp.Named](tuple fp.Labelled1[A1]) hlist.Cons[A1, hlist.Nil] {
						return hlist.Concat(tuple.Head(), hlist.Empty())
					}
		`)

		for i := 2; i < max.Product; i++ {
			fmt.Fprintf(f, "func HList%dLabelled [%s fp.Named]( tuple fp.Labelled%d[%s]) %s { ", i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i),
				genfp.ConsType(1, i, "hlist.Nil"))

			fmt.Fprintf(f, `
						return hlist.Of%d( tuple.Unapply() )
					}
					`, i)

		}

	})

}
