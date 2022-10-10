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
	return r.set.Contains(v)
}
func (r Set[V]) Size() int {
	return r.set.Size()
}
func (r Set[V]) Iterator() Iterator[V] {
	return r.set.Iterator()
}
func (r Set[V]) Incl(v V) Set[V] {
	return MakeSet(r.getEmpty, r.set.Incl(v))
}
func (r Set[V]) Excl(v V) Set[V] {
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
	return fmt.Sprint(r.set)
}

func (r Set[V]) IsEmpty() bool {
	return r.set.Size() == 0
}

func (r Set[V]) NonEmpty() bool {
	return r.set.Size() != 0
}

func MakeSet[V any](empty func() SetMinimal[V], s SetMinimal[V]) Set[V] {
	return Set[V]{empty, s}
}
