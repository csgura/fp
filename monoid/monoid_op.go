package monoid

//go:generate go run github.com/csgura/fp/internal/generator/monoid_gen
import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
)

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

func New[T any](zero fp.EmptyFunc[T], combine fp.SemigroupFunc[T]) fp.Monoid[T] {
	return monoid[T]{
		zero, combine,
	}
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
			return a.Concact(b)
		},
	)
}
