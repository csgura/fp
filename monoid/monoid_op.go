//go:generate go run github.com/csgura/fp/internal/generator/monoid_gen
package monoid

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/semigroup"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
)

func New[T any](zero fp.EmptyFunc[T], combine fp.SemigroupFunc[T]) fp.Monoid[T] {
	return monoid[T]{
		zero, combine,
	}
}

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
			return option.Applicative2(m.Combine).ApOption(a).ApOption(b)
		},
	)
}

func Try[T any](m fp.Monoid[T]) fp.Monoid[fp.Try[T]] {
	return New(
		func() fp.Try[T] {
			return try.Success(m.Empty())
		},
		func(a fp.Try[T], b fp.Try[T]) fp.Try[T] {
			return try.Applicative2(m.Combine).ApTry(a).ApTry(b)
		},
	)
}

func Future[T any](m fp.Monoid[T]) fp.Monoid[fp.Future[T]] {
	return New(
		func() fp.Future[T] {
			return future.Successful(m.Empty())
		},
		func(a fp.Future[T], b fp.Future[T]) fp.Future[T] {
			return future.Applicative2(m.Combine).ApFuture(a).ApFuture(b)
		},
	)
}

func Seq[T any]() fp.Monoid[fp.Seq[T]] {
	return New(
		func() fp.Seq[T] {
			return seq.Of[T]()
		},
		func(a fp.Seq[T], b fp.Seq[T]) fp.Seq[T] {
			return a.Concat(b)
		},
	)
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

func (r monoid[T]) Curried() fp.Func1[T, fp.Func1[T, T]] {
	return r.combine.Curried()
}

func Endo[T any]() fp.Monoid[fp.Endo[T]] {
	return New(
		func() fp.Endo[T] {
			return fp.Id[T]
		},
		semigroup.Endo[T](),
	)
}

func Dual[T any](m fp.Monoid[T]) fp.Monoid[fp.Dual[T]] {
	return New(
		func() fp.Dual[T] {
			return fp.Dual[T]{m.Empty()}
		},
		semigroup.Dual[T](m),
	)
}

func Eval[T any](m fp.Monoid[T]) fp.Monoid[lazy.Eval[T]] {
	return New(
		func() lazy.Eval[T] {
			return lazy.Done(m.Empty())
		},
		semigroup.Eval[T](m),
	)
}

var Any fp.Monoid[bool] = New(
	func() bool {
		return false
	},
	semigroup.Any,
)

var All fp.Monoid[bool] = New(
	func() bool {
		return true
	},
	semigroup.All,
)
