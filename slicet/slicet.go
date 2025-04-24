package slicet

import (
	"iter"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/slice"
	"github.com/csgura/fp/try"
)

func Of[A any](v ...A) fp.SliceT[A] {
	return try.Success(slice.Of(v...))
}

func Failure[A any](err error) fp.SliceT[A] {
	return try.Failure[fp.Slice[A]](err)
}

func Iterator[T any](optionT fp.SliceT[T]) fp.Iterator[T] {
	return iterator.FlatMap(try.Iterator(optionT), slice.Iterator)
}

func All[T any](optionT fp.SliceT[T]) iter.Seq[T] {
	return Iterator(optionT).All()
}

func MapKey[KA, KB, V any](s fp.SliceT[fp.Tuple2[KA, V]], f func(KA) KB) fp.SliceT[fp.Tuple2[KB, V]] {
	return Map(s, func(v fp.Tuple2[KA, V]) fp.Tuple2[KB, V] {
		return fp.Tuple2[KB, V]{
			I1: f(v.I1),
			I2: v.I2,
		}
	})
}

func FilterMapKey[KA, KB, V any](s fp.SliceT[fp.Tuple2[KA, V]], f func(KA) fp.Option[KB]) fp.SliceT[fp.Tuple2[KB, V]] {
	return FilterMap(s, func(v fp.Tuple2[KA, V]) fp.Option[fp.Tuple2[KB, V]] {
		return option.Zip(f(v.I1), option.Some(v.I2))
	})
}

func MapValue[K, VA, VB any](s fp.SliceT[fp.Tuple2[K, VA]], f func(VA) VB) fp.SliceT[fp.Tuple2[K, VB]] {
	return Map(s, func(v fp.Tuple2[K, VA]) fp.Tuple2[K, VB] {
		return fp.Tuple2[K, VB]{
			I1: v.I1,
			I2: f(v.I2),
		}
	})
}

func FilterMapValue[K, VA, VB any](s fp.SliceT[fp.Tuple2[K, VA]], f func(VA) fp.Option[VB]) fp.SliceT[fp.Tuple2[K, VB]] {
	return FilterMap(s, func(v fp.Tuple2[K, VA]) fp.Option[fp.Tuple2[K, VB]] {
		return option.Zip(option.Some(v.I1), f(v.I2))
	})
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.SliceT[T]] {
	return genfp.GenerateMonadTransformer[fp.SliceT[T]]{
		File:     "slicet_op.go",
		TypeParm: genfp.TypeOf[T](),
		GivenMonad: genfp.MonadFunctions{
			Pure: try.Success[T],
		},
		ExposureMonad: genfp.MonadFunctions{
			Pure:    slice.Pure[T],
			FlatMap: slice.FlatMap[T, U],
		},
		Sequence: func(v fp.Slice[fp.Try[T]]) fp.SliceT[T] {
			return try.FoldM(iterator.FromSeq(v), fp.Slice[T]{}, func(t1 fp.Slice[T], t2 fp.Try[T]) fp.SliceT[T] {
				return try.Map(t2, func(v T) fp.Slice[T] {
					return append(t1, v)
				})
			})

		},
		Transform: []any{
			slice.Filter[T],
			slice.Add[T],
			slice.Append[T],
			slice.Concat[T],
			slice.Drop[T],
			slice.Exists[T],
			slice.FilterNot[T],
			slice.Find[T],
			slice.ForAll[T],
			slice.Foreach[T],
			slice.Get[T],
			slice.Head[T],
			slice.Tail[T],
			slice.Init[T],
			slice.IsEmpty[T],
			slice.Last[T],
			slice.MakeString[T],
			slice.NonEmpty[T],
			slice.Reverse[T],
			slice.Size[T],
			slice.Take[T],
			slice.Fold[T, U],
			slice.Scan[T, U],
			slice.Sort[T],
			slice.Min[T],
			slice.Max[T],
			slice.FilterMap[T, U],
		},
	}
}

func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.SliceT[B]) fp.SliceT[B] {

	/*
		아 하스켈 코드 이해하기 빡세네

		// foldlM :: (Foldable t, Monad m) => (b -> a -> m b) -> b -> t a -> m b
		// foldlM f z0 xs = foldr c return xs z0
		// -- See Note [List fusion and continuations in 'c']
		// where c x k z = f z x >>= k

		type K = fp.Func1[B, fp.SliceT[B]]

		// 하스켈 foldr 은 대충 다음과 같은 시그니쳐
		foldr := as.Curried3(func(f func(A, K) K, z K, s fp.Iterator[A]) K {
			panic("")
		})

		// c 는  a -> ( b -> m b ) -> b -> m b  타입
		// k 가  b -> mb 에 해당
		c := func(x A, k K) K {
			return func(z B) fp.SliceT[B] {
				return FlatMap(f(z, x), k)
			}
		}
		return foldr(c)(Pure)(s)(zero)
	*/
	/*
		type K = fp.Func1[B, fp.SliceT[B]]
		return iterator.Fold(s, Pure[B], func(k K, x A) K {
			return func(z B) fp.SliceT[B] {
				return FlatMap(f(z, x), k)
			}
		})(zero)
	*/

	return iterator.Fold(s, Pure(zero), func(b fp.SliceT[B], a A) fp.SliceT[B] {
		return FlatMap(b, fp.Flip2(f)(a))
	})
}

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.SliceT[A]] {
	return genfp.GenerateMonadFunctions[fp.SliceT[A]]{
		File:     "slicet_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}
