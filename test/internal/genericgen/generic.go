package genericgen

import (
	"github.com/csgura/fp/genfp"
)

type Identity[T any] struct {
	v T
}

func Pure[T any](v T) Identity[T] {
	return Identity[T]{v}
}

func (r Identity[T]) FlatMap[S any](f func(T) Identity[S]) Identity[S] {
	return f(r.v)
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[A any]() genfp.GenerateMonadMethods[Identity[A]] {
	return genfp.GenerateMonadMethods[Identity[A]]{
		File:     "monad_method.go",
		TypeParm: genfp.TypeOf[A](),
		Pure:     Pure[A],
	}
}
