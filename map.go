package fp

type Set[V any] interface {
	Contains(v V) bool
	Size() int
	Iterator() Iterator[V]
	Incl(v V) Set[V]
	Excl(v V) Set[V]
}

// type Set[V comparable] map[V]bool

// func (r Set[V]) Contains(v V) bool {
// 	return r[v]
// }

// func (r Set[V]) Iterator() Iterator[V] {
// 	seq := Seq[V]{}
// 	for k := range r {
// 		seq = append(seq, k)
// 	}
// 	return seq.Iterator()
// }

// func SetOf[V comparable](v ...V) Set[V] {
// 	ret := Set[V]{}
// 	for _, e := range v {
// 		ret[e] = true
// 	}
// 	return ret
// }

type MapMinimal[K, V any] interface {
	Size() int
	Get(k K) Option[V]
	Removed(k ...K) Map[K, V]
	Updated(k K, v V) Map[K, V]
	Iterator() Iterator[Tuple2[K, V]]
}

type Map[K, V any] interface {
	MapMinimal[K, V]
	Contains(K) bool
	Keys() Iterator[K]
	Foreach(func(Tuple2[K, V]))
}

type MapAdaptor[K, V any] struct {
	MapMinimal[K, V]
}

func (r MapAdaptor[K, V]) Contains(k K) bool {
	return r.Get(k).IsDefined()
}

func (r MapAdaptor[K, V]) Keys() Iterator[K] {
	itr := r.Iterator()
	return MakeIterator(itr.HasNext, func() K {
		return itr.Next().I1
	})
}

func (r MapAdaptor[K, V]) Values() Iterator[V] {
	itr := r.Iterator()
	return MakeIterator(itr.HasNext, func() V {
		return itr.Next().I2
	})
}

func (r MapAdaptor[K, V]) Foreach(f func(Tuple2[K, V])) {
	r.Iterator().Foreach(f)
}
