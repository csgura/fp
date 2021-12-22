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
func Tuple6[A1, A2, A3, A4, A5, A6 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6]) fp.Monoid[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {
	return New(
		func() fp.Tuple6[A1, A2, A3, A4, A5, A6] {
			return product.Tuple6(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty())
		},
		func(t1 fp.Tuple6[A1, A2, A3, A4, A5, A6], t2 fp.Tuple6[A1, A2, A3, A4, A5, A6]) fp.Tuple6[A1, A2, A3, A4, A5, A6] {
			return product.Tuple6(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6))
		},
	)
}
func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7]) fp.Monoid[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {
	return New(
		func() fp.Tuple7[A1, A2, A3, A4, A5, A6, A7] {
			return product.Tuple7(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty())
		},
		func(t1 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7], t2 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) fp.Tuple7[A1, A2, A3, A4, A5, A6, A7] {
			return product.Tuple7(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7))
		},
	)
}
func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8]) fp.Monoid[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return New(
		func() fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8] {
			return product.Tuple8(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty())
		},
		func(t1 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8], t2 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8] {
			return product.Tuple8(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8))
		},
	)
}
func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9]) fp.Monoid[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return New(
		func() fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
			return product.Tuple9(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty())
		},
		func(t1 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9], t2 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
			return product.Tuple9(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9))
		},
	)
}
func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10]) fp.Monoid[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return New(
		func() fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
			return product.Tuple10(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty())
		},
		func(t1 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], t2 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
			return product.Tuple10(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10))
		},
	)
}
func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11]) fp.Monoid[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return New(
		func() fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
			return product.Tuple11(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty())
		},
		func(t1 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], t2 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
			return product.Tuple11(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11))
		},
	)
}
func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12]) fp.Monoid[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return New(
		func() fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
			return product.Tuple12(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty())
		},
		func(t1 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], t2 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
			return product.Tuple12(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12))
		},
	)
}
func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13]) fp.Monoid[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return New(
		func() fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
			return product.Tuple13(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty())
		},
		func(t1 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], t2 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
			return product.Tuple13(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13))
		},
	)
}
func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14]) fp.Monoid[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return New(
		func() fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
			return product.Tuple14(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty())
		},
		func(t1 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], t2 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
			return product.Tuple14(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14))
		},
	)
}
func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15]) fp.Monoid[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return New(
		func() fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
			return product.Tuple15(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty())
		},
		func(t1 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], t2 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
			return product.Tuple15(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15))
		},
	)
}
func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15], tins16 fp.Monoid[A16]) fp.Monoid[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return New(
		func() fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
			return product.Tuple16(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty(), tins16.Empty())
		},
		func(t1 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], t2 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
			return product.Tuple16(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15), tins16.Combine(t1.I16, t2.I16))
		},
	)
}
func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15], tins16 fp.Monoid[A16], tins17 fp.Monoid[A17]) fp.Monoid[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return New(
		func() fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
			return product.Tuple17(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty(), tins16.Empty(), tins17.Empty())
		},
		func(t1 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17], t2 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
			return product.Tuple17(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15), tins16.Combine(t1.I16, t2.I16), tins17.Combine(t1.I17, t2.I17))
		},
	)
}
func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15], tins16 fp.Monoid[A16], tins17 fp.Monoid[A17], tins18 fp.Monoid[A18]) fp.Monoid[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return New(
		func() fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
			return product.Tuple18(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty(), tins16.Empty(), tins17.Empty(), tins18.Empty())
		},
		func(t1 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18], t2 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
			return product.Tuple18(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15), tins16.Combine(t1.I16, t2.I16), tins17.Combine(t1.I17, t2.I17), tins18.Combine(t1.I18, t2.I18))
		},
	)
}
func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15], tins16 fp.Monoid[A16], tins17 fp.Monoid[A17], tins18 fp.Monoid[A18], tins19 fp.Monoid[A19]) fp.Monoid[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return New(
		func() fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
			return product.Tuple19(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty(), tins16.Empty(), tins17.Empty(), tins18.Empty(), tins19.Empty())
		},
		func(t1 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19], t2 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
			return product.Tuple19(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15), tins16.Combine(t1.I16, t2.I16), tins17.Combine(t1.I17, t2.I17), tins18.Combine(t1.I18, t2.I18), tins19.Combine(t1.I19, t2.I19))
		},
	)
}
func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15], tins16 fp.Monoid[A16], tins17 fp.Monoid[A17], tins18 fp.Monoid[A18], tins19 fp.Monoid[A19], tins20 fp.Monoid[A20]) fp.Monoid[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return New(
		func() fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
			return product.Tuple20(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty(), tins16.Empty(), tins17.Empty(), tins18.Empty(), tins19.Empty(), tins20.Empty())
		},
		func(t1 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20], t2 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
			return product.Tuple20(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15), tins16.Combine(t1.I16, t2.I16), tins17.Combine(t1.I17, t2.I17), tins18.Combine(t1.I18, t2.I18), tins19.Combine(t1.I19, t2.I19), tins20.Combine(t1.I20, t2.I20))
		},
	)
}
func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](tins1 fp.Monoid[A1], tins2 fp.Monoid[A2], tins3 fp.Monoid[A3], tins4 fp.Monoid[A4], tins5 fp.Monoid[A5], tins6 fp.Monoid[A6], tins7 fp.Monoid[A7], tins8 fp.Monoid[A8], tins9 fp.Monoid[A9], tins10 fp.Monoid[A10], tins11 fp.Monoid[A11], tins12 fp.Monoid[A12], tins13 fp.Monoid[A13], tins14 fp.Monoid[A14], tins15 fp.Monoid[A15], tins16 fp.Monoid[A16], tins17 fp.Monoid[A17], tins18 fp.Monoid[A18], tins19 fp.Monoid[A19], tins20 fp.Monoid[A20], tins21 fp.Monoid[A21]) fp.Monoid[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return New(
		func() fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
			return product.Tuple21(tins1.Empty(), tins2.Empty(), tins3.Empty(), tins4.Empty(), tins5.Empty(), tins6.Empty(), tins7.Empty(), tins8.Empty(), tins9.Empty(), tins10.Empty(), tins11.Empty(), tins12.Empty(), tins13.Empty(), tins14.Empty(), tins15.Empty(), tins16.Empty(), tins17.Empty(), tins18.Empty(), tins19.Empty(), tins20.Empty(), tins21.Empty())
		},
		func(t1 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21], t2 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
			return product.Tuple21(tins1.Combine(t1.I1, t2.I1), tins2.Combine(t1.I2, t2.I2), tins3.Combine(t1.I3, t2.I3), tins4.Combine(t1.I4, t2.I4), tins5.Combine(t1.I5, t2.I5), tins6.Combine(t1.I6, t2.I6), tins7.Combine(t1.I7, t2.I7), tins8.Combine(t1.I8, t2.I8), tins9.Combine(t1.I9, t2.I9), tins10.Combine(t1.I10, t2.I10), tins11.Combine(t1.I11, t2.I11), tins12.Combine(t1.I12, t2.I12), tins13.Combine(t1.I13, t2.I13), tins14.Combine(t1.I14, t2.I14), tins15.Combine(t1.I15, t2.I15), tins16.Combine(t1.I16, t2.I16), tins17.Combine(t1.I17, t2.I17), tins18.Combine(t1.I18, t2.I18), tins19.Combine(t1.I19, t2.I19), tins20.Combine(t1.I20, t2.I20), tins21.Combine(t1.I21, t2.I21))
		},
	)
}
