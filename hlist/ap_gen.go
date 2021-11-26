package hlist

func Ap2[A1, A2, R any](f func(a1 A1, a2 A2) R) func(Cons[A2, Cons[A1, Nil]]) R {
	return func(v Cons[A2, Cons[A1, Nil]]) R {
		rf := Ap1(func(a1 A1) R {
			return f(a1, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R) func(Cons[A3, Cons[A2, Cons[A1, Nil]]]) R {
	return func(v Cons[A3, Cons[A2, Cons[A1, Nil]]]) R {
		rf := Ap2(func(a1 A1, a2 A2) R {
			return f(a1, a2, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R) func(Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]) R {
	return func(v Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]) R {
		rf := Ap3(func(a1 A1, a2 A2, a3 A3) R {
			return f(a1, a2, a3, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) func(Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]) R {
	return func(v Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]) R {
		rf := Ap4(func(a1 A1, a2 A2, a3 A3, a4 A4) R {
			return f(a1, a2, a3, a4, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) func(Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]) R {
	return func(v Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]) R {
		rf := Ap5(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
			return f(a1, a2, a3, a4, a5, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) func(Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]) R {
	return func(v Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]) R {
		rf := Ap6(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
			return f(a1, a2, a3, a4, a5, a6, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) func(Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]) R {
	return func(v Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]) R {
		rf := Ap7(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
			return f(a1, a2, a3, a4, a5, a6, a7, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) func(Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]) R {
	return func(v Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]) R {
		rf := Ap8(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R) func(Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]) R {
	return func(v Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]) R {
		rf := Ap9(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R) func(Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]) R {
	return func(v Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]) R {
		rf := Ap10(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R) func(Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]) R {
	return func(v Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]) R {
		rf := Ap11(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R) func(Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]) R {
	return func(v Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]) R {
		rf := Ap12(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R) func(Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]) R {
	return func(v Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]) R {
		rf := Ap13(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R) func(Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]) R {
	return func(v Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]) R {
		rf := Ap14(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R) func(Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]) R {
		rf := Ap15(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R) func(Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]) R {
		rf := Ap16(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R) func(Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]) R {
		rf := Ap17(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R) func(Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]) R {
		rf := Ap18(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R) func(Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]) R {
		rf := Ap19(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R) func(Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]) R {
		rf := Ap20(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, v.Head())
		})

		return rf(v.Tail())
	}
}
func Ap22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R) func(Cons[A22, Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]]) R {
	return func(v Cons[A22, Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]]) R {
		rf := Ap21(func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, v.Head())
		})

		return rf(v.Tail())
	}
}
