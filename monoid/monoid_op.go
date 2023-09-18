//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package monoid

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/semigroup"
	"github.com/csgura/fp/try"
)

func New[T any](zero fp.EmptyFunc[T], combine fp.SemigroupFunc[T]) fp.Monoid[T] {
	return monoid[T]{
		zero, combine,
	}
}

var String = New(fp.Zero[string], func(a, b string) string {
	return a + b
})

func Sum[T fp.ImplicitOrd]() fp.Monoid[T] {
	return fp.SemigroupFunc[T](func(a, b T) T {
		return a + b
	})
}

func Product[T fp.ImplicitNum]() fp.Monoid[T] {
	return New(
		func() T {
			return 1
		},
		func(a, b T) T {
			return a * b
		},
	)
}

func Option[T any](m fp.Monoid[T]) fp.Monoid[fp.Option[T]] {
	return New(
		func() fp.Option[T] {
			return option.Some(m.Empty())
		},
		func(a fp.Option[T], b fp.Option[T]) fp.Option[T] {
			return option.Map2(a, b, m.Combine)
		},
	)
}

func Try[T any](m fp.Monoid[T]) fp.Monoid[fp.Try[T]] {
	return New(
		func() fp.Try[T] {
			return try.Success(m.Empty())
		},
		func(a fp.Try[T], b fp.Try[T]) fp.Try[T] {
			return try.Map2(a, b, m.Combine)
		},
	)
}

func Future[T any](m fp.Monoid[T]) fp.Monoid[fp.Future[T]] {
	return New(
		func() fp.Future[T] {
			return future.Successful(m.Empty())
		},
		func(a fp.Future[T], b fp.Future[T]) fp.Future[T] {
			return future.Map2(a, b, m.Combine)
		},
	)
}

func MergeSeq[T any]() fp.Monoid[fp.Seq[T]] {
	return New(
		fp.Zero[fp.Seq[T]],
		func(a fp.Seq[T], b fp.Seq[T]) fp.Seq[T] {
			return a.Concat(b)
		},
	)
}

func MergeSlice[T any]() fp.Monoid[[]T] {
	return IMap(MergeSeq[T](), func(b fp.Seq[T]) []T {
		return b
	}, as.Seq[T])
}

var HNil fp.Monoid[hlist.Nil] = fp.SemigroupFunc[hlist.Nil](func(a, b hlist.Nil) hlist.Nil {
	return hlist.Nil{}
})

func HCons[H any, T hlist.HList](hm fp.Monoid[H], tm fp.Monoid[T]) fp.Monoid[hlist.Cons[H, T]] {
	return New(
		func() hlist.Cons[H, T] {
			return hlist.Concat(hm.Empty(), tm.Empty())
		},
		func(a, b hlist.Cons[H, T]) hlist.Cons[H, T] {
			return hlist.Concat(hm.Combine(a.Head(), b.Head()), tm.Combine(a.Tail(), b.Tail()))
		},
	)
}

type monoid[T any] struct {
	zero    fp.EmptyFunc[T]
	combine fp.SemigroupFunc[T]
}

func (r monoid[T]) Empty() T {
	return r.zero()
}

func (r monoid[T]) Combine(a, b T) T {
	return r.combine(a, b)
}

func (r monoid[T]) ToMonoid(emptyFunc fp.EmptyFunc[T]) fp.Monoid[T] {
	return monoid[T]{emptyFunc, r.combine}
}

func (r monoid[T]) Curried() func(T) func(T) T {
	return r.combine.Curried()
}

func Endo[T any]() fp.Monoid[fp.Endo[T]] {
	return New(
		func() fp.Endo[T] {
			return fp.Id
		},
		semigroup.Endo[T]().Combine,
	)
}

func Dual[T any](m fp.Monoid[T]) fp.Monoid[fp.Dual[T]] {
	return New(
		func() fp.Dual[T] {
			return fp.Dual[T]{m.Empty()}
		},
		semigroup.Dual[T](m).Combine,
	)
}

func Eval[T any](m fp.Monoid[T]) fp.Monoid[lazy.Eval[T]] {
	return New(
		func() lazy.Eval[T] {
			return lazy.Done(m.Empty())
		},
		semigroup.Eval[T](m).Combine,
	)
}

var Any fp.Monoid[bool] = New(
	func() bool {
		return false
	},
	semigroup.Any.Combine,
)

var All fp.Monoid[bool] = New(
	func() bool {
		return true
	},
	semigroup.All.Combine,
)

func IMap[A, B any](instance fp.Monoid[A], fab func(A) B, fba func(B) A) fp.Monoid[B] {
	return New(func() B {
		return fab(instance.Empty())
	}, func(a, b B) B {
		return fab(instance.Combine(fba(a), fba(b)))
	})
}

type Derives[T any] interface {
	Target() T
}

func MergeMap[K, V any]() fp.Monoid[fp.Map[K, V]] {
	return New(
		fp.Zero[fp.Map[K, V]],
		func(a, b fp.Map[K, V]) fp.Map[K, V] {
			return a.Concat(b)
		})
}

func MergeSet[V any]() fp.Monoid[fp.Set[V]] {
	return New(
		fp.Zero[fp.Set[V]],
		func(a, b fp.Set[V]) fp.Set[V] {
			return a.Concat(b)
		})
}

func MergeGoMap[K comparable, V any]() fp.Monoid[map[K]V] {
	return New(func() map[K]V {
		return map[K]V{}
	}, func(a, b map[K]V) map[K]V {
		ret := map[K]V{}

		for k, v := range a {
			ret[k] = v
		}

		for k, v := range b {
			ret[k] = v
		}

		return ret
	})
}

func Ptr[T any](monoidT lazy.Eval[fp.Monoid[T]]) fp.Monoid[*T] {
	return New(
		fp.Zero[*T],
		func(a, b *T) *T {
			if a != nil && b != nil {
				ret := monoidT.Get().Combine(*a, *b)
				return &ret
			}
			if a == nil {
				return b
			}
			return a
		})
}

var Unit = New(fp.Zero[fp.Unit], func(a, b fp.Unit) fp.Unit {
	return fp.Unit{}
})

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/product", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]({{DeclTypeClassArgs 1 .N "fp.Monoid"}}) fp.Monoid[fp.{{TupleType .N}}] {
	return New(
		func() fp.{{TupleType .N}} {
			return product.Tuple{{.N}}({{ (Args "ins" 1 .N).Dot "Empty()"}})
		},
		func(t1 fp.{{TupleType .N}}, t2 fp.{{TupleType .N}}) fp.{{TupleType .N}} {
			return product.Tuple{{.N}}(
				{{- range $idx := Range 1 .N -}}
				ins{{$idx}}.Combine(t1.I{{$idx}}, t2.I{{$idx}}),
				{{- end -}}
			)
		},
	)
}
	`,
}
