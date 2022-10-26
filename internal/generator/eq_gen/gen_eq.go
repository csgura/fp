package main

import (
	"go/types"

	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/metafp"
)

func main() {

	metafp.Generate("eq", "tuple_gen.go", func(f metafp.Writer) {

		f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		f.Iteration(2, max.Product).Write(`
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Eq"}} ) fp.Eq[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New(
		func(t1 , t2 fp.{{TupleType .N}}) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}
		`, map[string]any{})

	})
}
