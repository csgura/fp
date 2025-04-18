// Code generated by monad_gen, DO NOT EDIT.
package seqt

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
)

func Pure[A any](a A) fp.SeqT[A] {
	return try.Success[fp.Seq[A]](seq.Pure[A](a))
}

func LiftT[A any](a fp.Try[A]) fp.SeqT[A] {
	return try.Map(a, seq.Pure[A])
}

func Map[A any, B any](t fp.SeqT[A], f func(A) B) fp.SeqT[B] {
	return try.Map(t, func(ma fp.Seq[A]) fp.Seq[B] {
		return seq.FlatMap[A, B](ma, func(a A) fp.Seq[B] {
			return seq.Pure[B](f(a))
		})
	})
}

func SubFlatMap[A any, B any](t fp.SeqT[A], f func(A) fp.Seq[B]) fp.SeqT[B] {
	return try.Map(t, func(ma fp.Seq[A]) fp.Seq[B] {
		return seq.FlatMap[A, B](ma, func(a A) fp.Seq[B] {
			return f(a)
		})
	})
}

func MapT[A any, B any](t fp.SeqT[A], f func(A) fp.Try[B]) fp.SeqT[B] {
	sequencef := func(v fp.Seq[fp.Try[B]]) fp.SeqT[B] {
		return try.FoldM(iterator.FromSeq(v), fp.Seq[B]{}, func(t1 fp.Seq[B], t2 fp.Try[B]) fp.SeqT[B] {
			return try.Map(t2, t1.Add)
		})
	}
	return try.FlatMap(Map(t, f), sequencef)
}

func FlatMap[A any, B any](t fp.SeqT[A], f func(A) fp.SeqT[B]) fp.SeqT[B] {

	flatten := func(v fp.Seq[fp.Seq[B]]) fp.Seq[B] {
		return seq.FlatMap[fp.Seq[B], B](v, fp.Id)
	}

	return try.Map(MapT(t, f), flatten)

}

func Filter[T any](seqT fp.SeqT[T], p func(v T) bool) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Filter(insideValue, p)
	})
}

func Add[T any](seqT fp.SeqT[T], item T) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Add(insideValue, item)
	})
}

func Append[T any](seqT fp.SeqT[T], items T) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Append(insideValue, items)
	})
}

func Concat[T any](seqT fp.SeqT[T], tail fp.Seq[T]) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Concat(insideValue, tail)
	})
}

func Drop[T any](seqT fp.SeqT[T], n int) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Drop(insideValue, n)
	})
}

func Exists[T any](seqT fp.SeqT[T], p func(v T) bool) fp.Try[bool] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) bool {
		return fp.Seq[T].Exists(insideValue, p)
	})
}

func FilterNot[T any](seqT fp.SeqT[T], p func(v T) bool) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].FilterNot(insideValue, p)
	})
}

func Find[T any](seqT fp.SeqT[T], p func(v T) bool) fp.Try[fp.Option[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Option[T] {
		return fp.Seq[T].Find(insideValue, p)
	})
}

func ForAll[T any](seqT fp.SeqT[T], p func(v T) bool) fp.Try[bool] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) bool {
		return fp.Seq[T].ForAll(insideValue, p)
	})
}

func Foreach[T any](seqT fp.SeqT[T], f func(v T)) {
	try.Map(seqT, func(insideValue fp.Seq[T]) error {
		fp.Seq[T].Foreach(insideValue, f)
		return nil
	})
}

func Get[T any](seqT fp.SeqT[T], idx int) fp.Try[fp.Option[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Option[T] {
		return fp.Seq[T].Get(insideValue, idx)
	})
}

func Head[T any](seqT fp.SeqT[T]) fp.Try[fp.Option[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Option[T] {
		return fp.Seq[T].Head(insideValue)
	})
}

func Tail[T any](seqT fp.SeqT[T]) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Tail(insideValue)
	})
}

func Init[T any](seqT fp.SeqT[T]) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Init(insideValue)
	})
}

func IsEmpty[T any](seqT fp.SeqT[T]) fp.Try[bool] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) bool {
		return fp.Seq[T].IsEmpty(insideValue)
	})
}

func Last[T any](seqT fp.SeqT[T]) fp.Try[fp.Option[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Option[T] {
		return fp.Seq[T].Last(insideValue)
	})
}

func MakeString[T any](seqT fp.SeqT[T], sep string) fp.Try[string] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) string {
		return fp.Seq[T].MakeString(insideValue, sep)
	})
}

func NonEmpty[T any](seqT fp.SeqT[T]) fp.Try[bool] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) bool {
		return fp.Seq[T].NonEmpty(insideValue)
	})
}

func Reverse[T any](seqT fp.SeqT[T]) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Reverse(insideValue)
	})
}

func Size[T any](seqT fp.SeqT[T]) fp.Try[int] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) int {
		return fp.Seq[T].Size(insideValue)
	})
}

func Take[T any](seqT fp.SeqT[T], n int) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return fp.Seq[T].Take(insideValue, n)
	})
}

func Fold[T any, U any](seqT fp.SeqT[T], zero U, f func(U, T) U) fp.Try[U] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) U {
		return seq.Fold[T, U](insideValue, zero, f)
	})
}

func Scan[T any, U any](seqT fp.SeqT[T], zero U, f func(U, T) U) fp.Try[fp.Seq[U]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[U] {
		return seq.Scan[T, U](insideValue, zero, f)
	})
}

func Sort[T any](seqT fp.SeqT[T], ord fp.Ord[T]) fp.Try[fp.Seq[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Seq[T] {
		return seq.Sort[T](insideValue, ord)
	})
}

func Min[T any](seqT fp.SeqT[T], ord fp.Ord[T]) fp.Try[fp.Option[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Option[T] {
		return seq.Min[T](insideValue, ord)
	})
}

func Max[T any](seqT fp.SeqT[T], ord fp.Ord[T]) fp.Try[fp.Option[T]] {
	return try.Map(seqT, func(insideValue fp.Seq[T]) fp.Option[T] {
		return seq.Max[T](insideValue, ord)
	})
}
