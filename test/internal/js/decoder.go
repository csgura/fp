package js

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type DecoderContext struct {
	workingObject fp.Option[map[string]json.RawMessage]
}

type Decoder[T any] interface {
	Decode(DecoderContext, string) fp.Try[T]
}

type DecoderFunc[T any] func(DecoderContext, string) fp.Try[T]

func (r DecoderFunc[T]) Decode(ctx DecoderContext, t string) fp.Try[T] {
	return r(ctx, t)
}

func NewDecoder[T any](f func(ctx DecoderContext, a string) fp.Try[T]) Decoder[T] {
	return DecoderFunc[T](f)
}

var DecoderString = NewDecoder(func(ctx DecoderContext, a string) fp.Try[string] {
	if a[0] == '"' {
		return try.Success(a[1 : len(a)-1])
	}
	return try.Failure[string](fmt.Errorf("invalid string literal"))
})

var DecoderBool = NewDecoder(func(ctx DecoderContext, a string) fp.Try[bool] {
	if a == "true" {
		return try.Success(true)
	} else if a == "false" {
		return try.Success(false)

	}
	return try.Failure[bool](fmt.Errorf("invalid boolean literal"))
})

func DecoderNumber[T fp.ImplicitNum]() Decoder[T] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[T] {
		r, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return try.Failure[T](err)
		}
		return try.Success(T(r))
	})
}

var DecoderTime = NewDecoder(func(ctx DecoderContext, a string) fp.Try[time.Time] {
	return try.FlatMap(DecoderString.Decode(ctx, a), func(v string) fp.Try[time.Time] {
		return try.Apply(time.Parse(time.RFC3339, v))
	})

})

var DecoderUnit = NewDecoder(func(ctx DecoderContext, a string) fp.Try[fp.Unit] {
	return try.Success(fp.Unit{})
})

var DecoderHNil Decoder[hlist.Nil] = NewDecoder(func(ctx DecoderContext, a string) fp.Try[hlist.Nil] {
	return try.Success(hlist.Empty())
})

func DecoderSlice[T any](decT Decoder[T]) Decoder[[]T] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[[]T] {
		var l []json.RawMessage
		err := json.Unmarshal([]byte(a), &l)
		if err != nil {
			return try.Failure[[]T](err)
		}

		return try.TraverseSeq(as.Seq(l), func(v json.RawMessage) fp.Try[T] {
			return decT.Decode(ctx, string(v))
		})
	})
}

func DecoderGoMap[T any](decT Decoder[T]) Decoder[map[string]T] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[map[string]T] {
		var l map[string]json.RawMessage
		err := json.Unmarshal([]byte(a), &l)
		if err != nil {
			return try.Failure[map[string]T](err)
		}

		ret := map[string]T{}
		for k, v := range l {
			tv := decT.Decode(ctx, string(v))
			if tv.IsFailure() {
				return try.Failure[map[string]T](tv.Failed().Get())
			}
			ret[k] = tv.Get()
		}
		return try.Success(ret)
	})
}

var DecoderGoMapAny = NewDecoder(func(ctx DecoderContext, a string) fp.Try[map[string]any] {
	var ret map[string]any
	err := json.Unmarshal([]byte(a), &ret)
	if err != nil {
		return try.Failure[map[string]any](err)
	}
	return try.Success(ret)
})

func DecoderGiven[T any]() Decoder[T] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[T] {
		var ret T
		err := json.Unmarshal([]byte(a), &ret)
		if err != nil {
			return try.Failure[T](err)
		}
		return try.Success(ret)
	})
}

func DecoderPtr[T any](decT lazy.Eval[Decoder[T]]) Decoder[*T] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[*T] {
		if a == "null" {
			return try.Success[*T](nil)
		}
		ret := decT.Get().Decode(ctx, a)
		return try.Map(ret, func(v T) *T {
			return &v
		})
	})
}

func DecoderHConsLabelled[H fp.Named, T hlist.HList](heq Decoder[H], teq Decoder[T]) Decoder[hlist.Cons[H, T]] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[hlist.Cons[H, T]] {

		var m map[string]json.RawMessage
		if ctx.workingObject.IsDefined() {
			m = ctx.workingObject.Get()
		} else {
			err := json.Unmarshal([]byte(a), &m)
			if err != nil {
				return try.Failure[hlist.Cons[H, T]](err)
			}
		}

		var h H
		toDecode := option.Map(option.Of(m[h.Name()]), func(v json.RawMessage) string {
			return string(v)
		}).OrElse("null")

		head := heq.Decode(ctx.WithNoneWorkingObject(), toDecode)

		return try.FlatMap(head, func(h H) fp.Try[hlist.Cons[H, T]] {
			return try.Map(teq.Decode(ctx.WithSomeWorkingObject(m), a), func(t T) hlist.Cons[H, T] {
				return hlist.Concat(h, t)
			})
		})
	})
}

// func IMap[A, B any](instance Decoder[A], fab func(A) B, fba func(B) A) Decoder[B] {
// 	return New(func(a B) string {
// 		return instance.Decode(fba(a))
// 	})
// }

func DecoderMap[T, U any](instance Decoder[T], fn func(T) U) Decoder[U] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[U] {
		return try.Map(instance.Decode(ctx, a), fn)
	})
}

// func DecoderLabelled1[A fp.Named](ins1 Decoder[A]) Decoder[fp.Labelled1[A]] {
// 	return NewDecoder(
// 		func(a fp.Labelled1[A]) string {
// 			return fmt.Sprintf(`{"%s" : %s}`, a.I1.Name(), ins1.Decode(a.I1))
// 		},
// 	)
// }

func DecoderNamed[T interface {
	fp.NamedField[A]
	WithValue(A) T
}, A any](enc Decoder[A]) Decoder[T] {
	return NewDecoder(func(ctx DecoderContext, a string) fp.Try[T] {
		ret := enc.Decode(ctx, a)
		return try.Map(ret, func(v A) T {
			var zero T
			return zero.WithValue(v)
		})
	})
}

func DecoderLabelled2[N1, N2 fp.Named](ins1 Decoder[N1], ins2 Decoder[N2]) Decoder[fp.Labelled2[N1, N2]] {

	return NewDecoder(
		func(ctx DecoderContext, a string) fp.Try[fp.Labelled2[N1, N2]] {
			m := map[string]any{}

			err := json.Unmarshal([]byte(a), &m)
			if err != nil {
				return try.Failure[fp.Labelled2[N1, N2]](err)
			}

			var a1 N1
			toDecode := option.Map(option.Of(m[a1.Name()]), func(v any) string {
				b, _ := json.Marshal(v)
				return string(b)
			}).OrElse("null")

			v1 := ins1.Decode(ctx, toDecode)

			var a2 N2
			toDecode = option.Map(option.Of(m[a2.Name()]), func(v any) string {
				b, _ := json.Marshal(v)
				return string(b)
			}).OrElse("null")

			v2 := ins2.Decode(ctx, toDecode)

			return try.Map2(v1, v2, func(a N1, b N2) fp.Labelled2[N1, N2] {
				return as.Labelled2(a, b)
			})
		},
	)
}
