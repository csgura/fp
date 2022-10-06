package common

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

func FuncTypeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "A%d", j)
	}
	return f.String()
}

type Monad string

func (r Monad) FuncChain(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		if j == until {
			fmt.Fprintf(f, "f%d fp.Func1[A%d,%s[R]]", j, j, r)
		} else {
			fmt.Fprintf(f, "f%d fp.Func1[A%d,%s[A%d]]", j, j, r, j+1)
		}
	}
	return f.String()
}

type Args string

func (r Args) Call(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d", r, j)
	}
	return f.String()
}

type writer struct {
	Package string
	Buffer  *bytes.Buffer
}

func (r *writer) Write(b []byte) (int, error) {
	return r.Buffer.Write(b)
}

func (r *writer) Iteration(start, end int) Range {
	return Range{r, start, end}
}

type Writer interface {
	io.Writer
	Iteration(start, end int) Range
}

func Generate(packname string, filename string, writeFunc func(w Writer)) {
	f := &writer{packname, &bytes.Buffer{}}

	fmt.Fprintf(f, "package %s\n\n", packname)
	writeFunc(f)

	formatted, err := format.Source(f.Buffer.Bytes())
	if err != nil {
		lines := strings.Split(f.Buffer.String(), "\n")
		for i := range lines {
			lines[i] = fmt.Sprintf("%d: %s", i, lines[i])
		}
		log.Print(strings.Join(lines, "\n"))
		log.Fatal("format error ", err)

		return
	}

	err = ioutil.WriteFile(filename, formatted, 0644)
	if err != nil {
		return
	}
}

func FuncDeclArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d A%d", j, j)
	}
	return f.String()
}

func FuncCallArgs(start, until int, prefixOpt ...string) string {

	prefix := "a"
	if len(prefixOpt) > 0 {
		prefix = prefixOpt[0]
	}
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d", prefix, j)
	}
	return f.String()
}

func FuncDeclTypeClassArgs(start, until int, typeClass string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "ins%d %s[A%d]", j, typeClass, j)
	}
	return f.String()
}

func TypeClassArgs(start, until int, typeClass string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s[A%d]", typeClass, j)
	}
	return f.String()
}

type Range struct {
	writer *writer
	start  int
	end    int
}

var defaultFunc = map[string]any{
	"TypeArgs":          FuncTypeArgs,
	"DeclArgs":          FuncDeclArgs,
	"CallArgs":          FuncCallArgs,
	"DeclTypeClassArgs": FuncDeclTypeClassArgs,
	"TupleType": func(n int) string {
		return fmt.Sprintf("Tuple%d[%s]", n, FuncTypeArgs(1, n))
	},
	"dec": func(n int) int {
		return n - 1
	},
}

func (r Range) Write(txt string, param map[string]any) {

	if param == nil {
		param = map[string]any{}
	}

	tpl, err := template.New("write").Funcs(defaultFunc).Parse(txt)
	if err == nil {
		for i := r.start; i < r.end; i++ {
			param["N"] = i
			err := tpl.Execute(r.writer, param)
			if err != nil {
				panic(err)
			}
		}
	} else {
		panic(err)
	}
}
