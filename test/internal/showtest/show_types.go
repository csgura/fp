package showtest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Person struct {
	Name string
	Age  int
}

// @fp.Derive
var _ show.Derives[fp.Show[Person]]

type Collection struct {
	Index       map[string]Person
	List        []Person
	Description *string
	Set         fp.Set[int]
	Option      fp.Option[string]
}

// @fp.Derive
var _ show.Derives[fp.Show[Collection]]
