package slice

import "github.com/csgura/fp"

type SortedMap[K, V any] struct {
	ord    fp.Ord[K]
	sorted []fp.Tuple2[K, V]
}

func (r SortedMap[K, V]) Size() int {
	return len(r.sorted)
}

func bsearch[K, V any](s []fp.Tuple2[K, V], ord fp.Ord[K], k K) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (low + high) / 2
		midVal := s[mid].I1

		cmp := ord.Compare(midVal, k)
		if cmp < 0 {
			low = mid + 1
		} else if cmp > 0 {
			high = mid - 1
		} else {
			return mid // key found
		}
	}
	return -(low + 1) // key not found.
}

// Returns a view of the portion of this map whose keys are greater than or equal to fromKey
func (r SortedMap[K, V]) TailMap(fromKey K) SortedMap[K, V] {
	idx := bsearch(r.sorted, r.ord, fromKey)
	if idx < 0 {
		idx = -(idx + 1)
	}
	return SortedMap[K, V]{
		ord:    r.ord,
		sorted: r.sorted[idx:],
	}
}

func (r SortedMap[K, V]) Get(k K) fp.Option[V] {
	idx := bsearch(r.sorted, r.ord, k)
	if idx < 0 {
		return fp.None[V]()
	}
	return fp.Some(r.sorted[idx].I2)
}

func (r SortedMap[K, V]) Iterator() fp.Iterator[fp.Tuple2[K, V]] {
	return fp.IteratorOfSeq(r.sorted)
}
