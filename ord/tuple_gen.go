package ord

import (
	"github.com/csgura/fp"
)

func Tuple2[A1, A2 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2]) fp.Ord[fp.Tuple2[A1, A2]] {
	return fp.LessFunc[fp.Tuple2[A1, A2]](func(t1 fp.Tuple2[A1, A2], t2 fp.Tuple2[A1, A2]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple1(tins2).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple3[A1, A2, A3 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3]) fp.Ord[fp.Tuple3[A1, A2, A3]] {
	return fp.LessFunc[fp.Tuple3[A1, A2, A3]](func(t1 fp.Tuple3[A1, A2, A3], t2 fp.Tuple3[A1, A2, A3]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple2(tins2, tins3).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple4[A1, A2, A3, A4 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4]) fp.Ord[fp.Tuple4[A1, A2, A3, A4]] {
	return fp.LessFunc[fp.Tuple4[A1, A2, A3, A4]](func(t1 fp.Tuple4[A1, A2, A3, A4], t2 fp.Tuple4[A1, A2, A3, A4]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple3(tins2, tins3, tins4).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple5[A1, A2, A3, A4, A5 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5]) fp.Ord[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return fp.LessFunc[fp.Tuple5[A1, A2, A3, A4, A5]](func(t1 fp.Tuple5[A1, A2, A3, A4, A5], t2 fp.Tuple5[A1, A2, A3, A4, A5]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple4(tins2, tins3, tins4, tins5).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple6[A1, A2, A3, A4, A5, A6 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6]) fp.Ord[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {
	return fp.LessFunc[fp.Tuple6[A1, A2, A3, A4, A5, A6]](func(t1 fp.Tuple6[A1, A2, A3, A4, A5, A6], t2 fp.Tuple6[A1, A2, A3, A4, A5, A6]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple5(tins2, tins3, tins4, tins5, tins6).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7]) fp.Ord[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {
	return fp.LessFunc[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]](func(t1 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7], t2 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple6(tins2, tins3, tins4, tins5, tins6, tins7).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8]) fp.Ord[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return fp.LessFunc[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]](func(t1 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8], t2 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple7(tins2, tins3, tins4, tins5, tins6, tins7, tins8).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9]) fp.Ord[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return fp.LessFunc[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]](func(t1 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9], t2 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple8(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10]) fp.Ord[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return fp.LessFunc[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]](func(t1 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], t2 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple9(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11]) fp.Ord[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return fp.LessFunc[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]](func(t1 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], t2 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple10(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12]) fp.Ord[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return fp.LessFunc[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]](func(t1 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], t2 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple11(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13]) fp.Ord[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return fp.LessFunc[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]](func(t1 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], t2 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple12(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14]) fp.Ord[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return fp.LessFunc[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]](func(t1 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], t2 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple13(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15]) fp.Ord[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return fp.LessFunc[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]](func(t1 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], t2 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple14(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16]) fp.Ord[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return fp.LessFunc[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]](func(t1 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], t2 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple15(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16], tins17 fp.Ord[A17]) fp.Ord[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return fp.LessFunc[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]](func(t1 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17], t2 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple16(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16], tins17 fp.Ord[A17], tins18 fp.Ord[A18]) fp.Ord[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return fp.LessFunc[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]](func(t1 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18], t2 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple17(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16], tins17 fp.Ord[A17], tins18 fp.Ord[A18], tins19 fp.Ord[A19]) fp.Ord[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return fp.LessFunc[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]](func(t1 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19], t2 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple18(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16], tins17 fp.Ord[A17], tins18 fp.Ord[A18], tins19 fp.Ord[A19], tins20 fp.Ord[A20]) fp.Ord[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return fp.LessFunc[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]](func(t1 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20], t2 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple19(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19, tins20).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16], tins17 fp.Ord[A17], tins18 fp.Ord[A18], tins19 fp.Ord[A19], tins20 fp.Ord[A20], tins21 fp.Ord[A21]) fp.Ord[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return fp.LessFunc[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]](func(t1 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21], t2 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple20(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19, tins20, tins21).Less(t1.Tail(), t2.Tail())
	})
}
func Tuple22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22 any](tins1 fp.Ord[A1], tins2 fp.Ord[A2], tins3 fp.Ord[A3], tins4 fp.Ord[A4], tins5 fp.Ord[A5], tins6 fp.Ord[A6], tins7 fp.Ord[A7], tins8 fp.Ord[A8], tins9 fp.Ord[A9], tins10 fp.Ord[A10], tins11 fp.Ord[A11], tins12 fp.Ord[A12], tins13 fp.Ord[A13], tins14 fp.Ord[A14], tins15 fp.Ord[A15], tins16 fp.Ord[A16], tins17 fp.Ord[A17], tins18 fp.Ord[A18], tins19 fp.Ord[A19], tins20 fp.Ord[A20], tins21 fp.Ord[A21], tins22 fp.Ord[A22]) fp.Ord[fp.Tuple22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22]] {
	return fp.LessFunc[fp.Tuple22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22]](func(t1 fp.Tuple22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22], t2 fp.Tuple22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22]) bool {
		if tins1.Less(t1.I1, t2.I1) {
			return true
		}
		if tins1.Less(t2.I1, t1.I1) {
			return false
		}
		return Tuple21(tins2, tins3, tins4, tins5, tins6, tins7, tins8, tins9, tins10, tins11, tins12, tins13, tins14, tins15, tins16, tins17, tins18, tins19, tins20, tins21, tins22).Less(t1.Tail(), t2.Tail())
	})
}
