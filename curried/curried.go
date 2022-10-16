//go:generate go run github.com/csgura/fp/internal/generator/curried_gen
package curried

import (
	"github.com/csgura/fp"
)

func Func1[A, R any](f func(A) R) fp.Func1[A, R] {
	return fp.Func1[A, R](f)
}

func Concat[A, B, R any](f fp.Func1[B, R]) fp.Func1[A, fp.Func1[B, R]] {
	return func(a A) fp.Func1[B, R] {
		return f
	}
}

func Flip[A, B, R any](f fp.Func1[A, fp.Func1[B, R]]) fp.Func1[B, fp.Func1[A, R]] {
	return func(b B) fp.Func1[A, R] {
		return func(a A) R {
			return f(a)(b)
		}
	}
}

func Compose2[A, B, GA, GR any](f fp.Func1[A, fp.Func1[B, GA]], g fp.Func1[GA, GR]) fp.Func1[A, fp.Func1[B, GR]] {

	return func(a A) fp.Func1[B, GR] {
		return func(b B) GR {
			return g(f(a)(b))
		}
	}
}
