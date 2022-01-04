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

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) R) fp.Func6[A1, A2, A3, A4, A5, A6, R] {
	return fp.Func6[A1, A2, A3, A4, A5, A6, R](f)
}

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) R) fp.Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return fp.Func7[A1, A2, A3, A4, A5, A6, A7, R](f)
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) R) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R](f)
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) R) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R](f)
}

func Curried2[A1, A2, R any](f func(A1, A2) R) fp.Func1[A1, fp.Func1[A2, R]] {
	return fp.Func2[A1, A2, R](f).Curried()
}

func Curried3[A1, A2, A3, R any](f func(A1, A2, A3) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]] {
	return fp.Func3[A1, A2, A3, R](f).Curried()
}

func Curried4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]] {
	return fp.Func4[A1, A2, A3, A4, R](f).Curried()
}

func Curried5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]] {
	return fp.Func5[A1, A2, A3, A4, A5, R](f).Curried()
}

func Curried6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]] {
	return fp.Func6[A1, A2, A3, A4, A5, A6, R](f).Curried()
}

func Curried7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]] {
	return fp.Func7[A1, A2, A3, A4, A5, A6, A7, R](f).Curried()
}

func Curried8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]] {
	return fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R](f).Curried()
}

func Curried9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]] {
	return fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R](f).Curried()
}

func UnTupled2[A1, A2, R any](f func(fp.Tuple2[A1, A2]) R) fp.Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return f(Tuple2(a1, a2))
	}
}

func UnTupled3[A1, A2, A3, R any](f func(fp.Tuple3[A1, A2, A3]) R) fp.Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return f(Tuple3(a1, a2, a3))
	}
}

func UnTupled4[A1, A2, A3, A4, R any](f func(fp.Tuple4[A1, A2, A3, A4]) R) fp.Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return f(Tuple4(a1, a2, a3, a4))
	}
}

func UnTupled5[A1, A2, A3, A4, A5, R any](f func(fp.Tuple5[A1, A2, A3, A4, A5]) R) fp.Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return f(Tuple5(a1, a2, a3, a4, a5))
	}
}

func UnTupled6[A1, A2, A3, A4, A5, A6, R any](f func(fp.Tuple6[A1, A2, A3, A4, A5, A6]) R) fp.Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return f(Tuple6(a1, a2, a3, a4, a5, a6))
	}
}

func UnTupled7[A1, A2, A3, A4, A5, A6, A7, R any](f func(fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) R) fp.Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return f(Tuple7(a1, a2, a3, a4, a5, a6, a7))
	}
}

func UnTupled8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) R) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return f(Tuple8(a1, a2, a3, a4, a5, a6, a7, a8))
	}
}

func UnTupled9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) R) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return f(Tuple9(a1, a2, a3, a4, a5, a6, a7, a8, a9))
	}
}
