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
func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7]) fp.Eq[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {
	return New(
		func(t1 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7], t2 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple6(tins2, tins3, tins4, tins5, tins6, tins7).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8]) fp.Eq[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return New(
		func(t1 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8], t2 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple7(tins2, tins3, tins4, tins5, tins6, tins7, tins8).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9]) fp.Eq[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return New(
		func(t1 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9], t2 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple8(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10]) fp.Eq[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return New(
		func(t1 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], t2 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple9(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11]) fp.Eq[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return New(
		func(t1 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], t2 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple10(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12]) fp.Eq[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return New(
		func(t1 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], t2 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple11(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13]) fp.Eq[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return New(
		func(t1 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], t2 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple12(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14]) fp.Eq[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return New(
		func(t1 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], t2 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple13(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15]) fp.Eq[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return New(
		func(t1 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], t2 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple14(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15], tins16 fp.Eq[A16]) fp.Eq[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return New(
		func(t1 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], t2 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple15(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15], tins16 fp.Eq[A16], tins17 fp.Eq[A17]) fp.Eq[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return New(
		func(t1 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17], t2 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple16(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15], tins16 fp.Eq[A16], tins17 fp.Eq[A17], tins18 fp.Eq[A18]) fp.Eq[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return New(
		func(t1 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18], t2 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple17(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15], tins16 fp.Eq[A16], tins17 fp.Eq[A17], tins18 fp.Eq[A18], tins19 fp.Eq[A19]) fp.Eq[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return New(
		func(t1 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19], t2 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple18(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15], tins16 fp.Eq[A16], tins17 fp.Eq[A17], tins18 fp.Eq[A18], tins19 fp.Eq[A19], tins20 fp.Eq[A20]) fp.Eq[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return New(
		func(t1 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20], t2 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple19(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19, tins20).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](tins1 fp.Eq[A1], tins2 fp.Eq[A2], tins3 fp.Eq[A3], tins4 fp.Eq[A4], tins5 fp.Eq[A5], tins6 fp.Eq[A6], tins7 fp.Eq[A7], tins8 fp.Eq[A8], tins9 fp.Eq[A9], tins10 fp.Eq[A10], tins11 fp.Eq[A11], tins12 fp.Eq[A12], tins13 fp.Eq[A13], tins14 fp.Eq[A14], tins15 fp.Eq[A15], tins16 fp.Eq[A16], tins17 fp.Eq[A17], tins18 fp.Eq[A18], tins19 fp.Eq[A19], tins20 fp.Eq[A20], tins21 fp.Eq[A21]) fp.Eq[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return New(
		func(t1 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21], t2 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) bool {
			return tins1.Eqv(t1.I1, t2.I1) && Tuple20(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19, tins20, tins21).Eqv(t1.Tail(), t2.Tail())
		},
	)
}
