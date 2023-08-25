// Code generated by gombok, DO NOT EDIT.
package showtest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/show"
	"github.com/csgura/fp/test/internal/recursive"
)

func ShowPerson() fp.Show[Person] {
	return show.Generic(
		as.Generic(
			"showtest.Person",
			"Struct",
			fp.Compose(
				func(v Person) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]] {
					i0, i1 := v.Name, v.Age
					return as.Labelled2(as.Named("Name", i0), as.Named("Age", i1))
				},
				as.HList2Labelled,
			),

			fp.Compose(
				product.LabelledFromHList2,
				func(t fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]]) Person {
					return Person{Name: t.I1.Value(), Age: t.I2.Value()}
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string]](show.String),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[int]](show.Int[int]()),
				show.HNil,
			),
		),
	)
}

func ShowCollection() fp.Show[Collection] {
	return show.Generic(
		as.Generic(
			"showtest.Collection",
			"Struct",
			fp.Compose(
				func(v Collection) fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]] {
					i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10 := v.Index, v.List, v.Description, v.Set, v.Option, v.NoDerive, v.Stringer, v.BoolPtr, v.NoMap, v.Alias, v.StringSeq
					return as.Labelled11(as.Named("Index", i0), as.Named("List", i1), as.Named("Description", i2), as.Named("Set", i3), as.Named("Option", i4), as.Named("NoDerive", i5), as.Named("Stringer", i6), as.Named("BoolPtr", i7), as.Named("NoMap", i8), as.Named("Alias", i9), as.Named("StringSeq", i10))
				},
				as.HList11Labelled,
			),

			fp.Compose(
				product.LabelledFromHList11,
				func(t fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]]) Collection {
					return Collection{Index: t.I1.Value(), List: t.I2.Value(), Description: t.I3.Value(), Set: t.I4.Value(), Option: t.I5.Value(), NoDerive: t.I6.Value(), Stringer: t.I7.Value(), BoolPtr: t.I8.Value(), NoMap: t.I9.Value(), Alias: t.I10.Value(), StringSeq: t.I11.Value()}
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[map[string]Person]](show.GoMap(show.String, ShowPerson())),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[[]Person]](show.Slice(ShowPerson())),
				show.HConsLabelled(
					show.Named[fp.RuntimeNamed[*string]](show.Ptr(lazy.Call(func() fp.Show[string] {
						return show.String
					}))),
					show.HConsLabelled(
						show.Named[fp.RuntimeNamed[fp.Set[int]]](show.Set(show.Int[int]())),
						show.HConsLabelled(
							show.Named[fp.RuntimeNamed[fp.Option[Person]]](show.Option(ShowPerson())),
							show.HConsLabelled(
								show.Named[fp.RuntimeNamed[NoDerive]](ShowNoDerive()),
								show.HConsLabelled(
									show.Named[fp.RuntimeNamed[HasStringMethod]](show.Given[HasStringMethod]()),
									show.HConsLabelled(
										show.Named[fp.RuntimeNamed[*bool]](show.Ptr(lazy.Call(func() fp.Show[bool] {
											return show.Bool
										}))),
										show.HConsLabelled(
											show.Named[fp.RuntimeNamed[map[string]NoDerive]](show.GoMap(show.String, ShowNoDerive())),
											show.HConsLabelled(
												show.Named[fp.RuntimeNamed[recursive.StringAlias]](ShowRecursiveStringAlias()),
												show.HConsLabelled(
													show.Named[fp.RuntimeNamed[fp.Seq[string]]](show.Seq(show.String)),
													show.HNil,
												),
											),
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
}

func ShowDupGenerate() fp.Show[DupGenerate] {
	return show.Generic(
		as.Generic(
			"showtest.DupGenerate",
			"Struct",
			fp.Compose(
				func(v DupGenerate) fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]] {
					i0, i1 := v.NoDerive, v.World
					return as.Labelled2(as.Named("NoDerive", i0), as.Named("World", i1))
				},
				as.HList2Labelled,
			),

			fp.Compose(
				product.LabelledFromHList2,
				func(t fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]]) DupGenerate {
					return DupGenerate{NoDerive: t.I1.Value(), World: t.I2.Value()}
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[NoDerive]](ShowNoDerive()),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[string]](show.String),
				show.HNil,
			),
		),
	)
}

func ShowHasTuple() fp.Show[HasTuple] {
	return show.Generic(
		as.Generic(
			"showtest.HasTuple",
			"Struct",
			fp.Compose(
				func(v HasTuple) fp.Labelled2[fp.RuntimeNamed[fp.Tuple2[string, int]], fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]] {
					i0, i1 := v.Entry, v.HList
					return as.Labelled2(as.Named("Entry", i0), as.Named("HList", i1))
				},
				as.HList2Labelled,
			),

			fp.Compose(
				product.LabelledFromHList2,
				func(t fp.Labelled2[fp.RuntimeNamed[fp.Tuple2[string, int]], fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]]) HasTuple {
					return HasTuple{Entry: t.I1.Value(), HList: t.I2.Value()}
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[fp.Tuple2[string, int]]](show.Generic(
				as.Generic(
					"fp.Tuple2",
					"Tuple",
					as.HList2,
					product.TupleFromHList2[string, int],
				),
				show.TupleHCons(
					show.String,
					show.TupleHCons(
						show.Int[int](),
						show.HNil,
					),
				),
			)),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]](show.HCons(show.String, show.HCons(show.Int[int](), show.HNil))),
				show.HNil,
			),
		),
	)
}

