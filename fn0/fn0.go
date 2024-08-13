package fn0

import "github.com/csgura/fp"

func Pure[A any](v A) fp.Func0[A] {
	return func(a1 fp.Unit) A {
		return v
	}
}

func Map[A, B any](m fp.Func0[A], fn fp.Func1[A, B]) fp.Func0[B] {
	return fp.Compose(m, fn)
}

func FlatMap[A, B any](m fp.Func0[A], fn fp.Func1[A, fp.Func0[B]]) fp.Func0[B] {
	return Flatten(Map(m, fn))
}

func Flatten[A any](m fp.Func0[fp.Func0[A]]) fp.Func0[A] {
	return func(u fp.Unit) A {
		return m(u)(u)
	}
}
