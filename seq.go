package fp

import "sort"

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

func (r Seq[T]) Head() Option[T] {
	if r.Size() > 0 {
		return Some[T]{r[0]}
	} else {
		return None[T]{}
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
	return r.Iterator().Take(n).ToSeq()
}

func (r Seq[T]) Drop(n int) Seq[T] {
	return r.Iterator().Drop(n).ToSeq()
}

func (r Seq[T]) Foreach(f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) Filter(p func(v T) bool) Seq[T] {
	return r.Iterator().Filter(p).ToSeq()
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

	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			return idx < r.Size()
		},
		GetNext: func() T {
			ret := r[idx]
			idx++
			return ret
		},
	}
}

type seqSorter[T any] struct {
	seq Seq[T]
	ord Ord[T]
}

func (p *seqSorter[T]) Len() int           { return len(p.seq) }
func (p *seqSorter[T]) Less(i, j int) bool { return p.ord.Less(p.seq[i], p.seq[j]) }
func (p *seqSorter[T]) Swap(i, j int)      { p.seq[i], p.seq[j] = p.seq[j], p.seq[i] }

func (r Seq[T]) Sort(ord Ord[T]) Seq[T] {
	ns := r.Concat(nil)
	sort.Sort(&seqSorter[T]{ns, ord})
	return ns
}
