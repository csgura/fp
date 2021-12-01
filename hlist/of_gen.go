package hlist

func Of2[A1, A2 any](a1 A1, a2 A2) Cons[A1, Cons[A2, Nil]] {
	return Concact(a1, Of1(a2))
}
func Of3[A1, A2, A3 any](a1 A1, a2 A2, a3 A3) Cons[A1, Cons[A2, Cons[A3, Nil]]] {
	return Concact(a1, Of2(a2, a3))
}
func Of4[A1, A2, A3, A4 any](a1 A1, a2 A2, a3 A3, a4 A4) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Nil]]]] {
	return Concact(a1, Of3(a2, a3, a4))
}
func Of5[A1, A2, A3, A4, A5 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Nil]]]]] {
	return Concact(a1, Of4(a2, a3, a4, a5))
}
func Of6[A1, A2, A3, A4, A5, A6 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Nil]]]]]] {
	return Concact(a1, Of5(a2, a3, a4, a5, a6))
}
func Of7[A1, A2, A3, A4, A5, A6, A7 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Nil]]]]]]] {
	return Concact(a1, Of6(a2, a3, a4, a5, a6, a7))
}
func Of8[A1, A2, A3, A4, A5, A6, A7, A8 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Nil]]]]]]]] {
	return Concact(a1, Of7(a2, a3, a4, a5, a6, a7, a8))
}
func Of9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Nil]]]]]]]]] {
	return Concact(a1, Of8(a2, a3, a4, a5, a6, a7, a8, a9))
}
func Of10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Nil]]]]]]]]]] {
	return Concact(a1, Of9(a2, a3, a4, a5, a6, a7, a8, a9, a10))
}
func Of11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Nil]]]]]]]]]]] {
	return Concact(a1, Of10(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11))
}
func Of12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Nil]]]]]]]]]]]] {
	return Concact(a1, Of11(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12))
}
func Of13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Nil]]]]]]]]]]]]] {
	return Concact(a1, Of12(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13))
}
func Of14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Nil]]]]]]]]]]]]]] {
	return Concact(a1, Of13(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14))
}
func Of15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Nil]]]]]]]]]]]]]]] {
	return Concact(a1, Of14(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15))
}
func Of16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Nil]]]]]]]]]]]]]]]] {
	return Concact(a1, Of15(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16))
}
func Of17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Nil]]]]]]]]]]]]]]]]] {
	return Concact(a1, Of16(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17))
}
func Of18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Nil]]]]]]]]]]]]]]]]]] {
	return Concact(a1, Of17(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18))
}
func Of19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Nil]]]]]]]]]]]]]]]]]]] {
	return Concact(a1, Of18(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19))
}
func Of20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Nil]]]]]]]]]]]]]]]]]]]] {
	return Concact(a1, Of19(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20))
}
func Of21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Nil]]]]]]]]]]]]]]]]]]]]] {
	return Concact(a1, Of20(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21))
}
func Of22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, Cons[A10, Cons[A11, Cons[A12, Cons[A13, Cons[A14, Cons[A15, Cons[A16, Cons[A17, Cons[A18, Cons[A19, Cons[A20, Cons[A21, Cons[A22, Nil]]]]]]]]]]]]]]]]]]]]]] {
	return Concact(a1, Of21(a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22))
}