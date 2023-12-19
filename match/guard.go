package match

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
)

// Case 로 시작하지 않는 함수들은  then 함수를 인자로 받지 않고,
// 타입을 분해하여,   타입안에 있는 값을 체크하는 guard 임

// And 로 끝나는 함수들은,  분해된 값을  다시 체크하는 guard를 인자로 받음.
func Any[T any]() fp.PartialFunc[T, T] {
	return as.PartialFunc(fp.Const[T](true), fp.Id[T])
}

func Equal[T comparable](v T) fp.PartialFunc[T, T] {
	return as.PartialFunc(eq.GivenValue(v), fp.Id[T])
}

func IsIn[T comparable](v ...T) fp.PartialFunc[T, T] {
	return as.PartialFunc(func(t T) bool {
		for _, i := range v {
			if t == i {
				return true
			}
		}

		return false
	}, fp.Id[T])
}

func Some[T any]() fp.PartialFunc[fp.Option[T], T] {
	return as.PartialFunc(fp.Option[T].IsDefined, fp.Option[T].Get)
}

func Success[T any]() fp.PartialFunc[fp.Try[T], T] {
	return as.PartialFunc(fp.Try[T].IsSuccess, fp.Try[T].Get)
}

func Left[L, R any]() fp.PartialFunc[fp.Either[L, R], L] {
	return as.PartialFunc(fp.Either[L, R].IsLeft, fp.Either[L, R].Left)
}

func Right[L, R any]() fp.PartialFunc[fp.Either[L, R], R] {
	return as.PartialFunc(fp.Either[L, R].IsRight, fp.Either[L, R].Get)

}

func LeftAnd[L, R, T2 any](down fp.PartialFunc[L, T2]) fp.PartialFunc[fp.Either[L, R], T2] {
	return combine(Left[L, R](), down)
}

func RightAnd[L, R, T2 any](down fp.PartialFunc[R, T2]) fp.PartialFunc[fp.Either[L, R], T2] {
	return combine(Right[L, R](), down)
}

func SuccessAnd[T, T2 any](down fp.PartialFunc[T, T2]) fp.PartialFunc[fp.Try[T], T2] {
	return combine(Success[T](), down)
}

func FailureAnd[T, T2 any](down fp.PartialFunc[error, T2]) fp.PartialFunc[fp.Try[T], T2] {
	return combine(Failure[T](), down)
}

func Failure[T any]() fp.PartialFunc[fp.Try[T], error] {
	return as.PartialFunc(fp.Try[T].IsFailure, func(t fp.Try[T]) error {
		return t.Failed().Get()
	})
}

func SomeAnd[T, T2 any](down fp.PartialFunc[T, T2]) fp.PartialFunc[fp.Option[T], T2] {
	return combine(Some[T](), down)
}

func None[T any]() fp.PartialFunc[fp.Option[T], fp.Unit] {
	return as.PartialFunc(fp.Option[T].IsEmpty, fp.Const[fp.Option[T]](fp.Unit{}))
}

func Tuple2[A1, A2, B1, B2 any](adown fp.PartialFunc[A1, A2], bdown fp.PartialFunc[B1, B2]) fp.PartialFunc[fp.Tuple2[A1, B1], fp.Tuple2[A2, B2]] {
	return as.PartialFunc(func(t fp.Tuple2[A1, B1]) bool {
		return adown.IsDefined(t.I1) && bdown.IsDefined(t.I2)
	}, func(t fp.Tuple2[A1, B1]) fp.Tuple2[A2, B2] {
		return as.Tuple(adown.Apply(t.I1), bdown.Apply(t.I2))
	})
}

func Cons[C fp.Cons[H1, T1], H1, H2, T1, T2 any](adown fp.PartialFunc[H1, H2], bdown fp.PartialFunc[T1, T2]) fp.PartialFunc[C, fp.Tuple2[H2, T2]] {
	return as.PartialFunc(func(t C) bool {
		return adown.IsDefined(t.Head()) && bdown.IsDefined(t.Tail())
	}, func(t C) fp.Tuple2[H2, T2] {
		return as.Tuple(adown.Apply(t.Head()), bdown.Apply(t.Tail()))
	})
}

func Head[C fp.Cons[H1, T1], H1, H2, T1 any](hdown fp.PartialFunc[H1, H2]) fp.PartialFunc[C, H2] {
	return as.PartialFunc(func(c C) bool {
		return hdown.IsDefined(c.Head())
	}, func(c C) H2 {
		return hdown.Apply(c.Head())
	})
}

func SeqHead[T, T2 any](hdown fp.PartialFunc[fp.Option[T], T2]) fp.PartialFunc[fp.Seq[T], T2] {
	return as.PartialFunc(func(c fp.Seq[T]) bool {
		return hdown.IsDefined(c.Head())
	}, func(c fp.Seq[T]) T2 {
		return hdown.Apply(c.Head())
	})
}
