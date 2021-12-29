package ord

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

func Tuple2[A1, A2 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2]) fp.Ord[fp.Tuple2[A1, A2]] {

	pt := Tuple1(ins2)

	return New(eq.New(func(a, b fp.Tuple2[A1, A2]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple2[A1, A2]](func(t1, t2 fp.Tuple2[A1, A2]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple3[A1, A2, A3 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3]) fp.Ord[fp.Tuple3[A1, A2, A3]] {

	pt := Tuple2(ins2, ins3)

	return New(eq.New(func(a, b fp.Tuple3[A1, A2, A3]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple3[A1, A2, A3]](func(t1, t2 fp.Tuple3[A1, A2, A3]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple4[A1, A2, A3, A4 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4]) fp.Ord[fp.Tuple4[A1, A2, A3, A4]] {

	pt := Tuple3(ins2, ins3, ins4)

	return New(eq.New(func(a, b fp.Tuple4[A1, A2, A3, A4]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple4[A1, A2, A3, A4]](func(t1, t2 fp.Tuple4[A1, A2, A3, A4]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple5[A1, A2, A3, A4, A5 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5]) fp.Ord[fp.Tuple5[A1, A2, A3, A4, A5]] {

	pt := Tuple4(ins2, ins3, ins4, ins5)

	return New(eq.New(func(a, b fp.Tuple5[A1, A2, A3, A4, A5]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple5[A1, A2, A3, A4, A5]](func(t1, t2 fp.Tuple5[A1, A2, A3, A4, A5]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6]) fp.Ord[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {

	pt := Tuple5(ins2, ins3, ins4, ins5, ins6)

	return New(eq.New(func(a, b fp.Tuple6[A1, A2, A3, A4, A5, A6]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple6[A1, A2, A3, A4, A5, A6]](func(t1, t2 fp.Tuple6[A1, A2, A3, A4, A5, A6]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7]) fp.Ord[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {

	pt := Tuple6(ins2, ins3, ins4, ins5, ins6, ins7)

	return New(eq.New(func(a, b fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]](func(t1, t2 fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8]) fp.Ord[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {

	pt := Tuple7(ins2, ins3, ins4, ins5, ins6, ins7, ins8)

	return New(eq.New(func(a, b fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]](func(t1, t2 fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9]) fp.Ord[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {

	pt := Tuple8(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)

	return New(eq.New(func(a, b fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]](func(t1, t2 fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10]) fp.Ord[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {

	pt := Tuple9(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10)

	return New(eq.New(func(a, b fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]](func(t1, t2 fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11]) fp.Ord[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {

	pt := Tuple10(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11)

	return New(eq.New(func(a, b fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]](func(t1, t2 fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12]) fp.Ord[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {

	pt := Tuple11(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12)

	return New(eq.New(func(a, b fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]](func(t1, t2 fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13]) fp.Ord[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {

	pt := Tuple12(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13)

	return New(eq.New(func(a, b fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]](func(t1, t2 fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14]) fp.Ord[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {

	pt := Tuple13(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14)

	return New(eq.New(func(a, b fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]](func(t1, t2 fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15]) fp.Ord[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {

	pt := Tuple14(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15)

	return New(eq.New(func(a, b fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]](func(t1, t2 fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15], ins16 fp.Ord[A16]) fp.Ord[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {

	pt := Tuple15(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16)

	return New(eq.New(func(a, b fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]](func(t1, t2 fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15], ins16 fp.Ord[A16], ins17 fp.Ord[A17]) fp.Ord[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {

	pt := Tuple16(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17)

	return New(eq.New(func(a, b fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]](func(t1, t2 fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15], ins16 fp.Ord[A16], ins17 fp.Ord[A17], ins18 fp.Ord[A18]) fp.Ord[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {

	pt := Tuple17(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18)

	return New(eq.New(func(a, b fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]](func(t1, t2 fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15], ins16 fp.Ord[A16], ins17 fp.Ord[A17], ins18 fp.Ord[A18], ins19 fp.Ord[A19]) fp.Ord[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {

	pt := Tuple18(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19)

	return New(eq.New(func(a, b fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]](func(t1, t2 fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15], ins16 fp.Ord[A16], ins17 fp.Ord[A17], ins18 fp.Ord[A18], ins19 fp.Ord[A19], ins20 fp.Ord[A20]) fp.Ord[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {

	pt := Tuple19(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20)

	return New(eq.New(func(a, b fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]](func(t1, t2 fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}

func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Ord[A1], ins2 fp.Ord[A2], ins3 fp.Ord[A3], ins4 fp.Ord[A4], ins5 fp.Ord[A5], ins6 fp.Ord[A6], ins7 fp.Ord[A7], ins8 fp.Ord[A8], ins9 fp.Ord[A9], ins10 fp.Ord[A10], ins11 fp.Ord[A11], ins12 fp.Ord[A12], ins13 fp.Ord[A13], ins14 fp.Ord[A14], ins15 fp.Ord[A15], ins16 fp.Ord[A16], ins17 fp.Ord[A17], ins18 fp.Ord[A18], ins19 fp.Ord[A19], ins20 fp.Ord[A20], ins21 fp.Ord[A21]) fp.Ord[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {

	pt := Tuple20(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20, ins21)

	return New(eq.New(func(a, b fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), fp.LessFunc[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]](func(t1, t2 fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(t1.Tail(), t2.Tail())
	}))
}
