// Code generated by template_gen, DO NOT EDIT.
package try

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

func Curried2[A1, A2, R any](f func(A1, A2) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Try[R]]] {
	return as.Curried2(func(a1 A1, a2 A2) fp.Try[R] {
		return Apply(f(a1, a2))
	})
}

func CurriedPtr2[A1, A2, R any](f func(A1, A2) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Try[R]]] {
	return as.Curried2(func(a1 A1, a2 A2) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2)), FromPtr)
	})
}

func Curried3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Try[R]]]] {
	return as.Curried3(func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		return Apply(f(a1, a2, a3))
	})
}

func CurriedPtr3[A1, A2, A3, R any](f func(A1, A2, A3) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Try[R]]]] {
	return as.Curried3(func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3)), FromPtr)
	})
}

func Curried4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Try[R]]]]] {
	return as.Curried4(func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		return Apply(f(a1, a2, a3, a4))
	})
}

func CurriedPtr4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Try[R]]]]] {
	return as.Curried4(func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3, a4)), FromPtr)
	})
}

func Curried5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Try[R]]]]]] {
	return as.Curried5(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		return Apply(f(a1, a2, a3, a4, a5))
	})
}

func CurriedPtr5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Try[R]]]]]] {
	return as.Curried5(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3, a4, a5)), FromPtr)
	})
}

func Curried6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Try[R]]]]]]] {
	return as.Curried6(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		return Apply(f(a1, a2, a3, a4, a5, a6))
	})
}

func CurriedPtr6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Try[R]]]]]]] {
	return as.Curried6(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3, a4, a5, a6)), FromPtr)
	})
}

func Curried7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Try[R]]]]]]]] {
	return as.Curried7(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		return Apply(f(a1, a2, a3, a4, a5, a6, a7))
	})
}

func CurriedPtr7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Try[R]]]]]]]] {
	return as.Curried7(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3, a4, a5, a6, a7)), FromPtr)
	})
}

func Curried8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Try[R]]]]]]]]] {
	return as.Curried8(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		return Apply(f(a1, a2, a3, a4, a5, a6, a7, a8))
	})
}

func CurriedPtr8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Try[R]]]]]]]]] {
	return as.Curried8(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3, a4, a5, a6, a7, a8)), FromPtr)
	})
}

func Curried9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Try[R]]]]]]]]]] {
	return as.Curried9(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		return Apply(f(a1, a2, a3, a4, a5, a6, a7, a8, a9))
	})
}

func CurriedPtr9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (*R, error)) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Try[R]]]]]]]]]] {
	return as.Curried9(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		return FlatMap(Apply(f(a1, a2, a3, a4, a5, a6, a7, a8, a9)), FromPtr)
	})
}
