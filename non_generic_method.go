//go:build !go1.27

package fp

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
