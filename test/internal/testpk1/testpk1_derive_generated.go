// Code generated by gombok, DO NOT EDIT.
package testpk1

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
	"time"
)

var EqWorld = eq.ContraMap(
	eq.Tuple2(eq.String, eq.Given[time.Time]()),
	World.AsTuple,
)

var EncoderWorld = js.EncoderContraMap(
	js.EncoderLabelled2(js.EncoderNamed[NameIsMessage[string]](js.EncoderString), js.EncoderNamed[NameIsTimestamp[time.Time]](js.EncoderTime)),
	World.AsLabelled,
)

var DecoderWorld = js.DecoderMap(
	js.DecoderLabelled2(js.DecoderNamed[NameIsMessage[string]](js.DecoderString), js.DecoderNamed[NameIsTimestamp[time.Time]](js.DecoderTime)),
	fp.Compose(
		as.Curried2(WorldBuilder.FromLabelled)(WorldBuilder{}),
		WorldBuilder.Build,
	),
)

var ShowWorld = show.Generic(
	as.Generic(
		"testpk1.World",
		fp.Compose(
			World.AsTuple,
			as.HList2[string, time.Time],
		),

		fp.Compose(
			product.TupleFromHList2[string, time.Time],
			fp.Compose(
				as.Curried2(WorldBuilder.FromTuple)(WorldBuilder{}),
				WorldBuilder.Build,
			),
		),
	),
	show.HCons(
		show.String,
		show.HCons(
			show.Time,
			show.HNil,
		),
	),
)

var EncoderHasOption = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NameIsMessage[string]](js.EncoderString),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NameIsAddr[fp.Option[string]]](js.EncoderOption(js.EncoderString)),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NameIsPhone[[]string]](js.EncoderSlice(js.EncoderString)),
				js.EncoderHConsLabelled(
					js.EncoderNamed[NameIsEmptySeq[[]int]](js.EncoderSlice(js.EncoderNumber[int]())),
					js.EncoderHNil,
				),
			),
		),
	),
	fp.Compose(
		HasOption.AsLabelled,
		as.HList4Labelled[NameIsMessage[string], NameIsAddr[fp.Option[string]], NameIsPhone[[]string], NameIsEmptySeq[[]int]],
	),
)

var EqAliasedStruct = eq.ContraMap(
	eq.Tuple2(eq.String, eq.Given[time.Time]()),
	AliasedStruct.AsTuple,
)

var ShowHListInsideHList = show.Generic(
	as.Generic(
		"testpk1.HListInsideHList",
		fp.Compose(
			HListInsideHList.AsTuple,
			as.HList3[fp.Tuple2[string, int], string, World],
		),

		fp.Compose(
			product.TupleFromHList3[fp.Tuple2[string, int], string, World],
			fp.Compose(
				as.Curried2(HListInsideHListBuilder.FromTuple)(HListInsideHListBuilder{}),
				HListInsideHListBuilder.Build,
			),
		),
	),
	show.HCons(
		show.Generic(
			as.Generic(
				"",
				as.HList2[string, int],
				product.TupleFromHList2[string, int],
			),
			show.HCons(
				show.String,
				show.HCons(
					show.Int[int](),
					show.HNil,
				),
			),
		),
		show.HCons(
			show.String,
			show.HCons(
				ShowWorld,
				show.HNil,
			),
		),
	),
)

var ReadHListInsideHList = read.Generic(
	as.Generic(
		"testpk1.HListInsideHList",
		fp.Compose(
			HListInsideHList.AsTuple,
			as.HList3[fp.Tuple2[string, int], string, World],
		),

		fp.Compose(
			product.TupleFromHList3[fp.Tuple2[string, int], string, World],
			fp.Compose(
				as.Curried2(HListInsideHListBuilder.FromTuple)(HListInsideHListBuilder{}),
				HListInsideHListBuilder.Build,
			),
		),
	),
	read.HCons(
		read.Generic(
			as.Generic(
				"",
				as.HList2[string, int],
				product.TupleFromHList2[string, int],
			),
			read.HCons(
				read.String,
				read.HCons(
					read.Int[int](),
					read.HNil,
				),
			),
		),
		read.HCons(
			read.String,
			read.HCons(
				ReadWorld,
				read.HNil,
			),
		),
	),
)

var ReadWorld = read.Generic(
	as.Generic(
		"testpk1.World",
		fp.Compose(
			World.AsTuple,
			as.HList2[string, time.Time],
		),

		fp.Compose(
			product.TupleFromHList2[string, time.Time],
			fp.Compose(
				as.Curried2(WorldBuilder.FromTuple)(WorldBuilder{}),
				WorldBuilder.Build,
			),
		),
	),
	read.HCons(
		read.String,
		read.HCons(
			read.Time,
			read.HNil,
		),
	),
)

var EqTestOrderedEq = eq.ContraMap(
	eq.Tuple2(EqSeq(eq.Given[int](), ord.Given[int]()), EqSeq(eq.Tuple2(eq.Given[int](), eq.Given[int]()), ord.Tuple2(ord.Given[int](), ord.Given[int]()))),
	TestOrderedEq.AsTuple,
)

var EqMapEq = eq.ContraMap(
	eq.Tuple2(eq.GoMap[string, World](EqWorld), eq.FpMap[string, World](EqWorld)),
	MapEq.AsTuple,
)

var MonoidSeqMonoid = monoid.IMap(
	monoid.Tuple4(monoid.String, monoid.MergeSeq[string](), monoid.MergeGoMap[string, int](), monoid.MergeMap[string, World]()),
	fp.Compose(
		as.Curried2(SeqMonoidBuilder.FromTuple)(SeqMonoidBuilder{}),
		SeqMonoidBuilder.Build,
	),
	SeqMonoid.AsTuple,
)

var EqMyInt = eq.ContraMap(
	eq.Given[int](),
	func(v MyInt) int {
		return int(v)
	},
)

func EqMySeq[T any](eqT fp.Eq[T]) fp.Eq[MySeq[T]] {
	return eq.ContraMap(
		eq.Slice(eqT),
		func(v MySeq[T]) []T {
			return []T(v)
		},
	)
}

func MonoidMySeq[T any]() fp.Monoid[MySeq[T]] {
	return monoid.IMap(
		monoid.MergeSlice[T](),
		func(v []T) MySeq[T] {
			return MySeq[T](v)
		},
		func(v MySeq[T]) []T {
			return []T(v)
		},
	)
}
