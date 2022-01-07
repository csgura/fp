package fp

import "fmt"

type Option[T any] struct {
	v *T
}

func (r Option[T]) Foreach(f func(v T)) {
	if r.IsDefined() {
		f(*r.v)
	}
}

func (r Option[T]) IsDefined() bool {
	return r.v != nil
}

func (r Option[T]) IsEmpty() bool {
	return !r.IsDefined()
}

func (r Option[T]) Get() T {
	return *r.v
}

func (r Option[T]) String() string {
	if r.IsDefined() {
		return fmt.Sprintf("Some(%v)", r.v)
	} else {
		return "None"
	}
}

func (r Option[T]) Filter(p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if p(*r.v) {
			return r
		}
	}
	return Option[T]{}

}

func (r Option[T]) OrElse(t T) T {
	if r.IsDefined() {
		return *r.v
	}
	return t
}
func (r Option[T]) OrElseGet(f func() T) T {
	if r.IsDefined() {
		return *r.v
	}
	return f()
}
func (r Option[T]) Or(f func() Option[T]) Option[T] {
	if r.IsDefined() {
		return r
	}
	return f()
}

func (r Option[T]) Recover(f func() T) Option[T] {
	if r.IsDefined() {
		return r
	}
	t := f()
	return Option[T]{&t}
}

func (r Option[T]) Exists(p func(v T) bool) bool {
	if r.IsDefined() {
		return p(*r.v)
	}
	return false
}
func (r Option[T]) ForAll(p func(v T) bool) bool {
	if r.IsDefined() {
		return p(*r.v)
	}
	return true
}

func (r Option[T]) ToSeq() Seq[T] {
	if r.IsDefined() {
		return Seq[T]{*r.v}
	}
	return nil
}

func (r Option[T]) Iterator() Iterator[T] {
	return MakeIterator(
		r.IsDefined,
		r.Get,
	)
}

func Some[T any](v T) Option[T] {
	return Option[T]{&v}
}
