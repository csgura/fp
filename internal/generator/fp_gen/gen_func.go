package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
)

func typeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	return f.String()
}

func tupleArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "r.I%d", j)
	}
	return f.String()
}

func consType(start, until int) string {
	ret := "hlist.Nil"
	for j := until; j >= start; j-- {
		ret = fmt.Sprintf("hlist.Cons[T%d, %s]", j, ret)
	}
	return ret
}

func main() {

	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package fp\n\n")

	for i := 1; i < 23; i++ {
		fmt.Fprintf(f, "type Func%d", i)
		fmt.Fprintf(f, "[")

		for j := 1; j <= i; j++ {
			if j != 1 {
				fmt.Fprintf(f, ",")
			}
			fmt.Fprintf(f, "A%d", j)
		}
		fmt.Fprintf(f, ",R any]")

		fmt.Fprintf(f, "func( ")

		for j := 1; j <= i; j++ {
			if j != 1 {
				fmt.Fprintf(f, ",")
			}
			fmt.Fprintf(f, "a%d A%d", j, j)
		}
		fmt.Fprintf(f, ") R\n\n")
	}

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile("func_gen.go", formatted, 0666)
	if err != nil {
		return
	}

	f = &bytes.Buffer{}

	fmt.Fprintf(f, "package fp\n\n")

	fmt.Fprintln(f, `
import (
	"github.com/csgura/fp/hlist"
)`)

	for i := 2; i < 23; i++ {
		fmt.Fprintf(f, "type Tuple%d", i)
		fmt.Fprintf(f, "[")

		for j := 1; j <= i; j++ {
			if j != 1 {
				fmt.Fprintf(f, ",")
			}
			fmt.Fprintf(f, "T%d", j)
		}
		fmt.Fprintf(f, " any] ")

		fmt.Fprintf(f, "struct {\n")

		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "    I%d T%d\n", j, j)
		}
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, `
func (r Tuple%d[%s]) Head() T1 {
	return r.I1;
}
`, i, typeArgs(1, i))

		fmt.Fprintf(f, `
func (r Tuple%d[%s]) Tail() Tuple%d[%s] {
	return Tuple%d[%s]{%s};
}
`, i, typeArgs(1, i), i-1, typeArgs(2, i), i-1, typeArgs(2, i), tupleArgs(2, i))

		fmt.Fprintf(f, `
func (r Tuple%d[%s]) ToHList() %s {
	return hlist.Concact( r.Head(), r.Tail().ToHList())
}
`, i, typeArgs(1, i), consType(1, i))
	}

	formatted, err = format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())

		log.Fatal("format error ", err)
		return
	}

	err = ioutil.WriteFile("tuple_gen.go", formatted, 0666)
	if err != nil {
		return
	}

}
