package show

import (
	"fmt"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
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

func Ptr[T any](tshow lazy.Eval[fp.Show[T]]) fp.Show[*T] {
	return New(func(pt *T) string {
		if pt != nil {
			return tshow.Get().Show(*pt)
		}
		return "nil"
	})
}

func Given[T any]() fp.Show[T] {
	return fp.Sprint[T]()
}

var HNil = New(func(hlist.Nil) string {
	return "Nil"
})

func Seq[T any](tshow fp.Show[T]) fp.Show[fp.Seq[T]] {
	return New(func(s fp.Seq[T]) string {
		return "Seq(" + seq.Map(s, tshow.Show).MakeString(",") + ")"
	})
}

func Set[V any](showv fp.Show[V]) fp.Show[fp.Set[V]] {
	return New(func(v fp.Set[V]) string {
		return "Set(" + iterator.Map(v.Iterator(), func(v V) string {
			return fmt.Sprintf("%s", showv.Show(v))
		}).MakeString(",") + ")"
	})
}

func Map[K, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[fp.Map[K, V]] {
	return New(func(v fp.Map[K, V]) string {
		return "Map(" + iterator.Map(v.Iterator(), func(t fp.Tuple2[K, V]) string {
			return fmt.Sprintf("%s: %s", showk.Show(t.I1), showv.Show(t.I2))
		}).MakeString(",") + ")"
	})
}

func GoMap[K comparable, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[map[K]V] {
	return New(func(v map[K]V) string {
		return "Map(" + iterator.Map(mutable.MapOf(v).Iterator(), func(t fp.Tuple2[K, V]) string {
			return fmt.Sprintf("%s: %s", showk.Show(t.I1), showv.Show(t.I2))
		}).MakeString(",") + ")"
	})
}

func Slice[T any](tshow fp.Show[T]) fp.Show[[]T] {
	return New(func(s []T) string {
		return "Seq(" + seq.Map(s, tshow.Show).MakeString(",") + ")"
	})
}

func Option[T any](tshow fp.Show[T]) fp.Show[fp.Option[T]] {
	return New(func(s fp.Option[T]) string {
		if s.IsDefined() {
			return fmt.Sprintf("Some(%s)", tshow.Show(s.Get()))
		}
		return "None"
	})
}

func Named[T fp.NamedField[A], A any](ashow fp.Show[A]) fp.Show[T] {
	return New(func(s T) string {
		return fmt.Sprintf("%s:%s", s.Name(), ashow.Show(s.Value()))
	})
}

func HConsLabelled[H fp.Named, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return New(func(list hlist.Cons[H, T]) string {

		hstr := hshow.Show(list.Head())
		tstr := tshow.Show(list.Tail())

		if tstr != "Nil" {
			return fmt.Sprintf("%s,%s", hstr, tstr)
		}
		return hstr
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
