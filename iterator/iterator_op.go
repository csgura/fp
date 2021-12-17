package iterator

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

func Of[T any](list ...T) fp.Iterator[T] {
	return fp.Seq[T](list).Iterator()
}

func Ap[T, U any](t fp.Iterator[fp.Func1[T, U]], a fp.Iterator[T]) fp.Iterator[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Iterator[U] {
		return Map(a, f)
	})
}

func Lift[T, U any](f func(v T) U) fp.Func1[fp.Iterator[T], fp.Iterator[U]] {
	return func(opt fp.Iterator[T]) fp.Iterator[U] {
		return Map(opt, f)
	}
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.Iterator[B]], f2 fp.Func1[B, fp.Iterator[C]]) fp.Func1[A, fp.Iterator[C]] {
	return func(a A) fp.Iterator[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.Iterator[B]], f2 fp.Func1[B, C]) fp.Func1[A, fp.Iterator[C]] {
	return func(a A) fp.Iterator[C] {
		return Map(f1(a), f2)
	}
}

func Flatten[T any](opt fp.Iterator[fp.Iterator[T]]) fp.Iterator[T] {
	return FlatMap(opt, func(v fp.Iterator[T]) fp.Iterator[T] {
		return v
	})
}

func Concat[T any](head T, tail fp.Iterator[T]) fp.Iterator[T] {
	return Of(head).Concat(tail)
}

func Map[T, U any](opt fp.Iterator[T], fn func(v T) U) fp.Iterator[U] {
	return fp.IteratorAdaptor[U]{
		IsHasNext: func() bool {
			return opt.HasNext()
		},
		GetNext: func() U {
			return fn(opt.Next())
		},
	}
}

func FlatMap[T, U any](opt fp.Iterator[T], fn func(v T) fp.Iterator[U]) fp.Iterator[U] {

	current := option.None[fp.Iterator[U]]()

	hasNext := func() bool {
		if option.Map(current, fp.Iterator[U].HasNext).OrElse(false) == true {
			return true
		}

		for opt.HasNext() {
			nextItr := fn(opt.Next())
			current = option.Some(nextItr)
			if nextItr.HasNext() {
				return true
			}
		}

		return false
	}

	return fp.IteratorAdaptor[U]{
		IsHasNext: func() bool {
			return hasNext()
		},
		GetNext: func() U {
			if hasNext() {
				return current.Get().Next()
			}
			panic("next on empty iterator")
		},
	}
}

func ToMap[K comparable, V any](itr fp.Iterator[fp.Tuple2[K, V]]) fp.Map[K, V] {
	ret := fp.Map[K, V]{}

	for itr.HasNext() {
		k, v := itr.Next().Unapply()
		ret[k] = v
	}

	return ret
}

func Zip[T, U any](a fp.Iterator[T], b fp.Iterator[U]) fp.Iterator[fp.Tuple2[T, U]] {
	return fp.IteratorAdaptor[fp.Tuple2[T, U]]{
		IsHasNext: func() bool {
			return a.HasNext() && b.HasNext()
		},
		GetNext: func() fp.Tuple2[T, U] {
			return as.Tuple(a.Next(), b.Next())
		},
	}
}

func Fold[A, B any](s fp.Iterator[A], zero B, f func(B, A) B) B {
	sum := zero
	for s.HasNext() {
		sum = f(sum, s.Next())
	}
	return sum
}
