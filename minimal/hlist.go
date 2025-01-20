package minimal

import "github.com/csgura/fp/hlist"

type HList interface {
	// Cons[_,_] | Nil
}

type Nil = hlist.Nil
type Cons[H any, T HList] struct {
	Head H
	Tail T
}

func Empty() Nil {
	return Nil{}
}

func Concat[H any, T HList](h H, t T) Cons[H, T] {
	return Cons[H, T]{h, t}
}

func IsNil(v HList) bool {
	return Nil{} == v
}
