package main

import (
	"bytes"
	"fmt"
	"go/types"

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
	genfp.Generate("option", "applicative_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/curried", "curried"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, "type MonadChain%d [H hlist.Header[HT], HT ", i)

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

			receiver := fmt.Sprintf("func (r MonadChain%d%s)", i, typeparams)

			if i < max.Flip {

				fmt.Fprintf(f, "%s Flip() MonadChain%d[H,HT,%s,R] {\n", receiver, i, flipTypeArgs(1, i))
				fmt.Fprintf(f, `
	return MonadChain%d[H, HT, %s, R]{
		r.h,
		Map(r.fn, curried.Flip[A1,A2,%s]),
	}

}
`, i, flipTypeArgs(1, i), genfp.CurriedType(3, i, "R"))
			}

			fmt.Fprintf(f, "%s FlatMap( a func(HT) fp.Option[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}`)

			fmt.Fprintf(f, "%s Map( a func(HT) A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListMap( a func(H) A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListFlatMap( a func(H) fp.Option[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}`)

			fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain%d%s{nh, Ap(r.fn, a)}
}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s Ap( a A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApOption(Some(a))

}`)

			fmt.Fprintf(f, "%s ApOptionFunc( a func() fp.Option[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a()
	})
	return r.ApOption(av)
}`)

			fmt.Fprintf(f, "%s ApFunc( a func() A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
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

			fmt.Fprintf(f, "func Chain%d[%s any](fn fp.Func%d[%s]) MonadChain%d[hlist.Nil, hlist.Nil, %s] {\n", i, tpr, i, tpr, i, tpr)
			fmt.Fprintf(f, "    return MonadChain%d[hlist.Nil, hlist.Nil, %s]{Some(hlist.Empty()), Some(curried.Func%d(fn))}\n", i, tpr, i)
			fmt.Fprintf(f, "}\n")

			// for j := 1; j <= i; j++ {
			// 	if j != 1 {
			// 		fmt.Fprintf(f, ",")
			// 	}
			// 	fmt.Fprintf(f, "a%d A%d", j, j)
			// }
			// fmt.Fprintf(f, ") R\n\n")

		}

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, "type ApplicativeFunctor%d [", i)

			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "A%d,", j)
			}
			fmt.Fprintf(f, "R any]")

			fmt.Fprintf(f, " struct {\n")
			fmt.Fprintf(f, "  fn fp.Option[")
			endBracket := "]"
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "fp.Func1[A%d, ", j)
				endBracket = endBracket + "]"
			}
			fmt.Fprintf(f, "R%s\n", endBracket)

			fmt.Fprintf(f, "}\n")

			typeparams := fmt.Sprintf("[%s,R]", genfp.FuncTypeArgs(1, i))
			nexttp := fmt.Sprintf("[%s,R]", genfp.FuncTypeArgs(2, i))

			receiver := fmt.Sprintf("func (r ApplicativeFunctor%d%s)", i, typeparams)

			fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `

	return ApplicativeFunctor%d%s{Ap(r.fn, a)}
}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s Ap( a A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApOption(Some(a))

}`)

			fmt.Fprintf(f, "%s ApOptionFunc( a func() fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
		return ApplicativeFunctor%d%s{ApFunc(r.fn, a)}

}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s ApFunc( a func() A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
		return r.ApOptionFunc(func() fp.Option[A1] {
			return Some(a())
		})
}`)

			tpr := ""
			for j := 1; j <= i; j++ {
				if j != 1 {
					tpr = tpr + ","
				}
				tpr = tpr + fmt.Sprintf("A%d", j)
			}
			tpr = tpr + ",R"

			fmt.Fprintf(f, "func Applicative%d[%s any](fn fp.Func%d[%s]) ApplicativeFunctor%d[%s] {\n", i, tpr, i, tpr, i, tpr)
			fmt.Fprintf(f, "    return ApplicativeFunctor%d[%s]{Some(curried.Func%d(fn))}\n", i, tpr, i)
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

	genfp.Generate("option", "func_gen.go", func(f genfp.Writer) {
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
			`, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), i, genfp.TypeClassArgs(1, i, "fp.Option"),
				genfp.FuncDeclTypeClassArgs(1, i, "fp.Option"),
				i-1, genfp.FuncDeclArgs(2, i),
				genfp.FuncCallArgs(1, i),
				genfp.FuncCallArgs(2, i, "ins"),
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
			`, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), i, genfp.TypeClassArgs(1, i, "fp.Option"),
				genfp.FuncDeclTypeClassArgs(1, i, "fp.Option"),
				i-1, genfp.FuncDeclArgs(2, i),
				genfp.FuncCallArgs(1, i),
				genfp.FuncCallArgs(2, i, "ins"),
			)

			fmt.Fprintf(f, `
				func Flap%d[%s,R any](tf fp.Option[%s]) %s {
					return func(a1 A1) %s {
						return Flap%d(Ap(tf, Some(a1)))
					}
				}
			`, i, genfp.FuncTypeArgs(1, i), genfp.CurriedType(1, i, "R"), genfp.CurriedType(1, i, "fp.Option[R]"),
				genfp.CurriedType(2, i, "fp.Option[R]"),
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
			`, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), i-1, genfp.FuncTypeArgs(2, i),
				genfp.FuncDeclArgs(2, i),
				genfp.FuncCallArgs(1, i),
			)

			fmt.Fprintf(f, `
				func FlatMethod%d[%s,R any](ta1 fp.Option[A1], fa1 func(%s) fp.Option[R]) fp.Func%d[%s, fp.Option[R]] {
					return func(%s) fp.Option[R] {
						return FlatMap(ta1, func(a1 A1) fp.Option[R] {
							return fa1(%s)
						})
					}
				}
			`, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i), i-1, genfp.FuncTypeArgs(2, i),
				genfp.FuncDeclArgs(2, i),
				genfp.FuncCallArgs(1, i),
			)
		}

		for i := 3; i < max.Compose; i++ {
			fmt.Fprintf(f, `
func Compose%d[%s,R any] ( %s ) fp.Func1[A1,fp.Option[R]] {
	return Compose2(f1, Compose%d(%s))
}
			`, i, genfp.FuncTypeArgs(1, i), genfp.Monad("fp.Option").FuncChain(1, i), i-1, genfp.Args("f").Call(2, i))
		}
	})
}
