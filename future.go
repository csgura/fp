package fp

import (
	"fmt"

	"github.com/csgura/fp/internal/atomic"
)

type Runnable interface {
	Run()
}

// RunnableFunc is converter which converts function to Runnable interface
type RunnableFunc func()

// Run is Runnable.Run
func (r RunnableFunc) Run() {
	r()
}

type Executor interface {
	ExecuteUnsafe(runnable Runnable)
}

// type Promise[T any] interface {
// 	Future() Future[T]
// 	Success(value T) bool
// 	Failure(err error) bool
// 	IsCompleted() bool
// 	Complete(result Try[T]) bool
// }

func NewPromise[T any]() Promise[T] {
	return Promise[T]{atomic.New()}
}

type Promise[T any] struct {
	status atomic.Reference
}

func (r Promise[T]) Future() Future[T] {
	return Future[T]{r}
}

func (r Promise[T]) Success(value T) bool {
	return r.Complete(Success(value))
}

func (r Promise[T]) Failure(err error) bool {
	return r.Complete(Failure[T](err))
}

func (r Promise[T]) Value() Try[T] {
	if r.status == nil {
		panic("Promise not initalized correctly")
	}
	switch v := r.status.Load().(type) {
	case Try[T]:
		return v
	}
	panic("Promise not completed")
}

func (r Promise[T]) IsCompleted() bool {
	if r.status == nil {
		return false
	}

	switch r.status.Load().(type) {
	case Try[T]:
		return true
	}
	return false
}
func (r Promise[T]) Complete(result Try[T]) bool {
	if r.status == nil {
		return false
	}

	ret, cbs := r.tryCompleteAndGetListeners(result)
	for _, cf := range cbs {
		cf(result)
	}
	return ret
}

type onCompleteFunc[T any] func(t Try[T])

func (r Promise[T]) tryCompleteAndGetListeners(v Try[T]) (bool, []onCompleteFunc[T]) {
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

	case Try[T]:
		return false, nil
	}
	panic("not possible")
}

func (r Promise[T]) dispatchOrAddCallback(cb onCompleteFunc[T]) {
	if r.status == nil {
		return
	}

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

	case Try[T]:
		cb(status)
	}
}

type Future[T any] struct {
	p Promise[T]
}

func (r Future[T]) String() string {
	if r.p.IsCompleted() {
		v := r.p.Value()
		return fmt.Sprintf("Future(%v)", v)
	} else {
		return fmt.Sprintf("Future[%s](not completed)", TypeName[T]())
	}

}
func (r Future[T]) OnFailure(cb func(err error), ctx ...Executor) {
	r.OnComplete(func(try Try[T]) {
		if !try.IsSuccess() {
			cb(try.Failed().Get())
		}
	}, ctx...)
}

func (r Future[T]) OnSuccess(cb func(success T), ctx ...Executor) {
	r.OnComplete(func(try Try[T]) {
		if try.IsSuccess() {
			cb(try.Get())
		}
	}, ctx...)

}

func (r Future[T]) Foreach(f func(v T), ctx ...Executor) {
	r.OnSuccess(f, ctx...)
}

type goExecutor struct{}

func (r goExecutor) ExecuteUnsafe(runnable Runnable) {
	go runnable.Run()
}

func getExecutor(ctx ...Executor) Executor {
	if len(ctx) == 0 || ctx[0] == nil {
		return goExecutor{}
	}
	return ctx[0]
}

func (r Future[T]) OnComplete(cb func(try Try[T]), ctx ...Executor) {
	r.p.dispatchOrAddCallback(func(t Try[T]) {
		getExecutor(ctx...).ExecuteUnsafe(RunnableFunc(func() {
			cb(t)
		}))
	})
}

func (r Future[T]) IsCompleted() bool {
	return r.p.IsCompleted()
}

func (r Future[T]) Value() Try[T] {
	if r.p.status == nil {
		panic("Future not completed")
	}
	switch v := r.p.status.Load().(type) {
	case Try[T]:
		return v
	}
	panic("Future not completed")
}

func (r Future[T]) Failed() Future[error] {
	np := NewPromise[error]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			np.Failure(ErrFutureNotFailed)
		} else {
			np.Success(t.Failed().Get())
		}
	})

	return np.Future()
}

func (r Future[T]) Recover(f func(err error) T, ctx ...Executor) Future[T] {
	np := NewPromise[T]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			np.Success(t.Get())
		} else {
			np.Success(f(t.Failed().Get()))
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) RecoverWith(f func(err error) Future[T], ctx ...Executor) Future[T] {
	np := NewPromise[T]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			np.Success(t.Get())
		} else {
			f(t.Failed().Get()).OnComplete(func(t Try[T]) {
				np.Complete(t)
			}, ctx...)
		}
	}, ctx...)

	return np.Future()
}

// type Future[T any] interface {
// 	OnFailure(cb func(err error), ctx ...Executor)
// 	OnSuccess(cb func(success T), ctx ...Executor)
// 	Foreach(f func(v T), ctx ...Executor)
// 	OnComplete(cb func(try Try[T]), ctx ...Executor)
// 	IsCompleted() bool
// 	Failed() Future[error]
// 	Recover(f func(err error) T, ctx ...Executor) Future[T]
// 	RecoverWith(f func(err error) Future[T], ctx ...Executor) Future[T]
// }
