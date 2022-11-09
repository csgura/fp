package testjson

import "github.com/csgura/fp/test/internal/js"

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.GenLabelled
type Root struct {
	a int
	b string
	c float64
	d bool
	e *int
	f []int
	g map[string]int
}

// @fp.Derive
var _ js.Derives[js.Encoder[Root]]
