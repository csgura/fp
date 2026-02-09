package fp

import (
	"fmt"
	"iter"
	"maps"
)

type MapBase[K, V any] interface {
	Size() int
	Get(k K) Option[V]
	Removed(k ...K) MapBase[K, V]
	Updated(k K, v V) MapBase[K, V]
	Iterator() Iterator[Tuple2[K, V]]
}

type MapBaseUpdatedWith[K, V any] interface {
	UpdatedWith(k K, remap func(Option[V]) Option[V]) MapBase[K, V]
}

type Map[K, V any] struct {
	Base MapBase[K, V]
}

func (r Map[K, V]) Size() int {
	if r.Base == nil {
		return 0
	}
	return r.Base.Size()
}

func (r Map[K, V]) IsEmpty() bool {
	return r.Size() == 0
}

func (r Map[K, V]) NonEmpty() bool {
	return r.Size() != 0
}

func (r Map[K, V]) Get(k K) Option[V] {
	if r.Base == nil {
		return None[V]()
	}
	return r.Base.Get(k)
}
func (r Map[K, V]) Removed(k ...K) Map[K, V] {
	if r.Base == nil {
		return r
	}
	return MakeMap(r.Base.Removed(k...))
}
func (r Map[K, V]) Updated(k K, v V) Map[K, V] {
	if r.Base == nil {
		return MakeMap(
			UnsafeGoMap[K, V]{
				k: v,
			},
		)
	}
	return MakeMap(r.Base.Updated(k, v))
}

func (r Map[K, V]) UpdatedWith(k K, remap func(Option[V]) Option[V]) Map[K, V] {
	if um, ok := r.Base.(MapBaseUpdatedWith[K, V]); ok {
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
	if r.Base == nil {
		return MakeIterator(func() bool {
			return false
		}, func() Tuple2[K, V] {
			panic(ErrIteratorEmpty)
		})
	}
	return r.Base.Iterator()
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
	if r.Base == nil {
		return "Map()"
	}
	return fmt.Sprint(r.Base)
}

func MakeMap[K, V any](Base MapBase[K, V]) Map[K, V] {
	return Map[K, V]{Base}
}

type UnsafeGoMap[K, V any] map[any]V

var _ MapBase[string, int] = UnsafeGoMap[string, int]{}

func (r UnsafeGoMap[K, V]) Get(k K) Option[V] {
	if v, ok := r[k]; ok {
		return Some(v)
	}
	return None[V]()
}

func (r UnsafeGoMap[K, V]) Size() int {
	return len(r)
}

func (r UnsafeGoMap[K, V]) Removed(k ...K) MapBase[K, V] {

	n := UnsafeGoMap[K, V]{}
	for ek, ev := range r {
		n[ek] = ev
	}

	for _, k := range k {
		delete(n, k)
	}
	return n
}

func (r UnsafeGoMap[K, V]) Updated(k K, v V) MapBase[K, V] {
	n := UnsafeGoMap[K, V]{}
	for ek, ev := range r {
		n[ek] = ev
	}

	n[k] = v
	return n
}

func (r UnsafeGoMap[K, V]) Iterator() Iterator[Tuple2[K, V]] {
	seq := []Tuple2[K, V]{}
	for k, v := range r {
		seq = append(seq, Tuple2[K, V]{k.(K), v})
	}
	return IteratorOfSeq(seq)
}

func seq2[K, V any](seq iter.Seq2[K, V]) iter.Seq[Tuple2[K, V]] {
	return func(yield func(Tuple2[K, V]) bool) {
		for k, v := range seq {
			if !yield(Tuple2[K, V]{
				I1: k,
				I2: v,
			}) {
				return
			}

		}
	}
}

func IteratorOfGoMap[K comparable, V any](m map[K]V) Iterator[Tuple2[K, V]] {
	// seq := []Tuple2[K, V]{}
	// for k, v := range m {
	// 	seq = append(seq, Tuple2[K, V]{k, v})
	// }
	// return IteratorOfSeq(seq)

	return MakePullIterator(seq2(maps.All(m)))
}
