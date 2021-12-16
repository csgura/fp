package as

import (
	"github.com/csgura/fp"
)

func Func1[A1, R any](f func(A1) R) fp.Func1[A1, R] {
	return fp.Func1[A1, R](f)
}

func Func2[A1, A2, R any](f func(A1, A2) R) fp.Func2[A1, A2, R] {
	return fp.Func2[A1, A2, R](f)
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) R) fp.Func3[A1, A2, A3, R] {
	return fp.Func3[A1, A2, A3, R](f)
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) R) fp.Func4[A1, A2, A3, A4, R] {
	return fp.Func4[A1, A2, A3, A4, R](f)
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) R) fp.Func5[A1, A2, A3, A4, A5, R] {
	return fp.Func5[A1, A2, A3, A4, A5, R](f)
}
