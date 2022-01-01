package mutable

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

type Set[V comparable] map[V]bool

var _ fp.SetMinimal[string] = Set[string]{}

func (r Set[V]) Contains(v V) bool {
	return r[v]
}

func (r Set[V]) Size() int {
	return len(r)
}

func (r Set[V]) Iterator() fp.Iterator[V] {
	seq := fp.Seq[V]{}
	for k := range r {
		seq = append(seq, k)
	}
	return seq.Iterator()
}

func (r Set[V]) Incl(v V) fp.SetMinimal[V] {
	ret := Set[V]{}
	for k, v := range r {
		ret[k] = v
	}
	ret[v] = true
	return ret
}

func (r Set[V]) Excl(v V) fp.SetMinimal[V] {
	ret := Set[V]{}
	for k, v := range r {
		ret[k] = v
	}
	delete(ret, v)
	return ret
}

func SetOf[V comparable](v ...V) Set[V] {
	ret := Set[V]{}
	for _, e := range v {
		ret[e] = true
	}
	return ret
}

type Map[K comparable, V any] map[K]V

var _ fp.MapMinimal[string, int] = Map[string, int]{}

func (r Map[K, V]) Get(k K) fp.Option[V] {
	if v, ok := r[k]; ok {
		return option.Some(v)
	}
	return option.None[V]()
}

func (r Map[K, V]) Size() int {
	return len(r)
}

func (r Map[K, V]) Removed(k ...K) fp.MapMinimal[K, V] {
	s := SetOf(k...)

	nm := Map[K, V]{}
	for k, v := range r {
		if !s.Contains(k) {
			nm[k] = v
		}
	}
	return nm
}

func (r Map[K, V]) Updated(k K, v V) fp.MapMinimal[K, V] {

	nm := Map[K, V]{}
	for k, v := range r {
		nm[k] = v
	}
	nm[k] = v
	return nm
}

func (r Map[K, V]) Iterator() fp.Iterator[fp.Tuple2[K, V]] {
	seq := fp.Seq[fp.Tuple2[K, V]]{}
	for k, v := range r {
		seq = append(seq, as.Tuple2(k, v))
	}
	return seq.Iterator()
}
