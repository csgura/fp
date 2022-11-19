package main

import (
	"go/types"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {

	genfp.Generate("ord", "tuple_gen.go", func(f genfp.Writer) {

		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/eq", "eq"))
		_ = f.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		f.Iteration(2, max.Product).Write(`
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Ord"}} ) fp.Ord[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New( eq.New( func( a, b fp.{{TupleType .N}} ) bool {
		return ins1.Eqv(a.Head(),b.Head()) && pt.Eqv(as.Tuple{{dec .N}}(a.Tail()), as.Tuple{{dec .N}}(b.Tail()))
	}), fp.LessFunc[fp.{{TupleType .N}}](func(t1 , t2 fp.{{TupleType .N}}) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(as.Tuple{{dec .N}}(t1.Tail()), as.Tuple{{dec .N}}(t2.Tail()))
	}))
}
		`, map[string]any{})

	})

}
