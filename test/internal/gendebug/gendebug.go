package gendebug

import (
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type Hello struct {
	hello    string
	world    int
	universe string
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[Hello]]
