package fp

import (
	"fmt"
	"strings"
)

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
	Append(buf []string, t T, option ShowOption) []string
	Stringer(t T, option ShowOption) fmt.Stringer
}

type ShowIndentFunc[T any] func(t T, option ShowOption) string

func (r ShowIndentFunc[T]) Show(t T) string {
	return r(t, ShowOption{})
}

func (r ShowIndentFunc[T]) ShowIndent(t T, opt ShowOption) string {
	return r(t, opt)
}

func (r ShowIndentFunc[T]) Append(buf []string, t T, option ShowOption) []string {
	return append(buf, r(t, option))
}

type stringerFunc func() string

func (r stringerFunc) String() string {
	return r()
}

func (r ShowIndentFunc[T]) Stringer(t T, option ShowOption) fmt.Stringer {
	return stringerFunc(func() string {
		return r.ShowIndent(t, option)
	})
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

func (r ShowFunc[T]) Append(buf []string, t T, option ShowOption) []string {
	return append(buf, r(t))
}

func (r ShowFunc[T]) Stringer(t T, option ShowOption) fmt.Stringer {
	return stringerFunc(func() string {
		return r.Show(t)
	})
}

type ShowAppendFunc[T any] func(buf []string, t T, option ShowOption) []string

func (r ShowAppendFunc[T]) Show(t T) string {
	return r.ShowIndent(t, ShowOption{})
}

func (r ShowAppendFunc[T]) ShowIndent(t T, opt ShowOption) string {
	ret := r(nil, t, opt)
	return strings.Join(ret, "")
}

func (r ShowAppendFunc[T]) Append(buf []string, t T, option ShowOption) []string {
	return r(buf, t, option)
}

func (r ShowAppendFunc[T]) Stringer(t T, option ShowOption) fmt.Stringer {
	return stringerFunc(func() string {
		return r.ShowIndent(t, option)
	})
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
