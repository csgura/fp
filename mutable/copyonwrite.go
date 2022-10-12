package mutable

import (
	"sync"
	"sync/atomic"

	"github.com/csgura/fp"
)

type CopyOnWriteMap[K, V any] struct {
	value atomic.Value
	lock  sync.Mutex
}

var _ fp.MapMinimal[string, int] = &CopyOnWriteMap[string, int]{}

func (r *CopyOnWriteMap[K, V]) load() fp.UnsafeGoMap[K, V] {
	m := r.value.Load()

	if m == nil {
		r.lock.Lock()
		defer r.lock.Unlock()

		m = r.value.Load()
		if m == nil {
			m = fp.UnsafeGoMap[K, V]{}
			r.value.Store(m)
		}
	}
	return m.(fp.UnsafeGoMap[K, V])
}

func (r *CopyOnWriteMap[K, V]) copyOnWrite(f func(om fp.UnsafeGoMap[K, V]) fp.UnsafeGoMap[K, V]) fp.UnsafeGoMap[K, V] {

	r.lock.Lock()
	defer r.lock.Unlock()

	m := r.value.Load()
	if m == nil {
		m = fp.UnsafeGoMap[K, V]{}
	}

	nm := f(m.(fp.UnsafeGoMap[K, V]))
	r.value.Store(nm)
	return nm
}

func (r *CopyOnWriteMap[K, V]) Get(k K) fp.Option[V] {

	return r.load().Get(k)

}

func (r *CopyOnWriteMap[K, V]) Size() int {
	return r.load().Size()
}

func unsafeSet[V any](v ...V) fp.UnsageGoSet[V] {
	ret := fp.UnsageGoSet[V]{}
	for _, e := range v {
		ret[e] = true
	}
	return ret
}

func (r *CopyOnWriteMap[K, V]) Removed(k ...K) fp.MapMinimal[K, V] {

	r.copyOnWrite(func(om fp.UnsafeGoMap[K, V]) fp.UnsafeGoMap[K, V] {
		nm := fp.UnsafeGoMap[K, V]{}
		s := unsafeSet(k...)

		for k, v := range om {
			if !s.Contains(k.(K)) {
				nm[k] = v
			}
		}
		return nm
	})

	return r
}

func (r *CopyOnWriteMap[K, V]) ComputeIfAbsent(k K, f func() V) V {

	ret := r.Get(k)
	if ret.IsDefined() {
		return ret.Get()
	}

	nv := f()
	r.copyOnWrite(func(om fp.UnsafeGoMap[K, V]) fp.UnsafeGoMap[K, V] {
		nm := fp.UnsafeGoMap[K, V]{}

		for k, v := range om {
			nm[k] = v

		}
		nm[k] = nv
		return nm
	})

	return r.Get(k).Get()
}

func (r *CopyOnWriteMap[K, V]) Updated(k K, v V) fp.MapMinimal[K, V] {

	r.copyOnWrite(func(om fp.UnsafeGoMap[K, V]) fp.UnsafeGoMap[K, V] {
		nm := fp.UnsafeGoMap[K, V]{}

		for k, v := range om {
			nm[k] = v

		}
		nm[k] = v
		return nm
	})

	return r
}

func (r *CopyOnWriteMap[K, V]) UpdatedWith(k K, remap func(fp.Option[V]) fp.Option[V]) fp.MapMinimal[K, V] {
	r.copyOnWrite(func(om fp.UnsafeGoMap[K, V]) fp.UnsafeGoMap[K, V] {

		ov := om.Get(k)
		nv := remap(ov)

		if nv.IsDefined() {
			nm := fp.UnsafeGoMap[K, V]{}

			for k, v := range om {
				nm[k] = v
			}
			nm[k] = nv.Get()
			return nm
		} else if ov.IsDefined() {
			nm := fp.UnsafeGoMap[K, V]{}

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
