package hash

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

func Tuple2[A1, A2 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2]) fp.Hashable[fp.Tuple2[A1, A2]] {

	pt := Tuple1(ins2)

	return New(eq.New(func(a, b fp.Tuple2[A1, A2]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple2[A1, A2]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple3[A1, A2, A3 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3]) fp.Hashable[fp.Tuple3[A1, A2, A3]] {

	pt := Tuple2(ins2, ins3)

	return New(eq.New(func(a, b fp.Tuple3[A1, A2, A3]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple3[A1, A2, A3]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple4[A1, A2, A3, A4 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4]) fp.Hashable[fp.Tuple4[A1, A2, A3, A4]] {

	pt := Tuple3(ins2, ins3, ins4)

	return New(eq.New(func(a, b fp.Tuple4[A1, A2, A3, A4]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple4[A1, A2, A3, A4]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple5[A1, A2, A3, A4, A5 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5]) fp.Hashable[fp.Tuple5[A1, A2, A3, A4, A5]] {

	pt := Tuple4(ins2, ins3, ins4, ins5)

	return New(eq.New(func(a, b fp.Tuple5[A1, A2, A3, A4, A5]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple5[A1, A2, A3, A4, A5]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6]) fp.Hashable[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {

	pt := Tuple5(ins2, ins3, ins4, ins5, ins6)

	return New(eq.New(func(a, b fp.Tuple6[A1, A2, A3, A4, A5, A6]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple6[A1, A2, A3, A4, A5, A6]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7]) fp.Hashable[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {

	pt := Tuple6(ins2, ins3, ins4, ins5, ins6, ins7)

	return New(eq.New(func(a, b fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8]) fp.Hashable[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {

	pt := Tuple7(ins2, ins3, ins4, ins5, ins6, ins7, ins8)

	return New(eq.New(func(a, b fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9]) fp.Hashable[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {

	pt := Tuple8(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)

	return New(eq.New(func(a, b fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10]) fp.Hashable[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {

	pt := Tuple9(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10)

	return New(eq.New(func(a, b fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11]) fp.Hashable[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {

	pt := Tuple10(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11)

	return New(eq.New(func(a, b fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12]) fp.Hashable[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {

	pt := Tuple11(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12)

	return New(eq.New(func(a, b fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13]) fp.Hashable[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {

	pt := Tuple12(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13)

	return New(eq.New(func(a, b fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14]) fp.Hashable[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {

	pt := Tuple13(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14)

	return New(eq.New(func(a, b fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15]) fp.Hashable[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {

	pt := Tuple14(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15)

	return New(eq.New(func(a, b fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16]) fp.Hashable[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {

	pt := Tuple15(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16)

	return New(eq.New(func(a, b fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17]) fp.Hashable[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {

	pt := Tuple16(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17)

	return New(eq.New(func(a, b fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18]) fp.Hashable[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {

	pt := Tuple17(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18)

	return New(eq.New(func(a, b fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18], ins19 fp.Hashable[A19]) fp.Hashable[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {

	pt := Tuple18(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19)

	return New(eq.New(func(a, b fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18], ins19 fp.Hashable[A19], ins20 fp.Hashable[A20]) fp.Hashable[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {

	pt := Tuple19(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20)

	return New(eq.New(func(a, b fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}

func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18], ins19 fp.Hashable[A19], ins20 fp.Hashable[A20], ins21 fp.Hashable[A21]) fp.Hashable[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {

	pt := Tuple20(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20, ins21)

	return New(eq.New(func(a, b fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) bool {
		return ins1.Eqv(a.Head(), b.Head()) && pt.Eqv(a.Tail(), b.Tail())
	}), func(t fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) uint32 {
		return ins1.Hash(t.Head())*31 + pt.Hash(t.Tail())
	})
}
