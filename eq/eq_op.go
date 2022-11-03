//go:generate go run github.com/csgura/fp/internal/generator/eq_gen
package eq

import (
	"bytes"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
)

func New[T any](f func(a, b T) bool) fp.Eq[T] {
	return fp.EqFunc[T](f)
}

var Bytes fp.Eq[[]byte] = New(bytes.Equal)

func Tuple1[A any](a fp.Eq[A]) fp.Eq[fp.Tuple1[A]] {
	return New(
		fp.EqFunc[fp.Tuple1[A]](func(t1 fp.Tuple1[A], t2 fp.Tuple1[A]) bool {
			return a.Eqv(t1.I1, t2.I1)
		}),
	)
}

func Option[T any](eq fp.Eq[T]) fp.Eq[fp.Option[T]] {
	return fp.EqFunc[fp.Option[T]](func(t1 fp.Option[T], t2 fp.Option[T]) bool {
		if t1.IsEmpty() && t2.IsEmpty() {
			return true
		}

		if t1.IsDefined() && t2.IsDefined() {
			return eq.Eqv(t1.Get(), t2.Get())
		}

		return false
	})
}

func Seq[T any](eq fp.Eq[T]) fp.Eq[fp.Seq[T]] {
	return New(func(a, b fp.Seq[T]) bool {
		if a.Size() != b.Size() {
			return false
		}

		for i := range a {
			if !eq.Eqv(a[i], b[i]) {
				return false
			}
		}
		return true
	})
}

func Slice[T any](eq fp.Eq[T]) fp.Eq[[]T] {
	return ContraMap(Seq(eq), as.Seq[T])
}

var HNil fp.Eq[hlist.Nil] = fp.EqGiven[hlist.Nil]()

func HCons[H any, T hlist.HList](heq fp.Eq[H], teq fp.Eq[T]) fp.Eq[hlist.Cons[H, T]] {
	return New(func(a, b hlist.Cons[H, T]) bool {
		return heq.Eqv(a.Head(), b.Head()) && teq.Eqv(a.Tail(), b.Tail())
	})
}

func Given[T comparable]() fp.Eq[T] {
	return fp.EqGiven[T]()
}

func Ptr[T any](eq fp.Eq[T]) fp.Eq[*T] {
	return New(func(a, b *T) bool {
		if a == nil && b == nil {
			return true
		}

		if a != nil && b != nil {
			return eq.Eqv(*a, *b)
		}

		return false
	})
}

func ContraMap[T, U any](instance fp.Eq[T], fn func(U) T) fp.Eq[U] {
	return New(func(a, b U) bool {
		return instance.Eqv(fn(a), fn(b))
	})
}

type Derives[T any] interface {
	Target() T
}

var String = Given[string]()

// for import test

// var SeqTuple2Correct = New(func(a, b fp.Seq[fp.Tuple2[string, int]]) bool {
// 	return a.Size() == b.Size()
// })

// var SeqTuple2False = New(func(a, b fp.Seq[fp.Tuple2[int, int]]) bool {
// 	return a.Size() == b.Size()
// })

// var SeqTuple3 = New(func(a, b fp.Seq[fp.Tuple3[int, int, int]]) bool {
// 	return a.Size() == b.Size()
// })

// func SeqTupleFuncCorrect[A, B any]() fp.Eq[fp.Seq[fp.Tuple2[A, B]]] {
// 	return New(func(a, b fp.Seq[fp.Tuple2[A, B]]) bool {
// 		return a.Size() == b.Size()
// 	})
// }

// func SeqTupleFuncFalse[B any]() fp.Eq[fp.Seq[fp.Tuple2[int, B]]] {
// 	return New(func(a, b fp.Seq[fp.Tuple2[int, B]]) bool {
// 		return a.Size() == b.Size()
// 	})
// }

// func Hello[A interface{ Hello() string }]() fp.Eq[A] {
// 	panic("")
// }
