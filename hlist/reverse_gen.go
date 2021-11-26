package hlist

func Reverse2[A1, A2 any](hl Cons[A1, Cons[A2, Nil]]) Cons[A2, Cons[A1, Nil]] {
	return Case2(hl, func(a1 A1, a2 A2) Cons[A2, Cons[A1, Nil]] {
		return Of2(a2, a1)
	})
}
func Reverse3[A1, A2, A3 any](hl Cons[A1, Cons[A2, Cons[A3, Nil]]]) Cons[A3, Cons[A2, Cons[A1, Nil]]] {
	return Case3(hl, func(a1 A1, a2 A2, a3 A3) Cons[A3, Cons[A2, Cons[A1, Nil]]] {
		return Of3(a3, a2, a1)
	})
}
func Reverse4[A1, A2, A3, A4 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Nil]]]]) Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]] {
	return Case4(hl, func(a1 A1, a2 A2, a3 A3, a4 A4) Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]] {
		return Of4(a4, a3, a2, a1)
	})
}
func Reverse5[A1, A2, A3, A4, A5 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Nil]]]]]) Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]] {
	return Case5(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]] {
		return Of5(a5, a4, a3, a2, a1)
	})
}
func Reverse6[A1, A2, A3, A4, A5, A6 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Nil]]]]]]) Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]] {
	return Case6(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]] {
		return Of6(a6, a5, a4, a3, a2, a1)
	})
}
func Reverse7[A1, A2, A3, A4, A5, A6, A7 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Nil]]]]]]]) Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]] {
	return Case7(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]] {
		return Of7(a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse8[A1, A2, A3, A4, A5, A6, A7, A8 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Nil]]]]]]]]) Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]] {
	return Case8(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]] {
		return Of8(a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Nil]]]]]]]]]) Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]] {
	return Case9(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]] {
		return Of9(a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Nil]]]]]]]]]]) Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]] {
	return Case10(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]] {
		return Of10(a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Nil]]]]]]]]]]]) Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]] {
	return Case11(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]] {
		return Of11(a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Nil]]]]]]]]]]]]) Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]] {
	return Case12(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]] {
		return Of12(a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Nil]]]]]]]]]]]]]) Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]] {
	return Case13(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]] {
		return Of13(a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Nil]]]]]]]]]]]]]]) Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]] {
	return Case14(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]] {
		return Of14(a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Nil]]]]]]]]]]]]]]]) Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]] {
	return Case15(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]] {
		return Of15(a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Nil]]]]]]]]]]]]]]]]) Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]] {
	return Case16(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]] {
		return Of16(a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Nil]]]]]]]]]]]]]]]]]) Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]] {
	return Case17(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]] {
		return Of17(a17, a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Nil]]]]]]]]]]]]]]]]]]) Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]] {
	return Case18(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]] {
		return Of18(a18, a17, a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Nil]]]]]]]]]]]]]]]]]]]) Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]] {
	return Case19(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]] {
		return Of19(a19, a18, a17, a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Nil]]]]]]]]]]]]]]]]]]]]) Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]] {
	return Case20(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]] {
		return Of20(a20, a19, a18, a17, a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Nil]]]]]]]]]]]]]]]]]]]]]) Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]] {
	return Case21(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]] {
		return Of21(a21, a20, a19, a18, a17, a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
func Reverse22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22 any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Cons[A22, Nil]]]]]]]]]]]]]]]]]]]]]]) Cons[A22, Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]] {
	return Case22(hl, func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) Cons[A22, Cons[A21, Cons[A20, Cons[A19, Cons[A18, Cons[A17, Cons[A16, Cons[A15, Cons[A14, Cons[A13, Cons[A12, Cons[A11, Cons[A10, Cons[A9, Cons[A8, Cons[A7, Cons[A6, Cons[A5, Cons[A4, Cons[A3, Cons[A2, Cons[A1, Nil]]]]]]]]]]]]]]]]]]]]]] {
		return Of22(a22, a21, a20, a19, a18, a17, a16, a15, a14, a13, a12, a11, a10, a9, a8, a7, a6, a5, a4, a3, a2, a1)
	})
}
