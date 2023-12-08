//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package ord

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
)

func New[T any](eqv fp.Eq[T], less fp.LessFunc[T]) fp.Ord[T] {
	return fp.CompareFunc[T](func(a, b T) int {
		if eqv.Eqv(a, b) {
			return 0
		}
		return less.Compare(a, b)
	})
}

func Tuple1[A any](a fp.Ord[A]) fp.Ord[fp.Tuple1[A]] {
	return New[fp.Tuple1[A]](
		fp.EqFunc[fp.Tuple1[A]](func(t1 fp.Tuple1[A], t2 fp.Tuple1[A]) bool {
			return a.Eqv(t1.I1, t2.I1)
		}),
		fp.LessFunc[fp.Tuple1[A]](func(t1 fp.Tuple1[A], t2 fp.Tuple1[A]) bool {
			return a.Less(t1.I1, t2.I1)
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

// java 의 Comparator 의 comparing 함수
func GivenField[S any, T fp.ImplicitOrd](getter func(S) T) fp.Ord[S] {
	return ContraMap(Given[T](), getter)
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

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/eq", Name: "eq"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Ord"}} ) fp.Ord[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New( eq.New( func( a, b fp.{{TupleType .N}} ) bool {
		return ins1.Eqv(a.Head(),b.Head()) && pt.Eqv(as.Tuple{{dec .N}}(a.Tail()), as.Tuple{{dec .N}}(b.Tail()))
	}), fp.LessFunc[fp.{{TupleType .N}}](func(t1 , t2 fp.{{TupleType .N}}) bool {
		if ins1.Less(t1.I1, t2.I1) {
			return true
		}
		if ins1.Less(t2.I1, t1.I1) {
			return false
		}
		return pt.Less(as.Tuple{{dec .N}}(t1.Tail()), as.Tuple{{dec .N}}(t2.Tail()))
	}))
}
	`,
}
