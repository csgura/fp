//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package future

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/seq"
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

func Constructors[T fp.Future[V], V any]() (failed func(error) fp.Future[V], successful func(V) fp.Future[V]) {
	return Failed[V], Successful[V]
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

func ApplyT[T any](f func() fp.Try[T], ctx ...fp.Executor) fp.Future[T] {
	p := promise.New[T]()

	getExecutor(ctx...).ExecuteUnsafe(fp.RunnableFunc(func() {
		defer func() {
			if err := recover(); err != nil {
				p.Failure(fp.PanicError(err))
			}
		}()

		rt := f()
		p.Complete(rt)
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
		return Map(a, f, ctx...)
	}, ctx...)
}

func ApFunc[T, U any](t fp.Future[fp.Func1[T, U]], a func() fp.Future[T], ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Future[U] {
		return Map(a(), f, ctx...)
	}, ctx...)
}

func Map[T, U any](opt fp.Future[T], f func(v T) U, ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(opt, func(v T) fp.Future[U] {
		return Successful(f(v))
	}, ctx...)
}

// haskell 의 <$
// map . const 와 같은 함수
func Replace[A, B any](ta fp.Future[A], b B) fp.Future[B] {
	return Map(ta, fp.Const[A](b))
}

// Map(ta , seq.Lift(f)) 와 동일
func MapSeqLift[A, B any](ta fp.Future[fp.Seq[A]], f func(v A) B, ctx ...fp.Executor) fp.Future[fp.Seq[B]] {
	return Map(ta, func(a fp.Seq[A]) fp.Seq[B] {
		return iterator.Map(iterator.FromSeq(a), f).ToSeq()
	}, ctx...)
}

// Map(ta , seq.Lift(f)) 와 동일
func MapSliceLift[A, B any](ta fp.Future[[]A], f func(v A) B, ctx ...fp.Executor) fp.Future[[]B] {
	return Map(ta, func(a []A) []B {
		return iterator.Map(iterator.FromSeq(a), f).ToSeq()
	}, ctx...)
}

func Map2[A, B, U any](a fp.Future[A], b fp.Future[B], f func(A, B) U, ctx ...fp.Executor) fp.Future[U] {
	return FlatMap(a, func(v1 A) fp.Future[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		}, ctx...)
	}, ctx...)
}

// fp.With 의 future 버젼
// fp.With 가 Flip 과 사실상 같은 것처럼
// FlapMap 의 Flip 버젼과 동일
// var b fp.Future[B]
// a := future.Successful(A{})
// a.FlatMap( future.With(A.WithB, b))
// 형태로 코딩 가능
func With[A, B any](withf func(A, B) A, v fp.Future[B], ctx ...fp.Executor) func(A) fp.Future[A] {
	return Flap(Map(v, fp.Flip2(withf), ctx...), ctx...)
}

func Lift[T, U any](f func(v T) U, ctx ...fp.Executor) func(fp.Future[T]) fp.Future[U] {
	return func(opt fp.Future[T]) fp.Future[U] {
		return Map(opt, f, ctx...)
	}
}

func LiftA2[A1, A2, R any](f func(A1, A2) R, ctx ...fp.Executor) func(fp.Future[A1], fp.Future[A2]) fp.Future[R] {
	return func(a1 fp.Future[A1], a2 fp.Future[A2]) fp.Future[R] {
		return Map2(a1, a2, f, ctx...)
	}
}

