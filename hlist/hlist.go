//go:generate go run github.com/csgura/fp/internal/generator/hlist_gen
package hlist

type Header[HT any] interface {
	Head() HT
}

type Cons[H, T any] interface {
	Head() H
	Tail() T
}

type Nil struct{}

func (r Nil) Head() Nil {
	return r
}

func (r Nil) Tail() Nil {
	return r
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

// func Case2[A, B, R, T any](list Cons[A, Cons[B, T]], f func(a A, b B) R) {
// 	f(list.Head(), list.Tail().Head())
// }

// func Case3[A, B, C, R, T any](list Cons[A, Cons[B, Cons[C, T]]], f func(a A, b B, c C) R) {
// 	f(list.Head(), list.Tail().Head(), list.Tail().Tail().Head())
// }
