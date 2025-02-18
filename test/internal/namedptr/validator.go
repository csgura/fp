package namedptr

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/minimal"
)

type Validator[T any] interface {
	Validate(T) error
}

type Derives[T any] interface{}

type validateFunc[T any] func(T) error

func (r validateFunc[T]) Validate(t T) error {
	return r(t)
}

func New[T any](f func(T) error) Validator[T] {
	return validateFunc[T](f)
}

var HNil = New(func(t hlist.Nil) error {
	return nil
})

func StructHCons[H any, T minimal.HList](hv Validator[H], tv Validator[T]) Validator[minimal.Cons[H, T]] {
	return New(func(hlist minimal.Cons[H, T]) error {
		err := hv.Validate(hlist.Head)
		if err != nil {
			return err
		}
		return tv.Validate(hlist.Tail)
	})
}

func ContraGeneric[A, Repr any](name string, kind string, reprIns Validator[Repr], to func(A) Repr) Validator[A] {
	return New(func(a A) error {
		return reprIns.Validate(to(a))
	})
}

func Named[T any](name fp.Named, v Validator[T]) Validator[T] {
	return New(func(t T) error {
		return nil
	})
}

func NamedInt(name fp.Named) Validator[int] {
	return New(func(i int) error {
		return nil
	})
}

func NamedSlice[T any](name fp.Named, v Validator[T]) Validator[[]T] {
	return New(func(t []T) error {
		return nil
	})
}

var Int = New(func(t int) error {
	return nil
})

func Struct1Container[T any](v Validator[T]) Validator[fp.Tuple1[T]] {
	return New(func(t fp.Tuple1[T]) error {
		return nil
	})
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive
var _ Derives[Validator[Hello]]

type Hello struct {
	v int
	s []int
}

type Container struct {
	List []Value
}

type Value struct {
	Present int
}

// @fp.Derive(recursive=true)
var _ Derives[Validator[Container]]
