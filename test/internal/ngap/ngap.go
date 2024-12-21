package ngap

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/seq"
)

type Derives[T any] interface {
	Target() T
}

type NgapType struct {
	Present int
	First   *int
	Second  *string
	Third   *float64
}

type Split[T any] interface {
	Split(t T) []T
}

type SplitFunc[T any] func(t T) []T

func (r SplitFunc[T]) Split(t T) []T {
	return r(t)
}

func New[T any](f SplitFunc[T]) Split[T] {
	return f
}

func IMap[A, B any](instance Split[A], fab func(A) B, fba func(B) A) Split[B] {
	return New(func(t B) []B {
		return seq.Map(instance.Split(fba(t)), fab)
	})
}

func Tuple4[A1, A2, A3 any]() Split[fp.Tuple4[int, *A1, *A2, *A3]] {
	type tp4 = fp.Tuple4[int, *A1, *A2, *A3]
	return New(func(t tp4) []tp4 {
		var ret []tp4
		if t.I2 != nil {
			ret = append(ret, as.Tuple4[int, *A1, *A2, *A3](1, t.I2, nil, nil))
		}

		if t.I3 != nil {
			ret = append(ret, as.Tuple4[int, *A1, *A2, *A3](2, nil, t.I3, nil))
		}

		if t.I4 != nil {
			ret = append(ret, as.Tuple4[int, *A1, *A2, *A3](3, nil, nil, t.I4))
		}
		return ret
	})
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive
var _ Derives[Split[NgapType]]
