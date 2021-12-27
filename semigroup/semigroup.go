package semigroup

import "github.com/csgura/fp"

func Sum[T fp.ImplicitOrd]() fp.SemigroupFunc[T] {
	return func(a, b T) T {
		return a + b
	}
}

func Product[T fp.ImplicitNum](a, b T) fp.SemigroupFunc[T] {
	return func(a, b T) T {
		return a * b
	}
}
