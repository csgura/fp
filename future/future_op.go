//go:generate go run github.com/csgura/fp/internal/generator/future_gen
package future

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
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
				p.Failure(fp.PanicError(err))
			}
		}()

		result := f()
		p.Success(result)
	}))

	return p.Future()
}

func Apply2[T any](f func() (T, error), ctx ...fp.ExecContext) fp.Future[T] {
	p := promise.New[T]()

	getExecuter(ctx...).Execute(fp.RunnableFunc(func() {
		defer func() {
			if err := recover(); err != nil {
				p.Failure(fp.PanicError(err))
			}
		}()

		result, err := f()
		if err != nil {
			p.Failure(err)
		} else {
			p.Success(result)
		}
	}))

	return p.Future()
}

func FromOption[T any](v fp.Option[T]) fp.Future[T] {
	if v.IsDefined() {
		return Successful(v.Get())
	} else {
		return Failed[T](fp.ErrOptionEmpty)
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
		return Map(a, f)
	}, ctx...)
}

func Map[T, U any](opt fp.Future[T], f func(v T) U, ctx ...fp.ExecContext) fp.Future[U] {
	return FlatMap(opt, func(v T) fp.Future[U] {
		return Successful(f(v))
	}, ctx...)
}

func Map2[A, B, U any](a fp.Future[A], b fp.Future[B], f func(A, B) U, ctx ...fp.ExecContext) fp.Future[U] {
	return FlatMap(a, func(v1 A) fp.Future[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		}, ctx...)
	}, ctx...)
}

func Lift[T, U any](f func(v T) U, ctx ...fp.ExecContext) fp.Func1[fp.Future[T], fp.Future[U]] {
	return func(opt fp.Future[T]) fp.Future[U] {
		return Map(opt, f, ctx...)
	}
}

func LiftA2[A1, A2, R any](f fp.Func2[A1, A2, R], ctx ...fp.ExecContext) fp.Func2[fp.Future[A1], fp.Future[A2], fp.Future[R]] {
	return func(a1 fp.Future[A1], a2 fp.Future[A2]) fp.Future[R] {
		return Ap(Ap(Successful(f.Curried()), a1, ctx...), a2, ctx...)
	}
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.Future[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.ExecContext) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(f1(a), f2, ctx...)
	}
}

func Compose2[A, B, C any](f1 fp.Func1[A, fp.Future[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.ExecContext) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(f1(a), f2, ctx...)
	}
}

func ComposeOption[A, B, C any](f1 fp.Func1[A, fp.Option[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.ExecContext) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(FromOption(f1(a)), f2, ctx...)
	}
}

func ComposeTry[A, B, C any](f1 fp.Func1[A, fp.Try[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.ExecContext) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(FromTry(f1(a)), f2, ctx...)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.Future[B]], f2 fp.Func1[B, C], ctx ...fp.ExecContext) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return Map(f1(a), f2, ctx...)
	}
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

func Zip[A, B any](c1 fp.Future[A], c2 fp.Future[B]) fp.Future[fp.Tuple2[A, B]] {
	return FlatMap(c1, func(v1 A) fp.Future[fp.Tuple2[A, B]] {
		return Map(c2, func(v2 B) fp.Tuple2[A, B] {
			return product.Tuple2(v1, v2)
		})
	})
}

func Zip3[A, B, C any](c1 fp.Future[A], c2 fp.Future[B], c3 fp.Future[C]) fp.Future[fp.Tuple3[A, B, C]] {
	return Applicative3(as.Tuple3[A, B, C]).
		ApFuture(c1).
		ApFuture(c2).
		ApFuture(c3)
}

func Sequence[T any](futureList fp.Seq[fp.Future[T]], ctx ...fp.ExecContext) fp.Future[fp.Seq[T]] {
	head, tail := futureList.UnSeq()
	if head.IsDefined() {
		return FlatMap(head.Get(), func(headResult T) fp.Future[fp.Seq[T]] {
			last := Sequence(tail, ctx...)
			return Map(last, func(tail fp.Seq[T]) fp.Seq[T] {
				return seq.Concat(headResult, tail)
			}, ctx...)

		}, ctx...)
	}
	return Successful(seq.Of[T]())
}

type ApplicativeFunctor1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[H, HT, A, R]) Map(a func(HT) A, ctx ...fp.ExecContext) fp.Future[R] {
	return r.FlatMap(func(h HT) fp.Future[A] {
		return Successful(a(h))
	}, ctx...)
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListMap(a func(H) A, ctx ...fp.ExecContext) fp.Future[R] {
	return r.HListFlatMap(func(h H) fp.Future[A] {
		return Successful(a(h))
	}, ctx...)
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListFlatMap(a func(H) fp.Future[A], ctx ...fp.ExecContext) fp.Future[R] {
	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) FlatMap(a func(HT) fp.Future[A], ctx ...fp.ExecContext) fp.Future[R] {
	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a(v.Head())
	}, ctx...)

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

func (r ApplicativeFunctor1[H, HT, A, R]) ApFutureFunc(a func() fp.Future[A], ctx ...fp.ExecContext) fp.Future[R] {

	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor1[H, HT, A, R]) ApTryFunc(a func() fp.Try[A], ctx ...fp.ExecContext) fp.Future[R] {

	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A], ctx ...fp.ExecContext) fp.Future[R] {

	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor1[H, HT, A, R]) ApFunc(a func() A, ctx ...fp.ExecContext) fp.Future[R] {

	av := Map(r.h, func(v H) A {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R] {
	return ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R]{Successful(hlist.Empty()), Successful(fn)}
}

func Await[T any](future fp.Future[T], timeout time.Duration) fp.Try[T] {
	ch := make(chan fp.Try[T], 1)

	timer := time.AfterFunc(timeout, func() {
		ch <- try.Failure[T](fp.Error(408, "future not completed within %s", timeout))
	})

	future.OnComplete(func(r fp.Try[T]) {
		timer.Stop()
		ch <- r
	})

	return <-ch
}

func Func0[R any](f func() (R, error), ctx ...fp.ExecContext) fp.Func1[fp.Unit, fp.Future[R]] {
	return func(fp.Unit) fp.Future[R] {
		return Apply2(f, ctx...)
	}
}
