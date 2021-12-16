package future

import (
	"github.com/csgura/fp"
)

func Func1[A1, R any](f func(A1) (R, error), exec ...fp.ExecContext) fp.Func1[A1, fp.Future[R]] {
	return func(a1 A1) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1)
		})
	}
}

func Func2[A1, A2, R any](f func(A1, A2) (R, error), exec ...fp.ExecContext) fp.Func2[A1, A2, fp.Future[R]] {
	return func(a1 A1, a2 A2) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2)
		})
	}
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error), exec ...fp.ExecContext) fp.Func3[A1, A2, A3, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3)
		})
	}
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error), exec ...fp.ExecContext) fp.Func4[A1, A2, A3, A4, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4)
		})
	}
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error), exec ...fp.ExecContext) fp.Func5[A1, A2, A3, A4, A5, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5)
		})
	}
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error), exec ...fp.ExecContext) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6)
		})
	}
}
