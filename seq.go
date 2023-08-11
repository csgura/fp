package fp

import (
	"bytes"
	"fmt"
)

type Seq[T any] []T

func (r Seq[T]) Size() int {
	return len(r)
}

func (r Seq[T]) IsEmpty() bool {
	return r.Size() == 0
}

func (r Seq[T]) NonEmpty() bool {
	return r.Size() > 0
}

func (r Seq[T]) Get(idx int) Option[T] {
	if r.Size() > idx {
		return Some(r[idx])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Head() Option[T] {
	if r.Size() > 0 {
		return Some(r[0])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Init() Seq[T] {
	if r.Size() > 1 {
		return r[:r.Size()-1]
	} else {
		return nil
	}
}

func (r Seq[T]) Last() Option[T] {
	if r.Size() > 0 {
		return Some(r[r.Size()-1])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Tail() Seq[T] {
	if r.Size() > 0 {
		return r[1:]
	} else {
		return nil
	}
}

func (r Seq[T]) UnSeq() (Option[T], Seq[T]) {
	if r.Size() > 0 {
		return r.Head(), r[1:]
	} else {
		return r.Head(), nil
	}
}

func (r Seq[T]) Take(n int) Seq[T] {
	if len(r) < n {
		return r
	}
	return r[0:n]
}

func (r Seq[T]) Drop(n int) Seq[T] {
	if len(r) < n {
		return nil
	}
	return r[n:]
}

func (r Seq[T]) Foreach(f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) Filter(p func(v T) bool) Seq[T] {
	ret := make([]T, 0, len(r))
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func (r Seq[T]) FilterNot(p func(v T) bool) Seq[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Seq[T]) Exists(p func(v T) bool) bool {
	for _, v := range r {
		if p(v) {
			return true
		}
	}
	return false
}

func (r Seq[T]) ForAll(p func(v T) bool) bool {
	for _, v := range r {
		if !p(v) {
			return false
		}
	}
	return true
}

func (r Seq[T]) Map(mf func(T) T) Seq[T] {
	var ret = make([]T, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v))
	}
	return ret
}

func (r Seq[T]) FlatMap(mf func(T) Seq[T]) Seq[T] {
	var ret = make([]T, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v)...)
	}
	return ret
}

func (r Seq[T]) Find(p func(v T) bool) Option[T] {
	for _, v := range r {
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Seq[T]) Add(item T) Seq[T] {
	return r.Append(item)
}

func (r Seq[T]) Append(items ...T) Seq[T] {
	tail := Seq[T](items)
	ret := make(Seq[T], r.Size()+tail.Size())

	for i := range r {
		ret[i] = r[i]
	}

	for i := range tail {
		ret[i+r.Size()] = tail[i]
	}

	return ret
}

func (r Seq[T]) Concat(tail Seq[T]) Seq[T] {
	ret := make(Seq[T], r.Size()+tail.Size())

	for i := range r {
		ret[i] = r[i]
	}

	for i := range tail {
		ret[i+r.Size()] = tail[i]
	}

	return ret
}

// func (r Seq[T]) Reduce(m Monoid[T]) T {
// 	if r.Size() == 0 {
// 		return m.Empty()
// 	}

// 	reduce := m.Empty()
// 	for i := 0; i < len(r); i++ {
// 		reduce = m.Combine(reduce, r[i])
// 	}
// 	return reduce
// }

func (r Seq[T]) Reverse() Seq[T] {
	ret := make(Seq[T], r.Size())

	for i := range r {
		ret[r.Size()-i-1] = r[i]
	}

	return ret
}

// func (r Seq[T]) Iterator() Iterator[T] {
// 	return iteratorFromSlice(r)
// }

func IteratorOfSeq[T any](r []T) Iterator[T] {
	idx := 0

	return MakeIterator(
		func() bool {
			return idx < len(r)
		},
		func() T {
			ret := r[idx]
			idx++
			return ret
		},
	)
}

func IteratorOfOption[T any](r Option[T]) Iterator[T] {
	first := true

	return MakeIterator(
		func() bool {
			return first && r.IsDefined()
		},
		func() T {
			if first && r.IsDefined() {
				first = false
				return r.Get()
			}
			panic("next on empty iterator")
		},
	)
}

func (r Seq[T]) MakeString(sep string) string {
	buf := &bytes.Buffer{}

	for i, v := range r {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(fmt.Sprint(v))
	}
	return buf.String()
}
