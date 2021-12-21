package promise

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/atomic"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

type promise[T any] struct {
	status atomic.Value
}

func (r *promise[T]) Future() fp.Future[T] {
	return future[T]{r}
}

func (r *promise[T]) Success(value T) bool {
	return r.Complete(try.Success(value))
}

func (r *promise[T]) Failure(err error) bool {
	return r.Complete(try.Failure[T](err))
}

func (r promise[T]) Value() fp.Option[fp.Try[T]] {
	switch v := r.status.Load().(type) {
	case fp.Try[T]:
		return option.Some(v)
	}
	return option.None[fp.Try[T]]()
}

func (r *promise[T]) IsCompleted() bool {
	switch r.status.Load().(type) {
	case fp.Try[T]:
		return true
	}
	return false
}
func (r *promise[T]) Complete(result fp.Try[T]) bool {
	ret, cbs := r.tryCompleteAndGetListeners(result)
	for _, cf := range cbs {
		cf(result)
	}
	return ret
}

func New[T any]() fp.Promise[T] {
	return &promise[T]{}
}

type onCompleteFunc[T any] func(t fp.Try[T])

func (r *promise[T]) tryCompleteAndGetListeners(v fp.Try[T]) (bool, []onCompleteFunc[T]) {
	ap := r.status.Get()
	switch status := ap.Value().(type) {
	case nil:
		if r.status.CompareAndSwap(ap, v) {
			return true, nil
		}
		return r.tryCompleteAndGetListeners(v)

	case []onCompleteFunc[T]:
		if r.status.CompareAndSwap(ap, v) {
			return true, status
		}
		return r.tryCompleteAndGetListeners(v)

	case fp.Try[T]:
		return false, nil
	}
	panic("not possible")
}

func (r *promise[T]) dispatchOrAddCallback(cb onCompleteFunc[T]) {
	ap := r.status.Get()
	switch status := ap.Value().(type) {
	case nil:
		if r.status.CompareAndSwap(ap, []onCompleteFunc[T]{cb}) {
			return
		}
		r.dispatchOrAddCallback(cb)
		return

	case []onCompleteFunc[T]:
		if r.status.CompareAndSwap(ap, append(status, cb)) {
			return
		}
		r.dispatchOrAddCallback(cb)
		return

	case fp.Try[T]:
		cb(status)
	}
}