func LiftM[A, R any](fa func(v A) fp.Future[R], ctx ...fp.Executor) func(fp.Future[A]) fp.Future[R] {
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
func LiftM2[A, B, R any](fab func(A, B) fp.Future[R], ctx ...fp.Executor) func(fp.Future[A], fp.Future[B]) fp.Future[R] {
	return func(a fp.Future[A], b fp.Future[B]) fp.Future[R] {
		return Flatten(Map2(a, b, fab, ctx...))
	}
}

func Compose[A, B, C any](f1 func(A) fp.Future[B], f2 func(B) fp.Future[C], ctx ...fp.Executor) func(A) fp.Future[C] {
	return func(a A) fp.Future[C] {
		return FlatMap(f1(a), f2, ctx...)
	}
}

func Compose2[A, B, C any](f1 func(A) fp.Future[B], f2 func(B) fp.Future[C], ctx ...fp.Executor) func(A) fp.Future[C] {
	return func(a A) fp.Future[C] {
		return FlatMap(f1(a), f2, ctx...)
	}
}

func ComposeOption[A, B, C any](f1 func(A) fp.Option[B], f2 func(B) fp.Future[C], ctx ...fp.Executor) func(A) fp.Future[C] {
	return func(a A) fp.Future[C] {
		return FlatMap(FromOption(f1(a)), f2, ctx...)
	}
}

func ComposeTry[A, B, C any](f1 func(A) fp.Try[B], f2 func(B) fp.Future[C], ctx ...fp.Executor) func(A) fp.Future[C] {
	return func(a A) fp.Future[C] {
		return FlatMap(FromTry(f1(a)), f2, ctx...)
	}
}

func ComposePure[A, B any](fab func(A) B, ctx ...fp.Executor) func(A) fp.Future[B] {
	return fp.Compose(fab, Successful)
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

func FlatMapTraverseSeq[A, B any](ta fp.Future[fp.Seq[A]], f func(v A) fp.Future[B], ctx ...fp.Executor) fp.Future[fp.Seq[B]] {
	return FlatMap(ta, TraverseSeqFunc(f, ctx...), ctx...)
}

func FlatMapTraverseSlice[A, B any](ta fp.Future[[]A], f func(v A) fp.Future[B], ctx ...fp.Executor) fp.Future[[]B] {
	return FlatMap(ta, TraverseSliceFunc(f, ctx...), ctx...)
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
func Flap[A, R any](tfa fp.Future[fp.Func1[A, R]], ctx ...fp.Executor) func(A) fp.Future[R] {
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
func FlapMap[A, B, R any](tfab func(A, B) R, ta fp.Future[A], ctx ...fp.Executor) func(B) fp.Future[R] {
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
func FlatFlapMap[A, B, R any](fab func(A, B) fp.Future[R], ta fp.Future[A], ctx ...fp.Executor) func(B) fp.Future[R] {
	return fp.Compose(FlapMap(fab, ta, ctx...), Flatten)
}

// FlatMap 과는 아규먼트 순서가 다른 함수로
// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
func Method1[A, B, R any](ta fp.Future[A], fab func(a A, b B) R, ctx ...fp.Executor) func(B) fp.Future[R] {
	return FlapMap(fab, ta, ctx...)
}

func FlatMethod1[A, B, R any](ta fp.Future[A], fab func(a A, b B) fp.Future[R], ctx ...fp.Executor) func(B) fp.Future[R] {
	return FlatFlapMap(fab, ta, ctx...)
}

func Method2[A, B, C, R any](ta fp.Future[A], fabc func(a A, b B, c C) R, ctx ...fp.Executor) func(B, C) fp.Future[R] {

	return func(b B, c C) fp.Future[R] {
		return Map(ta, func(a A) R {
			return fabc(a, b, c)
		}, ctx...)
	}
}

func FlatMethod2[A, B, C, R any](ta fp.Future[A], fabc func(a A, b B, c C) fp.Future[R]) func(B, C) fp.Future[R] {

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

func Sequence[T any](futureList []fp.Future[T], ctx ...fp.Executor) fp.Future[[]T] {
	ret := iterator.Fold(iterator.FromSlice(futureList), Successful(seq.Empty[T]()), LiftA2(fp.Seq[T].Add, ctx...))
	return Map(ret, fp.Seq[T].Widen, ctx...)
}

func SequenceIterator[T any](futureList fp.Iterator[fp.Future[T]], ctx ...fp.Executor) fp.Future[fp.Iterator[T]] {
	ret := iterator.Fold(futureList, Successful(seq.Empty[T]()), LiftA2(fp.Seq[T].Add, ctx...))
	return Map(ret, iterator.FromSeq, ctx...)

}

func traverse[T, U any](itr fp.Iterator[T], fn func(T) fp.Future[U], ctx ...fp.Executor) fp.Future[fp.Seq[U]] {
	ret := iterator.FoldFuture(itr, seq.Empty[U](), func(acc fp.Seq[U], v T) fp.Future[fp.Seq[U]] {
		return Map(fn(v), acc.Add, ctx...)
	})
	return ret
}

func Traverse[T, U any](itr fp.Iterator[T], fn func(T) fp.Future[U], ctx ...fp.Executor) fp.Future[fp.Iterator[U]] {
	ret := traverse(itr, fn, ctx...)
	return Map(ret, iterator.FromSeq, ctx...)
}

func TraverseSeq[T, U any](seq fp.Seq[T], fn func(T) fp.Future[U], ctx ...fp.Executor) fp.Future[fp.Seq[U]] {
	return traverse(fp.IteratorOfSeq(seq), fn, ctx...)
}

func TraverseSlice[T, U any](seq []T, fn func(T) fp.Future[U], ctx ...fp.Executor) fp.Future[[]U] {
	return Map(traverse(fp.IteratorOfSeq(seq), fn), fp.Seq[U].Widen, ctx...)
}

func TraverseFunc[A, R any](far func(A) fp.Future[R], ctx ...fp.Executor) func(fp.Iterator[A]) fp.Future[fp.Iterator[R]] {
	return func(iterA fp.Iterator[A]) fp.Future[fp.Iterator[R]] {
		return Traverse(iterA, far, ctx...)
	}
}

func TraverseSeqFunc[A, R any](far func(A) fp.Future[R], ctx ...fp.Executor) func(fp.Seq[A]) fp.Future[fp.Seq[R]] {
	return func(seqA fp.Seq[A]) fp.Future[fp.Seq[R]] {
		return TraverseSeq(seqA, far, ctx...)
	}
}

func TraverseSliceFunc[A, R any](far func(A) fp.Future[R], ctx ...fp.Executor) func([]A) fp.Future[[]R] {
	return func(seqA []A) fp.Future[[]R] {
		return TraverseSlice(seqA, far, ctx...)
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

	if future.IsCompleted() {
		return future.Value()
	}
	// value := future.Value()
	// if value.IsDefined() {
	// 	return value.Get()
	// }

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

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: genfp.Imports(
		"github.com/csgura/fp",
		"github.com/csgura/fp/curried",
		"github.com/csgura/fp/hlist",
	),
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r MonadChain{{.N}}[H, HT, {{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}MonadChain{{dec .N}}[hlist.Cons[A1, H], {{TypeArgs 1 .N}}, R]{{end}}


type MonadChain{{.N}}[H hlist.Header[HT], HT, {{TypeArgs 1 .N}}, R any] struct {
	h  fp.Future[H]
	fn fp.Future[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
{{template "Receiver" .}} Map(a func(HT) A1, ctx ...fp.Executor) {{template "Next" .}} {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
{{template "Receiver" .}} HListMap(a func(H) A1, ctx ...fp.Executor) {{template "Next" .}} {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
{{template "Receiver" .}} HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}

{{template "Receiver" .}} ApFuture(a fp.Future[A1]) {{template "Next" .}} {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return {{template "Next" .}}{nh, Ap(r.fn, a)}
}

{{template "Receiver" .}} ApTry(a fp.Try[A1]) {{template "Next" .}} {

	return r.ApFuture(FromTry(a))
}
{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	return r.ApFuture(FromOption(a))
}
{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApFuture(Successful(a))

}

{{template "Receiver" .}} ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}

{{template "Receiver" .}} ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
{{template "Receiver" .}} ApFunc(a func() A1, ctx ...fp.Executor) {{template "Next" .}} {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}

func Chain{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R] {
	return MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R]{Successful(hlist.Empty()), Successful(curried.Func{{.N}}(fn))}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: genfp.Imports(
		"github.com/csgura/fp",
		"github.com/csgura/fp/curried",
		"github.com/csgura/fp/hlist",
	),
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}ApplicativeFunctor{{dec .N}}[{{TypeArgs 2 .N}}, R]{{end}}

type ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R any] struct {
	fn fp.Future[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} ApFuture(a fp.Future[A1]) {{template "Next" .}} {
	return {{template "Next" .}}{Ap(r.fn, a)}
}

{{template "Receiver" .}} ApFutureAll({{DeclTypeClassArgs 1 .N "fp.Future"}}) fp.Future[R] {
	return r.
	{{- range (dec .N) -}}
		ApFuture(ins{{inc .}}).
	{{- end -}}
		ApFuture(ins{{.N}})
}

{{template "Receiver" .}} ApTry(a fp.Try[A1]) {{template "Next" .}} {
	return r.ApFuture(FromTry(a))
}

{{template "Receiver" .}} ApTryAll({{DeclTypeClassArgs 1 .N "fp.Try"}}) fp.Future[R] {
	return r.
	{{- range (dec .N) -}}
		ApTry(ins{{inc .}}).
	{{- end -}}
		ApTry(ins{{.N}})
}

{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {
	return r.ApFuture(FromOption(a))
}

{{template "Receiver" .}} ApOptionAll({{DeclTypeClassArgs 1 .N "fp.Option"}}) fp.Future[R] {
	return r.
	{{- range (dec .N) -}}
		ApOption(ins{{inc .}}).
	{{- end -}}
		ApOption(ins{{.N}})
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {
	return r.ApFuture(Successful(a))
}

{{template "Receiver" .}} ApAll({{DeclArgs 1 .N}}) fp.Future[R] {
	return r.
	{{- range (dec .N) -}}
		Ap(a{{inc .}}).
	{{- end -}}
		Ap(a{{.N}})
}

{{template "Receiver" .}} ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) {{template "Next" .}} {

	return {{template "Next" .}}{ApFunc(r.fn, a)}

}

{{template "Receiver" .}} ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) {{template "Next" .}} {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)

}

{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) {{template "Next" .}} {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}

{{template "Receiver" .}} ApFunc(a func() A1, ctx ...fp.Executor) {{template "Next" .}} {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}

func Applicative{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R] {
	return ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]{Successful(curried.Func{{.N}}(fn))}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func LiftA{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{DeclArgs 1 .N}}) R, exec ...fp.Executor) func({{TypeClassArgs 1 .N "fp.Future"}}) fp.Future[R] {
	return func({{DeclTypeClassArgs 1 .N "fp.Future"}}) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA{{dec .N}}(func({{DeclArgs 2 .N}}) R {
				return f({{CallArgs 1 .N}})
			}, exec...)({{CallArgs 2 .N "ins"}})
		}, exec...)
	}
}

func LiftM{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{DeclArgs 1 .N}}) fp.Future[R], exec ...fp.Executor) func({{TypeClassArgs 1 .N "fp.Future"}}) fp.Future[R] {
	return func({{DeclTypeClassArgs 1 .N "fp.Future"}}) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM{{dec .N}}(func({{DeclArgs 2 .N}}) fp.Future[R] {
				return f({{CallArgs 1 .N}})
			}, exec...)({{CallArgs 2 .N "ins"}})
		}, exec...)
	}
}

