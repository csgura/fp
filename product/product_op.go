//go:generate go run github.com/csgura/fp/internal/generator/product_gen
package product

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
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
