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
	message    string
	timestamp  time.Time
	Pub        string
	_notExport string
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

// @fp.Value
type NotUsedProblem struct {
	m MapEqParam[string, int]
}

// @fp.Derive
var _ eq.Derives[fp.Eq[NotUsedProblem]]

// @fp.Value
type Node struct {
	value string
	left  *Node
	right *Node
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Node]]

// @fp.Value
type NoPrivate struct {
	Value int
}

// @fp.Derive
var _ eq.Derives[fp.Eq[NoPrivate]]

// @fp.Value
// @fp.GenLabelled
type Over21 struct {
	i1  int
	i2  int
	i3  int
	i4  int
	i5  int
	i6  int
	i7  int
	i8  int
	i9  int
	i10 int

	i11 int
	i12 int
	i13 int
	i14 int
	i15 int
	i16 int
	i17 int
	i18 int
	i19 int
	i20 int

	i21 int
	i22 int
	i23 int
	i24 int
	i25 int
	i26 int
	i27 int
	i28 int
	i29 int
	i30 int
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Over21]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[Over21]]

// @fp.Derive
var _ read.Derives[read.Read[Over21]]

// @fp.Derive
var _ js.Derives[js.Encoder[Over21]]

// @fp.Derive
var _ js.Derives[js.Decoder[Over21]]

// @fp.Value
type DefinedOtherPackage struct {
	PubField  string
	privField string
	DupGetter string
}

func (r *DefinedOtherPackage) PtrRecv() {
	// do nothing
}

func (r *DefinedOtherPackage) PtrRecvRet() string {
	// do nothing
	return ""
}

func (r *DefinedOtherPackage) GetDupGetter() string {
	return "dup"
}
