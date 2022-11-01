// Code generated by gombok, DO NOT EDIT.
package testpk2

import (
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
	"github.com/csgura/fp/test/internal/testpk1"
	"time"
)

var EqPerson = eq.ContraMap(
	eq.Tuple8(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String), eq.Slice(eq.String), eq.HCons(eq.String, eq.HCons(eq.Given[int](), eq.HNil)), EqFpSeq(EqFloat64), eq.Bytes),
	Person.AsTuple,
)

var EqWallet = eq.ContraMap(
	eq.Tuple2(EqPerson, eq.Given[int64]()),
	Wallet.AsTuple,
)

func EqEntry[A comparable, B any, C fmt.Stringer, D interface {
	Hello() string
}](eqA fp.Eq[A], eqB fp.Eq[B], eqC fp.Eq[C], eqD fp.Eq[D]) fp.Eq[Entry[A, B, C, D]] {
	return eq.ContraMap(
		eq.Tuple3(eq.String, eqA, eq.Tuple2(eqA, eqB)),
		Entry[A, B, C, D].AsTuple,
	)
}

func MonoidEntry[A comparable, B any, C fmt.Stringer, D interface {
	Hello() string
}](monoidA fp.Monoid[A], monoidB fp.Monoid[B], monoidC fp.Monoid[C], monoidD fp.Monoid[D]) fp.Monoid[Entry[A, B, C, D]] {
	return monoid.IMap(
		monoid.Tuple3(monoid.String, monoidA, monoid.Tuple2(monoidA, monoidB)),
		fp.Compose(
			as.Curried2(EntryBuilder[A, B, C, D].FromTuple)(EntryBuilder[A, B, C, D]{}),
			EntryBuilder[A, B, C, D].Build,
		),
		Entry[A, B, C, D].AsTuple,
	)
}

var HashableKey = hash.ContraMap(
	hash.Tuple3(hash.Number[int](), hash.Number[float32](), hash.Bytes),
	Key.AsTuple,
)

var MonoidPoint = monoid.IMap(
	monoid.Tuple3(MonoidInt, MonoidInt, monoid.Tuple2(MonoidInt, MonoidInt)),
	fp.Compose(
		as.Curried2(PointBuilder.FromTuple)(PointBuilder{}),
		PointBuilder.Build,
	),
	Point.AsTuple,
)

var EqGreeting = eq.ContraMap(
	eq.Tuple2(EqTestpk1World, eq.String),
	Greeting.AsTuple,
)

var EncoderGreeting = js.EncoderContraMap(
	js.EncoderLabelled2(js.EncoderNamed[NameIsHello[testpk1.World]](testpk1.EncoderWorld), js.EncoderNamed[NameIsLanguage[string]](js.EncoderString)),
	Greeting.AsLabelled,
)

var DecoderGreeting = js.DecoderMap(
	js.DecoderLabelled2(js.DecoderNamed[NameIsHello[testpk1.World]](testpk1.DecoderWorld), js.DecoderNamed[NameIsLanguage[string]](js.DecoderString)),
	fp.Compose(
		as.Curried2(GreetingBuilder.FromLabelled)(GreetingBuilder{}),
		GreetingBuilder.Build,
	),
)

var EncoderThree = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NameIsOne[int]](js.EncoderNumber[int]()),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NameIsTwo[string]](js.EncoderString),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NameIsThree[float64]](js.EncoderNumber[float64]()),
				js.EncoderHNil,
			),
		),
	),
	fp.Compose(
		Three.AsLabelled,
		as.HList3Labelled[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]],
	),
)

var DecoderThree = js.DecoderMap(
	js.DecoderHConsLabelled(
		js.DecoderNamed[NameIsOne[int]](js.DecoderNumber[int]()),
		js.DecoderHConsLabelled(
			js.DecoderNamed[NameIsTwo[string]](js.DecoderString),
			js.DecoderHConsLabelled(
				js.DecoderNamed[NameIsThree[float64]](js.DecoderNumber[float64]()),
				js.DecoderHNil,
			),
		),
	),

	fp.Compose(
		as.Func2(
			hlist.Case3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64], hlist.Nil, fp.Labelled3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]]],
		).ApplyLast(
			as.Labelled3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]],
		),
		fp.Compose(
			as.Curried2(ThreeBuilder.FromLabelled)(ThreeBuilder{}),
			ThreeBuilder.Build,
		),
	),
)

var ShowThree = show.Generic(
	as.Generic(
		"testpk2.Three",
		fp.Compose(
			Three.AsTuple,
			as.HList3[int, string, float64],
		),

		fp.Compose(
			as.Func2(
				hlist.Case3[int, string, float64, hlist.Nil, fp.Tuple3[int, string, float64]],
			).ApplyLast(
				as.Tuple3[int, string, float64],
			),
			fp.Compose(
				as.Curried2(ThreeBuilder.FromTuple)(ThreeBuilder{}),
				ThreeBuilder.Build,
			),
		),
	),
	show.HCons(
		show.Given[int](),
		show.HCons(
			show.String,
			show.HCons(
				show.Given[float64](),
				show.HNil,
			),
		),
	),
)

var ReadThree = read.Generic(
	as.Generic(
		"testpk2.Three",
		fp.Compose(
			Three.AsTuple,
			as.HList3[int, string, float64],
		),

		fp.Compose(
			as.Func2(
				hlist.Case3[int, string, float64, hlist.Nil, fp.Tuple3[int, string, float64]],
			).ApplyLast(
				as.Tuple3[int, string, float64],
			),
			fp.Compose(
				as.Curried2(ThreeBuilder.FromTuple)(ThreeBuilder{}),
				ThreeBuilder.Build,
			),
		),
	),
	read.HCons(
		read.Int[int](),
		read.HCons(
			read.String,
			read.HCons(
				read.Float[float64](),
				read.HNil,
			),
		),
	),
)

var EqTestpk1World = eq.ContraMap(
	eq.Tuple2(eq.String, eq.Given[time.Time]()),
	testpk1.World.AsTuple,
)

func EqTestpk1Wrapper[T any](eqT fp.Eq[T]) fp.Eq[testpk1.Wrapper[T]] {
	return eq.ContraMap(
		eq.Tuple1(eqT),
		testpk1.Wrapper[T].AsTuple,
	)
}
