// Code generated by gombok, DO NOT EDIT.
package testpk1

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
	"time"
)

var EqWorld = eq.ContraMap(
	eq.Tuple3(eq.String, eq.Given[time.Time](), eq.String),
	World.AsTuple,
)

var EncoderWorld = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NamedMessage[string]](js.EncoderString),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedTimestamp[time.Time]](js.EncoderTime),
			js.EncoderHConsLabelled(
				js.EncoderNamed[PubNamedPub[string]](js.EncoderString),
				js.EncoderHNil,
			),
		),
	),
	fp.Compose(
		World.AsLabelled,
		as.HList3Labelled[NamedMessage[string], NamedTimestamp[time.Time], PubNamedPub[string]],
	),
)

var DecoderWorld = js.DecoderMap(
	js.DecoderHConsLabelled(
		js.DecoderNamed[NamedMessage[string]](js.DecoderString),
		js.DecoderHConsLabelled(
			js.DecoderNamed[NamedTimestamp[time.Time]](js.DecoderTime),
			js.DecoderHConsLabelled(
				js.DecoderNamed[PubNamedPub[string]](js.DecoderString),
				js.DecoderHNil,
			),
		),
	),

	fp.Compose(
		product.LabelledFromHList3[NamedMessage[string], NamedTimestamp[time.Time], PubNamedPub[string]],
		fp.Compose(
			as.Curried2(WorldBuilder.FromLabelled)(WorldBuilder{}),
			WorldBuilder.Build,
		),
	),
)

var ShowWorld = show.Generic(
	as.Generic(
		"testpk1.World",
		"Struct",
		fp.Compose(
			World.AsTuple,
			as.HList3[string, time.Time, string],
		),

		fp.Compose(
			product.TupleFromHList3[string, time.Time, string],
			fp.Compose(
				as.Curried2(WorldBuilder.FromTuple)(WorldBuilder{}),
				WorldBuilder.Build,
			),
		),
	),
	show.StructHCons(
		show.String,
		show.StructHCons(
			show.Time,
			show.StructHCons(
				show.String,
				show.HNil,
			),
		),
	),
)

var EncoderHasOption = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NamedMessage[string]](js.EncoderString),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedAddr[fp.Option[string]]](js.EncoderOption(js.EncoderString)),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NamedPhone[[]string]](js.EncoderSlice(js.EncoderString)),
				js.EncoderHConsLabelled(
					js.EncoderNamed[NamedEmptySeq[[]int]](js.EncoderSlice(js.EncoderNumber[int]())),
					js.EncoderHNil,
				),
			),
		),
	),
	fp.Compose(
		HasOption.AsLabelled,
		as.HList4Labelled[NamedMessage[string], NamedAddr[fp.Option[string]], NamedPhone[[]string], NamedEmptySeq[[]int]],
	),
)

var EqAliasedStruct = eq.ContraMap(
	eq.Tuple3(eq.String, eq.Given[time.Time](), eq.String),
	AliasedStruct.AsTuple,
)

