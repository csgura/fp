package eq

import (
	"github.com/csgura/fp"
)

func Tuple2[A1, A2 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2]) fp.Eq[fp.Tuple2[A1, A2]] {
	return New(
		func(t1 fp.Tuple2[A1, A2], t2 fp.Tuple2[A1, A2]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple1(tins2).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple3[A1, A2, A3 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3]) fp.Eq[fp.Tuple3[A1, A2, A3]] {
	return New(
		func(t1 fp.Tuple3[A1, A2, A3], t2 fp.Tuple3[A1, A2, A3]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple2(tins2, tins3).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple4[A1, A2, A3, A4 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4]) fp.Eq[fp.Tuple4[A1, A2, A3, A4]] {
	return New(
		func(t1 fp.Tuple4[A1, A2, A3, A4], t2 fp.Tuple4[A1, A2, A3, A4]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple3(tins2, tins3, tins4).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple5[A1, A2, A3, A4, A5 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5]) fp.Eq[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return New(
		func(t1 fp.Tuple5[A1, A2, A3, A4, A5], t2 fp.Tuple5[A1, A2, A3, A4, A5]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple4(tins2, tins3, tins4, tins5).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple6[A1, A2, A3, A4, A5, A6 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6]) fp.Eq[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {
	return New(
		func(t1 fp.Tuple6[A1, A2, A3, A4, A5, A6], t2 fp.Tuple6[A1, A2, A3, A4, A5, A6]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple5(tins2, tins3, tins4, tins5, tins6).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
