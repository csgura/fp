//go:build !go1.27

package fp

import (
	"bytes"
	"fmt"
	"net/http"
)

func (r Option[T]) Map(mf func(T) T) Option[T] {
	if r.IsDefined() {
		r.v = mf(r.v)
	}
	return r
}

func (r Option[T]) FlatMap(mf func(T) Option[T]) Option[T] {
	if r.IsDefined() {
		return mf(r.v)
	}
	return r
}

func (r Try[T]) Map(mf func(T) T) Try[T] {
	if r.IsSuccess() {
		r.v = mf(r.v)
	}
	return r
}

func (r Try[T]) FlatMap(mf func(T) Try[T]) Try[T] {
	if r.IsSuccess() {
		return mf(r.v)
	}
	return r
}

func (r Future[T]) Map(mf func(T) T, ctx ...Executor) Future[T] {
	np := NewPromise[T]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			np.Success(mf(t.Get()))
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) FlatMap(mf func(T) Future[T], ctx ...Executor) Future[T] {
	np := NewPromise[T]()

	r.OnComplete(func(t Try[T]) {
		if t.IsSuccess() {
			mf(t.Get()).OnComplete(func(t Try[T]) {
				np.Complete(t)
			}, ctx...)
		} else {
			np.Failure(t.Failed().Get())
		}
	}, ctx...)

	return np.Future()
}

func (r Seq[T]) Map(mf func(T) T) Seq[T] {
	var ret = make([]T, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v))
	}
	return ret
}

func (r Seq[T]) FlatMap(mf func(T) Seq[T]) Seq[T] {
	var ret = make([]T, 0, len(r))
	for _, v := range r {
		ret = append(ret, mf(v)...)
	}
	return ret
}

func (r Option[T]) All() GoIter[T] {
	return func(f func(T) bool) {
		if r.IsDefined() {
			f(r.Get())
		}
	}
}

func (r Option[T]) Foreach(f func(v T)) {
	if r.IsDefined() {
		f(r.Get())
	}
}

func (r Option[T]) IsDefined() bool {
	return r.present
}

func (r Option[T]) IsEmpty() bool {
	return !r.IsDefined()
}

func (r Option[T]) Get() T {
	if r.IsDefined() {
		return r.v
	}
	panic(ErrOptionEmpty)
}

func (r Option[T]) Filter(p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if p(r.Get()) {
			return r
		}
	}
	return None[T]()

}

func (r Option[T]) FilterNot(p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if !p(r.Get()) {
			return r
		}
	}
	return None[T]()

}
func (r Option[T]) OrElse(t T) T {
	if r.IsDefined() {
		return r.Get()
	}
	return t
}

func (r Option[T]) OrZero() T {
	return r.OrElseGet(Zero[T])
}

func (r Option[T]) OrElseGet(f func() T) T {
	if r.IsDefined() {
		return r.Get()
	}
	return f()
}
func (r Option[T]) Or(f func() Option[T]) Option[T] {
	if r.IsDefined() {
		return r
	}
	return f()
}

func (r Option[T]) OrOption(v Option[T]) Option[T] {
	if r.IsDefined() {
		return r
	}
	return v
}

