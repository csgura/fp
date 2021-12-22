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

	generate("future", "applicative_gen.go", func(f io.Writer) {

		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)`)

		for i := 2; i < max.Func; i++ {

			fmt.Fprintf(f, `
type ApplicativeFunctor%d [H hlist.Header[HT], HT , %s , R any] struct {
	h fp.Future[H]
	fn fp.Future[%s]
}
`,
				i,
				typeArgs(1, i),
				curriedType(1, i),
			)

			receiver := fmt.Sprintf("func (r ApplicativeFunctor%d[H,HT,%s,R])", i, typeArgs(1, i))
			nexttp := fmt.Sprintf("[hlist.Cons[A1,H], %s, R]", typeArgs(1, i))

			fmt.Fprintf(f, "%s Shift() ApplicativeFunctor%d[H,HT,%s,A1,R] {\n", receiver, i, typeArgs(2, i))
			fmt.Fprintf(f, `
				nf := fp.Compose(curried.Revert%d[%s, R], fp.Compose(fp.Func%d[%s, R].Shift, fp.Func%d[%s, A1, R].Curried))
				return ApplicativeFunctor%d[H, HT, %s, A1, R]{
					r.h,
					Map(r.fn, nf),
				}

			}
			`, i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(2, i), i, typeArgs(2, i))

			fmt.Fprintf(f, "%s FlatMap( a func(HT) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}`)

			fmt.Fprintf(f, "%s Map( a func(HT) A1,ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}`)

			fmt.Fprintf(f, "%s HListMap( a func(H) A1, ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}`)

			fmt.Fprintf(f, "%s HListFlatMap( a func(H) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}`)

			fmt.Fprintf(f, "%s ApFuture( a fp.Future[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	nh := FlatMap(r.h, func(hv H) fp.Future[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor%d%s{nh, Ap(r.fn, a)}
}
`, i-1, nexttp)

			fmt.Fprintf(f, "%s ApTry( a fp.Try[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	return r.ApFuture(FromTry(a))
}
`)

			fmt.Fprintf(f, "%s ApOption( a fp.Option[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	return r.ApFuture(FromOption(a))
}
`)

			fmt.Fprintf(f, "%s Ap( a A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.ApFuture(Successful(a))

}`)

			fmt.Fprintf(f, "%s ApFutureFunc( a func() fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}`)

			fmt.Fprintf(f, "%s ApTryFunc( a func() fp.Try[A1], ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}`)

			fmt.Fprintf(f, "%s ApOptionFunc( a func() fp.Option[A1], ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}`)

			fmt.Fprintf(f, "%s ApFunc( a func() A1, ctx ...fp.ExecContext) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}`)

			fmt.Fprintf(f, "func Applicative%d[%s , R any](fn fp.Func%d[%s,R]) ApplicativeFunctor%d[hlist.Nil, hlist.Nil, %s,R] {\n", i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(1, i))
			fmt.Fprintf(f, "    return ApplicativeFunctor%d[hlist.Nil, hlist.Nil, %s,R]{Successful(hlist.Empty()), Successful(curried.Func%d(fn))}\n", i, typeArgs(1, i), i)
			fmt.Fprintf(f, "}\n")
		}
	})

	generate("future", "func_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Func%d[%s,R any]( f func(%s) (R,error) , exec ... fp.ExecContext) fp.Func%d[%s,fp.Future[R]] {
	return func(%s) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(%s)
		})
	}
}
`, i, typeArgs(1, i), typeArgs(1, i), i, typeArgs(1, i), funcDeclArgs(1, i), funcCallArgs(1, i))

			fmt.Fprintf(f, `
func Unit%d[%s any]( f func(%s) (error) , exec ... fp.ExecContext) fp.Func%d[%s,fp.Future[fp.Unit]] {
	return func(%s) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(%s)
			return fp.Unit{}, err
		})
	}
}
`, i, typeArgs(1, i), typeArgs(1, i), i, typeArgs(1, i), funcDeclArgs(1, i), funcCallArgs(1, i))

		}

		for i := 3; i < max.Compose; i++ {
			fmt.Fprintf(f, `
func Compose%d[%s,R any] ( %s , exec ...fp.ExecContext ) fp.Func1[A1,fp.Future[R]] {
	return Compose2(f1, Compose%d(%s, exec...), exec...)
}
			`, i, common.FuncTypeArgs(1, i), common.Monad("fp.Future").FuncChain(1, i), i-1, common.Args("f").Call(2, i))
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
