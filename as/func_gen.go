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

func Curried1[A1, R any](f func(A1) R) fp.Func1[A1, R] {
	return fp.Func1[A1, R](f).Curried()
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
