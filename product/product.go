package product

import "github.com/csgura/fp"

func Tuple1[A any](a A) fp.Tuple1[A] {
	return fp.Tuple1[A]{I1: a}
}

func Tuple2[A, B any](a A, b B) fp.Tuple2[A, B] {
	return fp.Tuple2[A, B]{I1: a, I2: b}
}

func Tuple3[A, B, C any](a A, b B, c C) fp.Tuple3[A, B, C] {
	return fp.Tuple3[A, B, C]{I1: a, I2: b, I3: c}
}
