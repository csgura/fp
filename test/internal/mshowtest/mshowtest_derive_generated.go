// Code generated by gombok, DO NOT EDIT.
package mshowtest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/minimal"
	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/test/internal/recursive"
)

func ShowPerson() mshow.Show[Person] {
	return mshow.Generic(
		fp.Generic[Person, minimal.Tuple2[string, int]]{
			Type: "mshowtest.Person",
			Kind: "Struct",
			To: func(v Person) minimal.Tuple2[string, int] {
				return minimal.Tuple2[string, int]{
					I1: v.Name,
					I2: v.Age,
				}
			},
			From: func(t minimal.Tuple2[string, int]) Person {
				return Person{
					Name: t.I1,
					Age:  t.I2,
				}
			},
		},
		mshow.Struct2([]fp.Named{as.NameTag(`Name`, ``), as.NameTag(`Age`, ``)}, mshow.String, mshow.Int[int]()),
	)
}

func ShowCollection() mshow.Show[Collection] {
	return mshow.Generic(
		fp.Generic[Collection, minimal.Tuple11[map[string]Person, []Person, *string, fp.Set[int], fp.Option[Person], NoDerive, HasStringMethod, *bool, map[string]NoDerive, recursive.StringAlias, fp.Seq[string]]]{
			Type: "mshowtest.Collection",
			Kind: "Struct",
			To: func(v Collection) minimal.Tuple11[map[string]Person, []Person, *string, fp.Set[int], fp.Option[Person], NoDerive, HasStringMethod, *bool, map[string]NoDerive, recursive.StringAlias, fp.Seq[string]] {
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
			From: func(t minimal.Tuple11[map[string]Person, []Person, *string, fp.Set[int], fp.Option[Person], NoDerive, HasStringMethod, *bool, map[string]NoDerive, recursive.StringAlias, fp.Seq[string]]) Collection {
				return Collection{
					Index:       t.I1,
					List:        t.I2,
					Description: t.I3,
					Set:         t.I4,
					Option:      t.I5,
					NoDerive:    t.I6,
					Stringer:    t.I7,
					BoolPtr:     t.I8,
					NoMap:       t.I9,
					Alias:       t.I10,
					StringSeq:   t.I11,
				}
			},
		},
		mshow.Struct11([]fp.Named{as.NameTag(`Index`, ``), as.NameTag(`List`, ``), as.NameTag(`Description`, ``), as.NameTag(`Set`, ``), as.NameTag(`Option`, ``), as.NameTag(`NoDerive`, ``), as.NameTag(`Stringer`, ``), as.NameTag(`BoolPtr`, ``), as.NameTag(`NoMap`, ``), as.NameTag(`Alias`, ``), as.NameTag(`StringSeq`, ``)}, mshow.GoMap(mshow.String, ShowPerson()), mshow.Slice(ShowPerson()), mshow.Ptr(lazy.Call(func() mshow.Show[string] {
			return mshow.String
		})), mshow.Set(mshow.Int[int]()), mshow.Option(ShowPerson()), ShowNoDerive(), mshow.Given[HasStringMethod](), mshow.Ptr(lazy.Call(func() mshow.Show[bool] {
			return mshow.Bool
		})), mshow.GoMap(mshow.String, ShowNoDerive()), ShowRecursiveStringAlias(), mshow.Seq(mshow.String)),
	)
}

func ShowDupGenerate() mshow.Show[DupGenerate] {
	return mshow.Generic(
		fp.Generic[DupGenerate, minimal.Tuple2[NoDerive, string]]{
			Type: "mshowtest.DupGenerate",
			Kind: "Struct",
			To: func(v DupGenerate) minimal.Tuple2[NoDerive, string] {
				return minimal.Tuple2[NoDerive, string]{
					I1: v.NoDerive,
					I2: v.World,
				}
			},
			From: func(t minimal.Tuple2[NoDerive, string]) DupGenerate {
				return DupGenerate{
					NoDerive: t.I1,
					World:    t.I2,
				}
			},
		},
		mshow.Struct2([]fp.Named{as.NameTag(`NoDerive`, ``), as.NameTag(`World`, ``)}, ShowNoDerive(), mshow.String),
	)
}

func ShowHasTuple() mshow.Show[HasTuple] {
	return mshow.Generic(
		fp.Generic[HasTuple, minimal.Tuple2[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]]{
			Type: "mshowtest.HasTuple",
			Kind: "Struct",
			To: func(v HasTuple) minimal.Tuple2[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]] {
				return minimal.Tuple2[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]{
					I1: v.Entry,
					I2: v.HList,
				}
			},
			From: func(t minimal.Tuple2[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]) HasTuple {
				return HasTuple{
					Entry: t.I1,
					HList: t.I2,
				}
			},
		},
		mshow.Struct2([]fp.Named{as.NameTag(`Entry`, ``), as.NameTag(`HList`, ``)}, mshow.Generic(
			fp.Generic[fp.Tuple2[string, int], hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]{
				Type: "fp.Tuple2",
				Kind: "Tuple",
				To:   as.HList2[string, int],
				From: product.TupleFromHList2[string, int],
			},
			mshow.TupleHCons(
				mshow.String,
				mshow.TupleHCons(
					mshow.Int[int](),
					mshow.HNil,
				),
			),
		), mshow.HCons(mshow.String, mshow.HCons(mshow.Int[int](), mshow.HNil))),
	)
}

