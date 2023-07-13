package showtest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/show"
	"github.com/csgura/fp/test/internal/recursive"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Person struct {
	Name string
	Age  int
}

// @fp.Derive
var _ show.Derives[fp.Show[Person]]

type NoDerive struct {
	Hello string
}

type Collection struct {
	Index       map[string]Person
	List        []Person
	Description *string
	Set         fp.Set[int]
	Option      fp.Option[Person]
	NoDerive    NoDerive
	Stringer    HasStringMethod
	BoolPtr     *bool
	NoMap       map[string]NoDerive
	Alias       recursive.StringAlias
}

// @fp.Derive(recursive=true)
var _ show.Derives[fp.Show[Collection]]

type HasStringMethod struct {
	There string
}

func (r HasStringMethod) String() string {
	return r.There
}

type DupGenerate struct {
	NoDerive NoDerive
	World    string
}

// @fp.Derive(recursive=true)
var _ show.Derives[fp.Show[DupGenerate]]
