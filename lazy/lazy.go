package lazy

import (
	"github.com/csgura/fp"
)

func Value[T any](v T) fp.Lazy[T] {
	return fp.LazyFunc(func() T {
		return v
	})
}

func Eval[T any](f func() T) fp.Lazy[T] {
	return fp.LazyFunc(f)
}

func Func0[T any](f fp.Func1[fp.Unit, T]) fp.Lazy[T] {
	return Eval(func() T {
		return f(fp.Unit{})
	})
}
