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
func Lift10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Nil]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Nil]]]]]]]]]]) R {
		rf := Lift9(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10)
		})

		return rf(v.Tail())
	}
}
func Rift10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R) func(Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]) R {
	return func(v Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]) R {
		rf := Rift9(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Nil]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Nil]]]]]]]]]]]) R {
		rf := Lift10(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		})

		return rf(v.Tail())
	}
}
func Rift11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R) func(Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]) R {
	return func(v Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]) R {
		rf := Rift10(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Nil]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Nil]]]]]]]]]]]]) R {
		rf := Lift11(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		})

		return rf(v.Tail())
	}
}
func Rift12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R) func(Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]) R {
	return func(v Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]) R {
		rf := Rift11(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Nil]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Nil]]]]]]]]]]]]]) R {
		rf := Lift12(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		})

		return rf(v.Tail())
	}
}
func Rift13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R) func(Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]) R {
	return func(v Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]) R {
		rf := Rift12(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Nil]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Nil]]]]]]]]]]]]]]) R {
		rf := Lift13(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		})

		return rf(v.Tail())
	}
}
func Rift14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R) func(Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]) R {
	return func(v Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]) R {
		rf := Rift13(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Nil]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Nil]]]]]]]]]]]]]]]) R {
		rf := Lift14(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		})

		return rf(v.Tail())
	}
}
func Rift15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R) func(Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]) R {
	return func(v Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]) R {
		rf := Rift14(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Nil]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Nil]]]]]]]]]]]]]]]]) R {
		rf := Lift15(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		})

		return rf(v.Tail())
	}
}
func Rift16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R) func(Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]) R {
		rf := Rift15(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Nil]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Nil]]]]]]]]]]]]]]]]]) R {
		rf := Lift16(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
		})

		return rf(v.Tail())
	}
}
func Rift17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R) func(Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]) R {
		rf := Rift16(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Nil]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Nil]]]]]]]]]]]]]]]]]]) R {
		rf := Lift17(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
		})

		return rf(v.Tail())
	}
}
func Rift18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R) func(Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]) R {
		rf := Rift17(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Nil]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Nil]]]]]]]]]]]]]]]]]]]) R {
		rf := Lift18(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
		})

		return rf(v.Tail())
	}
}
func Rift19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R) func(Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]) R {
		rf := Rift18(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Nil]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Nil]]]]]]]]]]]]]]]]]]]]) R {
		rf := Lift19(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
		})

		return rf(v.Tail())
	}
}
func Rift20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R) func(Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]) R {
		rf := Rift19(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Nil]]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Nil]]]]]]]]]]]]]]]]]]]]]) R {
		rf := Lift20(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
		})

		return rf(v.Tail())
	}
}
func Rift21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R) func(Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]) R {
		rf := Rift20(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, v.Head())
		})

		return rf(v.Tail())
	}
}
func Lift22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R) func(Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Cons[A22, Nil]]]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Cons[A22, Nil]]]]]]]]]]]]]]]]]]]]]]) R {
		rf := Lift21(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R {
			return f(v.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
		})

		return rf(v.Tail())
	}
}
func Rift22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R) func(Cons[A22, Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A22, Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]]) R {
		rf := Rift21(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, v.Head())
		})

		return rf(v.Tail())
	}
}
