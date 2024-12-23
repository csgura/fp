package ngap

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/seq"
)

type Derives[T any] interface {
	Target() T
}

type NgapType struct {
	Present int
	First   *int     `aper:"id=20"`
	Second  *string  `aper:"id=30"`
	Third   *float64 `aper:"id=40"`
}

// @fp.Value
// @fp.GenLabelled
type NgapValue struct {
	present int
	first   *int     `aper:"id=20"`
	second  *string  `aper:"id=30"`
	third   *float64 `aper:"id=40"`
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

type SplitTag[T any] interface {
	SplitTag(t T) []string
}

type SplitTagFunc[T any] func(t T) []string

func (r SplitTagFunc[T]) SplitTag(t T) []string {
	return r(t)
}

func NewSplitTag[T any](f SplitTagFunc[T]) SplitTag[T] {
	return f
}

func IMap[A, B any](instance Split[A], fab func(A) B, fba func(B) A) Split[B] {
	return New(func(t B) []B {
		return seq.Map(instance.Split(fba(t)), fab)
	})
}

func SplitTagContraMap[A, B any](instance SplitTag[A], fba func(B) A) SplitTag[B] {
	return NewSplitTag(func(t B) []string {
		return instance.SplitTag(fba(t))
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

var HNil = NewSplitTag(func(v hlist.Nil) []string {
	return nil
})

func HConsLabelled[H fp.Named, T hlist.HList](hshow SplitTag[H], tshow SplitTag[T]) SplitTag[hlist.Cons[H, T]] {
	return NewSplitTag(func(t hlist.Cons[H, T]) []string {
		tlist := tshow.SplitTag(hlist.Tail(t))
		h := hshow.SplitTag(t.Head())
		return append(h, tlist...)
	})
}

func Named[T fp.NamedField[*A], A any]() SplitTag[T] {
	return NewSplitTag(func(t T) []string {
		if t.Value() != nil {
			return seq.Of(t.Tag())
		}
		return nil
	})
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive
var _ Derives[Split[NgapType]]

// @fp.Derive
var _ Derives[SplitTag[NgapType]]

// @fp.Derive
var _ Derives[SplitTag[NgapValue]]
