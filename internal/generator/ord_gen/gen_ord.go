package main

import (
	"go/types"

	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/metafp"
)

func main() {

	metafp.Generate("ord", "tuple_gen.go", func(f metafp.Writer) {

		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/eq", "eq"))

		f.Iteration(2, max.Product).Write(`
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Ord"}} ) fp.Ord[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New( eq.New( func( a, b fp.{{TupleType .N}} ) bool {
		return ins1.Eqv(a.Head(),b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.{{TupleType .N}}](func(t1 , t2 fp.{{TupleType .N}}) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}
		`, map[string]any{})

	})

}
