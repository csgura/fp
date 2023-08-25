//go:generate go run github.com/csgura/fp/internal/generator/hlist_gen
package hlist

import (
	"fmt"
)

// Sealed is contraints interface type to force some argument type to be one of Cons[_,_] | Nil
// but go does not support existential type
// since it has non public method sealed(),  nothing can implement this interface except Cons and Nil
type HList interface {
	IsNil() bool
}

// Header is constrains interface type,  enforce Head type of Cons is HT
type Header[HT any] interface {
	HList
	Head() HT
}

type Nil struct {
}

func (r Nil) Head() Nil {
	return r
}

func (r Nil) Tail() Nil {
	return r
}

func (r Nil) IsNil() bool {
	return true
}

func (r Nil) String() string {
	return "Nil"
}

type Cons[H any, T HList] struct {
	head H
	tail T
}

func (r Cons[H, T]) Head() H {
	return r.head
}

func (r Cons[H, T]) Tail() T {
	return r.tail
}

func (r Cons[H, T]) IsNil() bool {
	return false
}

func (r Cons[H, T]) String() string {
	return fmt.Sprintf("%v :: %v", r.head, r.tail)
}

func Concat[H any, T HList](h H, t T) Cons[H, T] {
	return Cons[H, T]{h, t}
}

func Of1[H any](h H) Cons[H, Nil] {
	return Concat(h, Nil(Nil{}))
}

func Empty() Nil {
	return Nil{}
}

func Lift1[A, R any](f func(a A) R) func(v Cons[A, Nil]) R {
	return func(v Cons[A, Nil]) R {
		return f(v.Head())
	}
}

func Rift1[A, R any](f func(a A) R) func(v Cons[A, Nil]) R {
	return func(v Cons[A, Nil]) R {
		return f(v.Head())
	}
}

func Case1[A1 any, T HList, R any](hl Cons[A1, T], f func(a1 A1) R) R {
	return f(hl.Head())
}

// func Reverse1[A1 any](hl Cons[A1, Nil]) Cons[A1, Nil] {
// 	return hl
// }

// func Reverse2[A1, A2 any](hl Cons[A1, Cons[A2, Nil]]) Cons[A2, Cons[A1, Nil]] {
// 	panic("")
// 	//return Concat(Reverse1(hl.Tail()), hl.Head())
// }

// func Reverse3[A1, A2, A3 any](hl Cons[A1, Cons[A2, Cons[A3, Nil]]]) Cons[A3, Cons[A2, Cons[A1, Nil]]] {
// 	//panic("")

// }

func Unapply[H any, T HList](list Cons[H, T]) (H, T) {
	return list.Head(), list.Tail()
}
