package docexample

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/show"
)

//go:generate gombok

// @fp.Value
type Person struct {
	name string
	age  int
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Person]]

// @fp.Derive
var _ hash.Derives[fp.Hashable[Person]]

func (r Person) Eq(other Person) bool {
	return HashablePerson.Eqv(r, other)
}

func (r Person) Hashcode() uint32 {
	return HashablePerson.Hash(r)
}

// @fp.Value
// @fp.Json
type Address struct {
	country string
	city    string
	street  string
}

// @fp.Value
// @fp.GenLabelled
type Car struct {
	company string
	model   string
	year    int
}

// @fp.Value
type Entry[A comparable, B any] struct {
	key   A
	value B
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Car]]

// @fp.Value
type CarsOwned struct {
	owner Person
	cars  fp.Seq[Car]
}

var EqFpSeqCar = eq.New(func(a, b fp.Seq[Car]) bool {
	ordCar := ord.New(EqCar, func(a, b Car) bool {
		return a.year < b.year
	})
	asorted := seq.Sort(a, ordCar)

	bsorted := seq.Sort(a, ordCar)

	return eq.Seq(EqCar).Eqv(asorted, bsorted)
})

// @fp.Derive
var _ eq.Derives[fp.Eq[CarsOwned]]

// @fp.Value
type User struct {
	name   string
	email  fp.Option[string]
	active bool
}

// @fp.Derive
var _ show.Derives[fp.Show[Address]]

// @fp.Derive
var _ js.Derives[js.Encoder[Car]]
