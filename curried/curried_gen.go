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
func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]] {
		return Func6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
			return f(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}
func Revert7[A1, A2, A3, A4, A5, A6, A7, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]) fp.Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)
	}
}
func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]] {
		return Func7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}
func Revert8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)
	}
}
func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) R) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]] {
		return Func8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}
func Revert9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)
	}
}

func Flip2[A1, A2, A3, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A1, R]]] {
	return Func3(
		func(a2 A2, a3 A3, a1 A1) R {
			return f(a1)(a2)(a3)
		},
	)
}

func Flip3[A1, A2, A3, A4, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A1, R]]]] {
	return Func4(
		func(a2 A2, a3 A3, a4 A4, a1 A1) R {
			return f(a1)(a2)(a3)(a4)
		},
	)
}

func Flip4[A1, A2, A3, A4, A5, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A1, R]]]]] {
	return Func5(
		func(a2 A2, a3 A3, a4 A4, a5 A5, a1 A1) R {
			return f(a1)(a2)(a3)(a4)(a5)
		},
	)
}

func Flip5[A1, A2, A3, A4, A5, A6, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A1, R]]]]]] {
	return Func6(
		func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a1 A1) R {
			return f(a1)(a2)(a3)(a4)(a5)(a6)
		},
	)
}

func Flip6[A1, A2, A3, A4, A5, A6, A7, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A1, R]]]]]]] {
	return Func7(
		func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a1 A1) R {
			return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)
		},
	)
}

func Flip7[A1, A2, A3, A4, A5, A6, A7, A8, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A1, R]]]]]]]] {
	return Func8(
		func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a1 A1) R {
			return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)
		},
	)
}

func Flip8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A1, R]]]]]]]]] {
	return Func9(
		func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a1 A1) R {
			return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)
		},
	)
}
