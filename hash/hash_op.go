//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package hash

import (
	"hash/fnv"
	"hash/maphash"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
)

func hashUint64(value uint64) uint32 {
	hash := value
	for value > 0xffffffff {
		value /= 0xffffffff
		hash ^= value
	}
	return uint32(hash)
}

type hasher[T any] struct {
	fp.Eq[T]
	f func(T) uint32
}

func (r hasher[T]) Hash(t T) uint32 {
	return r.f(t)
}

func New[T any](eq fp.Eq[T], f func(T) uint32) fp.Hashable[T] {
	return hasher[T]{eq, f}
}

func Number[T fp.ImplicitNum]() fp.Hashable[T] {
	return New(eq.Given[T](), func(key T) uint32 {
		return hashUint64(uint64(key))
	})
}

func Given[T comparable]() fp.Hashable[T] {
	seed := maphash.MakeSeed()
	return GivenSeed[T](seed)
}

func GivenSeed[T comparable](seed maphash.Seed) fp.Hashable[T] {
	return New(eq.Given[T](), func(t T) uint32 {
		return hashUint64(maphash.Comparable(seed, t))
	})
}

const prime32 = 16777619
const offset32 = 2166136261

var String fp.Hashable[string] = New(eq.Given[string](), func(value string) uint32 {

	var hash uint32 = offset32
	for i, value := 0, value; i < len(value); i++ {
		hash *= prime32
		hash ^= uint32(value[i])
	}
	return hash
})

var Bytes fp.Hashable[[]byte] = New(eq.Bytes, func(value []byte) uint32 {
	w := fnv.New32()
	w.Write(value)
	return w.Sum32()

	// var hash uint32
	// for i, value := 0, value; i < len(value); i++ {
	// 	hash = 31*hash + uint32(value[i])
	// }
	// return hash
})

func Tuple1[A1 any](ins1 fp.Hashable[A1]) fp.Hashable[fp.Tuple1[A1]] {
	return New(eq.Tuple1[A1](ins1), func(t fp.Tuple1[A1]) uint32 {
		return ins1.Hash(t.Head())
	})
}

var HNil fp.Hashable[hlist.Nil] = New(eq.HNil, func(t hlist.Nil) uint32 {
	return 0
})

func HCons[H any, T hlist.HList](heq fp.Hashable[H], teq fp.Hashable[T]) fp.Hashable[hlist.Cons[H, T]] {
	return New(eq.HCons[H, T](heq, teq), func(a hlist.Cons[H, T]) uint32 {
		if hlist.IsNil(hlist.Tail(a)) {
			return heq.Hash(a.Head())
		}
		return heq.Hash(a.Head())*31 + teq.Hash(hlist.Tail(a))
	})
}

func Seq[T any](hashT fp.Hashable[T]) fp.Hashable[fp.Seq[T]] {
	return New(eq.Seq[T](hashT), func(a fp.Seq[T]) uint32 {
		var h uint32
		for _, t := range a {
			h = h*31 + hashT.Hash(t)
		}
		return h
	})
}

func Slice[T any](hashT fp.Hashable[T]) fp.Hashable[[]T] {
	return New(eq.Slice[T](hashT), func(a fp.Slice[T]) uint32 {
		var h uint32
		for _, t := range a {
			h = h*31 + hashT.Hash(t)
		}
		return h
	})
}

func Ptr[T any](hashT lazy.Eval[fp.Hashable[T]]) fp.Hashable[*T] {
	return New(eq.Ptr(lazy.Call(func() fp.Eq[T] {
		return hashT.Get()
	})), func(a *T) uint32 {
		if a == nil {
			return 0
		}

		return hashT.Get().Hash(*a)
	})
}

func Option[T any](hashT fp.Hashable[T]) fp.Hashable[fp.Option[T]] {
	return New(eq.Option[T](hashT), func(a fp.Option[T]) uint32 {
		if a.IsEmpty() {
			return 0
		}

		return hashT.Hash(a.Get())
	})
}

func ContraMap[T, U any](teq fp.Hashable[T], fn func(U) T) fp.Hashable[U] {
	return New(eq.ContraMap[T](teq, fn), func(a U) uint32 {
		return teq.Hash(fn(a))
	})
}

type Derives[T any] interface {
	Target() T
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/eq", Name: "eq"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Hashable"}} ) fp.Hashable[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New( eq.New( func( a, b fp.{{TupleType .N}} ) bool {
		return ins1.Eqv(a.Head(),b.Head()) && pt.Eqv(as.Tuple{{dec .N}}(a.Tail()), as.Tuple{{dec .N}}(b.Tail()))
	}), func(t fp.{{TupleType .N}} ) uint32 {
		return ins1.Hash(t.Head()) * 31 + pt.Hash(as.Tuple{{dec .N}}(t.Tail()))
	})
}
	`,
}
