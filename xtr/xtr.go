package xtr

// head([ 1,2,3 ] ) ->  1
func Head[T interface{ Head() V }, V any](t T) V {
	return t.Head()
}

// init([1,2,3]) -> [1,2]
func Init[T interface{ Init() V }, V any](t T) V {
	return t.Init()
}

// last([1,2,3]) ->  3
func Last[T interface{ Last() V }, V any](t T) V {
	return t.Last()
}

// tail([1,2,3]) -> [2,3]
func Tail[T interface{ Tail() V }, V any](t T) V {
	return t.Tail()
}
