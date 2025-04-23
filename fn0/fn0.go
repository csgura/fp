package fn0

import (
	"sync"

	"github.com/csgura/fp"
)

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

func Memoize[A any](f fp.Func0[A]) fp.Func0[A] {
	once := sync.Once{}
	var ret A
	return func(u fp.Unit) A {
		once.Do(func() {
			ret = f(u)
		})
		return ret
	}
}

func Get[A any](f fp.Func0[A]) A {
	return f(fp.Unit{})
}

func IntoSupplier[A any](f fp.Func0[A]) fp.Supplier[A] {
	return func() A {
		return f(fp.Unit{})
	}
}

func FromSupplier[A any](f fp.Supplier[A]) fp.Func0[A] {
	return func(fp.Unit) A {
		return f()
	}
}
