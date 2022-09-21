package fp

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

func (r Seq[T]) Tail() List[T] {
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

func (r Seq[T]) Unapply() (Option[T], List[T]) {
	if r.Size() > 0 {
		return r.Head(), r[1:]
	} else {
		return r.Head(), nil
	}
}

func (r Seq[T]) Take(n int) Seq[T] {
	return iteratorToSeq(r.Iterator().Take(n))
}

func (r Seq[T]) Drop(n int) Seq[T] {
	return iteratorToSeq(r.Iterator().Drop(n))
}

func (r Seq[T]) Foreach(f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) Filter(p func(v T) bool) Seq[T] {
	return iteratorToSeq(r.Iterator().Filter(p))
}

func (r Seq[T]) Exists(p func(v T) bool) bool {
	return r.Iterator().Exists(p)
}

func (r Seq[T]) ForAll(p func(v T) bool) bool {
	return r.Iterator().ForAll(p)
}

func (r Seq[T]) Find(p func(v T) bool) Option[T] {
	return r.Iterator().Find(p)
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

func (r Seq[T]) Reduce(m Monoid[T]) T {
	if r.Size() == 0 {
		return m.Empty()
	}

	reduce := m.Empty()
	for i := 0; i < len(r); i++ {
		reduce = m.Combine(reduce, r[i])
	}
	return reduce
}

func (r Seq[T]) Reverse() Seq[T] {
	ret := make(Seq[T], r.Size())

	for i := range r {
		ret[r.Size()-i-1] = r[i]
	}

	return ret
}

func (r Seq[T]) Iterator() Iterator[T] {
	idx := 0

	return MakeIterator(
		func() bool {
			return idx < r.Size()
		},
		func() T {
			ret := r[idx]
			idx++
			return ret
		},
	)
}
