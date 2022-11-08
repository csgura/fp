package show

import (
	"fmt"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/seq"
)

type Derives[T any] interface {
}

func New[T any](f func(T) string) fp.Show[T] {
	return fp.ShowFunc[T](f)
}

var Time = New(func(t time.Time) string {
	return t.Format(time.RFC3339)
})

var String = New(func(t string) string {
	return fmt.Sprintf(`"%s"`, t)
})

func Int[T fp.ImplicitInt]() fp.Show[T] {
	return fp.Sprint[T]()
}

func Number[T fp.ImplicitNum]() fp.Show[T] {
	return fp.Sprint[T]()
}

func Given[T any]() fp.Show[T] {
	return fp.Sprint[T]()
}

var HNil = New(func(hlist.Nil) string {
	return "Nil"
})

func Seq[T any](tshow fp.Show[T]) fp.Show[fp.Seq[T]] {
	return New(func(s fp.Seq[T]) string {
		return "[" + seq.Map(s, tshow.Show).MakeString(",") + "]"
	})
}

func HCons[H any, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return New(func(list hlist.Cons[H, T]) string {

		hstr := hshow.Show(list.Head())
		tstr := tshow.Show(list.Tail())

		return fmt.Sprintf("%s :: %s", hstr, tstr)
	})
}

func ContraMap[A, B any](ashow fp.Show[A], fba func(B) A) fp.Show[B] {
	return New(func(b B) string {
		return ashow.Show(fba(b))
	})
}

func Generic[A, Repr any](gen fp.Generic[A, Repr], reprShow fp.Show[Repr]) fp.Show[A] {
	return New(func(a A) string {
		return fmt.Sprintf("%s(%s)", gen.Type, reprShow.Show(gen.To(a)))
	})
}
