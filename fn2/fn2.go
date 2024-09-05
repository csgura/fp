package fn2

func Pa1[A, B, R any](f func(A, B) R, a A) func(B) R {
	return func(b B) R {
		return f(a, b)
	}
}

func Partial[A, B, R any](f func(A, B) R, a A) func(B) R {
	return Pa1(f, a)
}

func PaLast1[A, B, R any](f func(A, B) R, b B) func(A) R {
	return func(a A) R {
		return f(a, b)
	}
}
