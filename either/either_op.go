package either

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
)

func Left[L, R any](l L) fp.Either[L, R] {
	return left[L, R]{l}
}

func Right[L, R any](r R) fp.Either[L, R] {
	return right[L, R]{r}
}

func Ap[L, R, R1 any](t fp.Either[L, fp.Func1[R, R1]], a fp.Either[L, R]) fp.Either[L, R1] {
	return FlatMap(t, func(f fp.Func1[R, R1]) fp.Either[L, R1] {
		return Map(a.(fp.Either[L, R]), func(a R) R1 {
			return f(a)
		})
	})
}

func Map[L, R, R1 any](opt fp.Either[L, R], f func(v R) R1) fp.Either[L, R1] {
	return FlatMap(opt, func(v R) fp.Either[L, R1] {
		return Right[L](f(v))
	})
}

func FlatMap[L, R, R1 any](opt fp.Either[L, R], fn func(v R) fp.Either[L, R1]) fp.Either[L, R1] {
	if opt.IsRight() {
		return fn(opt.Get())
	}
	return Left[L, R1](opt.Left().Get())
}

func Flatten[L, R any](opt fp.Either[L, fp.Either[L, R]]) fp.Either[L, R] {
	return FlatMap(opt, func(v fp.Either[L, R]) fp.Either[L, R] {
		return v
	})
}

type left[L, R any] struct {
	v L
}

func (r left[L, R]) IsLeft() bool {
	return true
}
func (r left[L, R]) IsRight() bool {
	return false
}
func (r left[L, R]) Left() fp.Option[L] {
	return option.Some(r.v)
}
func (r left[L, R]) Right() fp.Option[R] {
	return option.None[R]()
}
func (r left[L, R]) Swap() fp.Either[R, L] {
	return Right[R](r.v)
}
func (r left[L, R]) Get() R {
	panic("Either.left")
}
func (r left[L, R]) Foreach(f func(v R)) {

}
func (r left[L, R]) OrElse(t R) R {
	return t
}
func (r left[L, R]) OrElseGet(f func() R) R {
	return f()
}
func (r left[L, R]) Exists(p func(v R) bool) bool {
	return false
}
func (r left[L, R]) ForAll(p func(v R) bool) bool {
	return true
}

type right[L, R any] struct {
	v R
}

func (r right[L, R]) IsLeft() bool {
	return false
}
func (r right[L, R]) IsRight() bool {
	return true
}
func (r right[L, R]) Left() fp.Option[L] {
	return option.None[L]()
}
func (r right[L, R]) Right() fp.Option[R] {
	return option.Some(r.v)
}
func (r right[L, R]) Swap() fp.Either[R, L] {
	return Left[R, L](r.v)
}
func (r right[L, R]) Get() R {
	return r.v
}
func (r right[L, R]) Foreach(f func(v R)) {
	f(r.v)
}
func (r right[L, R]) OrElse(t R) R {
	return r.v
}
func (r right[L, R]) OrElseGet(f func() R) R {
	return r.v
}
func (r right[L, R]) Exists(p func(v R) bool) bool {
	return p(r.v)
}
func (r right[L, R]) ForAll(p func(v R) bool) bool {
	return p(r.v)
}
