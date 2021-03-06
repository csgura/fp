package hlist

func Case2[A1, A2 any, T HList, R any](hl Cons[A1, Cons[A2, T]], f func(a1 A1, a2 A2) R) R {
	return Case1(hl.Tail(), func(a2 A2) R {
		return f(hl.Head(), a2)
	})
}
func Case3[A1, A2, A3 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, T]]], f func(a1 A1, a2 A2, a3 A3) R) R {
	return Case2(hl.Tail(), func(a2 A2, a3 A3) R {
		return f(hl.Head(), a2, a3)
	})
}
func Case4[A1, A2, A3, A4 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, T]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4) R) R {
	return Case3(hl.Tail(), func(a2 A2, a3 A3, a4 A4) R {
		return f(hl.Head(), a2, a3, a4)
	})
}
func Case5[A1, A2, A3, A4, A5 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, T]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) R {
	return Case4(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5) R {
		return f(hl.Head(), a2, a3, a4, a5)
	})
}
func Case6[A1, A2, A3, A4, A5, A6 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, T]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) R {
	return Case5(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return f(hl.Head(), a2, a3, a4, a5, a6)
	})
}
func Case7[A1, A2, A3, A4, A5, A6, A7 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, T]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) R {
	return Case6(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7)
	})
}
func Case8[A1, A2, A3, A4, A5, A6, A7, A8 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, T]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) R {
	return Case7(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8)
	})
}
func Case9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any, T HList, R any](hl Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Cons[A6, Cons[A7, Cons[A8, Cons[A9, T]]]]]]]]], f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) R {
	return Case8(hl.Tail(), func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return f(hl.Head(), a2, a3, a4, a5, a6, a7, a8, a9)
	})
}
