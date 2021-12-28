package main

import (
	"fmt"

	"github.com/csgura/fp/internal/generator/common"
	"github.com/csgura/fp/internal/max"
)

func main() {

	common.Generate("lazy", "tailcall_gen.go", func(f common.Writer) {
		fmt.Fprintln(f, `
import (
	"github.com/csgura/fp"
)`)

		f.Iteration(1, max.Func).Write(`
func TailCall{{.N}}[{{call .FuncTypeArgs 1 .N}}, R any]( f fp.Func{{.N}}[{{call .FuncTypeArgs 1 .N}} , Eval[R]], {{call .FuncDeclArgs 1 .N}} ) Eval[R] {
	return TailCall( func() Eval[R] {
		return f({{call .FuncCallArgs 1 .N}})
	})
}
		`, map[string]any{})

	})
}
