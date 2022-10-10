package fp

import "fmt"

type MapMinimal[K, V any] interface {
	Size() int
	Get(k K) Option[V]
	Removed(k ...K) MapMinimal[K, V]
	Updated(k K, v V) MapMinimal[K, V]
	Iterator() Iterator[Tuple2[K, V]]
}

type MapMinimalUpdatedWith[K, V any] interface {
	UpdatedWith(k K, remap func(Option[V]) Option[V]) MapMinimal[K, V]
}

type Map[K, V any] struct {
	minimal MapMinimal[K, V]
}

func (r Map[K, V]) Size() int {
	return r.minimal.Size()
}

func (r Map[K, V]) IsEmpty() bool {
	return r.minimal.Size() == 0
}

func (r Map[K, V]) NonEmpty() bool {
	return r.minimal.Size() != 0
}

func (r Map[K, V]) Get(k K) Option[V] {
	return r.minimal.Get(k)
}
func (r Map[K, V]) Removed(k ...K) Map[K, V] {
	return MakeMap(r.minimal.Removed(k...))
}
func (r Map[K, V]) Updated(k K, v V) Map[K, V] {
	return MakeMap(r.minimal.Updated(k, v))
}

func (r Map[K, V]) UpdatedWith(k K, remap func(Option[V]) Option[V]) Map[K, V] {
	if um, ok := r.minimal.(MapMinimalUpdatedWith[K, V]); ok {
		return MakeMap(um.UpdatedWith(k, remap))
	}

	v := r.Get(k)
	nv := remap(v)

	if nv.IsDefined() {
		return r.Updated(k, nv.Get())
	} else if v.IsDefined() {
		return r.Removed(k)
	}
	return r
}

func (r Map[K, V]) Iterator() Iterator[Tuple2[K, V]] {
	return r.minimal.Iterator()
}

func (r Map[K, V]) Contains(k K) bool {
	return r.Get(k).IsDefined()
}

func (r Map[K, V]) Keys() Iterator[K] {
	itr := r.Iterator()
	return MakeIterator(itr.HasNext, func() K {
		return itr.Next().I1
	})
}

func (r Map[K, V]) Values() Iterator[V] {
	itr := r.Iterator()
	return MakeIterator(itr.HasNext, func() V {
		return itr.Next().I2
	})
}

func (r Map[K, V]) Foreach(f func(Tuple2[K, V])) {
	r.Iterator().Foreach(f)
}

func (r Map[K, V]) Concat(other Iterable[Tuple2[K, V]]) Map[K, V] {
	var ret Map[K, V] = r
	itr := other.Iterator()
	for itr.HasNext() {
		next := itr.Next()
		ret = ret.Updated(next.I1, next.I2)
	}

	return ret
}

func (r Map[K, V]) String() string {
	return fmt.Sprint(r.minimal)
}

func MakeMap[K, V any](minimal MapMinimal[K, V]) Map[K, V] {
	return Map[K, V]{minimal}
}
