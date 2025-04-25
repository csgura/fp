package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
func _[T, U, V any, K comparable]() genfp.GenerateMonadTransformer[fp.SeqT[T]] {
	return genfp.GenerateMonadTransformer[fp.SeqT[T]]{
		File:     "gendebug_generated.go",
		TypeParm: genfp.TypeOf[T](),
		GivenMonad: genfp.MonadFunctions{
			Pure: try.Success[T],
		},
		ExposureMonad: genfp.MonadFunctions{
			Pure:    seq.Pure[T],
			FlatMap: seq.FlatMap[T, U],
		},
		Sequence: func(v fp.Seq[fp.Try[T]]) fp.SeqT[T] {
			return try.FoldM(iterator.FromSeq(v), fp.Seq[T]{}, func(t1 fp.Seq[T], t2 fp.Try[T]) fp.SeqT[T] {
				return try.Map(t2, t1.Add)
			})

		},
		Transform: []any{
			seq.ToGoMap[K, V],
			seq.FoldTry[T, U],
		},
	}
}
