package main

import (
	"bytes"
	"fmt"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func flipTypeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}

		if j == start {
			fmt.Fprintf(f, "A%d", j+1)
		} else if j == start+1 {
			fmt.Fprintf(f, "A%d", j-1)
		} else {
			fmt.Fprintf(f, "A%d", j)
		}
	}
	return f.String()
}

func main() {

	genfp.Generate("iterator", "func_gen.go", func(f genfp.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 3; i < max.Func; i++ {

			fmt.Fprintf(f, `
				func Flap%d[%s,R any](tf fp.Iterator[%s]) %s {
					return func(a1 A1) %s {
						return Flap%d(Ap(tf, Of(a1)))
					}
				}
			`, i, genfp.FuncTypeArgs(1, i), genfp.CurriedType(1, i, "R"), genfp.CurriedType(1, i, "fp.Iterator[R]"),
				genfp.CurriedType(2, i, "fp.Iterator[R]"),
				i-1,
			)

			fmt.Fprintf(f, `
				func Method%d[%s,R any](ta1 fp.Iterator[A1], fa1 func(%s) R) func(%s) fp.Iterator[R] {
					return func(%s) fp.Iterator[R] {
						return Map(ta1, func(a1 A1) R {
							return fa1(%s)
						})
					}
				}
			`, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), genfp.FuncTypeArgs(2, i),
				genfp.FuncDeclArgs(2, i),
				genfp.FuncCallArgs(1, i),
			)

		}
	})

}
