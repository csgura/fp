package testpk1

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.Json
// @fp.GenLabelled
type World struct {
	message   string
	timestamp time.Time
}

// @fp.Derive
var _ eq.Derives[fp.Eq[World]]

// @fp.Derive
var _ js.Derives[js.Encoder[World]]

// @fp.Derive
var _ js.Derives[js.Decoder[World]]

// @fp.Derive
var _ show.Derives[fp.Show[World]]

// @fp.Value
// @fp.GenLabelled
type HasOption struct {
	message  string
	addr     fp.Option[string]
	phone    []string
	emptySeq []int
}

// @fp.Derive
var _ js.Derives[js.Encoder[HasOption]]

// @fp.Value
type CustomValue struct {
	a string
	b int
}

func (r CustomValue) A() string {
	return "hello" + r.a
}

func (r CustomValue) WithB(v int) CustomValue {
	if v > 0 {
		r.b = v
	}
	return r
}

type CustomValueBuilder CustomValue

func (r CustomValueBuilder) B(v int) CustomValueBuilder {
	if v > 0 {
		r.b = v
	}
	return r
}

// // @fp.Value
// type NotDerivable struct {
// 	a int
// 	b interface {
// 		Hello() string
// 	}
// }

// // @fp.Derive
// var _ monoid.Derives[fp.Monoid[NotDerivable]]

// var MonoidInt = monoid.Sum[int]()

// @fp.Value
type AliasedStruct World

// @fp.Derive
var _ eq.Derives[fp.Eq[AliasedStruct]]

// @fp.Value
type HListInsideHList struct {
	tp    fp.Tuple2[string, int]
	value string
	hello World
}

// @fp.Derive
var _ show.Derives[fp.Show[HListInsideHList]]

// @fp.Derive
var _ read.Derives[read.Read[HListInsideHList]]

// @fp.Derive
var _ read.Derives[read.Read[World]]

// @fp.Value
type Wrapper[T any] struct {
	unwrap T
}

// @fp.Value
type TestOrderedEq struct {
	list  fp.Seq[int]
	tlist fp.Seq[fp.Tuple2[int, int]]
}

func EqSeq[T any](eqT fp.Eq[T], ordT fp.Ord[T]) fp.Eq[fp.Seq[T]] {
	return eq.New(func(a, b fp.Seq[T]) bool {
		asorted := seq.Sort(a, ordT)
		bsorted := seq.Sort(b, ordT)
		return eq.Seq(eqT).Eqv(asorted, bsorted)
	})
}

// @fp.ImportGiven
var _ ord.Derives[fp.Ord[any]]

// @fp.Derive
var _ eq.Derives[fp.Eq[TestOrderedEq]]

// @fp.Value
type MapEq struct {
	m  map[string]World
	m2 fp.Map[string, World]
}

// @fp.Derive
var _ eq.Derives[fp.Eq[MapEq]]

// @fp.Value
type SeqMonoid struct {
	v  string
	s  fp.Seq[string]
	m  map[string]int
	m2 fp.Map[string, World]
}

// @fp.Derive
var _ monoid.Derives[fp.Monoid[SeqMonoid]]

type MyInt int

type MySeq[T any] fp.Seq[T]

// @fp.Derive
var _ eq.Derives[fp.Eq[MyInt]]

// @fp.Derive
var _ eq.Derives[fp.Eq[MySeq[any]]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[MySeq[any]]]

// @fp.Value
type MapEqParam[K, V any] struct {
	m fp.Map[K, V]
}

// @fp.Derive
var _ eq.Derives[fp.Eq[MapEqParam[any, any]]]
