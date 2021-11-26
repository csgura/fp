package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
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

func reversConsType(start, until int) string {
	ret := "Nil"
	for j := start; j <= until; j++ {
		ret = fmt.Sprintf("Cons[A%d, %s]", j, ret)
	}
	return ret
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

	fmt.Fprintf(f, "package hlist\n\n")

	// 	fmt.Fprintln(f, `
	// import (
	// 	"github.com/csgura/fp"
	// )`)

	for i := 2; i < 23; i++ {

		fmt.Fprintf(f, "func Ap%d [%s, R any]( f func(%s) R ) func(%s) R { ", i, typeArgs(1, i), funcDeclArgs(1, i), reversConsType(1, i))

		fmt.Fprintf(f, `
	return func(v %s) R {
		rf := Ap%d(func(%s) R {
			return f(%s, v.Head())
		})

		return rf(v.Tail())
	}
}	
`, reversConsType(1, i), i-1, funcDeclArgs(1, i-1), funcCallArgs(1, i-1))

	}

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		log.Print(f.String())
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile("ap_gen.go", formatted, 0666)
	if err != nil {
		return
	}
}
