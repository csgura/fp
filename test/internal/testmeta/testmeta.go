package testmeta

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

func Named[V fp.NamedField[*A], A any]() fp.Eq[V] {
	panic("not implemented")
}

type Person struct {
	name string
	age  int
}

var _ hash.Derives[fp.Hashable[Person]]
var _ show.Derives[fp.Show[Person]]

type HasTuple struct {
	Entry fp.Tuple2[string, int]
	HList hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
}

var _ show.Derives[fp.Show[HasTuple]]

var Intlist = hlist.Concat(10, hlist.Empty())

type LocalPerson struct {
	Name string
	age  int
}

// @fp.Derive
var _ js.Derives[js.Decoder[LocalPerson]]

var StringNamed = as.NamedWithTag[string]("hello", "world", "")

var carOpt = NamedOptOfCar[int]{}

var _ fp.NamedField[fp.Option[int]] = NamedOptOfCar[int]{}

type NamedOptOfCar[T comparable] fp.Tuple1[fp.Option[T]]

func (r NamedOptOfCar[T]) Name() string {
	return "opt"
}
func (r NamedOptOfCar[T]) Value() fp.Option[T] {
	return r.I1
}
func (r NamedOptOfCar[T]) Tag() string {
	return ``
}
func (r NamedOptOfCar[T]) Static() bool {
	return true
}
