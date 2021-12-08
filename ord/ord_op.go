package ord

import "github.com/csgura/fp"

func Tuple2[A, B any](a fp.Ord[A], b fp.Ord[B]) fp.Ord[fp.Tuple2[A, B]] {
	return fp.LessFunc[fp.Tuple2[A, B]](func(t1 fp.Tuple2[A, B], t2 fp.Tuple2[A, B]) bool {
		if a.Less(t1.I1, t2.I1) {
			return true
		}
		return b.Less(t1.I2, t2.I2)
	})
}

func Tuple3[A, B, C any](a fp.Ord[A], b fp.Ord[B], c fp.Ord[C]) fp.Ord[fp.Tuple3[A, B, C]] {
	return fp.LessFunc[fp.Tuple3[A, B, C]](func(t1 fp.Tuple3[A, B, C], t2 fp.Tuple3[A, B, C]) bool {
		if a.Less(t1.I1, t2.I1) {
			return true
		}
		return Tuple2(b, c).Less(t1.Tail(), t2.Tail())
	})
}
