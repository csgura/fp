package clone

import (
	"github.com/csgura/fp"

	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

type Derives[T any] interface {
}

func New[T any](f func(T) T) fp.Clone[T] {
	return fp.CloneFunc[T](f)
}

func Ptr[T any](tshow lazy.Eval[fp.Clone[T]]) fp.Clone[*T] {
	return New(func(pt *T) *T {
		if pt == nil {
			return nil
		}
		var t = *pt
		return &t
	})
}

func Given[T any]() fp.Clone[T] {
	return New(func(t T) T {
		return t
	})
}

var HNil = New(func(hlist.Nil) hlist.Nil {
	return hlist.Empty()
})

func Seq[T any](tclone fp.Clone[T]) fp.Clone[fp.Seq[T]] {
	return New(func(s fp.Seq[T]) fp.Seq[T] {
		return seq.Map(s, tclone.Clone)
	})
}

func GoMap[K comparable, V any](clonek fp.Clone[K], clonev fp.Clone[V]) fp.Clone[map[K]V] {
	return New(func(s map[K]V) map[K]V {
		ret := map[K]V{}
		for k, v := range s {
			ret[clonek.Clone(k)] = clonev.Clone(v)
		}
		return ret
	})
}

func Slice[T any](tclone fp.Clone[T]) fp.Clone[[]T] {
	return New(func(s []T) []T {
		return seq.Map(s, tclone.Clone)
	})
}

func Option[T any](tclone fp.Clone[T]) fp.Clone[fp.Option[T]] {
	return New(func(s fp.Option[T]) fp.Option[T] {
		return option.Map(s, tclone.Clone)
	})
}

func HCons[H any, T hlist.HList](hclone fp.Clone[H], tclone fp.Clone[T]) fp.Clone[hlist.Cons[H, T]] {
	return New(func(list hlist.Cons[H, T]) hlist.Cons[H, T] {

		h := hclone.Clone(list.Head())
		t := tclone.Clone(list.Tail())

		return hlist.Concat(h, t)

	})
}

func Generic[A, Repr any](gen fp.Generic[A, Repr], reprClone fp.Clone[Repr]) fp.Clone[A] {
	return New(func(a A) A {
		return gen.From(reprClone.Clone(gen.To(a)))
	})
}
