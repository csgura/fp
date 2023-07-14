package fp

import "fmt"

type ShowOption struct {
	Indent        string
	OmitEmpty     bool
	currentIndent string
}

func (r ShowOption) CurrentIndent() string {
	return r.currentIndent
}

func (r ShowOption) IncreaseIndent() ShowOption {
	r.currentIndent = r.currentIndent + r.Indent
	return r
}

type Show[T any] interface {
	Show(t T) string
	ShowIndent(t T, option ShowOption) string
}

type ShowIndentFunc[T any] func(t T, option ShowOption) string

func (r ShowIndentFunc[T]) Show(t T) string {
	return r(t, ShowOption{})
}

func (r ShowIndentFunc[T]) ShowIndent(t T, opt ShowOption) string {
	return r(t, opt)
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

func (r ShowFunc[T]) ShowIndent(t T, opt ShowOption) string {
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

type Clone[T any] interface {
	Clone(t T) T
}

type CloneFunc[T any] func(t T) T

func (r CloneFunc[T]) Clone(t T) T {
	return r(t)
}
