package hlist

func Lift2[A1, A2, R any](f func(a1 A1, a2 A2) R) func(Cons[A1, Cons[A2, Nil]]) R {
	return func(v Cons[A1, Cons[A2, Nil]]) R {
		rf := Lift1(func(a2 A2) R {
			return f(v.Head(), a2)
		})

		return rf(v.Tail())
	}
}
func Rift2[A1, A2, R any](f func(a1 A1, a2 A2) R) func(Cons[A2, Cons[A1, Nil]]) R {
	return func(v Cons[A2, Cons[A1, Nil]]) R {
		rf := Rift1(func(a1 A1) R {
			return f(a1, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R) func(Cons[A1, Cons[A2, Cons[A3, Nil]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Nil]]]) R {
		rf := Lift2(func(a2 A2, a3 A3) R {
			return f(v.Head(), a2, a3)
		})

		return rf(v.Tail())
	}
}
func Rift3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R) func(Cons[A3, Cons[A2, Cons[A1, Nil]]]) R {
	return func(v Cons[A3, Cons[A2, Cons[A1, Nil]]]) R {
		rf := Rift2(func(a1 A1, a2 A2) R {
			return f(a1, a2, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Nil]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Nil]]]]) R {
		rf := Lift3(func(a2 A2, a3 A3, a4 A4) R {
			return f(v.Head(), a2, a3, a4)
		})

		return rf(v.Tail())
	}
}
func Rift4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R) func(Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]) R {
	return func(v Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]) R {
		rf := Rift3(func(a1 A1, a2 A2, a3 A3) R {
			return f(a1, a2, a3, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Nil]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Nil]]]]]) R {
		rf := Lift4(func(a2 A2, a3 A3, a4 A4, a5 A5) R {
			return f(v.Head(), a2, a3, a4, a5)
		})

		return rf(v.Tail())
	}
}
func Rift5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) func(Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]) R {
	return func(v Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]) R {
		rf := Rift4(func(a1 A1, a2 A2, a3 A3, a4 A4) R {
			return f(a1, a2, a3, a4, v.Head())
		})

		return rf(v.Tail())
	}
}
