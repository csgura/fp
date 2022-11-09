package js

import (
	"fmt"
	"strings"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

type Encoder[T any] interface {
	Encode(T) fp.Option[string]
}

type EncoderFunc[T any] func(T) fp.Option[string]

func (r EncoderFunc[T]) Encode(t T) fp.Option[string] {
	return r(t)
}

type Derives[T any] interface {
}

func NewEncoder[T any](f func(a T) fp.Option[string]) Encoder[T] {
	return EncoderFunc[T](f)
}

var EncoderString = NewEncoder(func(a string) fp.Option[string] {
	if a != "" {
		return option.Some(fmt.Sprintf(`"%s"`, a))
	}
	return option.None[string]()
})

func EncoderNumber[T fp.ImplicitNum]() Encoder[T] {
	return NewEncoder(func(a T) fp.Option[string] {
		return option.Some(fmt.Sprintf("%v", a))
	})
}

var EncoderTime = NewEncoder(func(a time.Time) fp.Option[string] {
	return EncoderString.Encode(a.Format(time.RFC3339))
})

var EncoderUnit = NewEncoder(func(a fp.Unit) fp.Option[string] {
	return option.None[string]()
})

var EncoderHNil Encoder[hlist.Nil] = NewEncoder(func(a hlist.Nil) fp.Option[string] {
	return option.None[string]()
})

func EncoderSeq[T any](enc Encoder[T]) Encoder[fp.Seq[T]] {
	return NewEncoder(func(s fp.Seq[T]) fp.Option[string] {
		if len(s) == 0 {
			return option.None[string]()
		}
		return option.Some("[" + seq.Map(s, func(v T) string {
			return enc.Encode(v).OrElse("null")
		}).MakeString(",") + "]")
	})
}

var EncoderBool = NewEncoder(func(a bool) fp.Option[string] {
	if a {
		return option.Some("true")
	}
	return option.Some("false")
})

func EncoderPtr[T any](encT Encoder[T]) Encoder[*T] {
	return NewEncoder(func(a *T) fp.Option[string] {
		if a != nil {
			return encT.Encode(*a)
		}
		return option.Some("null")
	})
}

func EncoderSlice[T any](enc Encoder[T]) Encoder[[]T] {
	return NewEncoder(func(s []T) fp.Option[string] {
		if len(s) == 0 {
			return option.None[string]()
		}
		return option.Some("[" + seq.Map(s, func(v T) string {
			return enc.Encode(v).OrElse("null")
		}).MakeString(",") + "]")
	})
}
func EncoderOption[T any](enc Encoder[T]) Encoder[fp.Option[T]] {
	return NewEncoder(func(opt fp.Option[T]) fp.Option[string] {
		if opt.IsDefined() {
			return enc.Encode(opt.Get())

		}
		return option.None[string]()
	})
}

func EncoderNamed[T fp.NamedField[A], A any](enc Encoder[A]) Encoder[T] {
	return NewEncoder(func(a T) fp.Option[string] {

		return option.Map(enc.Encode(a.Value()), func(v string) string {
			return fmt.Sprintf(`"%s":%s`, a.Name(), v)
		})
	})
}

func EncoderHConsLabelled[H fp.Named, T hlist.HList](heq Encoder[H], teq Encoder[T]) Encoder[hlist.Cons[H, T]] {
	return NewEncoder(func(a hlist.Cons[H, T]) fp.Option[string] {

		head := heq.Encode(a.Head())
		tail := teq.Encode(a.Tail())

		if head.IsDefined() && tail.IsDefined() {
			return option.Some(fmt.Sprintf(`{%s,%s}`, head.Get(),
				strings.Trim(tail.Get(), "{}")))
		}

		if head.IsDefined() {
			return option.Some(fmt.Sprintf("{%s}", head.Get()))
		}

		return option.Map(tail, func(v string) string {
			return fmt.Sprintf("{%s}", v)
		})
	})
}

func EncoderGoMap[V any](encV Encoder[V]) Encoder[map[string]V] {
	return NewEncoder(func(a map[string]V) fp.Option[string] {
		list := fp.Seq[string]{}
		for k, v := range a {
			vstr := encV.Encode(v)
			if vstr.IsDefined() {
				list = list.Append(fmt.Sprintf(`"%s":%s`, k, vstr.Get()))
			}
		}

		if list.Size() > 0 {
			return option.Some("{" + list.MakeString(",") + "}")
		}
		return option.None[string]()
	})
}

// func IMap[A, B any](instance Encoder[A], fab func(A) B, fba func(B) A) Encoder[B] {
// 	return New(func(a B) string {
// 		return instance.Encode(fba(a))
// 	})
// }

func EncoderContraMap[T, U any](instance Encoder[T], fn func(U) T) Encoder[U] {
	return NewEncoder(func(a U) fp.Option[string] {
		return instance.Encode(fn(a))
	})
}

func EncoderLabelled1[A fp.Named](ins1 Encoder[A]) Encoder[fp.Labelled1[A]] {
	return NewEncoder(
		func(a fp.Labelled1[A]) fp.Option[string] {
			return option.Map(ins1.Encode(a.I1), func(v string) string {
				return fmt.Sprintf(`{"%s"}`, v)
			})

		},
	)
}

func EncoderLabelled2[A1, A2 fp.Named](ins1 Encoder[A1], ins2 Encoder[A2]) Encoder[fp.Labelled2[A1, A2]] {

	return NewEncoder(
		func(a fp.Labelled2[A1, A2]) fp.Option[string] {
			i1 := ins1.Encode(a.I1)
			i2 := ins2.Encode(a.I2)

			if i1.IsDefined() && i2.IsDefined() {
				return option.Some(fmt.Sprintf(`{%s,%s}`, i1.Get(), i2.Get()))
			}

			if i1.IsDefined() {
				return option.Some(fmt.Sprintf("{%s}", i1.Get()))
			}

			return option.Map(i2, func(v string) string {
				return fmt.Sprintf("{%s}", v)
			})
		},
	)
}