func ShowEmbeddedStruct() fp.Show[EmbeddedStruct] {
	return show.Generic(
		as.Generic(
			"showtest.EmbeddedStruct",
			"Struct",
			fp.Compose(
				func(v EmbeddedStruct) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
					Level int
					Stage string
				}]] {
					i0, i1 := v.Unapply()
					return as.Labelled2(as.Named("hello", i0), as.Named("world", i1))
				},
				as.HList2Labelled,
			),

			fp.Compose(
				product.LabelledFromHList2,
				func(t fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
					Level int
					Stage string
				}]]) EmbeddedStruct {
					return EmbeddedStructBuilder{}.Apply(t.I1.Value(), t.I2.Value()).Build()
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string]](show.String),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[struct {
					Level int
					Stage string
				}]](show.Generic(
					as.Generic(
						"struct",
						"Struct",
						fp.Compose(
							func(v struct {
								Level int
								Stage string
							}) fp.Labelled2[fp.RuntimeNamed[int], fp.RuntimeNamed[string]] {
								i0, i1 := v.Level, v.Stage
								return as.Labelled2(as.Named("Level", i0), as.Named("Stage", i1))
							},
							as.HList2Labelled,
						),

						fp.Compose(
							product.LabelledFromHList2,
							func(t fp.Labelled2[fp.RuntimeNamed[int], fp.RuntimeNamed[string]]) struct {
								Level int
								Stage string
							} {
								return struct {
									Level int
									Stage string
								}{Level: t.I1.Value(), Stage: t.I2.Value()}
							},
						),
					),
					show.HConsLabelled(
						show.Named[fp.RuntimeNamed[int]](show.Int[int]()),
						show.HConsLabelled(
							show.Named[fp.RuntimeNamed[string]](show.String),
							show.HNil,
						),
					),
				)),
				show.HNil,
			),
		),
	)
}

func ShowEmbeddedTypeParamStruct[T any](showT fp.Show[T]) fp.Show[EmbeddedTypeParamStruct[T]] {
	return show.Generic(
		as.Generic(
			"showtest.EmbeddedTypeParamStruct",
			"Struct",
			fp.Compose(
				func(v EmbeddedTypeParamStruct[T]) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
					Level T
					Stage string
				}]] {
					i0, i1 := v.Unapply()
					return as.Labelled2(as.Named("hello", i0), as.Named("world", i1))
				},
				as.HList2Labelled,
			),

			fp.Compose(
				product.LabelledFromHList2,
				func(t fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
					Level T
					Stage string
				}]]) EmbeddedTypeParamStruct[T] {
					return EmbeddedTypeParamStructBuilder[T]{}.Apply(t.I1.Value(), t.I2.Value()).Build()
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string]](show.String),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[struct {
					Level T
					Stage string
				}]](show.Generic(
					as.Generic(
						"struct",
						"Struct",
						fp.Compose(
							func(v struct {
								Level T
								Stage string
							}) fp.Labelled2[fp.RuntimeNamed[T], fp.RuntimeNamed[string]] {
								i0, i1 := v.Level, v.Stage
								return as.Labelled2(as.Named("Level", i0), as.Named("Stage", i1))
							},
							as.HList2Labelled,
						),

						fp.Compose(
							product.LabelledFromHList2,
							func(t fp.Labelled2[fp.RuntimeNamed[T], fp.RuntimeNamed[string]]) struct {
								Level T
								Stage string
							} {
								return struct {
									Level T
									Stage string
								}{Level: t.I1.Value(), Stage: t.I2.Value()}
							},
						),
					),
					show.HConsLabelled(
						show.Named[fp.RuntimeNamed[T]](showT),
						show.HConsLabelled(
							show.Named[fp.RuntimeNamed[string]](show.String),
							show.HNil,
						),
					),
				)),
				show.HNil,
			),
		),
	)
}

func ShowRecursiveStringAlias() fp.Show[recursive.StringAlias] {
	return show.Generic(
		as.Generic(
			"recursive.StringAlias",
			"NewType",
			func(v recursive.StringAlias) string {
				return string(v)
			},
			func(v string) recursive.StringAlias {
				return recursive.StringAlias(v)
			},
		),
		show.String,
	)
}

func ShowNoDerive() fp.Show[NoDerive] {
	return show.Generic(
		as.Generic(
			"showtest.NoDerive",
			"Struct",
			fp.Compose(
				func(v NoDerive) fp.Labelled1[fp.RuntimeNamed[string]] {
					i0 := v.Hello
					return as.Labelled1(as.Named("Hello", i0))
				},
				as.HList1Labelled,
			),

			fp.Compose(
				product.LabelledFromHList1,
				func(t fp.Labelled1[fp.RuntimeNamed[string]]) NoDerive {
					return NoDerive{Hello: t.I1.Value()}
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string]](show.String),
			show.HNil,
		),
	)
}
