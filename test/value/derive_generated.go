package value

import (
	"github.com/csgura/fp/eq"
)

var EqPerson = eq.ContraMap(eq.Tuple4(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String)), Person.AsTuple)
