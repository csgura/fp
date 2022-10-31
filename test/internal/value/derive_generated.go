// Code generated by gombok, DO NOT EDIT.
package value

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/hello"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
)

var EqPerson = eq.ContraMap(
	eq.Tuple8(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String), eq.Slice(eq.String), eq.HCons(eq.String, eq.HCons(eq.Given[int](), eq.HNil)), EqFpSeq(EqFloat64), eq.Bytes),
	Person.AsTuple,
)

var EqWallet = eq.ContraMap(
	eq.Tuple2(EqPerson, eq.Given[int64]()),
	Wallet.AsTuple,
)

func EqEntry[A interface {
	String() string
}, B any](eqA fp.Eq[A], eqB fp.Eq[B]) fp.Eq[Entry[A, B]] {
	return eq.ContraMap(
		eq.Tuple3(eq.String, eqA, eq.Tuple2(eqA, eqB)),
		Entry[A, B].AsTuple,
	)
}

func MonoidEntry[A interface {
	String() string
}, B any](monoidA fp.Monoid[A], monoidB fp.Monoid[B]) fp.Monoid[Entry[A, B]] {
	return monoid.IMap(
		monoid.Tuple3(monoid.String, monoidA, monoid.Tuple2(monoidA, monoidB)),
		fp.Compose(
			as.Curried2(EntryBuilder[A, B].FromTuple)(EntryBuilder[A, B]{}),
			EntryBuilder[A, B].Build,
		),
		Entry[A, B].AsTuple,
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
	eq.Tuple2(hello.EqWorld, eq.String),
	Greeting.AsTuple,
)

var EncoderGreeting = js.EncoderContraMap(
	js.EncoderLabelled2(js.EncoderNamed[NameIsHello[hello.World]](hello.EncoderWorld), js.EncoderNamed[NameIsLanguage[string]](js.EncoderString)),
	Greeting.AsLabelled,
)

var DecoderGreeting = js.DecoderMap(
	js.DecoderLabelled2(js.DecoderNamed[NameIsHello[hello.World]](hello.DecoderWorld), js.DecoderNamed[NameIsLanguage[string]](js.DecoderString)),
	fp.Compose(
		as.Curried2(GreetingBuilder.FromLabelled)(GreetingBuilder{}),
		GreetingBuilder.Build,
	),
)

var EncoderThree = js.EncoderContraMap(
	js.EncoderContraMap(
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
		as.HList3Labelled[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]],
	),
	Three.AsLabelled,
)

var DecoderThree = js.DecoderMap(
	js.DecoderMap(
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
		as.Func2(
			hlist.Case3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64], hlist.Nil, fp.Labelled3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]]],
		).ApplyLast(
			as.Labelled3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]],
		),
	),
	fp.Compose(
		as.Curried2(ThreeBuilder.FromLabelled)(ThreeBuilder{}),
		ThreeBuilder.Build,
	),
)

var ShowThree = show.Generic(
	as.Generic(
		Three.AsTuple,
		fp.Compose(
			as.Curried2(ThreeBuilder.FromTuple)(ThreeBuilder{}),
			ThreeBuilder.Build,
		),
	),
	show.ContraMap(
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
		as.HList3[int, string, float64],
	),
)

var ReadThree = read.Generic(
	as.Generic(
		Three.AsTuple,
		fp.Compose(
			as.Curried2(ThreeBuilder.FromTuple)(ThreeBuilder{}),
			ThreeBuilder.Build,
		),
	),
	read.Map(
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
		as.Func2(
			hlist.Case3[int, string, float64, hlist.Nil, fp.Tuple3[int, string, float64]],
		).ApplyLast(
			as.Tuple3[int, string, float64],
		),
	),
)
