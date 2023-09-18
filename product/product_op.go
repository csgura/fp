//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package product

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
)

func Tuple2[A1, A2 any](a1 A1, a2 A2) fp.Tuple2[A1, A2] {
	return fp.Tuple2[A1, A2]{
		I1: a1,
		I2: a2,
	}
}

func FromHNil(hlist.Nil) fp.Unit {
	return fp.Unit{}
}

func TupleFromHList1[A1 any](list hlist.Cons[A1, hlist.Nil]) fp.Tuple1[A1] {
	return fp.Tuple1[A1]{
		I1: list.Head(),
	}
}

func LabelledFromHList1[A1 fp.Named](list hlist.Cons[A1, hlist.Nil]) fp.Labelled1[A1] {
	return as.Labelled1(list.Head())
}

func MapKey[K, V, R any](t fp.Tuple2[K, V], mapf func(K) R) fp.Tuple2[R, V] {
	return as.Tuple2(mapf(t.I1), t.I2)
}

func MapValue[K, V, R any](t fp.Tuple2[K, V], mapf func(V) R) fp.Tuple2[K, R] {
	return as.Tuple2(t.I1, mapf(t.I2))
}

func LiftKey[K, V, R any](mapf func(K, V) R) fp.Func1[fp.Tuple2[K, V], fp.Tuple2[R, V]] {
	return func(a1 fp.Tuple2[K, V]) fp.Tuple2[R, V] {
		return as.Tuple2(mapf(a1.I1, a1.I2), a1.I2)
	}
}

func LiftValue[K, V, R any](mapf func(K, V) R) fp.Func1[fp.Tuple2[K, V], fp.Tuple2[K, R]] {
	return func(a1 fp.Tuple2[K, V]) fp.Tuple2[K, R] {
		return as.Tuple2(a1.I1, mapf(a1.I1, a1.I2))
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  3,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]({{DeclArgs 1 .N}}) fp.{{TupleType .N}} {
	return fp.{{TupleType .N}}{
		{{- range $idx := Range 1 .N}}
			I{{$idx}}: a{{$idx}},
		{{- end}}
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func TupleFromHList{{.N}}[{{TypeArgs 1 .N}} any](list {{ConsType 1 .N "hlist.Nil"}}) fp.{{TupleType .N}} {
	tail := TupleFromHList{{dec .N}}(list.Tail())
	return Tuple{{.N}}(list.Head(), {{CallArgs 1 (dec .N) "tail.I"}})
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func LabelledFromHList{{.N}}[{{TypeArgs 1 .N}} fp.Named](list {{ConsType 1 .N "hlist.Nil"}}) fp.Labelled{{.N}}[{{TypeArgs 1 .N}}] {
	tail := LabelledFromHList{{dec .N}}(list.Tail())
	return as.Labelled{{.N}}(list.Head(), {{CallArgs 1 (dec .N) "tail.I"}})
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func Lift{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) func(fp.{{TupleType .N}}) R {
	return func(t fp.Tuple{{.N}}[{{TypeArgs 1 .N}}]) R {
		return f(t.Unapply())
	}
}
	`,
}
