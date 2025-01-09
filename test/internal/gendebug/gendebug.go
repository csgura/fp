package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type HasTuple struct {
	Tuple fp.Tuple2[string, int]
	Entry fp.Entry[int]
}

// @fp.Derive
var _ show.Derives[fp.Show[HasTuple]]
