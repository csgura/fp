package monoid

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/product"
)

func Tuple2[A1, A2 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2]) fp.Monoid[fp.Tuple2[A1, A2]] {
	return New(
		func() fp.Tuple2[A1, A2] {
			return product.Tuple2(tins1.Empty(), tins2.Empty())
		},
		func(t1 fp.Tuple2[A1, A2], t2 fp.Tuple2[A1, A2]) fp.Tuple2[A1, A2] {
			return product.Tuple2(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2))
		},
	)
}
func Tuple3[A1, A2, A3 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3]) fp.Monoid[fp.Tuple3[A1, A2, A3]] {
	return New(
		func() fp.Tuple3[A1, A2, A3] {
			return product.Tuple3(tins1.Empty(), tins2.Empty(), tins3.Empty())
		},
		func(t1 fp.Tuple3[A1, A2, A3], t2 fp.Tuple3[A1, A2, A3]) fp.Tuple3[A1, A2, A3] {
			return product.Tuple3(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3))
		},
	)
}
func Tuple4[A1, A2, A3, A4 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4]) fp.Monoid[fp.Tuple4[A1, A2, A3, A4]] {
	return New(
		func() fp.Tuple4[A1, A2, A3, A4] {
			return product.Tuple4(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty())
		},
		func(t1 fp.Tuple4[A1, A2, A3, A4], t2 fp.Tuple4[A1, A2, A3, A4]) fp.Tuple4[A1, A2, A3, A4] {
			return product.Tuple4(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4))
		},
	)
}
func Tuple5[A1, A2, A3, A4, A5 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5]) fp.Monoid[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return New(
		func() fp.Tuple5[A1, A2, A3, A4, A5] {
			return product.Tuple5(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty())
		},
		func(t1 fp.Tuple5[A1, A2, A3, A4, A5], t2 fp.Tuple5[A1, A2, A3, A4, A5]) fp.Tuple5[A1, A2, A3, A4, A5] {
			return product.Tuple5(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5))
		},
	)
}
