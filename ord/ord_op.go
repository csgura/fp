//go:generate go run github.com/csgura/fp/internal/generator/ord_gen
package ord

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
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
		return option.Applicative2(m.Less).ApOption(t1).ApOption(t2).OrElse(!t1.IsDefined())
	})
}

func Seq[T any](ord fp.Ord[T]) fp.Ord[fp.Seq[T]] {
	return eq.Seq[T](ord).ToOrd(func(a, b fp.Seq[T]) bool {
		last := fp.Min(a.Size(), b.Size())
		for i := 0; i < last; i++ {
			if ord.Less(a[i], b[i]) {
				return true
			}
		}
		return a.Size() < b.Size()
	})
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
