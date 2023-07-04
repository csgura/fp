//go:generate go run github.com/csgura/fp/internal/generator/try_gen
package try

import (
	"fmt"
	"runtime/debug"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/product"
)

func Success[T any](t T) fp.Try[T] {
	return fp.Success(t)
}

func Failure[T any](err error) fp.Try[T] {
	return fp.Failure[T](err)
}

func FromOption[T any](v fp.Option[T]) fp.Try[T] {
	if v.IsDefined() {
		return Success(v.Get())
	} else {
		return Failure[T](fp.ErrOptionEmpty)
	}
}

type Panic interface {
	error
	Panic() any
	Stack() []byte
}

type panicError struct {
	cause any
	stack []byte
}

func (r *panicError) Error() string {
	return fmt.Sprintf("%v %v", r.cause, string(r.stack))
}

func (r *panicError) Stack() []byte {
	return r.stack
}

func (r *panicError) Panic() any {
	return r.cause
}

func Of[T any](f func() T) (ret fp.Try[T]) {
	defer func() {
		if p := recover(); p != nil {
			ret = Failure[T](&panicError{p, debug.Stack()})
		}
	}()

	ret = Success(f())
	return
}

func Apply[T any](v T, err error) fp.Try[T] {
	if err != nil {
		return Failure[T](err)
	}
	return Success(v)
}

func Compose[A, B, C any](f1 func(A) fp.Try[B], f2 func(B) fp.Try[C]) func(A) fp.Try[C] {
	return func(a A) fp.Try[C] {
		return FlatMap(f1(a), f2)
	}
}

