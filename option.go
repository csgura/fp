package fp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ptr[T any] = *T

type PtrT[T any] = Try[Ptr[T]]

type OptionT[T any] = Try[Option[T]]

type Option[T any] struct {
	present bool
	v       T
}

type GoIter[V any] = func(yield func(V) bool)

func (r Option[T]) All() GoIter[T] {
	return func(f func(T) bool) {
		if r.IsDefined() {
			f(r.Get())
		}
	}
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

func (r Option[T]) Unapply() (T, bool) {
	if r.IsDefined() {
		return r.Get(), true
	} else {
		var zero T
		return zero, false
	}
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

func (r Option[T]) Map(mf func(T) T) Option[T] {
	if r.IsDefined() {
		r.v = mf(r.v)
	}
	return r
}

func (r Option[T]) FlatMap(mf func(T) Option[T]) Option[T] {
	if r.IsDefined() {
		return mf(r.v)
	}
	return r
}

func (r Option[T]) FilterNot(p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if !p(r.Get()) {
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

func (r Option[T]) OrZero() T {
	return r.OrElseGet(Zero[T])
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

func (r Option[T]) OrOption(v Option[T]) Option[T] {
	if r.IsDefined() {
		return r
	}
	return v
}

func (r Option[T]) OrPtr(v *T) Option[T] {
	if r.IsDefined() {
		return r
	}
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func (r Option[T]) Recover(f func() T) Option[T] {
	if r.IsDefined() {
		return r
	}
	t := f()
	return Option[T]{true, t}
}

// func (r Option[T]) Exists(p func(v T) bool) bool {
// 	if r.IsDefined() {
// 		return p(r.Get())
// 	}
// 	return false
// }
// func (r Option[T]) ForAll(p func(v T) bool) bool {
// 	if r.IsDefined() {
// 		return p(r.Get())
// 	}
// 	return true
// }

func (r Option[T]) ToSeq() []T {
	if r.IsDefined() {
		return []T{r.Get()}
	}
	return nil
}

// func (r Option[T]) Iterator() Iterator[T] {
// 	return r.ToSeq().Iterator()
// }

func Some[T any](v T) Option[T] {
	return Option[T]{true, v}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (r *Option[T]) UnmarshalJSON(b []byte) error {
	if r == nil {
		return Error(http.StatusBadRequest, "target ptr is nil")
	}
	if len(b) > 0 {
		if b[0] != 'n' {
			var t T
			err := json.Unmarshal(b, &t)
			if err == nil {
				*r = Some(t)
			}
			return err
		}
	}
	*r = None[T]()

	return nil
}

func (r Option[T]) MarshalJSON() ([]byte, error) {
	if r.IsDefined() {
		return json.Marshal(r.Get())
	}

	return []byte("null"), nil
}

func (r Option[T]) Ptr() *T {
	if r.IsDefined() {
		return &r.v
	}

	return nil
}

func (r Option[T]) Exists(p func(v T) bool) bool {
	return r.IsDefined() && p(r.v)
}

func (r Option[T]) ForAll(p func(v T) bool) bool {
	return r.IsEmpty() || p(r.v)
}
