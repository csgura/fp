package match

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

type CaseBlock[T, R any] interface {
	Match(T) fp.Option[R]
}

type MatcherFunc[T, R any] func(T) fp.Option[R]

func (r MatcherFunc[T, R]) Match(v T) fp.Option[R] {
	return r(v)
}

func Of[V, R any](v V, c ...CaseBlock[V, R]) R {
	for _, m := range c {
		opt := m.Match(v)
		if opt.IsDefined() {
			return opt.Get()
		}
	}
	panic("case not matched")
}

type Matcher[V, T any] func(V) fp.Option[T]

func Case[V, T, R any](guard Matcher[V, T], then func(T) R) CaseBlock[V, R] {
	return MatcherFunc[V, R](func(v V) fp.Option[R] {
		return option.Map(guard(v), then)
	})
}

func Any[T any](v T) fp.Option[T] {
	return option.Some(v)
}

func Equal[T comparable](v T) Matcher[T, T] {
	return func(t T) fp.Option[T] {
		if v == t {
			return option.Some(v)
		}
		return option.None[T]()
	}
}

func Some[T any](v fp.Option[T]) fp.Option[T] {
	return v
}

func Success[T any](v fp.Try[T]) fp.Option[T] {
	return option.FromTry(v)
}

func Left[L, R any](v fp.Either[L, R]) fp.Option[L] {
	if v.IsLeft() {
		return option.Some(v.Left())
	}
	return option.None[L]()
}

func Right[L, R any](v fp.Either[L, R]) fp.Option[R] {
	if v.IsRight() {
		return option.Some(v.Get())
	}
	return option.None[R]()
}

func LeftAnd[L, R, T2 any](down Matcher[L, T2]) Matcher[fp.Either[L, R], T2] {
	return fp.Compose(Left[L, R], option.LiftM(down))
}

func RightAnd[L, R, T2 any](down Matcher[R, T2]) Matcher[fp.Either[L, R], T2] {
	return fp.Compose(Right[L], option.LiftM(down))
}

func SuccessAnd[T, T2 any](down Matcher[T, T2]) Matcher[fp.Try[T], T2] {
	return fp.Compose(Success, option.LiftM(down))
}

func FailureAnd[T, T2 any](down Matcher[error, T2]) Matcher[fp.Try[T], T2] {
	return fp.Compose(Failure[T], option.LiftM(down))
}

func Failure[T any](v fp.Try[T]) fp.Option[error] {
	return option.FromTry(v.Failed())
}

func SomeAnd[T, T2 any](down Matcher[T, T2]) Matcher[fp.Option[T], T2] {
	return fp.Compose(Some, option.LiftM(down))
}

func None[T any](v fp.Option[T]) fp.Option[fp.Unit] {
	if v.IsEmpty() {
		return option.Some(fp.Unit{})
	}
	return option.None[fp.Unit]()
}

func Tuple2[A1, A2, B1, B2 any](adown Matcher[A1, A2], bdown Matcher[B1, B2]) Matcher[fp.Tuple2[A1, B1], fp.Tuple2[A2, B2]] {
	return func(t fp.Tuple2[A1, B1]) fp.Option[fp.Tuple2[A2, B2]] {
		return option.Map2(adown(t.I1), bdown(t.I2), as.Tuple)
	}
}
