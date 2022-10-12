package fp

import "fmt"

type SetMinimal[V any] interface {
	Contains(v V) bool
	Size() int
	Iterator() Iterator[V]
	Incl(v V) SetMinimal[V]
	Excl(v V) SetMinimal[V]
}

type Set[V any] struct {
	getEmpty func() SetMinimal[V]
	set      SetMinimal[V]
}

func (r Set[V]) Contains(v V) bool {
	if r.set == nil {
		return false
	}
	return r.set.Contains(v)
}
func (r Set[V]) Size() int {
	if r.set == nil {
		return 0
	}
	return r.set.Size()
}
func (r Set[V]) Iterator() Iterator[V] {

	if r.set == nil {
		return MakeIterator(func() bool {
			return false
		}, func() V {
			panic("next on empty iterator")
		})
	}

	return r.set.Iterator()
}
func (r Set[V]) Incl(v V) Set[V] {
	if r.set == nil && r.getEmpty == nil {
		return MakeSet[V](func() SetMinimal[V] {
			return UnsageGoSet[V]{}
		}, UnsageGoSet[V]{
			v: true,
		})
	}

	if r.set == nil {
		return MakeSet(r.getEmpty, r.getEmpty().Incl(v))
	}

	return MakeSet(r.getEmpty, r.set.Incl(v))
}

func (r Set[V]) Excl(v V) Set[V] {
	if r.set == nil {
		return r
	}
	return MakeSet(r.getEmpty, r.set.Excl(v))
}

func (r Set[V]) Foreach(f func(V)) {
	r.Iterator().Foreach(f)
}
func (r Set[V]) Concat(other Iterable[V]) Set[V] {
	var ret Set[V] = r
	itr := other.Iterator()
	for itr.HasNext() {
		ret = ret.Incl(itr.Next())
	}
	return ret
}

func (r Set[V]) SubsetOf(other Set[V]) bool {
	return r.Iterator().ForAll(other.Contains)
}

func (r Set[V]) Diff(other Set[V]) Set[V] {
	ret := r.getEmpty()

	itr := r.Iterator()
	for itr.HasNext() {
		e := itr.Next()
		if other.Contains(e) == false {
			ret = ret.Incl(e)
		}
	}

	return MakeSet(r.getEmpty, ret)
}
func (r Set[V]) Intersect(other Set[V]) Set[V] {
	ret := r.getEmpty()

	itr := r.Iterator()
	for itr.HasNext() {
		e := itr.Next()
		if other.Contains(e) {
			ret = ret.Incl(e)
		}
	}

	return MakeSet(r.getEmpty, ret)
}

func (r Set[V]) String() string {
	if r.set == nil {
		return "Set()"
	}
	return fmt.Sprint(r.set)
}

func (r Set[V]) IsEmpty() bool {
	return r.Size() == 0
}

func (r Set[V]) NonEmpty() bool {
	return r.Size() != 0
}

func MakeSet[V any](empty func() SetMinimal[V], s SetMinimal[V]) Set[V] {
	return Set[V]{empty, s}
}

type UnsageGoSet[V any] map[any]bool

var _ SetMinimal[string] = UnsageGoSet[string]{}

func (r UnsageGoSet[V]) Contains(v V) bool {
	return r[v]
}

func (r UnsageGoSet[V]) Size() int {
	return len(r)
}

func (r UnsageGoSet[V]) Iterator() Iterator[V] {
	seq := Seq[V]{}
	for k := range r {
		seq = append(seq, k.(V))
	}
	return seq.Iterator()
}

func (r UnsageGoSet[V]) Incl(v V) SetMinimal[V] {
	n := UnsageGoSet[V]{}
	for ek, ev := range r {
		n[ek] = ev
	}
	n[v] = true
	return n
}

func (r UnsageGoSet[V]) Excl(v V) SetMinimal[V] {
	n := UnsageGoSet[V]{}
	for ek, ev := range r {
		n[ek] = ev
	}
	delete(n, v)
	return n
}
