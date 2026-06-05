package fp

import (
	"iter"
	"net/http"
	"runtime"
)

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] struct {
	hasNext func() bool
	next    func() T
	concat  []Iterator[T]
}

var ErrIteratorEmpty = Error(http.StatusNotFound, "next on empty iterator")

// range over func
//
//	func (r Iterator[T]) All() func(yield func(T) bool) bool {
//		return func(yield func(T) bool) bool {
//			for r.HasNext() {
//				v := r.Next()
//				if !yield(v) {
//					return false
//				}
//			}
//			return true
//		}
//	}

func MakeIterator[T any](has func() bool, next func() T) Iterator[T] {
	return Iterator[T]{
		hasNext: has,
		next:    next,
	}
}

type pull[T any] struct {
	nextfn func() (T, bool)
	val    T
	ok     bool
}

func (r *pull[T]) hasNext() bool {
	return r.ok
}

func (r *pull[T]) next() T {
	if !r.ok {
		panic(ErrIteratorEmpty)
	}

	ret := r.val

	r.val, r.ok = r.nextfn()

	return ret
}

func MakePullIterator[T any](seq iter.Seq[T]) Iterator[T] {
	nextfn, stopfn := iter.Pull(seq)

	nv, ok := nextfn()

	ret := &pull[T]{
		nextfn: nextfn,
		val:    nv,
		ok:     ok,
	}

	runtime.AddCleanup(ret, func(s func()) {
		s()
	}, stopfn)

	//runtime.SetFinalizer(ret, (*pull[T]).Close)

	return MakeIterator(ret.hasNext, ret.next)

}
