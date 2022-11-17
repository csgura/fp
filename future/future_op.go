//go:generate go run github.com/csgura/fp/internal/generator/future_gen
package future

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/try"
)

type goExecutor struct{}

func (r goExecutor) ExecuteUnsafe(runnable fp.Runnable) {
	go runnable.Run()
}

func getExecutor(ctx ...fp.Executor) fp.Executor {
	if len(ctx) == 0 || ctx[0] == nil {
		return goExecutor{}
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

func Apply[T any](f func() T, ctx ...fp.Executor) fp.Future[T] {
	p := promise.New[T]()

	getExecutor(ctx...).ExecuteUnsafe(fp.RunnableFunc(func() {
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

func Apply2[T any](f func() (T, error), ctx ...fp.Executor) fp.Future[T] {
	p := promise.New[T]()

	getExecutor(ctx...).ExecuteUnsafe(fp.RunnableFunc(func() {
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

func Ap[T, U any](t fp.Future[fp.Func1[T, U]], a fp.Future[T], ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Future[U] {
		return Map(a, f)
	}, ctx...)
}

func ApFunc[T, U any](t fp.Future[fp.Func1[T, U]], a func() fp.Future[T], ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Future[U] {
		return Map(a(), f)
	}, ctx...)
}

func Map[T, U any](opt fp.Future[T], f func(v T) U, ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(opt, func(v T) fp.Future[U] {
		return Successful(f(v))
	}, ctx...)
}

func Map2[A, B, U any](a fp.Future[A], b fp.Future[B], f func(A, B) U, ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(a, func(v1 A) fp.Future[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		}, ctx...)
	}, ctx...)
}

func Lift[T, U any](f func(v T) U, ctx ...fp.Executor) fp.Func1[fp.Future[T], fp.Future[U]] {
	return func(opt fp.Future[T]) fp.Future[U] {
		return Map(opt, f, ctx...)
	}
}

func LiftA2[A1, A2, R any](f fp.Func2[A1, A2, R], ctx ...fp.Executor) fp.Func2[fp.Future[A1], fp.Future[A2], fp.Future[R]] {
	return func(a1 fp.Future[A1], a2 fp.Future[A2]) fp.Future[R] {
		return Map2(a1, a2, f, ctx...)
	}
}

func LiftM[A, R any](fa func(v A) fp.Future[R], ctx ...fp.Executor) fp.Func1[fp.Future[A], fp.Future[R]] {
	return func(ta fp.Future[A]) fp.Future[R] {
		return Flatten(Map(ta, fa, ctx...))
	}
}

// (a -> b -> m r) -> m a -> m b -> m r
// 하스켈에서는  liftM2 와 liftA2 는 같은 함수이고
// 위와 같은 함수는 존재하지 않음.
// hoogle 에서 검색해 보면 , liftJoin2 , bindM2 등의 이름으로 정의된 것이 있음.
// 하지만 ,  fp 패키지에서도   LiftA2 와 LiftM2 를 동일하게 하는 것은 낭비이고
// M 은 Monad 라는 뜻인데, Monad는 Flatten, FlatMap 의 의미가 있으니까
// LiftM2 를 다음과 같이 정의함.
func LiftM2[A, B, R any](fab fp.Func2[A, B, fp.Future[R]], ctx ...fp.Executor) fp.Func2[fp.Future[A], fp.Future[B], fp.Future[R]] {
	return func(a fp.Future[A], b fp.Future[B]) fp.Future[R] {
		return Flatten(Map2(a, b, fab, ctx...))
	}
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.Future[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.Executor) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(f1(a), f2, ctx...)
	}
}

func Compose2[A, B, C any](f1 fp.Func1[A, fp.Future[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.Executor) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(f1(a), f2, ctx...)
	}
}

func ComposeOption[A, B, C any](f1 fp.Func1[A, fp.Option[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.Executor) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(FromOption(f1(a)), f2, ctx...)
	}
}

func ComposeTry[A, B, C any](f1 fp.Func1[A, fp.Try[B]], f2 fp.Func1[B, fp.Future[C]], ctx ...fp.Executor) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return FlatMap(FromTry(f1(a)), f2, ctx...)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.Future[B]], f2 fp.Func1[B, C], ctx ...fp.Executor) fp.Func1[A, fp.Future[C]] {
	return func(a A) fp.Future[C] {
		return Map(f1(a), f2, ctx...)
	}
}

func FlatMap[T, U any](opt fp.Future[T], fn func(v T) fp.Future[U], ctx ...fp.Executor) fp.Future[U] {
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

func Transform[T, U any](opt fp.Future[T], fn func(v fp.Try[T]) fp.Try[U], ctx ...fp.Executor) fp.Future[U] {
	np := promise.New[U]()

	opt.OnComplete(func(t fp.Try[T]) {
		np.Complete(fn(t))
	}, ctx...)

	return np.Future()
}

func TransformWith[T, U any](opt fp.Future[T], fn func(v fp.Try[T]) fp.Future[U], ctx ...fp.Executor) fp.Future[U] {
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

// 하스켈 : m( a -> r ) -> a -> m r
// 스칼라 : M[ A => r ] => A => M[R]
// 하스켈이나 스칼라의 기본 패키지에는 이런 기능을 하는 함수가 없는데,
// hoogle 에서 검색해 보면
// https://hoogle.haskell.org/?hoogle=m%20(%20a%20-%3E%20b)%20-%3E%20a%20-%3E%20m%20b
// ?? 혹은 flap 이라는 이름으로 정의된 함수가 있음
func Flap[A, R any](tfa fp.Future[fp.Func1[A, R]], ctx ...fp.Executor) fp.Func1[A, fp.Future[R]] {
	return func(a A) fp.Future[R] {
		return Ap(tfa, Successful(a), ctx...)
	}
}

// 하스켈 : m( a -> b -> r ) -> a -> b -> m r
func Flap2[A, B, R any](tfab fp.Future[fp.Func1[A, fp.Func1[B, R]]], ctx ...fp.Executor) fp.Func1[A, fp.Func1[B, fp.Future[R]]] {
	return func(a A) fp.Func1[B, fp.Future[R]] {
		return Flap(Ap(tfab, Successful(a)), ctx...)
	}
}

// (a -> b -> r) -> m a -> b -> m r
// Map 호출 후에 Flap 을 호출 한 것
//
// https://hoogle.haskell.org/?hoogle=%28+a+-%3E+b+-%3E++r+%29+-%3E+m+a+-%3E++b+-%3E+m+r+&scope=set%3Astackage
// liftOp 라는 이름으로 정의된 것이 있음
func FlapMap[A, B, R any](tfab func(A, B) R, ta fp.Future[A], ctx ...fp.Executor) fp.Func1[B, fp.Future[R]] {
	// 	return Flap(Map(a, as.Func2(tfab).Curried()))
	return func(b B) fp.Future[R] {
		return Map(ta, func(a A) R {
			return tfab(a, b)
		}, ctx...)
	}
}

// ( a -> b -> m r) -> m a -> b -> m r
//
//	Flatten . FlapMap
//
// https://hoogle.haskell.org/?hoogle=(%20a%20-%3E%20b%20-%3E%20m%20r%20)%20-%3E%20m%20a%20-%3E%20%20b%20-%3E%20m%20r%20
// om , ==<<  이름으로 정의된 것이 있음
func FlatFlapMap[A, B, R any](fab func(A, B) fp.Future[R], ta fp.Future[A], ctx ...fp.Executor) fp.Func1[B, fp.Future[R]] {
	return fp.Compose(FlapMap(fab, ta, ctx...), Flatten[R])
}

// FlatMap 과는 아규먼트 순서가 다른 함수로
// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
func Method1[A, B, R any](ta fp.Future[A], fab func(a A, b B) R, ctx ...fp.Executor) fp.Func1[B, fp.Future[R]] {
	return FlapMap(fab, ta, ctx...)
}

func FlatMethod1[A, B, R any](ta fp.Future[A], fab func(a A, b B) fp.Future[R], ctx ...fp.Executor) fp.Func1[B, fp.Future[R]] {
	return FlatFlapMap(fab, ta, ctx...)
}

func Method2[A, B, C, R any](ta fp.Future[A], fabc func(a A, b B, c C) R, ctx ...fp.Executor) fp.Func2[B, C, fp.Future[R]] {

	return func(b B, c C) fp.Future[R] {
		return Map(ta, func(a A) R {
			return fabc(a, b, c)
		}, ctx...)
	}
}

func FlatMethod2[A, B, C, R any](ta fp.Future[A], fabc func(a A, b B, c C) fp.Future[R]) fp.Func2[B, C, fp.Future[R]] {

	return func(b B, c C) fp.Future[R] {
		return FlatMap(ta, func(a A) fp.Future[R] {
			return fabc(a, b, c)
		})
	}
}

func Zip[A, B any](c1 fp.Future[A], c2 fp.Future[B]) fp.Future[fp.Tuple2[A, B]] {
	return Map2(c1, c2, product.Tuple2[A, B])
}

func Zip3[A, B, C any](c1 fp.Future[A], c2 fp.Future[B], c3 fp.Future[C]) fp.Future[fp.Tuple3[A, B, C]] {
	return LiftA3(as.Tuple3[A, B, C])(c1, c2, c3)
}

func Sequence[T any](futureList fp.Seq[fp.Future[T]], ctx ...fp.Executor) fp.Future[fp.Seq[T]] {
	return Map(SequenceIterator(fp.IteratorOfSeq(futureList)), fp.Compose(fp.Iterator[T].ToSeq, as.Seq[T]))

}

func SequenceIterator[T any](futureList fp.Iterator[fp.Future[T]], ctx ...fp.Executor) fp.Future[fp.Iterator[T]] {
	return iterator.Fold(futureList, Successful(iterator.Empty[T]()), LiftA2(fp.Iterator[T].Appended, ctx...))
}

func Traverse[T, U any](itr fp.Iterator[T], fn func(T) fp.Future[U], ctx ...fp.Executor) fp.Future[fp.Iterator[U]] {
	return iterator.Fold(itr, Successful(iterator.Empty[U]()), func(tryItr fp.Future[fp.Iterator[U]], v T) fp.Future[fp.Iterator[U]] {
		return FlatMap(tryItr, func(acc fp.Iterator[U]) fp.Future[fp.Iterator[U]] {
			return Map(fn(v), func(v U) fp.Iterator[U] {
				return acc.Concat(iterator.Of(v))
			}, ctx...)
		}, ctx...)
	})
}

func TraverseSeq[T, U any](seq fp.Seq[T], fn func(T) fp.Future[U], ctx ...fp.Executor) fp.Future[fp.Seq[U]] {
	return Map(Traverse(fp.IteratorOfSeq(seq), fn), fp.Compose(fp.Iterator[U].ToSeq, as.Seq[U]), ctx...)
}

func TraverseFunc[A, R any](far func(A) fp.Future[R], ctx ...fp.Executor) fp.Func1[fp.Iterator[A], fp.Future[fp.Iterator[R]]] {
	return func(iterA fp.Iterator[A]) fp.Future[fp.Iterator[R]] {
		return Traverse(iterA, far, ctx...)
	}
}

func TraverseSeqFunc[A, R any](far func(A) fp.Future[R], ctx ...fp.Executor) fp.Func1[fp.Seq[A], fp.Future[fp.Seq[R]]] {
	return func(seqA fp.Seq[A]) fp.Future[fp.Seq[R]] {
		return TraverseSeq(seqA, far, ctx...)
	}
}

type MonadChain1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A, R]]
}

func (r MonadChain1[H, HT, A, R]) Map(a func(HT) A, ctx ...fp.Executor) fp.Future[R] {
	return r.FlatMap(func(h HT) fp.Future[A] {
		return Successful(a(h))
	}, ctx...)
}

func (r MonadChain1[H, HT, A, R]) HListMap(a func(H) A, ctx ...fp.Executor) fp.Future[R] {
	return r.HListFlatMap(func(h H) fp.Future[A] {
		return Successful(a(h))
	}, ctx...)
}

func (r MonadChain1[H, HT, A, R]) HListFlatMap(a func(H) fp.Future[A], ctx ...fp.Executor) fp.Future[R] {
	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}

func (r MonadChain1[H, HT, A, R]) FlatMap(a func(HT) fp.Future[A], ctx ...fp.Executor) fp.Future[R] {
	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a(v.Head())
	}, ctx...)

	return r.ApFuture(av)
}

