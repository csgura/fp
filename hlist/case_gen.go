package hlist

func Case2[A1, A2, T, R any](hl Cons[A1, Cons[A2, T]], f func(a1 A1, a2 A2) R) R {
	return Case1(hl.Tail(), func(a2 A2) R {
		return f(hl.Head(), a2)
	})
}
func Case3[A1, A2, A3, T, R any](hl Cons[A1, Cons[A2, Cons[A3, T]]], f func(a1 A1, a2 A2, a3 A3) R) R {
	return Case2(hl.Tail(), func(a2 A2, a3 A3) R {
		return f(hl.Head(), a2, a3)
	})
}
func Case4[A1, A2, A3, A4, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, T]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4) R) R {
	return Case3(hl.Tail(), func(a2 A2, a3 A3, a4 A4) R {
		return f(hl.Head(), a2, a3, a4)
	})
}
func Case5[A1, A2, A3, A4, A5, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, T]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) R {
	return Case4(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5) R {
		return f(hl.Head(), a2, a3, a4, a5)
	})
}
func Case6[A1, A2, A3, A4, A5, A6, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, T]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) R {
	return Case5(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return f(hl.Head(), a2, a3, a4, a5, a6)
	})
}
func Case7[A1, A2, A3, A4, A5, A6, A7, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, T]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) R {
	return Case6(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7)
	})
}
func Case8[A1, A2, A3, A4, A5, A6, A7, A8, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, T]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) R {
	return Case7(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8)
	})
}
func Case9[A1, A2, A3, A4, A5, A6, A7, A8, A9, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, T]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) R {
	return Case8(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9)
	})
}
func Case10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, T]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R) R {
	return Case9(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10)
	})
}
func Case11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, T]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R) R {
	return Case10(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	})
}
func Case12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, T]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R) R {
	return Case11(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	})
}
func Case13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, T]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R) R {
	return Case12(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	})
}
func Case14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, T]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R) R {
	return Case13(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	})
}
func Case15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, T]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R) R {
	return Case14(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	})
}
func Case16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, T]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R) R {
	return Case15(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	})
}
func Case17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, T]]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R) R {
	return Case16(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	})
}
func Case18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, T]]]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R) R {
	return Case17(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	})
}
func Case19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, T]]]]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R) R {
	return Case18(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	})
}
func Case20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, T]]]]]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R) R {
	return Case19(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	})
}
func Case21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, T]]]]]]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R) R {
	return Case20(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	})
}
func Case22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, T, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Cons[A22, T]]]]]]]]]]]]]]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R) R {
	return Case21(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	})
}
