package gendebug

import (
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	world    string
	universe string
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[Hello]]
