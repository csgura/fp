// Code generated by gombok, DO NOT EDIT.
package ngap

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

func SplitNgapType() Split[NgapType] {
	return IMap(
		Tuple4[int, string, float64](),
		func(t fp.Tuple4[int, *int, *string, *float64]) NgapType {
			return NgapType{
				Present: t.I1,
				First:   t.I2,
				Second:  t.I3,
				Third:   t.I4,
			}
		},
		func(v NgapType) fp.Tuple4[int, *int, *string, *float64] {
			return as.Tuple4(v.Present, v.First, v.Second, v.Third)
		},
	)
}

func SplitTagNgapType() SplitTag[NgapType] {
	return SplitTagContraMap(
		HConsLabelled(
			NamedInt[fp.RuntimeNamed[int]](),
			HConsLabelled(
				NamedPtr[fp.RuntimeNamed[*int], int](),
				HConsLabelled(
					NamedPtr[fp.RuntimeNamed[*string], string](),
					HConsLabelled(
						NamedPtr[fp.RuntimeNamed[*float64], float64](),
						HNil,
					),
				),
			),
		),
		fp.Compose(
			func(v NgapType) fp.Labelled4[fp.RuntimeNamed[int], fp.RuntimeNamed[*int], fp.RuntimeNamed[*string], fp.RuntimeNamed[*float64]] {
				i0, i1, i2, i3 := v.Present, v.First, v.Second, v.Third
				return as.Labelled4(as.NamedWithTag("Present", i0, ``), as.NamedWithTag("First", i1, `aper:"id=20"`), as.NamedWithTag("Second", i2, `aper:"id=30"`), as.NamedWithTag("Third", i3, `aper:"id=40"`))
			},
			as.HList4Labelled,
		),
	)
}

func SplitTagNgapValue() SplitTag[NgapValue] {
	return SplitTagContraMap(
		HConsLabelled(
			NamedInt[NamedPresent[int]](),
			HConsLabelled(
				NamedPtr[NamedFirst[*int], int](),
				HConsLabelled(
					NamedPtr[NamedSecond[*string], string](),
					HConsLabelled(
						NamedPtr[NamedThird[*float64], float64](),
						HNil,
					),
				),
			),
		),
		fp.Compose(
			NgapValue.AsLabelled,
			as.HList4Labelled,
		),
	)
}