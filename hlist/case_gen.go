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
