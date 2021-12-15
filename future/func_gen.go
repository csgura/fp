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

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (R, error), exec ...fp.ExecContext) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R, error), exec ...fp.ExecContext) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R, error), exec ...fp.ExecContext) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R, error), exec ...fp.ExecContext) fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		})
	}
}

func Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R, error), exec ...fp.ExecContext) fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		})
	}
}

func Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R, error), exec ...fp.ExecContext) fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		})
	}
}

func Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R, error), exec ...fp.ExecContext) fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		})
	}
}

func Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R, error), exec ...fp.ExecContext) fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		})
	}
}

func Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R, error), exec ...fp.ExecContext) fp.Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		})
	}
}

func Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R, error), exec ...fp.ExecContext) fp.Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		})
	}
}

func Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R, error), exec ...fp.ExecContext) fp.Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
		})
	}
}

func Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R, error), exec ...fp.ExecContext) fp.Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
		})
	}
}

func Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R, error), exec ...fp.ExecContext) fp.Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
		})
	}
}

func Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R, error), exec ...fp.ExecContext) fp.Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
		})
	}
}

func Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R, error), exec ...fp.ExecContext) fp.Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
		})
	}
}

func Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R, error), exec ...fp.ExecContext) fp.Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
		})
	}
}
