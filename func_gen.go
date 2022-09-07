package fp

type Func3[A1, A2, A3, R any] func(a1 A1, a2 A2, a3 A3) R

type Func4[A1, A2, A3, A4, R any] func(a1 A1, a2 A2, a3 A3, a4 A4) R

type Func5[A1, A2, A3, A4, A5, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R

type Func6[A1, A2, A3, A4, A5, A6, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R

type Func7[A1, A2, A3, A4, A5, A6, A7, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R

type Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R

type Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R

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

func (r Func3[A1, A2, A3, R]) ApplyFirst(a1 A1) Func2[A2, A3, R] {
	return func(a2 A2, a3 A3) R {
		return r(a1, a2, a3)
	}
}

func (r Func3[A1, A2, A3, R]) ApplyFirst2(a1 A1, a2 A2) Func1[A3, R] {
	return func(a3 A3) R {
		return r(a1, a2, a3)
	}
}

func (r Func3[A1, A2, A3, R]) ApplyLast(a3 A3) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3)
	}
}

func (r Func3[A1, A2, A3, R]) ApplyLast2(a2 A2, a3 A3) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3)
	}
}

func (r Func3[A1, A2, A3, R]) ApplySecond(a2 A2) Func2[A1, A3, R] {
	return func(a1 A1, a3 A3) R {
		return r(a1, a2, a3)
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

func (r Func4[A1, A2, A3, A4, R]) ApplyFirst(a1 A1) Func3[A2, A3, A4, R] {
	return func(a2 A2, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplyFirst2(a1 A1, a2 A2) Func2[A3, A4, R] {
	return func(a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplyFirst3(a1 A1, a2 A2, a3 A3) Func1[A4, R] {
	return func(a4 A4) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplyLast(a4 A4) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplyLast2(a3 A3, a4 A4) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplyLast3(a2 A2, a3 A3, a4 A4) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplySecond(a2 A2) Func3[A1, A3, A4, R] {
	return func(a1 A1, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func4[A1, A2, A3, A4, R]) ApplyThird(a3 A3) Func3[A1, A2, A4, R] {
	return func(a1 A1, a2 A2, a4 A4) R {
		return r(a1, a2, a3, a4)
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

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyFirst(a1 A1) Func4[A2, A3, A4, A5, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyFirst2(a1 A1, a2 A2) Func3[A3, A4, A5, R] {
	return func(a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyFirst3(a1 A1, a2 A2, a3 A3) Func2[A4, A5, R] {
	return func(a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyFirst4(a1 A1, a2 A2, a3 A3, a4 A4) Func1[A5, R] {
	return func(a5 A5) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyLast(a5 A5) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyLast2(a4 A4, a5 A5) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyLast3(a3 A3, a4 A4, a5 A5) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyLast4(a2 A2, a3 A3, a4 A4, a5 A5) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplySecond(a2 A2) Func4[A1, A3, A4, A5, R] {
	return func(a1 A1, a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyThird(a3 A3) Func4[A1, A2, A4, A5, R] {
	return func(a1 A1, a2 A2, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) ApplyFourth(a4 A4) Func4[A1, A2, A3, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a5 A5) R {
		return r(a1, a2, a3, a4, a5)
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

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFirst(a1 A1) Func5[A2, A3, A4, A5, A6, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFirst2(a1 A1, a2 A2) Func4[A3, A4, A5, A6, R] {
	return func(a3 A3, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFirst3(a1 A1, a2 A2, a3 A3) Func3[A4, A5, A6, R] {
	return func(a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFirst4(a1 A1, a2 A2, a3 A3, a4 A4) Func2[A5, A6, R] {
	return func(a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFirst5(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Func1[A6, R] {
	return func(a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyLast(a6 A6) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyLast2(a5 A5, a6 A6) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyLast3(a4 A4, a5 A5, a6 A6) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyLast4(a3 A3, a4 A4, a5 A5, a6 A6) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyLast5(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplySecond(a2 A2) Func5[A1, A3, A4, A5, A6, R] {
	return func(a1 A1, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyThird(a3 A3) Func5[A1, A2, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFourth(a4 A4) Func5[A1, A2, A3, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) ApplyFifth(a5 A5) Func5[A1, A2, A3, A4, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6)
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

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFirst(a1 A1) Func6[A2, A3, A4, A5, A6, A7, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFirst2(a1 A1, a2 A2) Func5[A3, A4, A5, A6, A7, R] {
	return func(a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFirst3(a1 A1, a2 A2, a3 A3) Func4[A4, A5, A6, A7, R] {
	return func(a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFirst4(a1 A1, a2 A2, a3 A3, a4 A4) Func3[A5, A6, A7, R] {
	return func(a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFirst5(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Func2[A6, A7, R] {
	return func(a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFirst6(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Func1[A7, R] {
	return func(a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyLast(a7 A7) Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyLast2(a6 A6, a7 A7) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyLast3(a5 A5, a6 A6, a7 A7) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyLast4(a4 A4, a5 A5, a6 A6, a7 A7) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyLast5(a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyLast6(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplySecond(a2 A2) Func6[A1, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyThird(a3 A3) Func6[A1, A2, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFourth(a4 A4) Func6[A1, A2, A3, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplyFifth(a5 A5) Func6[A1, A2, A3, A4, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplySixth(a6 A6) Func6[A1, A2, A3, A4, A5, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
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

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst(a1 A1) Func7[A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst2(a1 A1, a2 A2) Func6[A3, A4, A5, A6, A7, A8, R] {
	return func(a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst3(a1 A1, a2 A2, a3 A3) Func5[A4, A5, A6, A7, A8, R] {
	return func(a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst4(a1 A1, a2 A2, a3 A3, a4 A4) Func4[A5, A6, A7, A8, R] {
	return func(a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst5(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Func3[A6, A7, A8, R] {
	return func(a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst6(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Func2[A7, A8, R] {
	return func(a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFirst7(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Func1[A8, R] {
	return func(a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast(a8 A8) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast2(a7 A7, a8 A8) Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast3(a6 A6, a7 A7, a8 A8) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast4(a5 A5, a6 A6, a7 A7, a8 A8) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast5(a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast6(a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyLast7(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplySecond(a2 A2) Func7[A1, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyThird(a3 A3) Func7[A1, A2, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFourth(a4 A4) Func7[A1, A2, A3, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplyFifth(a5 A5) Func7[A1, A2, A3, A4, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplySixth(a6 A6) Func7[A1, A2, A3, A4, A5, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplySeventh(a7 A7) Func7[A1, A2, A3, A4, A5, A6, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
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

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst(a1 A1) Func8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst2(a1 A1, a2 A2) Func7[A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst3(a1 A1, a2 A2, a3 A3) Func6[A4, A5, A6, A7, A8, A9, R] {
	return func(a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst4(a1 A1, a2 A2, a3 A3, a4 A4) Func5[A5, A6, A7, A8, A9, R] {
	return func(a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst5(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Func4[A6, A7, A8, A9, R] {
	return func(a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst6(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Func3[A7, A8, A9, R] {
	return func(a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst7(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Func2[A8, A9, R] {
	return func(a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFirst8(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Func1[A9, R] {
	return func(a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast(a9 A9) Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast2(a8 A8, a9 A9) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast3(a7 A7, a8 A8, a9 A9) Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast4(a6 A6, a7 A7, a8 A8, a9 A9) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast5(a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast6(a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast7(a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyLast8(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Func1[A1, R] {
	return func(a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplySecond(a2 A2) Func8[A1, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyThird(a3 A3) Func8[A1, A2, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFourth(a4 A4) Func8[A1, A2, A3, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyFifth(a5 A5) Func8[A1, A2, A3, A4, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplySixth(a6 A6) Func8[A1, A2, A3, A4, A5, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a7 A7, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplySeventh(a7 A7) Func8[A1, A2, A3, A4, A5, A6, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a8 A8, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplyEighth(a8 A8) Func8[A1, A2, A3, A4, A5, A6, A7, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a9 A9) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
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
