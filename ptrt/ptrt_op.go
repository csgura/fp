// Code generated by monad_gen, DO NOT EDIT.
package ptrt

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/ptr"
	"github.com/csgura/fp/try"
)

func Pure[A any](a A) fp.PtrT[A] {
	return try.Success[fp.Ptr[A]](ptr.Pure[A](a))
}

func LiftT[A any](a fp.Try[A]) fp.PtrT[A] {
	return try.Map(a, ptr.Pure[A])
}

func Map[A any, B any](t fp.PtrT[A], f func(A) B) fp.PtrT[B] {
	return try.Map(t, func(ma fp.Ptr[A]) fp.Ptr[B] {
		return ptr.FlatMap[A, B](ma, func(a A) fp.Ptr[B] {
			return ptr.Pure[B](f(a))
		})
	})
}

func SubFlatMap[A any, B any](t fp.PtrT[A], f func(A) fp.Ptr[B]) fp.PtrT[B] {
	return try.Map(t, func(ma fp.Ptr[A]) fp.Ptr[B] {
		return ptr.FlatMap[A, B](ma, func(a A) fp.Ptr[B] {
			return f(a)
		})
	})
}

func MapT[A any, B any](t fp.PtrT[A], f func(A) fp.Try[B]) fp.PtrT[B] {
	sequencef := func(v fp.Ptr[fp.Try[B]]) fp.PtrT[B] {
		if ptr.IsDefined(v) {
			return try.Map(ptr.Get(v), ptr.Some)
		}
		return try.Success(ptr.None[B]())
	}
	return try.FlatMap(Map(t, f), sequencef)
}

func FlatMap[A any, B any](t fp.PtrT[A], f func(A) fp.PtrT[B]) fp.PtrT[B] {

	flatten := func(v fp.Ptr[fp.Ptr[B]]) fp.Ptr[B] {
		return ptr.FlatMap[fp.Ptr[B], B](v, fp.Id)
	}

	return try.Map(MapT(t, f), flatten)

}

func IsDefined[T any](ptrT fp.PtrT[T]) fp.Try[bool] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) bool {
		return ptr.IsDefined[T](insideValue)
	})
}

func IsEmpty[T any](ptrT fp.PtrT[T]) fp.Try[bool] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) bool {
		return ptr.IsEmpty[T](insideValue)
	})
}

func Filter[T any](ptrT fp.PtrT[T], p func(v T) bool) fp.Try[fp.Ptr[T]] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) fp.Ptr[T] {
		return ptr.Filter[T](insideValue, p)
	})
}

func OrElse[T any](ptrT fp.PtrT[T], t T) fp.Try[T] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) T {
		return ptr.OrElse[T](insideValue, t)
	})
}

func OrZero[T any](ptrT fp.PtrT[T]) fp.Try[T] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) T {
		return ptr.OrZero[T](insideValue)
	})
}

func OrElseGet[T any](ptrT fp.PtrT[T], f func() T) fp.Try[T] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) T {
		return ptr.OrElseGet[T](insideValue, f)
	})
}

func Or[T any](ptrT fp.PtrT[T], f func() fp.Ptr[T]) fp.Try[fp.Ptr[T]] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) fp.Ptr[T] {
		return ptr.Or[T](insideValue, f)
	})
}

func OrOption[T any](ptrT fp.PtrT[T], v fp.Option[T]) fp.Try[fp.Ptr[T]] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) fp.Ptr[T] {
		return ptr.OrOption[T](insideValue, v)
	})
}

func OrPtr[T any](ptrT fp.PtrT[T], v fp.Ptr[T]) fp.Try[fp.Ptr[T]] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) fp.Ptr[T] {
		return ptr.OrPtr[T](insideValue, v)
	})
}

func Recover[T any](ptrT fp.PtrT[T], f func() T) fp.Try[fp.Ptr[T]] {
	return try.Map(ptrT, func(insideValue fp.Ptr[T]) fp.Ptr[T] {
		return ptr.Recover[T](insideValue, f)
	})
}
