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
