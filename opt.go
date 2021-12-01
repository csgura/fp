package fp

import "fmt"

type Some[T any] struct {
	v T
}

func (r Some[T]) Foreach(f func(v T)) {
	f(r.v)
}

func (r Some[T]) IsDefined() bool {
	return true
}

func (r Some[T]) Get() T {
	return r.v
}

func (r Some[T]) String() string {
	return fmt.Sprintf("Some(%v)", r.v)
}

func (r Some[T]) Filter(p func(v T) bool) Option[T] {

	if p(r.v) {
		return r
	}
	return None[T]{}

}

func (r Some[T]) OrElse(t T) T {
	return r.v
}
func (r Some[T]) OrElseGet(f func() T) T {
	return r.v
}
func (r Some[T]) Or(f func() Option[T]) Option[T] {
	return r
}

func (r Some[T]) Recover(f func() T) Option[T] {
	return r
}

type None[T any] struct{}

func (r None[T]) Foreach(f func(v T)) {

}

func (r None[T]) IsDefined() bool {
	return false
}

func (r None[T]) Get() T {
	panic("Option.empty")
}

func (r None[T]) String() string {
	return "None"
}

func (r None[T]) Filter(p func(v T) bool) Option[T] {
	return r
}

func (r None[T]) OrElse(t T) T {
	return t
}
func (r None[T]) OrElseGet(f func() T) T {
	return f()
}
func (r None[T]) Or(f func() Option[T]) Option[T] {
	return f()
}

func (r None[T]) Recover(f func() T) Option[T] {
	return Some[T]{f()}
}
