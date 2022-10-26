package value

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/hello"
	"github.com/csgura/fp/test/internal/js"
)

var EqPerson = eq.ContraMap(eq.Tuple8(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String), eq.Slice(eq.String), eq.HCons(eq.String, eq.HCons(eq.Given[int](), eq.HNil)), EqFpSeq(EqFloat64), eq.Bytes), Person.AsTuple)

var EqWallet = eq.ContraMap(eq.Tuple2(EqPerson, eq.Given[int64]()), Wallet.AsTuple)

func EqEntry[A interface{ String() string }, B any](eqA fp.Eq[A], eqB fp.Eq[B]) fp.Eq[Entry[A, B]] {
	return eq.ContraMap(eq.Tuple3(eq.String, eqA, eq.Tuple2(eqA, eqB)), Entry[A, B].AsTuple)
}

func MonoidEntry[A interface{ String() string }, B any](monoidA fp.Monoid[A], monoidB fp.Monoid[B]) fp.Monoid[Entry[A, B]] {
	return monoid.IMap(monoid.Tuple3(monoid.String, monoidA, monoid.Tuple2(monoidA, monoidB)), fp.Compose(
		as.Curried2(EntryBuilder[A, B].FromTuple)(EntryBuilder[A, B]{}), EntryBuilder[A, B].Build),
		Entry[A, B].AsTuple)
}

var HashableKey = hash.ContraMap(hash.Tuple3(hash.Number[int](), hash.Number[float32](), hash.Bytes), Key.AsTuple)

var MonoidPoint = monoid.IMap(monoid.Tuple3(MonoidInt, MonoidInt, monoid.Tuple2(MonoidInt, MonoidInt)), fp.Compose(
	as.Curried2(PointBuilder.FromTuple)(PointBuilder{}), PointBuilder.Build),
	Point.AsTuple)

var EqGreeting = eq.ContraMap(eq.Tuple2(hello.EqWorld, eq.String), Greeting.AsTuple)

var EncoderGreeting = js.ContraMap(js.Labelled2(hello.EncoderWorld, js.String), Greeting.AsLabelled)
