package future

import (
	"github.com/csgura/fp"
)

func Func1[A1, R any](f func(A1) (R, error), exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return func(a1 A1) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1)
		})
	}
}

func Unit1[A1 any](f func(A1) error, exec ...fp.Executor) fp.Func1[A1, fp.Future[fp.Unit]] {
	return func(a1 A1) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1)
			return fp.Unit{}, err
		})
	}
}

func Func2[A1, A2, R any](f func(A1, A2) (R, error), exec ...fp.Executor) fp.Func2[A1, A2, fp.Future[R]] {
	return func(a1 A1, a2 A2) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2)
		})
	}
}

func Unit2[A1, A2 any](f func(A1, A2) error, exec ...fp.Executor) fp.Func2[A1, A2, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2)
			return fp.Unit{}, err
		})
	}
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error), exec ...fp.Executor) fp.Func3[A1, A2, A3, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3)
		})
	}
}

func Unit3[A1, A2, A3 any](f func(A1, A2, A3) error, exec ...fp.Executor) fp.Func3[A1, A2, A3, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3)
			return fp.Unit{}, err
		})
	}
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error), exec ...fp.Executor) fp.Func4[A1, A2, A3, A4, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4)
		})
	}
}

func Unit4[A1, A2, A3, A4 any](f func(A1, A2, A3, A4) error, exec ...fp.Executor) fp.Func4[A1, A2, A3, A4, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4)
			return fp.Unit{}, err
		})
	}
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error), exec ...fp.Executor) fp.Func5[A1, A2, A3, A4, A5, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5)
		})
	}
}

func Unit5[A1, A2, A3, A4, A5 any](f func(A1, A2, A3, A4, A5) error, exec ...fp.Executor) fp.Func5[A1, A2, A3, A4, A5, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5)
			return fp.Unit{}, err
		})
	}
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error), exec ...fp.Executor) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6)
		})
	}
}

func Unit6[A1, A2, A3, A4, A5, A6 any](f func(A1, A2, A3, A4, A5, A6) error, exec ...fp.Executor) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6)
			return fp.Unit{}, err
		})
	}
}

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (R, error), exec ...fp.Executor) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func Unit7[A1, A2, A3, A4, A5, A6, A7 any](f func(A1, A2, A3, A4, A5, A6, A7) error, exec ...fp.Executor) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6, a7)
			return fp.Unit{}, err
		})
	}
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R, error), exec ...fp.Executor) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func Unit8[A1, A2, A3, A4, A5, A6, A7, A8 any](f func(A1, A2, A3, A4, A5, A6, A7, A8) error, exec ...fp.Executor) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6, a7, a8)
			return fp.Unit{}, err
		})
	}
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R, error), exec ...fp.Executor) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func Unit9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) error, exec ...fp.Executor) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			return fp.Unit{}, err
		})
	}
}

func Compose3[A1, A2, A3, R any](f1 fp.Func1[A1, fp.Future[A2]], f2 fp.Func1[A2, fp.Future[A3]], f3 fp.Func1[A3, fp.Future[R]], exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose2(f2, f3, exec...), exec...)
}

func Compose4[A1, A2, A3, A4, R any](f1 fp.Func1[A1, fp.Future[A2]], f2 fp.Func1[A2, fp.Future[A3]], f3 fp.Func1[A3, fp.Future[A4]], f4 fp.Func1[A4, fp.Future[R]], exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose3(f2, f3, f4, exec...), exec...)
}

func Compose5[A1, A2, A3, A4, A5, R any](f1 fp.Func1[A1, fp.Future[A2]], f2 fp.Func1[A2, fp.Future[A3]], f3 fp.Func1[A3, fp.Future[A4]], f4 fp.Func1[A4, fp.Future[A5]], f5 fp.Func1[A5, fp.Future[R]], exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose4(f2, f3, f4, f5, exec...), exec...)
}
