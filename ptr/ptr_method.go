package ptr

import (
	"fmt"

	"github.com/csgura/fp"
)

// Option[T] 의 method 로 있던 것
func All[T any](r fp.Ptr[T]) func(func(T) bool) {
	return func(f func(T) bool) {
		if IsDefined(r) {
			f(Get(r))
		}
	}
}

func Foreach[T any](r fp.Ptr[T], f func(v T)) {
	if IsDefined(r) {
		f(Get(r))
	}
}

func IsDefined[T any](r fp.Ptr[T]) bool {
	return r != nil
}

func IsEmpty[T any](r fp.Ptr[T]) bool {
	return r == nil
}

func Get[T any](r fp.Ptr[T]) T {
	if IsDefined(r) {
		return *r
	}
	panic("Ptr.empty")
}

func Unapply[T any](r fp.Ptr[T]) (T, bool) {
	if IsDefined(r) {
		return Get(r), true
	} else {
		var zero T
		return zero, false
	}
}

func String[T any](r fp.Ptr[T]) string {
	if IsDefined(r) {
		return fmt.Sprintf("Some(%v)", *r)
	} else {
		return "None"
	}
}

func Filter[T any](r fp.Ptr[T], p func(v T) bool) fp.Ptr[T] {
	if IsDefined(r) {
		if p(Get(r)) {
			return r
		}
	}
	return None[T]()

}

func FilterNot[T any](r fp.Ptr[T], p func(v T) bool) fp.Ptr[T] {
	if IsDefined(r) {
		if !p(Get(r)) {
			return r
		}
	}
	return None[T]()

}
func OrElse[T any](r fp.Ptr[T], t T) T {
	if IsDefined(r) {
		return Get(r)
	}
	return t
}

func OrZero[T any](r fp.Ptr[T]) T {
	return OrElseGet(r, fp.Zero[T])
}

func OrElseGet[T any](r fp.Ptr[T], f func() T) T {
	if IsDefined(r) {
		return Get(r)
	}
	return f()
}
func Or[T any](r fp.Ptr[T], f func() fp.Ptr[T]) fp.Ptr[T] {
	if IsDefined(r) {
		return r
	}
	return f()
}

func OrOption[T any](r fp.Ptr[T], v fp.Option[T]) fp.Ptr[T] {
	if IsDefined(r) {
		return r
	}
	return v.Ptr()
}

func OrPtr[T any](r fp.Ptr[T], v fp.Ptr[T]) fp.Ptr[T] {
	if IsDefined(r) {
		return r
	}
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func Recover[T any](r fp.Ptr[T], f func() T) fp.Ptr[T] {
	if IsDefined(r) {
		return r
	}
	t := f()
	return &t
}

func ToSlice[T any](r fp.Ptr[T]) fp.Slice[T] {
	if IsDefined(r) {
		return []T{Get(r)}
	}
	return nil
}

func Exists[T any](r fp.Ptr[T], p func(v T) bool) bool {
	return IsDefined(r) && p(*r)
}

func ForAll[T any](r fp.Ptr[T], p func(v T) bool) bool {
	return IsEmpty(r) || p(*r)
}