func Compose2[A, B, C any](f1 func(A) fp.Try[B], f2 func(B) fp.Try[C]) func(A) fp.Try[C] {
	return func(a A) fp.Try[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposeOption[A, B, C any](f1 func(A) fp.Option[B], f2 func(B) fp.Try[C]) func(A) fp.Try[C] {
	return func(a A) fp.Try[C] {
		return FlatMap(FromOption(f1(a)), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Try[B] {
	return fp.Compose(fab, Success[B])
}

var Unit fp.Try[fp.Unit] = Success(fp.Unit{})

func Ap[A, B any](tfab fp.Try[fp.Func1[A, B]], ta fp.Try[A]) fp.Try[B] {
	return FlatMap(tfab, func(fab fp.Func1[A, B]) fp.Try[B] {
		return Map(ta, fab)
	})
}

func ApFunc[A, B any](tfab fp.Try[fp.Func1[A, B]], ta func() fp.Try[A]) fp.Try[B] {
	return FlatMap(tfab, func(fab fp.Func1[A, B]) fp.Try[B] {
		return Map(ta(), fab)
	})
}

func Map[A, B any](ta fp.Try[A], f func(v A) B) fp.Try[B] {
	return FlatMap(ta, func(a A) fp.Try[B] {
		return Success(f(a))
	})
}

func Map2[A, B, R any](ta fp.Try[A], tb fp.Try[B], fab func(A, B) R) fp.Try[R] {
	return FlatMap(ta, func(a A) fp.Try[R] {
		return Map(tb, func(b B) R {
			return fab(a, b)
		})
	})
}

// func Map[T, U any](opt fp.Try[T], f func(v T) U) fp.Try[U] {
// 	return Ap(Success(as.Func1(f)), opt)
// }

func Lift[A, R any](fa func(v A) R) func(fp.Try[A]) fp.Try[R] {
	return func(ta fp.Try[A]) fp.Try[R] {
		return Map(ta, fa)
	}
}

func LiftA2[A, B, R any](fab func(A, B) R) func(fp.Try[A], fp.Try[B]) fp.Try[R] {
	return func(a fp.Try[A], b fp.Try[B]) fp.Try[R] {
		return Map2(a, b, fab)
	}
}

func LiftM[A, R any](fa func(v A) fp.Try[R]) func(fp.Try[A]) fp.Try[R] {
	return func(ta fp.Try[A]) fp.Try[R] {
		return Flatten(Map(ta, fa))
	}
}

// (a -> b -> m r) -> m a -> m b -> m r
// 하스켈에서는  liftM2 와 liftA2 는 같은 함수이고
// 위와 같은 함수는 존재하지 않음.
// hoogle 에서 검색해 보면 , liftJoin2 , bindM2 등의 이름으로 정의된 것이 있음.
// 하지만 ,  fp 패키지에서도   LiftA2 와 LiftM2 를 동일하게 하는 것은 낭비이고
// M 은 Monad 라는 뜻인데, Monad는 Flatten, FlatMap 의 의미가 있으니까
// LiftM2 를 다음과 같이 정의함.
func LiftM2[A, B, R any](fab func(A, B) fp.Try[R]) func(fp.Try[A], fp.Try[B]) fp.Try[R] {
	return func(a fp.Try[A], b fp.Try[B]) fp.Try[R] {
		return Flatten(Map2(a, b, fab))
	}
}

func FlatMap[A, B any](ta fp.Try[A], fn func(v A) fp.Try[B]) fp.Try[B] {
	if ta.IsSuccess() {
		return fn(ta.Get())
	}
	return Failure[B](ta.Failed().Get())
}

func Flatten[A any](tta fp.Try[fp.Try[A]]) fp.Try[A] {
	return FlatMap(tta, func(v fp.Try[A]) fp.Try[A] {
		return v
	})
}

// 하스켈 : m( a -> r ) -> a -> m r
// 스칼라 : M[ A => r ] => A => M[R]
// 하스켈이나 스칼라의 기본 패키지에는 이런 기능을 하는 함수가 없는데,
// hoogle 에서 검색해 보면
// https://hoogle.haskell.org/?hoogle=m%20(%20a%20-%3E%20b)%20-%3E%20a%20-%3E%20m%20b
// ?? 혹은 flap 이라는 이름으로 정의된 함수가 있음
func Flap[A, R any](tfa fp.Try[fp.Func1[A, R]]) func(A) fp.Try[R] {
	return func(a A) fp.Try[R] {
		return Ap(tfa, Success(a))
	}
}

// 하스켈 : m( a -> b -> r ) -> a -> b -> m r
func Flap2[A, B, R any](tfab fp.Try[fp.Func1[A, fp.Func1[B, R]]]) fp.Func1[A, fp.Func1[B, fp.Try[R]]] {
	return func(a A) fp.Func1[B, fp.Try[R]] {
		return Flap(Ap(tfab, Success(a)))
	}
}

// (a -> b -> r) -> m a -> b -> m r
// Map 호출 후에 Flap 을 호출 한 것
//
// https://hoogle.haskell.org/?hoogle=%28+a+-%3E+b+-%3E++r+%29+-%3E+m+a+-%3E++b+-%3E+m+r+&scope=set%3Astackage
// liftOp 라는 이름으로 정의된 것이 있음
func FlapMap[A, B, R any](tfab func(A, B) R, a fp.Try[A]) func(B) fp.Try[R] {
	return Flap(Map(a, as.Curried2(tfab)))
}

// ( a -> b -> m r) -> m a -> b -> m r
//
//	Flatten . FlapMap
//
// https://hoogle.haskell.org/?hoogle=(%20a%20-%3E%20b%20-%3E%20m%20r%20)%20-%3E%20m%20a%20-%3E%20%20b%20-%3E%20m%20r%20
// om , ==<<  이름으로 정의된 것이 있음
func FlatFlapMap[A, B, R any](fab func(A, B) fp.Try[R], ta fp.Try[A]) func(B) fp.Try[R] {
	return fp.Compose(FlapMap(fab, ta), Flatten[R])
}

// FlatMap 과는 아규먼트 순서가 다른 함수로
// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
func Method1[A, B, R any](ta fp.Try[A], fab func(a A, b B) R) func(B) fp.Try[R] {
	return FlapMap(fab, ta)
}

func FlatMethod1[A, B, R any](ta fp.Try[A], fab func(a A, b B) fp.Try[R]) func(B) fp.Try[R] {
	return FlatFlapMap(fab, ta)
}

func Method2[A, B, C, R any](ta fp.Try[A], fabc func(a A, b B, c C) R) func(B, C) fp.Try[R] {

	return curried.Revert2(Flap2(Map(ta, as.Curried3(fabc))))
	// return func(b B, c C) fp.Try[R] {
	// 	return Map(a, func(a A) R {
	// 		return cf(a, b, c)
	// 	})
	// }
}

func FlatMethod2[A, B, C, R any](ta fp.Try[A], fabc func(a A, b B, c C) fp.Try[R]) func(B, C) fp.Try[R] {

	return curried.Revert2(curried.Compose2(Flap2(Map(ta, as.Curried3(fabc))), Flatten[R]))

	// return func(b B, c C) fp.Try[R] {
	// 	return FlatMap(ta, func(a A) fp.Try[R] {
	// 		return cf(a, b, c)
	// 	})
	// }
}

func Zip[A, B any](ta fp.Try[A], tb fp.Try[B]) fp.Try[fp.Tuple2[A, B]] {
	return Map2(ta, tb, product.Tuple2[A, B])
}

func Zip3[A, B, C any](ta fp.Try[A], tb fp.Try[B], tc fp.Try[C]) fp.Try[fp.Tuple3[A, B, C]] {
	return LiftA3(as.Tuple3[A, B, C])(ta, tb, tc)
}

func SequenceIterator[A any](ita fp.Iterator[fp.Try[A]]) fp.Try[fp.Iterator[A]] {
	return iterator.Fold(ita, Success(iterator.Empty[A]()), LiftA2(fp.Iterator[A].Appended))
}

func Traverse[A, R any](ia fp.Iterator[A], fn func(A) fp.Try[R]) fp.Try[fp.Iterator[R]] {
	return iterator.Fold(ia, Success(iterator.Empty[R]()), func(tir fp.Try[fp.Iterator[R]], a A) fp.Try[fp.Iterator[R]] {
		// return ApFunc(Ap(Success(as.Curried2(fp.Iterator[R].Appended)), tir), lazy.Func1(fn)(a))
		return FlatMap(tir, func(acc fp.Iterator[R]) fp.Try[fp.Iterator[R]] {
			return Map(fn(a), acc.Appended)
		})
	})
}

func TraverseOption[A, R any](opta fp.Option[A], fa func(A) fp.Try[R]) fp.Try[fp.Option[R]] {
	return Map(Traverse(fp.IteratorOfOption(opta), fa), fp.Iterator[R].NextOption)
}

func TraverseSeq[A, R any](sa []A, fa func(A) fp.Try[R]) fp.Try[[]R] {
	return Map(Traverse(fp.IteratorOfSeq(sa), fa), fp.Iterator[R].ToSeq)
}

func TraverseFunc[A, R any](far func(A) fp.Try[R]) func(fp.Iterator[A]) fp.Try[fp.Iterator[R]] {
	return func(iterA fp.Iterator[A]) fp.Try[fp.Iterator[R]] {
		return Traverse(iterA, far)
	}
}

func TraverseSeqFunc[A, R any](far func(A) fp.Try[R]) func([]A) fp.Try[[]R] {
	return func(seqA []A) fp.Try[[]R] {
		return TraverseSeq(seqA, far)
	}
}

func Sequence[A any](tsa []fp.Try[A]) fp.Try[[]A] {
	return Map(SequenceIterator(fp.IteratorOfSeq(tsa)), fp.Iterator[A].ToSeq)
}

func Fold[A, B any](ta fp.Try[A], bzero B, fba func(B, A) B) B {
	if ta.IsFailure() {
		return bzero
	}

	return fba(bzero, ta.Get())
}

func FoldRight[A, B any](ta fp.Try[A], bzero B, fab func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if ta.IsFailure() {
		return lazy.Done(bzero)
	}

	return fab(ta.Get(), lazy.Done(bzero))
}

func ToSeq[A any](ta fp.Try[A]) fp.Seq[A] {
	if ta.IsSuccess() {
		return fp.Seq[A]{ta.Get()}
	}
	return nil
}

func Iterator[A any](ta fp.Try[A]) fp.Iterator[A] {
	return fp.IteratorOfSeq(ToSeq(ta))
}

type MonadChain1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A, R]]
}

func (r MonadChain1[H, HT, A, R]) Map(a func(HT) A) fp.Try[R] {
	return r.FlatMap(func(h HT) fp.Try[A] {
		return Success(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListMap(a func(H) A) fp.Try[R] {
	return r.HListFlatMap(func(h H) fp.Try[A] {
		return Success(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListFlatMap(a func(H) fp.Try[A]) fp.Try[R] {
	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a(v)
	})

	return r.ApTry(av)
}

func (r MonadChain1[H, HT, A, R]) FlatMap(a func(HT) fp.Try[A]) fp.Try[R] {
	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a(v.Head())
	})

	return r.ApTry(av)
}

func (r MonadChain1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Try[R] {
	return r.ApTry(FromOption(a))
}

func (r MonadChain1[H, HT, A, R]) ApTry(a fp.Try[A]) fp.Try[R] {
	return Ap(r.fn, a)
}

func (r MonadChain1[H, HT, A, R]) Ap(a A) fp.Try[R] {
	return r.ApTry(Success(a))
}

func (r MonadChain1[H, HT, A, R]) ApTryFunc(a func() fp.Try[A]) fp.Try[R] {

	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Try[R] {

	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain1[H, HT, A, R]) ApFunc(a func() A) fp.Try[R] {

	av := Map(r.h, func(v H) A {
		return a()
	})
	return r.ApTry(av)
}

func Chain1[A, R any](fn fp.Func1[A, R]) MonadChain1[hlist.Nil, hlist.Nil, A, R] {
	return MonadChain1[hlist.Nil, hlist.Nil, A, R]{Success(hlist.Empty()), Success(fn)}
}

type ApplicativeFunctor1[A, R any] struct {
	fn fp.Try[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[A, R]) ApOption(a fp.Option[A]) fp.Try[R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor1[A, R]) ApTry(a fp.Try[A]) fp.Try[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[A, R]) Ap(a A) fp.Try[R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor1[A, R]) ApTryFunc(a func() fp.Try[A]) fp.Try[R] {

	return ApFunc(r.fn, a)

}
func (r ApplicativeFunctor1[A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Try[R] {
	return r.ApTryFunc(func() fp.Try[A] {
		return FromOption(a())
	})
}
func (r ApplicativeFunctor1[A, R]) ApFunc(a func() A) fp.Try[R] {

	return r.ApTryFunc(func() fp.Try[A] {
		return Success(a())
	})

}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[A, R] {
	return ApplicativeFunctor1[A, R]{Success(fn)}
}

func Func0[R any](f func() (R, error)) fp.Func1[fp.Unit, fp.Try[R]] {
	return func(fp.Unit) fp.Try[R] {
		ret, err := f()
		return Apply(ret, err)
	}
}
