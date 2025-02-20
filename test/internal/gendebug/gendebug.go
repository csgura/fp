package gendebug

import (
	"github.com/csgura/fp/minimal"
	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/test/internal/namedptr"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Value struct {
	Present int
	Value   int
}

type Container struct {
	v Value
}

// func ValidatorStruct1[A any](as mshow.Show[A]) namedptr.Validator[minimal.Tuple1[A]] {
// 	return namedptr.New(func(t minimal.Tuple1[A]) error {
// 		return nil
// 	})
// }

func ShowStruct1[A any](as namedptr.Validator[A]) mshow.Show[minimal.Tuple1[A]] {
	panic("")
}

// @fp.Derive(recusive=true)
var _ mshow.Derives[mshow.Show[Container]]

// @fp.Derive
var _ namedptr.Derives[namedptr.Validator[Value]]

// // @fp.Derive(recusive=true)
// var _ namedptr.Derives[namedptr.Validator[Container]]
