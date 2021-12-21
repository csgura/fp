package lazy

import (
	"sync"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

type Val[T any] interface {
	Get() T
}

type ValFunc[T any] func() T

func (r ValFunc[T]) Get() T {
	return r()
}

func Eval[T any](f func() T) Val[T] {
	return Func0(as.Func0(f))
}

func Func0[T any](f fp.Func1[fp.Unit, T]) Val[T] {
	once := sync.Once{}
	var ret T
	return ValFunc[T](func() T {
		once.Do(func() {
			ret = f(fp.Unit{})
		})
		return ret
	})
}
