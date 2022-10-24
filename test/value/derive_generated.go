package value

import (
	"github.com/csgura/fp/eq"
)

var EqPerson = eq.ContraMap(eq.Tuple7(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String), eq.Slice(eq.String), eq.HCons(eq.String, eq.HCons(eq.Given[int](), eq.HNil)), EqFpSeq(EqFloat64)), Person.AsTuple)
