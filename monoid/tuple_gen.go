package monoid

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/product"
)

func Tuple2[A1, A2 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2]) fp.Monoid[fp.Tuple2[A1, A2]] {
	return New(
		func() fp.Tuple2[A1, A2] {
			return product.Tuple2(ins1.Empty(), ins2.Empty())
		},
		func(t1 fp.Tuple2[A1, A2], t2 fp.Tuple2[A1, A2]) fp.Tuple2[A1, A2] {
			return product.Tuple2(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2))
		},
	)
}
func Tuple3[A1, A2, A3 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3]) fp.Monoid[fp.Tuple3[A1, A2, A3]] {
	return New(
		func() fp.Tuple3[A1, A2, A3] {
			return product.Tuple3(ins1.Empty(), ins2.Empty(), ins3.Empty())
		},
		func(t1 fp.Tuple3[A1, A2, A3], t2 fp.Tuple3[A1, A2, A3]) fp.Tuple3[A1, A2, A3] {
			return product.Tuple3(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3))
		},
	)
}
func Tuple4[A1, A2, A3, A4 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4]) fp.Monoid[fp.Tuple4[A1, A2, A3, A4]] {
	return New(
		func() fp.Tuple4[A1, A2, A3, A4] {
			return product.Tuple4(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty())
		},
		func(t1 fp.Tuple4[A1, A2, A3, A4], t2 fp.Tuple4[A1, A2, A3, A4]) fp.Tuple4[A1, A2, A3, A4] {
			return product.Tuple4(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4))
		},
	)
}
func Tuple5[A1, A2, A3, A4, A5 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5]) fp.Monoid[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return New(
		func() fp.Tuple5[A1, A2, A3, A4, A5] {
			return product.Tuple5(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty())
		},
		func(t1 fp.Tuple5[A1, A2, A3, A4, A5], t2 fp.Tuple5[A1, A2, A3, A4, A5]) fp.Tuple5[A1, A2, A3, A4, A5] {
			return product.Tuple5(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5))
		},
	)
}
func Tuple6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6]) fp.Monoid[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {
	return New(
		func() fp.Tuple6[A1, A2, A3, A4, A5, A6] {
			return product.Tuple6(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty())
		},
		func(t1 fp.Tuple6[A1, A2, A3, A4, A5, A6], t2 fp.Tuple6[A1, A2, A3, A4, A5, A6]) fp.Tuple6[A1, A2, A3, A4, A5, A6] {
			return product.Tuple6(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6))
		},
	)
}
func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7]) fp.Monoid[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {
	return New(
		func() fp.Tuple7[A1, A2, A3, A4, A5, A6, A7] {
			return product.Tuple7(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty())
		},
		func(t1 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7], t2 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) fp.Tuple7[A1, A2, A3, A4, A5, A6, A7] {
			return product.Tuple7(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7))
		},
	)
}
func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8]) fp.Monoid[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return New(
		func() fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8] {
			return product.Tuple8(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty())
		},
		func(t1 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8], t2 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8] {
			return product.Tuple8(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8))
		},
	)
}
func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9]) fp.Monoid[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return New(
		func() fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
			return product.Tuple9(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty())
		},
		func(t1 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9], t2 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
			return product.Tuple9(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9))
		},
	)
}
func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10]) fp.Monoid[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return New(
		func() fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
			return product.Tuple10(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty())
		},
		func(t1 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], t2 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
			return product.Tuple10(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10))
		},
	)
}
func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11]) fp.Monoid[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return New(
		func() fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
			return product.Tuple11(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty())
		},
		func(t1 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], t2 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
			return product.Tuple11(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11))
		},
	)
}
func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12]) fp.Monoid[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return New(
		func() fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
			return product.Tuple12(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty())
		},
		func(t1 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], t2 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
			return product.Tuple12(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12))
		},
	)
}
func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13]) fp.Monoid[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return New(
		func() fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
			return product.Tuple13(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty())
		},
		func(t1 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], t2 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
			return product.Tuple13(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13))
		},
	)
}
func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14]) fp.Monoid[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return New(
		func() fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
			return product.Tuple14(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty())
		},
		func(t1 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], t2 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
			return product.Tuple14(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14))
		},
	)
}
func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15]) fp.Monoid[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return New(
		func() fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
			return product.Tuple15(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty())
		},
		func(t1 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], t2 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
			return product.Tuple15(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15))
		},
	)
}
func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15], ins16 fp.Monoid[A16]) fp.Monoid[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return New(
		func() fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
			return product.Tuple16(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty(), ins16.Empty())
		},
		func(t1 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], t2 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
			return product.Tuple16(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15), ins16.Combine(t1.I16, t2.I16))
		},
	)
}
func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15], ins16 fp.Monoid[A16], ins17 fp.Monoid[A17]) fp.Monoid[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return New(
		func() fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
			return product.Tuple17(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty(), ins16.Empty(), ins17.Empty())
		},
		func(t1 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17], t2 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
			return product.Tuple17(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15), ins16.Combine(t1.I16, t2.I16), ins17.Combine(t1.I17, t2.I17))
		},
	)
}
func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15], ins16 fp.Monoid[A16], ins17 fp.Monoid[A17], ins18 fp.Monoid[A18]) fp.Monoid[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return New(
		func() fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
			return product.Tuple18(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty(), ins16.Empty(), ins17.Empty(), ins18.Empty())
		},
		func(t1 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18], t2 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
			return product.Tuple18(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15), ins16.Combine(t1.I16, t2.I16), ins17.Combine(t1.I17, t2.I17), ins18.Combine(t1.I18, t2.I18))
		},
	)
}
func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15], ins16 fp.Monoid[A16], ins17 fp.Monoid[A17], ins18 fp.Monoid[A18], ins19 fp.Monoid[A19]) fp.Monoid[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return New(
		func() fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
			return product.Tuple19(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty(), ins16.Empty(), ins17.Empty(), ins18.Empty(), ins19.Empty())
		},
		func(t1 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19], t2 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
			return product.Tuple19(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15), ins16.Combine(t1.I16, t2.I16), ins17.Combine(t1.I17, t2.I17), ins18.Combine(t1.I18, t2.I18), ins19.Combine(t1.I19, t2.I19))
		},
	)
}
func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15], ins16 fp.Monoid[A16], ins17 fp.Monoid[A17], ins18 fp.Monoid[A18], ins19 fp.Monoid[A19], ins20 fp.Monoid[A20]) fp.Monoid[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return New(
		func() fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
			return product.Tuple20(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty(), ins16.Empty(), ins17.Empty(), ins18.Empty(), ins19.Empty(), ins20.Empty())
		},
		func(t1 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20], t2 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
			return product.Tuple20(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15), ins16.Combine(t1.I16, t2.I16), ins17.Combine(t1.I17, t2.I17), ins18.Combine(t1.I18, t2.I18), ins19.Combine(t1.I19, t2.I19), ins20.Combine(t1.I20, t2.I20))
		},
	)
}
func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Monoid[A1], ins2 fp.Monoid[A2], ins3 fp.Monoid[A3], ins4 fp.Monoid[A4], ins5 fp.Monoid[A5], ins6 fp.Monoid[A6], ins7 fp.Monoid[A7], ins8 fp.Monoid[A8], ins9 fp.Monoid[A9], ins10 fp.Monoid[A10], ins11 fp.Monoid[A11], ins12 fp.Monoid[A12], ins13 fp.Monoid[A13], ins14 fp.Monoid[A14], ins15 fp.Monoid[A15], ins16 fp.Monoid[A16], ins17 fp.Monoid[A17], ins18 fp.Monoid[A18], ins19 fp.Monoid[A19], ins20 fp.Monoid[A20], ins21 fp.Monoid[A21]) fp.Monoid[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return New(
		func() fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
			return product.Tuple21(ins1.Empty(), ins2.Empty(), ins3.Empty(), ins4.Empty(), ins5.Empty(), ins6.Empty(), ins7.Empty(), ins8.Empty(), ins9.Empty(), ins10.Empty(), ins11.Empty(), ins12.Empty(), ins13.Empty(), ins14.Empty(), ins15.Empty(), ins16.Empty(), ins17.Empty(), ins18.Empty(), ins19.Empty(), ins20.Empty(), ins21.Empty())
		},
		func(t1 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21], t2 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
			return product.Tuple21(ins1.Combine(t1.I1, t2.I1), ins2.Combine(t1.I2, t2.I2), ins3.Combine(t1.I3, t2.I3), ins4.Combine(t1.I4, t2.I4), ins5.Combine(t1.I5, t2.I5), ins6.Combine(t1.I6, t2.I6), ins7.Combine(t1.I7, t2.I7), ins8.Combine(t1.I8, t2.I8), ins9.Combine(t1.I9, t2.I9), ins10.Combine(t1.I10, t2.I10), ins11.Combine(t1.I11, t2.I11), ins12.Combine(t1.I12, t2.I12), ins13.Combine(t1.I13, t2.I13), ins14.Combine(t1.I14, t2.I14), ins15.Combine(t1.I15, t2.I15), ins16.Combine(t1.I16, t2.I16), ins17.Combine(t1.I17, t2.I17), ins18.Combine(t1.I18, t2.I18), ins19.Combine(t1.I19, t2.I19), ins20.Combine(t1.I20, t2.I20), ins21.Combine(t1.I21, t2.I21))
		},
	)
}