func (r MonadChain1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Future[R] {
	return r.ApFuture(FromOption(a))
}

func (r MonadChain1[H, HT, A, R]) ApTry(a fp.Try[A]) fp.Future[R] {
	return r.ApFuture(FromTry(a))
}

func (r MonadChain1[H, HT, A, R]) ApFuture(a fp.Future[A]) fp.Future[R] {
	return Ap(r.fn, a)
}

func (r MonadChain1[H, HT, A, R]) Ap(a A) fp.Future[R] {
	return r.ApFuture(Successful(a))
}

func (r MonadChain1[H, HT, A, R]) ApFutureFunc(a func() fp.Future[A], ctx ...fp.Executor) fp.Future[R] {

	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain1[H, HT, A, R]) ApTryFunc(a func() fp.Try[A], ctx ...fp.Executor) fp.Future[R] {

	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A], ctx ...fp.Executor) fp.Future[R] {

	av := FlatMap(r.h, func(v H) fp.Future[A] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain1[H, HT, A, R]) ApFunc(a func() A, ctx ...fp.Executor) fp.Future[R] {

	av := Map(r.h, func(v H) A {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}

func Chain1[A, R any](fn fp.Func1[A, R]) MonadChain1[hlist.Nil, hlist.Nil, A, R] {
	return MonadChain1[hlist.Nil, hlist.Nil, A, R]{Successful(hlist.Empty()), Successful(fn)}
}

type ApplicativeFunctor1[A, R any] struct {
	fn fp.Future[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[A, R]) ApOption(a fp.Option[A]) fp.Future[R] {
	return r.ApFuture(FromOption(a))
}

func (r ApplicativeFunctor1[A, R]) ApTry(a fp.Try[A]) fp.Future[R] {
	return r.ApFuture(FromTry(a))
}

func (r ApplicativeFunctor1[A, R]) ApFuture(a fp.Future[A]) fp.Future[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[A, R]) Ap(a A) fp.Future[R] {
	return r.ApFuture(Successful(a))
}

func (r ApplicativeFunctor1[A, R]) ApFutureFunc(a func() fp.Future[A], ctx ...fp.Executor) fp.Future[R] {

	return ApFunc(r.fn, a, ctx...)
}
func (r ApplicativeFunctor1[A, R]) ApTryFunc(a func() fp.Try[A], ctx ...fp.Executor) fp.Future[R] {

	return r.ApFutureFunc(func() fp.Future[A] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor1[A, R]) ApOptionFunc(a func() fp.Option[A], ctx ...fp.Executor) fp.Future[R] {

	return r.ApFutureFunc(func() fp.Future[A] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor1[A, R]) ApFunc(a func() A, ctx ...fp.Executor) fp.Future[R] {

	return r.ApFutureFunc(func() fp.Future[A] {
		return Successful(a())
	}, ctx...)
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[A, R] {
	return ApplicativeFunctor1[A, R]{Successful(fn)}
}

func Await[T any](future fp.Future[T], timeout time.Duration) fp.Try[T] {
	value := future.Value()
	if value.IsDefined() {
		return value.Get()
	}

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

func Func0[R any](f func() (R, error), ctx ...fp.Executor) fp.Func1[fp.Unit, fp.Future[R]] {
	return func(fp.Unit) fp.Future[R] {
		return Apply2(f, ctx...)
	}
}
