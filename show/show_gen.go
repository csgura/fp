// Code generated by gombok, DO NOT EDIT.
package show

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/iterator"
)

func Labelled3[A1, A2, A3 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3]) fp.Show[fp.Labelled3[A1, A2, A3]] {
	return NewAppend(func(buf []string, t fp.Labelled3[A1, A2, A3], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled4[A1, A2, A3, A4 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4]) fp.Show[fp.Labelled4[A1, A2, A3, A4]] {
	return NewAppend(func(buf []string, t fp.Labelled4[A1, A2, A3, A4], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled5[A1, A2, A3, A4, A5 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5]) fp.Show[fp.Labelled5[A1, A2, A3, A4, A5]] {
	return NewAppend(func(buf []string, t fp.Labelled5[A1, A2, A3, A4, A5], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled6[A1, A2, A3, A4, A5, A6 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6]) fp.Show[fp.Labelled6[A1, A2, A3, A4, A5, A6]] {
	return NewAppend(func(buf []string, t fp.Labelled6[A1, A2, A3, A4, A5, A6], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled7[A1, A2, A3, A4, A5, A6, A7 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7]) fp.Show[fp.Labelled7[A1, A2, A3, A4, A5, A6, A7]] {
	return NewAppend(func(buf []string, t fp.Labelled7[A1, A2, A3, A4, A5, A6, A7], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled8[A1, A2, A3, A4, A5, A6, A7, A8 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8]) fp.Show[fp.Labelled8[A1, A2, A3, A4, A5, A6, A7, A8]] {
	return NewAppend(func(buf []string, t fp.Labelled8[A1, A2, A3, A4, A5, A6, A7, A8], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled9[A1, A2, A3, A4, A5, A6, A7, A8, A9 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9]) fp.Show[fp.Labelled9[A1, A2, A3, A4, A5, A6, A7, A8, A9]] {
	return NewAppend(func(buf []string, t fp.Labelled9[A1, A2, A3, A4, A5, A6, A7, A8, A9], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10]) fp.Show[fp.Labelled10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]] {
	return NewAppend(func(buf []string, t fp.Labelled10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11]) fp.Show[fp.Labelled11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]] {
	return NewAppend(func(buf []string, t fp.Labelled11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12]) fp.Show[fp.Labelled12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]] {
	return NewAppend(func(buf []string, t fp.Labelled12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13]) fp.Show[fp.Labelled13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]] {
	return NewAppend(func(buf []string, t fp.Labelled13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14]) fp.Show[fp.Labelled14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]] {
	return NewAppend(func(buf []string, t fp.Labelled14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15]) fp.Show[fp.Labelled15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]] {
	return NewAppend(func(buf []string, t fp.Labelled15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15], ins16 fp.Show[A16]) fp.Show[fp.Labelled16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]] {
	return NewAppend(func(buf []string, t fp.Labelled16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
			ins16.Append(nil, t.I16, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15], ins16 fp.Show[A16], ins17 fp.Show[A17]) fp.Show[fp.Labelled17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17]] {
	return NewAppend(func(buf []string, t fp.Labelled17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
			ins16.Append(nil, t.I16, opt),
			ins17.Append(nil, t.I17, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15], ins16 fp.Show[A16], ins17 fp.Show[A17], ins18 fp.Show[A18]) fp.Show[fp.Labelled18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18]] {
	return NewAppend(func(buf []string, t fp.Labelled18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
			ins16.Append(nil, t.I16, opt),
			ins17.Append(nil, t.I17, opt),
			ins18.Append(nil, t.I18, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15], ins16 fp.Show[A16], ins17 fp.Show[A17], ins18 fp.Show[A18], ins19 fp.Show[A19]) fp.Show[fp.Labelled19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19]] {
	return NewAppend(func(buf []string, t fp.Labelled19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
			ins16.Append(nil, t.I16, opt),
			ins17.Append(nil, t.I17, opt),
			ins18.Append(nil, t.I18, opt),
			ins19.Append(nil, t.I19, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15], ins16 fp.Show[A16], ins17 fp.Show[A17], ins18 fp.Show[A18], ins19 fp.Show[A19], ins20 fp.Show[A20]) fp.Show[fp.Labelled20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20]] {
	return NewAppend(func(buf []string, t fp.Labelled20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
			ins16.Append(nil, t.I16, opt),
			ins17.Append(nil, t.I17, opt),
			ins18.Append(nil, t.I18, opt),
			ins19.Append(nil, t.I19, opt),
			ins20.Append(nil, t.I20, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Labelled21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 fp.Named](ins1 fp.Show[A1], ins2 fp.Show[A2], ins3 fp.Show[A3], ins4 fp.Show[A4], ins5 fp.Show[A5], ins6 fp.Show[A6], ins7 fp.Show[A7], ins8 fp.Show[A8], ins9 fp.Show[A9], ins10 fp.Show[A10], ins11 fp.Show[A11], ins12 fp.Show[A12], ins13 fp.Show[A13], ins14 fp.Show[A14], ins15 fp.Show[A15], ins16 fp.Show[A16], ins17 fp.Show[A17], ins18 fp.Show[A18], ins19 fp.Show[A19], ins20 fp.Show[A20], ins21 fp.Show[A21]) fp.Show[fp.Labelled21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21]] {
	return NewAppend(func(buf []string, t fp.Labelled21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			ins1.Append(nil, t.I1, opt),
			ins2.Append(nil, t.I2, opt),
			ins3.Append(nil, t.I3, opt),
			ins4.Append(nil, t.I4, opt),
			ins5.Append(nil, t.I5, opt),
			ins6.Append(nil, t.I6, opt),
			ins7.Append(nil, t.I7, opt),
			ins8.Append(nil, t.I8, opt),
			ins9.Append(nil, t.I9, opt),
			ins10.Append(nil, t.I10, opt),
			ins11.Append(nil, t.I11, opt),
			ins12.Append(nil, t.I12, opt),
			ins13.Append(nil, t.I13, opt),
			ins14.Append(nil, t.I14, opt),
			ins15.Append(nil, t.I15, opt),
			ins16.Append(nil, t.I16, opt),
			ins17.Append(nil, t.I17, opt),
			ins18.Append(nil, t.I18, opt),
			ins19.Append(nil, t.I19, opt),
			ins20.Append(nil, t.I20, opt),
			ins21.Append(nil, t.I21, opt),
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}
