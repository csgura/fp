package hello

import (
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/test/internal/js"
	"time"
)

var EqWorld = eq.ContraMap(eq.Tuple2(eq.String, eq.Given[time.Time]()), World.AsTuple)

var EncoderWorld = js.ContraMap(js.Labelled2(js.String, js.Time), World.AsLabelled)
