// Code generated by monad_gen, DO NOT EDIT.
package either

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/iterator"
)

func Traverse[L any, A any, R any](ia fp.Iterator[A], fn func(A) fp.Either[L, R]) fp.Either[L, fp.Iterator[R]] {
	return Map(FoldM(ia, fp.Seq[R]{}, func(acc fp.Seq[R], a A) fp.Either[L, fp.Seq[R]] {
		return Map(fn(a), acc.Add)
	}), iterator.FromSeq)
}

func TraverseSeq[L any, A any, R any](sa fp.Seq[A], fa func(A) fp.Either[L, R]) fp.Either[L, fp.Seq[R]] {
	return FoldM(fp.IteratorOfSeq(sa), fp.Seq[R]{}, func(acc fp.Seq[R], a A) fp.Either[L, fp.Seq[R]] {
		return Map(fa(a), acc.Add)
	})
}

func TraverseSlice[L any, A any, R any](sa []A, fa func(A) fp.Either[L, R]) fp.Either[L, []R] {
	return Map(TraverseSeq(sa, fa), fp.Seq[R].Widen)
}

func TraverseFunc[L any, A any, R any](far func(A) fp.Either[L, R]) func(fp.Iterator[A]) fp.Either[L, fp.Iterator[R]] {
	return func(iterA fp.Iterator[A]) fp.Either[L, fp.Iterator[R]] {
		return Traverse(iterA, far)
	}
}

func TraverseSeqFunc[L any, A any, R any](far func(A) fp.Either[L, R]) func(fp.Seq[A]) fp.Either[L, fp.Seq[R]] {
	return func(seqA fp.Seq[A]) fp.Either[L, fp.Seq[R]] {
		return TraverseSeq(seqA, far)
	}
}

func TraverseSliceFunc[L any, A any, R any](far func(A) fp.Either[L, R]) func([]A) fp.Either[L, []R] {
	return func(seqA []A) fp.Either[L, []R] {
		return TraverseSlice(seqA, far)
	}
}

func FlatMapTraverseSeq[L any, A any, B any](ta fp.Either[L, fp.Seq[A]], f func(v A) fp.Either[L, B]) fp.Either[L, fp.Seq[B]] {
	return FlatMap(ta, TraverseSeqFunc(f))
}

func FlatMapTraverseSlice[L any, A any, B any](ta fp.Either[L, []A], f func(v A) fp.Either[L, B]) fp.Either[L, []B] {
	return FlatMap(ta, TraverseSliceFunc(f))
}

func Sequence[L any, A any](tsa []fp.Either[L, A]) fp.Either[L, []A] {
	ret := FoldM(iterator.FromSeq(tsa), fp.Seq[A]{}, func(t1 fp.Seq[A], t2 fp.Either[L, A]) fp.Either[L, fp.Seq[A]] {
		return Map(t2, t1.Add)
	})

	return Map(ret, fp.Seq[A].Widen)
}

func SequenceIterator[L any, A any](ita fp.Iterator[fp.Either[L, A]]) fp.Either[L, fp.Iterator[A]] {
	ret := FoldM(ita, fp.Seq[A]{}, func(t1 fp.Seq[A], t2 fp.Either[L, A]) fp.Either[L, fp.Seq[A]] {
		return Map(t2, t1.Add)
	})
	return Map(ret, iterator.FromSeq)

}
