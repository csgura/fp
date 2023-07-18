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
		return Map(a, f)
	})
}

func Map[L, R, R1 any](opt fp.Either[L, R], f func(v R) R1) fp.Either[L, R1] {
	return FlatMap(opt, func(v R) fp.Either[L, R1] {
		return Right[L](f(v))
	})
}

func Map2[L, R1, R2, R3 any](a fp.Either[L, R1], b fp.Either[L, R2], f func(R1, R2) R3) fp.Either[L, R3] {
	return FlatMap(a, func(v1 R1) fp.Either[L, R3] {
		return Map(b, func(v2 R2) R3 {
			return f(v1, v2)
		})
	})
}

func Compose[L, A, B, C any](f1 func(A) fp.Either[L, B], f2 func(B) fp.Either[L, C]) func(A) fp.Either[L, C] {
	return func(a A) fp.Either[L, C] {
		return FlatMap(f1(a), f2)
	}
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

func Fold[L, R, V any](e fp.Either[L, R], fl func(L) V, fr func(R) V) V {
	if e.IsLeft() {
		return fl(e.Left().Get())
	}
	return fr(e.Right().Get())
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
