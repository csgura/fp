package lazy

import (
	"github.com/csgura/fp"
)

type Val[T any] interface {
	Get() T
}

type ValFunc[T any] func() T

func (r ValFunc[T]) Get() T {
	return r()
}

func Eval[T any](f func() T) Val[T] {
	return fp.LazyFunc(f)
}

func Func0[T any](f fp.Func1[fp.Unit, T]) Val[T] {
	return Eval(func() T {
		return f(fp.Unit{})
	})
}
