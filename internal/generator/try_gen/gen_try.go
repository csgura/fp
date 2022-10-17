package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"

	"github.com/csgura/fp/internal/generator/common"
	"github.com/csgura/fp/internal/max"
)

func generate(packname string, filename string, writeFunc func(w io.Writer)) {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package %s\n\n", packname)
	writeFunc(f)

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile(filename, formatted, 0644)
	if err != nil {
		return
	}
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

func funcDeclArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d A%d", j, j)
	}
	return f.String()
}

func funcCallArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d", j)
	}
	return f.String()
}

func reverseFuncCallArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := until; j >= start; j-- {
		if j != until {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d", j)
	}
	return f.String()
}

func reversConsType(start, until int) string {
	ret := "Nil"
	for j := start; j <= until; j++ {
		ret = fmt.Sprintf("Cons[A%d, %s]", j, ret)
	}
	return ret
}

func consType(start, until int, last string) string {
	ret := last
	for j := until; j >= start; j-- {
		ret = fmt.Sprintf("Cons[A%d, %s]", j, ret)
	}
	return ret
}

func main() {

	generate("try", "applicative_gen.go", func(f io.Writer) {

		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)`)

		for i := 2; i < max.Func; i++ {

			fmt.Fprintf(f, `
type MonadChain%d [H hlist.Header[HT], HT , %s , R any] struct {
	h fp.Try[H]
	fn fp.Try[%s]
}
`,
				i,
				typeArgs(1, i),
				curriedType(1, i),
			)

			receiver := fmt.Sprintf("func (r MonadChain%d[H,HT,%s,R])", i, typeArgs(1, i))
			nexttp := fmt.Sprintf("[hlist.Cons[A1,H], %s, R]", typeArgs(1, i))

			if i < max.Flip {

				fmt.Fprintf(f, "%s Flip() MonadChain%d[H,HT,%s,R] {\n", receiver, i, flipTypeArgs(1, i))
				fmt.Fprintf(f, `
	return MonadChain%d[H, HT, %s, R]{
		r.h,
		Map(r.fn, curried.Flip[A1,A2,%s]),
	}

}
`, i, flipTypeArgs(1, i), curriedType(3, i))
			}

			fmt.Fprintf(f, "%s FlatMap( a func(HT) fp.Try[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "%s Map( a func(HT) A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListMap( a func(H) A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListFlatMap( a func(H) fp.Try[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "%s ApTry( a fp.Try[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain%d%s{nh, Ap(r.fn, a)}
}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	return r.ApTry(FromOption(a))
}
`)

			fmt.Fprintf(f, "%s Ap( a A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApTry(Success(a))

}`)

			fmt.Fprintf(f, "%s ApTryFunc( a func() fp.Try[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "%s ApOptionFunc( a func() fp.Option[A1]) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "%s ApFunc( a func() A1) MonadChain%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "func Chain%d[%s , R any](fn fp.Func%d[%s,R]) MonadChain%d[hlist.Nil, hlist.Nil, %s,R] {\n", i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(1, i))
			fmt.Fprintf(f, "    return MonadChain%d[hlist.Nil, hlist.Nil, %s,R]{Success(hlist.Empty()), Success(curried.Func%d(fn))}\n", i, typeArgs(1, i), i)
			fmt.Fprintf(f, "}\n")
		}

		for i := 2; i < max.Func; i++ {

			fmt.Fprintf(f, `
