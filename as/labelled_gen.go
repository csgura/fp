package as

import (
	"github.com/csgura/fp"
)

func Labelled1[A1 any](ins1 fp.Field[A1]) fp.Labelled1[A1] {
	return fp.Labelled1[A1]{
		I1: ins1,
	}
}
func Labelled2[A1, A2 any](ins1 fp.Field[A1], ins2 fp.Field[A2]) fp.Labelled2[A1, A2] {
	return fp.Labelled2[A1, A2]{
		I1: ins1,
		I2: ins2,
	}
}
func Labelled3[A1, A2, A3 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3]) fp.Labelled3[A1, A2, A3] {
	return fp.Labelled3[A1, A2, A3]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
	}
}
func Labelled4[A1, A2, A3, A4 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4]) fp.Labelled4[A1, A2, A3, A4] {
	return fp.Labelled4[A1, A2, A3, A4]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
		I4: ins4,
	}
}
func Labelled5[A1, A2, A3, A4, A5 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5]) fp.Labelled5[A1, A2, A3, A4, A5] {
	return fp.Labelled5[A1, A2, A3, A4, A5]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
		I4: ins4,
		I5: ins5,
	}
}
func Labelled6[A1, A2, A3, A4, A5, A6 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6]) fp.Labelled6[A1, A2, A3, A4, A5, A6] {
	return fp.Labelled6[A1, A2, A3, A4, A5, A6]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
		I4: ins4,
		I5: ins5,
		I6: ins6,
	}
}
func Labelled7[A1, A2, A3, A4, A5, A6, A7 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7]) fp.Labelled7[A1, A2, A3, A4, A5, A6, A7] {
	return fp.Labelled7[A1, A2, A3, A4, A5, A6, A7]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
		I4: ins4,
		I5: ins5,
		I6: ins6,
		I7: ins7,
	}
}
func Labelled8[A1, A2, A3, A4, A5, A6, A7, A8 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8]) fp.Labelled8[A1, A2, A3, A4, A5, A6, A7, A8] {
	return fp.Labelled8[A1, A2, A3, A4, A5, A6, A7, A8]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
		I4: ins4,
		I5: ins5,
		I6: ins6,
		I7: ins7,
		I8: ins8,
	}
}
func Labelled9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9]) fp.Labelled9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	return fp.Labelled9[A1, A2, A3, A4, A5, A6, A7, A8, A9]{
		I1: ins1,
		I2: ins2,
		I3: ins3,
		I4: ins4,
		I5: ins5,
		I6: ins6,
		I7: ins7,
		I8: ins8,
		I9: ins9,
	}
}
func Labelled10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10]) fp.Labelled10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	return fp.Labelled10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
	}
}
func Labelled11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11]) fp.Labelled11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	return fp.Labelled11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
	}
}
func Labelled12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12]) fp.Labelled12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	return fp.Labelled12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
	}
}
func Labelled13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13]) fp.Labelled13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	return fp.Labelled13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
	}
}
func Labelled14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14]) fp.Labelled14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	return fp.Labelled14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
	}
}
func Labelled15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15]) fp.Labelled15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	return fp.Labelled15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
	}
}
func Labelled16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15], ins16 fp.Field[A16]) fp.Labelled16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	return fp.Labelled16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
		I16: ins16,
	}
}
func Labelled17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15], ins16 fp.Field[A16], ins17 fp.Field[A17]) fp.Labelled17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17] {
	return fp.Labelled17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
		I16: ins16,
		I17: ins17,
	}
}
func Labelled18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15], ins16 fp.Field[A16], ins17 fp.Field[A17], ins18 fp.Field[A18]) fp.Labelled18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18] {
	return fp.Labelled18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
		I16: ins16,
		I17: ins17,
		I18: ins18,
	}
}
func Labelled19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15], ins16 fp.Field[A16], ins17 fp.Field[A17], ins18 fp.Field[A18], ins19 fp.Field[A19]) fp.Labelled19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19] {
	return fp.Labelled19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
		I16: ins16,
		I17: ins17,
		I18: ins18,
		I19: ins19,
	}
}
func Labelled20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15], ins16 fp.Field[A16], ins17 fp.Field[A17], ins18 fp.Field[A18], ins19 fp.Field[A19], ins20 fp.Field[A20]) fp.Labelled20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20] {
	return fp.Labelled20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
		I16: ins16,
		I17: ins17,
		I18: ins18,
		I19: ins19,
		I20: ins20,
	}
}
func Labelled21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](ins1 fp.Field[A1], ins2 fp.Field[A2], ins3 fp.Field[A3], ins4 fp.Field[A4], ins5 fp.Field[A5], ins6 fp.Field[A6], ins7 fp.Field[A7], ins8 fp.Field[A8], ins9 fp.Field[A9], ins10 fp.Field[A10], ins11 fp.Field[A11], ins12 fp.Field[A12], ins13 fp.Field[A13], ins14 fp.Field[A14], ins15 fp.Field[A15], ins16 fp.Field[A16], ins17 fp.Field[A17], ins18 fp.Field[A18], ins19 fp.Field[A19], ins20 fp.Field[A20], ins21 fp.Field[A21]) fp.Labelled21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21] {
	return fp.Labelled21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]{
		I1:  ins1,
		I2:  ins2,
		I3:  ins3,
		I4:  ins4,
		I5:  ins5,
		I6:  ins6,
		I7:  ins7,
		I8:  ins8,
		I9:  ins9,
		I10: ins10,
		I11: ins11,
		I12: ins12,
		I13: ins13,
		I14: ins14,
		I15: ins15,
		I16: ins16,
		I17: ins17,
		I18: ins18,
		I19: ins19,
		I20: ins20,
		I21: ins21,
	}
}