func Flap{{.N}}[{{TypeArgs 1 .N}}, R any](tf fp.Future[{{CurriedFunc 1 .N "R"}}], exec ...fp.Executor) {{CurriedFunc 1 .N "fp.Future[R]"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "fp.Future[R]"}} {
		return Flap{{dec .N}}(Ap(tf, Successful(a1)), exec...)
	}
}

func Method{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Future[A1], fa1 func({{DeclArgs 1 .N}}) R, exec ...fp.Executor) func({{TypeArgs 2 .N}}) fp.Future[R] {
	return func({{DeclArgs 2 .N}}) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1({{CallArgs 1 .N}})
		}, exec...)
	}
}

func FlatMethod{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Future[A1], fa1 func({{DeclArgs 1 .N}}) fp.Future[R], exec ...fp.Executor) func({{TypeArgs 2 .N}}) fp.Future[R] {
	return func({{DeclArgs 2 .N}}) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1({{CallArgs 1 .N}})
		}, exec...)
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  1,
	Until: genfp.MaxFunc,
	Template: `
func Func{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) (R, error), exec ...fp.Executor) fp.Func{{.N}}[{{TypeArgs 1 .N}}, fp.Future[R]] {
	return func({{DeclArgs 1 .N}}) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f({{CallArgs 1 .N}})
		})
	}
}

func Unit{{.N}}[{{TypeArgs 1 .N}} any](f func({{TypeArgs 1 .N}}) error, exec ...fp.Executor) fp.Func{{.N}}[{{TypeArgs 1 .N}}, fp.Future[fp.Unit]] {
	return func({{DeclArgs 1 .N}}) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f({{CallArgs 1 .N}})
			return fp.Unit{}, err
		})
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxCompose,
	Template: `
func Compose{{.N}}[{{TypeArgs 1 .N}}, R any]({{(Monad "fp.Future").FuncChain 1 .N}}, exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose{{dec .N}}({{CallArgs 2 .N "f"}}, exec...), exec...)
}
	`,
}
