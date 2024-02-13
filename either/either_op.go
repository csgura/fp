package either

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[L, R any]() genfp.GenerateMonadFunctions[fp.Either[L, R]] {
	return genfp.GenerateMonadFunctions[fp.Either[L, R]]{
		File:     "either_monad.go",
		TypeParm: genfp.TypeOf[R](),
	}
}

// @internal.Generate
func _[L, R any]() genfp.GenerateTraverseFunctions[fp.Either[L, R]] {
	return genfp.GenerateTraverseFunctions[fp.Either[L, R]]{
		File:     "either_traverse.go",
		TypeParm: genfp.TypeOf[R](),
	}
}

func Left[L, R any](l L) fp.Either[L, R] {
	return fp.Left[L, R](l)
}

func NotRight[R, L any](l L) fp.Either[L, R] {
	return fp.Left[L, R](l)
}

func Right[L, R any](r R) fp.Either[L, R] {
	return fp.Right[L, R](r)
}

func Pure[L, R any](r R) fp.Either[L, R] {
	return fp.Right[L, R](r)
}

func Swap[L, R any](e fp.Either[L, R]) fp.Either[R, L] {
	if e.IsRight() {
		return Left[R, L](e.Get())
	}
	return Right[R](e.Left())
}

func FlatMap[L, R, R1 any](opt fp.Either[L, R], fn func(v R) fp.Either[L, R1]) fp.Either[L, R1] {
	if opt.IsRight() {
		return fn(opt.Get())
	}
	return Left[L, R1](opt.Left())
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

// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldM[L, R, B any](s fp.Iterator[R], zero B, f func(B, R) fp.Either[L, B]) fp.Either[L, B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsRight() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return Right[L](sum)
}
