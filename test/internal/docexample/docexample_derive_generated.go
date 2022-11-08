// Code generated by gombok, DO NOT EDIT.
package docexample

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/show"
)

var EqPerson = eq.ContraMap(
	eq.Tuple2(eq.String, eq.Given[int]()),
	Person.AsTuple,
)

var HashablePerson = hash.ContraMap(
	hash.Tuple2(hash.String, hash.Number[int]()),
	Person.AsTuple,
)

var EqCar = eq.ContraMap(
	eq.Tuple3(eq.String, eq.String, eq.Given[int]()),
	Car.AsTuple,
)

var OrdCar = ord.ContraMap(
	ord.Tuple3(ord.Given[string](), ord.Given[string](), ord.Given[int]()),
	Car.AsTuple,
)

var EqCarsOwned = eq.ContraMap(
	eq.Tuple2(EqPerson, eq.Seq(EqCar)),
	CarsOwned.AsTuple,
)

var ShowAddress = show.Generic(
	as.Generic(
		"docexample.Address",
		fp.Compose(
			Address.AsTuple,
			as.HList3[string, string, string],
		),

		fp.Compose(
			as.Func2(
				hlist.Case3[string, string, string, hlist.Nil, fp.Tuple3[string, string, string]],
			).ApplyLast(
				as.Tuple3[string, string, string],
			),
			fp.Compose(
				as.Curried2(AddressBuilder.FromTuple)(AddressBuilder{}),
				AddressBuilder.Build,
			),
		),
	),
	show.HCons(
		show.String,
		show.HCons(
			show.String,
			show.HCons(
				show.String,
				show.HNil,
			),
		),
	),
)

var EncoderCar = js.EncoderContraMap(
	js.EncoderHConsLabelled(
		js.EncoderNamed[NameIsCompany[string]](js.EncoderString),
		js.EncoderHConsLabelled(
			js.EncoderNamed[NameIsModel[string]](js.EncoderString),
			js.EncoderHConsLabelled(
				js.EncoderNamed[NameIsYear[int]](js.EncoderNumber[int]()),
				js.EncoderHNil,
			),
		),
	),
	fp.Compose(
		Car.AsLabelled,
		as.HList3Labelled[NameIsCompany[string], NameIsModel[string], NameIsYear[int]],
	),
)
