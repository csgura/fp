//go:build !go1.27

package optiont

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

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
			fp.Option[T].IsEmpty,
			fp.Option[T].Filter,
			fp.Option[T].OrElse,
			fp.Option[T].OrZero,
			fp.Option[T].OrElseGet,
			fp.Option[T].Or,
			fp.Option[T].OrOption,
			fp.Option[T].OrPtr,
			fp.Option[T].Recover,
			fp.Option[T].Foreach,
		},
	}
}
