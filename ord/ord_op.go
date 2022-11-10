//go:generate go run github.com/csgura/fp/internal/generator/ord_gen
package ord

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
)

type ord[T any] struct {
	eqv  fp.Eq[T]
	less fp.LessFunc[T]
}

func (r ord[T]) Eqv(a, b T) bool {
	return r.eqv.Eqv(a, b)
}

func (r ord[T]) Less(a, b T) bool {
	return r.less(a, b)
}

func (r ord[T]) ToOrd(less fp.LessFunc[T]) fp.Ord[T] {
	return ord[T]{
		r.eqv, less,
	}
}

func New[T any](eqv fp.Eq[T], less fp.LessFunc[T]) fp.Ord[T] {
	return ord[T]{
		eqv, less,
	}
}

func Tuple1[A any](a fp.Ord[A]) fp.Ord[fp.Tuple1[A]] {
	return New[fp.Tuple1[A]](
		fp.EqFunc[fp.Tuple1[A]](func(t1 fp.Tuple1[A], t2 fp.Tuple1[A]) bool {
			return a.Eqv(t1.I1, t2.I1)
		}),
		fp.LessFunc[fp.Tuple1[A]](func(t1 fp.Tuple1[A], t2 fp.Tuple1[A]) bool {
			if a.Less(t1.I1, t2.I1) {
				return true
			}
			return false
		}),
	)
}

func Option[T any](m fp.Ord[T]) fp.Ord[fp.Option[T]] {
	return fp.LessFunc[fp.Option[T]](func(t1 fp.Option[T], t2 fp.Option[T]) bool {
		if !t1.IsDefined() && !t2.IsDefined() {
			return false
		}
		return option.Map2(t1, t2, m.Less).OrElse(t1.IsEmpty())
	})
}

func Seq[T any](ord fp.Ord[T]) fp.Ord[fp.Seq[T]] {
	return New(eq.Seq[T](ord), func(a, b fp.Seq[T]) bool {
		last := fp.Min(a.Size(), b.Size())
		for i := 0; i < last; i++ {
			if ord.Less(a[i], b[i]) {
				return true
			}
		}
		return a.Size() < b.Size()
	})
}

func Slice[T any](ord fp.Ord[T]) fp.Ord[[]T] {
	return ContraMap(Seq(ord), as.Seq[T])
}

var HNil fp.Ord[hlist.Nil] = New(fp.EqGiven[hlist.Nil](), func(a, b hlist.Nil) bool { return false })

func HCons[H any, T hlist.HList](heq fp.Ord[H], teq fp.Ord[T]) fp.Ord[hlist.Cons[H, T]] {
	return New(eq.HCons[H, T](heq, teq), func(a, b hlist.Cons[H, T]) bool {
		if heq.Less(a.Head(), b.Head()) {
			return true
		}

		if heq.Less(b.Head(), a.Head()) {
			return false
		}

		return teq.Less(a.Tail(), b.Tail())
	})
}

func Given[T fp.ImplicitOrd]() fp.Ord[T] {
	return fp.LessGiven[T]()
}

func ContraMap[T, U any](instance fp.Ord[T], fn func(U) T) fp.Ord[U] {
	return New(eq.ContraMap[T](instance, fn), func(a, b U) bool {
		return instance.Less(fn(a), fn(b))
	})
}

type Derives[T any] interface{ Target() T }

// nil goes to first
func Ptr[T any](ordT lazy.Eval[fp.Ord[T]]) fp.Ord[*T] {
	return New(eq.Ptr(lazy.Call(func() fp.Eq[T] {
		return ordT.Get()
	})), func(a, b *T) bool {

		if a != nil && b != nil {
			return ordT.Get().Less(*a, *b)
		}

		if a == nil && b == nil {
			return false
		}

		return a == nil
	})
}
