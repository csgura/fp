package fp

import "fmt"

type Show[T any] interface {
	Show(t T) string
}

type ShowFunc[T any] func(T) string

func Sprint[T any]() Show[T] {
	return ShowFunc[T](func(t T) string {
		return fmt.Sprint(t)
	})
}

func (r ShowFunc[T]) Show(t T) string {
	return r(t)
}

type Eq[T any] interface {
	Eqv(a T, b T) bool
}

type EqFunc[T any] func(a, b T) bool

func (r EqFunc[T]) Eqv(a, b T) bool {
	return r(a, b)
}

func EqGiven[T comparable]() Eq[T] {
	return EqFunc[T](func(a, b T) bool {
		return a == b
	})
}

type Hashable[T any] interface {
	Eq[T]
	Hash(T) uint32
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

func LessGiven[T ImplicitOrd]() Ord[T] {
	return LessFunc[T](func(a, b T) bool {
		return a < b
	})
}

type ord[T any] struct {
	eqv  Eq[T]
	less LessFunc[T]
}

func (r ord[T]) Eqv(a, b T) bool {
	return r.Eqv(a, b)
}

func (r ord[T]) Less(a, b T) bool {
	return r.less(a, b)
}
