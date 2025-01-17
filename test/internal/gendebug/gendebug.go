package gendebug

import "github.com/csgura/fp/mshow"

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	Hello string
	World int
}

type World struct {
	Hello Hello
}

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[World]]
