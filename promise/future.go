package promise

import (
	"fmt"

	"github.com/csgura/fp"
)

type goExecuter struct{}

func (r goExecuter) Execute(runnable fp.Runnable) {
	go runnable.Run()
}

func getExecuter(ctx ...fp.ExecContext) fp.ExecContext {
	if len(ctx) == 0 {
		return goExecuter{}
	}
	return ctx[0]
}

type future[T any] struct {
	p *promise[T]
}

func (r future[T]) OnFailure(cb func(err error), ctx ...fp.ExecContext) {
	r.OnComplete(func(try fp.Try[T]) {
		if !try.IsSuccess() {
			cb(try.Failed().Get())
		}
	}, ctx...)
}

func (r future[T]) OnSuccess(cb func(success T), ctx ...fp.ExecContext) {
	r.OnComplete(func(try fp.Try[T]) {
		if try.IsSuccess() {
			cb(try.Get())
		}
	}, ctx...)

}

func (r future[T]) Foreach(f func(v T), ctx ...fp.ExecContext) {

}

func (r future[T]) OnComplete(cb func(try fp.Try[T]), ctx ...fp.ExecContext) {
	r.p.dispatchOrAddCallback(func(t fp.Try[T]) {
		getExecuter(ctx...).Execute(fp.RunnableFunc(func() {
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
			np.Failure(fmt.Errorf("future.Failed not completed with a error"))
		} else {
			np.Success(t.Failed().Get())
		}
	})

	return np.Future()
}

func (r future[T]) Recover(f func(err error) T, ctx ...fp.ExecContext) fp.Future[T] {
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

func (r future[T]) RecoverWith(f func(err error) fp.Future[T], ctx ...fp.ExecContext) fp.Future[T] {
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
