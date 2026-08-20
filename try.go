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
	if err == nil {
		panic("Failure error is nil")
	}
	var zero T
	return Try[T]{false, zero, err}
}

func (r Try[T]) String() string {
	if r.success {
		return fmt.Sprintf("Success(%v)", r.v)
	}
	return fmt.Sprintf("Failure(%v)", r.err)
}

func (r Try[T]) All[_ Phantom[T]]() GoIter[T] {
	return func(f func(T) bool) {
		if r.success {
			f(r.Get())
		}
	}
}

func (r Try[T]) IsSuccess[_ Phantom[T]]() bool {
	return r.success
}

func (r Try[T]) IsFailure[_ Phantom[T]]() bool {
	return !r.success
}

func (r Try[T]) Get[_ Phantom[T]]() T {
	if r.success {
		return r.v
	}
	if r.err == nil {
		panic(Error(http.StatusNotAcceptable, "Try not initialized correctly"))
	}
	panic(ErrTryNotFailed)
}

func (r Try[T]) Unapply() (T, error) {
	if r.success {
		return r.v, nil
	} else if r.err == nil {
		var zero T
		return zero, Error(http.StatusNotAcceptable, "Try not initialized correctly")
	} else {
		var zero T
		return zero, r.err
	}
}

func (r Try[T]) MapError[_ Phantom[T]](mf func(error) error) Try[T] {
	if !r.success {
		r.err = mf(r.err)
	}
	return r
}

func (r Try[T]) Foreach[_ Phantom[T]](f func(v T)) {
	if r.success {
		f(r.v)
	}
}
func (r Try[T]) Failed[_ Phantom[T]]() Try[error] {
	if r.success {
		return Failure[error](ErrTryNotFailed)
	}
	if r.err == nil {
		return Failure[error](Error(http.StatusNotAcceptable, "Try not initialized correctly"))
	}
	return Success(r.err)
}

func (r Try[T]) OrElse[_ Phantom[T]](t T) T {
	if r.success {
		return r.v
	}
	return t
}

func (r Try[T]) OrZero[_ Phantom[T]]() T {
	if r.success {
		return r.v
	}
	var zero T
	return zero
}

func (r Try[T]) OrElseGet[_ Phantom[T]](f func() T) T {
	if r.success {
		return r.v
	}
	return f()
}

func (r Try[T]) Or[_ Phantom[T]](f func() Try[T]) Try[T] {
	if r.success {
		return r
	}
	return f()
}

func (r Try[T]) OrTry[_ Phantom[T]](v Try[T]) Try[T] {
	if r.success {
		return r
	}
	return v
}

func (r Try[T]) Recover[_ Phantom[T]](f func(err error) T) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}
	return Success(f(err))

}

func (r Try[T]) RecoverCase[_ Phantom[T]](isDefinedAt func(error) bool, then func(error) T) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}

	if isDefinedAt(err) {
		return Success(then(err))
	}

	return r
}

func (r Try[T]) RecoverCaseWith[_ Phantom[T]](isDefinedAt func(error) bool, then func(error) Try[T]) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}

	if isDefinedAt(err) {
		return then(err)
	}

	return r
}

func (r Try[T]) RecoverWith[_ Phantom[T]](f func(err error) Try[T]) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}
	return f(err)
}

func (r Try[T]) ToSeq[_ Phantom[T]]() []T {
	if r.success {
		return []T{r.v}
	}
	return nil
}

func (r Try[T]) Map[R any](mf func(T) R) Try[R] {
	if r.success {
		return Success(mf(r.v))
	}
	return Failure[R](r.err)
}

func (r Try[T]) FlatMap[R any](mf func(T) Try[R]) Try[R] {
	if r.success {
		return mf(r.v)
	}
	return Failure[R](r.err)
}

func (r Try[T]) Replace[R any](o R) Try[R] {
	_, err := r.Unapply()
	if err == nil {
		return Success(o)
	}
	return Failure[R](err)
}

func (r Try[T]) ReplaceS[R any](f func() R) Try[R] {
	return r.Map(func(t T) R {
		return f()
	})
}

func (r Try[T]) Void[R any]() Try[Unit] {
	return r.Replace(Unit{})
}

func (r Try[T]) Map2[U, R any](other Try[U], f func(T, U) R) Try[R] {
	return r.FlatMap(func(t T) Try[R] {
		return other.Map(func(u U) R {
			return f(t, u)
		})
	})
}

func (r Try[T]) IntoOption[_ Phantom[T]]() Option[T] {
	if r.success {
		return Some(r.v)
	}
	return None[T]()
}

func (r Try[T]) IntoFuture[_ Phantom[T]]() Future[T] {
	p := NewPromise[T]()
	if r.success {
		p.Success(r.v)
		return p.Future()
	}
	_, err := r.Unapply()
	p.Failure(err)
	return p.Future()
}

func (r Try[T]) TraverseF[R any](f func(T) Future[R]) Future[R] {
	if r.success {
		return f(r.v)
	}
	p := NewPromise[R]()
	_, err := r.Unapply()
	p.Failure(err)
	return p.Future()
}

func (r Try[T]) Fold[A any](zero A, f func(A, T) A) A {
	if r.success {
		return f(zero, r.v)
	}
	return zero
}

func (r Try[T]) Either[R any](errorcase func(error) R, successcase func(T) R) R {
	v, err := r.Unapply()
	if err == nil {
		return successcase(v)
	}
	return errorcase(err)
}
