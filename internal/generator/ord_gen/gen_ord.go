package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"strings"

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

func funcDeclTypeClassArgs(start, until int, typeClass string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "tins%d %s[A%d]", j, typeClass, j)
	}
	return f.String()
}

func funcCallTypeClassArgs(start, until int, method string, argf func(n int) string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "tins%d.%s(%s)", j, method, argf(j))
	}
	return f.String()
}

func tupleNType(n int) string {
	return fmt.Sprintf("fp.Tuple%d[%s]", n, typeArgs(1, n))
}

func noTypeclassArgs(n int) string {
	return ""
}
func main() {

	generate("ord", "tuple_gen.go", func(f io.Writer) {

		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 2; i < max.Product; i++ {

			fmt.Fprintf(f, "func Tuple%d [%s any]( %s ) fp.Ord[fp.Tuple%d[%s]] { ", i, typeArgs(1, i), funcDeclTypeClassArgs(1, i, "fp.Ord"), i, typeArgs(1, i))

			tuple := tupleNType(i)

			tail := strings.ReplaceAll(funcCallArgs(2, i), "a", "tins")
			fmt.Fprintf(f, `
		return New[%s](
				fp.EqFunc[%s](func(t1 %s, t2 %s) bool {
					return tins1.Eqv(t1.I1, t2.I1) && Tuple%d(%s).Eqv(t1.Tail(), t2.Tail())
				}),
			fp.LessFunc[%s](func(t1 %s, t2 %s) bool {
				if tins1.Less(t1.I1, t2.I1) {
					return true
				}
				if tins1.Less(t2.I1, t1.I1) {
					return false
				}
				return Tuple%d(%s).Less(t1.Tail(), t2.Tail())
			}),
		)
}
			`, tuple, tuple, tuple, tuple, i-1, tail,
				tuple, tuple, tuple, i-1, tail,
			)
		}
	})
}
