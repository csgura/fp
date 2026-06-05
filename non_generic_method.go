//go:build !go1.27

package fp

import "net/http"

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

func (r Future[T]) Map(mf func(T) T, ctx ...Executor) Future[T] {
	np := NewPromise[T]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			np.Success(mf(t.Get()))
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) FlatMap(mf func(T) Future[T], ctx ...Executor) Future[T] {
	np := NewPromise[T]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			mf(t.Get()).OnComplete(func(t Try[T]) {
				np.Complete(t)
			}, ctx...)
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func (r Seq[T]) Map(mf func(T) T) Seq[T] {
	var ret = make([]T, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v))
	}
	return ret
}

func (r Seq[T]) FlatMap(mf func(T) Seq[T]) Seq[T] {
	var ret = make([]T, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v)...)
	}
	return ret
}

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
	panic(ErrOptionEmpty)
}

func (r Option[T]) Filter(p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if p(r.Get()) {
			return r
		}
	}
	return None[T]()

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

func (r Option[T]) ToSeq() []T {
	if r.IsDefined() {
		return []T{r.Get()}
	}
	return nil
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

func (r Try[T]) All() GoIter[T] {
	return func(f func(T) bool) {
		if r.IsSuccess() {
			f(r.Get())
		}
	}
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

func (r Try[T]) MapError(mf func(error) error) Try[T] {
	if r.IsFailure() {
		r.err = mf(r.err)
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

func (r Try[T]) OrTry(v Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return v
}

func (r Try[T]) Recover(f func(err error) T) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return Success(f(r.Failed().Get()))

}

func (r Try[T]) RecoverCase(isDefinedAt func(error) bool, then func(error) T) Try[T] {
	if r.IsSuccess() {
		return r
	}

	if isDefinedAt(r.Failed().Get()) {
		return Success(then(r.Failed().Get()))
	}

	return r
}

func (r Try[T]) RecoverCaseWith(isDefinedAt func(error) bool, then func(error) Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}

	if isDefinedAt(r.Failed().Get()) {
		return then(r.Failed().Get())
	}

	return r
}

func (r Try[T]) RecoverWith(f func(err error) Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return f(r.Failed().Get())
}

func (r Try[T]) ToSeq() []T {
	if r.IsSuccess() {
		return []T{r.Get()}
	}
	return nil
}
