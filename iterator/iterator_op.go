package iterator

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
)

func List[T any](list fp.List[T]) fp.Iterator[T] {
	current := list

	return fp.MakeIterator(
		func() bool {
			return current.Head().IsDefined()
		},
		func() T {
			ret := current.Head().Get()
			current = current.Tail()
			return ret
		},
	)
}

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
	return fp.MakeIterator(
		func() bool {
			return opt.HasNext()
		},
		func() U {
			return fn(opt.Next())
		},
	)
}

func Map2[A, B, U any](a fp.Iterator[A], b fp.Iterator[B], f func(A, B) U) fp.Iterator[U] {
	return FlatMap(a, func(v1 A) fp.Iterator[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
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

	return fp.MakeIterator(
		hasNext,
		func() U {
			if hasNext() {
				return current.Get().Next()
			}
			panic("next on empty iterator")
		},
	)
}

func ToMap[K, V any](itr fp.Iterator[fp.Tuple2[K, V]], hasher fp.Hashable[K]) fp.Map[K, V] {
	ret := immutable.MapBuilder[K, V](hasher)

	for itr.HasNext() {
		k, v := itr.Next().Unapply()
		ret = ret.Add(k, v)
	}

	return fp.MakeMap(ret.Build())
}

func ToSet[V any](itr fp.Iterator[V], hasher fp.Hashable[V]) fp.Set[V] {
	ret := immutable.SetBuilder(hasher)

	for itr.HasNext() {
		v := itr.Next()
		ret = ret.Add(v)
	}

	return ret.Build()
}

func Zip[T, U any](a fp.Iterator[T], b fp.Iterator[U]) fp.Iterator[fp.Tuple2[T, U]] {
	return fp.MakeIterator(
		func() bool {
			return a.HasNext() && b.HasNext()
		},
		func() fp.Tuple2[T, U] {
			return as.Tuple(a.Next(), b.Next())
		},
	)
}

func Fold[A, B any](s fp.Iterator[A], zero B, f func(B, A) B) B {
	sum := zero
	for s.HasNext() {
		sum = f(sum, s.Next())
	}
	return sum
}

func FoldRight[A, B any](s fp.Iterator[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	head := s.Next()

	v := lazy.TailCall(func() lazy.Eval[B] {
		return FoldRight(s, zero, f)
	})

	return f(head, v)

}

func GroupBy[A any, K any](s fp.Iterator[A], keyFunc func(A) K) mutable.Map[K, fp.Seq[A]] {

	ret := map[any]fp.Seq[A]{}

	return Fold(s, ret, func(b map[any]fp.Seq[A], a A) map[any]fp.Seq[A] {
		k := keyFunc(a)
		b[k] = b[k].Append(a)
		return b
	})
}

func Scan[A, B any](s fp.Iterator[A], zero B, f func(B, A) B) fp.Iterator[B] {

	first := true
	sum := zero
	hasNext := func() bool {
		if first {
			return true
		}
		return s.HasNext()
	}
	return fp.MakeIterator(
		hasNext,
		func() B {
			if hasNext() {
				if first {
					first = false
					return sum
				}
				sum = f(sum, s.Next())
				return sum
			}
			panic("next on empty iterator")
		},
	)
}

func Range(from, exclusive int) fp.Iterator[int] {
	i := from
	return fp.MakeIterator(
		func() bool {
			return i < exclusive
		},
		func() int {
			if i < exclusive {
				ret := i
				i++
				return ret
			}
			panic("next on empty iterator")
		},
	)
}

func RangeClosed(from, inclusive int) fp.Iterator[int] {
	i := from
	return fp.MakeIterator(
		func() bool {
			return i <= inclusive
		},
		func() int {
			if i <= inclusive {
				ret := i
				i++
				return ret
			}
			panic("next on empty iterator")
		},
	)
}
