package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	Hello string
	world int
}

type LocalEmbedPrivate struct {
	Name    string
	Age     int
	privacy string
	hello   Hello
}

// @fp.Derive(recursive=true)
var _ show.Derives[fp.Show[LocalEmbedPrivate]]
