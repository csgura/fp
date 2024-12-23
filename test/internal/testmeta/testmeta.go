package testmeta

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/test/internal/show"
)

func Named[V fp.NamedField[*A], A any]() fp.Eq[V] {
	panic("not implemented")
}

// @fp.Value
type Person struct {
	name string
	age  int
}

var _ hash.Derives[fp.Hashable[Person]]
var _ show.Derives[fp.Show[Person]]