// Code generated by gombok, DO NOT EDIT.
package clonetest

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/clone"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"time"
)

func CloneCloneStruct() fp.Clone[CloneStruct] {
	return clone.Generic(
		fp.Generic[CloneStruct, fp.Tuple2[string, int]]{
			Type: "clonetest.CloneStruct",
			Kind: "Struct",
			To:   CloneStruct.AsTuple,
			From: fp.Compose(
				as.Curried2(CloneStructBuilder.FromTuple)(CloneStructBuilder{}),
				CloneStructBuilder.Build,
			),
		},
		clone.Tuple2(clone.Given[string](), clone.Given[int]()),
	)
}

func CloneHasReference() fp.Clone[HasReference] {
	return clone.Generic(
		fp.Generic[HasReference, fp.Tuple8[*string, []int, map[string]int, RecursiveDerive, time.Time, MySeq, ValueStruct, CloneStruct]]{
			Type: "clonetest.HasReference",
			Kind: "Struct",
			To: func(v HasReference) fp.Tuple8[*string, []int, map[string]int, RecursiveDerive, time.Time, MySeq, ValueStruct, CloneStruct] {
				return fp.Tuple8[*string, []int, map[string]int, RecursiveDerive, time.Time, MySeq, ValueStruct, CloneStruct]{
					I1: v.A,
					I2: v.S,
					I3: v.M,
					I4: v.RD,
					I5: v.T,
					I6: v.MS,
					I7: v.VS,
					I8: v.CS,
				}
			},
			From: func(t fp.Tuple8[*string, []int, map[string]int, RecursiveDerive, time.Time, MySeq, ValueStruct, CloneStruct]) HasReference {
				return HasReference{
					A:  t.I1,
					S:  t.I2,
					M:  t.I3,
					RD: t.I4,
					T:  t.I5,
					MS: t.I6,
					VS: t.I7,
					CS: t.I8,
				}
			},
		},
		clone.Tuple8(clone.Ptr(lazy.Call(func() fp.Clone[string] {
			return clone.Given[string]()
		})), clone.Slice(clone.Given[int]()), clone.GoMap(clone.Given[string](), clone.Given[int]()), CloneRecursiveDerive(), clone.Given[time.Time](), CloneMySeq(), CloneValueStruct(), CloneCloneStruct()),
	)
}

func CloneRecursiveDerive() fp.Clone[RecursiveDerive] {
	return clone.Generic(
		fp.Generic[RecursiveDerive, hlist.Cons[[]string, hlist.Nil]]{
			Type: "clonetest.RecursiveDerive",
			Kind: "Struct",
			To: func(v RecursiveDerive) hlist.Cons[[]string, hlist.Nil] {
				i0 := v.S
				h1 := hlist.Empty()
				h0 := hlist.Concat(i0, h1)
				return h0
			},
			From: func(hl0 hlist.Cons[[]string, hlist.Nil]) RecursiveDerive {
				i0 := hlist.Head(hl0)
				return RecursiveDerive{S: i0}
			},
		},
		clone.HCons(
			clone.Slice(clone.Given[string]()),
			clone.HNil,
		),
	)
}

func CloneMySeq() fp.Clone[MySeq] {
	return clone.Generic(
		fp.Generic[MySeq, []string]{
			Type: "clonetest.MySeq",
			Kind: "NewType",
			To: func(v MySeq) []string {
				return []string(v)
			},
			From: func(v []string) MySeq {
				return MySeq(v)
			},
		},
		clone.Slice(clone.Given[string]()),
	)
}

func CloneValueStruct() fp.Clone[ValueStruct] {
	return clone.Generic(
		fp.Generic[ValueStruct, fp.Tuple2[string, int]]{
			Type: "clonetest.ValueStruct",
			Kind: "Struct",
			To:   ValueStruct.AsTuple,
			From: fp.Compose(
				as.Curried2(ValueStructBuilder.FromTuple)(ValueStructBuilder{}),
				ValueStructBuilder.Build,
			),
		},
		clone.Tuple2(clone.Given[string](), clone.Given[int]()),
	)
}
