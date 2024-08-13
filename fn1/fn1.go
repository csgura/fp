package fn1

import "github.com/csgura/fp"

func Pure[X, A any](v A) fp.Func1[X, A] {
	return fp.Const[X](v)
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
