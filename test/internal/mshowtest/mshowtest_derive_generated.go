// Code generated by gombok, DO NOT EDIT.
package mshowtest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/minimal"
	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/test/internal/recursive"
)

func ShowPerson() mshow.Show[Person] {
	return mshow.ContraGeneric(
		"mshowtest.Person",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`Name`, ``), as.NameTag(`Age`, ``)}, mshow.String, mshow.Int[int]()),
		func(v Person) minimal.Tuple2[string, int] {
			return minimal.Tuple2[string, int]{
				I1: v.Name,
				I2: v.Age,
			}
		},
	)
}

//go:noinline
func ShowCollection() mshow.Show[Collection] {
	return mshow.ContraGeneric(
		"mshowtest.Collection",
		"Struct",
		mshow.Struct11([]fp.Named{as.NameTag(`Index`, ``), as.NameTag(`List`, ``), as.NameTag(`Description`, ``), as.NameTag(`Set`, ``), as.NameTag(`Option`, ``), as.NameTag(`NoDerive`, ``), as.NameTag(`Stringer`, ``), as.NameTag(`BoolPtr`, ``), as.NameTag(`NoMap`, ``), as.NameTag(`Alias`, ``), as.NameTag(`StringSeq`, ``)}, mshow.GoMap(mshow.String, ShowPerson()), mshow.Slice(ShowPerson()), mshow.Ptr(lazy.Call(func() mshow.Show[string] {
			return mshow.String
		})), mshow.Set(mshow.Int[int]()), mshow.Option(ShowPerson()), ShowNoDerive(), mshow.Given[HasStringMethod](), mshow.Ptr(lazy.Call(func() mshow.Show[bool] {
			return mshow.Bool
		})), mshow.GoMap(mshow.String, ShowNoDerive()), ShowRecursiveStringAlias(), mshow.Seq(mshow.String)),
		func(v Collection) minimal.Tuple11[map[string]Person, []Person, *string, fp.Set[int], fp.Option[Person], NoDerive, HasStringMethod, *bool, map[string]NoDerive, recursive.StringAlias, fp.Seq[string]] {
			return minimal.Tuple11[map[string]Person, []Person, *string, fp.Set[int], fp.Option[Person], NoDerive, HasStringMethod, *bool, map[string]NoDerive, recursive.StringAlias, fp.Seq[string]]{
				I1:  v.Index,
				I2:  v.List,
				I3:  v.Description,
				I4:  v.Set,
				I5:  v.Option,
				I6:  v.NoDerive,
				I7:  v.Stringer,
				I8:  v.BoolPtr,
				I9:  v.NoMap,
				I10: v.Alias,
				I11: v.StringSeq,
			}
		},
	)
}

func ShowDupGenerate() mshow.Show[DupGenerate] {
	return mshow.ContraGeneric(
		"mshowtest.DupGenerate",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`NoDerive`, ``), as.NameTag(`World`, ``)}, ShowNoDerive(), mshow.String),
		func(v DupGenerate) minimal.Tuple2[NoDerive, string] {
			return minimal.Tuple2[NoDerive, string]{
				I1: v.NoDerive,
				I2: v.World,
			}
		},
	)
}

func ShowHasTuple() mshow.Show[HasTuple] {
	return mshow.ContraGeneric(
		"mshowtest.HasTuple",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`Entry`, ``), as.NameTag(`HList`, ``)}, mshow.ContraGeneric(
			"fp.Tuple2",
			"Tuple",
			mshow.TupleHCons(
				mshow.String,
				mshow.TupleHCons(
					mshow.Int[int](),
					mshow.HNil,
				),
			),
			as.HList2[string, int],
		), mshow.HCons(mshow.String, mshow.HCons(mshow.Int[int](), mshow.HNil))),
		func(v HasTuple) minimal.Tuple2[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]] {
			return minimal.Tuple2[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]{
				I1: v.Entry,
				I2: v.HList,
			}
		},
	)
}

func ShowEmbeddedStruct() mshow.Show[EmbeddedStruct] {
	return mshow.ContraGeneric(
		"mshowtest.EmbeddedStruct",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`hello`, ``), as.NameTag(`world`, ``)}, mshow.String, mshow.ContraGeneric(
			"struct",
			"Struct",
			mshow.Struct2([]fp.Named{as.NameTag(`Level`, ``), as.NameTag(`Stage`, ``)}, mshow.Int[int](), mshow.String),
			func(v struct {
				Level int
				Stage string
			}) minimal.Tuple2[int, string] {
				return minimal.Tuple2[int, string]{
					I1: v.Level,
					I2: v.Stage,
				}
			},
		)),
		func(v EmbeddedStruct) minimal.Tuple2[string, struct {
			Level int
			Stage string
		}] {
			return minimal.Tuple2[string, struct {
				Level int
				Stage string
			}]{
				I1: v.hello,
				I2: v.world,
			}
		},
	)
}

func ShowEmbeddedTypeParamStruct[T any](showT mshow.Show[T]) mshow.Show[EmbeddedTypeParamStruct[T]] {
	return mshow.ContraGeneric(
		"mshowtest.EmbeddedTypeParamStruct",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`hello`, ``), as.NameTag(`world`, ``)}, mshow.String, mshow.ContraGeneric(
			"struct",
			"Struct",
			mshow.Struct2([]fp.Named{as.NameTag(`Level`, ``), as.NameTag(`Stage`, ``)}, showT, mshow.String),
			func(v struct {
				Level T
				Stage string
			}) minimal.Tuple2[T, string] {
				return minimal.Tuple2[T, string]{
					I1: v.Level,
					I2: v.Stage,
				}
			},
		)),
		func(v EmbeddedTypeParamStruct[T]) minimal.Tuple2[string, struct {
			Level T
			Stage string
		}] {
			return minimal.Tuple2[string, struct {
				Level T
				Stage string
			}]{
				I1: v.hello,
				I2: v.world,
			}
		},
	)
}

func ShowEmptyStruct() mshow.Show[EmptyStruct] {
	return mshow.ContraGeneric(
		"mshowtest.EmptyStruct",
		"Struct",
		mshow.HNil,
		func(EmptyStruct) hlist.Nil {
			return hlist.Empty()
		},
	)
}

func ShowHasAliasType() mshow.Show[HasAliasType] {
	return mshow.ContraGeneric(
		"mshowtest.HasAliasType",
		"Struct",
		mshow.Struct1([]fp.Named{as.NameTag(`Data`, ``)}, mshow.Slice(mshow.Int[byte]())),
		func(v HasAliasType) minimal.Tuple1[TLV] {
			return minimal.Tuple1[TLV]{
				I1: v.Data,
			}
		},
	)
}

//go:noinline
func ShowNoDerive() mshow.Show[NoDerive] {
	return mshow.ContraGeneric(
		"mshowtest.NoDerive",
		"Struct",
		mshow.Struct1([]fp.Named{as.NameTag(`Hello`, ``)}, mshow.String),
		func(v NoDerive) minimal.Tuple1[string] {
			return minimal.Tuple1[string]{
				I1: v.Hello,
			}
		},
	)
}

//go:noinline
func ShowRecursiveStringAlias() mshow.Show[recursive.StringAlias] {
	return mshow.ContraGeneric(
		"recursive.StringAlias",
		"NewType",
		mshow.String,
		func(v recursive.StringAlias) string {
			return string(v)
		},
	)
}
