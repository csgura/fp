//go:build go1.27

package fp

func (r Option[T]) Map[R any](mf func(T) R) Option[R] {
	if r.IsDefined() {
		return Some(mf(r.v))
	}
	return None[R]()
}

func (r Option[T]) FlatMap[R any](mf func(T) Option[R]) Option[R] {
	if r.IsDefined() {
		return mf(r.v)
	}
	return None[R]()
}

func (r Option[T]) Replace[R any](o R) Option[R] {
	return r.Map(Const[T](o))
}

func (r Option[T]) ReplaceS[R any](f func() R) Option[R] {
	return r.Map(func(t T) R {
		return f()
	})
}

func (r Option[T]) Void[R any]() Option[Unit] {
	return r.Replace(Unit{})
}

func (r Option[T]) Map2[U, R any](other Option[U], f func(T, U) R) Option[R] {
	return r.FlatMap(func(t T) Option[R] {
		return other.Map(func(u U) R {
			return f(t, u)
		})
	})
}

func (r Option[T]) IntoTry[E error](err func() E) Try[T] {
	if r.IsDefined() {
		return Success(r.Get())
	}
	return Failure[T](err())
}

func (r Option[T]) IntoFuture[E error](err func() E) Future[T] {
	p := NewPromise[T]()
	if r.IsDefined() {
		p.Success(r.Get())
		return p.Future()
	}
	p.Failure(err())
	return p.Future()
}

func (r Option[T]) TraverseT[R any](f func(T) Try[R]) Try[Option[R]] {
	if r.IsDefined() {
		return f(r.Get()).Map(Some)
	}
	return Success(None[R]())
}

func (r Option[T]) TraverseF[R any](f func(T) Future[R]) Future[Option[R]] {
	if r.IsDefined() {
		return f(r.Get()).Map(Some)
	}
	p := NewPromise[Option[R]]()
	p.Success(None[R]())
	return p.Future()
}

func (r Try[T]) Map[R any](mf func(T) R) Try[R] {
	if r.IsSuccess() {
		return Success(mf(r.v))
	}
	return Failure[R](r.err)
}

func (r Try[T]) FlatMap[R any](mf func(T) Try[R]) Try[R] {
	if r.IsSuccess() {
		return mf(r.v)
	}
	return Failure[R](r.err)
}

func (r Try[T]) Replace[R any](o R) Try[R] {
	return r.Map(Const[T](o))
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
	if r.IsSuccess() {
		return Some(r.Get())
	}
	return None[T]()
}

func (r Try[T]) IntoFuture[_ Phantom[T]]() Future[T] {
	p := NewPromise[T]()
	if r.IsSuccess() {
		p.Success(r.Get())
		return p.Future()
	}
	p.Failure(r.Failed().Get())
	return p.Future()
}

func (r Try[T]) TraverseF[R any](f func(T) Future[R]) Future[R] {
	if r.IsSuccess() {
		return f(r.Get())
	}
	p := NewPromise[R]()
	p.Failure(r.Failed().Get())
	return p.Future()
}

func (r Future[T]) Map[R any](mf func(T) R, ctx ...Executor) Future[R] {
	np := NewPromise[R]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			np.Success(mf(t.Get()))
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) FlatMap[R any](mf func(T) Future[R], ctx ...Executor) Future[R] {
	np := NewPromise[R]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			mf(t.Get()).OnComplete(func(t Try[R]) {
				np.Complete(t)
			}, ctx...)
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) Replace[R any](o R) Future[R] {
	return r.Map(Const[T](o))
}

func (r Future[T]) ReplaceS[R any](f func() R) Future[R] {
	return r.Map(func(t T) R {
		return f()
	})
}

func (r Future[T]) Void[R any]() Future[Unit] {
	return r.Replace(Unit{})
}

func (r Seq[T]) Map[R any](mf func(T) R) Seq[R] {
	var ret = make([]R, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v))
	}
	return ret
}

func (r Seq[T]) FlatMap[R any](mf func(T) Seq[R]) Seq[R] {
	var ret = make([]R, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v)...)
	}
	return ret
}

func (r Seq[T]) Map2[U, R any](other Seq[U], f func(T, U) R) Seq[R] {
	return r.FlatMap(func(t T) Seq[R] {
		return other.Map(func(u U) R {
			return f(t, u)
		})
	})
}

func (r Seq[T]) FoldT[ACC any](zero ACC, f func(ACC, T) Try[ACC]) Try[ACC] {
	sum := zero
	for _, v := range r {
		t := f(sum, v)
		if t.IsSuccess() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return Success(sum)
}

func (r Seq[T]) Fold[ACC any](zero ACC, f func(ACC, T) ACC) ACC {
	sum := zero
	for _, v := range r {
		sum = f(sum, v)
	}
	return sum
}

func (r Seq[T]) FoldF[ACC any](zero ACC, f func(ACC, T) Future[ACC], ctx ...Executor) Future[ACC] {
	p := NewPromise[ACC]()
	p.Success(zero)

	return r.Fold(p.Future(), func(accf Future[ACC], t T) Future[ACC] {
		return accf.FlatMap(func(acc ACC) Future[ACC] {
			return f(acc, t)
		}, ctx...)
	})
}

func (r Seq[T]) TraverseT[R any](f func(T) Try[R]) Try[Seq[R]] {
	return r.FoldT(nil, func(a Seq[R], t T) Try[Seq[R]] {
		return f(t).Map(a.Add)
	})
}

func (r Seq[T]) TraverseF[R any](f func(T) Future[R], ctx ...Executor) Future[Seq[R]] {
	return r.FoldF(nil, func(acc Seq[R], t T) Future[Seq[R]] {
		return f(t).Map(acc.Add, ctx...)
	}, ctx...)
}
