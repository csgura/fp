package main

import (
	"go/types"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {

	genfp.Generate("hash", "tuple_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/eq", "eq"))

		f.Iteration(2, max.Product).Write(`
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Hashable"}} ) fp.Hashable[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New( eq.New( func( a, b fp.{{TupleType .N}} ) bool {
		return ins1.Eqv(a.Head(),b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.{{TupleType .N}} ) uint32 {
		return ins1.Hash(t.Head()) * 31 + pt.Hash(t.Tail())
	})
}
		`, map[string]any{})

	})
}
