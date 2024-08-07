// Code generated by monad_gen, DO NOT EDIT.
package try

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
)

func PureOptionT[A any](a A) fp.Try[fp.Option[A]] {
	return Pure(option.Pure[A](a))
}

func LiftOptionT[A any](a fp.Try[A]) fp.Try[fp.Option[A]] {
	return Map(a, option.Pure[A])
}

func MapOptionT[A any, B any](t fp.Try[fp.Option[A]], f func(A) B) fp.Try[fp.Option[B]] {
	return Map(t, func(ma fp.Option[A]) fp.Option[B] {
		return option.FlatMap[A, B](ma, func(a A) fp.Option[B] {
			return option.Pure[B](f(a))
		})
	})
}

func SubFlatMapOptionT[A any, B any](t fp.Try[fp.Option[A]], f func(A) fp.Option[B]) fp.Try[fp.Option[B]] {
	return Map(t, func(ma fp.Option[A]) fp.Option[B] {
		return option.FlatMap[A, B](ma, func(a A) fp.Option[B] {
			return f(a)
		})
	})
}

func TraverseOptionT[A any, B any](t fp.Try[fp.Option[A]], f func(A) fp.Try[B]) fp.Try[fp.Option[B]] {
	sequencef := func(v fp.Option[fp.Try[B]]) fp.Try[fp.Option[B]] {
		if v.IsDefined() {
			return Map(v.Get(), option.Some)
		}
		return Success(fp.Option[B]{})
	}
	return FlatMap(MapOptionT(t, f), sequencef)
}

func FlatMapOptionT[A any, B any](t fp.Try[fp.Option[A]], f func(A) fp.Try[fp.Option[B]]) fp.Try[fp.Option[B]] {

	flatten := func(v fp.Option[fp.Option[B]]) fp.Option[B] {
		return option.FlatMap[fp.Option[B], B](v, fp.Id)
	}

	return Map(TraverseOptionT(t, f), flatten)

}

func FilterOptionT[T any](optionT fp.Try[fp.Option[T]], p func(v T) bool) fp.Try[fp.Option[T]] {
	return Map(optionT, func(insideValue fp.Option[T]) fp.Option[T] {
		return fp.Option[T].Filter(insideValue, p)
	})
}

func OrElseOptionT[T any](optionT fp.Try[fp.Option[T]], t T) fp.Try[T] {
	return Map(optionT, func(insideValue fp.Option[T]) T {
		return fp.Option[T].OrElse(insideValue, t)
	})
}

func OrZeroOptionT[T any](optionT fp.Try[fp.Option[T]]) fp.Try[T] {
	return Map(optionT, func(insideValue fp.Option[T]) T {
		return fp.Option[T].OrZero(insideValue)
	})
}

func OrElseGetOptionT[T any](optionT fp.Try[fp.Option[T]], f func() T) fp.Try[T] {
	return Map(optionT, func(insideValue fp.Option[T]) T {
		return fp.Option[T].OrElseGet(insideValue, f)
	})
}

func OrOptionT[T any](optionT fp.Try[fp.Option[T]], f func() fp.Option[T]) fp.Try[fp.Option[T]] {
	return Map(optionT, func(insideValue fp.Option[T]) fp.Option[T] {
		return fp.Option[T].Or(insideValue, f)
	})
}

func OrOptionOptionT[T any](optionT fp.Try[fp.Option[T]], v fp.Option[T]) fp.Try[fp.Option[T]] {
	return Map(optionT, func(insideValue fp.Option[T]) fp.Option[T] {
		return fp.Option[T].OrOption(insideValue, v)
	})
}

func OrPtrOptionT[T any](optionT fp.Try[fp.Option[T]], v *T) fp.Try[fp.Option[T]] {
	return Map(optionT, func(insideValue fp.Option[T]) fp.Option[T] {
		return fp.Option[T].OrPtr(insideValue, v)
	})
}

func RecoverOptionT[T any](optionT fp.Try[fp.Option[T]], f func() T) fp.Try[fp.Option[T]] {
	return Map(optionT, func(insideValue fp.Option[T]) fp.Option[T] {
		return fp.Option[T].Recover(insideValue, f)
	})
}

func FoldOptionT[T any, U any](optionT fp.Try[fp.Option[T]], zero U, f func(U, T) U) fp.Try[U] {
	return Map(optionT, func(insideValue fp.Option[T]) U {
		return option.Fold[T, U](insideValue, zero, f)
	})
}
