package try

import (
	"github.com/csgura/fp"
)

func Func1[A1, R any](f func(A1) (R, error)) fp.Func1[A1, fp.Try[R]] {
	return func(a1 A1) fp.Try[R] {
		ret, err := f(a1)
		return Apply(ret, err)
	}
}

func Func2[A1, A2, R any](f func(A1, A2) (R, error)) fp.Func2[A1, A2, fp.Try[R]] {
	return func(a1 A1, a2 A2) fp.Try[R] {
		ret, err := f(a1, a2)
		return Apply(ret, err)
	}
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error)) fp.Func3[A1, A2, A3, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		ret, err := f(a1, a2, a3)
		return Apply(ret, err)
	}
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error)) fp.Func4[A1, A2, A3, A4, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4)
		return Apply(ret, err)
	}
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error)) fp.Func5[A1, A2, A3, A4, A5, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5)
		return Apply(ret, err)
	}
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error)) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6)
		return Apply(ret, err)
	}
}
