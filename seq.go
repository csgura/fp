package fp

import (
	"bytes"
	"fmt"
)

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

func (r Seq[T]) Widen[_ Phantom[T]]() []T {
	return r
}

func (r Seq[T]) Size[_ Phantom[T]]() int {
	return len(r)
}

func (r Seq[T]) IsEmpty[_ Phantom[T]]() bool {
	return len(r) == 0
}

func (r Seq[T]) NonEmpty[_ Phantom[T]]() bool {
	return len(r) > 0
}

func (r Seq[T]) Get[_ Phantom[T]](idx int) Option[T] {
	if len(r) > idx {
		return Some(r[idx])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Head[_ Phantom[T]]() Option[T] {
	if len(r) > 0 {
		return Some(r[0])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Init[_ Phantom[T]]() Seq[T] {
	if len(r) > 1 {
		return r[:len(r)-1]
	} else {
		return nil
	}
}

func (r Seq[T]) Last[_ Phantom[T]]() Option[T] {
	if len(r) > 0 {
		return Some(r[len(r)-1])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Tail[_ Phantom[T]]() Seq[T] {
	if len(r) > 0 {
		return r[1:]
	} else {
		return nil
	}
}

func (r Seq[T]) UnSeq[_ Phantom[T]]() (Option[T], Seq[T]) {
	if len(r) > 0 {
		return r.Head(), r[1:]
	} else {
		return r.Head(), nil
	}
}

func (r Seq[T]) Take[_ Phantom[T]](n int) Seq[T] {
	if len(r) < n {
		return r
	}
	return r[0:n]
}

func (r Seq[T]) Drop[_ Phantom[T]](n int) Seq[T] {
	if len(r) < n {
		return nil
	}
	return r[n:]
}

func (r Seq[T]) Foreach[_ Phantom[T]](f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) Filter[_ Phantom[T]](p func(v T) bool) Seq[T] {
	ret := make([]T, 0, len(r))
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func (r Seq[T]) FilterMap[S any](p func(v T) Option[S]) Seq[S] {
	ret := make([]S, 0, len(r))
	for _, v := range r {
		e := p(v)
		if e.IsDefined() {
			ret = append(ret, e.Get())
		}
	}
	return ret
}

func (r Seq[T]) FilterNot[_ Phantom[T]](p func(v T) bool) Seq[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Seq[T]) Exists[_ Phantom[T]](p func(v T) bool) bool {
	for _, v := range r {
		if p(v) {
			return true
		}
	}
	return false
}

func (r Seq[T]) ForAll[_ Phantom[T]](p func(v T) bool) bool {
	for _, v := range r {
		if !p(v) {
			return false
		}
	}
	return true
}

func (r Seq[T]) Find[_ Phantom[T]](p func(v T) bool) Option[T] {
	for _, v := range r {
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Seq[T]) Add[_ Phantom[T]](item T) Seq[T] {
	return r.Append(item)
}

func (r Seq[T]) Append[_ Phantom[T]](items ...T) Seq[T] {
	if len(items) > 0 {
		tail := Seq[T](items)
		ret := make(Seq[T], len(r)+tail.Size())

		copy(ret, r)

		for i := range tail {
			ret[i+len(r)] = tail[i]
		}

		return ret
	}
	return r
}

func (r Seq[T]) Concat[_ Phantom[T]](tail Seq[T]) Seq[T] {
	if len(tail) > 0 {
		ret := make(Seq[T], len(r)+tail.Size())

		copy(ret, r)

		for i := range tail {
			ret[i+len(r)] = tail[i]
		}

		return ret
	}
	return r
}

// func (r Seq[T]) Reduce(m Monoid[T]) T {
// 	if len(r) == 0 {
// 		return m.Empty()
// 	}

// 	reduce := m.Empty()
// 	for i := 0; i < len(r); i++ {
// 		reduce = m.Combine(reduce, r[i])
// 	}
// 	return reduce
// }

func (r Seq[T]) Reverse[_ Phantom[T]]() Seq[T] {
	ret := make(Seq[T], len(r))

	for i := range r {
		ret[len(r)-i-1] = r[i]
	}

	return ret
}

func (r Seq[T]) MakeString[_ Phantom[T]](sep string) string {
	buf := &bytes.Buffer{}

	for i, v := range r {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(fmt.Sprint(v))
	}
	return buf.String()
}

func (r Seq[T]) Map[R any](mf func(T) R) Seq[R] {
	var ret = make([]R, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v))
	}
	return ret
}

func (r Seq[T]) FlatMap[R any](mf func(T) Seq[R]) Seq[R] {
	var ret = make([]R, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v)...)
	}
	return ret
}

func (r Seq[T]) Map2[U, R any](other Seq[U], f func(T, U) R) Seq[R] {
	return r.FlatMap(func(t T) Seq[R] {
		return other.Map(func(u U) R {
			return f(t, u)
		})
	})
}

func (r Seq[T]) FoldT[ACC any](zero ACC, f func(ACC, T) Try[ACC]) Try[ACC] {
	sum := zero
	for _, v := range r {
		t := f(sum, v)
		if t.IsSuccess() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return Success(sum)
}

func (r Seq[T]) Fold[ACC any](zero ACC, f func(ACC, T) ACC) ACC {
	sum := zero
	for _, v := range r {
		sum = f(sum, v)
	}
	return sum
}

func (r Seq[T]) FoldF[ACC any](zero ACC, f func(ACC, T) Future[ACC], ctx ...Executor) Future[ACC] {
	p := NewPromise[ACC]()
	p.Success(zero)

	return r.Fold(p.Future(), func(accf Future[ACC], t T) Future[ACC] {
		return accf.FlatMap(func(acc ACC) Future[ACC] {
			return f(acc, t)
		}, ctx...)
	})
}

func (r Seq[T]) TraverseT[R any](f func(T) Try[R]) Try[Seq[R]] {
	return r.FoldT(nil, func(a Seq[R], t T) Try[Seq[R]] {
		return f(t).Map(a.Add)
	})
}

func (r Seq[T]) TraverseF[R any](f func(T) Future[R], ctx ...Executor) Future[Seq[R]] {
	return r.FoldF(nil, func(acc Seq[R], t T) Future[Seq[R]] {
		return f(t).Map(acc.Add, ctx...)
	}, ctx...)
}
