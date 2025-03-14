package mshowtest

import (
	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/test/internal/showtest"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive
var _ mshow.Derives[mshow.Show[showtest.Person]]

// @fp.Derive(recursive=true,noinline)
var _ mshow.Derives[mshow.Show[showtest.Collection]]

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[showtest.DupGenerate]]

// @fp.Derive
var _ mshow.Derives[mshow.Show[showtest.HasTuple]]

// @fp.Value
type EmbeddedStruct struct {
	hello string
	world struct {
		Level int
		Stage string
	}
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[EmbeddedStruct]]

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

// @fp.Derive
var _ mshow.Derives[mshow.Show[showtest.EmptyStruct]]

// @fp.Derive
var _ mshow.Derives[mshow.Show[showtest.HasAliasType]]
