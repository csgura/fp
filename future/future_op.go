//go:generate go run github.com/csgura/fp/internal/generator/future_gen
package future

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/seq"
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

func Apply[T any](f func() T, ctx ...fp.ExecContext) fp.Future[T] {
	p := promise.New[T]()

	getExecuter(ctx...).Execute(fp.RunnableFunc(func() {
		defer func() {
			if err := recover(); err != nil {
				p.Failure(fmt.Errorf("panic occurred : %v", err))
			}
		}()

		result := f()
		p.Success(result)
	}))

	return p.Future()
}

func FromOption[T any](v fp.Option[T]) fp.Future[T] {
	if v.IsDefined() {
		return Successful(v.Get())
	} else {
		return Failed[T](fmt.Errorf("Option.empty"))
	}
}

func FromTry[T any](v fp.Try[T]) fp.Future[T] {
	if v.IsSuccess() {
		return Successful(v.Get())
	} else {
		return Failed[T](v.Failed().Get())
	}
}

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

func Sequence[T any](futureList fp.Seq[fp.Future[T]], ctx ...fp.ExecContext) fp.Future[fp.Seq[T]] {
	head := futureList.Head()
	if head.IsDefined() {
		return FlatMap(head.Get(), func(headResult T) fp.Future[fp.Seq[T]] {
			last := Sequence(futureList.Tail(), ctx...)
			return Map(last, func(tail fp.Seq[T]) fp.Seq[T] {
				return seq.Concact(headResult, tail)
			}, ctx...)

		}, ctx...)
	}
	return Successful(seq.Of[T]())
}

type ApplicativeFunctor1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[H, HT, A, R]) Map(a func(HT) A) fp.Future[R] {
	return r.FlatMap(func(h HT) fp.Future[A] {
		return Successful(a(h))
	})
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListMap(a func(H) A) fp.Future[R] {
	return r.HListFlatMap(func(h H) fp.Future[A] {
		return Successful(a(h))
	})
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListFlatMap(a func(H) fp.Future[A]) fp.Future[R] {
	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a(v)
	})

	return r.ApFuture(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) FlatMap(a func(HT) fp.Future[A]) fp.Future[R] {
	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a(v.Head())
	})

	return r.ApFuture(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Future[R] {
	return r.ApFuture(FromOption(a))
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApTry(a fp.Try[A]) fp.Future[R] {
	return r.ApFuture(FromTry(a))
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApFuture(a fp.Future[A]) fp.Future[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[H, HT, A, R]) Ap(a A) fp.Future[R] {
	return r.ApFuture(Successful(a))
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R] {
	return ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R]{Successful(hlist.Empty()), Successful(fn)}
}
