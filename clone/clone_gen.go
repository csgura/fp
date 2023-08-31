// Code generated by gombok, DO NOT EDIT.
package clone

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

func Tuple3[A1, A2, A3 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3]) fp.Clone[fp.Tuple3[A1, A2, A3]] {
	return New(func(t fp.Tuple3[A1, A2, A3]) fp.Tuple3[A1, A2, A3] {
		return as.Tuple3(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
		)
	})
}

func Tuple4[A1, A2, A3, A4 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4]) fp.Clone[fp.Tuple4[A1, A2, A3, A4]] {
	return New(func(t fp.Tuple4[A1, A2, A3, A4]) fp.Tuple4[A1, A2, A3, A4] {
		return as.Tuple4(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
		)
	})
}

func Tuple5[A1, A2, A3, A4, A5 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5]) fp.Clone[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return New(func(t fp.Tuple5[A1, A2, A3, A4, A5]) fp.Tuple5[A1, A2, A3, A4, A5] {
		return as.Tuple5(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
		)
	})
}

func Tuple6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6]) fp.Clone[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {
	return New(func(t fp.Tuple6[A1, A2, A3, A4, A5, A6]) fp.Tuple6[A1, A2, A3, A4, A5, A6] {
		return as.Tuple6(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
		)
	})
}

func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7]) fp.Clone[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {
	return New(func(t fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) fp.Tuple7[A1, A2, A3, A4, A5, A6, A7] {
		return as.Tuple7(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
		)
	})
}

func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8]) fp.Clone[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return New(func(t fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8] {
		return as.Tuple8(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
		)
	})
}

func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9]) fp.Clone[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return New(func(t fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
		return as.Tuple9(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
		)
	})
}

func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10]) fp.Clone[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return New(func(t fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
		return as.Tuple10(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
		)
	})
}

func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11]) fp.Clone[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return New(func(t fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
		return as.Tuple11(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
		)
	})
}

func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12]) fp.Clone[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return New(func(t fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
		return as.Tuple12(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
		)
	})
}

func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13]) fp.Clone[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return New(func(t fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
		return as.Tuple13(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
		)
	})
}

func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14]) fp.Clone[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return New(func(t fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
		return as.Tuple14(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
		)
	})
}

func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15]) fp.Clone[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return New(func(t fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
		return as.Tuple15(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
		)
	})
}

func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15], ins16 fp.Clone[A16]) fp.Clone[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return New(func(t fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
		return as.Tuple16(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
			ins16.Clone(t.I16),
		)
	})
}

func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15], ins16 fp.Clone[A16], ins17 fp.Clone[A17]) fp.Clone[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return New(func(t fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
		return as.Tuple17(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
			ins16.Clone(t.I16),
			ins17.Clone(t.I17),
		)
	})
}

func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15], ins16 fp.Clone[A16], ins17 fp.Clone[A17], ins18 fp.Clone[A18]) fp.Clone[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return New(func(t fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
		return as.Tuple18(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
			ins16.Clone(t.I16),
			ins17.Clone(t.I17),
			ins18.Clone(t.I18),
		)
	})
}

func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15], ins16 fp.Clone[A16], ins17 fp.Clone[A17], ins18 fp.Clone[A18], ins19 fp.Clone[A19]) fp.Clone[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return New(func(t fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
		return as.Tuple19(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
			ins16.Clone(t.I16),
			ins17.Clone(t.I17),
			ins18.Clone(t.I18),
			ins19.Clone(t.I19),
		)
	})
}

func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15], ins16 fp.Clone[A16], ins17 fp.Clone[A17], ins18 fp.Clone[A18], ins19 fp.Clone[A19], ins20 fp.Clone[A20]) fp.Clone[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return New(func(t fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
		return as.Tuple20(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
			ins16.Clone(t.I16),
			ins17.Clone(t.I17),
			ins18.Clone(t.I18),
			ins19.Clone(t.I19),
			ins20.Clone(t.I20),
		)
	})
}

func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2], ins3 fp.Clone[A3], ins4 fp.Clone[A4], ins5 fp.Clone[A5], ins6 fp.Clone[A6], ins7 fp.Clone[A7], ins8 fp.Clone[A8], ins9 fp.Clone[A9], ins10 fp.Clone[A10], ins11 fp.Clone[A11], ins12 fp.Clone[A12], ins13 fp.Clone[A13], ins14 fp.Clone[A14], ins15 fp.Clone[A15], ins16 fp.Clone[A16], ins17 fp.Clone[A17], ins18 fp.Clone[A18], ins19 fp.Clone[A19], ins20 fp.Clone[A20], ins21 fp.Clone[A21]) fp.Clone[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return New(func(t fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
		return as.Tuple21(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
			ins3.Clone(t.I3),
			ins4.Clone(t.I4),
			ins5.Clone(t.I5),
			ins6.Clone(t.I6),
			ins7.Clone(t.I7),
			ins8.Clone(t.I8),
			ins9.Clone(t.I9),
			ins10.Clone(t.I10),
			ins11.Clone(t.I11),
			ins12.Clone(t.I12),
			ins13.Clone(t.I13),
			ins14.Clone(t.I14),
			ins15.Clone(t.I15),
			ins16.Clone(t.I16),
			ins17.Clone(t.I17),
			ins18.Clone(t.I18),
			ins19.Clone(t.I19),
			ins20.Clone(t.I20),
			ins21.Clone(t.I21),
		)
	})
}
