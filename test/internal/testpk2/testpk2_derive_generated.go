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
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
	"github.com/csgura/fp/test/internal/testpk1"
)

func EqPerson() fp.Eq[Person] {
	return eq.ContraMap(
		eq.Tuple8(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String), eq.Slice(eq.String), eq.HCons(eq.String, eq.HCons(eq.Given[int](), eq.HNil)), EqFpSeq(EqFloat64), eq.Bytes),
		Person.AsTuple,
	)
}

func EqWallet() fp.Eq[Wallet] {
	return eq.ContraMap(
		eq.Tuple2(EqPerson(), eq.Given[int64]()),
		Wallet.AsTuple,
	)
}

func EqEntry[A comparable, B any, C fmt.Stringer, D interface {
	Hello() string
}](eqA fp.Eq[A], eqB fp.Eq[B]) fp.Eq[Entry[A, B, C, D]] {
	return eq.ContraMap(
		eq.Tuple3(eq.String, eqA, eq.Tuple2(eqA, eqB)),
		Entry[A, B, C, D].AsTuple,
	)
}

func MonoidEntry[A comparable, B any, C fmt.Stringer, D interface {
	Hello() string
}](monoidA fp.Monoid[A], monoidB fp.Monoid[B]) fp.Monoid[Entry[A, B, C, D]] {
	return monoid.IMap(
		monoid.Tuple3(monoid.String, monoidA, monoid.Tuple2(monoidA, monoidB)),
		fp.Compose(
			as.Curried2(EntryBuilder[A, B, C, D].FromTuple)(EntryBuilder[A, B, C, D]{}),
			EntryBuilder[A, B, C, D].Build,
		),
		Entry[A, B, C, D].AsTuple,
	)
}

func HashableKey() fp.Hashable[Key] {
	return hash.ContraMap(
		hash.Tuple3(hash.Number[int](), hash.Number[float32](), hash.Bytes),
		Key.AsTuple,
	)
}

func MonoidPoint() fp.Monoid[Point] {
	return monoid.IMap(
		monoid.Tuple3(MonoidInt, MonoidInt, monoid.Tuple2(MonoidInt, MonoidInt)),
		fp.Compose(
			as.Curried2(PointBuilder.FromTuple)(PointBuilder{}),
			PointBuilder.Build,
		),
		Point.AsTuple,
	)
}

func EqGreeting() fp.Eq[Greeting] {
	return eq.ContraMap(
		eq.Tuple2(EqTestpk1World(), eq.String),
		Greeting.AsTuple,
	)
}

func EncoderGreeting() js.Encoder[Greeting] {
	return js.EncoderContraMap(
		js.EncoderLabelled2(js.EncoderNamed[NamedHelloOfGreeting](EncoderTestpk1World()), js.EncoderNamed[NamedLanguageOfGreeting](js.EncoderString)),
		Greeting.AsLabelled,
	)
}

func DecoderGreeting() js.Decoder[Greeting] {
	return js.DecoderMap(
		js.DecoderLabelled2(js.DecoderNamed[NamedHelloOfGreeting](testpk1.DecoderWorld()), js.DecoderNamed[NamedLanguageOfGreeting](js.DecoderString)),
		fp.Compose(
			as.Curried2(GreetingBuilder.FromLabelled)(GreetingBuilder{}),
			GreetingBuilder.Build,
		),
	)
}

func EncoderThree() js.Encoder[Three] {
	return js.EncoderContraMap(
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedOneOfThree](js.EncoderNumber[int]()),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NamedTwoOfThree](js.EncoderString),
				js.EncoderHConsLabelled(
					js.EncoderNamed[NamedThreeOfThree](js.EncoderNumber[float64]()),
					js.EncoderHNil,
				),
			),
		),
		fp.Compose(
			Three.AsLabelled,
			as.HList3Labelled,
		),
	)
}

func DecoderThree() js.Decoder[Three] {
	return js.DecoderMap(
		js.DecoderHConsLabelled(
			js.DecoderNamed[NamedOneOfThree](js.DecoderNumber[int]()),
			js.DecoderHConsLabelled(
				js.DecoderNamed[NamedTwoOfThree](js.DecoderString),
				js.DecoderHConsLabelled(
					js.DecoderNamed[NamedThreeOfThree](js.DecoderNumber[float64]()),
					js.DecoderHNil,
				),
			),
		),

		fp.Compose(
			product.LabelledFromHList3,
			fp.Compose(
				as.Curried2(ThreeBuilder.FromLabelled)(ThreeBuilder{}),
				ThreeBuilder.Build,
			),
		),
	)
}

func ShowThree() fp.Show[Three] {
	return show.Generic(
		fp.Generic[Three, hlist.Cons[int, hlist.Cons[string, hlist.Cons[float64, hlist.Nil]]]]{
			Type: "testpk2.Three",
			Kind: "Struct",
			To: fp.Compose(
				Three.AsTuple,
				as.HList3[int, string, float64],
			),
			From: fp.Compose(
				product.TupleFromHList3,
				fp.Compose(
					as.Curried2(ThreeBuilder.FromTuple)(ThreeBuilder{}),
					ThreeBuilder.Build,
				),
			),
		},
		show.StructHCons(
			show.Int[int](),
			show.StructHCons(
				show.String,
				show.StructHCons(
					show.Number[float64](),
					show.HNil,
				),
			),
		),
	)
}

func ReadThree() read.Read[Three] {
	return read.Generic(
		fp.Generic[Three, hlist.Cons[int, hlist.Cons[string, hlist.Cons[float64, hlist.Nil]]]]{
			Type: "testpk2.Three",
			Kind: "Struct",
			To: fp.Compose(
				Three.AsTuple,
				as.HList3[int, string, float64],
			),
			From: fp.Compose(
				product.TupleFromHList3,
				fp.Compose(
					as.Curried2(ThreeBuilder.FromTuple)(ThreeBuilder{}),
					ThreeBuilder.Build,
				),
			),
		},
		read.TupleHCons(
			read.Int[int](),
			read.TupleHCons(
				read.String,
				read.TupleHCons(
					read.Float[float64](),
					read.TupleHNil,
				),
			),
		),
	)
}

func EqTestpk1World() fp.Eq[testpk1.World] {
	return eq.ContraMap(
		eq.Tuple3(eq.String, eq.Time, eq.String),
		testpk1.World.AsTuple,
	)
}

func EqTestpk1Wrapper[T any](eqT fp.Eq[T]) fp.Eq[testpk1.Wrapper[T]] {
	return eq.ContraMap(
		eq.Tuple1(eqT),
		testpk1.Wrapper[T].AsTuple,
	)
}

func EqTree() fp.Eq[Tree] {
	return eq.ContraMap(
		eq.Tuple1(testpk1.EqNode()),
		Tree.AsTuple,
	)
}

func EncoderTestpk1World() js.Encoder[testpk1.World] {
	return js.EncoderContraMap(
		js.EncoderHConsLabelled(
			js.EncoderNamed[testpk1.NamedMessageOfWorld](js.EncoderString),
			js.EncoderHConsLabelled(
				js.EncoderNamed[testpk1.NamedTimestampOfWorld](js.EncoderTime),
				js.EncoderHConsLabelled(
					js.EncoderNamed[testpk1.PubNamedPubOfWorld](js.EncoderString),
					js.EncoderHNil,
				),
			),
		),
		fp.Compose(
			testpk1.World.AsLabelled,
			as.HList3Labelled,
		),
	)
}
