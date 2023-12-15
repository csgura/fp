package match

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

// Case 로 시작하지 않는 함수들은  then 함수를 인자로 받지 않고,
// 타입을 분해하여,   타입안에 있는 값을 체크하는 guard 임

// And 로 끝나는 함수들은,  분해된 값을  다시 체크하는 guard를 인자로 받음.
func Any[T any](v T) fp.Option[T] {
	return option.Some(v)
}

func Equal[T comparable](v T) PartialFunction[T, T] {
	return func(t T) fp.Option[T] {
		if v == t {
			return option.Some(t)
		}
		return option.None[T]()
	}
}

func IsIn[T comparable](v ...T) PartialFunction[T, T] {
	return func(t T) fp.Option[T] {
		for _, i := range v {
			if t == i {
				return option.Some(t)
			}
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

func LeftAnd[L, R, T2 any](down PartialFunction[L, T2]) PartialFunction[fp.Either[L, R], T2] {
	return fp.Compose(Left[L, R], option.LiftM(down))
}

func RightAnd[L, R, T2 any](down PartialFunction[R, T2]) PartialFunction[fp.Either[L, R], T2] {
	return fp.Compose(Right[L], option.LiftM(down))
}

func SuccessAnd[T, T2 any](down PartialFunction[T, T2]) PartialFunction[fp.Try[T], T2] {
	return fp.Compose(Success, option.LiftM(down))
}

func FailureAnd[T, T2 any](down PartialFunction[error, T2]) PartialFunction[fp.Try[T], T2] {
	return fp.Compose(Failure[T], option.LiftM(down))
}

func Failure[T any](v fp.Try[T]) fp.Option[error] {
	return option.FromTry(v.Failed())
}

func SomeAnd[T, T2 any](down PartialFunction[T, T2]) PartialFunction[fp.Option[T], T2] {
	return fp.Compose(Some, option.LiftM(down))
}

func None[T any](v fp.Option[T]) fp.Option[fp.Unit] {
	if v.IsEmpty() {
		return option.Some(fp.Unit{})
	}
	return option.None[fp.Unit]()
}

func Tuple2[A1, A2, B1, B2 any](adown PartialFunction[A1, A2], bdown PartialFunction[B1, B2]) PartialFunction[fp.Tuple2[A1, B1], fp.Tuple2[A2, B2]] {
	return func(t fp.Tuple2[A1, B1]) fp.Option[fp.Tuple2[A2, B2]] {
		return option.Map2(adown(t.I1), bdown(t.I2), as.Tuple)
	}
}

func Cons[C fp.Cons[H1, T1], H1, H2, T1, T2 any](adown PartialFunction[H1, H2], bdown PartialFunction[T1, T2]) PartialFunction[C, fp.Tuple2[H2, T2]] {
	return func(t C) fp.Option[fp.Tuple2[H2, T2]] {
		return option.Map2(adown(t.Head()), bdown(t.Tail()), as.Tuple)
	}
}

func Head[C fp.Cons[H1, T1], H1, H2, T1 any](hdown PartialFunction[H1, H2]) PartialFunction[C, H2] {
	return func(c C) fp.Option[H2] {
		return hdown(c.Head())
	}
}

func SeqHead[T, T2 any](hdown PartialFunction[fp.Option[T], T2]) PartialFunction[fp.Seq[T], T2] {
	return func(c fp.Seq[T]) fp.Option[T2] {
		return hdown(c.Head())
	}
}
