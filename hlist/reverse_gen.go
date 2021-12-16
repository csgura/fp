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
