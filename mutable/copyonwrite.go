package mutable

import (
	"sync"
	"sync/atomic"

	"github.com/csgura/fp"
)

type CopyOnWriteMap[K comparable, V any] struct {
	value atomic.Value
	lock  sync.Mutex
}

var _ fp.MapBase[string, int] = &CopyOnWriteMap[string, int]{}

func (r *CopyOnWriteMap[K, V]) load() Map[K, V] {
	m := r.value.Load()

	if m == nil {
		r.lock.Lock()
		defer r.lock.Unlock()

		m = r.value.Load()
		if m == nil {
			m = Map[K, V]{}
			r.value.Store(m)
		}
	}
	return m.(Map[K, V])
}

func (r *CopyOnWriteMap[K, V]) copyOnWrite(f func(om Map[K, V]) Map[K, V]) Map[K, V] {

	r.lock.Lock()
	defer r.lock.Unlock()

	m := r.value.Load()
	if m == nil {
		m = Map[K, V]{}
	}

	nm := f(m.(Map[K, V]))
	r.value.Store(nm)
	return nm
}

func (r *CopyOnWriteMap[K, V]) Get(k K) fp.Option[V] {

	return r.load().Get(k)

}

func (r *CopyOnWriteMap[K, V]) Size() int {
	return r.load().Size()
}

func unsafeSet[V comparable](v ...V) Set[V] {
	ret := Set[V]{}
	for _, e := range v {
		ret[e] = true
	}
	return ret
}

func (r *CopyOnWriteMap[K, V]) Removed(k ...K) fp.MapBase[K, V] {

	r.copyOnWrite(func(om Map[K, V]) Map[K, V] {
		nm := Map[K, V]{}
		s := unsafeSet(k...)

		for k, v := range om {
			if !s.Contains(k) {
				nm[k] = v
			}
		}
		return nm
	})

	return r
}

func (r *CopyOnWriteMap[K, V]) ComputeIfAbsent(k K, f func() V) V {
	return r.ComputeIf(k, func(V) bool {
		return false
	}, f)
}

func (r *CopyOnWriteMap[K, V]) ComputeIf(k K, pred func(V) bool, f func() V) V {

	ret := r.Get(k).FilterNot(pred)
	if ret.IsDefined() {
		return ret.Get()
	}

	nv := f()
	r.copyOnWrite(func(om Map[K, V]) Map[K, V] {
		nm := Map[K, V]{}

		for k, v := range om {
			nm[k] = v

		}
		nm[k] = nv
		return nm
	})

	return r.Get(k).Get()
}

func (r *CopyOnWriteMap[K, V]) Updated(k K, v V) fp.MapBase[K, V] {

	r.copyOnWrite(func(om Map[K, V]) Map[K, V] {
		nm := Map[K, V]{}

		for k, v := range om {
			nm[k] = v

		}
		nm[k] = v
		return nm
	})

	return r
}

func (r *CopyOnWriteMap[K, V]) UpdatedWith(k K, remap func(fp.Option[V]) fp.Option[V]) fp.MapBase[K, V] {
	r.copyOnWrite(func(om Map[K, V]) Map[K, V] {

		ov := om.Get(k)
		nv := remap(ov)

		if nv.IsDefined() {
			nm := Map[K, V]{}

			for k, v := range om {
				nm[k] = v
			}
			nm[k] = nv.Get()
			return nm
		} else if ov.IsDefined() {
			nm := Map[K, V]{}

			for k, v := range om {
				nm[k] = v
			}
			delete(nm, k)
			return nm
		}
		return om
	})

	return r
}

func (r *CopyOnWriteMap[K, V]) Iterator() fp.Iterator[fp.Tuple2[K, V]] {
	return r.load().Iterator()
}
