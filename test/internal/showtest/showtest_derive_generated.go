// Code generated by gombok, DO NOT EDIT.
package showtest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/show"
	"github.com/csgura/fp/test/internal/recursive"
)

var ShowPerson = show.Generic(
	as.Generic(
		"showtest.Person",
		"Struct",
		fp.Compose(
			func(v Person) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]] {
				i0, i1 := v.Name, v.Age
				return as.Labelled2(fp.RuntimeNamed[string]{I1: "Name", I2: i0}, fp.RuntimeNamed[int]{I1: "Age", I2: i1})
			},
			as.HList2Labelled[fp.RuntimeNamed[string], fp.RuntimeNamed[int]],
		),

		fp.Compose(
			product.LabelledFromHList2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]],
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

var ShowCollection = show.Generic(
	as.Generic(
		"showtest.Collection",
		"Struct",
		fp.Compose(
			func(v Collection) fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]] {
				i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10 := v.Index, v.List, v.Description, v.Set, v.Option, v.NoDerive, v.Stringer, v.BoolPtr, v.NoMap, v.Alias, v.StringSeq
				return as.Labelled11(fp.RuntimeNamed[map[string]Person]{I1: "Index", I2: i0}, fp.RuntimeNamed[[]Person]{I1: "List", I2: i1}, fp.RuntimeNamed[*string]{I1: "Description", I2: i2}, fp.RuntimeNamed[fp.Set[int]]{I1: "Set", I2: i3}, fp.RuntimeNamed[fp.Option[Person]]{I1: "Option", I2: i4}, fp.RuntimeNamed[NoDerive]{I1: "NoDerive", I2: i5}, fp.RuntimeNamed[HasStringMethod]{I1: "Stringer", I2: i6}, fp.RuntimeNamed[*bool]{I1: "BoolPtr", I2: i7}, fp.RuntimeNamed[map[string]NoDerive]{I1: "NoMap", I2: i8}, fp.RuntimeNamed[recursive.StringAlias]{I1: "Alias", I2: i9}, fp.RuntimeNamed[fp.Seq[string]]{I1: "StringSeq", I2: i10})
			},
			as.HList11Labelled[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]],
		),

		fp.Compose(
			product.LabelledFromHList11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]],
			func(t fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]]) Collection {
				return Collection{Index: t.I1.Value(), List: t.I2.Value(), Description: t.I3.Value(), Set: t.I4.Value(), Option: t.I5.Value(), NoDerive: t.I6.Value(), Stringer: t.I7.Value(), BoolPtr: t.I8.Value(), NoMap: t.I9.Value(), Alias: t.I10.Value(), StringSeq: t.I11.Value()}
			},
		),
	),
	show.HConsLabelled(
		show.Named[fp.RuntimeNamed[map[string]Person]](show.GoMap(show.String, ShowPerson)),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[[]Person]](show.Slice(ShowPerson)),
			show.HConsLabelled(
				show.Named[fp.RuntimeNamed[*string]](show.Ptr(lazy.Call(func() fp.Show[string] {
					return show.String
				}))),
				show.HConsLabelled(
					show.Named[fp.RuntimeNamed[fp.Set[int]]](show.Set(show.Int[int]())),
					show.HConsLabelled(
						show.Named[fp.RuntimeNamed[fp.Option[Person]]](show.Option(ShowPerson)),
						show.HConsLabelled(
							show.Named[fp.RuntimeNamed[NoDerive]](ShowNoDerive),
							show.HConsLabelled(
								show.Named[fp.RuntimeNamed[HasStringMethod]](show.Given[HasStringMethod]()),
								show.HConsLabelled(
									show.Named[fp.RuntimeNamed[*bool]](show.Ptr(lazy.Call(func() fp.Show[bool] {
										return show.Bool
									}))),
									show.HConsLabelled(
										show.Named[fp.RuntimeNamed[map[string]NoDerive]](show.GoMap(show.String, ShowNoDerive)),
										show.HConsLabelled(
											show.Named[fp.RuntimeNamed[recursive.StringAlias]](ShowRecursiveStringAlias),
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

var ShowDupGenerate = show.Generic(
	as.Generic(
		"showtest.DupGenerate",
		"Struct",
		fp.Compose(
			func(v DupGenerate) fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]] {
				i0, i1 := v.NoDerive, v.World
				return as.Labelled2(fp.RuntimeNamed[NoDerive]{I1: "NoDerive", I2: i0}, fp.RuntimeNamed[string]{I1: "World", I2: i1})
			},
			as.HList2Labelled[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]],
		),

		fp.Compose(
			product.LabelledFromHList2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]],
			func(t fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]]) DupGenerate {
				return DupGenerate{NoDerive: t.I1.Value(), World: t.I2.Value()}
			},
		),
	),
	show.HConsLabelled(
		show.Named[fp.RuntimeNamed[NoDerive]](ShowNoDerive),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string]](show.String),
			show.HNil,
		),
	),
)

var ShowRecursiveStringAlias = show.Generic(
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

var ShowNoDerive = show.Generic(
	as.Generic(
		"showtest.NoDerive",
		"Struct",
		fp.Compose(
			func(v NoDerive) fp.Labelled1[fp.RuntimeNamed[string]] {
				i0 := v.Hello
				return as.Labelled1(fp.RuntimeNamed[string]{I1: "Hello", I2: i0})
			},
			as.HList1Labelled[fp.RuntimeNamed[string]],
		),

		fp.Compose(
			product.LabelledFromHList1[fp.RuntimeNamed[string]],
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
