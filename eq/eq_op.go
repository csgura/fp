//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package eq

import (
	"bytes"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
)

func New[T any](f func(a, b T) bool) fp.Eq[T] {
	return fp.EqFunc[T](f)
}

var Time fp.Eq[time.Time] = New(time.Time.Equal)
var Bytes fp.Eq[[]byte] = New(bytes.Equal)
var String = Given[string]()
var PtrBytes fp.Eq[*[]byte] = Ptr(lazy.Done(Bytes))

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
	return New(func(a, b fp.Slice[T]) bool {
		if len(a) != len(b) {
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

var HNil fp.Eq[hlist.Nil] = fp.EqGiven[hlist.Nil]()

func HCons[H any, T hlist.HList](heq fp.Eq[H], teq fp.Eq[T]) fp.Eq[hlist.Cons[H, T]] {
	return New(func(a, b hlist.Cons[H, T]) bool {
		return heq.Eqv(hlist.Head(a), hlist.Head(b)) && teq.Eqv(hlist.Tail(a), hlist.Tail(b))
	})
}

func Given[T comparable]() fp.Eq[T] {
	return fp.EqGiven[T]()
}

func GivenValue[T comparable](a T) fp.Predicate[T] {
	return func(b T) bool {
		return Given[T]().Eqv(a, b)
	}
}

func GivenPtr[T comparable](a *T) fp.Predicate[*T] {
	return func(b *T) bool {
		return PtrGiven[T]().Eqv(a, b)
	}
}

func GivenFieldValue[S any, T comparable](getter func(S) T, a T) fp.Predicate[S] {
	return func(s S) bool {
		return Given[T]().Eqv(getter(s), a)
	}
}

func GivenFieldPtr[S any, T comparable](getter func(S) *T, a *T) fp.Predicate[S] {
	return func(s S) bool {
		return PtrGiven[T]().Eqv(getter(s), a)
	}
}

func Ptr[T any](eq lazy.Eval[fp.Eq[T]]) fp.Eq[*T] {
	return New(func(a, b *T) bool {
		if a == nil && b == nil {
			return true
		}

		if a != nil && b != nil {
			return eq.Get().Eqv(*a, *b)
		}

		return false
	})
}

func PtrGiven[T comparable]() fp.Eq[*T] {
	return Ptr(lazy.Done(Given[T]()))
}

func ContraMap[T, U any](instance fp.Eq[T], fn func(U) T) fp.Eq[U] {
	return New(func(a, b U) bool {
		return instance.Eqv(fn(a), fn(b))
	})
}

type Derives[T any] interface {
	Target() T
}

func GoMap[K comparable, V any](eqV fp.Eq[V]) fp.Eq[map[K]V] {
	return New(func(a, b map[K]V) bool {
		if len(a) != len(b) {
			return false
		}

		for k, av := range a {
			bv, ok := b[k]
			if !ok {
				return false
			}

			if !eqV.Eqv(av, bv) {
				return false
			}
		}
		return true
	})
}

func FpMap[K, V any](eqV fp.Eq[V]) fp.Eq[fp.Map[K, V]] {
	return New(func(a, b fp.Map[K, V]) bool {
		if a.Size() != b.Size() {
			return false
		}

		return a.Iterator().ForAll(func(v fp.Tuple2[K, V]) bool {
			bv := b.Get(v.I1)
			if bv.IsEmpty() {
				return false
			}

			return eqV.Eqv(v.I2, bv.Get())
		})
	})
}

func NotNilAnd[A any](pf fp.Predicate[A]) fp.Predicate[*A] {
	return func(a *A) bool {
		if a == nil {
			return false
		}
		return pf(*a)
	}
}

func NilOr[A any](pf fp.Predicate[A]) fp.Predicate[*A] {
	return func(a *A) bool {
		if a == nil {
			return true
		}
		return pf(*a)
	}
}

func NotZero[A comparable](v A) bool {
	var zero A

	return v != zero
}

func NotZeroAnd[A comparable](pf fp.Predicate[A]) fp.Predicate[A] {
	var zero A
	return func(a A) bool {
		if a == zero {
			return false
		}
		return pf(a)
	}
}

func ZeroOr[A comparable](pf fp.Predicate[A]) fp.Predicate[A] {
	var zero A

	return func(a A) bool {
		if a == zero {
			return true
		}
		return pf(a)
	}
}

// as.Func2(fp.Option[A].Exists).ApplyLast 와 같은 함수
func SomeAnd[A any](pf fp.Predicate[A]) fp.Predicate[fp.Option[A]] {
	//return as.Func2(fp.Option[A].Exists).ApplyLast(pf)
	//return fp.Predicate[fp.Option[A]](fp.Flip2(fp.Option[A].Exists)(pf))

	return func(a fp.Option[A]) bool {
		if a.IsEmpty() {
			return false
		}
		return pf(a.Get())
	}
}

// as.Func2(fp.Option[A].ForAll).ApplyLast  와 같은 함수
func NoneOr[A any](pf fp.Predicate[A]) fp.Predicate[fp.Option[A]] {
	return func(a fp.Option[A]) bool {
		if a.IsEmpty() {
			return true
		}
		return pf(a.Get())
	}
}

func FieldNotNilAnd[A, B any](getter func(A) *B, pf fp.Predicate[B]) fp.Predicate[A] {
	return func(a A) bool {
		p := getter(a)
		if p == nil {
			return false
		}
		return pf(*p)
	}
}

func FieldNilOr[A, B any](getter func(A) *B, pf fp.Predicate[B]) fp.Predicate[A] {
	return func(a A) bool {
		p := getter(a)
		if p == nil {
			return true
		}
		return pf(*p)
	}
}

func FieldSomeAnd[A, B any](getter func(A) fp.Option[B], pf fp.Predicate[B]) fp.Predicate[A] {
	return func(a A) bool {
		p := getter(a)
		if p.IsEmpty() {
			return false
		}
		return pf(p.Get())
	}
}

func FieldNoneOr[A, B any](getter func(A) fp.Option[B], pf fp.Predicate[B]) fp.Predicate[A] {
	return func(a A) bool {
		p := getter(a)
		if p.IsEmpty() {
			return true
		}
		return pf(p.Get())
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]( {{DeclTypeClassArgs 1 .N "fp.Eq"}} ) fp.Eq[fp.{{TupleType .N}}] {

	pt := Tuple{{dec .N}}({{CallArgs 2 .N "ins"}})

	return New(
		func(t1 , t2 fp.{{TupleType .N}}) bool {
			return ins1.Eqv(t1.I1, t2.I1) && pt.Eqv(as.Tuple{{dec .N}}(t1.Tail()), as.Tuple{{dec .N}}(t2.Tail()))
		},
	)
}
	`,
}
