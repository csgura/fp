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
func Lift6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Nil]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Nil]]]]]]) R {
		rf := Lift5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
			return f(v.Head(), a2, a3, a4, a5, a6)
		})

		return rf(v.Tail())
	}
}
func Rift6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) func(Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]) R {
	return func(v Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]) R {
		rf := Rift5(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
			return f(a1, a2, a3, a4, a5, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Nil]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Nil]]]]]]]) R {
		rf := Lift6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7)
		})

		return rf(v.Tail())
	}
}
func Rift7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) func(Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]) R {
	return func(v Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]) R {
		rf := Rift6(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
			return f(a1, a2, a3, a4, a5, a6, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Nil]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Nil]]]]]]]]) R {
		rf := Lift7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8)
		})

		return rf(v.Tail())
	}
}
func Rift8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) func(Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]) R {
	return func(v Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]) R {
		rf := Rift7(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
			return f(a1, a2, a3, a4, a5, a6, a7, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Nil]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Nil]]]]]]]]]) R {
		rf := Lift8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9)
		})

		return rf(v.Tail())
	}
}
func Rift9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) func(Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]) R {
	return func(v Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]) R {
		rf := Rift8(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, v.Head())
		})

		return rf(v.Tail())
	}
}
