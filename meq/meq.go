package meq

import (
	"bytes"
	"time"

	"github.com/csgura/fp"
)

type Eq[T any] = func(T, T) bool

var Time Eq[time.Time] = time.Time.Equal
var Bytes Eq[[]byte] = bytes.Equal
var String Eq[string] = Given[string]()
var PtrBytes Eq[*[]byte] = Ptr(Bytes)

func New[T any](f func(T, T) bool) Eq[T] {
	return f
}

func Ptr[T any](v Eq[T]) Eq[*T] {
	return func(a, b *T) bool {
		if a == nil && b == nil {
			return true
		}

		if a != nil && b != nil {
			return v(*a, *b)
		}
		return false
	}
}

func Given[T comparable]() Eq[T] {
	return func(t1, t2 T) bool {
		return t1 == t2
	}
}

func PtrGiven[T comparable]() Eq[*T] {
	return Ptr(Given[T]())
}

func UnPtr[T any](eqp Eq[*T]) Eq[T] {
	return func(t1, t2 T) bool {
		return eqp(&t1, &t2)
	}
}

func Slice[T any](eq Eq[T]) Eq[fp.Slice[T]] {
	return func(a, b fp.Slice[T]) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if !eq(a[i], b[i]) {
				return false
			}
		}
		return true
	}
}

func GoMap[K comparable, V any](eqV Eq[V]) Eq[map[K]V] {
	return func(a, b map[K]V) bool {
		if len(a) != len(b) {
			return false
		}

		for k, av := range a {
			bv, ok := b[k]
			if !ok {
				return false
			}

			if !eqV(av, bv) {
				return false
			}
		}
		return true
	}
}
