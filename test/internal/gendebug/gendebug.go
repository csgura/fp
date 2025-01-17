package gendebug

import "github.com/csgura/fp/mshow"

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	Hello string
	World int
}

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[Hello]]
