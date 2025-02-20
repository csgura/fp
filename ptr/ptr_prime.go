package ptr

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

func Map[A, B any](pa fp.Ptr[A], f func(A) B) fp.Ptr[B] {
	if pa == nil {
		return nil
	}
	ret := f(*pa)
	return &ret
}

func FlatMap[A, B any](pa fp.Ptr[A], f func(A) fp.Ptr[B]) fp.Ptr[B] {
	if pa == nil {
		return nil
	}
	ret := f(*pa)
	return ret
}

// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.Ptr[B]) fp.Ptr[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t != nil {
			sum = *t
		} else {
			return t
		}
	}
	return Pure(sum)
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.Ptr[A]] {
	return genfp.GenerateMonadFunctions[fp.Ptr[A]]{
		File:     "ptr_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateTraverseFunctions[fp.Ptr[A]] {
	return genfp.GenerateTraverseFunctions[fp.Ptr[A]]{
		File:     "ptr_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}
