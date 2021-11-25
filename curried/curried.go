package curried

import (
	"github.com/csgura/fp"
)

func Func1[A, R any](f func(A) R) fp.Func1[A, R] {
	return fp.Func1[A, R](f)
}

func Func2[A, B, R any](f func(A, B) R) fp.Func1[A, fp.Func1[B, R]] {
	return func(a A) fp.Func1[B, R] {
		return Func1(func(b B) R {
			return f(a, b)
		})
	}
}

func Func3[A, B, C, R any](f func(A, B, C) R) fp.Func1[A, fp.Func1[B, fp.Func1[C, R]]] {
	return func(a A) fp.Func1[B, fp.Func1[C, R]] {
		return Func2(func(b B, c C) R {
			return f(a, b, c)
		})
	}
}
