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
func TailCall{{.N}}[{{TypeArgs 1 .N}}, R any]( f fp.Func{{.N}}[{{TypeArgs 1 .N}} , Eval[R]], {{DeclArgs 1 .N}} ) Eval[R] {
	return TailCall( func() Eval[R] {
		return f({{CallArgs 1 .N}})
	})
}
		`, map[string]any{})

	})
}
