package hash

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

func Tuple2[A1, A2 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2]) fp.Hashable[fp.Tuple2[A1, A2]] {
	return New(eq.Tuple2[A1, A2](ins1, ins2), func(t fp.Tuple2[A1, A2]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple1(ins2).Hash(t.Tail())
	})
}

func Tuple3[A1, A2, A3 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3]) fp.Hashable[fp.Tuple3[A1, A2, A3]] {
	return New(eq.Tuple3[A1, A2, A3](ins1, ins2, ins3), func(t fp.Tuple3[A1, A2, A3]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple2(ins2, ins3).Hash(t.Tail())
	})
}

func Tuple4[A1, A2, A3, A4 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4]) fp.Hashable[fp.Tuple4[A1, A2, A3, A4]] {
	return New(eq.Tuple4[A1, A2, A3, A4](ins1, ins2, ins3, ins4), func(t fp.Tuple4[A1, A2, A3, A4]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple3(ins2, ins3, ins4).Hash(t.Tail())
	})
}

func Tuple5[A1, A2, A3, A4, A5 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5]) fp.Hashable[fp.Tuple5[A1, A2, A3, A4, A5]] {
	return New(eq.Tuple5[A1, A2, A3, A4, A5](ins1, ins2, ins3, ins4, ins5), func(t fp.Tuple5[A1, A2, A3, A4, A5]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple4(ins2, ins3, ins4, ins5).Hash(t.Tail())
	})
}

func Tuple6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6]) fp.Hashable[fp.Tuple6[A1, A2, A3, A4, A5, A6]] {
	return New(eq.Tuple6[A1, A2, A3, A4, A5, A6](ins1, ins2, ins3, ins4, ins5, ins6), func(t fp.Tuple6[A1, A2, A3, A4, A5, A6]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple5(ins2, ins3, ins4, ins5, ins6).Hash(t.Tail())
	})
}

func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7]) fp.Hashable[fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]] {
	return New(eq.Tuple7[A1, A2, A3, A4, A5, A6, A7](ins1, ins2, ins3, ins4, ins5, ins6, ins7), func(t fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple6(ins2, ins3, ins4, ins5, ins6, ins7).Hash(t.Tail())
	})
}

func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8]) fp.Hashable[fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return New(eq.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8), func(t fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple7(ins2, ins3, ins4, ins5, ins6, ins7, ins8).Hash(t.Tail())
	})
}

func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9]) fp.Hashable[fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return New(eq.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9), func(t fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple8(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9).Hash(t.Tail())
	})
}

func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10]) fp.Hashable[fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return New(eq.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10), func(t fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple9(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10).Hash(t.Tail())
	})
}

func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11]) fp.Hashable[fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return New(eq.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11), func(t fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple10(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11).Hash(t.Tail())
	})
}

func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12]) fp.Hashable[fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return New(eq.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12), func(t fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple11(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12).Hash(t.Tail())
	})
}

func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13]) fp.Hashable[fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return New(eq.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13), func(t fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple12(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13).Hash(t.Tail())
	})
}

func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14]) fp.Hashable[fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return New(eq.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14), func(t fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple13(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14).Hash(t.Tail())
	})
}

func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15]) fp.Hashable[fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return New(eq.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15), func(t fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple14(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15).Hash(t.Tail())
	})
}

func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16]) fp.Hashable[fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return New(eq.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16), func(t fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple15(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16).Hash(t.Tail())
	})
}

func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17]) fp.Hashable[fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return New(eq.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17), func(t fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple16(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17).Hash(t.Tail())
	})
}

func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18]) fp.Hashable[fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return New(eq.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18), func(t fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple17(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18).Hash(t.Tail())
	})
}

func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18], ins19 fp.Hashable[A19]) fp.Hashable[fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return New(eq.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19), func(t fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple18(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19).Hash(t.Tail())
	})
}

func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18], ins19 fp.Hashable[A19], ins20 fp.Hashable[A20]) fp.Hashable[fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return New(eq.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20), func(t fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple19(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20).Hash(t.Tail())
	})
}

func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Hashable[A1], ins2 fp.Hashable[A2], ins3 fp.Hashable[A3], ins4 fp.Hashable[A4], ins5 fp.Hashable[A5], ins6 fp.Hashable[A6], ins7 fp.Hashable[A7], ins8 fp.Hashable[A8], ins9 fp.Hashable[A9], ins10 fp.Hashable[A10], ins11 fp.Hashable[A11], ins12 fp.Hashable[A12], ins13 fp.Hashable[A13], ins14 fp.Hashable[A14], ins15 fp.Hashable[A15], ins16 fp.Hashable[A16], ins17 fp.Hashable[A17], ins18 fp.Hashable[A18], ins19 fp.Hashable[A19], ins20 fp.Hashable[A20], ins21 fp.Hashable[A21]) fp.Hashable[fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return New(eq.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21](ins1, ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20, ins21), func(t fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) uint32 {
		return ins1.Hash(t.Head())*31 + Tuple20(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9, ins10, ins11, ins12, ins13, ins14, ins15, ins16, ins17, ins18, ins19, ins20, ins21).Hash(t.Tail())
	})
}
