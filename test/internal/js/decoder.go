package js

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

type Decoder[T any] interface {
	Decode(string) fp.Try[T]
}

type DecoderFunc[T any] func(string) fp.Try[T]

func (r DecoderFunc[T]) Decode(t string) fp.Try[T] {
	return r(t)
}

func NewDecoder[T any](f func(a string) fp.Try[T]) Decoder[T] {
	return DecoderFunc[T](f)
}

var DecoderString = NewDecoder(func(a string) fp.Try[string] {
	if a[0] == '"' {
		return try.Success(a[1 : len(a)-1])
	}
	return try.Failure[string](fmt.Errorf("invalid string literal"))
})

func DecoderNumber[T fp.ImplicitNum]() Decoder[T] {
	return NewDecoder(func(a string) fp.Try[T] {
		r, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return try.Failure[T](err)
		}
		return try.Success(T(r))
	})
}

var DecoderTime = NewDecoder(func(a string) fp.Try[time.Time] {
	return try.FlatMap(DecoderString.Decode(a), func(v string) fp.Try[time.Time] {
		return try.Apply(time.Parse(time.RFC3339, v))
	})

})

var DecoderUnit = NewDecoder(func(a string) fp.Try[fp.Unit] {
	return try.Success(fp.Unit{})
})

var DecoderHNil Decoder[hlist.Nil] = NewDecoder(func(a string) fp.Try[hlist.Nil] {
	return try.Success(hlist.Empty())
})

func DecoderHCons[H any, T hlist.HList](heq Decoder[H], teq Decoder[T]) Decoder[hlist.Cons[H, T]] {
	return NewDecoder(func(a string) fp.Try[hlist.Cons[H, T]] {
		commaIdx := strings.Index(a, ",")
		head := heq.Decode(a[:commaIdx])
		tailstr := ""
		if commaIdx > 0 {
			tailstr = a[commaIdx+1:]

		}
		return try.FlatMap(head, func(h H) fp.Try[hlist.Cons[H, T]] {
			return try.Map(teq.Decode(tailstr), func(t T) hlist.Cons[H, T] {
				return hlist.Concat(h, t)
			})
		})
	})
}

func DecoderHConsLabelled[H fp.Named, T hlist.HList](heq Decoder[H], teq Decoder[T]) Decoder[hlist.Cons[H, T]] {
	return NewDecoder(func(a string) fp.Try[hlist.Cons[H, T]] {

		m := map[string]any{}

		err := json.Unmarshal([]byte(a), &m)
		if err != nil {
			return try.Failure[hlist.Cons[H, T]](err)
		}

		var h H
		toDecode := option.Map(option.Of(m[h.Name()]), func(v any) string {
			b, _ := json.Marshal(v)
			return string(b)
		}).OrElse("null")

		head := heq.Decode(toDecode)

		return try.FlatMap(head, func(h H) fp.Try[hlist.Cons[H, T]] {
			return try.Map(teq.Decode(a), func(t T) hlist.Cons[H, T] {
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
	return NewDecoder(func(a string) fp.Try[U] {
		return try.Map(instance.Decode(a), fn)
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
	return NewDecoder(func(a string) fp.Try[T] {
		ret := enc.Decode(a)
		return try.Map(ret, func(v A) T {
			var zero T
			return zero.WithValue(v)
		})
	})
}

func DecoderLabelled2[N1, N2 fp.Named](ins1 Decoder[N1], ins2 Decoder[N2]) Decoder[fp.Labelled2[N1, N2]] {

	return NewDecoder(
		func(a string) fp.Try[fp.Labelled2[N1, N2]] {
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

			v1 := ins1.Decode(toDecode)

			var a2 N2
			toDecode = option.Map(option.Of(m[a2.Name()]), func(v any) string {
				b, _ := json.Marshal(v)
				return string(b)
			}).OrElse("null")

			v2 := ins2.Decode(toDecode)

			return try.Map2(v1, v2, func(a N1, b N2) fp.Labelled2[N1, N2] {
				return as.Labelled2(a, b)
			})
		},
	)
}
