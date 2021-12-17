package fp

type Set[V comparable] map[V]bool

func (r Set[V]) Contains(v V) bool {
	return r[v]
}

func (r Set[V]) Iterator() Iterator[V] {
	seq := Seq[V]{}
	for k := range r {
		seq = append(seq, k)
	}
	return seq.Iterator()
}

func SetOf[V comparable](v ...V) Set[V] {
	ret := Set[V]{}
	for _, e := range v {
		ret[e] = true
	}
	return ret
}

type Map[K comparable, V any] map[K]V

func (r Map[K, V]) Get(k K) Option[V] {
	if v, ok := r[k]; ok {
		return Some[V]{v}
	}
	return None[V]{}
}

func (r Map[K, V]) Removed(k ...K) Map[K, V] {
	s := SetOf(k...)

	nm := Map[K, V]{}
	for k, v := range r {
		if !s.Contains(k) {
			nm[k] = v
		}
	}
	return nm
}

func (r Map[K, V]) Updated(k K, v V) Map[K, V] {

	nm := Map[K, V]{}
	for k, v := range r {
		nm[k] = v
	}
	nm[k] = v
	return nm
}

func (r Map[K, V]) Iterator() Iterator[Tuple2[K, V]] {
	seq := Seq[Tuple2[K, V]]{}
	for k, v := range r {
		seq = append(seq, Tuple2[K, V]{k, v})
	}
	return seq.Iterator()
}
