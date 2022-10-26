package hello

import (
	"github.com/csgura/fp/eq"
	"time"
)

var EqWorld = eq.ContraMap(eq.Tuple2(eq.String, eq.Given[time.Time]()), World.AsTuple)
