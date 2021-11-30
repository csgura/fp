//go:generate go run github.com/csgura/fp/internal/generator/hlist_gen
package hlist

import "fmt"

// Sealed is contraints interface type to force some argument type to be one of Cons[_,_] | Nil
// but go does not support existential type
// since it has non public method isNil(),  nothing can implement this interface except Cons and Nil
type Sealed interface {
	isNil() bool
	HasTail() bool
	// Cons[_,_] | Nil
}

// Header is constrains interface type,  enforce Head type of Cons is HT
type Header[HT any] interface {
	Sealed
	Head() HT
}

// Cons means H :: T
// zero value of Cons[H,T] is not allowed.
// so Cons defined as interface type
type Cons[H, T any] interface {
	Sealed
	Header[H]
	Tail() T
}

type Nil struct{}

func (r Nil) Head() Nil {
	return r
}

func (r Nil) Tail() Nil {
	return r
}

func (r Nil) isNil() bool {
	return true
}

func (r Nil) String() string {
	return "Nil"
}

func (r Nil) HasTail() bool {
	return false
}

type hlistImpl[H, T any] struct {
	head H
	tail T
}

func (r hlistImpl[H, T]) Head() H {
	return r.head
}

func (r hlistImpl[H, T]) Tail() T {
	return r.tail
}

func (r hlistImpl[H, T]) isNil() bool {
	return false
}

func (r hlistImpl[H, T]) HasTail() bool {
	return Nil{} != any(r.tail)
}

func (r hlistImpl[H, T]) String() string {
	return fmt.Sprintf("%v :: %v", r.head, r.tail)
}

func hlist[H any, T Sealed](h H, t T) Cons[H, T] {
	return hlistImpl[H, T]{h, t}
}

func Concact[H any, T Sealed](h H, t T) Cons[H, T] {
	return hlist(h, t)
}

func Of1[H any](h H) Cons[H, Nil] {
	return hlist(h, Nil{})
}

func Empty() Nil {
	return Nil{}
}

func Ap1[A, R any](f func(a A) R) func(v Cons[A, Nil]) R {
	return func(v Cons[A, Nil]) R {
		return f(v.Head())
	}
}

func Case1[A1, T, R any](hl Cons[A1, T], f func(a1 A1) R) R {
	return f(hl.Head())
}

// func Reverse1[A1 any](hl Cons[A1, Nil]) Cons[A1, Nil] {
// 	return hl
// }

// func Reverse2[A1, A2 any](hl Cons[A1, Cons[A2, Nil]]) Cons[A2, Cons[A1, Nil]] {
// 	panic("")
// 	//return Concact(Reverse1(hl.Tail()), hl.Head())
// }

// func Reverse3[A1, A2, A3 any](hl Cons[A1, Cons[A2, Cons[A3, Nil]]]) Cons[A3, Cons[A2, Cons[A1, Nil]]] {
// 	//panic("")

// }
