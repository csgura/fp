package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/clone"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type MySeq []string

type HasReference struct {
	MS MySeq
}

// @fp.Derive(recursive=true)
var _ clone.Derives[fp.Clone[HasReference]]
