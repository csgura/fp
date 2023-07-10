package fp

import (
	"fmt"
	"net/http"
)

type Try[T any] struct {
	success bool
	v       T
	err     error
}

func Success[T any](t T) Try[T] {
	return Try[T]{true, t, nil}
}

func Failure[T any](err error) Try[T] {
	var zero T
	return Try[T]{false, zero, err}
}

func (r Try[T]) IsSuccess() bool {
	return r.success
}

func (r Try[T]) IsFailure() bool {
	return !r.IsSuccess()
}

func (r Try[T]) Get() T {
	if r.IsSuccess() {
		return r.v
	}
	panic(r.Failed().Get())
}

func (r Try[T]) Unapply() (T, error) {
	if r.IsSuccess() {
		return r.Get(), nil
	} else {
		var zero T
		return zero, r.Failed().Get()
	}
}

func (r Try[T]) Map(mf func(T) T) Try[T] {
	if r.IsSuccess() {
		r.v = mf(r.v)
	}
	return r
}

func (r Try[T]) FlatMap(mf func(T) Try[T]) Try[T] {
	if r.IsSuccess() {
		return mf(r.v)
	}
	return r
}

func (r Try[T]) Foreach(f func(v T)) {
	if r.IsSuccess() {
		f(r.Get())
	}
}
func (r Try[T]) Failed() Try[error] {
	if r.IsSuccess() {
		return Failure[error](ErrTryNotFailed)
	}
	if r.err == nil {
		return Failure[error](Error(http.StatusNotAcceptable, "Try not initialized correctly"))
	}
	return Success(r.err)
}
func (r Try[T]) OrElse(t T) T {
	if r.IsSuccess() {
		return r.Get()
	}
	return t
}

func (r Try[T]) OrZero() T {
	return r.OrElseGet(Zero[T])
}

func (r Try[T]) OrElseGet(f func() T) T {
	if r.IsSuccess() {
		return r.Get()
	}
	return f()
}
func (r Try[T]) Or(f func() Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return f()
}

func (r Try[T]) Recover(f func(err error) T) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return Success(f(r.Failed().Get()))

}
func (r Try[T]) RecoverWith(f func(err error) Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return f(r.Failed().Get())
}

// func (r Try[T]) ToOption() Option[T] {
// 	if r.IsSuccess() {
// 		return Some(r.v)
// 	}
// 	return None[T]()
// }

func (r Try[T]) String() string {
	if r.IsSuccess() {
		return fmt.Sprintf("Success(%v)", r.Get())
	}
	return fmt.Sprintf("Failure(%v)", r.Failed().Get())
}

func (r Try[T]) ToSeq() []T {
	if r.IsSuccess() {
		return []T{r.Get()}
	}
	return nil
}

// func (r Try[T]) Iterator() Iterator[T] {
// 	return r.ToSeq().Iterator()
// }
