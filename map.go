package fp

import "fmt"

type MapMinimal[K, V any] interface {
	Size() int
	Get(k K) Option[V]
	Removed(k ...K) MapMinimal[K, V]
	Updated(k K, v V) MapMinimal[K, V]
	Iterator() Iterator[Tuple2[K, V]]
}

type Map[K, V any] interface {
	IsEmpty() bool
	NonEmpty() bool
	Size() int
	Get(k K) Option[V]
	Removed(k ...K) Map[K, V]
	Updated(k K, v V) Map[K, V]
	UpdatedWith(k K, remap func(Option[V]) Option[V]) Map[K, V]
	Iterator() Iterator[Tuple2[K, V]]
	Contains(K) bool
	Keys() Iterator[K]
	Values() Iterator[V]
	Foreach(func(Tuple2[K, V]))
	Concat(Iterable[Tuple2[K, V]]) Map[K, V]
}

type MapOps[K, V any] struct {
	Map MapMinimal[K, V]
}

func (r MapOps[K, V]) Size() int {
	return r.Map.Size()
}

func (r MapOps[K, V]) IsEmpty() bool {
	return r.Map.Size() == 0
}

func (r MapOps[K, V]) NonEmpty() bool {
	return r.Map.Size() != 0
}

func (r MapOps[K, V]) Get(k K) Option[V] {
	return r.Map.Get(k)
}
func (r MapOps[K, V]) Removed(k ...K) Map[K, V] {
	return MakeMap(r.Map.Removed(k...))
}
func (r MapOps[K, V]) Updated(k K, v V) Map[K, V] {
	return MakeMap(r.Map.Updated(k, v))
}

func (r MapOps[K, V]) UpdatedWith(k K, remap func(Option[V]) Option[V]) Map[K, V] {
	v := r.Get(k)
	nv := remap(v)

	if nv.IsDefined() {
		return r.Updated(k, nv.Get())
	} else if v.IsDefined() {
		return r.Removed(k)
	}
	return r
}

func (r MapOps[K, V]) Iterator() Iterator[Tuple2[K, V]] {
	return r.Map.Iterator()
}

func (r MapOps[K, V]) Contains(k K) bool {
	return r.Get(k).IsDefined()
}

func (r MapOps[K, V]) Keys() Iterator[K] {
	itr := r.Iterator()
	return MakeIterator(itr.HasNext, func() K {
		return itr.Next().I1
	})
}

func (r MapOps[K, V]) Values() Iterator[V] {
	itr := r.Iterator()
	return MakeIterator(itr.HasNext, func() V {
		return itr.Next().I2
	})
}

func (r MapOps[K, V]) Foreach(f func(Tuple2[K, V])) {
	r.Iterator().Foreach(f)
}

func (r MapOps[K, V]) Concat(other Iterable[Tuple2[K, V]]) Map[K, V] {
	var ret Map[K, V] = r
	itr := other.Iterator()
	for itr.HasNext() {
		next := itr.Next()
		ret = ret.Updated(next.I1, next.I2)
	}

	return ret
}

func (r MapOps[K, V]) String() string {
	return fmt.Sprint(r.Map)
}

func MakeMap[K, V any](minimal MapMinimal[K, V]) Map[K, V] {
	return MapOps[K, V]{minimal}
}
