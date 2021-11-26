//go:generate go run github.com/csgura/fp/internal/generator/hlist_gen
package hlist

import "fmt"

type Header[HT any] interface {
	Head() HT
}

type Cons[H, T any] interface {
	Head() H
	Tail() T
	IsNil() bool
}

type Nil struct{}

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

func (r hlistImpl[H, T]) IsNil() bool {
	return false
}

func (r hlistImpl[H, T]) String() string {
	return fmt.Sprintf("%v :: %v", r.head, r.tail)
}

func hlist[H, T any](h H, t T) Cons[H, T] {
	return hlistImpl[H, T]{h, t}
}

func Concact[H, T any](h H, t T) Cons[H, T] {
	return hlist(h, t)
}

func Of[H any](h H) Cons[H, Nil] {
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
