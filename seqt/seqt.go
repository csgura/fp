package seqt

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
)

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.SeqT[T]] {
	return genfp.GenerateMonadTransformer[fp.SeqT[T]]{
		File:     "seqt_op.go",
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
			fp.Seq[T].Filter,
			fp.Seq[T].Add,
			fp.Seq[T].Append,
			fp.Seq[T].Concat,
			fp.Seq[T].Drop,
			fp.Seq[T].Exists,
			fp.Seq[T].FilterNot,
			fp.Seq[T].Find,
			fp.Seq[T].ForAll,
			fp.Seq[T].Foreach,
			fp.Seq[T].Get,
			fp.Seq[T].Head,
			fp.Seq[T].Tail,
			fp.Seq[T].Init,
			fp.Seq[T].IsEmpty,
			fp.Seq[T].Last,
			fp.Seq[T].MakeString,
			fp.Seq[T].NonEmpty,
			fp.Seq[T].Reverse,
			fp.Seq[T].Size,
			fp.Seq[T].Take,
			seq.Fold[T, U],
			seq.Scan[T, U],
			seq.Sort[T],
			seq.Min[T],
			seq.Max[T],
		},
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.SeqT[A]] {
	return genfp.GenerateMonadFunctions[fp.SeqT[A]]{
		File:     "seqt_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}
