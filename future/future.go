package future

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/promise"
)

func Successful[T any](v T) fp.Future[T] {
	p := promise.New[T]()
	p.Success(v)
	return p.Future()
}

func Failed[T any](err error) fp.Future[T] {
	p := promise.New[T]()
	p.Failure(err)
	return p.Future()
}

var Unit fp.Future[fp.Unit] = Successful(fp.Unit{})

func Ap[T, U any](t fp.Future[fp.Func1[T, U]], a fp.Future[T], ctx ...fp.ExecContext) fp.Future[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Future[U] {
		return Map(a.(fp.Future[T]), func(a T) U {
			return f(a)
		})
	}, ctx...)
}

func Map[T, U any](opt fp.Future[T], f func(v T) U, ctx ...fp.ExecContext) fp.Future[U] {
	return FlatMap(opt, func(v T) fp.Future[U] {
		return Successful(f(v))
	}, ctx...)
}

func FlatMap[T, U any](opt fp.Future[T], fn func(v T) fp.Future[U], ctx ...fp.ExecContext) fp.Future[U] {
	np := promise.New[U]()

	opt.OnComplete(func(t fp.Try[T]) {
		if t.IsSuccess() {
			fn(t.Get()).OnComplete(func(t fp.Try[U]) {
				np.Complete(t)
			}, ctx...)
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func Transform[T, U any](opt fp.Future[T], fn func(v fp.Try[T]) fp.Try[U], ctx ...fp.ExecContext) fp.Future[U] {
	np := promise.New[U]()

	opt.OnComplete(func(t fp.Try[T]) {
		np.Complete(fn(t))
	}, ctx...)

	return np.Future()
}

func TransformWith[T, U any](opt fp.Future[T], fn func(v fp.Try[T]) fp.Future[U], ctx ...fp.ExecContext) fp.Future[U] {
	np := promise.New[U]()

	opt.OnComplete(func(t fp.Try[T]) {
		fn(t).OnComplete(func(t fp.Try[U]) {
			np.Complete(t)
		}, ctx...)
	}, ctx...)

	return np.Future()
}

func Flatten[T any](opt fp.Future[fp.Future[T]]) fp.Future[T] {
	return FlatMap(opt, func(v fp.Future[T]) fp.Future[T] {
		return v
	})
}
