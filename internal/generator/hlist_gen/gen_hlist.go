package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
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

func main() {

	generate("hlist", "ap_gen.go", func(f io.Writer) {

		for i := 2; i < 23; i++ {

			fmt.Fprintf(f, "func Lift%d [%s, R any]( f func(%s) R ) func(%s) R { ", i, typeArgs(1, i), funcDeclArgs(1, i), consType(1, i, "Nil"))

			fmt.Fprintf(f, `
	return func(v %s) R {
		rf := Lift%d(func(%s) R {
			return f(v.Head(), %s)
		})

		return rf(v.Tail())
	}
}	
`, consType(1, i, "Nil"), i-1, funcDeclArgs(1+1, i), funcCallArgs(1+1, i))

			fmt.Fprintf(f, "func Rift%d [%s, R any]( f func(%s) R ) func(%s) R { ", i, typeArgs(1, i), funcDeclArgs(1, i), reversConsType(1, i))

			fmt.Fprintf(f, `
	return func(v %s) R {
		rf := Rift%d(func(%s) R {
			return f(%s, v.Head())
		})

		return rf(v.Tail())
	}
}	
`, reversConsType(1, i), i-1, funcDeclArgs(1, i-1), funcCallArgs(1, i-1))

		}
	})

	generate("hlist", "case_gen.go", func(f io.Writer) {
		for i := 2; i < 23; i++ {

			fmt.Fprintf(f, "func Case%d [%s any, T Sealed, R any](hl %s,  f func(%s) R ) R { ", i, typeArgs(1, i), consType(1, i, "T"), funcDeclArgs(1, i))

			fmt.Fprintf(f, `
	return Case%d(hl.Tail(), func(%s) R {
		return f(hl.Head(), %s)
	})
}	
`, i-1, funcDeclArgs(2, i), funcCallArgs(2, i))

		}
	})

	generate("hlist", "of_gen.go", func(f io.Writer) {
		for i := 2; i < 23; i++ {

			fmt.Fprintf(f, "func Of%d [%s any](%s) %s { ", i, typeArgs(1, i), funcDeclArgs(1, i), consType(1, i, "Nil"))

			fmt.Fprintf(f, `
	return Concat(a1, Of%d(%s))
}	
`, i-1, funcCallArgs(2, i))

		}
	})

	generate("hlist", "reverse_gen.go", func(f io.Writer) {
		for i := 2; i < 23; i++ {

			fmt.Fprintf(f, "func Reverse%d [%s any](hl %s ) %s { ", i, typeArgs(1, i), consType(1, i, "Nil"), reversConsType(1, i))

			fmt.Fprintf(f, `
		return Case%d(hl, func(%s) %s {
			return Of%d(%s)
		})
	}
	`, i, funcDeclArgs(1, i), reversConsType(1, i), i, reverseFuncCallArgs(1, i))

		}
	})

	// 	fmt.Fprintln(f, `
	// import (
	// 	"github.com/csgura/fp"
	// )`)

}