var ShowHListInsideHList = show.Generic(
	as.Generic(
		"testpk1.HListInsideHList",
		"Struct",
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
	show.StructHCons(
		show.Generic(
			as.Generic(
				"fp.Tuple2",
				"Tuple",
				as.HList2[string, int],
				product.TupleFromHList2[string, int],
			),
			show.TupleHCons(
				show.String,
				show.TupleHCons(
					show.Int[int](),
					show.HNil,
				),
			),
		),
		show.StructHCons(
			show.String,
			show.StructHCons(
				ShowWorld,
				show.HNil,
			),
		),
	),
)

var ReadHListInsideHList = read.Generic(
	as.Generic(
		"testpk1.HListInsideHList",
		"Struct",
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
	read.TupleHCons(
		read.Generic(
			as.Generic(
				"fp.Tuple2",
				"Tuple",
				as.HList2[string, int],
				product.TupleFromHList2[string, int],
			),
			read.TupleHCons(
				read.String,
				read.TupleHCons(
					read.Int[int](),
					read.TupleHNill,
				),
			),
		),
		read.TupleHCons(
			read.String,
			read.TupleHCons(
				ReadWorld,
				read.TupleHNill,
			),
		),
	),
)

var ReadWorld = read.Generic(
	as.Generic(
		"testpk1.World",
		"Struct",
		fp.Compose(
			World.AsTuple,
			as.HList3[string, time.Time, string],
		),

		fp.Compose(
			product.TupleFromHList3[string, time.Time, string],
			fp.Compose(
				as.Curried2(WorldBuilder.FromTuple)(WorldBuilder{}),
				WorldBuilder.Build,
			),
		),
	),
	read.TupleHCons(
		read.String,
		read.TupleHCons(
			read.Time,
			read.TupleHCons(
				read.String,
				read.TupleHNill,
			),
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

func EqMapEqParam[K any, V any](eqV fp.Eq[V]) fp.Eq[MapEqParam[K, V]] {
	return eq.ContraMap(
		eq.Tuple1(eq.FpMap[K, V](eqV)),
		MapEqParam[K, V].AsTuple,
	)
}

var EqNotUsedProblem = eq.ContraMap(
	eq.Tuple1(EqMapEqParam[string, int](eq.Given[int]())),
	NotUsedProblem.AsTuple,
)

func EqNode() fp.Eq[Node] {
	return eq.ContraMap(
		eq.Tuple3(eq.String, eq.Ptr(lazy.Call(func() fp.Eq[Node] {
			return EqNode()
		})), eq.Ptr(lazy.Call(func() fp.Eq[Node] {
			return EqNode()
		}))),
		Node.AsTuple,
	)
}

var EqNoPrivate = eq.ContraMap(
	eq.Tuple1(eq.Given[int]()),
	NoPrivate.AsTuple,
)

var EqOver21 = eq.ContraMap(
	eq.HCons(
		eq.Given[int](),
		eq.HCons(
			eq.Given[int](),
			eq.HCons(
				eq.Given[int](),
				eq.HCons(
					eq.Given[int](),
					eq.HCons(
						eq.Given[int](),
						eq.HCons(
							eq.Given[int](),
							eq.HCons(
								eq.Given[int](),
								eq.HCons(
									eq.Given[int](),
									eq.HCons(
										eq.Given[int](),
										eq.HCons(
											eq.Given[int](),
											eq.HCons(
												eq.Given[int](),
												eq.HCons(
													eq.Given[int](),
													eq.HCons(
														eq.Given[int](),
														eq.HCons(
															eq.Given[int](),
															eq.HCons(
																eq.Given[int](),
																eq.HCons(
																	eq.Given[int](),
																	eq.HCons(
																		eq.Given[int](),
																		eq.HCons(
																			eq.Given[int](),
																			eq.HCons(
																				eq.Given[int](),
																				eq.HCons(
																					eq.Given[int](),
																					eq.HCons(
																						eq.Given[int](),
																						eq.HCons(
																							eq.Given[int](),
																							eq.HCons(
																								eq.Given[int](),
																								eq.HCons(
																									eq.Given[int](),
																									eq.HCons(
																										eq.Given[int](),
																										eq.HCons(
																											eq.Given[int](),
																											eq.HCons(
																												eq.Given[int](),
																												eq.HCons(
																													eq.Given[int](),
																													eq.HCons(
																														eq.Given[int](),
																														eq.HCons(
																															eq.Given[int](),
																															eq.HNil,
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	),
	func(v Over21) hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]] {
		i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29 := v.Unapply()
		return hlist.Concat(i0,
			hlist.Concat(i1,
				hlist.Concat(i2,
					hlist.Concat(i3,
						hlist.Concat(i4,
							hlist.Concat(i5,
								hlist.Concat(i6,
									hlist.Concat(i7,
										hlist.Concat(i8,
											hlist.Concat(i9,
												hlist.Concat(i10,
													hlist.Concat(i11,
														hlist.Concat(i12,
															hlist.Concat(i13,
																hlist.Concat(i14,
																	hlist.Concat(i15,
																		hlist.Concat(i16,
																			hlist.Concat(i17,
																				hlist.Concat(i18,
																					hlist.Concat(i19,
																						hlist.Concat(i20,
																							hlist.Concat(i21,
																								hlist.Concat(i22,
																									hlist.Concat(i23,
																										hlist.Concat(i24,
																											hlist.Concat(i25,
																												hlist.Concat(i26,
																													hlist.Concat(i27,
																														hlist.Concat(i28,
																															hlist.Concat(i29,
																																hlist.Empty(),
																															),
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		)
	},
)

var MonoidOver21 = monoid.IMap(
	monoid.HCons(
		monoid.Product[int](),
		monoid.HCons(
			monoid.Product[int](),
			monoid.HCons(
				monoid.Product[int](),
				monoid.HCons(
					monoid.Product[int](),
					monoid.HCons(
						monoid.Product[int](),
						monoid.HCons(
							monoid.Product[int](),
							monoid.HCons(
								monoid.Product[int](),
								monoid.HCons(
									monoid.Product[int](),
									monoid.HCons(
										monoid.Product[int](),
										monoid.HCons(
											monoid.Product[int](),
											monoid.HCons(
												monoid.Product[int](),
												monoid.HCons(
													monoid.Product[int](),
													monoid.HCons(
														monoid.Product[int](),
														monoid.HCons(
															monoid.Product[int](),
															monoid.HCons(
																monoid.Product[int](),
																monoid.HCons(
																	monoid.Product[int](),
																	monoid.HCons(
																		monoid.Product[int](),
																		monoid.HCons(
																			monoid.Product[int](),
																			monoid.HCons(
																				monoid.Product[int](),
																				monoid.HCons(
																					monoid.Product[int](),
																					monoid.HCons(
																						monoid.Product[int](),
																						monoid.HCons(
																							monoid.Product[int](),
																							monoid.HCons(
																								monoid.Product[int](),
																								monoid.HCons(
																									monoid.Product[int](),
																									monoid.HCons(
																										monoid.Product[int](),
																										monoid.HCons(
																											monoid.Product[int](),
																											monoid.HCons(
																												monoid.Product[int](),
																												monoid.HCons(
																													monoid.Product[int](),
																													monoid.HCons(
																														monoid.Product[int](),
																														monoid.HCons(
																															monoid.Product[int](),
																															monoid.HNil,
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	),
	func(hl0 hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]) Over21 {
		i0, hl1 := hlist.Unapply(hl0)
		i1, hl2 := hlist.Unapply(hl1)
		i2, hl3 := hlist.Unapply(hl2)
		i3, hl4 := hlist.Unapply(hl3)
		i4, hl5 := hlist.Unapply(hl4)
		i5, hl6 := hlist.Unapply(hl5)
		i6, hl7 := hlist.Unapply(hl6)
		i7, hl8 := hlist.Unapply(hl7)
		i8, hl9 := hlist.Unapply(hl8)
		i9, hl10 := hlist.Unapply(hl9)
		i10, hl11 := hlist.Unapply(hl10)
		i11, hl12 := hlist.Unapply(hl11)
		i12, hl13 := hlist.Unapply(hl12)
		i13, hl14 := hlist.Unapply(hl13)
		i14, hl15 := hlist.Unapply(hl14)
		i15, hl16 := hlist.Unapply(hl15)
		i16, hl17 := hlist.Unapply(hl16)
		i17, hl18 := hlist.Unapply(hl17)
		i18, hl19 := hlist.Unapply(hl18)
		i19, hl20 := hlist.Unapply(hl19)
		i20, hl21 := hlist.Unapply(hl20)
		i21, hl22 := hlist.Unapply(hl21)
		i22, hl23 := hlist.Unapply(hl22)
		i23, hl24 := hlist.Unapply(hl23)
		i24, hl25 := hlist.Unapply(hl24)
		i25, hl26 := hlist.Unapply(hl25)
		i26, hl27 := hlist.Unapply(hl26)
		i27, hl28 := hlist.Unapply(hl27)
		i28, hl29 := hlist.Unapply(hl28)
		i29 := hl29.Head()
		return Over21Builder{}.Apply(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29).Build()
	},
	func(v Over21) hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]] {
		i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29 := v.Unapply()
		return hlist.Concat(i0,
			hlist.Concat(i1,
				hlist.Concat(i2,
					hlist.Concat(i3,
						hlist.Concat(i4,
							hlist.Concat(i5,
								hlist.Concat(i6,
									hlist.Concat(i7,
										hlist.Concat(i8,
											hlist.Concat(i9,
												hlist.Concat(i10,
													hlist.Concat(i11,
														hlist.Concat(i12,
															hlist.Concat(i13,
																hlist.Concat(i14,
																	hlist.Concat(i15,
																		hlist.Concat(i16,
																			hlist.Concat(i17,
																				hlist.Concat(i18,
																					hlist.Concat(i19,
																						hlist.Concat(i20,
																							hlist.Concat(i21,
																								hlist.Concat(i22,
																									hlist.Concat(i23,
																										hlist.Concat(i24,
																											hlist.Concat(i25,
																												hlist.Concat(i26,
																													hlist.Concat(i27,
																														hlist.Concat(i28,
																															hlist.Concat(i29,
																																hlist.Empty(),
																															),
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		)
	},
)

var ReadOver21 = read.Generic(
	as.Generic(
		"testpk1.Over21",
		"Struct",
		func(v Over21) hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]] {
			i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29 := v.Unapply()
			return hlist.Concat(i0,
				hlist.Concat(i1,
					hlist.Concat(i2,
						hlist.Concat(i3,
							hlist.Concat(i4,
								hlist.Concat(i5,
									hlist.Concat(i6,
										hlist.Concat(i7,
											hlist.Concat(i8,
												hlist.Concat(i9,
													hlist.Concat(i10,
														hlist.Concat(i11,
															hlist.Concat(i12,
																hlist.Concat(i13,
																	hlist.Concat(i14,
																		hlist.Concat(i15,
																			hlist.Concat(i16,
																				hlist.Concat(i17,
																					hlist.Concat(i18,
																						hlist.Concat(i19,
																							hlist.Concat(i20,
																								hlist.Concat(i21,
																									hlist.Concat(i22,
																										hlist.Concat(i23,
																											hlist.Concat(i24,
																												hlist.Concat(i25,
																													hlist.Concat(i26,
																														hlist.Concat(i27,
																															hlist.Concat(i28,
																																hlist.Concat(i29,
																																	hlist.Empty(),
																																),
																															),
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			)
		},
		func(hl0 hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Cons[int, hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]) Over21 {
			i0, hl1 := hlist.Unapply(hl0)
			i1, hl2 := hlist.Unapply(hl1)
			i2, hl3 := hlist.Unapply(hl2)
			i3, hl4 := hlist.Unapply(hl3)
			i4, hl5 := hlist.Unapply(hl4)
			i5, hl6 := hlist.Unapply(hl5)
			i6, hl7 := hlist.Unapply(hl6)
			i7, hl8 := hlist.Unapply(hl7)
			i8, hl9 := hlist.Unapply(hl8)
			i9, hl10 := hlist.Unapply(hl9)
			i10, hl11 := hlist.Unapply(hl10)
			i11, hl12 := hlist.Unapply(hl11)
			i12, hl13 := hlist.Unapply(hl12)
			i13, hl14 := hlist.Unapply(hl13)
			i14, hl15 := hlist.Unapply(hl14)
			i15, hl16 := hlist.Unapply(hl15)
			i16, hl17 := hlist.Unapply(hl16)
			i17, hl18 := hlist.Unapply(hl17)
			i18, hl19 := hlist.Unapply(hl18)
			i19, hl20 := hlist.Unapply(hl19)
			i20, hl21 := hlist.Unapply(hl20)
			i21, hl22 := hlist.Unapply(hl21)
			i22, hl23 := hlist.Unapply(hl22)
			i23, hl24 := hlist.Unapply(hl23)
			i24, hl25 := hlist.Unapply(hl24)
			i25, hl26 := hlist.Unapply(hl25)
			i26, hl27 := hlist.Unapply(hl26)
			i27, hl28 := hlist.Unapply(hl27)
			i28, hl29 := hlist.Unapply(hl28)
			i29 := hl29.Head()
			return Over21Builder{}.Apply(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29).Build()
		},
	),
	read.TupleHCons(
		read.Int[int](),
		read.TupleHCons(
			read.Int[int](),
			read.TupleHCons(
				read.Int[int](),
				read.TupleHCons(
					read.Int[int](),
					read.TupleHCons(
						read.Int[int](),
						read.TupleHCons(
							read.Int[int](),
							read.TupleHCons(
								read.Int[int](),
								read.TupleHCons(
									read.Int[int](),
									read.TupleHCons(
										read.Int[int](),
										read.TupleHCons(
											read.Int[int](),
											read.TupleHCons(
												read.Int[int](),
												read.TupleHCons(
													read.Int[int](),
													read.TupleHCons(
														read.Int[int](),
														read.TupleHCons(
															read.Int[int](),
															read.TupleHCons(
																read.Int[int](),
																read.TupleHCons(
																	read.Int[int](),
																	read.TupleHCons(
																		read.Int[int](),
																		read.TupleHCons(
																			read.Int[int](),
																			read.TupleHCons(
																				read.Int[int](),
																				read.TupleHCons(
																					read.Int[int](),
																					read.TupleHCons(
																						read.Int[int](),
																						read.TupleHCons(
																							read.Int[int](),
																							read.TupleHCons(
																								read.Int[int](),
																								read.TupleHCons(
																									read.Int[int](),
																									read.TupleHCons(
																										read.Int[int](),
																										read.TupleHCons(
																											read.Int[int](),
																											read.TupleHCons(
																												read.Int[int](),
																												read.TupleHCons(
																													read.Int[int](),
																													read.TupleHCons(
																														read.Int[int](),
																														read.TupleHCons(
																															read.Int[int](),
																															read.TupleHNill,
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	),
)

var EncoderOver21 = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
		js.EncoderHConsLabelled(
			js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
			js.EncoderHConsLabelled(
				js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
				js.EncoderHConsLabelled(
					js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
					js.EncoderHConsLabelled(
						js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
						js.EncoderHConsLabelled(
							js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
							js.EncoderHConsLabelled(
								js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
								js.EncoderHConsLabelled(
									js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
									js.EncoderHConsLabelled(
										js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
										js.EncoderHConsLabelled(
											js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
											js.EncoderHConsLabelled(
												js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
												js.EncoderHConsLabelled(
													js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
													js.EncoderHConsLabelled(
														js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
														js.EncoderHConsLabelled(
															js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
															js.EncoderHConsLabelled(
																js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																js.EncoderHConsLabelled(
																	js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																	js.EncoderHConsLabelled(
																		js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																		js.EncoderHConsLabelled(
																			js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																			js.EncoderHConsLabelled(
																				js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																				js.EncoderHConsLabelled(
																					js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																					js.EncoderHConsLabelled(
																						js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																						js.EncoderHConsLabelled(
																							js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																							js.EncoderHConsLabelled(
																								js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																								js.EncoderHConsLabelled(
																									js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																									js.EncoderHConsLabelled(
																										js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																										js.EncoderHConsLabelled(
																											js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																											js.EncoderHConsLabelled(
																												js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																												js.EncoderHConsLabelled(
																													js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																													js.EncoderHConsLabelled(
																														js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																														js.EncoderHConsLabelled(
																															js.EncoderNamed[fp.RuntimeNamed[int]](js.EncoderNumber[int]()),
																															js.EncoderHNil,
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	),
	func(v Over21) hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]] {
		i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29 := v.Unapply()
		return hlist.Concat(fp.RuntimeNamed[int]{I1: "i1", I2: i0},
			hlist.Concat(fp.RuntimeNamed[int]{I1: "i2", I2: i1},
				hlist.Concat(fp.RuntimeNamed[int]{I1: "i3", I2: i2},
					hlist.Concat(fp.RuntimeNamed[int]{I1: "i4", I2: i3},
						hlist.Concat(fp.RuntimeNamed[int]{I1: "i5", I2: i4},
							hlist.Concat(fp.RuntimeNamed[int]{I1: "i6", I2: i5},
								hlist.Concat(fp.RuntimeNamed[int]{I1: "i7", I2: i6},
									hlist.Concat(fp.RuntimeNamed[int]{I1: "i8", I2: i7},
										hlist.Concat(fp.RuntimeNamed[int]{I1: "i9", I2: i8},
											hlist.Concat(fp.RuntimeNamed[int]{I1: "i10", I2: i9},
												hlist.Concat(fp.RuntimeNamed[int]{I1: "i11", I2: i10},
													hlist.Concat(fp.RuntimeNamed[int]{I1: "i12", I2: i11},
														hlist.Concat(fp.RuntimeNamed[int]{I1: "i13", I2: i12},
															hlist.Concat(fp.RuntimeNamed[int]{I1: "i14", I2: i13},
																hlist.Concat(fp.RuntimeNamed[int]{I1: "i15", I2: i14},
																	hlist.Concat(fp.RuntimeNamed[int]{I1: "i16", I2: i15},
																		hlist.Concat(fp.RuntimeNamed[int]{I1: "i17", I2: i16},
																			hlist.Concat(fp.RuntimeNamed[int]{I1: "i18", I2: i17},
																				hlist.Concat(fp.RuntimeNamed[int]{I1: "i19", I2: i18},
																					hlist.Concat(fp.RuntimeNamed[int]{I1: "i20", I2: i19},
																						hlist.Concat(fp.RuntimeNamed[int]{I1: "i21", I2: i20},
																							hlist.Concat(fp.RuntimeNamed[int]{I1: "i22", I2: i21},
																								hlist.Concat(fp.RuntimeNamed[int]{I1: "i23", I2: i22},
																									hlist.Concat(fp.RuntimeNamed[int]{I1: "i24", I2: i23},
																										hlist.Concat(fp.RuntimeNamed[int]{I1: "i25", I2: i24},
																											hlist.Concat(fp.RuntimeNamed[int]{I1: "i26", I2: i25},
																												hlist.Concat(fp.RuntimeNamed[int]{I1: "i27", I2: i26},
																													hlist.Concat(fp.RuntimeNamed[int]{I1: "i28", I2: i27},
																														hlist.Concat(fp.RuntimeNamed[int]{I1: "i29", I2: i28},
																															hlist.Concat(fp.RuntimeNamed[int]{I1: "i30", I2: i29},
																																hlist.Empty(),
																															),
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		)
	},
)

var DecoderOver21 = js.DecoderMap(
	js.DecoderHConsLabelled(
		js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
		js.DecoderHConsLabelled(
			js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
			js.DecoderHConsLabelled(
				js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
				js.DecoderHConsLabelled(
					js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
					js.DecoderHConsLabelled(
						js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
						js.DecoderHConsLabelled(
							js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
							js.DecoderHConsLabelled(
								js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
								js.DecoderHConsLabelled(
									js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
									js.DecoderHConsLabelled(
										js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
										js.DecoderHConsLabelled(
											js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
											js.DecoderHConsLabelled(
												js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
												js.DecoderHConsLabelled(
													js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
													js.DecoderHConsLabelled(
														js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
														js.DecoderHConsLabelled(
															js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
															js.DecoderHConsLabelled(
																js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																js.DecoderHConsLabelled(
																	js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																	js.DecoderHConsLabelled(
																		js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																		js.DecoderHConsLabelled(
																			js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																			js.DecoderHConsLabelled(
																				js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																				js.DecoderHConsLabelled(
																					js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																					js.DecoderHConsLabelled(
																						js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																						js.DecoderHConsLabelled(
																							js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																							js.DecoderHConsLabelled(
																								js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																								js.DecoderHConsLabelled(
																									js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																									js.DecoderHConsLabelled(
																										js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																										js.DecoderHConsLabelled(
																											js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																											js.DecoderHConsLabelled(
																												js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																												js.DecoderHConsLabelled(
																													js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																													js.DecoderHConsLabelled(
																														js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																														js.DecoderHConsLabelled(
																															js.DecoderNamed[fp.RuntimeNamed[int]](js.DecoderNumber[int]()),
																															js.DecoderHNil,
																														),
																													),
																												),
																											),
																										),
																									),
																								),
																							),
																						),
																					),
																				),
																			),
																		),
																	),
																),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	),
	func(hl0 hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Cons[fp.RuntimeNamed[int], hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]) Over21 {
		i0, hl1 := hlist.Unapply(hl0)
		i1, hl2 := hlist.Unapply(hl1)
		i2, hl3 := hlist.Unapply(hl2)
		i3, hl4 := hlist.Unapply(hl3)
		i4, hl5 := hlist.Unapply(hl4)
		i5, hl6 := hlist.Unapply(hl5)
		i6, hl7 := hlist.Unapply(hl6)
		i7, hl8 := hlist.Unapply(hl7)
		i8, hl9 := hlist.Unapply(hl8)
		i9, hl10 := hlist.Unapply(hl9)
		i10, hl11 := hlist.Unapply(hl10)
		i11, hl12 := hlist.Unapply(hl11)
		i12, hl13 := hlist.Unapply(hl12)
		i13, hl14 := hlist.Unapply(hl13)
		i14, hl15 := hlist.Unapply(hl14)
		i15, hl16 := hlist.Unapply(hl15)
		i16, hl17 := hlist.Unapply(hl16)
		i17, hl18 := hlist.Unapply(hl17)
		i18, hl19 := hlist.Unapply(hl18)
		i19, hl20 := hlist.Unapply(hl19)
		i20, hl21 := hlist.Unapply(hl20)
		i21, hl22 := hlist.Unapply(hl21)
		i22, hl23 := hlist.Unapply(hl22)
		i23, hl24 := hlist.Unapply(hl23)
		i24, hl25 := hlist.Unapply(hl24)
		i25, hl26 := hlist.Unapply(hl25)
		i26, hl27 := hlist.Unapply(hl26)
		i27, hl28 := hlist.Unapply(hl27)
		i28, hl29 := hlist.Unapply(hl28)
		i29 := hl29.Head()
		return Over21Builder{}.Apply(i0.Value(), i1.Value(), i2.Value(), i3.Value(), i4.Value(), i5.Value(), i6.Value(), i7.Value(), i8.Value(), i9.Value(), i10.Value(), i11.Value(), i12.Value(), i13.Value(), i14.Value(), i15.Value(), i16.Value(), i17.Value(), i18.Value(), i19.Value(), i20.Value(), i21.Value(), i22.Value(), i23.Value(), i24.Value(), i25.Value(), i26.Value(), i27.Value(), i28.Value(), i29.Value()).Build()
	},
)