type ApplicativeFunctor%d [%s , R any] struct {
	fn fp.Try[%s]
}
`,
				i,
				typeArgs(1, i),
				curriedType(1, i),
			)

			receiver := fmt.Sprintf("func (r ApplicativeFunctor%d[%s,R])", i, typeArgs(1, i))
			nexttp := fmt.Sprintf("[%s, R]", typeArgs(2, i))

			if i < max.Flip {

				fmt.Fprintf(f, "%s Flip() ApplicativeFunctor%d[%s,R] {\n", receiver, i, flipTypeArgs(1, i))
				fmt.Fprintf(f, `
	return ApplicativeFunctor%d[%s, R]{
		Map(r.fn, curried.Flip[A1,A2,%s]),
	}

}
`, i, flipTypeArgs(1, i), curriedType(3, i))
			}

			fmt.Fprintf(f, "%s ApTry( a fp.Try[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `

	return ApplicativeFunctor%d%s{Ap(r.fn, a)}
}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	return r.ApTry(FromOption(a))
}
`)

			fmt.Fprintf(f, "%s Ap( a A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApTry(Success(a))

}`)

			fmt.Fprintf(f, "%s ApTryFunc( a func() fp.Try[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
		return ApplicativeFunctor%d%s{ApFunc(r.fn, a)}

}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s ApOptionFunc( a func() fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}`)

			fmt.Fprintf(f, "%s ApFunc( a func() A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}`)

			fmt.Fprintf(f, "func Applicative%d[%s , R any](fn fp.Func%d[%s,R]) ApplicativeFunctor%d[%s,R] {\n", i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(1, i))
			fmt.Fprintf(f, "    return ApplicativeFunctor%d[%s,R]{Success(curried.Func%d(fn))}\n", i, typeArgs(1, i), i)
			fmt.Fprintf(f, "}\n")
		}

	})

	generate("try", "func_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 3; i < max.Func; i++ {
			fmt.Fprintf(f, `
				func LiftA%d[%s,R any]( f func(%s) R ) fp.Func%d[%s,fp.Try[R]] {
					return func(%s) fp.Try[R] {

						return FlatMap(ins1, func(a1 A1) fp.Try[R] {
							return LiftA%d(func(%s) R {
								return f(%s)
							})(%s)
						})
					}
				}
			`, i, typeArgs(1, i), funcDeclArgs(1, i), i, common.TypeClassArgs(1, i, "fp.Try"),
				common.FuncDeclTypeClassArgs(1, i, "fp.Try"),
				i-1, funcDeclArgs(2, i),
				funcCallArgs(1, i),
				common.FuncCallArgs(2, i, "ins"),
			)

			fmt.Fprintf(f, `
				func LiftM%d[%s,R any]( f func(%s) fp.Try[R] ) fp.Func%d[%s,fp.Try[R]] {
					return func(%s) fp.Try[R] {

						return FlatMap(ins1, func(a1 A1) fp.Try[R] {
							return LiftM%d(func(%s) fp.Try[R] {
								return f(%s)
							})(%s)
						})
					}
				}
			`, i, typeArgs(1, i), funcDeclArgs(1, i), i, common.TypeClassArgs(1, i, "fp.Try"),
				common.FuncDeclTypeClassArgs(1, i, "fp.Try"),
				i-1, funcDeclArgs(2, i),
				funcCallArgs(1, i),
				common.FuncCallArgs(2, i, "ins"),
			)

			fmt.Fprintf(f, `
				func Flap%d[%s,R any](tf fp.Try[%s]) %s {
					return func(a1 A1) %s {
						return Flap%d(Ap(tf, Success(a1)))
					}
				}
			`, i, typeArgs(1, i), curriedType(1, i), common.CurriedType(1, i, "fp.Try[R]"),
				common.CurriedType(2, i, "fp.Try[R]"),
				i-1,
			)

			fmt.Fprintf(f, `
				func Method%d[%s,R any](ta1 fp.Try[A1], fa1 func(%s) R) fp.Func%d[%s, fp.Try[R]] {
					return func(%s) fp.Try[R] {
						return Map(ta1, func(a1 A1) R {
							return fa1(%s)
						})
					}
				}
			`, i, typeArgs(1, i), funcDeclArgs(1, i), i-1, typeArgs(2, i),
				funcDeclArgs(2, i),
				funcCallArgs(1, i),
			)

			fmt.Fprintf(f, `
				func FlatMethod%d[%s,R any](ta1 fp.Try[A1], fa1 func(%s) fp.Try[R]) fp.Func%d[%s, fp.Try[R]] {
					return func(%s) fp.Try[R] {
						return FlatMap(ta1, func(a1 A1) fp.Try[R] {
							return fa1(%s)
						})
					}
				}
			`, i, typeArgs(1, i), funcDeclArgs(1, i), i-1, typeArgs(2, i),
				funcDeclArgs(2, i),
				funcCallArgs(1, i),
			)
		}

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Func%d[%s,R any]( f func(%s) (R,error)) fp.Func%d[%s,fp.Try[R]] {
	return func(%s) fp.Try[R] {
		ret , err := f(%s)
		return Apply(ret,err)
	}
}
`, i, typeArgs(1, i), typeArgs(1, i), i, typeArgs(1, i), funcDeclArgs(1, i), funcCallArgs(1, i))

			fmt.Fprintf(f, `
func Unit%d[%s any]( f func(%s) error) fp.Func%d[%s,fp.Try[fp.Unit]] {
	return func(%s) fp.Try[fp.Unit] {
		err := f(%s)
		return Apply(fp.Unit{},err)
	}
}
`, i, typeArgs(1, i), typeArgs(1, i), i, typeArgs(1, i), funcDeclArgs(1, i), funcCallArgs(1, i))

		}

		for i := 3; i < max.Compose; i++ {
			fmt.Fprintf(f, `
func Compose%d[%s,R any] ( %s ) fp.Func1[A1,fp.Try[R]] {
	return Compose2(f1, Compose%d(%s))
}
			`, i, common.FuncTypeArgs(1, i), common.Monad("fp.Try").FuncChain(1, i), i-1, common.Args("f").Call(2, i))
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
