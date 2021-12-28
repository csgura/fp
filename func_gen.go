package fp

type Func2[A1, A2, R any] func(a1 A1, a2 A2) R

type Func3[A1, A2, A3, R any] func(a1 A1, a2 A2, a3 A3) R

type Func4[A1, A2, A3, A4, R any] func(a1 A1, a2 A2, a3 A3, a4 A4) R

type Func5[A1, A2, A3, A4, A5, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R

type Func6[A1, A2, A3, A4, A5, A6, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R

type Func7[A1, A2, A3, A4, A5, A6, A7, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R

type Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R

type Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R

func (r Func2[A1, A2, R]) Tupled() Func1[Tuple2[A1, A2], R] {
	return func(t Tuple2[A1, A2]) R {
		return r(t.Unapply())
	}
}

func (r Func2[A1, A2, R]) Curried() Func1[A1, Func1[A2, R]] {
	return func(a1 A1) Func1[A2, R] {
		return Func1[A2, R](func(a2 A2) R {
			return r(a1, a2)
		}).Curried()
	}
}

func (r Func2[A1, A2, R]) Shift() Func2[A2, A1, R] {
	return func(a2 A2, a1 A1) R {
		return r(a1, a2)
	}
}

func (r Func3[A1, A2, A3, R]) Tupled() Func1[Tuple3[A1, A2, A3], R] {
	return func(t Tuple3[A1, A2, A3]) R {
		return r(t.Unapply())
	}
}

func (r Func3[A1, A2, A3, R]) Curried() Func1[A1, Func1[A2, Func1[A3, R]]] {
	return func(a1 A1) Func1[A2, Func1[A3, R]] {
		return Func2[A2, A3, R](func(a2 A2, a3 A3) R {
			return r(a1, a2, a3)
		}).Curried()
	}
}

func (r Func4[A1, A2, A3, A4, R]) Tupled() Func1[Tuple4[A1, A2, A3, A4], R] {
	return func(t Tuple4[A1, A2, A3, A4]) R {
		return r(t.Unapply())
	}
}

func (r Func4[A1, A2, A3, A4, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, R]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, R]]] {
		return Func3[A2, A3, A4, R](func(a2 A2, a3 A3, a4 A4) R {
			return r(a1, a2, a3, a4)
		}).Curried()
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) Tupled() Func1[Tuple5[A1, A2, A3, A4, A5], R] {
	return func(t Tuple5[A1, A2, A3, A4, A5]) R {
		return r(t.Unapply())
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, R]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, R]]]] {
		return Func4[A2, A3, A4, A5, R](func(a2 A2, a3 A3, a4 A4, a5 A5) R {
			return r(a1, a2, a3, a4, a5)
		}).Curried()
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) Tupled() Func1[Tuple6[A1, A2, A3, A4, A5, A6], R] {
	return func(t Tuple6[A1, A2, A3, A4, A5, A6]) R {
		return r(t.Unapply())
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, R]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, R]]]]] {
		return Func5[A2, A3, A4, A5, A6, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
			return r(a1, a2, a3, a4, a5, a6)
		}).Curried()
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) Tupled() Func1[Tuple7[A1, A2, A3, A4, A5, A6, A7], R] {
	return func(t Tuple7[A1, A2, A3, A4, A5, A6, A7]) R {
		return r(t.Unapply())
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, R]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, R]]]]]] {
		return Func6[A2, A3, A4, A5, A6, A7, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
			return r(a1, a2, a3, a4, a5, a6, a7)
		}).Curried()
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Tupled() Func1[Tuple8[A1, A2, A3, A4, A5, A6, A7, A8], R] {
	return func(t Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) R {
		return r(t.Unapply())
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, R]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, R]]]]]]] {
		return Func7[A2, A3, A4, A5, A6, A7, A8, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8)
		}).Curried()
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Tupled() Func1[Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9], R] {
	return func(t Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) R {
		return r(t.Unapply())
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, R]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, R]]]]]]]] {
		return Func8[A2, A3, A4, A5, A6, A7, A8, A9, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}).Curried()
	}
}

func Compose3[A1, A2, A3, R any](f1 Func1[A1, A2], f2 Func1[A2, A3], f3 Func1[A3, R]) Func1[A1, R] {
	return Compose2(f1, Compose2(f2, f3))
}

func Compose4[A1, A2, A3, A4, R any](f1 Func1[A1, A2], f2 Func1[A2, A3], f3 Func1[A3, A4], f4 Func1[A4, R]) Func1[A1, R] {
	return Compose2(f1, Compose3(f2, f3, f4))
}

func Compose5[A1, A2, A3, A4, A5, R any](f1 Func1[A1, A2], f2 Func1[A2, A3], f3 Func1[A3, A4], f4 Func1[A4, A5], f5 Func1[A5, R]) Func1[A1, R] {
	return Compose2(f1, Compose4(f2, f3, f4, f5))
}

func Nop1[A1, A2, R any](f func(A2) R) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return f(a2)
	}
}

func Nop2[A1, A2, A3, R any](f func(A3) R) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return f(a3)
	}
}

func Nop3[A1, A2, A3, A4, R any](f func(A4) R) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return f(a4)
	}
}

func Nop4[A1, A2, A3, A4, A5, R any](f func(A5) R) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return f(a5)
	}
}

func Nop5[A1, A2, A3, A4, A5, A6, R any](f func(A6) R) Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return f(a6)
	}
}

func Nop6[A1, A2, A3, A4, A5, A6, A7, R any](f func(A7) R) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return f(a7)
	}
}

func Nop7[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A8) R) Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return f(a8)
	}
}

func Nop8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A9) R) Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return f(a9)
	}
}
