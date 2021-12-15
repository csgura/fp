//go:generate go run github.com/csgura/fp/internal/generator/try_gen
package try

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/product"
)

func Success[T any](t T) fp.Try[T] {
	return success[T]{t}
}

func Failure[T any](err error) fp.Try[T] {
	return failure[T]{err}
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

var Unit fp.Try[fp.Unit] = Success(fp.Unit{})

func Ap[T, U any](t fp.Try[fp.Func1[T, U]], a fp.Try[T]) fp.Try[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Try[U] {
		return Map(a.(fp.Try[T]), func(a T) U {
			return f(a)
		})
	})
}

func Map[T, U any](opt fp.Try[T], f func(v T) U) fp.Try[U] {
	return FlatMap(opt, func(v T) fp.Try[U] {
		return Success(f(v))
	})
}

func Lift[T, U any](f func(v T) U) fp.Func1[fp.Try[T], fp.Try[U]] {
	return func(opt fp.Try[T]) fp.Try[U] {
		return Map(opt, f)
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

type success[T any] struct {
	v T
}

func (r success[T]) IsSuccess() bool {
	return true
}
func (r success[T]) Get() T {
	return r.v
}

func (r success[T]) Unapply() (T, error) {
	return r.v, nil
}

func (r success[T]) Foreach(f func(v T)) {
	f(r.v)
}
func (r success[T]) Failed() fp.Try[error] {
	return failure[error]{fp.ErrTryNotFailed}
}
func (r success[T]) OrElse(t T) T {
	return r.v
}
func (r success[T]) OrElseGet(func() T) T {
	return r.v
}
func (r success[T]) Or(func() fp.Try[T]) fp.Try[T] {
	return r
}
func (r success[T]) Recover(func(err error) T) fp.Try[T] {
	return r

}
func (r success[T]) RecoverWith(func(err error) fp.Try[T]) fp.Try[T] {
	return r
}
func (r success[T]) ToOption() fp.Option[T] {
	return option.Some(r.v)
}

func (r success[T]) String() string {
	return fmt.Sprintf("Success(%v)", r.Get())
}

type failure[T any] struct {
	err error
}

func (r failure[T]) IsSuccess() bool {
	return false
}
func (r failure[T]) Get() T {
	panic("not possible")
}

func (r failure[T]) Unapply() (T, error) {
	var zero T
	return zero, r.err
}

func (r failure[T]) Foreach(f func(v T)) {

}
func (r failure[T]) Failed() fp.Try[error] {
	return success[error]{r.err}
}
func (r failure[T]) OrElse(t T) T {
	return t
}
func (r failure[T]) OrElseGet(f func() T) T {
	return f()
}
func (r failure[T]) Or(f func() fp.Try[T]) fp.Try[T] {
	return f()
}
func (r failure[T]) Recover(f func(err error) T) fp.Try[T] {
	return success[T]{f(r.err)}

}
func (r failure[T]) RecoverWith(f func(err error) fp.Try[T]) fp.Try[T] {
	return f(r.err)
}
func (r failure[T]) ToOption() fp.Option[T] {
	return option.None[T]()
}

func (r failure[T]) String() string {
	return fmt.Sprintf("Failure(%v)", r.err)
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

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R] {
	return ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R]{Success(hlist.Empty()), Success(fn)}
}

func Func0[R any](f func() (R, error)) fp.Func0[fp.Try[R]] {
	return func() fp.Try[R] {
		ret, err := f()
		return Apply(ret, err)
	}
}
