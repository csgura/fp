package fp

import "sync"

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
	NextOption() Option[T]
	ToSeq() Seq[T]
	Take(n int) Iterator[T]
	TakeWhile(p func(T) bool) Iterator[T]
	Drop(n int) Iterator[T]
	DropWhile(p func(T) bool) Iterator[T]
	Filter(p func(T) bool) Iterator[T]
	FilterNot(p func(T) bool) Iterator[T]
	Find(p func(T) bool) Option[T]
	Foreach(p func(T))
	Concat(tail Iterator[T]) Iterator[T]
	Reduce(m Monoid[T]) T
	Exists(p func(v T) bool) bool
	ForAll(p func(v T) bool) bool
	IsEmpty() bool
	NonEmpty() bool
	TapEach(p func(T)) Iterator[T]
	Span(p func(T) bool) Tuple2[Iterator[T], Iterator[T]]
	Partition(p func(T) bool) Tuple2[Iterator[T], Iterator[T]]
	Duplicate() Tuple2[Iterator[T], Iterator[T]]
}

var _ Iterator[int] = IteratorAdaptor[int]{}

type IteratorAdaptor[T any] struct {
	IsHasNext func() bool
	GetNext   func() T
}

func (r IteratorAdaptor[T]) HasNext() bool {
	return r.IsHasNext()
}

func (r IteratorAdaptor[T]) Next() T {
	return r.GetNext()
}

func (r IteratorAdaptor[T]) NextOption() Option[T] {
	if r.HasNext() {
		return Some[T]{r.GetNext()}
	}
	return None[T]{}
}

func (r IteratorAdaptor[T]) ToSeq() Seq[T] {
	ret := Seq[T]{}
	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func (r IteratorAdaptor[T]) Take(n int) Iterator[T] {

	i := 0
	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			if i < n {
				return r.HasNext()
			}
			return false
		},
		GetNext: func() T {
			i++
			return r.Next()
		},
	}
}

func (r IteratorAdaptor[T]) nextOnEmpty() T {
	panic("next on empty iterator")
}

func (r IteratorAdaptor[T]) TakeWhile(p func(T) bool) Iterator[T] {

	breaking := false
	var fv Option[T] = None[T]{}

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
				fv = Some[T]{v}
				return true
			}
			breaking = true
		}
		return false
	}

	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			return hasNext()
		},
		GetNext: func() T {

			if hasNext() {
				ret := fv.Get()
				fv = None[T]{}
				return ret
			}
			return r.nextOnEmpty()
		},
	}
}

func (r IteratorAdaptor[T]) Drop(n int) Iterator[T] {

	for i := 0; i < n && r.HasNext(); i++ {
		r.Next()
	}

	return r
}

func (r IteratorAdaptor[T]) DropWhile(p func(T) bool) Iterator[T] {

	found := false
	var first Option[T] = None[T]{}
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
				first = Some[T]{v}
				found = true
				return true
			}
		}
		return false
	}

	return IteratorAdaptor[T]{
		IsHasNext: hasNext,
		GetNext: func() T {
			if hasNext() {
				if first.IsDefined() {
					ret := first.Get()
					first = None[T]{}
					return ret
				}

				if found {
					return r.Next()
				}

			}
			return r.nextOnEmpty()
		},
	}
}

func (r IteratorAdaptor[T]) Filter(p func(T) bool) Iterator[T] {

	first := true
	var fv Option[T] = None[T]{}

	hasNext := func() bool {
		if first {
			fv = r.Find(p)
			first = false
		}
		return fv.IsDefined()
	}
	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			return hasNext()
		},
		GetNext: func() T {
			if hasNext() {

				ret := fv.Get()
				fv = r.Find(p)
				return ret
			}
			return r.nextOnEmpty()
		},
	}
}

func (r IteratorAdaptor[T]) FilterNot(p func(T) bool) Iterator[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r IteratorAdaptor[T]) Find(p func(T) bool) Option[T] {
	for r.HasNext() {
		v := r.Next()
		if p(v) {
			return Some[T]{v}
		}
	}
	return None[T]{}
}

func (r IteratorAdaptor[T]) Foreach(p func(T)) {
	for r.HasNext() {
		v := r.Next()
		p(v)
	}
}

func (r IteratorAdaptor[T]) TapEach(p func(T)) Iterator[T] {
	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			return r.IsHasNext()
		},
		GetNext: func() T {
			ret := r.GetNext()
			p(ret)
			return ret
		},
	}
}

func (r IteratorAdaptor[T]) Concat(tail Iterator[T]) Iterator[T] {

	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			return r.HasNext() || tail.HasNext()
		},
		GetNext: func() T {
			if r.HasNext() {
				return r.Next()
			}
			return tail.Next()
		},
	}
}

func (r IteratorAdaptor[T]) Reduce(m Monoid[T]) T {
	ret := m.Empty()
	for r.HasNext() {
		v := r.Next()
		m.Combine(ret, v)
	}
	return ret
}

func (r IteratorAdaptor[T]) Exists(p func(v T) bool) bool {
	for r.HasNext() {
		if p(r.Next()) {
			return true
		}
	}

	return false
}

func (r IteratorAdaptor[T]) ForAll(p func(v T) bool) bool {
	for r.HasNext() {
		if !p(r.Next()) {
			return false
		}
	}
	return true
}

func (r IteratorAdaptor[T]) IsEmpty() bool {
	return !r.HasNext()
}

func (r IteratorAdaptor[T]) NonEmpty() bool {
	return r.HasNext()
}

func (r IteratorAdaptor[T]) Duplicate() Tuple2[Iterator[T], Iterator[T]] {
	lock := sync.Mutex{}

	queue := Seq[T]{}

	leftAhead := true

	left := IteratorAdaptor[T]{
		IsHasNext: func() bool {
			lock.Lock()
			defer lock.Unlock()

			if leftAhead || queue.IsEmpty() {
				return r.IsHasNext()
			}
			// queue not empty
			return true
		},
		GetNext: func() T {
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
			ret := queue.Head()
			queue = queue.Tail()
			return ret.Get()
		},
	}

	right := IteratorAdaptor[T]{
		IsHasNext: func() bool {
			lock.Lock()
			defer lock.Unlock()

			if !leftAhead || queue.IsEmpty() {
				return r.IsHasNext()
			}
			// queue not empty
			return true
		},
		GetNext: func() T {
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
			ret := queue.Head()
			queue = queue.Tail()
			return ret.Get()
		},
	}

	return Tuple2[Iterator[T], Iterator[T]]{left, right}
}

func (r IteratorAdaptor[T]) Span(p func(T) bool) Tuple2[Iterator[T], Iterator[T]] {
	left, right := r.Duplicate().Unapply()

	return Tuple2[Iterator[T], Iterator[T]]{left.TakeWhile(p), right.DropWhile(p)}

}

func (r IteratorAdaptor[T]) Partition(p func(T) bool) Tuple2[Iterator[T], Iterator[T]] {
	left, right := r.Duplicate().Unapply()

	return Tuple2[Iterator[T], Iterator[T]]{left.Filter(p), right.FilterNot(p)}

}
