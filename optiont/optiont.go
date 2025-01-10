package optiont

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func Some[A any](v A) fp.OptionT[A] {
	return Pure(v)
}

func None[A any]() fp.OptionT[A] {
	return try.Success(option.None[A]())
}

func Fold[T any, U any](optionT fp.OptionT[T], zero U, f func(U, T) U) U {
	if optionT.IsFailure() {
		return zero
	}
	return option.Fold(optionT.Get(), zero, f)
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.OptionT[T]] {
	return genfp.GenerateMonadTransformer[fp.OptionT[T]]{
		File:     "optiont_op.go",
		TypeParm: genfp.TypeOf[T](),
		GivenMonad: genfp.MonadFunctions{
			Pure:    try.Pure[T],
			FlatMap: try.FlatMap[T, U],
		},
		ExposureMonad: genfp.MonadFunctions{
			Pure:    option.Pure[T],
			FlatMap: option.FlatMap[T, U],
		},
		Sequence: func(v fp.Option[fp.Try[T]]) fp.OptionT[T] {
			if v.IsDefined() {
				return try.Map(v.Get(), option.Some)
			}
			return try.Success(fp.Option[T]{})
		},
		Transform: []any{
			fp.Option[T].Filter,
			fp.Option[T].OrElse,
			fp.Option[T].OrZero,
			fp.Option[T].OrElseGet,
			fp.Option[T].Or,
			fp.Option[T].OrOption,
			fp.Option[T].OrPtr,
			fp.Option[T].Recover,
		},
	}
}

func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.OptionT[B]) fp.OptionT[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsSuccess() && t.Get().IsDefined() {
			sum = t.Get().Get()
		} else {
			return t
		}
	}
	return Pure(sum)
}

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.OptionT[A]] {
	return genfp.GenerateMonadFunctions[fp.OptionT[A]]{
		File:     "optiont_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateTraverseFunctions[fp.OptionT[A]] {
	return genfp.GenerateTraverseFunctions[fp.OptionT[A]]{
		File:     "optiont_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}
