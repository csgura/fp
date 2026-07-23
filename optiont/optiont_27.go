//go:build go1.27

package optiont

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

type Type[T any] fp.OptionT[T]

func Trans[T any](v fp.OptionT[T]) Type[T] {
	return Type[T](v)
}

func (r Type[T]) Try() fp.OptionT[T] {
	return fp.OptionT[T](r)
}
func (r Type[T]) Map[R any](f func(T) R) Type[R] {
	return Trans(Map(r.Try(), f))
}

func (r Type[T]) FlatMap[R any](f func(T) Type[R]) Type[R] {
	return Trans(FlatMap[T, R](r.Try(), func(v T) fp.OptionT[R] {
		return f(v).Try()
	}))
}

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.OptionT[T]] {
	return genfp.GenerateMonadTransformer[fp.OptionT[T]]{
		File:     "optiont_op.go",
		TypeParm: genfp.TypeOf[T](),
		GivenMonad: genfp.MonadFunctions{
			Pure: try.Success[T],
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
			fp.Option[T].IsDefined,
			fp.Option[T].IsEmpty[fp.Phantom[T]],
			fp.Option[T].Filter[fp.Phantom[T]],
			fp.Option[T].OrElse,
			fp.Option[T].OrZero,
			fp.Option[T].OrElseGet[fp.Phantom[T]],
			fp.Option[T].Or[fp.Phantom[T]],
			fp.Option[T].OrOption[fp.Phantom[T]],
			fp.Option[T].OrPtr[fp.Phantom[T]],
			fp.Option[T].Recover,
			fp.Option[T].Foreach[fp.Phantom[T]],
		},
	}
}
