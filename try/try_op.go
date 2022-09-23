//go:generate go run github.com/csgura/fp/internal/generator/try_gen
package try

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
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

func Apply[T any](v T, err error) fp.Try[T] {
	if err != nil {
		return Failure[T](err)
	}
	return Success(v)
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.Try[B]], f2 fp.Func1[B, fp.Try[C]]) fp.Func1[A, fp.Try[C]] {
	return func(a A) fp.Try[C] {
		return FlatMap(f1(a), f2)
	}
}

func Compose2[A, B, C any](f1 fp.Func1[A, fp.Try[B]], f2 fp.Func1[B, fp.Try[C]]) fp.Func1[A, fp.Try[C]] {
	return func(a A) fp.Try[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposeOption[A, B, C any](f1 fp.Func1[A, fp.Option[B]], f2 fp.Func1[B, fp.Try[C]]) fp.Func1[A, fp.Try[C]] {
	return func(a A) fp.Try[C] {
		return FlatMap(FromOption(f1(a)), f2)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.Try[B]], f2 fp.Func1[B, C]) fp.Func1[A, fp.Try[C]] {
	return func(a A) fp.Try[C] {
		return Map(f1(a), f2)
	}
}

var Unit fp.Try[fp.Unit] = Success(fp.Unit{})

func Ap[T, U any](t fp.Try[fp.Func1[T, U]], a fp.Try[T]) fp.Try[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Try[U] {
		return Map(a, f)
	})
}

func Map[T, U any](opt fp.Try[T], f func(v T) U) fp.Try[U] {
	return FlatMap(opt, func(v T) fp.Try[U] {
		return Success(f(v))
	})
}

func Map2[A, B, U any](a fp.Try[A], b fp.Try[B], f func(A, B) U) fp.Try[U] {
	return FlatMap(a, func(v1 A) fp.Try[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

// func Map[T, U any](opt fp.Try[T], f func(v T) U) fp.Try[U] {
// 	return Ap(Success(as.Func1(f)), opt)
// }

func Lift[T, U any](f func(v T) U) fp.Func1[fp.Try[T], fp.Try[U]] {
	return func(opt fp.Try[T]) fp.Try[U] {
		return Map(opt, f)
	}
}

func LiftA2[A1, A2, R any](f fp.Func2[A1, A2, R]) fp.Func2[fp.Try[A1], fp.Try[A2], fp.Try[R]] {
	return func(a1 fp.Try[A1], a2 fp.Try[A2]) fp.Try[R] {
		return Ap(Ap(Success(f.Curried()), a1), a2)
	}
}

func FlatMap[T, U any](opt fp.Try[T], fn func(v T) fp.Try[U]) fp.Try[U] {
	if opt.IsSuccess() {
		return fn(opt.Get())
	}
	return Failure[U](opt.Failed().Get())
}

func Flatten[T any](opt fp.Try[fp.Try[T]]) fp.Try[T] {
	return FlatMap(opt, func(v fp.Try[T]) fp.Try[T] {
		return v
	})
}

func Zip[A, B any](c1 fp.Try[A], c2 fp.Try[B]) fp.Try[fp.Tuple2[A, B]] {
	return FlatMap(c1, func(v1 A) fp.Try[fp.Tuple2[A, B]] {
		return Map(c2, func(v2 B) fp.Tuple2[A, B] {
			return product.Tuple2(v1, v2)
		})
	})
}

func Zip3[A, B, C any](c1 fp.Try[A], c2 fp.Try[B], c3 fp.Try[C]) fp.Try[fp.Tuple3[A, B, C]] {
	return Applicative3(as.Tuple3[A, B, C]).
		ApTry(c1).
		ApTry(c2).
		ApTry(c3)
}

func SequenceIterator[T any](tryItr fp.Iterator[fp.Try[T]]) fp.Try[fp.Iterator[T]] {

	return iterator.Fold(tryItr, Success(iterator.Empty[T]()), func(list fp.Try[fp.Iterator[T]], v fp.Try[T]) fp.Try[fp.Iterator[T]] {
		return Map2(list, v, func(l fp.Iterator[T], e T) fp.Iterator[T] {
			return l.Concat(iterator.Of(e))
		})
	})
}

func Traverse[T, U any](itr fp.Iterator[T], fn func(T) fp.Try[U]) fp.Try[fp.Iterator[U]] {
	return iterator.Fold(itr, Success(iterator.Empty[U]()), func(tryItr fp.Try[fp.Iterator[U]], v T) fp.Try[fp.Iterator[U]] {
		return FlatMap(tryItr, func(acc fp.Iterator[U]) fp.Try[fp.Iterator[U]] {
			return Map(fn(v), func(v U) fp.Iterator[U] {
				return acc.Concat(iterator.Of(v))
			})
		})
	})
}

func TraverseSeq[T, U any](seq fp.Seq[T], fn func(T) fp.Try[U]) fp.Try[fp.Seq[U]] {
	return Map(Traverse(seq.Iterator(), fn), fp.Iterator[U].ToSeq)
}

func Sequence[T any](trySeq fp.Seq[fp.Try[T]]) fp.Try[fp.Seq[T]] {
	return Map(SequenceIterator(trySeq.Iterator()), fp.Iterator[T].ToSeq)
}

func Fold[A, B any](s fp.Try[A], zero B, f func(B, A) B) B {
	if s.IsFailure() {
		return zero
	}

	return f(zero, s.Get())
}

func FoldRight[A, B any](s fp.Try[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsFailure() {
		return lazy.Done(zero)
	}

	return f(s.Get(), lazy.Done(zero))
}

type ApplicativeFunctor1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[H, HT, A, R]) Map(a func(HT) A) fp.Try[R] {
	return r.FlatMap(func(h HT) fp.Try[A] {
		return Success(a(h))
	})
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListMap(a func(H) A) fp.Try[R] {
	return r.HListFlatMap(func(h H) fp.Try[A] {
		return Success(a(h))
	})
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListFlatMap(a func(H) fp.Try[A]) fp.Try[R] {
	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a(v)
	})

	return r.ApTry(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) FlatMap(a func(HT) fp.Try[A]) fp.Try[R] {
	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a(v.Head())
	})

	return r.ApTry(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Try[R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApTry(a fp.Try[A]) fp.Try[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[H, HT, A, R]) Ap(a A) fp.Try[R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApTryFunc(a func() fp.Try[A]) fp.Try[R] {

	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a()
	})
	return r.ApTry(av)
}
func (r ApplicativeFunctor1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Try[R] {

	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r ApplicativeFunctor1[H, HT, A, R]) ApFunc(a func() A) fp.Try[R] {

	av := Map(r.h, func(v H) A {
		return a()
	})
	return r.ApTry(av)
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R] {
	return ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R]{Success(hlist.Empty()), Success(fn)}
}

func Func0[R any](f func() (R, error)) fp.Func1[fp.Unit, fp.Try[R]] {
	return func(fp.Unit) fp.Try[R] {
		ret, err := f()
		return Apply(ret, err)
	}
}
