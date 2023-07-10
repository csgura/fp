// Code generated by gombok, DO NOT EDIT.
package recursive

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/test/internal/show"
)

var ShowNormalStruct = show.Generic(
	as.Generic(
		"recursive.NormalStruct",
		fp.Compose(
			func(v NormalStruct) fp.Tuple3[string, int, string] {
				return as.Tuple3(v.Name, v.Age, v.Address)
			},
			as.HList3[string, int, string],
		),

		fp.Compose(
			product.TupleFromHList3[string, int, string],
			func(t fp.Tuple3[string, int, string]) NormalStruct {
				return NormalStruct{
					Name:    t.I1,
					Age:     t.I2,
					Address: t.I3,
				}
			},
		),
	),
	show.HCons(
		show.String,
		show.HCons(
			show.Int[int](),
			show.HCons(
				show.String,
				show.HNil,
			),
		),
	),
)
