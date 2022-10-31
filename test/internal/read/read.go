package read

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/try"
)

type Derives[T any] interface {
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type Result[T any] struct {
	value   T
	remains string
}

func MapResult[A, B any](a Result[A], fab func(A) B) Result[B] {
	return ResultMutable[B]{
		Value:   fab(a.Value()),
		Remains: a.Remains(),
	}.AsImmutable()
}

type Read[T any] interface {
	Read(str string) fp.Try[T]
	Reads(str string) fp.Try[Result[T]]
}

type ReadFunc[T any] func(string) fp.Try[Result[T]]

func (r ReadFunc[T]) Read(str string) fp.Try[T] {
	return try.Map(r.Reads(str), Result[T].Value)
}

func (r ReadFunc[T]) Reads(str string) fp.Try[Result[T]] {
	return r(str)
}

func New[T any](f func(string) fp.Try[Result[T]]) Read[T] {
	return ReadFunc[T](f)
}

func readTokens(s string) Result[string] {
	s = strings.TrimSpace(s)

	r := []rune(s)

	if len(r) > 0 && r[0] == '"' {
		for i := 1; i < len(r); i++ {
			if r[i] == '"' {
				return ResultMutable[string]{
					Value:   string(r[1:i]),
					Remains: string(r[i+1:]),
				}.AsImmutable()
			}
		}
	}

	for i := 0; i < len(r); i++ {
		if unicode.IsSpace(r[i]) || r[i] == ',' || r[i] == '(' || r[i] == ')' || r[i] == ':' {
			return ResultMutable[string]{
				Value:   string(r[:i]),
				Remains: string(r[i:]),
			}.AsImmutable()
		}
	}
	return ResultMutable[string]{
		Value:   s,
		Remains: "",
	}.AsImmutable()
}

var String = New(func(s string) fp.Try[Result[string]] {
	return try.Success(readTokens(s))
})

func UInt[T fp.ImplicitUInt]() Read[T] {
	return New(func(s string) fp.Try[Result[T]] {
		t := readTokens(s)
		n, err := strconv.ParseUint(t.Value(), 0, 64)
		if err != nil {
			return try.Failure[Result[T]](err)
		}
		return try.Success(ResultMutable[T]{
			Value:   T(n),
			Remains: t.Remains(),
		}.AsImmutable())
	})
}

func Int[T fp.ImplicitInt]() Read[T] {
	return New(func(s string) fp.Try[Result[T]] {
		t := readTokens(s)

		n, err := strconv.ParseInt(t.Value(), 0, 64)
		if err != nil {
			return try.Failure[Result[T]](err)
		}
		return try.Success(ResultMutable[T]{
			Value:   T(n),
			Remains: t.Remains(),
		}.AsImmutable())
	})
}

func Float[T fp.ImplicitFloat]() Read[T] {
	return New(func(s string) fp.Try[Result[T]] {
		t := readTokens(s)

		n, err := strconv.ParseFloat(t.Value(), 64)
		if err != nil {
			return try.Failure[Result[T]](err)
		}
		return try.Success(ResultMutable[T]{
			Value:   T(n),
			Remains: t.Remains(),
		}.AsImmutable())
	})
}

var HNil = New(func(s string) fp.Try[Result[hlist.Nil]] {
	r := readTokens(s)
	if r.Value() == "Nil" {
		return try.Success(ResultMutable[hlist.Nil]{
			Value:   hlist.Empty(),
			Remains: r.Remains(),
		}.AsImmutable())
	}
	return try.Failure[Result[hlist.Nil]](fmt.Errorf("expected Nil but %s", r.Value()))
})

func skipColonColon(s string) string {
	idx := strings.Index(s, "::")
	if idx >= 0 {
		return strings.TrimSpace(s[idx+2:])
	}
	return s
}

func HCons[H any, T hlist.HList](hread Read[H], tread Read[T]) Read[hlist.Cons[H, T]] {
	return New(func(s string) fp.Try[Result[hlist.Cons[H, T]]] {
		hres := hread.Reads(s)
		return try.FlatMap(hres, func(hr Result[H]) fp.Try[Result[hlist.Cons[H, T]]] {
			nextHead := skipColonColon(hr.Remains())
			return try.Map(tread.Reads(nextHead), func(tr Result[T]) Result[hlist.Cons[H, T]] {
				return Result[hlist.Cons[H, T]]{
					value:   hlist.Concat(hr.value, tr.value),
					remains: tr.remains,
				}
			})
		})
	})
}

func Map[A, B any](aread Read[A], fab func(A) B) Read[B] {
	return New(func(s string) fp.Try[Result[B]] {
		return try.Map(aread.Reads(s), func(r Result[A]) Result[B] {
			return MapResult(r, fab)
		})
	})
}

func Generic[T, Repr any](gen fp.Generic[T, Repr], reprRead Read[Repr]) Read[T] {
	return New(func(s string) fp.Try[Result[T]] {
		var t T
		tpname := fmt.Sprintf("%T", t)
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, tpname) {
			s = strings.Trim(s[len(tpname)+1:], "()")
			res := reprRead.Reads(s)
			return try.Map(res, func(r Result[Repr]) Result[T] {
				return MapResult(r, gen.From)
			})
		}
		return try.Failure[Result[T]](fmt.Errorf("expected type name %s but %s", tpname, s))
	})
}
