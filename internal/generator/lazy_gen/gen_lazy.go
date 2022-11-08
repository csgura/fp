package main

import (
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {

	genfp.Generate("lazy", "tailcall_gen.go", func(f genfp.Writer) {

		f.Iteration(1, max.Func).Write(`
func TailCall{{.N}}[{{TypeArgs 1 .N}}, R any]( f func({{TypeArgs 1 .N}}) Eval[R], {{DeclArgs 1 .N}} ) Eval[R] {
	return TailCall( func() Eval[R] {
		return f({{CallArgs 1 .N}})
	})
}
		`, map[string]any{})

	})
}
