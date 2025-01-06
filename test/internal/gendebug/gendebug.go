package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type LocalEmbedPrivate struct {
	Name    string
	Age     int
	privacy string
	NoName  struct {
		Hello string
		world int
	}
}

// @fp.Derive
var _ eq.Derives[fp.Eq[LocalEmbedPrivate]]
