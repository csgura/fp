// Code generated by gombok, DO NOT EDIT.
package showorder

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/show"
)

func ShowHello() fp.Show[Hello] {
	return show.Generic(
		as.Generic(
			"showorder.Hello",
			"Struct",
			func(v Hello) fp.Labelled4[fp.RuntimeNamed[int], fp.RuntimeNamed[string], fp.RuntimeNamed[any], fp.RuntimeNamed[World]] {
				i0, i1, i2, i3 := v.A, v.B, v.AN, v.World
				return as.Labelled4(as.NamedWithTag("A", i0, ``), as.NamedWithTag("B", i1, ``), as.NamedWithTag("AN", i2, ``), as.NamedWithTag("World", i3, ``))
			},
			func(t fp.Labelled4[fp.RuntimeNamed[int], fp.RuntimeNamed[string], fp.RuntimeNamed[any], fp.RuntimeNamed[World]]) Hello {
				return Hello{A: t.I1.Value(), B: t.I2.Value(), AN: t.I3.Value(), World: t.I4.Value()}
			},
		),
		show.Labelled4(show.Named[fp.RuntimeNamed[int], int](show.Int[int]()), show.Named[fp.RuntimeNamed[string], string](show.String), show.Named[fp.RuntimeNamed[any], any](ShowAny[any]()), show.Named[fp.RuntimeNamed[World], World](ShowWorld())),
	)
}

func ShowWorld() fp.Show[World] {
	return show.Generic(
		as.Generic(
			"showorder.World",
			"Struct",
			fp.Compose(
				func(v World) fp.Labelled1[fp.RuntimeNamed[string]] {
					i0 := v.Loc
					return as.Labelled1(as.NamedWithTag("Loc", i0, ``))
				},
				as.HList1Labelled,
			),

			fp.Compose(
				product.LabelledFromHList1,
				func(t fp.Labelled1[fp.RuntimeNamed[string]]) World {
					return World{Loc: t.I1.Value()}
				},
			),
		),
		show.HConsLabelled(
			show.Named[fp.RuntimeNamed[string], string](show.String),
			show.HNil,
		),
	)
}
