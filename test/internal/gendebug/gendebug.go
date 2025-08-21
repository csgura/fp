package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type World struct {
	Value string
}

type Hello struct {
	A *World
}

// @fp.Derive(recursive=true)
var _ eq.Derives[fp.Eq[Hello]]
