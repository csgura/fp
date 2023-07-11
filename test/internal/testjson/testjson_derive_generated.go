// Code generated by gombok, DO NOT EDIT.
package testjson

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/test/internal/js"
)

var EncoderRoot = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NamedA[int]](js.EncoderNumber[int]()),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedB[string]](js.EncoderString),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NamedC[float64]](js.EncoderNumber[float64]()),
				js.EncoderHConsLabelled(
					js.EncoderNamed[NamedD[bool]](js.EncoderBool),
					js.EncoderHConsLabelled(
						js.EncoderNamed[NamedE[*int]](js.EncoderPtr(lazy.Call(func() js.Encoder[int] {
							return js.EncoderNumber[int]()
						}))),
						js.EncoderHConsLabelled(
							js.EncoderNamed[NamedF[[]int]](js.EncoderSlice(js.EncoderNumber[int]())),
							js.EncoderHConsLabelled(
								js.EncoderNamed[NamedG[map[string]int]](js.EncoderGoMap(js.EncoderNumber[int]())),
								js.EncoderHConsLabelled(
									js.EncoderNamed[NamedH[Child]](EncoderChild),
									js.EncoderHNil,
								),
							),
						),
					),
				),
			),
		),
	),
	fp.Compose(
		Root.AsLabelled,
		as.HList8Labelled[NamedA[int], NamedB[string], NamedC[float64], NamedD[bool], NamedE[*int], NamedF[[]int], NamedG[map[string]int], NamedH[Child]],
	),
)

var DecoderRoot = js.DecoderMap(
	js.DecoderHConsLabelled(
		js.DecoderNamed[NamedA[int]](js.DecoderNumber[int]()),
		js.DecoderHConsLabelled(
			js.DecoderNamed[NamedB[string]](js.DecoderString),
			js.DecoderHConsLabelled(
				js.DecoderNamed[NamedC[float64]](js.DecoderNumber[float64]()),
				js.DecoderHConsLabelled(
					js.DecoderNamed[NamedD[bool]](js.DecoderBool),
					js.DecoderHConsLabelled(
						js.DecoderNamed[NamedE[*int]](js.DecoderPtr(lazy.Call(func() js.Decoder[int] {
							return js.DecoderNumber[int]()
						}))),
						js.DecoderHConsLabelled(
							js.DecoderNamed[NamedF[[]int]](js.DecoderSlice(js.DecoderNumber[int]())),
							js.DecoderHConsLabelled(
								js.DecoderNamed[NamedG[map[string]int]](js.DecoderGoMap(js.DecoderNumber[int]())),
								js.DecoderHConsLabelled(
									js.DecoderNamed[NamedH[Child]](DecoderChild),
									js.DecoderHNil,
								),
							),
						),
					),
				),
			),
		),
	),

	fp.Compose(
		product.LabelledFromHList8[NamedA[int], NamedB[string], NamedC[float64], NamedD[bool], NamedE[*int], NamedF[[]int], NamedG[map[string]int], NamedH[Child]],
		fp.Compose(
			as.Curried2(RootBuilder.FromLabelled)(RootBuilder{}),
			RootBuilder.Build,
		),
	),
)

var EncoderChild = js.EncoderContraMap(
	js.EncoderLabelled2(js.EncoderNamed[NamedA[map[string]any]](js.EncoderGoMapAny), js.EncoderNamed[NamedB[any]](js.EncoderGiven[any]())),
	Child.AsLabelled,
)

var DecoderChild = js.DecoderMap(
	js.DecoderLabelled2(js.DecoderNamed[NamedA[map[string]any]](js.DecoderGoMapAny), js.DecoderNamed[NamedB[any]](js.DecoderGiven[any]())),
	fp.Compose(
		as.Curried2(ChildBuilder.FromLabelled)(ChildBuilder{}),
		ChildBuilder.Build,
	),
)

func EncoderNode() js.Encoder[Node] {
	return js.EncoderContraMap(
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedName[string]](js.EncoderString),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NamedLeft[*Node]](js.EncoderPtr(lazy.Call(func() js.Encoder[Node] {
					return EncoderNode()
				}))),
				js.EncoderHConsLabelled(
					js.EncoderNamed[NamedRight[*Node]](js.EncoderPtr(lazy.Call(func() js.Encoder[Node] {
						return EncoderNode()
					}))),
					js.EncoderHNil,
				),
			),
		),
		fp.Compose(
			Node.AsLabelled,
			as.HList3Labelled[NamedName[string], NamedLeft[*Node], NamedRight[*Node]],
		),
	)
}

var EncoderTree = js.EncoderContraMap(
	js.EncoderLabelled1(js.EncoderNamed[NamedRoot[*Node]](js.EncoderPtr(lazy.Call(func() js.Encoder[Node] {
		return EncoderNode()
	})))),
	Tree.AsLabelled,
)

func EncoderEntry[V any](encoderV js.Encoder[V]) js.Encoder[Entry[V]] {
	return js.EncoderContraMap(
		js.EncoderLabelled2(js.EncoderNamed[NamedName[string]](js.EncoderString), js.EncoderNamed[NamedValue[V]](encoderV)),
		Entry[V].AsLabelled,
	)
}

func EncoderNotUsedParam[K any, V any](encoderV js.Encoder[V]) js.Encoder[NotUsedParam[K, V]] {
	return js.EncoderContraMap(
		js.EncoderLabelled2(js.EncoderNamed[NamedParam[string]](js.EncoderString), js.EncoderNamed[NamedValue[V]](encoderV)),
		NotUsedParam[K, V].AsLabelled,
	)
}

var EncoderMovie = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NamedName[string]](js.EncoderString),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NamedCasting[Entry[string]]](EncoderEntry(js.EncoderString)),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NamedNotUsed[NotUsedParam[int, string]]](EncoderNotUsedParam[int, string](js.EncoderString)),
				js.EncoderHNil,
			),
		),
	),
	fp.Compose(
		Movie.AsLabelled,
		as.HList3Labelled[NamedName[string], NamedCasting[Entry[string]], NamedNotUsed[NotUsedParam[int, string]]],
	),
)

var EncoderNoPrivate = js.EncoderContraMap(
	js.EncoderLabelled1(js.EncoderNamed[PubNamedRoot[string]](js.EncoderString)),
	NoPrivate.AsLabelled,
)

var DecoderNoPrivate = js.DecoderMap(
	js.DecoderHConsLabelled(
		js.DecoderNamed[PubNamedRoot[string]](js.DecoderString),
		js.DecoderHNil,
	),

	fp.Compose(
		product.LabelledFromHList1[PubNamedRoot[string]],
		fp.Compose(
			as.Curried2(NoPrivateBuilder.FromLabelled)(NoPrivateBuilder{}),
			NoPrivateBuilder.Build,
		),
	),
)
