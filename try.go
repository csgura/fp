package fp

import (
	"fmt"
	"net/http"
)

type Try[T any] struct {
	v   *T
	err error
}

func Success[T any](t T) Try[T] {
	return Try[T]{&t, nil}
}

func Failure[T any](err error) Try[T] {
	return Try[T]{nil, err}
}

func (r Try[T]) IsSuccess() bool {
	return r.v != nil
}

func (r Try[T]) IsFailure() bool {
	return !r.IsSuccess()
}

func (r Try[T]) Get() T {
	return *r.v
}

func (r Try[T]) Unapply() (T, error) {
	if r.IsSuccess() {
		return *r.v, nil
	} else {
		var zero T
		var err = r.err

		if err == nil {
			err = Error(http.StatusNotAcceptable, "Try not initialized correctly")
		}

		return zero, err
	}
}

func (r Try[T]) Foreach(f func(v T)) {
	if r.IsSuccess() {
		f(*r.v)
	}
}
func (r Try[T]) Failed() Try[error] {
	if r.IsSuccess() {
		return Success(ErrTryNotFailed)
	}
	return Success(r.err)
}
func (r Try[T]) OrElse(t T) T {
	if r.IsSuccess() {
		return *r.v
	}
	return t
}

func (r Try[T]) OrElseGet(f func() T) T {
	if r.IsSuccess() {
		return *r.v
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
	return Success(f(r.err))

}
func (r Try[T]) RecoverWith(f func(err error) Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return f(r.err)
}

func (r Try[T]) ToOption() Option[T] {
	if r.IsSuccess() {
		return Option[T]{r.v}
	}
	return Option[T]{}
}

func (r Try[T]) ToSeq() Seq[T] {
	if r.IsSuccess() {
		return Seq[T]{*r.v}
	}
	return nil
}

func (r Try[T]) String() string {
	if r.IsSuccess() {
		return fmt.Sprintf("Success(%v)", r.Get())
	}
	return fmt.Sprintf("Failure(%v)", r.err)
}

func (r Try[T]) Iterator() Iterator[T] {
	return MakeIterator(
		r.IsSuccess,
		r.Get,
	)
}
