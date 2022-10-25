package semigroup

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/lazy"
)

func Sum[T fp.ImplicitOrd]() fp.SemigroupFunc[T] {
	return func(a, b T) T {
		return a + b
	}
}

func Product[T fp.ImplicitNum](a, b T) fp.SemigroupFunc[T] {
	return func(a, b T) T {
		return a * b
	}
}

func Endo[T any]() fp.SemigroupFunc[fp.Endo[T]] {
	return func(a, b fp.Endo[T]) fp.Endo[T] {
		f := fp.Compose(b.AsFunc(), a.AsFunc())
		return fp.Endo[T](f)
	}
}

func Dual[T any](sg fp.Semigroup[T]) fp.SemigroupFunc[fp.Dual[T]] {
	return func(a, b fp.Dual[T]) fp.Dual[T] {
		return fp.Dual[T]{sg.Combine(b.GetDual, a.GetDual)}
	}
}

func Eval[T any](sg fp.Semigroup[T]) fp.SemigroupFunc[lazy.Eval[T]] {
	return func(a, b lazy.Eval[T]) lazy.Eval[T] {
		return lazy.Map2(a, b, sg.Combine)
	}
}

var Any fp.SemigroupFunc[bool] = fp.SemigroupFunc[bool](func(a, b bool) bool {
	return a || b
})

var All fp.SemigroupFunc[bool] = fp.SemigroupFunc[bool](func(a, b bool) bool {
	return a || b
})

func IMap[A, B any](instance fp.Semigroup[A], fab func(A) B, fba func(B) A) fp.SemigroupFunc[B] {
	return func(a, b B) B {
		return fab(instance.Combine(fba(a), fba(b)))
	}
}

type Derives[T any] interface {
	Target() T
}
