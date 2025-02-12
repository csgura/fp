package testmeta

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/minimal"
	"github.com/csgura/fp/mshow"
	fshow "github.com/csgura/fp/show"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/ngap"
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
	Entry fp.Entry[int]
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

type NamedTypeParam[S any] fp.Tuple1[S]

func (r NamedTypeParam[S]) Name() string {
	return "opt"
}
func (r NamedTypeParam[S]) Value() S {
	return r.I1
}
func (r NamedTypeParam[S]) Tag() string {
	return ``
}
func (r NamedTypeParam[S]) Static() bool {
	return true
}

var fp4fn = ngap.Tuple4[*int, *string, *float64]
var tp4 = fp.Tuple4[int, *int, *string, *float64]{}

// @test.Summon
var hello js.Encoder[string]

var ShowNil = mshow.New(func(minimal.Nil) string {
	return "nil"
})

func returnError() error {
	return nil
}

func returnAny() any {
	return nil
}

func ShowAny[T any]() fp.Show[T] {
	return fshow.New(func(t T) string {
		return fmt.Sprint(t)
	})
}
