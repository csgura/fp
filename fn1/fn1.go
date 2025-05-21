package fn1

import (
	"sync"

	"github.com/csgura/fp"
)

func Pure[X, A any](v A) fp.Func1[X, A] {
	return fp.Const[X](v)
}

func Pipe[A, R any](a A, f func(A) R) R {
	return f(a)
}

func Apply[A, R any](f fp.Func1[A, R], a A) R {
	return f(a)
}

func Map[X, A, B any](m fp.Func1[X, A], fn fp.Func1[A, B]) fp.Func1[X, B] {
	return fp.Compose(m, fn)
}

func FlatMap[X, A, B any](m fp.Func1[X, A], fn fp.Func1[A, fp.Func1[X, B]]) fp.Func1[X, B] {
	return Flatten(Map(m, fn))
}

func Flatten[X, A any](m fp.Func1[X, fp.Func1[X, A]]) fp.Func1[X, A] {
	return func(u X) A {
		return m(u)(u)
	}
}

func Get[X any]() fp.Func1[X, X] {
	return func(x X) X {
		return x
	}
}

func WithArg[X, A any](fn fp.Func1[X, fp.Func1[X, A]]) fp.Func1[X, A] {
	return FlatMap(Get[X](), fn)
}

func Memoize[A, B any](f func(A) B) fp.Func1[A, B] {
	once := sync.Once{}
	var ret B
	return func(a A) B {
		once.Do(func() {
			ret = f(a)
		})
		return ret
	}
}

func Pa1[A, R any](f fp.Func1[A, R], a A) func() R {
	return func() R {
		return f(a)
	}
}
