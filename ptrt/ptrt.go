package ptrt

import (
	"iter"
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/ptr"
	"github.com/csgura/fp/try"
)

func Some[A any](v A) fp.PtrT[A] {
	return Pure(v)
}

func None[A any]() fp.PtrT[A] {
	return try.Success(ptr.None[A]())
}

func Failure[T any](err error) fp.PtrT[T] {
	return try.Failure[fp.Ptr[T]](err)
}

func isNil(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func Of[T any](v T) fp.PtrT[T] {
	var i any = v
	if i == nil {
		return None[T]()
	}

	rv := reflect.ValueOf(i)
	if isNil(rv) {
		return None[T]()
	}
	return Some(v)
}

func NonZero[T comparable](t T) fp.PtrT[T] {
	if t == fp.Zero[T]() {
		return None[T]()
	}
	return Some(t)
}

func NonEmptySlice[T ~[]E, E any](t T) fp.PtrT[T] {
	if len(t) == 0 {
		return None[T]()
	}
	return Some(t)
}

func FromTry[A any](v fp.Try[A]) fp.PtrT[A] {
	return try.Map(v, ptr.Some)
}

func Fold[T any, U any](ptrT fp.PtrT[T], zero U, f func(U, T) U) U {
	if ptrT.IsFailure() {
		return zero
	}
	return ptr.Fold(ptrT.Get(), zero, f)
}

func Iterator[T any](optionT fp.PtrT[T]) fp.Iterator[T] {
	return iterator.FlatMap(try.Iterator(optionT), ptr.Iterator)
}

func All[T any](optionT fp.PtrT[T]) iter.Seq[T] {
	return Iterator(optionT).All()
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.PtrT[T]] {
	return genfp.GenerateMonadTransformer[fp.PtrT[T]]{
		File:     "ptrt_op.go",
		TypeParm: genfp.TypeOf[T](),
		GivenMonad: genfp.MonadFunctions{
			Pure: try.Success[T],
		},
		ExposureMonad: genfp.MonadFunctions{
			Pure:    ptr.Pure[T],
			FlatMap: ptr.FlatMap[T, U],
		},
		Sequence: func(v fp.Ptr[fp.Try[T]]) fp.PtrT[T] {
			if ptr.IsDefined(v) {
				return try.Map(ptr.Get(v), ptr.Some)
			}
			return try.Success(ptr.None[T]())
		},
		Transform: []any{
			ptr.IsDefined[T],
			ptr.IsEmpty[T],
			ptr.Filter[T],
			ptr.OrElse[T],
			ptr.OrZero[T],
			ptr.OrElseGet[T],
			ptr.Or[T],
			ptr.OrOption[T],
			ptr.OrPtr[T],
			ptr.Recover[T],
		},
	}
}

func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.PtrT[B]) fp.PtrT[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsSuccess() && ptr.IsDefined(t.Get()) {
			sum = ptr.Get(t.Get())
		} else {
			return t
		}
	}
	return Pure(sum)
}

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.PtrT[A]] {
	return genfp.GenerateMonadFunctions[fp.PtrT[A]]{
		File:     "ptrt_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateTraverseFunctions[fp.PtrT[A]] {
	return genfp.GenerateTraverseFunctions[fp.PtrT[A]]{
		File:     "ptrt_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}