func ShowEmbeddedStruct() mshow.Show[EmbeddedStruct] {
	return mshow.Generic(
		fp.Generic[EmbeddedStruct, minimal.Tuple2[string, struct {
			Level int
			Stage string
		}]]{
			Type: "mshowtest.EmbeddedStruct",
			Kind: "Struct",
			To: func(v EmbeddedStruct) minimal.Tuple2[string, struct {
				Level int
				Stage string
			}] { return minimal.Tuple2[string, struct {
				Level int
				Stage string
			}]{
				I1: v.hello,
				I2: v.world,
			} },
			From: func(t minimal.Tuple2[string, struct {
				Level int
				Stage string
			}]) EmbeddedStruct {
				return EmbeddedStruct{
					hello: t.I1,
					world: t.I2,
				}
			},
		},
		mshow.Struct2([]fp.Named{as.NameTag(`hello`, ``), as.NameTag(`world`, ``)}, mshow.String, mshow.Generic(
			fp.Generic[struct {
				Level int
				Stage string
			}, minimal.Tuple2[int, string]]{
				Type: "struct",
				Kind: "Struct",
				To: func(v struct {
					Level int
					Stage string
				}) minimal.Tuple2[int, string] {
					return minimal.AsTuple2(v.Level, v.Stage)
				},
				From: func(t minimal.Tuple2[int, string]) struct {
					Level int
					Stage string
				} {
					return struct {
						Level int
						Stage string
					}{
						Level: t.I1,
						Stage: t.I2,
					}
				},
			},
			mshow.Struct2([]fp.Named{as.NameTag(`Level`, ``), as.NameTag(`Stage`, ``)}, mshow.Int[int](), mshow.String),
		)),
	)
}

func ShowEmbeddedTypeParamStruct[T any](showT mshow.Show[T]) mshow.Show[EmbeddedTypeParamStruct[T]] {
	return mshow.Generic(
		fp.Generic[EmbeddedTypeParamStruct[T], minimal.Tuple2[string, struct {
			Level T
			Stage string
		}]]{
			Type: "mshowtest.EmbeddedTypeParamStruct",
			Kind: "Struct",
			To: func(v EmbeddedTypeParamStruct[T]) minimal.Tuple2[string, struct {
				Level T
				Stage string
			}] { return minimal.Tuple2[string, struct {
				Level T
				Stage string
			}]{
				I1: v.hello,
				I2: v.world,
			} },
			From: func(t minimal.Tuple2[string, struct {
				Level T
				Stage string
			}]) EmbeddedTypeParamStruct[T] {
				return EmbeddedTypeParamStruct[T]{
					hello: t.I1,
					world: t.I2,
				}
			},
		},
		mshow.Struct2([]fp.Named{as.NameTag(`hello`, ``), as.NameTag(`world`, ``)}, mshow.String, mshow.Generic(
			fp.Generic[struct {
				Level T
				Stage string
			}, minimal.Tuple2[T, string]]{
				Type: "struct",
				Kind: "Struct",
				To: func(v struct {
					Level T
					Stage string
				}) minimal.Tuple2[T, string] {
					return minimal.AsTuple2(v.Level, v.Stage)
				},
				From: func(t minimal.Tuple2[T, string]) struct {
					Level T
					Stage string
				} {
					return struct {
						Level T
						Stage string
					}{
						Level: t.I1,
						Stage: t.I2,
					}
				},
			},
			mshow.Struct2([]fp.Named{as.NameTag(`Level`, ``), as.NameTag(`Stage`, ``)}, showT, mshow.String),
		)),
	)
}

func ShowEmptyStruct() mshow.Show[EmptyStruct] {
	return mshow.Generic(
		fp.Generic[EmptyStruct, hlist.Nil]{
			Type: "mshowtest.EmptyStruct",
			Kind: "Struct",
			To: func(EmptyStruct) hlist.Nil {
				return hlist.Empty()
			},
			From: func(hlist.Nil) EmptyStruct {
				return EmptyStruct{}
			},
		},
		mshow.HNil,
	)
}

func ShowHasAliasType() mshow.Show[HasAliasType] {
	return mshow.Generic(
		fp.Generic[HasAliasType, minimal.Tuple1[TLV]]{
			Type: "mshowtest.HasAliasType",
			Kind: "Struct",
			To: func(v HasAliasType) minimal.Tuple1[TLV] {
				return minimal.Tuple1[TLV]{
					I1: v.Data,
				}
			},
			From: func(t minimal.Tuple1[TLV]) HasAliasType {
				return HasAliasType{
					Data: t.I1,
				}
			},
		},
		mshow.Struct1([]fp.Named{as.NameTag(`Data`, ``)}, mshow.Slice(mshow.Int[byte]())),
	)
}

func ShowNoDerive() mshow.Show[NoDerive] {
	return mshow.Generic(
		fp.Generic[NoDerive, minimal.Tuple1[string]]{
			Type: "mshowtest.NoDerive",
			Kind: "Struct",
			To: func(v NoDerive) minimal.Tuple1[string] {
				return minimal.Tuple1[string]{
					I1: v.Hello,
				}
			},
			From: func(t minimal.Tuple1[string]) NoDerive {
				return NoDerive{
					Hello: t.I1,
				}
			},
		},
		mshow.Struct1([]fp.Named{as.NameTag(`Hello`, ``)}, mshow.String),
	)
}

func ShowRecursiveStringAlias() mshow.Show[recursive.StringAlias] {
	return mshow.Generic(
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
		mshow.String,
	)
}
