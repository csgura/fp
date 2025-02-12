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
		fp.Generic[Person, fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]]]{
			Type: "showtest.Person",
			Kind: "Struct",
			To: func(v Person) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]] {
				i0, i1 := v.Name, v.Age
				return as.Labelled2(as.NamedWithTag("Name", i0, ``), as.NamedWithTag("Age", i1, ``))
			},
			From: func(t fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[int]]) Person {
				return Person{Name: t.I1.Value(), Age: t.I2.Value()}
			},
		},
		show.Labelled2(show.Named[fp.RuntimeNamed[string]](show.String), show.Named[fp.RuntimeNamed[int]](show.Int[int]())),
	)
}

func ShowCollection() fp.Show[Collection] {
	return show.Generic(
		fp.Generic[Collection, fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]]]{
			Type: "showtest.Collection",
			Kind: "Struct",
			To: func(v Collection) fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]] {
				i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10 := v.Index, v.List, v.Description, v.Set, v.Option, v.NoDerive, v.Stringer, v.BoolPtr, v.NoMap, v.Alias, v.StringSeq
				return as.Labelled11(as.NamedWithTag("Index", i0, ``), as.NamedWithTag("List", i1, ``), as.NamedWithTag("Description", i2, ``), as.NamedWithTag("Set", i3, ``), as.NamedWithTag("Option", i4, ``), as.NamedWithTag("NoDerive", i5, ``), as.NamedWithTag("Stringer", i6, ``), as.NamedWithTag("BoolPtr", i7, ``), as.NamedWithTag("NoMap", i8, ``), as.NamedWithTag("Alias", i9, ``), as.NamedWithTag("StringSeq", i10, ``))
			},
			From: func(t fp.Labelled11[fp.RuntimeNamed[map[string]Person], fp.RuntimeNamed[[]Person], fp.RuntimeNamed[*string], fp.RuntimeNamed[fp.Set[int]], fp.RuntimeNamed[fp.Option[Person]], fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[HasStringMethod], fp.RuntimeNamed[*bool], fp.RuntimeNamed[map[string]NoDerive], fp.RuntimeNamed[recursive.StringAlias], fp.RuntimeNamed[fp.Seq[string]]]) Collection {
				return Collection{Index: t.I1.Value(), List: t.I2.Value(), Description: t.I3.Value(), Set: t.I4.Value(), Option: t.I5.Value(), NoDerive: t.I6.Value(), Stringer: t.I7.Value(), BoolPtr: t.I8.Value(), NoMap: t.I9.Value(), Alias: t.I10.Value(), StringSeq: t.I11.Value()}
			},
		},
		show.Labelled11(show.Named[fp.RuntimeNamed[map[string]Person]](show.GoMap(show.String, ShowPerson())), show.Named[fp.RuntimeNamed[[]Person]](show.Slice(ShowPerson())), show.Named[fp.RuntimeNamed[*string]](show.Ptr(lazy.Call(func() fp.Show[string] {
			return show.String
		}))), show.Named[fp.RuntimeNamed[fp.Set[int]]](show.Set(show.Int[int]())), show.Named[fp.RuntimeNamed[fp.Option[Person]]](show.Option(ShowPerson())), show.Named[fp.RuntimeNamed[NoDerive]](ShowNoDerive()), show.Named[fp.RuntimeNamed[HasStringMethod]](show.Given[HasStringMethod]()), show.Named[fp.RuntimeNamed[*bool]](show.Ptr(lazy.Call(func() fp.Show[bool] {
			return show.Bool
		}))), show.Named[fp.RuntimeNamed[map[string]NoDerive]](show.GoMap(show.String, ShowNoDerive())), show.Named[fp.RuntimeNamed[recursive.StringAlias]](ShowRecursiveStringAlias()), show.Named[fp.RuntimeNamed[fp.Seq[string]]](show.Seq(show.String))),
	)
}

