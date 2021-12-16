package product

import (
	"github.com/csgura/fp"
)

func Tuple2[A1, A2 any](a1 A1, a2 A2) fp.Tuple2[A1, A2] {
	return fp.Tuple2[A1, A2]{
		I1: a1,
		I2: a2,
	}
}
func Tuple3[A1, A2, A3 any](a1 A1, a2 A2, a3 A3) fp.Tuple3[A1, A2, A3] {
	return fp.Tuple3[A1, A2, A3]{
		I1: a1,
		I2: a2,
		I3: a3,
	}
}
func Tuple4[A1, A2, A3, A4 any](a1 A1, a2 A2, a3 A3, a4 A4) fp.Tuple4[A1, A2, A3, A4] {
	return fp.Tuple4[A1, A2, A3, A4]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
	}
}
func Tuple5[A1, A2, A3, A4, A5 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Tuple5[A1, A2, A3, A4, A5] {
	return fp.Tuple5[A1, A2, A3, A4, A5]{
		I1: a1,
		I2: a2,
		I3: a3,
		I4: a4,
		I5: a5,
	}
}
