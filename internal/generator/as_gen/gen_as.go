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

	generate("as", "func_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 1; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Func%d[%s,R any]( f func(%s) R) fp.Func%d[%s,R] {
	return fp.Func%d[%s,R](f)
}
`, i, typeArgs(1, i), typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(1, i))

		}

		fmt.Fprintf(f, `
func Curried2[A1, A2, R any](f fp.Func2[A1,A2,R]) fp.Func1[A1, fp.Func1[A2, R]] {
	return func(a1 A1) fp.Func1[A2, R] {
		return func(a2 A2) R {
			return f(a1, a2)
		}
	}
}		
		`)

		for i := 3; i < max.Func; i++ {

			fmt.Fprintf(f, `
func Curried%d[%s,R any]( f fp.Func%d[%s, R]) %s {
	return func(a1 A1) %s {
		return Curried%d( func(%s) R {
			return f(%s)
		} )
	}	
}
`, i, typeArgs(1, i), i, typeArgs(1, i), curriedType(1, i), curriedType(2, i), i-1, common.FuncDeclArgs(2, i), common.FuncCallArgs(1, i))

		}

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, `
func UnTupled%d[%s,R any]( f func(fp.Tuple%d[%s]) R) fp.Func%d[%s,R] {
	return func(%s) R {
		return f(Tuple%d(%s))
	}
}
`, i, typeArgs(1, i), i, typeArgs(1, i), i, typeArgs(1, i), funcDeclArgs(1, i), i, funcCallArgs(1, i))

		}

	})

	generate("as", "tuple_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		for i := 1; i < max.Product; i++ {

			fmt.Fprintf(f, "func Tuple%d [%s any]( %s ) fp.Tuple%d[%s] { ", i, typeArgs(1, i), funcDeclArgs(1, i), i, typeArgs(1, i))

			fmt.Fprintf(f, "  return fp.Tuple%d[%s] {\n", i, typeArgs(1, i))
			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d: a%d,\n", j, j)
			}
			fmt.Fprintf(f, `}
		}
`)

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
