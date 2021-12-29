package eq

import (
	"github.com/csgura/fp"
)

func Tuple2[A1, A2 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2]) fp.Eq[fp.Tuple2[A1, A2]] {

	pt := Tuple1(ins2)

	return New(
		func(t1, t2 fp.Tuple2[A1, A2]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple3[A1, A2, A3 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3]) fp.Eq[fp.Tuple3[A1, A2, A3]] {

	pt := Tuple2(ins2, ins3)

	return New(
		func(t1, t2 fp.Tuple3[A1, A2, A3]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple4[A1, A2, A3, A4 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4]) fp.Eq[fp.Tuple4[A1, A2, A3, A4]] {

	pt := Tuple3(ins2, ins3, ins4)

	return New(
		func(t1, t2 fp.Tuple4[A1, A2, A3, A4]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple5[A1, A2, A3, A4, A5 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5]) fp.Eq[fp.Tuple5[A1, A2, A3, A4, A5]] {

	pt := Tuple4(ins2, ins3, ins4, ins5)

	return New(
		func(t1, t2 fp.Tuple5[A1, A2, A3, A4, A5]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6]) fp.Eq[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {

	pt := Tuple5(ins2, ins3, ins4, ins5, ins6)

	return New(
		func(t1, t2 fp.Tuple6[A1, A2, A3, A4, A5, A6]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7]) fp.Eq[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {

	pt := Tuple6(ins2, ins3, ins4, ins5, ins6, ins7)

	return New(
		func(t1, t2 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8]) fp.Eq[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {

	pt := Tuple7(ins2, ins3, ins4, ins5, ins6, ins7, ins8)

	return New(
		func(t1, t2 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9]) fp.Eq[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {

	pt := Tuple8(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)

	return New(
		func(t1, t2 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10]) fp.Eq[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {

	pt := Tuple9(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10)

	return New(
		func(t1, t2 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11]) fp.Eq[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {

	pt := Tuple10(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11)

	return New(
		func(t1, t2 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12]) fp.Eq[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {

	pt := Tuple11(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12)

	return New(
		func(t1, t2 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13]) fp.Eq[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {

	pt := Tuple12(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13)

	return New(
		func(t1, t2 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14]) fp.Eq[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {

	pt := Tuple13(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14)

	return New(
		func(t1, t2 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15]) fp.Eq[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {

	pt := Tuple14(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15)

	return New(
		func(t1, t2 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15], ins16 fp.Eq[A16]) fp.Eq[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {

	pt := Tuple15(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16)

	return New(
		func(t1, t2 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15], ins16 fp.Eq[A16], ins17 fp.Eq[A17]) fp.Eq[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {

	pt := Tuple16(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17)

	return New(
		func(t1, t2 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15], ins16 fp.Eq[A16], ins17 fp.Eq[A17], ins18 fp.Eq[A18]) fp.Eq[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {

	pt := Tuple17(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18)

	return New(
		func(t1, t2 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15], ins16 fp.Eq[A16], ins17 fp.Eq[A17], ins18 fp.Eq[A18], ins19 fp.Eq[A19]) fp.Eq[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {

	pt := Tuple18(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19)

	return New(
		func(t1, t2 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15], ins16 fp.Eq[A16], ins17 fp.Eq[A17], ins18 fp.Eq[A18], ins19 fp.Eq[A19], ins20 fp.Eq[A20]) fp.Eq[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {

	pt := Tuple19(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20)

	return New(
		func(t1, t2 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}

func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Eq[A1], ins2 fp.Eq[A2], ins3 fp.Eq[A3], ins4 fp.Eq[A4], ins5 fp.Eq[A5], ins6 fp.Eq[A6], ins7 fp.Eq[A7], ins8 fp.Eq[A8], ins9 fp.Eq[A9], ins10 fp.Eq[A10], ins11 fp.Eq[A11], ins12 fp.Eq[A12], ins13 fp.Eq[A13], ins14 fp.Eq[A14], ins15 fp.Eq[A15], ins16 fp.Eq[A16], ins17 fp.Eq[A17], ins18 fp.Eq[A18], ins19 fp.Eq[A19], ins20 fp.Eq[A20], ins21 fp.Eq[A21]) fp.Eq[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {

	pt := Tuple20(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20, ins21)

	return New(
		func(t1, t2 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(t1.Tail(), t2.Tail())
		},
	)
}
