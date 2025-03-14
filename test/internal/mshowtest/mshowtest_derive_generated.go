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
	"github.com/csgura/fp/test/internal/showtest"
)

func ShowShowtestPerson() mshow.Show[showtest.Person] {
	return mshow.ContraGeneric(
		"showtest.Person",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`Name`, ``), as.NameTag(`Age`, ``)}, mshow.String, mshow.Int[int]()),
		func(v showtest.Person) minimal.Tuple2[string, int] {
			return minimal.Tuple2[string, int]{
				I1: v.Name,
				I2: v.Age,
			}
		},
	)
}

//go:noinline
func ShowShowtestCollection() mshow.Show[showtest.Collection] {
	return mshow.ContraGeneric(
		"showtest.Collection",
		"Struct",
		mshow.Struct12([]fp.Named{as.NameTag(`Index`, ``), as.NameTag(`List`, ``), as.NameTag(`Description`, ``), as.NameTag(`Set`, ``), as.NameTag(`Option`, ``), as.NameTag(`NoDerive`, ``), as.NameTag(`Stringer`, ``), as.NameTag(`BoolPtr`, ``), as.NameTag(`NoMap`, ``), as.NameTag(`Alias`, ``), as.NameTag(`StringSeq`, ``), as.NameTag(`Array`, ``)}, mshow.GoMap(mshow.String, ShowShowtestPerson()), mshow.Slice(ShowShowtestPerson()), mshow.Ptr(lazy.Call(func() mshow.Show[string] {
			return mshow.String
		})), mshow.Set(mshow.Int[int]()), mshow.Option(ShowShowtestPerson()), ShowShowtestNoDerive(), mshow.Given[showtest.HasStringMethod](), mshow.Ptr(lazy.Call(func() mshow.Show[bool] {
			return mshow.Bool
		})), mshow.GoMap(mshow.String, ShowShowtestNoDerive()), ShowRecursiveStringAlias(), mshow.Seq(mshow.String), mshow.ContraGeneric(
			"[4]bool",
			"Conversion",
			mshow.Slice(mshow.Bool),
			func(v [4]bool) []bool {
				return v[:]
			},
		)),
		func(v showtest.Collection) minimal.Tuple12[map[string]showtest.Person, []showtest.Person, *string, fp.Set[int], fp.Option[showtest.Person], showtest.NoDerive, showtest.HasStringMethod, *bool, map[string]showtest.NoDerive, recursive.StringAlias, fp.Seq[string], [4]bool] {
			return minimal.Tuple12[map[string]showtest.Person, []showtest.Person, *string, fp.Set[int], fp.Option[showtest.Person], showtest.NoDerive, showtest.HasStringMethod, *bool, map[string]showtest.NoDerive, recursive.StringAlias, fp.Seq[string], [4]bool]{
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
				I12: v.Array,
			}
		},
	)
}

func ShowShowtestDupGenerate() mshow.Show[showtest.DupGenerate] {
	return mshow.ContraGeneric(
		"showtest.DupGenerate",
		"Struct",
		mshow.Struct2([]fp.Named{as.NameTag(`NoDerive`, ``), as.NameTag(`World`, ``)}, ShowShowtestNoDerive(), mshow.String),
		func(v showtest.DupGenerate) minimal.Tuple2[showtest.NoDerive, string] {
			return minimal.Tuple2[showtest.NoDerive, string]{
				I1: v.NoDerive,
				I2: v.World,
			}
		},
	)
}

func ShowShowtestHasTuple() mshow.Show[showtest.HasTuple] {
	return mshow.ContraGeneric(
		"showtest.HasTuple",
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
		), mshow.TupleHCons(mshow.String, mshow.TupleHCons(mshow.Int[int](), mshow.HNil))),
		func(v showtest.HasTuple) minimal.Tuple2[fp.Entry[int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]] {
			return minimal.Tuple2[fp.Entry[int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]{
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

func ShowShowtestEmptyStruct() mshow.Show[showtest.EmptyStruct] {
	return mshow.ContraGeneric(
		"showtest.EmptyStruct",
		"Struct",
		mshow.HNil,
		func(showtest.EmptyStruct) hlist.Nil {
			return hlist.Empty()
		},
	)
}

func ShowShowtestHasAliasType() mshow.Show[showtest.HasAliasType] {
	return mshow.ContraGeneric(
		"showtest.HasAliasType",
		"Struct",
		mshow.Struct1([]fp.Named{as.NameTag(`Data`, ``)}, mshow.Slice(mshow.Int[byte]())),
		func(v showtest.HasAliasType) minimal.Tuple1[showtest.TLV] {
			return minimal.Tuple1[showtest.TLV]{
				I1: v.Data,
			}
		},
	)
}

//go:noinline
func ShowShowtestNoDerive() mshow.Show[showtest.NoDerive] {
	return mshow.ContraGeneric(
		"showtest.NoDerive",
		"Struct",
		mshow.Struct1([]fp.Named{as.NameTag(`Hello`, ``)}, mshow.String),
		func(v showtest.NoDerive) minimal.Tuple1[string] {
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
