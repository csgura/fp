package gendebug

import (
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type HasTuple struct {
	HList hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[HasTuple]]
