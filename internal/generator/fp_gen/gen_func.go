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
