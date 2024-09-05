package binop

// functions for binary operator : func(T,T) T

func Partial[T any](f func(T, T) T, t T) func(T) T {
	return Pa1(f, t)
}

func Pa1[T any](f func(T, T) T, t1 T) func(T) T {
	return func(t2 T) T {
		return f(t1, t2)
	}
}

func PaLast1[T any](f func(T, T) T, t2 T) func(T) T {
	return func(t1 T) T {
		return f(t1, t2)
	}
}
