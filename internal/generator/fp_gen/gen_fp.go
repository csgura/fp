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

func curriedType(start, until int) string {
	f := &bytes.Buffer{}
	endBracket := ""
	for j := start; j <= until; j++ {
		fmt.Fprintf(f, "Func1[A%d, ", j)
		endBracket = endBracket + "]"
	}
	fmt.Fprintf(f, "R%s", endBracket)

	return f.String()
}

func funcTypeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "A%d", j)
	}
	return f.String()
}

func tupleTypeArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	return f.String()
}

func formatStr(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ",")
		}
		fmt.Fprintf(f, "%s", "%v")
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

func funcChain(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		if j == until {
			fmt.Fprintf(f, "f%d Func1[A%d,R]", j, j)
		} else {
			fmt.Fprintf(f, "f%d Func1[A%d,A%d]", j, j, j+1)
		}
	}
	return f.String()
}

func generate(packname string, filename string, writeFunc func(w io.Writer)) {
	f := &bytes.Buffer{}

	fmt.Fprintf(f, "package %s\n\n", packname)
	writeFunc(f)

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		lines := strings.Split(f.String(), "\n")
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

var ordinalName = []string{
	"Zero",
	"First",
	"Second",
	"Third",
	"Fourth",
	"Fifth",
	"Sixth",
	"Seventh",
	"Eighth",
	"Ninth",
	"Tenth",
}

func main() {
	generate("fp", "func_gen.go", func(f io.Writer) {
		for i := 3; i < max.Func; i++ {
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
		for i := 3; i < max.Func; i++ {

			fmt.Fprintf(f, `
func(r Func%d[%s,R]) Tupled() Func1[Tuple%d[%s],R] {
	return func(t Tuple%d[%s]) R {
		return r(t.Unapply())
	}
}
`, i, funcTypeArgs(1, i), i, funcTypeArgs(1, i), i, funcTypeArgs(1, i))

			fmt.Fprintf(f, `
func(r Func%d[%s,R]) Curried() %s {
	return func(a1 A1) %s {
		return Func%d[%s,R](func(%s) R {
			return r(%s)
		}).Curried()
	}	
}
`, i, funcTypeArgs(1, i), curriedType(1, i), curriedType(2, i), i-1, funcTypeArgs(2, i), funcDeclArgs(2, i), funcCallArgs(1, i))

			for j := i - 1; j < i; j++ {
				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyFirst%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), j, funcDeclArgs(1, j), i-j, funcTypeArgs(j+1, i), funcDeclArgs(j+1, i), funcCallArgs(1, i))
			}

			for j := i - 1; j < i; j++ {
				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyLast%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), j, funcDeclArgs(i-j+1, i), i-j, funcTypeArgs(1, i-j), funcDeclArgs(1, i-j), funcCallArgs(1, i))
			}

			/* 일부만 아규먼트 적용하는 함수는
			   너무 많은 code 를 생성해내서, 삭제

			   			fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyFirst(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), funcDeclArgs(1, 1), i-1, funcTypeArgs(2, i), funcDeclArgs(2, i), funcCallArgs(1, i))

			   			for j := 2; j < i; j++ {
			   				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyFirst%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), j, funcDeclArgs(1, j), i-j, funcTypeArgs(j+1, i), funcDeclArgs(j+1, i), funcCallArgs(1, i))
			   			}

			   			fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyLast(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), funcDeclArgs(i, i), i-1, funcTypeArgs(1, i-1), funcDeclArgs(1, i-1), funcCallArgs(1, i))

			   			for j := 2; j < i; j++ {
			   				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyLast%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), j, funcDeclArgs(i-j+1, i), i-j, funcTypeArgs(1, i-j), funcDeclArgs(1, i-j), funcCallArgs(1, i))

			   			}

			   			for j := 2; j < i; j++ {
			   				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) Apply%s(%s) Func%d[%s,%s,R] {
			   	return func(%s,%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, funcTypeArgs(1, i), ordinalName[j], funcDeclArgs(j, j), i-1, funcTypeArgs(1, j-1), funcTypeArgs(j+1, i), funcDeclArgs(1, j-1), funcDeclArgs(j+1, i), funcCallArgs(1, i))

			   			}

			*/

			// 			if i < max.Shift {
			// 				fmt.Fprintf(f, `
			// func(r Func%d[%s,R]) Shift() Func%d[%s,A1,R] {
			// 	return func(%s , a1 A1) R {
			// 		return r(%s)
			// 	}
			// }
			// `, i, funcTypeArgs(1, i), i, funcTypeArgs(2, i), funcDeclArgs(2, i), funcCallArgs(1, i))
			// 			}

		}

		for i := 3; i < max.Compose; i++ {
			fmt.Fprintf(f, `
func Compose%d[%s,R any] ( %s ) Func1[A1,R] {
	return Compose2(f1, Compose%d(%s))
}
			`, i, funcTypeArgs(1, i), funcChain(1, i), i-1, strings.ReplaceAll(funcCallArgs(2, i), "a", "f"))
		}

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Nop%d[%s,R any] ( f func(A%d) R ) Func%d[%s,R] {
	return func(%s) R {
		return f(a%d)
	}
}
			`, i-1, funcTypeArgs(1, i), i, i, funcTypeArgs(1, i), funcDeclArgs(1, i), i)
		}
	})

	generate("fp", "tuple_gen.go", func(f io.Writer) {
		fmt.Fprintln(f, `
import (
	"fmt"
	"github.com/csgura/fp/hlist"

)`)

		for i := 2; i < max.Product; i++ {
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
`, i, tupleTypeArgs(1, i))

			fmt.Fprintf(f, `
func (r Tuple%d[%s]) Tail() Tuple%d[%s] {
	return Tuple%d[%s]{%s};
}
`, i, tupleTypeArgs(1, i), i-1, tupleTypeArgs(2, i), i-1, tupleTypeArgs(2, i), tupleArgs(2, i))

			fmt.Fprintf(f, `
			func (r Tuple%d[%s]) ToHList() %s {
				return hlist.Concat( r.Head(), r.Tail().ToHList())
			}
			`, i, tupleTypeArgs(1, i), consType(1, i))

			fmt.Fprintf(f, `
func (r Tuple%d[%s]) String() string {
	return fmt.Sprintf("(%s)", %s)
}
`, i, tupleTypeArgs(1, i), formatStr(1, i), tupleArgs(1, i))

			fmt.Fprintf(f, `
func (r Tuple%d[%s]) Unapply() (%s) {
	return %s
}
`, i, tupleTypeArgs(1, i), tupleTypeArgs(1, i), tupleArgs(1, i))

		}

	})

}
