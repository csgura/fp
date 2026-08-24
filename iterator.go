package fp

import (
	"bytes"
	"fmt"
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

func (r Iterator[T]) All[_ Phantom[T]]() func(func(T) bool) {
	return func(f func(T) bool) {
		for r.hasNext() {
			if !f(r.Next()) {
				return
			}
		}
	}
}

func (r Iterator[T]) ToSeq[_ Phantom[T]]() []T {
	ret := []T{}
	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func (r Iterator[T]) Count[_ Phantom[T]]() int {
	ret := 0
	for r.HasNext() {
		r.Next()
		ret++
	}
	return ret
}

// func (r Iterator[T]) ToList() List[T] {

// 	head := r.NextOption()

// 	return MakeList(func() Option[T] {
// 		return head
// 	}, func() List[T] {
// 		return r.ToList()
// 	})
// }

func (r Iterator[T]) MakeString[_ Phantom[T]](sep string) string {
	buf := &bytes.Buffer{}

	first := true
	for r.HasNext() {
		if !first {
			buf.WriteString(sep)
		} else {
			first = false
		}

		v := r.Next()
		buf.WriteString(fmt.Sprint(v))

	}

	return buf.String()
}

func (r Iterator[T]) HasNext[_ Phantom[T]]() bool {
	if r.hasNext == nil {
		return false
	}

	return r.hasNext()
}

func (r Iterator[T]) Next[_ Phantom[T]]() T {
	return r.next()
}

func (r Iterator[T]) NextOption[_ Phantom[T]]() Option[T] {
	if r.HasNext() {
		v := r.next()
		return Some(v)
	}
	return None[T]()
}

func (r Iterator[T]) Take[_ Phantom[T]](n int) Iterator[T] {

	i := 0
	hasNext := func() bool {
		if i < n {
			return r.HasNext()
		}
		return false
	}
	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {
				i++
				return r.Next()
			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) nextOnEmpty[_ Phantom[T]]() T {
	panic(ErrIteratorEmpty)
}

func (r Iterator[T]) TakeWhile[_ Phantom[T]](p func(T) bool) Iterator[T] {

	breaking := false
	var fv Option[T] = None[T]()

	hasNext := func() bool {
		if breaking {
			return false
		}

		if fv.IsDefined() {
			return true
		}

		if r.HasNext() {
			v := r.Next()
			if p(v) {
				fv = Some(v)
				return true
			}
			breaking = true
		}
		return false
	}

	return MakeIterator(
		hasNext,
		func() T {

			if hasNext() {
				ret := fv.Get()
				fv = None[T]()
				return ret
			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) Drop[_ Phantom[T]](n int) Iterator[T] {

	for i := 0; i < n && r.HasNext(); i++ {
		r.Next()
	}

	return r
}

func (r Iterator[T]) DropWhile[_ Phantom[T]](p func(T) bool) Iterator[T] {

	found := false
	var first Option[T] = None[T]()
	hasNext := func() bool {
		if first.IsDefined() {
			return true
		}

		if found {
			return r.HasNext()
		}

		for r.HasNext() {
			v := r.Next()
			if !p(v) {
				first = Some(v)
				found = true
				return true
			}
		}
		return false
	}

	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {
				if first.IsDefined() {
					ret := first.Get()
					first = None[T]()
					return ret
				}

				if found {
					return r.Next()
				}

			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) Filter[_ Phantom[T]](p func(T) bool) Iterator[T] {

	first := true
	var fv Option[T] = None[T]()

	hasNext := func() bool {
		if first {
			fv = r.Find(p)
			first = false
		}
		return fv.IsDefined()
	}

	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {

				ret := fv.Get()
				fv = r.Find(p)
				return ret
			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) FilterMap[S any](fn func(T) Option[S]) Iterator[S] {
	return r.FlatMap(Compose(fn, IteratorOfOption))
}

func (r Iterator[T]) FilterNot[_ Phantom[T]](p func(T) bool) Iterator[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Iterator[T]) Find[_ Phantom[T]](p func(T) bool) Option[T] {
	for r.HasNext() {
		v := r.Next()
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Iterator[T]) Foreach[_ Phantom[T]](p func(T)) {
	for r.HasNext() {
		v := r.Next()
		p(v)
	}
}

func (r Iterator[T]) TapEach[_ Phantom[T]](p func(T)) Iterator[T] {
	return MakeIterator(
		func() bool {
			return r.HasNext()
		},
		func() T {
			ret := r.next()
			p(ret)
			return ret
		},
	)
}

func (r Iterator[T]) Appended[_ Phantom[T]](elem T) Iterator[T] {
	return r.Concat(IteratorOfSeq([]T{elem}))
}

func (r Iterator[T]) Concat[_ Phantom[T]](tail Iterator[T]) Iterator[T] {

	alliter := r.concat
	if len(alliter) == 0 {
		alliter = make([]Iterator[T], 1, 2+len(tail.concat))
		alliter[0] = r
	}

	if len(tail.concat) > 0 {
		alliter = append(alliter, tail.concat...)
	} else {
		alliter = append(alliter, tail)
	}

	currentItr := Some(alliter[0])
	remainItr := alliter[1:]

	currentNextChecked := false

	currentNext := func() bool {

		if currentNextChecked {
			return true
		}

		if currentItr.IsEmpty() {
			return false
		}

		if currentItr.Get().HasNext() {
			currentNextChecked = true
			return true
		}

		for i, itr := range remainItr {
			if itr.HasNext() {
				currentItr = Some(itr)
				remainItr = remainItr[i+1:]
				currentNextChecked = true
				return true
			}
		}

		currentItr = None[Iterator[T]]()

		return false
	}

	ret := MakeIterator(
		func() bool {

			return currentNext()
		},
		func() T {
			if currentNext() {
				currentNextChecked = false
				return currentItr.Get().Next()
			}
			panic(ErrIteratorEmpty)
		},
	)
	ret.concat = alliter

	return ret
}

func (r Iterator[T]) Map[R any](mf func(T) R) Iterator[R] {
	return MakeIterator(
		func() bool {
			return r.HasNext()
		},
		func() R {
			return mf(r.Next())
		},
	)
}

func (r Iterator[T]) FlatMap[R any](mf func(T) Iterator[R]) Iterator[R] {
	current := None[Iterator[R]]()

	hasNext := func() bool {
		if current.IsDefined() && current.Get().HasNext() {
			return true
		}

		for r.HasNext() {
			nextItr := mf(r.Next())
			current = Some(nextItr)
			if nextItr.HasNext() {
				return true
			}
		}

		return false
	}

	return MakeIterator(
		hasNext,
		func() R {
			if hasNext() {
				return current.Get().Next()
			}
			panic(ErrIteratorEmpty)
		},
	)
}

// func (r Iterator[T]) Reduce(m Monoid[T]) T {
// 	ret := m.Empty()
// 	for r.HasNext() {
// 		v := r.Next()
// 		m.Combine(ret, v)
// 	}
// 	return ret
// }

func (r Iterator[T]) Exists[_ Phantom[T]](p func(v T) bool) bool {
	for r.HasNext() {
		if p(r.Next()) {
			return true
		}
	}

	return false
}

func (r Iterator[T]) ForAll[_ Phantom[T]](p func(v T) bool) bool {
	for r.HasNext() {
		if !p(r.Next()) {
			return false
		}
	}
	return true
}

func (r Iterator[T]) IsEmpty[_ Phantom[T]]() bool {
	return !r.HasNext()
}

func (r Iterator[T]) NonEmpty[_ Phantom[T]]() bool {
	return r.HasNext()
}
