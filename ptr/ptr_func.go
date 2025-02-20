package ptr

import (
	"github.com/csgura/fp"
)

func Pure[A any](a A) fp.Ptr[A] {
	return &a
}

func Some[T any](v T) fp.Ptr[T] {
	return &v
}

func None[T any]() fp.Ptr[T] {
	return nil
}

// 아규먼트를 무시하고 항상 None 을 리턴
func ConstNone[A, B any](a A) fp.Ptr[B] {
	return nil
}

func NonZero[T comparable](t T) fp.Ptr[T] {
	if t == fp.Zero[T]() {
		return None[T]()
	}
	return Some(t)
}

func FromTry[T any](t fp.Try[T]) fp.Ptr[T] {
	if t.IsSuccess() {
		return Some(t.Get())
	}
	return None[T]()
}

func Fold[A, B any](s fp.Ptr[A], zero B, f func(B, A) B) B {
	if s == nil {
		return zero
	}

	return f(zero, *s)
}

func ToSeq[T any](r fp.Ptr[T]) fp.Seq[T] {
	if r != nil {
		return fp.Seq[T]{*r}
	}
	return nil
}

func Iterator[T any](r fp.Ptr[T]) fp.Iterator[T] {
	return fp.IteratorOfSeq(ToSeq(r))
}
