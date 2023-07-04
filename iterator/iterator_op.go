//go:generate go run github.com/csgura/fp/internal/generator/itr_gen
package iterator

import (
	"sync"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
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

func Empty[T any]() fp.Iterator[T] {
	return fp.MakeIterator(func() bool {
		return false
	}, func() T {
		panic("next on empty iterator")
	})
}

func Of[T any](list ...T) fp.Iterator[T] {
	return fp.IteratorOfSeq(list)
}

func FromSeq[T any](seq []T) fp.Iterator[T] {
	return fp.IteratorOfSeq(seq)
}

func FromPtr[T any](ptr *T) fp.Iterator[T] {
	if ptr == nil {
		return Empty[T]()
	} else {
		return Of(*ptr)
	}
}

func FromMap[K comparable, V any](m map[K]V) fp.Iterator[fp.Tuple2[K, V]] {
	return fp.IteratorOfGoMap(m)
}

func FromMapKey[K comparable, V any](m map[K]V) fp.Iterator[K] {
	return mutable.MapOf(m).Keys()
}

func FromMapValue[K comparable, V any](m map[K]V) fp.Iterator[V] {
	return mutable.MapOf(m).Values()
}

func Ap[T, U any](t fp.Iterator[fp.Func1[T, U]], a fp.Iterator[T]) fp.Iterator[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Iterator[U] {
		return Map(a, f)
	})
}

func Lift[T, U any](f func(v T) U) func(fp.Iterator[T]) fp.Iterator[U] {
	return func(opt fp.Iterator[T]) fp.Iterator[U] {
		return Map(opt, f)
	}
}

func Compose[A, B, C any](f1 func(A) fp.Iterator[B], f2 func(B) fp.Iterator[C]) func(A) fp.Iterator[C] {
	return func(a A) fp.Iterator[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Iterator[B] {
	return func(a A) fp.Iterator[B] {
		return Of(fab(a))
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

func FilterMap[T, U any](opt fp.Iterator[T], fn func(v T) fp.Option[U]) fp.Iterator[U] {
	return FlatMap(opt, fp.Compose(fn, fp.IteratorOfOption[U]))
}

func FlatMap[T, U any](opt fp.Iterator[T], fn func(v T) fp.Iterator[U]) fp.Iterator[U] {

	current := fp.None[fp.Iterator[U]]()

	hasNext := func() bool {
		if current.IsDefined() && current.Get().HasNext() {
			return true
		}

		for opt.HasNext() {
			nextItr := fn(opt.Next())
			current = fp.Some(nextItr)
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

func Flap[A, R any](tfa fp.Iterator[fp.Func1[A, R]]) func(A) fp.Iterator[R] {
	return func(a A) fp.Iterator[R] {
		return Ap(tfa, Of(a))
	}
}

func Flap2[A, B, R any](tfab fp.Iterator[fp.Func1[A, fp.Func1[B, R]]]) fp.Func1[A, fp.Func1[B, fp.Iterator[R]]] {
	return func(a A) fp.Func1[B, fp.Iterator[R]] {
		return Flap(Ap(tfab, Of(a)))
	}
}

func FlapMap[A, B, R any](tfab func(A, B) R, a fp.Iterator[A]) func(B) fp.Iterator[R] {
	return Flap(Map(a, as.Curried2(tfab)))
}

func Method1[A, B, R any](ta fp.Iterator[A], fab func(a A, b B) R) func(B) fp.Iterator[R] {
	return FlapMap(fab, ta)
}

func Method2[A, B, C, R any](ta fp.Iterator[A], fabc func(a A, b B, c C) R) func(B, C) fp.Iterator[R] {
	return curried.Revert2(Flap2(Map(ta, as.Curried3(fabc))))
}

func ToMap[K, V any](itr fp.Iterator[fp.Tuple2[K, V]], hasher fp.Hashable[K]) fp.Map[K, V] {
	ret := immutable.MapBuilder[K, V](hasher)

	for itr.HasNext() {
		k, v := itr.Next().Unapply()
		ret = ret.Add(k, v)
	}

	return ret.Build()
}

func ToGoMap[K comparable, V any](itr fp.Iterator[fp.Tuple2[K, V]]) map[K]V {
	ret := map[K]V{}
	for itr.HasNext() {
		k, v := itr.Next().Unapply()
		ret[k] = v
	}
	return ret
}

func ToSet[V any](itr fp.Iterator[V], hasher fp.Hashable[V]) fp.Set[V] {
	ret := immutable.SetBuilder(hasher)

	for itr.HasNext() {
		v := itr.Next()
		ret = ret.Add(v)
	}

	return ret.Build()
}

func ToGoSet[V comparable](itr fp.Iterator[V]) mutable.Set[V] {
	ret := map[V]bool{}
	for itr.HasNext() {
		k := itr.Next()
		ret[k] = true
	}
	return ret
}

func ToList[V any](itr fp.Iterator[V]) fp.List[V] {
	head := itr.NextOption()

	return fp.MakeList(func() fp.Option[V] {
		return head
	}, func() fp.List[V] {

		return ToList(itr)
	})
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

func Reduce[T any](r fp.Iterator[T], m fp.Monoid[T]) T {
	ret := m.Empty()
	for r.HasNext() {
		v := r.Next()
		m.Combine(ret, v)
	}
	return ret
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

func GroupBy[A any, K comparable](s fp.Iterator[A], keyFunc func(A) K) map[K]fp.Seq[A] {

	ret := map[K]fp.Seq[A]{}

	return Fold(s, ret, func(b map[K]fp.Seq[A], a A) map[K]fp.Seq[A] {
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

func Duplicate[T any](r fp.Iterator[T]) (fp.Iterator[T], fp.Iterator[T]) {
	lock := sync.Mutex{}

	queue := []T{}

	leftAhead := true

	unseq := func(r []T) (fp.Option[T], []T) {
		if len(r) > 0 {
			return fp.Some(r[0]), r[1:]
		} else {
			return fp.None[T](), nil
		}
	}

	left := fp.MakeIterator(
		func() bool {
			lock.Lock()
			defer lock.Unlock()

			if leftAhead || len(queue) == 0 {
				return r.HasNext()
			}
			// queue not empty
			return true
		},
		func() T {
			lock.Lock()
			defer lock.Unlock()

			if len(queue) == 0 {
				leftAhead = true
			}
			if leftAhead {
				ret := r.Next()
				queue = append(queue, ret)
				return ret
			}

			// leftAhead == false means queue not empty
			head, tail := unseq(queue)
			queue = tail
			return head.Get()
		},
	)

	right := fp.MakeIterator(
		func() bool {
			lock.Lock()
			defer lock.Unlock()

			if !leftAhead || len(queue) == 0 {
				return r.HasNext()
			}
			// queue not empty
			return true
		},
		func() T {
			lock.Lock()
			defer lock.Unlock()

			if len(queue) == 0 {
				// right ahead
				leftAhead = false
			}
			if !leftAhead {
				ret := r.Next()
				queue = append(queue, ret)
				return ret
			}
			// rightAhead means queue not empty
			head, tail := unseq(queue)
			queue = tail
			return head.Get()
		},
	)

	return left, right
}

func Span[T any](r fp.Iterator[T], p func(T) bool) (fp.Iterator[T], fp.Iterator[T]) {
	left, right := Duplicate(r)

	return left.TakeWhile(p), right.DropWhile(p)

}

func Partition[T any](r fp.Iterator[T], p func(T) bool) (fp.Iterator[T], fp.Iterator[T]) {
	left, right := Duplicate(r)

	return left.Filter(p), right.FilterNot(p)

}
