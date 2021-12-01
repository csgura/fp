package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
)

func main() {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package option\n\n")

	fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)`)

	for i := 2; i < 23; i++ {
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
	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concact(av, hv)
		})
	})

	return ApplicativeFunctor%d%s{nh, Ap(r.fn, a)}
}
`, i-1, nexttp)

		fmt.Fprintf(f, "%s Ap( a A1) ApplicativeFunctor%d%s {\n", receiver, i-1, nexttp)
		fmt.Fprintln(f, `
	return r.ApOption(Some(a))

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

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile("applicative_gen.go", formatted, 0666)
	if err != nil {
		return
	}
}