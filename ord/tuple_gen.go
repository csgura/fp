package ord

import (
	"github.com/csgura/fp"
)

func Tuple2[A1, A2 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2]) fp.Ord[fp.Tuple2[A1, A2]] {
	return New[fp.Tuple2[A1, A2]](
		fp.EqFunc[fp.Tuple2[A1, A2]](func(t1 fp.Tuple2[A1, A2], t2 fp.Tuple2[A1, A2]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple1(tins2).Eqv(t1.Tail(), t2.Tail())
		}),
		fp.LessFunc[fp.Tuple2[A1, A2]](func(t1 fp.Tuple2[A1, A2], t2 fp.Tuple2[A1, A2]) bool {
			if tins1.Less(t1.I1, t2.I1) {
				return true
			}
			if tins1.Less(t2.I1, t1.I1) {
				return false
			}
			return Tuple1(tins2).Less(t1.Tail(), t2.Tail())
		}),
	)
}
func Tuple3[A1, A2, A3 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3]) fp.Ord[fp.Tuple3[A1, A2, A3]] {
	return New[fp.Tuple3[A1, A2, A3]](
		fp.EqFunc[fp.Tuple3[A1, A2, A3]](func(t1 fp.Tuple3[A1, A2, A3], t2 fp.Tuple3[A1, A2, A3]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple2(tins2, tins3).Eqv(t1.Tail(), t2.Tail())
		}),
		fp.LessFunc[fp.Tuple3[A1, A2, A3]](func(t1 fp.Tuple3[A1, A2, A3], t2 fp.Tuple3[A1, A2, A3]) bool {
			if tins1.Less(t1.I1, t2.I1) {
				return true
			}
			if tins1.Less(t2.I1, t1.I1) {
				return false
			}
			return Tuple2(tins2, tins3).Less(t1.Tail(), t2.Tail())
		}),
	)
}
func Tuple4[A1, A2, A3, A4 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4]) fp.Ord[fp.Tuple4[A1, A2, A3, A4]] {
	return New[fp.Tuple4[A1, A2, A3, A4]](
		fp.EqFunc[fp.Tuple4[A1, A2, A3, A4]](func(t1 fp.Tuple4[A1, A2, A3, A4], t2 fp.Tuple4[A1, A2, A3, A4]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple3(tins2, tins3, tins4).Eqv(t1.Tail(), t2.Tail())
		}),
		fp.LessFunc[fp.Tuple4[A1, A2, A3, A4]](func(t1 fp.Tuple4[A1, A2, A3, A4], t2 fp.Tuple4[A1, A2, A3, A4]) bool {
			if tins1.Less(t1.I1, t2.I1) {
				return true
			}
			if tins1.Less(t2.I1, t1.I1) {
				return false
			}
			return Tuple3(tins2, tins3, tins4).Less(t1.Tail(), t2.Tail())
		}),
	)
}
func Tuple5[A1, A2, A3, A4, A5 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5]) fp.Ord[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return New[fp.Tuple5[A1, A2, A3, A4, A5]](
		fp.EqFunc[fp.Tuple5[A1, A2, A3, A4, A5]](func(t1 fp.Tuple5[A1, A2, A3, A4, A5], t2 fp.Tuple5[A1, A2, A3, A4, A5]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple4(tins2, tins3, tins4, tins5).Eqv(t1.Tail(), t2.Tail())
		}),
		fp.LessFunc[fp.Tuple5[A1, A2, A3, A4, A5]](func(t1 fp.Tuple5[A1, A2, A3, A4, A5], t2 fp.Tuple5[A1, A2, A3, A4, A5]) bool {
			if tins1.Less(t1.I1, t2.I1) {
				return true
			}
			if tins1.Less(t2.I1, t1.I1) {
				return false
			}
			return Tuple4(tins2, tins3, tins4, tins5).Less(t1.Tail(), t2.Tail())
		}),
	)
}
