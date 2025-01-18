package minimal

import "github.com/csgura/fp/hlist"

type HList interface {
	// Cons[_,_] | Nil
}

type Nil = hlist.Nil
type Cons[H any, T HList] struct {
	head H
	tail T
}

func Head[H any, T HList](r Cons[H, T]) H {
	return r.head
}

func Tail[H any, T HList](r Cons[H, T]) T {
	return r.tail
}

func Empty() Nil {
	return Nil{}
}

func Concat[H any, T HList](h H, t T) Cons[H, T] {
	return Cons[H, T]{h, t}
}

func Unapply[H any, T HList](list Cons[H, T]) (H, T) {
	return Head(list), Tail(list)
}

func IsNil[T HList](v T) bool {
	switch any(v).(type) {
	case Nil:
		return true
	}
	return false
}