func (r Option[T]) OrPtr(v *T) Option[T] {
	if r.IsDefined() {
		return r
	}
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func (r Option[T]) ToSeq() []T {
	if r.IsDefined() {
		return []T{r.Get()}
	}
	return nil
}

func (r Option[T]) Ptr() *T {
	if r.IsDefined() {
		return &r.v
	}

	return nil
}

func (r Option[T]) Exists(p func(v T) bool) bool {
	return r.IsDefined() && p(r.v)
}

func (r Option[T]) ForAll(p func(v T) bool) bool {
	return r.IsEmpty() || p(r.v)
}

func (r Try[T]) All() GoIter[T] {
	return func(f func(T) bool) {
		if r.IsSuccess() {
			f(r.Get())
		}
	}
}

func (r Try[T]) IsSuccess() bool {
	return r.success
}

func (r Try[T]) IsFailure() bool {
	return !r.IsSuccess()
}

func (r Try[T]) Get() T {
	if r.IsSuccess() {
		return r.v
	}
	panic(r.Failed().Get())
}

func (r Try[T]) Unapply() (T, error) {
	if r.IsSuccess() {
		return r.Get(), nil
	} else {
		var zero T
		return zero, r.Failed().Get()
	}
}

func (r Try[T]) MapError(mf func(error) error) Try[T] {
	if r.IsFailure() {
		r.err = mf(r.err)
	}
	return r
}

func (r Try[T]) Foreach(f func(v T)) {
	if r.IsSuccess() {
		f(r.Get())
	}
}
func (r Try[T]) Failed() Try[error] {
	if r.IsSuccess() {
		return Failure[error](ErrTryNotFailed)
	}
	if r.err == nil {
		return Failure[error](Error(http.StatusNotAcceptable, "Try not initialized correctly"))
	}
	return Success(r.err)
}
func (r Try[T]) OrElse(t T) T {
	if r.IsSuccess() {
		return r.Get()
	}
	return t
}

func (r Try[T]) OrZero() T {
	return r.OrElseGet(Zero[T])
}

func (r Try[T]) OrElseGet(f func() T) T {
	if r.IsSuccess() {
		return r.Get()
	}
	return f()
}

func (r Try[T]) Or(f func() Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return f()
}

func (r Try[T]) OrTry(v Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return v
}

func (r Try[T]) Recover(f func(err error) T) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return Success(f(r.Failed().Get()))

}

func (r Try[T]) RecoverCase(isDefinedAt func(error) bool, then func(error) T) Try[T] {
	if r.IsSuccess() {
		return r
	}

	if isDefinedAt(r.Failed().Get()) {
		return Success(then(r.Failed().Get()))
	}

	return r
}

func (r Try[T]) RecoverCaseWith(isDefinedAt func(error) bool, then func(error) Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}

	if isDefinedAt(r.Failed().Get()) {
		return then(r.Failed().Get())
	}

	return r
}

func (r Try[T]) RecoverWith(f func(err error) Try[T]) Try[T] {
	if r.IsSuccess() {
		return r
	}
	return f(r.Failed().Get())
}

func (r Try[T]) ToSeq() []T {
	if r.IsSuccess() {
		return []T{r.Get()}
	}
	return nil
}

func (r Seq[T]) Widen() []T {
	return r
}

func (r Seq[T]) Size() int {
	return len(r)
}

func (r Seq[T]) IsEmpty() bool {
	return r.Size() == 0
}

func (r Seq[T]) NonEmpty() bool {
	return r.Size() > 0
}

