package docexample

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/show"
)

// not work go:generate gombok
//go:generate go run github.com/csgura/fp/cmd/gombok

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
	return HashablePerson().Eqv(r, other)
}

func (r Person) Hashcode() uint32 {
	return HashablePerson().Hash(r)
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
	company string `column:"company"`
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

// @fp.Derive
var _ ord.Derives[fp.Ord[Car]]

func EqSortSeq[T any](eqT fp.Eq[T], ordT fp.Ord[T]) fp.Eq[fp.Seq[T]] {
	return eq.New(func(a, b fp.Seq[T]) bool {

		asorted := seq.Sort(a, ordT)

		bsorted := seq.Sort(a, ordT)

		return eq.Seq(eqT).Eqv(asorted, bsorted)
	})
}

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

// @fp.Deref
type MapEntry[K, V any] fp.Tuple2[K, V]

// @fp.Deref
type OptionalInt fp.Option[int]

// @fp.Deref
type OptionalStringer[T fmt.Stringer] fp.Option[T]
