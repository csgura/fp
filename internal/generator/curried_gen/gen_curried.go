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

	for i := 2; i < 23; i++ {

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
