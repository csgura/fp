package hlist

func Of2[A1, A2 any](a1 A1, a2 A2) Cons[A1, Cons[A2, Nil]] {
	return Concat(a1, Of1(a2))
}
func Of3[A1, A2, A3 any](a1 A1, a2 A2, a3 A3) Cons[A1, Cons[A2, Cons[A3, Nil]]] {
	return Concat(a1, Of2(a2, a3))
}
func Of4[A1, A2, A3, A4 any](a1 A1, a2 A2, a3 A3, a4 A4) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Nil]]]] {
	return Concat(a1, Of3(a2, a3, a4))
}
func Of5[A1, A2, A3, A4, A5 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Cons[A1, Cons[A2, Cons[A3, Cons[A4, Cons[A5, Nil]]]]] {
	return Concat(a1, Of4(a2, a3, a4, a5))
}
