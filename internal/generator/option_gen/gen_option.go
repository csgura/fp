package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/csgura/fp/internal/generator/common"
	"github.com/csgura/fp/internal/max"
)

func typeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "A%d", j)
	}
	return f.String()
}

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

func curriedType(start, until int) string {
	f := &bytes.Buffer{}
	endBracket := ""
	for j := start; j <= until; j++ {
		fmt.Fprintf(f, "fp.Func1[A%d, ", j)
		endBracket = endBracket + "]"
	}
	fmt.Fprintf(f, "R%s", endBracket)

	return f.String()
}

func generate(packname string, filename string, writeFunc func(w io.Writer)) {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package %s\n\n", packname)
	writeFunc(f)

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		lines := strings.Split(f.String(), "\n")
		for i := range lines {
			lines[i] = fmt.Sprintf("%d: %s", i, lines[i])
		}
		log.Print(strings.Join(lines, "\n"))
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile(filename, formatted, 0644)
	if err != nil {
		return
	}
}

func main() {
	generate("option", "applicative_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)`)

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, "type ApplicativeFunctor%d [H hlist.Header[HT], HT ", i)

			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, ", A%d", j)
			}
			fmt.Fprintf(f, ",R any]")

			fmt.Fprintf(f, " struct {\n")
			fmt.Fprintf(f, "  h fp.Option[H]\n")
			fmt.Fprintf(f, "  fn fp.Option[")
			endBracket := "]"
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "fp.Func1[A%d, ", j)
				endBracket = endBracket + "]"
			}
			fmt.Fprintf(f, "R%s\n", endBracket)

			fmt.Fprintf(f, "}\n")

			typeparams := "[H, HT"

			nexttp := "[hlist.Cons[A1,H]"
			for j := 1; j <= i; j++ {
				typeparams = fmt.Sprintf("%s , A%d", typeparams, j)
				nexttp = fmt.Sprintf("%s, A%d", nexttp, j)
			}
			typeparams = typeparams + ", R]"
			nexttp = nexttp + ", R]"

			receiver := fmt.Sprintf("func (r ApplicativeFunctor%d%s)", i, typeparams)

			if i < max.Flip {

				fmt.Fprintf(f, "%s Flip() ApplicativeFunctor%d[H,HT,%s,R] {\n", receiver, i, flipTypeArgs(1, i))
				fmt.Fprintf(f, `
	return ApplicativeFunctor%d[H, HT, %s, R]{
		r.h,
		Map(r.fn, curried.Flip[A1,A2,%s]),
	}

}
`, i, flipTypeArgs(1, i), curriedType(3, i))
			}

			fmt.Fprintf(f, "%s FlatMap( a func(HT) fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}`)

			fmt.Fprintf(f, "%s Map( a func(HT) A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListMap( a func(H) A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListFlatMap( a func(H) fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}`)

			fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor%d%s{nh, Ap(r.fn, a)}
}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s Ap( a A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApOption(Some(a))

}`)

			fmt.Fprintf(f, "%s ApOptionFunc( a func() fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a()
	})
	return r.ApOption(av)
}`)

			fmt.Fprintf(f, "%s ApFunc( a func() A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApOption(av)
}`)

			tpr := ""
			for j := 1; j <= i; j++ {
				if j != 1 {
					tpr = tpr + ","
				}
				tpr = tpr + fmt.Sprintf("A%d", j)
			}
			tpr = tpr + ",R"

			fmt.Fprintf(f, "func Applicative%d[%s any](fn fp.Func%d[%s]) ApplicativeFunctor%d[hlist.Nil, hlist.Nil, %s] {\n", i, tpr, i, tpr, i, tpr)
			fmt.Fprintf(f, "    return ApplicativeFunctor%d[hlist.Nil, hlist.Nil, %s]{Some(hlist.Empty()), Some(curried.Func%d(fn))}\n", i, tpr, i)
			fmt.Fprintf(f, "}\n")

			// for j := 1; j <= i; j++ {
			// 	if j != 1 {
			// 		fmt.Fprintf(f, ",")
			// 	}
			// 	fmt.Fprintf(f, "a%d A%d", j, j)
			// }
			// fmt.Fprintf(f, ") R\n\n")
		}

	})

	generate("option", "func_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 3; i < max.Func; i++ {
			fmt.Fprintf(f, `
				func LiftA%d[%s,R any]( f func(%s) R ) fp.Func%d[%s,fp.Option[R]] {
					return func(%s) fp.Option[R] {

						return FlatMap(ins1, func(a1 A1) fp.Option[R] {
							return LiftA%d(func(%s) R {
								return f(%s)
							})(%s)
						})
					}
				}
			`, i, typeArgs(1, i), common.FuncDeclArgs(1, i), i, common.TypeClassArgs(1, i, "fp.Option"),
				common.FuncDeclTypeClassArgs(1, i, "fp.Option"),
				i-1, common.FuncDeclArgs(2, i),
				common.FuncCallArgs(1, i),
				common.FuncCallArgs(2, i, "ins"),
			)

			fmt.Fprintf(f, `
				func LiftM%d[%s,R any]( f func(%s) fp.Option[R] ) fp.Func%d[%s,fp.Option[R]] {
					return func(%s) fp.Option[R] {

						return FlatMap(ins1, func(a1 A1) fp.Option[R] {
							return LiftM%d(func(%s) fp.Option[R] {
								return f(%s)
							})(%s)
						})
					}
				}
			`, i, typeArgs(1, i), common.FuncDeclArgs(1, i), i, common.TypeClassArgs(1, i, "fp.Option"),
				common.FuncDeclTypeClassArgs(1, i, "fp.Option"),
				i-1, common.FuncDeclArgs(2, i),
				common.FuncCallArgs(1, i),
				common.FuncCallArgs(2, i, "ins"),
			)

			fmt.Fprintf(f, `
				func Flap%d[%s,R any](tf fp.Option[%s]) %s {
					return func(a1 A1) %s {
						return Flap%d(Ap(tf, Some(a1)))
					}
				}
			`, i, typeArgs(1, i), curriedType(1, i), common.CurriedType(1, i, "fp.Option[R]"),
				common.CurriedType(2, i, "fp.Option[R]"),
				i-1,
			)

			fmt.Fprintf(f, `
				func Method%d[%s,R any](ta1 fp.Option[A1], fa1 func(%s) R) fp.Func%d[%s, fp.Option[R]] {
					return func(%s) fp.Option[R] {
						return Map(ta1, func(a1 A1) R {
							return fa1(%s)
						})
					}
				}
			`, i, typeArgs(1, i), common.FuncDeclArgs(1, i), i-1, typeArgs(2, i),
				common.FuncDeclArgs(2, i),
				common.FuncCallArgs(1, i),
			)

			fmt.Fprintf(f, `
				func FlatMethod%d[%s,R any](ta1 fp.Option[A1], fa1 func(%s) fp.Option[R]) fp.Func%d[%s, fp.Option[R]] {
					return func(%s) fp.Option[R] {
						return FlatMap(ta1, func(a1 A1) fp.Option[R] {
							return fa1(%s)
						})
					}
				}
			`, i, typeArgs(1, i), common.FuncDeclArgs(1, i), i-1, typeArgs(2, i),
				common.FuncDeclArgs(2, i),
				common.FuncCallArgs(1, i),
			)
		}

		for i := 3; i < max.Compose; i++ {
			fmt.Fprintf(f, `
func Compose%d[%s,R any] ( %s ) fp.Func1[A1,fp.Option[R]] {
	return Compose2(f1, Compose%d(%s))
}
			`, i, common.FuncTypeArgs(1, i), common.Monad("fp.Option").FuncChain(1, i), i-1, common.Args("f").Call(2, i))
		}
	})
}
