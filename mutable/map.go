package mutable

import (
	"maps"

	"github.com/csgura/fp"
)

type Set[V comparable] map[V]bool

var _ fp.SetMinimal[string] = Set[string]{}

func (r Set[V]) Contains(v V) bool {
	return r[v]
}

func (r Set[V]) Size() int {
	return len(r)
}

func (r Set[V]) ToSeq() []V {
	seq := []V{}
	for k := range r {
		seq = append(seq, k)
	}
	return seq
}

func (r Set[V]) Iterator() fp.Iterator[V] {
	// seq := []V{}
	// for k := range r {
	// 	seq = append(seq, k)
	// }
	// return fp.IteratorOfSeq(seq)

	return fp.MakePullIterator(maps.Keys(r))
}

func (r Set[V]) Incl(v V) fp.SetMinimal[V] {
	r[v] = true
	return r
}

func (r Set[V]) Excl(v V) fp.SetMinimal[V] {
	delete(r, v)
	return r
}

func AsFpSet[V comparable](s Set[V]) fp.Set[V] {
	return fp.MakeSet(func() fp.SetMinimal[V] {
		return Set[V]{}
	}, s)
}
func SetOf[V comparable](v ...V) fp.Set[V] {
	ret := Set[V]{}
	for _, e := range v {
		ret[e] = true
	}
	return AsFpSet(ret)
}

func EmptySet[V comparable]() fp.Set[V] {
	return SetOf[V]()
}

func MapOf[K comparable, V any](m map[K]V) fp.Map[K, V] {
	return fp.MakeMap(Map[K, V](m))
}

func EmptyMap[K comparable, V any]() fp.Map[K, V] {
	return fp.MakeMap(Map[K, V]{})
}

type Map[K comparable, V any] map[K]V

var _ fp.MapBase[string, int] = Map[string, int]{}

func (r Map[K, V]) Get(k K) fp.Option[V] {
	if v, ok := r[k]; ok {
		return fp.Some(v)
	}
	return fp.None[V]()
}

func (r Map[K, V]) Size() int {
	return len(r)
}

func (r Map[K, V]) Removed(k ...K) fp.MapBase[K, V] {

	for _, k := range k {
		delete(r, k)
	}
	return r
}

func (r Map[K, V]) Updated(k K, v V) fp.MapBase[K, V] {

	r[k] = v
	return r
}

func (r Map[K, V]) Iterator() fp.Iterator[fp.Tuple2[K, V]] {
	// seq := []fp.Tuple2[K, V]{}
	// for k, v := range r {
	// 	seq = append(seq, fp.Tuple2[K, V]{
	// 		I1: k,
	// 		I2: v,
	// 	})
	// }
	// return fp.IteratorOfSeq(seq)

	return fp.IteratorOfGoMap(r)
}

func (r Map[K, V]) All() fp.GoIter[fp.Tuple2[K, V]] {
	return func(yield func(fp.Tuple2[K, V]) bool) {
		for k, v := range r {
			if !yield(fp.Tuple2[K, V]{I1: k, I2: v}) {
				return
			}
		}
	}
}
