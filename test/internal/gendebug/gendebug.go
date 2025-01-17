package gendebug

import (
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type EmbeddedTypeParamStruct[T any] struct {
	hello string
	world struct {
		Level T
		Stage string
	}
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[EmbeddedTypeParamStruct[any]]]