func (r Seq[T]) Get(idx int) Option[T] {
	if r.Size() > idx {
		return Some(r[idx])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Head() Option[T] {
	if r.Size() > 0 {
		return Some(r[0])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Init() Seq[T] {
	if r.Size() > 1 {
		return r[:r.Size()-1]
	} else {
		return nil
	}
}

func (r Seq[T]) Last() Option[T] {
	if r.Size() > 0 {
		return Some(r[r.Size()-1])
	} else {
		return None[T]()
	}
}

func (r Seq[T]) Tail() Seq[T] {
	if r.Size() > 0 {
		return r[1:]
	} else {
		return nil
	}
}

func (r Seq[T]) UnSeq() (Option[T], Seq[T]) {
	if r.Size() > 0 {
		return r.Head(), r[1:]
	} else {
		return r.Head(), nil
	}
}

func (r Seq[T]) Take(n int) Seq[T] {
	if len(r) < n {
		return r
	}
	return r[0:n]
}

func (r Seq[T]) Drop(n int) Seq[T] {
	if len(r) < n {
		return nil
	}
	return r[n:]
}

func (r Seq[T]) Foreach(f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) Filter(p func(v T) bool) Seq[T] {
	ret := make([]T, 0, len(r))
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func (r Seq[T]) FilterNot(p func(v T) bool) Seq[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Seq[T]) Exists(p func(v T) bool) bool {
	for _, v := range r {
		if p(v) {
			return true
		}
	}
	return false
}

func (r Seq[T]) ForAll(p func(v T) bool) bool {
	for _, v := range r {
		if !p(v) {
			return false
		}
	}
	return true
}

func (r Seq[T]) Find(p func(v T) bool) Option[T] {
	for _, v := range r {
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Seq[T]) Add(item T) Seq[T] {
	return r.Append(item)
}

func (r Seq[T]) Append(items ...T) Seq[T] {
	if len(items) > 0 {
		tail := Seq[T](items)
		ret := make(Seq[T], r.Size()+tail.Size())

		copy(ret, r)

		for i := range tail {
			ret[i+r.Size()] = tail[i]
		}

		return ret
	}
	return r
}

func (r Seq[T]) Concat(tail Seq[T]) Seq[T] {
	if len(tail) > 0 {
		ret := make(Seq[T], r.Size()+tail.Size())

		copy(ret, r)

		for i := range tail {
			ret[i+r.Size()] = tail[i]
		}

		return ret
	}
	return r
}

// func (r Seq[T]) Reduce(m Monoid[T]) T {
// 	if r.Size() == 0 {
// 		return m.Empty()
// 	}

// 	reduce := m.Empty()
// 	for i := 0; i < len(r); i++ {
// 		reduce = m.Combine(reduce, r[i])
// 	}
// 	return reduce
// }

func (r Seq[T]) Reverse() Seq[T] {
	ret := make(Seq[T], r.Size())

	for i := range r {
		ret[r.Size()-i-1] = r[i]
	}

	return ret
}

func (r Seq[T]) MakeString(sep string) string {
	buf := &bytes.Buffer{}

	for i, v := range r {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(fmt.Sprint(v))
	}
	return buf.String()
}

func (r Iterator[T]) All() func(func(T) bool) {
	return func(f func(T) bool) {
		for r.hasNext() {
			if !f(r.Next()) {
				return
			}
		}
	}
}

func (r Iterator[T]) ToSeq() []T {
	ret := []T{}
	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func (r Iterator[T]) Count() int {
	ret := 0
	for r.HasNext() {
		r.Next()
		ret++
	}
	return ret
}

// func (r Iterator[T]) ToList() List[T] {

// 	head := r.NextOption()

// 	return MakeList(func() Option[T] {
// 		return head
// 	}, func() List[T] {
// 		return r.ToList()
// 	})
// }

func (r Iterator[T]) MakeString(sep string) string {
	buf := &bytes.Buffer{}

	first := true
	for r.HasNext() {
		if !first {
			buf.WriteString(sep)
		} else {
			first = false
		}

		v := r.Next()
		buf.WriteString(fmt.Sprint(v))

	}

	return buf.String()
}

func (r Iterator[T]) HasNext() bool {
	if r.hasNext == nil {
		return false
	}

	return r.hasNext()
}

func (r Iterator[T]) Next() T {
	return r.next()
}

func (r Iterator[T]) NextOption() Option[T] {
	if r.HasNext() {
		v := r.next()
		return Some(v)
	}
	return None[T]()
}

func (r Iterator[T]) Take(n int) Iterator[T] {

	i := 0
	hasNext := func() bool {
		if i < n {
			return r.HasNext()
		}
		return false
	}
	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {
				i++
				return r.Next()
			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) nextOnEmpty() T {
	panic(ErrIteratorEmpty)
}

func (r Iterator[T]) TakeWhile(p func(T) bool) Iterator[T] {

	breaking := false
	var fv Option[T] = None[T]()

	hasNext := func() bool {
		if breaking {
			return false
		}

		if fv.IsDefined() {
			return true
		}

		if r.HasNext() {
			v := r.Next()
			if p(v) {
				fv = Some(v)
				return true
			}
			breaking = true
		}
		return false
	}

	return MakeIterator(
		hasNext,
		func() T {

			if hasNext() {
				ret := fv.Get()
				fv = None[T]()
				return ret
			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) Drop(n int) Iterator[T] {

	for i := 0; i < n && r.HasNext(); i++ {
		r.Next()
	}

	return r
}

func (r Iterator[T]) DropWhile(p func(T) bool) Iterator[T] {

	found := false
	var first Option[T] = None[T]()
	hasNext := func() bool {
		if first.IsDefined() {
			return true
		}

		if found {
			return r.HasNext()
		}

		for r.HasNext() {
			v := r.Next()
			if !p(v) {
				first = Some(v)
				found = true
				return true
			}
		}
		return false
	}

	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {
				if first.IsDefined() {
					ret := first.Get()
					first = None[T]()
					return ret
				}

				if found {
					return r.Next()
				}

			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) Filter(p func(T) bool) Iterator[T] {

	first := true
	var fv Option[T] = None[T]()

	hasNext := func() bool {
		if first {
			fv = r.Find(p)
			first = false
		}
		return fv.IsDefined()
	}

	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {

				ret := fv.Get()
				fv = r.Find(p)
				return ret
			}
			return r.nextOnEmpty()
		},
	)
}

func (r Iterator[T]) FilterNot(p func(T) bool) Iterator[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Iterator[T]) Find(p func(T) bool) Option[T] {
	for r.HasNext() {
		v := r.Next()
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Iterator[T]) Foreach(p func(T)) {
	for r.HasNext() {
		v := r.Next()
		p(v)
	}
}

func (r Iterator[T]) TapEach(p func(T)) Iterator[T] {
	return MakeIterator(
		func() bool {
			return r.HasNext()
		},
		func() T {
			ret := r.next()
			p(ret)
			return ret
		},
	)
}

func (r Iterator[T]) Appended(elem T) Iterator[T] {
	return r.Concat(IteratorOfSeq([]T{elem}))
}

func (r Iterator[T]) Concat(tail Iterator[T]) Iterator[T] {

	alliter := r.concat
	if len(alliter) == 0 {
		alliter = make([]Iterator[T], 1, 2+len(tail.concat))
		alliter[0] = r
	}

	if len(tail.concat) > 0 {
		alliter = append(alliter, tail.concat...)
	} else {
		alliter = append(alliter, tail)
	}

	currentItr := Some(alliter[0])
	remainItr := alliter[1:]

	currentNextChecked := false

	currentNext := func() bool {

		if currentNextChecked {
			return true
		}

		if currentItr.IsEmpty() {
			return false
		}

		if currentItr.Get().HasNext() {
			currentNextChecked = true
			return true
		}

		for i, itr := range remainItr {
			if itr.HasNext() {
				currentItr = Some(itr)
				remainItr = remainItr[i+1:]
				currentNextChecked = true
				return true
			}
		}

		currentItr = None[Iterator[T]]()

		return false
	}

	ret := MakeIterator(
		func() bool {

			return currentNext()
		},
		func() T {
			if currentNext() {
				currentNextChecked = false
				return currentItr.Get().Next()
			}
			panic(ErrIteratorEmpty)
		},
	)
	ret.concat = alliter

	return ret
}

func (r Iterator[T]) Map(mf func(T) T) Iterator[T] {
	return MakeIterator(
		func() bool {
			return r.HasNext()
		},
		func() T {
			return mf(r.Next())
		},
	)
}

func (r Iterator[T]) FlatMap(mf func(T) Iterator[T]) Iterator[T] {
	current := None[Iterator[T]]()

	hasNext := func() bool {
		if current.IsDefined() && current.Get().HasNext() {
			return true
		}

		for r.HasNext() {
			nextItr := mf(r.Next())
			current = Some(nextItr)
			if nextItr.HasNext() {
				return true
			}
		}

		return false
	}

	return MakeIterator(
		hasNext,
		func() T {
			if hasNext() {
				return current.Get().Next()
			}
			panic(ErrIteratorEmpty)
		},
	)
}

// func (r Iterator[T]) Reduce(m Monoid[T]) T {
// 	ret := m.Empty()
// 	for r.HasNext() {
// 		v := r.Next()
// 		m.Combine(ret, v)
// 	}
// 	return ret
// }

func (r Iterator[T]) Exists(p func(v T) bool) bool {
	for r.HasNext() {
		if p(r.Next()) {
			return true
		}
	}

	return false
}

func (r Iterator[T]) ForAll(p func(v T) bool) bool {
	for r.HasNext() {
		if !p(r.Next()) {
			return false
		}
	}
	return true
}

func (r Iterator[T]) IsEmpty() bool {
	return !r.HasNext()
}

func (r Iterator[T]) NonEmpty() bool {
	return r.HasNext()
}