func ShowDupGenerate() fp.Show[DupGenerate] {
	return show.Generic(
		fp.Generic[DupGenerate, fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]]]{
			Type: "showtest.DupGenerate",
			Kind: "Struct",
			To: func(v DupGenerate) fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]] {
				i0, i1 := v.NoDerive, v.World
				return as.Labelled2(as.NamedWithTag("NoDerive", i0, ``), as.NamedWithTag("World", i1, ``))
			},
			From: func(t fp.Labelled2[fp.RuntimeNamed[NoDerive], fp.RuntimeNamed[string]]) DupGenerate {
				return DupGenerate{NoDerive: t.I1.Value(), World: t.I2.Value()}
			},
		},
		show.Labelled2(show.Named[fp.RuntimeNamed[NoDerive]](ShowNoDerive()), show.Named[fp.RuntimeNamed[string]](show.String)),
	)
}

func ShowHasTuple() fp.Show[HasTuple] {
	return show.Generic(
		fp.Generic[HasTuple, fp.Labelled2[fp.RuntimeNamed[fp.Entry[int]], fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]]]{
			Type: "showtest.HasTuple",
			Kind: "Struct",
			To: func(v HasTuple) fp.Labelled2[fp.RuntimeNamed[fp.Entry[int]], fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]] {
				i0, i1 := v.Entry, v.HList
				return as.Labelled2(as.NamedWithTag("Entry", i0, ``), as.NamedWithTag("HList", i1, ``))
			},
			From: func(t fp.Labelled2[fp.RuntimeNamed[fp.Entry[int]], fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]]) HasTuple {
				return HasTuple{Entry: t.I1.Value(), HList: t.I2.Value()}
			},
		},
		show.Labelled2(show.Named[fp.RuntimeNamed[fp.Entry[int]]](show.Generic(
			fp.Generic[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]{
				Type: "fp.Tuple2",
				Kind: "Tuple",
				To:   as.HList2[string, int],
				From: product.TupleFromHList2[string, int],
			},
			show.TupleHCons(
				show.String,
				show.TupleHCons(
					show.Int[int](),
					show.HNil,
				),
			),
		)), show.Named[fp.RuntimeNamed[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]](show.TupleHCons(show.String, show.TupleHCons(show.Int[int](), show.HNil)))),
	)
}

func ShowEmbeddedStruct() fp.Show[EmbeddedStruct] {
	return show.Generic(
		fp.Generic[EmbeddedStruct, fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
			Level int
			Stage string
		}]]]{
			Type: "showtest.EmbeddedStruct",
			Kind: "Struct",
			To: func(v EmbeddedStruct) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
				Level int
				Stage string
			}]] { i0, i1 := v.Unapply(); return as.Labelled2(as.NamedWithTag("hello", i0, ``), as.NamedWithTag("world", i1, ``)) },
			From: func(t fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
				Level int
				Stage string
			}]]) EmbeddedStruct {
				return EmbeddedStructBuilder{}.Apply(t.I1.Value(), t.I2.Value()).Build()
			},
		},
		show.Labelled2(show.Named[fp.RuntimeNamed[string]](show.String), show.Named[fp.RuntimeNamed[struct {
			Level int
			Stage string
		}]](show.Generic(
			fp.Generic[struct {
				Level int
				Stage string
			}, fp.Labelled2[fp.RuntimeNamed[int], fp.RuntimeNamed[string]]]{
				Type: "struct",
				Kind: "Struct",
				To: func(v struct {
					Level int
					Stage string
				}) fp.Labelled2[fp.RuntimeNamed[int], fp.RuntimeNamed[string]] {
					i0, i1 := v.Level, v.Stage
					return as.Labelled2(as.NamedWithTag("Level", i0, ``), as.NamedWithTag("Stage", i1, ``))
				},
				From: func(t fp.Labelled2[fp.RuntimeNamed[int], fp.RuntimeNamed[string]]) struct {
					Level int
					Stage string
				} {
					return struct {
						Level int
						Stage string
					}{Level: t.I1.Value(), Stage: t.I2.Value()}
				},
			},
			show.Labelled2(show.Named[fp.RuntimeNamed[int]](show.Int[int]()), show.Named[fp.RuntimeNamed[string]](show.String)),
		))),
	)
}

