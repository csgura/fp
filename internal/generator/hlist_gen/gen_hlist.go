package main

import (
	"fmt"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func consType(start, until int, last string) string {
	ret := last
	for j := until; j >= start; j-- {
		ret = fmt.Sprintf("Cons[A%d, %s]", j, ret)
	}
	return ret
}

func reversConsType(start, until int) string {
	ret := "Nil"
	for j := start; j <= until; j++ {
		ret = fmt.Sprintf("Cons[A%d, %s]", j, ret)
	}
	return ret
}

func main() {

	genfp.Generate("hlist", "lift_gen.go", func(f genfp.Writer) {

		for i := 2; i < max.Func; i++ {

			fmt.Fprintf(f, "func Lift%d [%s, R any]( f func(%s) R ) func(%s) R { ", i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), consType(1, i, "Nil"))

			fmt.Fprintf(f, `
	return func(v %s) R {
		rf := Lift%d(func(%s) R {
			return f(v.Head(), %s)
		})

		return rf(v.Tail())
	}
}	
`, consType(1, i, "Nil"), i-1, genfp.FuncDeclArgs(1+1, i), genfp.FuncCallArgs(1+1, i))

			fmt.Fprintf(f, "func Rift%d [%s, R any]( f func(%s) R ) func(%s) R { ", i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), reversConsType(1, i))

			fmt.Fprintf(f, `
	return func(v %s) R {
		rf := Rift%d(func(%s) R {
			return f(%s, v.Head())
		})

		return rf(v.Tail())
	}
}	
`, reversConsType(1, i), i-1, genfp.FuncDeclArgs(1, i-1), genfp.FuncCallArgs(1, i-1))

		}
	})

	genfp.Generate("hlist", "case_gen.go", func(f genfp.Writer) {
		for i := 2; i < max.Product; i++ {

			fmt.Fprintf(f, "func Case%d [%s any, T HList, R any](hl %s,  f func(%s) R ) R { ", i, genfp.FuncTypeArgs(1, i), consType(1, i, "T"), genfp.FuncDeclArgs(1, i))

			fmt.Fprintf(f, `
	return Case%d(hl.Tail(), func(%s) R {
		return f(hl.Head(), %s)
	})
}	
`, i-1, genfp.FuncDeclArgs(2, i), genfp.FuncCallArgs(2, i))

		}
	})

	genfp.Generate("hlist", "of_gen.go", func(f genfp.Writer) {
		for i := 2; i < max.Product; i++ {

			fmt.Fprintf(f, "func Of%d [%s any](%s) %s { ", i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), consType(1, i, "Nil"))

			fmt.Fprintf(f, `
	return Concat(a1, Of%d(%s))
}	
`, i-1, genfp.FuncCallArgs(2, i))

		}
	})

	genfp.Generate("hlist", "reverse_gen.go", func(f genfp.Writer) {
		for i := 2; i < max.Func; i++ {

			fmt.Fprintf(f, "func Reverse%d [%s any](hl %s ) %s { ", i, genfp.FuncTypeArgs(1, i), consType(1, i, "Nil"), reversConsType(1, i))

			fmt.Fprintf(f, `
		return Case%d(hl, func(%s) %s {
			return Of%d(%s)
		})
	}
	`, i, genfp.FuncDeclArgs(1, i), reversConsType(1, i), i, genfp.ReverseFuncCallArgs(1, i))

		}
	})

	// 	fmt.Fprintln(f, `
	// import (
	// 	"github.com/csgura/fp"
	// )`)

}
