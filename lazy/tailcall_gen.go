package lazy

func TailCall1[A1, R any](f func(A1) Eval[R], a1 A1) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1)
	})
}

func TailCall2[A1, A2, R any](f func(A1, A2) Eval[R], a1 A1, a2 A2) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2)
	})
}

func TailCall3[A1, A2, A3, R any](f func(A1, A2, A3) Eval[R], a1 A1, a2 A2, a3 A3) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3)
	})
}

func TailCall4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) Eval[R], a1 A1, a2 A2, a3 A3, a4 A4) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3, a4)
	})
}

func TailCall5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) Eval[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3, a4, a5)
	})
}

func TailCall6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) Eval[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3, a4, a5, a6)
	})
}

func TailCall7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) Eval[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3, a4, a5, a6, a7)
	})
}

func TailCall8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) Eval[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3, a4, a5, a6, a7, a8)
	})
}

func TailCall9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) Eval[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) Eval[R] {
	return TailCall(func() Eval[R] {
		return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	})
}