func ShowEmbeddedTypeParamStruct[T any](showT fp.Show[T]) fp.Show[EmbeddedTypeParamStruct[T]] {
	return show.Generic(
		fp.Generic[EmbeddedTypeParamStruct[T], fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
			Level T
			Stage string
		}]]]{
			Type: "showtest.EmbeddedTypeParamStruct",
			Kind: "Struct",
			To: func(v EmbeddedTypeParamStruct[T]) fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
				Level T
				Stage string
			}]] { i0, i1 := v.Unapply(); return as.Labelled2(as.NamedWithTag("hello", i0, ``), as.NamedWithTag("world", i1, ``)) },
			From: func(t fp.Labelled2[fp.RuntimeNamed[string], fp.RuntimeNamed[struct {
				Level T
				Stage string
			}]]) EmbeddedTypeParamStruct[T] {
				return EmbeddedTypeParamStructBuilder[T]{}.Apply(t.I1.Value(), t.I2.Value()).Build()
			},
		},
		show.Labelled2(show.Named[fp.RuntimeNamed[string]](show.String), show.Named[fp.RuntimeNamed[struct {
			Level T
			Stage string
		}]](show.Generic(
			fp.Generic[struct {
				Level T
				Stage string
			}, fp.Labelled2[fp.RuntimeNamed[T], fp.RuntimeNamed[string]]]{
				Type: "struct",
				Kind: "Struct",
				To: func(v struct {
					Level T
					Stage string
				}) fp.Labelled2[fp.RuntimeNamed[T], fp.RuntimeNamed[string]] {
					i0, i1 := v.Level, v.Stage
					return as.Labelled2(as.NamedWithTag("Level", i0, ``), as.NamedWithTag("Stage", i1, ``))
				},
				From: func(t fp.Labelled2[fp.RuntimeNamed[T], fp.RuntimeNamed[string]]) struct {
					Level T
					Stage string
				} {
					return struct {
						Level T
						Stage string
					}{Level: t.I1.Value(), Stage: t.I2.Value()}
				},
			},
			show.Labelled2(show.Named[fp.RuntimeNamed[T]](showT), show.Named[fp.RuntimeNamed[string]](show.String)),
		))),
	)
}

func ShowNoDerive() fp.Show[NoDerive] {
	return show.Generic(
		fp.Generic[NoDerive, hlist.Cons[fp.RuntimeNamed[string], hlist.Nil]]{
			Type: "showtest.NoDerive",
			Kind: "Struct",
			To: fp.Compose(
				func(v NoDerive) fp.Labelled1[fp.RuntimeNamed[string]] {
					i0 := v.Hello
					return as.Labelled1(as.NamedWithTag("Hello", i0, ``))
				},
				as.HList1Labelled,
			),
			From: fp.Compose(
				product.LabelledFromHList1,
				func(t fp.Labelled1[fp.RuntimeNamed[string]]) NoDerive {
					return NoDerive{Hello: t.I1.Value()}
				},
			),
		},
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string]](show.String),
			show.HNil,
		),
	)
}

func ShowRecursiveStringAlias() fp.Show[recursive.StringAlias] {
	return show.Generic(
		fp.Generic[recursive.StringAlias, string]{
			Type: "recursive.StringAlias",
			Kind: "NewType",
			To: func(v recursive.StringAlias) string {
				return string(v)
			},
			From: func(v string) recursive.StringAlias {
				return recursive.StringAlias(v)
			},
		},
		show.String,
	)
}
