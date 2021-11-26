package hlist

import "github.com/csgura/fp"

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

func Ap2[A, B, R any](f func(a A, b B) R) func(v Cons[B, Cons[A, Nil]]) R {
	return func(v Cons[B, Cons[A, Nil]]) R {
		return f(v.Tail().Head(), v.Head())
	}
}

func Ap3[A, B, C, R any](f func(a A, b B, c C) R) func(v Cons[C, Cons[B, Cons[A, Nil]]]) R {
	return func(v Cons[C, Cons[B, Cons[A, Nil]]]) R {
		rf := Ap2(func(a A, b B) R {
			return f(a, b, v.Head())
		})

		return rf(v.Tail())
	}
}

func Ap4[A, B, C, D, R any](f func(a A, b B, c C, d D) R) func(v Cons[D, Cons[C, Cons[B, Cons[A, Nil]]]]) R {
	return func(v Cons[D, Cons[C, Cons[B, Cons[A, Nil]]]]) R {
		rf := Ap3(func(a A, b B, c C) R {
			return f(a, b, c, v.Head())
		})

		return rf(v.Tail())
	}
}

// func Case2[A, B, R, T any](list Cons[A, Cons[B, T]], f func(a A, b B) R) {
// 	f(list.Head(), list.Tail().Head())
// }

// func Case3[A, B, C, R, T any](list Cons[A, Cons[B, Cons[C, T]]], f func(a A, b B, c C) R) {
// 	f(list.Head(), list.Tail().Head(), list.Tail().Tail().Head())
// }

func Case2[T, A, B, R any](f func(a A, b B) R) fp.Func1[Cons[A, Cons[B, T]], R] {
	return func(v Cons[A, Cons[B, T]]) R {
		return f(v.Head(), v.Tail().Head())
	}
}

func Case3[A, B, C, R, T any](list Cons[A, Cons[B, Cons[C, T]]], f func(a A, b B, c C) R) {
	f(list.Head(), list.Tail().Head(), list.Tail().Tail().Head())
}
