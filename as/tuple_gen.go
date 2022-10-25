package as

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
)

func Tuple1[A1 any](a1 A1) fp.Tuple1[A1] {
	return fp.Tuple1[A1]{
		I1: a1,
	}
}
func Tuple2[A1, A2 any](a1 A1, a2 A2) fp.Tuple2[A1, A2] {
	return fp.Tuple2[A1, A2]{
		I1: a1,
		I2: a2,
	}
}
func Tuple3[A1, A2, A3 any](a1 A1, a2 A2, a3 A3) fp.Tuple3[A1, A2, A3] {
	return fp.Tuple3[A1, A2, A3]{
		I1: a1,
		I2: a2,
		I3: a3,
	}
}
func Tuple4[A1, A2, A3, A4 any](a1 A1, a2 A2, a3 A3, a4 A4) fp.Tuple4[A1, A2, A3, A4] {
	return fp.Tuple4[A1, A2, A3, A4]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
	}
}
func Tuple5[A1, A2, A3, A4, A5 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Tuple5[A1, A2, A3, A4, A5] {
	return fp.Tuple5[A1, A2, A3, A4, A5]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
		I5: a5,
	}
}
func Tuple6[A1, A2, A3, A4, A5, A6 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Tuple6[A1, A2, A3, A4, A5, A6] {
	return fp.Tuple6[A1, A2, A3, A4, A5, A6]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
		I5: a5,
		I6: a6,
	}
}
func Tuple7[A1, A2, A3, A4, A5, A6, A7 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Tuple7[A1, A2, A3, A4, A5, A6, A7] {
	return fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
		I5: a5,
		I6: a6,
		I7: a7,
	}
}
func Tuple8[A1, A2, A3, A4, A5, A6, A7, A8 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8] {
	return fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
		I5: a5,
		I6: a6,
		I7: a7,
		I8: a8,
	}
}
func Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	return fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
		I5: a5,
		I6: a6,
		I7: a7,
		I8: a8,
		I9: a9,
	}
}
func Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	return fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
	}
}
func Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	return fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
	}
}
func Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	return fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
	}
}
func Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	return fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
	}
}
func Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	return fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
	}
}
func Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	return fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
	}
}
func Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	return fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
		I16: a16,
	}
}
func Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
	return fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
		I16: a16,
		I17: a17,
	}
}
func Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
	return fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
		I16: a16,
		I17: a17,
		I18: a18,
	}
}
func Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
	return fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
		I16: a16,
		I17: a17,
		I18: a18,
		I19: a19,
	}
}
func Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
	return fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
		I16: a16,
		I17: a17,
		I18: a18,
		I19: a19,
		I20: a20,
	}
}
func Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
	return fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]{
		I1:  a1,
		I2:  a2,
		I3:  a3,
		I4:  a4,
		I5:  a5,
		I6:  a6,
		I7:  a7,
		I8:  a8,
		I9:  a9,
		I10: a10,
		I11: a11,
		I12: a12,
		I13: a13,
		I14: a14,
		I15: a15,
		I16: a16,
		I17: a17,
		I18: a18,
		I19: a19,
		I20: a20,
		I21: a21,
	}
}

func HList1[A1 any](tuple fp.Tuple1[A1]) hlist.Cons[A1, hlist.Nil] {
	return hlist.Concat(tuple.Head(), hlist.Empty())
}
func HList2[A1, A2 any](tuple fp.Tuple2[A1, A2]) hlist.Cons[A1, hlist.Cons[A2, hlist.Nil]] {
	return hlist.Concat(tuple.Head(), HList1(tuple.Tail()))
}
func HList3[A1, A2, A3 any](tuple fp.Tuple3[A1, A2, A3]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Nil]]] {
	return hlist.Concat(tuple.Head(), HList2(tuple.Tail()))
}
func HList4[A1, A2, A3, A4 any](tuple fp.Tuple4[A1, A2, A3, A4]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Nil]]]] {
	return hlist.Concat(tuple.Head(), HList3(tuple.Tail()))
}
func HList5[A1, A2, A3, A4, A5 any](tuple fp.Tuple5[A1, A2, A3, A4, A5]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Nil]]]]] {
	return hlist.Concat(tuple.Head(), HList4(tuple.Tail()))
}
func HList6[A1, A2, A3, A4, A5, A6 any](tuple fp.Tuple6[A1, A2, A3, A4, A5, A6]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Nil]]]]]] {
	return hlist.Concat(tuple.Head(), HList5(tuple.Tail()))
}
func HList7[A1, A2, A3, A4, A5, A6, A7 any](tuple fp.Tuple7[A1, A2, A3, A4, A5, A6, A7]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Nil]]]]]]] {
	return hlist.Concat(tuple.Head(), HList6(tuple.Tail()))
}
func HList8[A1, A2, A3, A4, A5, A6, A7, A8 any](tuple fp.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Nil]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList7(tuple.Tail()))
}
func HList9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](tuple fp.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Nil]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList8(tuple.Tail()))
}
func HList10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](tuple fp.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Nil]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList9(tuple.Tail()))
}
func HList11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](tuple fp.Tuple11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Nil]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList10(tuple.Tail()))
}
func HList12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](tuple fp.Tuple12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Nil]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList11(tuple.Tail()))
}
func HList13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](tuple fp.Tuple13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Nil]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList12(tuple.Tail()))
}
func HList14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](tuple fp.Tuple14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Nil]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList13(tuple.Tail()))
}
func HList15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](tuple fp.Tuple15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Nil]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList14(tuple.Tail()))
}
func HList16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](tuple fp.Tuple16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Cons[A16, hlist.Nil]]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList15(tuple.Tail()))
}
func HList17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](tuple fp.Tuple17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Cons[A16, hlist.Cons[A17, hlist.Nil]]]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList16(tuple.Tail()))
}
func HList18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](tuple fp.Tuple18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Cons[A16, hlist.Cons[A17, hlist.Cons[A18, hlist.Nil]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList17(tuple.Tail()))
}
func HList19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](tuple fp.Tuple19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Cons[A16, hlist.Cons[A17, hlist.Cons[A18, hlist.Cons[A19, hlist.Nil]]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList18(tuple.Tail()))
}
func HList20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](tuple fp.Tuple20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Cons[A16, hlist.Cons[A17, hlist.Cons[A18, hlist.Cons[A19, hlist.Cons[A20, hlist.Nil]]]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList19(tuple.Tail()))
}
func HList21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](tuple fp.Tuple21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]) hlist.Cons[A1, hlist.Cons[A2, hlist.Cons[A3, hlist.Cons[A4, hlist.Cons[A5, hlist.Cons[A6, hlist.Cons[A7, hlist.Cons[A8, hlist.Cons[A9, hlist.Cons[A10, hlist.Cons[A11, hlist.Cons[A12, hlist.Cons[A13, hlist.Cons[A14, hlist.Cons[A15, hlist.Cons[A16, hlist.Cons[A17, hlist.Cons[A18, hlist.Cons[A19, hlist.Cons[A20, hlist.Cons[A21, hlist.Nil]]]]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(tuple.Head(), HList20(tuple.Tail()))
}
