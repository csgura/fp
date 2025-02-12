// Code generated by monad_gen, DO NOT EDIT.
package optiont

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/iterator"
)

func Traverse[A any, R any](ia fp.Iterator[A], fn func(A) fp.OptionT[R]) fp.OptionT[fp.Iterator[R]] {
	return Map(FoldM(ia, fp.Seq[R]{}, func(acc fp.Seq[R], a A) fp.OptionT[fp.Seq[R]] {
		return Map(fn(a), acc.Add)
	}), iterator.FromSeq)
}

func TraverseSeq[A any, R any](sa fp.Seq[A], fa func(A) fp.OptionT[R]) fp.OptionT[fp.Seq[R]] {
	return FoldM(fp.IteratorOfSeq(sa), fp.Seq[R]{}, func(acc fp.Seq[R], a A) fp.OptionT[fp.Seq[R]] {
		return Map(fa(a), acc.Add)
	})
}

func TraverseSlice[A any, R any](sa []A, fa func(A) fp.OptionT[R]) fp.OptionT[[]R] {
	return Map(TraverseSeq(sa, fa), fp.Seq[R].Widen)
}

func TraverseFunc[A any, R any](far func(A) fp.OptionT[R]) func(fp.Iterator[A]) fp.OptionT[fp.Iterator[R]] {
	return func(iterA fp.Iterator[A]) fp.OptionT[fp.Iterator[R]] {
		return Traverse(iterA, far)
	}
}

func TraverseSeqFunc[A any, R any](far func(A) fp.OptionT[R]) func(fp.Seq[A]) fp.OptionT[fp.Seq[R]] {
	return func(seqA fp.Seq[A]) fp.OptionT[fp.Seq[R]] {
		return TraverseSeq(seqA, far)
	}
}

func TraverseSliceFunc[A any, R any](far func(A) fp.OptionT[R]) func([]A) fp.OptionT[[]R] {
	return func(seqA []A) fp.OptionT[[]R] {
		return TraverseSlice(seqA, far)
	}
}

func FlatMapTraverseSeq[A any, B any](ta fp.OptionT[fp.Seq[A]], f func(v A) fp.OptionT[B]) fp.OptionT[fp.Seq[B]] {
	return FlatMap(ta, TraverseSeqFunc(f))
}

func FlatMapTraverseSlice[A any, B any](ta fp.OptionT[[]A], f func(v A) fp.OptionT[B]) fp.OptionT[[]B] {
	return FlatMap(ta, TraverseSliceFunc(f))
}

func Sequence[A any](tsa []fp.OptionT[A]) fp.OptionT[[]A] {
	ret := FoldM(iterator.FromSeq(tsa), fp.Seq[A]{}, func(t1 fp.Seq[A], t2 fp.OptionT[A]) fp.OptionT[fp.Seq[A]] {
		return Map(t2, t1.Add)
	})

	return Map(ret, fp.Seq[A].Widen)
}

func SequenceIterator[A any](ita fp.Iterator[fp.OptionT[A]]) fp.OptionT[fp.Iterator[A]] {
	ret := FoldM(ita, fp.Seq[A]{}, func(t1 fp.Seq[A], t2 fp.OptionT[A]) fp.OptionT[fp.Seq[A]] {
		return Map(t2, t1.Add)
	})
	return Map(ret, iterator.FromSeq)

}
