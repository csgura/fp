package fp

type Func2[A1, A2, R any] func(a1 A1, a2 A2) R

type Func3[A1, A2, A3, R any] func(a1 A1, a2 A2, a3 A3) R

type Func4[A1, A2, A3, A4, R any] func(a1 A1, a2 A2, a3 A3, a4 A4) R

type Func5[A1, A2, A3, A4, A5, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R

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

func (r Func3[A1, A2, A3, R]) Shift() Func3[A2, A3, A1, R] {
	return func(a2 A2, a3 A3, a1 A1) R {
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

func (r Func4[A1, A2, A3, A4, R]) Shift() Func4[A2, A3, A4, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a1 A1) R {
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

func (r Func5[A1, A2, A3, A4, A5, R]) Shift() Func5[A2, A3, A4, A5, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a1 A1) R {
		return r(a1, a2, a3, a4, a5)
	}
}
