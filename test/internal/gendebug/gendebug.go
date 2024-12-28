package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/test/internal/js"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.GenLabelled
type Car[S any, T comparable] struct {
	company string `column:"company"`
	model   string
	year    int
	spec    S
	opt     fp.Option[T]
}

// @fp.Derive
var _ js.Derives[js.Encoder[Car[any, any]]]
