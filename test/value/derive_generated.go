package value

import (
	"github.com/csgura/fp/eq"
)

var EqPerson = eq.ContraMap(eq.Tuple2(eq.String, eq.Given[int]()), Person.AsTuple)
