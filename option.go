package fp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ref[T any] struct {
	ref *T
}

var ErrInvalidRef = Error(http.StatusBadRequest, "fp.Ref not initialized correctly")

func (r Ref[T]) Get() T {
	if r.ref == nil {
		panic(ErrInvalidRef)
	}
	return *r.ref
}

func RefOf[T any](v *T) Ref[T] {
	return Ref[T]{
		ref: v,
	}
}

type Ptr[T any] = *T

type PtrT[T any] = Try[Ptr[T]]

type OptionT[T any] = Try[Option[T]]

type Option[T any] struct {
	present bool
	v       T
}

type GoIter[V any] = func(yield func(V) bool)

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
	if r.present {
		return json.Marshal(r.v)
	}

	return []byte("null"), nil
}

func (r Option[T]) Unapply() (T, bool) {
	if r.present {
		return r.v, true
	} else {
		var zero T
		return zero, false
	}
}

func (r Option[T]) Recover(f func() T) Option[T] {
	if r.present {
		return r
	}
	t := f()
	return Option[T]{true, t}
}

func (r Option[T]) String() string {
	if r.present {
		return fmt.Sprintf("Some(%v)", r.v)
	} else {
		return "None"
	}
}
