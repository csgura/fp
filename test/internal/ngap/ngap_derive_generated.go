// Code generated by gombok, DO NOT EDIT.
package ngap

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

func SplitNgapType() Split[NgapType] {
	return IMap(
		Tuple4[int, string,float64](),
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
