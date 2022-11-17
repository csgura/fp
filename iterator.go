package fp

import (
	"bytes"
	"fmt"
	"sync"
)

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] struct {
	hasNext func() bool
	next    func() T
}

func (r Iterator[T]) Head() Option[T] {
	if r.HasNext() {
		return Some(r.Next())
	}
	return None[T]()
}

func (r Iterator[T]) First() Option[T] {
	if r.HasNext() {
		return Some(r.Next())
	}
	return None[T]()
}

func (r Iterator[T]) ToSeq() []T {
	ret := []T{}
	for r.HasNext() {
		ret = append(ret, r.Next())
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

func (r Iterator[T]) MakeString(sep string) string {
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

func (r Iterator[T]) HasNext() bool {
	if r.hasNext == nil {
		return false
	}

	return r.hasNext()
}

func (r Iterator[T]) Next() T {
	return r.next()
}

func (r Iterator[T]) NextOption() Option[T] {
	if r.HasNext() {
		v := r.next()
		return Some(v)
	}
	return None[T]()
}

func iteratorToSeq[T any](r Iterator[T], capa int) []T {
	ret := make([]T, 0, capa)

	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func (r Iterator[T]) Take(n int) Iterator[T] {

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

func (r Iterator[T]) nextOnEmpty() T {
	panic("next on empty iterator")
}

func (r Iterator[T]) TakeWhile(p func(T) bool) Iterator[T] {

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

func (r Iterator[T]) Drop(n int) Iterator[T] {

	for i := 0; i < n && r.HasNext(); i++ {
		r.Next()
	}

	return r
}

func (r Iterator[T]) DropWhile(p func(T) bool) Iterator[T] {

	found := false
	var first Option[T] = None[T]()
	hasNext := func() bool {
		if found {
			return r.HasNext()
		}
		if first.IsDefined() {
			return true
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

func (r Iterator[T]) Filter(p func(T) bool) Iterator[T] {

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

func (r Iterator[T]) FilterNot(p func(T) bool) Iterator[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Iterator[T]) Find(p func(T) bool) Option[T] {
	for r.HasNext() {
		v := r.Next()
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Iterator[T]) Foreach(p func(T)) {
	for r.HasNext() {
		v := r.Next()
		p(v)
	}
}

func (r Iterator[T]) TapEach(p func(T)) Iterator[T] {
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

func (r Iterator[T]) Appended(elem T) Iterator[T] {
	return r.Concat(IteratorOfSeq([]T{elem}))
}

func (r Iterator[T]) Concat(tail Iterator[T]) Iterator[T] {

	return MakeIterator(
		func() bool {
			return r.HasNext() || tail.HasNext()
		},
		func() T {
			if r.HasNext() {
				return r.Next()
			}
			return tail.Next()
		},
	)
}

func (r Iterator[T]) Reduce(m Monoid[T]) T {
	ret := m.Empty()
	for r.HasNext() {
		v := r.Next()
		m.Combine(ret, v)
	}
	return ret
}

func (r Iterator[T]) Exists(p func(v T) bool) bool {
	for r.HasNext() {
		if p(r.Next()) {
			return true
		}
	}

	return false
}

func (r Iterator[T]) ForAll(p func(v T) bool) bool {
	for r.HasNext() {
		if !p(r.Next()) {
			return false
		}
	}
	return true
}

func (r Iterator[T]) IsEmpty() bool {
	return !r.HasNext()
}

func (r Iterator[T]) NonEmpty() bool {
	return r.HasNext()
}

func (r Iterator[T]) Duplicate() (Iterator[T], Iterator[T]) {
	lock := sync.Mutex{}

	queue := Seq[T]{}

	leftAhead := true

	left := MakeIterator(
		func() bool {
			lock.Lock()
			defer lock.Unlock()

			if leftAhead || queue.IsEmpty() {
				return r.HasNext()
			}
			// queue not empty
			return true
		},
		func() T {
			lock.Lock()
			defer lock.Unlock()

			if queue.IsEmpty() {
				leftAhead = true
			}
			if leftAhead {
				ret := r.Next()
				queue = queue.Append(ret)
				return ret
			}

			// leftAhead == false means queue not empty
			head, tail := queue.UnSeq()
			queue = tail
			return head.Get()
		},
	)

	right := MakeIterator(
		func() bool {
			lock.Lock()
			defer lock.Unlock()

			if !leftAhead || queue.IsEmpty() {
				return r.HasNext()
			}
			// queue not empty
			return true
		},
		func() T {
			lock.Lock()
			defer lock.Unlock()

			if queue.IsEmpty() {
				// right ahead
				leftAhead = false
			}
			if !leftAhead {
				ret := r.Next()
				queue = queue.Append(ret)
				return ret
			}
			// rightAhead means queue not empty
			head, tail := queue.UnSeq()
			queue = tail
			return head.Get()
		},
	)

	return left, right
}

func (r Iterator[T]) Span(p func(T) bool) (Iterator[T], Iterator[T]) {
	left, right := r.Duplicate()

	return left.TakeWhile(p), right.DropWhile(p)

}

func (r Iterator[T]) Partition(p func(T) bool) (Iterator[T], Iterator[T]) {
	left, right := r.Duplicate()

	return left.Filter(p), right.FilterNot(p)

}

func MakeIterator[T any](has func() bool, next func() T) Iterator[T] {
	return Iterator[T]{
		hasNext: has,
		next:    next,
	}
}
