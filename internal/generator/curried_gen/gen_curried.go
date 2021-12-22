package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"

	"github.com/csgura/fp/internal/max"
)

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

func curriedCallArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {

		fmt.Fprintf(f, "(a%d)", j)
	}
	return f.String()
}

func callFunc(nargs int) string {
	f := &bytes.Buffer{}

	argTypes := ""
	args := ""

	for i := 2; i <= nargs; i++ {
		if i != 2 {
			argTypes = argTypes + ","
			args = args + ","
		}
		argTypes = argTypes + fmt.Sprintf("a%d A%d", i, i)
		args = args + fmt.Sprintf("a%d", i)

	}

	fmt.Fprintf(f, `func(%s) R {
		    return f(a1, %s)
	    }`, argTypes, args)

	return f.String()
}

func main() {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package curried\n\n")

	fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

	for i := 2; i < max.Func; i++ {

		args := ""

		for j := 1; j <= i; j++ {
			if j != 1 {
				args = args + ","
			}
			args = args + fmt.Sprintf("A%d", j)
		}
		fmt.Fprintf(f, "func Func%d [%s, R any]( f func(%s) R ) %s { ", i, args, args, curriedType(1, i))

		fmt.Fprintf(f, `
	return func(a1 A1) %s {
		return Func%d(%s)
	}	
}	
`, curriedType(2, i), i-1, callFunc(i))

		fmt.Fprintf(f, "func Revert%d [%s, R any]( f %s ) fp.Func%d[%s,R] { ", i, args, curriedType(1, i), i, args)

		fmt.Fprintf(f, `
	return func(%s) R {
		return f%s
	}	
}	
`, funcDeclArgs(1, i), curriedCallArgs(1, i))

	}

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile("curried_gen.go", formatted, 0666)
	if err != nil {
		return
	}
}
