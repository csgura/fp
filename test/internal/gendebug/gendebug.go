package gendebug

import "github.com/csgura/fp/mshow"

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	Index    map[string]string
	world    bool
	location string
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[Hello]]
