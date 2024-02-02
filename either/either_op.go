package either

import (
	"encoding/json"

	"github.com/csgura/fp"
)

func Left[L, R any](l L) fp.Either[L, R] {
	return left[L, R]{l}
}

func NotRight[R, L any](l L) fp.Either[L, R] {
	return left[L, R]{l}
}

func Right[L, R any](r R) fp.Either[L, R] {
	return right[L, R]{r}
}

func Swap[L, R any](e fp.Either[L, R]) fp.Either[R, L] {
	if e.IsRight() {
		return Left[R, L](e.Get())
	}
	return Right[R](e.Left())
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
	return Left[L, R1](opt.Left())
}

func Flatten[L, R any](opt fp.Either[L, fp.Either[L, R]]) fp.Either[L, R] {
	return FlatMap(opt, func(v fp.Either[L, R]) fp.Either[L, R] {
		return v
	})
}

func Fold[L, R, V any](e fp.Either[L, R], fl func(L) V, fr func(R) V) V {
	if e.IsLeft() {
		return fl(e.Left())
	}
	return fr(e.Get())
}

func Foreach[L, R any](e fp.Either[L, R], f func(R)) {
	if e.IsRight() {
		f(e.Get())
	}
}

func OrElse[L, R any](e fp.Either[L, R], t R) R {
	if e.IsLeft() {
		return t
	}
	return e.Get()
}

func OrElseGet[L, R any](e fp.Either[L, R], f func() R) R {
	if e.IsLeft() {
		return f()
	}
	return e.Get()
}

func Exists[L, R any](e fp.Either[L, R], p func(v R) bool) bool {
	if e.IsRight() {
		return p(e.Get())
	}
	return false
}
func ForAll[L, R any](e fp.Either[L, R], p func(v R) bool) bool {
	if e.IsRight() {
		return p(e.Get())
	}
	return true
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
func (r left[L, R]) Left() L {
	return r.v
}

func (r left[L, R]) Get() R {
	panic("Either.left")
}

func (r left[L, R]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.v)
}

func (r left[L, R]) Recover(f func() R) fp.Either[L, R] {
	return Right[L, R](f())
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
func (r right[L, R]) Left() L {
	panic("Either.right")
}

func (r right[L, R]) Get() R {
	return r.v
}

func (r right[L, R]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.v)
}

func (r right[L, R]) Recover(f func() R) fp.Either[L, R] {
	return r
}
