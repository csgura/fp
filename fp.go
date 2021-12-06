//go:generate go run github.com/csgura/fp/internal/generator/fp_gen
package fp

import (
	"fmt"
	"reflect"

	"github.com/csgura/fp/hlist"
)

type Unit struct {
}

func (r Unit) String() string {
	return "()"
}

type Tuple1[T1 any] struct {
	I1 T1
}

func (r Tuple1[T1]) Head() T1 {
	return r.I1
}

func (r Tuple1[T1]) Tail() Unit {
	return Unit{}
}

func (r Tuple1[T1]) ToHList() hlist.Cons[T1, hlist.Nil] {
	return hlist.Concact(r.Head(), hlist.Empty())
}

type Func0[R any] func() R

func Println[T any](v T) {
	fmt.Println(v)
}

func ToString[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

func TypeName[T any]() string {
	var zero *T
	return reflect.TypeOf(zero).Elem().String()
}

type ImplicitNum interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type ImplicitOrd interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Min[T ImplicitOrd](a1 T, a2 T) T {
	if a1 < a2 {
		return a1
	}
	return a2
}

type Semigroup[T any] interface {
	Combine(a T, b T) T
}

type Monoid[T any] interface {
	Semigroup[T]
	Empty() T
}

type SemigroupFunc[T any] func(a T, b T) T

func (r SemigroupFunc[T]) Empty() T {
	var zero T
	return zero
}

func (r SemigroupFunc[T]) Combine(a T, b T) T {
	return r(a, b)
}

type EmptyFunc[T any] func() T

func (r EmptyFunc[T]) Empty() T {
	return r()
}

type Eq[T any] interface {
	Eqv(a T, b T) bool
}

type Ord[T any] interface {
	Eq[T]
	Less(a T, b T) bool
}

type LessFunc[T any] func(a, b T) bool

func (r LessFunc[T]) Eqv(a, b T) bool {
	return r(a, b) == false && r(b, a) == false
}

func (r LessFunc[T]) Less(a, b T) bool {
	return r(a, b)
}

func Sum[T ImplicitOrd]() Monoid[T] {
	return SemigroupFunc[T](func(a, b T) T {
		return a + b
	})
}

type monoid[T any] struct {
	zero    EmptyFunc[T]
	combine SemigroupFunc[T]
}

func (r monoid[T]) Empty() T {
	return r.zero()
}

func (r monoid[T]) Combine(a, b T) T {
	return r.combine(a, b)
}

func Product[T ImplicitNum]() Monoid[T] {
	return monoid[T]{
		zero: func() T {
			return 1
		},
		combine: func(a, b T) T {
			return a * b
		},
	}
}

func Less[T ImplicitOrd]() Ord[T] {
	return LessFunc[T](func(a, b T) bool {
		return a < b
	})
}
