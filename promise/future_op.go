package promise

import (
	"fmt"

	"github.com/csgura/fp"
)

type goExecutor struct{}

func (r goExecutor) ExecuteUnsafe(runnable fp.Runnable) {
	go runnable.Run()
}

func getExecuter(ctx ...fp.Executor) fp.Executor {
	if len(ctx) == 0 {
		return goExecutor{}
	}
	return ctx[0]
}

type future[T any] struct {
	p *promise[T]
}

func (r future[T]) String() string {
	v := r.p.Value()
	if v.IsDefined() {
		return fmt.Sprintf("fp.Future(%v)", v.Get())
	}
	return fmt.Sprintf("fp.Future[%s](not completed)", fp.TypeName[T]())
}
func (r future[T]) OnFailure(cb func(err error), ctx ...fp.Executor) {
	r.OnComplete(func(try fp.Try[T]) {
		if !try.IsSuccess() {
			cb(try.Failed().Get())
		}
	}, ctx...)
}

func (r future[T]) OnSuccess(cb func(success T), ctx ...fp.Executor) {
	r.OnComplete(func(try fp.Try[T]) {
		if try.IsSuccess() {
			cb(try.Get())
		}
	}, ctx...)

}

func (r future[T]) Foreach(f func(v T), ctx ...fp.Executor) {
	r.OnSuccess(f, ctx...)
}

func (r future[T]) OnComplete(cb func(try fp.Try[T]), ctx ...fp.Executor) {
	r.p.dispatchOrAddCallback(func(t fp.Try[T]) {
		getExecuter(ctx...).ExecuteUnsafe(fp.RunnableFunc(func() {
			cb(t)
		}))
	})
}

func (r future[T]) IsCompleted() bool {
	return r.p.IsCompleted()
}

func (r future[T]) Value() fp.Option[fp.Try[T]] {
	return r.p.Value()
}

func (r future[T]) Failed() fp.Future[error] {
	np := New[error]()

	r.OnComplete(func(t fp.Try[T]) {
		if t.IsSuccess() {
			np.Failure(fp.ErrFutureNotFailed)
		} else {
			np.Success(t.Failed().Get())
		}
	})

	return np.Future()
}

func (r future[T]) Recover(f func(err error) T, ctx ...fp.Executor) fp.Future[T] {
	np := New[T]()

	r.OnComplete(func(t fp.Try[T]) {
		if t.IsSuccess() {
			np.Success(t.Get())
		} else {
			np.Success(f(t.Failed().Get()))
		}
	}, ctx...)

	return np.Future()
}

func (r future[T]) RecoverWith(f func(err error) fp.Future[T], ctx ...fp.Executor) fp.Future[T] {
	np := New[T]()

	r.OnComplete(func(t fp.Try[T]) {
		if t.IsSuccess() {
			np.Success(t.Get())
		} else {
			f(t.Failed().Get()).OnComplete(func(t fp.Try[T]) {
				np.Complete(t)
			}, ctx...)
		}
	}, ctx...)

	return np.Future()
}
