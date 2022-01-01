package fp

import "fmt"

type SetMinimal[V any] interface {
	Contains(v V) bool
	Size() int
	Iterator() Iterator[V]
	Incl(v V) SetMinimal[V]
	Excl(v V) SetMinimal[V]
}

type Set[V any] interface {
	IsEmpty() bool
	NonEmpty() bool
	Contains(v V) bool
	Size() int
	Iterator() Iterator[V]
	Incl(v V) Set[V]
	Excl(v V) Set[V]
	Foreach(func(V))
	Concat(Iterable[V]) Set[V]
	SubsetOf(other Set[V]) bool
	Diff(other Set[V]) Set[V]
	Intersect(other Set[V]) Set[V]
}

type SetOps[V any] struct {
	GetEmpty func() SetMinimal[V]
	Set      SetMinimal[V]
}

func (r SetOps[V]) Contains(v V) bool {
	return r.Set.Contains(v)
}
func (r SetOps[V]) Size() int {
	return r.Set.Size()
}
func (r SetOps[V]) Iterator() Iterator[V] {
	return r.Set.Iterator()
}
func (r SetOps[V]) Incl(v V) Set[V] {
	return MakeSet(r.GetEmpty, r.Set.Incl(v))
}
func (r SetOps[V]) Excl(v V) Set[V] {
	return MakeSet(r.GetEmpty, r.Set.Excl(v))
}
func (r SetOps[V]) Foreach(f func(V)) {
	r.Iterator().Foreach(f)
}
func (r SetOps[V]) Concat(other Iterable[V]) Set[V] {
	var ret Set[V] = r
	itr := other.Iterator()
	for itr.HasNext() {
		ret = ret.Incl(itr.Next())
	}
	return ret
}

func (r SetOps[V]) SubsetOf(other Set[V]) bool {
	return r.Iterator().ForAll(other.Contains)
}

func (r SetOps[V]) Diff(other Set[V]) Set[V] {
	ret := r.GetEmpty()

	itr := r.Iterator()
	for itr.HasNext() {
		e := itr.Next()
		if other.Contains(e) == false {
			ret = ret.Incl(e)
		}
	}

	return MakeSet(r.GetEmpty, ret)
}
func (r SetOps[V]) Intersect(other Set[V]) Set[V] {
	ret := r.GetEmpty()

	itr := r.Iterator()
	for itr.HasNext() {
		e := itr.Next()
		if other.Contains(e) {
			ret = ret.Incl(e)
		}
	}

	return MakeSet(r.GetEmpty, ret)
}

func (r SetOps[V]) String() string {
	return fmt.Sprint(r.Set)
}

func (r SetOps[V]) IsEmpty() bool {
	return r.Set.Size() == 0
}

func (r SetOps[V]) NonEmpty() bool {
	return r.Set.Size() != 0
}

func MakeSet[V any](empty func() SetMinimal[V], s SetMinimal[V]) Set[V] {
	return SetOps[V]{empty, s}
}
