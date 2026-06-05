package fp

type Slice[T any] = []T

type SliceT[T any] = Try[Slice[T]]

func SliceCasting[To ~[]T, From ~[]T, T any](a From) To {
	return To(a)
}

type SeqT[T any] = Try[Seq[T]]

type Seq[T any] []T

// func (r Seq[T]) Iterator() Iterator[T] {
// 	return iteratorFromSlice(r)
// }

func IteratorOfSeq[T any](r []T) Iterator[T] {
	idx := 0

	return MakeIterator(
		func() bool {
			return idx < len(r)
		},
		func() T {
			if idx < len(r) {
				ret := r[idx]
				idx++
				return ret
			}
			panic(ErrIteratorEmpty)
		},
	)
}

func IteratorOfOption[T any](r Option[T]) Iterator[T] {
	first := true

	return MakeIterator(
		func() bool {
			return first && r.IsDefined()
		},
		func() T {
			if first && r.IsDefined() {
				first = false
				return r.Get()
			}
			panic(ErrIteratorEmpty)
		},
	)
}
