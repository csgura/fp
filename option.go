package fp

import (
	"encoding/json"
	"fmt"
)

type Option[T any] struct {
	present bool
	v       T
}

func (r Option[T]) Foreach(f func(v T)) {
	if r.IsDefined() {
		f(r.Get())
	}
}

func (r Option[T]) IsDefined() bool {
	return r.present
}

func (r Option[T]) IsEmpty() bool {
	return !r.IsDefined()
}

func (r Option[T]) Get() T {
	if r.IsDefined() {
		return r.v
	}
	panic("Option.empty")
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
		if p(r.Get()) {
			return r
		}
	}
	return None[T]()

}

func (r Option[T]) OrElse(t T) T {
	if r.IsDefined() {
		return r.Get()
	}
	return t
}
func (r Option[T]) OrElseGet(f func() T) T {
	if r.IsDefined() {
		return r.Get()
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
	return Option[T]{true, t}
}

func (r Option[T]) Exists(p func(v T) bool) bool {
	if r.IsDefined() {
		return p(r.Get())
	}
	return false
}
func (r Option[T]) ForAll(p func(v T) bool) bool {
	if r.IsDefined() {
		return p(r.Get())
	}
	return true
}

func (r Option[T]) ToSeq() Seq[T] {
	if r.IsDefined() {
		return Seq[T]{r.Get()}
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
	return Option[T]{true, v}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (r *Option[T]) UnmarshalJSON(b []byte) error {
	if len(b) > 0 {
		if b[0] != 'n' {
			err := json.Unmarshal(b, &r.v)
			if err == nil {
				r.present = true
			}
			return err
		}
	}
	return nil
}

func (r Option[T]) MarshalJSON() ([]byte, error) {
	if r.IsDefined() {
		return json.Marshal(r.Get())
	}

	return []byte("null"), nil
}
