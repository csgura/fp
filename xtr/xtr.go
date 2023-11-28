package xtr

func Head[T interface{ Head() V }, V any](t T) V {
	return t.Head()
}

func Last[T interface{ Last() V }, V any](t T) V {
	return t.Last()
}
