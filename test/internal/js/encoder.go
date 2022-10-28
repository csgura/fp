package js

import (
	"fmt"
	"strings"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
)

type Encoder[T any] interface {
	Encode(T) string
}

type EncoderFunc[T any] func(T) string

func (r EncoderFunc[T]) Encode(t T) string {
	return r(t)
}

type Derives[T any] interface {
}

func NewEncoder[T any](f func(a T) string) Encoder[T] {
	return EncoderFunc[T](f)
}

var EncoderString = NewEncoder(func(a string) string {
	return fmt.Sprintf(`"%s"`, a)
})

func EncoderNumber[T fp.ImplicitNum]() Encoder[T] {
	return NewEncoder(func(a T) string {
		return fmt.Sprintf("%v", a)
	})
}

var EncoderTime = NewEncoder(func(a time.Time) string {
	return EncoderString.Encode(a.Format(time.RFC3339))
})

var EncoderUnit = NewEncoder(func(a fp.Unit) string {
	return "null"
})

var EncoderHNil Encoder[hlist.Nil] = NewEncoder(func(a hlist.Nil) string {
	return ""
})

func EncoderHCons[H any, T hlist.HList](heq Encoder[H], teq Encoder[T]) Encoder[hlist.Cons[H, T]] {
	return NewEncoder(func(a hlist.Cons[H, T]) string {
		return heq.Encode(a.Head()) + "," + teq.Encode(a.Tail())
	})
}

func EncoderHConsLabelled[H fp.Named, T hlist.HList](heq Encoder[H], teq Encoder[T]) Encoder[hlist.Cons[H, T]] {
	return NewEncoder(func(a hlist.Cons[H, T]) string {
		if a.Tail().IsNil() {
			return fmt.Sprintf(`{"%s":%s}`, a.Head().Name(), heq.Encode(a.Head()))
		}
		return fmt.Sprintf(`{"%s":%s,%s}`, a.Head().Name(), heq.Encode(a.Head()),
			strings.Trim(teq.Encode(a.Tail()), "{}"),
		)
	})
}

// func IMap[A, B any](instance Encoder[A], fab func(A) B, fba func(B) A) Encoder[B] {
// 	return New(func(a B) string {
// 		return instance.Encode(fba(a))
// 	})
// }

func EncoderContraMap[T, U any](instance Encoder[T], fn func(U) T) Encoder[U] {
	return NewEncoder(func(a U) string {
		return instance.Encode(fn(a))
	})
}

func EncoderLabelled1[A fp.Named](ins1 Encoder[A]) Encoder[fp.Labelled1[A]] {
	return NewEncoder(
		func(a fp.Labelled1[A]) string {
			return fmt.Sprintf(`{"%s" : %s}`, a.I1.Name(), ins1.Encode(a.I1))
		},
	)
}

func EncoderNamed[T fp.NamedField[A], A any](enc Encoder[A]) Encoder[T] {
	return NewEncoder(func(a T) string {
		return enc.Encode(a.Value())
	})
}

func EncoderLabelled2[A1, A2 fp.Named](ins1 Encoder[A1], ins2 Encoder[A2]) Encoder[fp.Labelled2[A1, A2]] {

	return NewEncoder(
		func(a fp.Labelled2[A1, A2]) string {
			return fmt.Sprintf(`{"%s":%s,"%s":%s}`,
				a.I1.Name(), ins1.Encode(a.I1),
				a.I2.Name(), ins2.Encode(a.I2),
			)
		},
	)
}
