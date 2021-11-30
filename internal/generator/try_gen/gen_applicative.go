package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
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

	generate("try", "applicative_gen.go", func(f io.Writer) {

		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)`)

		for i := 2; i < 23; i++ {

			fmt.Fprintf(f, `
type ApplicativeFunctor%d [H hlist.Header[HT], HT , %s , R any] struct {
	h fp.Try[H]
	fn fp.Try[%s]
}
`,
				i,
				typeArgs(1, i),
				curriedType(1, i),
			)

			receiver := fmt.Sprintf("func (r ApplicativeFunctor%d[H,HT,%s,R])", i, typeArgs(1, i))
			nexttp := fmt.Sprintf("[hlist.Cons[A1,H], %s, R]", typeArgs(1, i))

			fmt.Fprintf(f, "%s FlatMap( a func(HT) fp.Try[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "%s Map( a func(HT) A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListMap( a func(H) A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}`)

			fmt.Fprintf(f, "%s HListFlatMap( a func(H) fp.Try[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintln(f, `
	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}`)

			fmt.Fprintf(f, "%s ApTry( a fp.Try[A1]) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
			fmt.Fprintf(f, `
	nh := FlatMap(r.h, func(hv H) fp.Try[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concact(av, hv)
		})
	})

	return ApplicativeFunctor%d%s{nh, Ap(r.fn, a)}
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

			fmt.Fprintf(f, "func Applicative%d[%s , R any](fn fp.Func%d[%s,R]) ApplicativeFunctor%d[hlist.Nil, hlist.Nil, %s,R] {\n", i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(1, i))
			fmt.Fprintf(f, "    return ApplicativeFunctor%d[hlist.Nil, hlist.Nil, %s,R]{Success(hlist.Empty()), Success(curried.Func%d(fn))}\n", i, typeArgs(1, i), i)
			fmt.Fprintf(f, "}\n")
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
