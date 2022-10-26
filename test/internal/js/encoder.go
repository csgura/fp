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

func New[T any](f func(a T) string) Encoder[T] {
	return EncoderFunc[T](f)
}

var String = New(func(a string) string {
	return fmt.Sprintf(`"%s"`, a)
})

func Number[T fp.ImplicitNum]() Encoder[T] {
	return New(func(a T) string {
		return fmt.Sprintf("%v", a)
	})
}

var Time = New(func(a time.Time) string {
	return String.Encode(a.Format(time.RFC3339))
})

var Unit = New(func(a fp.Unit) string {
	return "null"
})

var HNil Encoder[hlist.Nil] = New(func(a hlist.Nil) string {
	return ""
})

func HCons[H any, T hlist.HList](heq Encoder[H], teq Encoder[T]) Encoder[hlist.Cons[H, T]] {
	return New(func(a hlist.Cons[H, T]) string {
		return heq.Encode(a.Head()) + "," + teq.Encode(a.Tail())
	})
}

func HConsLabelled[H any, T hlist.HList](heq Encoder[H], teq Encoder[T]) Encoder[hlist.Cons[fp.Field[H], T]] {
	return New(func(a hlist.Cons[fp.Field[H], T]) string {
		if a.Tail().IsNil() {
			return fmt.Sprintf(`{"%s":%s}`, a.Head().Name, heq.Encode(a.Head().Value))
		}
		return fmt.Sprintf(`{"%s":%s,%s}`, a.Head().Name, heq.Encode(a.Head().Value),
			strings.Trim(teq.Encode(a.Tail()), "{}"),
		)
	})
}

// func IMap[A, B any](instance Encoder[A], fab func(A) B, fba func(B) A) Encoder[B] {
// 	return New(func(a B) string {
// 		return instance.Encode(fba(a))
// 	})
// }

func ContraMap[T, U any](instance Encoder[T], fn func(U) T) Encoder[U] {
	return New(func(a U) string {
		return instance.Encode(fn(a))
	})
}

func Labelled1[A any](ins1 Encoder[A]) Encoder[fp.Labelled1[A]] {
	return New(
		func(a fp.Labelled1[A]) string {
			return fmt.Sprintf(`{"%s" : %s}`, a.I1.Name, ins1.Encode(a.I1.Value))
		},
	)
}

func Labelled2[A1, A2 any](ins1 Encoder[A1], ins2 Encoder[A2]) Encoder[fp.Labelled2[A1, A2]] {

	return New(
		func(a fp.Labelled2[A1, A2]) string {
			return fmt.Sprintf(`{"%s":%s,"%s":%s}`,
				a.I1.Name, ins1.Encode(a.I1.Value),
				a.I2.Name, ins2.Encode(a.I2.Value),
			)
		},
	)
}
