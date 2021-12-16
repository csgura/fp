package curried

import (
	"github.com/csgura/fp"
)

func Func2[A1, A2, R any](f func(A1, A2) R) fp.Func1[A1, fp.Func1[A2, R]] {
	return func(a1 A1) fp.Func1[A2, R] {
		return Func1(func(a2 A2) R {
			return f(a1, a2)
		})
	}
}
func Revert2[A1, A2, R any](f fp.Func1[A1, fp.Func1[A2, R]]) fp.Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return f(a1)(a2)
	}
}
func Func3[A1, A2, A3, R any](f func(A1, A2, A3) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, R]] {
		return Func2(func(a2 A2, a3 A3) R {
			return f(a1, a2, a3)
		})
	}
}
func Revert3[A1, A2, A3, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]) fp.Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return f(a1)(a2)(a3)
	}
}
func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]] {
		return Func3(func(a2 A2, a3 A3, a4 A4) R {
			return f(a1, a2, a3, a4)
		})
	}
}
func Revert4[A1, A2, A3, A4, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]) fp.Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return f(a1)(a2)(a3)(a4)
	}
}
func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]] {
		return Func4(func(a2 A2, a3 A3, a4 A4, a5 A5) R {
			return f(a1, a2, a3, a4, a5)
		})
	}
}
func Revert5[A1, A2, A3, A4, A5, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]) fp.Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return f(a1)(a2)(a3)(a4)(a5)
	}
}
func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]] {
		return Func5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
			return f(a1, a2, a3, a4, a5, a6)
		})
	}
}
func Revert6[A1, A2, A3, A4, A5, A6, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]) fp.Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)
	}
}
