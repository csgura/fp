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
				js.EncoderNamed[NamedPubPub[string]](js.EncoderString),
				js.EncoderHNil,
			),
		),
	),
	fp.Compose(
		World.AsLabelled,
		as.HList3Labelled[NamedMessage[string], NamedTimestamp[time.Time], NamedPubPub[string]],
	),
)

var DecoderWorld = js.DecoderMap(
	js.DecoderHConsLabelled(
		js.DecoderNamed[NamedMessage[string]](js.DecoderString),
		js.DecoderHConsLabelled(
			js.DecoderNamed[NamedTimestamp[time.Time]](js.DecoderTime),
			js.DecoderHConsLabelled(
				js.DecoderNamed[NamedPubPub[string]](js.DecoderString),
				js.DecoderHNil,
			),
		),
	),

	fp.Compose(
		product.LabelledFromHList3[NamedMessage[string], NamedTimestamp[time.Time], NamedPubPub[string]],
		fp.Compose(
			as.Curried2(WorldBuilder.FromLabelled)(WorldBuilder{}),
			WorldBuilder.Build,
		),
	),
)

var ShowWorld = show.Generic(
	as.Generic(
		"testpk1.World",
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
	show.HCons(
		show.String,
		show.HCons(
			show.Time,
			show.HCons(
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
	read.HCons(
		read.String,
		read.HCons(
			read.Time,
			read.HCons(
				read.String,
				read.HNil,
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
	read.HCons(
		read.Int[int](),
		read.HCons(
			read.Int[int](),
			read.HCons(
				read.Int[int](),
				read.HCons(
					read.Int[int](),
					read.HCons(
						read.Int[int](),
						read.HCons(
							read.Int[int](),
							read.HCons(
								read.Int[int](),
								read.HCons(
									read.Int[int](),
									read.HCons(
										read.Int[int](),
										read.HCons(
											read.Int[int](),
											read.HCons(
												read.Int[int](),
												read.HCons(
													read.Int[int](),
													read.HCons(
														read.Int[int](),
														read.HCons(
															read.Int[int](),
															read.HCons(
																read.Int[int](),
																read.HCons(
																	read.Int[int](),
																	read.HCons(
																		read.Int[int](),
																		read.HCons(
																			read.Int[int](),
																			read.HCons(
																				read.Int[int](),
																				read.HCons(
																					read.Int[int](),
																					read.HCons(
																						read.Int[int](),
																						read.HCons(
																							read.Int[int](),
																							read.HCons(
																								read.Int[int](),
																								read.HCons(
																									read.Int[int](),
																									read.HCons(
																										read.Int[int](),
																										read.HCons(
																											read.Int[int](),
																											read.HCons(
																												read.Int[int](),
																												read.HCons(
																													read.Int[int](),
																													read.HCons(
																														read.Int[int](),
																														read.HCons(
																															read.Int[int](),
																															read.HNil,
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
		js.EncoderNamed[NamedI1[int]](js.EncoderNumber[int]()),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedI2[int]](js.EncoderNumber[int]()),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NamedI3[int]](js.EncoderNumber[int]()),
				js.EncoderHConsLabelled(
					js.EncoderNamed[NamedI4[int]](js.EncoderNumber[int]()),
					js.EncoderHConsLabelled(
						js.EncoderNamed[NamedI5[int]](js.EncoderNumber[int]()),
						js.EncoderHConsLabelled(
							js.EncoderNamed[NamedI6[int]](js.EncoderNumber[int]()),
							js.EncoderHConsLabelled(
								js.EncoderNamed[NamedI7[int]](js.EncoderNumber[int]()),
								js.EncoderHConsLabelled(
									js.EncoderNamed[NamedI8[int]](js.EncoderNumber[int]()),
									js.EncoderHConsLabelled(
										js.EncoderNamed[NamedI9[int]](js.EncoderNumber[int]()),
										js.EncoderHConsLabelled(
											js.EncoderNamed[NamedI10[int]](js.EncoderNumber[int]()),
											js.EncoderHConsLabelled(
												js.EncoderNamed[NamedI11[int]](js.EncoderNumber[int]()),
												js.EncoderHConsLabelled(
													js.EncoderNamed[NamedI12[int]](js.EncoderNumber[int]()),
													js.EncoderHConsLabelled(
														js.EncoderNamed[NamedI13[int]](js.EncoderNumber[int]()),
														js.EncoderHConsLabelled(
															js.EncoderNamed[NamedI14[int]](js.EncoderNumber[int]()),
															js.EncoderHConsLabelled(
																js.EncoderNamed[NamedI15[int]](js.EncoderNumber[int]()),
																js.EncoderHConsLabelled(
																	js.EncoderNamed[NamedI16[int]](js.EncoderNumber[int]()),
																	js.EncoderHConsLabelled(
																		js.EncoderNamed[NamedI17[int]](js.EncoderNumber[int]()),
																		js.EncoderHConsLabelled(
																			js.EncoderNamed[NamedI18[int]](js.EncoderNumber[int]()),
																			js.EncoderHConsLabelled(
																				js.EncoderNamed[NamedI19[int]](js.EncoderNumber[int]()),
																				js.EncoderHConsLabelled(
																					js.EncoderNamed[NamedI20[int]](js.EncoderNumber[int]()),
																					js.EncoderHConsLabelled(
																						js.EncoderNamed[NamedI21[int]](js.EncoderNumber[int]()),
																						js.EncoderHConsLabelled(
																							js.EncoderNamed[NamedI22[int]](js.EncoderNumber[int]()),
																							js.EncoderHConsLabelled(
																								js.EncoderNamed[NamedI23[int]](js.EncoderNumber[int]()),
																								js.EncoderHConsLabelled(
																									js.EncoderNamed[NamedI24[int]](js.EncoderNumber[int]()),
																									js.EncoderHConsLabelled(
																										js.EncoderNamed[NamedI25[int]](js.EncoderNumber[int]()),
																										js.EncoderHConsLabelled(
																											js.EncoderNamed[NamedI26[int]](js.EncoderNumber[int]()),
																											js.EncoderHConsLabelled(
																												js.EncoderNamed[NamedI27[int]](js.EncoderNumber[int]()),
																												js.EncoderHConsLabelled(
																													js.EncoderNamed[NamedI28[int]](js.EncoderNumber[int]()),
																													js.EncoderHConsLabelled(
																														js.EncoderNamed[NamedI29[int]](js.EncoderNumber[int]()),
																														js.EncoderHConsLabelled(
																															js.EncoderNamed[NamedI30[int]](js.EncoderNumber[int]()),
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
	func(v Over21) hlist.Cons[NamedI1[int], hlist.Cons[NamedI2[int], hlist.Cons[NamedI3[int], hlist.Cons[NamedI4[int], hlist.Cons[NamedI5[int], hlist.Cons[NamedI6[int], hlist.Cons[NamedI7[int], hlist.Cons[NamedI8[int], hlist.Cons[NamedI9[int], hlist.Cons[NamedI10[int], hlist.Cons[NamedI11[int], hlist.Cons[NamedI12[int], hlist.Cons[NamedI13[int], hlist.Cons[NamedI14[int], hlist.Cons[NamedI15[int], hlist.Cons[NamedI16[int], hlist.Cons[NamedI17[int], hlist.Cons[NamedI18[int], hlist.Cons[NamedI19[int], hlist.Cons[NamedI20[int], hlist.Cons[NamedI21[int], hlist.Cons[NamedI22[int], hlist.Cons[NamedI23[int], hlist.Cons[NamedI24[int], hlist.Cons[NamedI25[int], hlist.Cons[NamedI26[int], hlist.Cons[NamedI27[int], hlist.Cons[NamedI28[int], hlist.Cons[NamedI29[int], hlist.Cons[NamedI30[int], hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]] {
		i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29 := v.Unapply()
		return hlist.Concat(NamedI1[int]{i0},
			hlist.Concat(NamedI2[int]{i1},
				hlist.Concat(NamedI3[int]{i2},
					hlist.Concat(NamedI4[int]{i3},
						hlist.Concat(NamedI5[int]{i4},
							hlist.Concat(NamedI6[int]{i5},
								hlist.Concat(NamedI7[int]{i6},
									hlist.Concat(NamedI8[int]{i7},
										hlist.Concat(NamedI9[int]{i8},
											hlist.Concat(NamedI10[int]{i9},
												hlist.Concat(NamedI11[int]{i10},
													hlist.Concat(NamedI12[int]{i11},
														hlist.Concat(NamedI13[int]{i12},
															hlist.Concat(NamedI14[int]{i13},
																hlist.Concat(NamedI15[int]{i14},
																	hlist.Concat(NamedI16[int]{i15},
																		hlist.Concat(NamedI17[int]{i16},
																			hlist.Concat(NamedI18[int]{i17},
																				hlist.Concat(NamedI19[int]{i18},
																					hlist.Concat(NamedI20[int]{i19},
																						hlist.Concat(NamedI21[int]{i20},
																							hlist.Concat(NamedI22[int]{i21},
																								hlist.Concat(NamedI23[int]{i22},
																									hlist.Concat(NamedI24[int]{i23},
																										hlist.Concat(NamedI25[int]{i24},
																											hlist.Concat(NamedI26[int]{i25},
																												hlist.Concat(NamedI27[int]{i26},
																													hlist.Concat(NamedI28[int]{i27},
																														hlist.Concat(NamedI29[int]{i28},
																															hlist.Concat(NamedI30[int]{i29},
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
		js.DecoderNamed[NamedI1[int]](js.DecoderNumber[int]()),
		js.DecoderHConsLabelled(
			js.DecoderNamed[NamedI2[int]](js.DecoderNumber[int]()),
			js.DecoderHConsLabelled(
				js.DecoderNamed[NamedI3[int]](js.DecoderNumber[int]()),
				js.DecoderHConsLabelled(
					js.DecoderNamed[NamedI4[int]](js.DecoderNumber[int]()),
					js.DecoderHConsLabelled(
						js.DecoderNamed[NamedI5[int]](js.DecoderNumber[int]()),
						js.DecoderHConsLabelled(
							js.DecoderNamed[NamedI6[int]](js.DecoderNumber[int]()),
							js.DecoderHConsLabelled(
								js.DecoderNamed[NamedI7[int]](js.DecoderNumber[int]()),
								js.DecoderHConsLabelled(
									js.DecoderNamed[NamedI8[int]](js.DecoderNumber[int]()),
									js.DecoderHConsLabelled(
										js.DecoderNamed[NamedI9[int]](js.DecoderNumber[int]()),
										js.DecoderHConsLabelled(
											js.DecoderNamed[NamedI10[int]](js.DecoderNumber[int]()),
											js.DecoderHConsLabelled(
												js.DecoderNamed[NamedI11[int]](js.DecoderNumber[int]()),
												js.DecoderHConsLabelled(
													js.DecoderNamed[NamedI12[int]](js.DecoderNumber[int]()),
													js.DecoderHConsLabelled(
														js.DecoderNamed[NamedI13[int]](js.DecoderNumber[int]()),
														js.DecoderHConsLabelled(
															js.DecoderNamed[NamedI14[int]](js.DecoderNumber[int]()),
															js.DecoderHConsLabelled(
																js.DecoderNamed[NamedI15[int]](js.DecoderNumber[int]()),
																js.DecoderHConsLabelled(
																	js.DecoderNamed[NamedI16[int]](js.DecoderNumber[int]()),
																	js.DecoderHConsLabelled(
																		js.DecoderNamed[NamedI17[int]](js.DecoderNumber[int]()),
																		js.DecoderHConsLabelled(
																			js.DecoderNamed[NamedI18[int]](js.DecoderNumber[int]()),
																			js.DecoderHConsLabelled(
																				js.DecoderNamed[NamedI19[int]](js.DecoderNumber[int]()),
																				js.DecoderHConsLabelled(
																					js.DecoderNamed[NamedI20[int]](js.DecoderNumber[int]()),
																					js.DecoderHConsLabelled(
																						js.DecoderNamed[NamedI21[int]](js.DecoderNumber[int]()),
																						js.DecoderHConsLabelled(
																							js.DecoderNamed[NamedI22[int]](js.DecoderNumber[int]()),
																							js.DecoderHConsLabelled(
																								js.DecoderNamed[NamedI23[int]](js.DecoderNumber[int]()),
																								js.DecoderHConsLabelled(
																									js.DecoderNamed[NamedI24[int]](js.DecoderNumber[int]()),
																									js.DecoderHConsLabelled(
																										js.DecoderNamed[NamedI25[int]](js.DecoderNumber[int]()),
																										js.DecoderHConsLabelled(
																											js.DecoderNamed[NamedI26[int]](js.DecoderNumber[int]()),
																											js.DecoderHConsLabelled(
																												js.DecoderNamed[NamedI27[int]](js.DecoderNumber[int]()),
																												js.DecoderHConsLabelled(
																													js.DecoderNamed[NamedI28[int]](js.DecoderNumber[int]()),
																													js.DecoderHConsLabelled(
																														js.DecoderNamed[NamedI29[int]](js.DecoderNumber[int]()),
																														js.DecoderHConsLabelled(
																															js.DecoderNamed[NamedI30[int]](js.DecoderNumber[int]()),
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
	func(hl0 hlist.Cons[NamedI1[int], hlist.Cons[NamedI2[int], hlist.Cons[NamedI3[int], hlist.Cons[NamedI4[int], hlist.Cons[NamedI5[int], hlist.Cons[NamedI6[int], hlist.Cons[NamedI7[int], hlist.Cons[NamedI8[int], hlist.Cons[NamedI9[int], hlist.Cons[NamedI10[int], hlist.Cons[NamedI11[int], hlist.Cons[NamedI12[int], hlist.Cons[NamedI13[int], hlist.Cons[NamedI14[int], hlist.Cons[NamedI15[int], hlist.Cons[NamedI16[int], hlist.Cons[NamedI17[int], hlist.Cons[NamedI18[int], hlist.Cons[NamedI19[int], hlist.Cons[NamedI20[int], hlist.Cons[NamedI21[int], hlist.Cons[NamedI22[int], hlist.Cons[NamedI23[int], hlist.Cons[NamedI24[int], hlist.Cons[NamedI25[int], hlist.Cons[NamedI26[int], hlist.Cons[NamedI27[int], hlist.Cons[NamedI28[int], hlist.Cons[NamedI29[int], hlist.Cons[NamedI30[int], hlist.Nil]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]) Over21 {
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
