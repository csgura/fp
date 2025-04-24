package seqt

import (
	"iter"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/xtr"
)

func Of[A any](v ...A) fp.SeqT[A] {
	return try.Success(seq.Of(v...))
}

func Failure[A any](err error) fp.SeqT[A] {
	return try.Failure[fp.Seq[A]](err)
}

func Iterator[T any](optionT fp.SeqT[T]) fp.Iterator[T] {
	return iterator.FlatMap(try.Iterator(optionT), seq.Iterator)
}

func All[T any](optionT fp.SeqT[T]) iter.Seq[T] {
	return Iterator(optionT).All()
}

func MapKey[KA, KB, V any](s fp.SeqT[fp.Tuple2[KA, V]], f func(KA) KB) fp.SeqT[fp.Tuple2[KB, V]] {
	return Map(s, func(v fp.Tuple2[KA, V]) fp.Tuple2[KB, V] {
		return fp.Tuple2[KB, V]{
			I1: f(v.I1),
			I2: v.I2,
		}
	})
}

func FilterMapKey[KA, KB, V any](s fp.SeqT[fp.Tuple2[KA, V]], f func(KA) fp.Option[KB]) fp.SeqT[fp.Tuple2[KB, V]] {
	return FilterMap(s, func(v fp.Tuple2[KA, V]) fp.Option[fp.Tuple2[KB, V]] {
		return option.Zip(f(v.I1), option.Some(v.I2))
	})
}

func MapValue[K, VA, VB any](s fp.SeqT[fp.Tuple2[K, VA]], f func(VA) VB) fp.SeqT[fp.Tuple2[K, VB]] {
	return Map(s, func(v fp.Tuple2[K, VA]) fp.Tuple2[K, VB] {
		return fp.Tuple2[K, VB]{
			I1: v.I1,
			I2: f(v.I2),
		}
	})
}

func FilterMapValue[K, VA, VB any](s fp.SeqT[fp.Tuple2[K, VA]], f func(VA) fp.Option[VB]) fp.SeqT[fp.Tuple2[K, VB]] {
	return FilterMap(s, func(v fp.Tuple2[K, VA]) fp.Option[fp.Tuple2[K, VB]] {
		return option.Zip(option.Some(v.I1), f(v.I2))
	})
}

func PartitionEithers[L, R any](r fp.SeqT[fp.Either[L, R]]) (fp.SeqT[L], fp.SeqT[R]) {
	ret := try.Map(r, func(a fp.Seq[fp.Either[L, R]]) fp.Tuple2[fp.Seq[L], fp.Seq[R]] {
		return as.Tuple(seq.PartitionEithers(a))
	})

	return try.Map(ret, xtr.Head), try.Map(ret, xtr.Last)
}

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
			seq.FilterMap[T, U],
		},
	}
}

func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.SeqT[B]) fp.SeqT[B] {

	/*
		아 하스켈 코드 이해하기 빡세네

		// foldlM :: (Foldable t, Monad m) => (b -> a -> m b) -> b -> t a -> m b
		// foldlM f z0 xs = foldr c return xs z0
		// -- See Note [List fusion and continuations in 'c']
		// where c x k z = f z x >>= k

		type K = fp.Func1[B, fp.SeqT[B]]

		// 하스켈 foldr 은 대충 다음과 같은 시그니쳐
		foldr := as.Curried3(func(f func(A, K) K, z K, s fp.Iterator[A]) K {
			panic("")
		})

		// c 는  a -> ( b -> m b ) -> b -> m b  타입
		// k 가  b -> mb 에 해당
		c := func(x A, k K) K {
			return func(z B) fp.SeqT[B] {
				return FlatMap(f(z, x), k)
			}
		}
		return foldr(c)(Pure)(s)(zero)
	*/
	/*
		type K = fp.Func1[B, fp.SeqT[B]]
		return iterator.Fold(s, Pure[B], func(k K, x A) K {
			return func(z B) fp.SeqT[B] {
				return FlatMap(f(z, x), k)
			}
		})(zero)
	*/

	return iterator.Fold(s, Pure(zero), func(b fp.SeqT[B], a A) fp.SeqT[B] {
		return FlatMap(b, fp.Flip2(f)(a))
	})
}

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.SeqT[A]] {
	return genfp.GenerateMonadFunctions[fp.SeqT[A]]{
		File:     "seqt_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}
