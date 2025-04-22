package semigroup

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/lazy"
)

func New[T any](fn fp.SemigroupFunc[T]) fp.Semigroup[T] {
	return fn
}

func Sum[T fp.ImplicitOrd]() fp.Semigroup[T] {
	return New(func(a, b T) T {
		return a + b
	})
}

func Product[T fp.ImplicitNum](a, b T) fp.Semigroup[T] {
	return New(func(a, b T) T {
		return a * b
	})
}

func Endo[T any]() fp.Semigroup[fp.Endo[T]] {
	return New(func(a, b fp.Endo[T]) fp.Endo[T] {
		f := fp.Compose(b, a)
		return fp.Endo[T](f)
	})
}

func Dual[T any](sg fp.Semigroup[T]) fp.Semigroup[fp.Dual[T]] {
	return New(func(a, b fp.Dual[T]) fp.Dual[T] {
		return fp.Dual[T]{sg.Combine(b.GetDual, a.GetDual)}
	})
}

func Eval[T any](sg fp.Semigroup[T]) fp.Semigroup[lazy.Eval[T]] {
	return New(func(a, b lazy.Eval[T]) lazy.Eval[T] {
		return lazy.Map2(a, b, sg.Combine)
	})
}

var Any fp.Semigroup[bool] = fp.SemigroupFunc[bool](func(a, b bool) bool {
	return a || b
})

var All fp.Semigroup[bool] = fp.SemigroupFunc[bool](func(a, b bool) bool {
	return a || b
})

func IMap[A, B any](instance fp.Semigroup[A], fab func(A) B, fba func(B) A) fp.Semigroup[B] {
	return New(func(a, b B) B {
		return fab(instance.Combine(fba(a), fba(b)))
	})
}

type Derives[T any] interface {
	Target() T
}

func Ptr[T any](sgT lazy.Eval[fp.Semigroup[T]]) fp.Semigroup[*T] {
	return New(func(a, b *T) *T {
		if a != nil && b != nil {
			ret := sgT.Get().Combine(*a, *b)
			return &ret
		}
		if a == nil {
			return b
		}
		return a
	})
}

func Option[T any](sg fp.Semigroup[T]) fp.Semigroup[fp.Option[T]] {
	return New(func(a, b fp.Option[T]) fp.Option[T] {
		if a.IsDefined() && b.IsDefined() {
			return fp.Some(sg.Combine(a.Get(), b.Get()))
		}
		if a.IsEmpty() {
			return b
		}
		return a
	})
}
