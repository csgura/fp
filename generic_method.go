//go:build go1.27

package fp

import (
	"bytes"
	"fmt"
	"net/http"
)

func (r Option[T]) Map[R any](mf func(T) R) Option[R] {
	if r.IsDefined() {
		return Some(mf(r.v))
	}
	return None[R]()
}

func (r Option[T]) FlatMap[R any](mf func(T) Option[R]) Option[R] {
	if r.IsDefined() {
		return mf(r.v)
	}
	return None[R]()
}

func (r Option[T]) Replace[R any](o R) Option[R] {
	return r.Map(Const[T](o))
}

func (r Option[T]) ReplaceS[R any](f func() R) Option[R] {
	return r.Map(func(t T) R {
		return f()
	})
}

func (r Option[T]) Void[R any]() Option[Unit] {
	return r.Replace(Unit{})
}

func (r Option[T]) Map2[U, R any](other Option[U], f func(T, U) R) Option[R] {
	return r.FlatMap(func(t T) Option[R] {
		return other.Map(func(u U) R {
			return f(t, u)
		})
	})
}

func (r Option[T]) IntoTry[E error](err func() E) Try[T] {
	if r.IsDefined() {
		return Success(r.Get())
	}
	return Failure[T](err())
}

func (r Option[T]) IntoFuture[E error](err func() E) Future[T] {
	p := NewPromise[T]()
	if r.IsDefined() {
		p.Success(r.Get())
		return p.Future()
	}
	p.Failure(err())
	return p.Future()
}

func (r Option[T]) TraverseT[R any](f func(T) Try[R]) Try[Option[R]] {
	if r.IsDefined() {
		return f(r.Get()).Map(Some)
	}
	return Success(None[R]())
}

func (r Option[T]) TraverseF[R any](f func(T) Future[R]) Future[Option[R]] {
	if r.IsDefined() {
		return f(r.Get()).Map(Some)
	}
	p := NewPromise[Option[R]]()
	p.Success(None[R]())
	return p.Future()
}

func (r Option[T]) All[_ Phantom[T]]() GoIter[T] {
	return func(f func(T) bool) {
		if r.IsDefined() {
			f(r.Get())
		}
	}
}

func (r Option[T]) Foreach[_ Phantom[T]](f func(v T)) {
	if r.IsDefined() {
		f(r.Get())
	}
}

func (r Option[T]) IsDefined[_ Phantom[T]]() bool {
	return r.present
}

func (r Option[T]) IsEmpty[_ Phantom[T]]() bool {
	return !r.IsDefined()
}

func (r Option[T]) Get[_ Phantom[T]]() T {
	if r.IsDefined() {
		return r.v
	}
	panic(ErrOptionEmpty)
}

func (r Option[T]) Filter[_ Phantom[T]](p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if p(r.Get()) {
			return r
		}
	}
	return None[T]()

}

func (r Option[T]) FilterNot[_ Phantom[T]](p func(v T) bool) Option[T] {
	if r.IsDefined() {
		if !p(r.Get()) {
			return r
		}
	}
	return None[T]()

}
func (r Option[T]) OrElse[_ Phantom[T]](t T) T {
	if r.IsDefined() {
		return r.Get()
	}
	return t
}

func (r Option[T]) OrZero[_ Phantom[T]]() T {
	return r.OrElseGet(Zero[T])
}

func (r Option[T]) OrElseGet[_ Phantom[T]](f func() T) T {
	if r.IsDefined() {
		return r.Get()
	}
	return f()
}
func (r Option[T]) Or[_ Phantom[T]](f func() Option[T]) Option[T] {
	if r.IsDefined() {
		return r
	}
	return f()
}

func (r Option[T]) OrOption[_ Phantom[T]](v Option[T]) Option[T] {
	if r.IsDefined() {
		return r
	}
	return v
}

func (r Option[T]) OrPtr[_ Phantom[T]](v *T) Option[T] {
	if r.IsDefined() {
		return r
	}
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func (r Option[T]) ToSeq[_ Phantom[T]]() []T {
	if r.IsDefined() {
		return []T{r.Get()}
	}
	return nil
}

func (r Option[T]) Ptr[_ Phantom[T]]() *T {
	if r.IsDefined() {
		return &r.v
	}

	return nil
}

func (r Option[T]) Exists[_ Phantom[T]](p func(v T) bool) bool {
	return r.IsDefined() && p(r.v)
}

func (r Option[T]) ForAll[_ Phantom[T]](p func(v T) bool) bool {
	return r.IsEmpty() || p(r.v)
}

func (r Try[T]) All[_ Phantom[T]]() GoIter[T] {
	return func(f func(T) bool) {
		if r.success {
			f(r.Get())
		}
	}
}

func (r Try[T]) IsSuccess[_ Phantom[T]]() bool {
	return r.success
}

func (r Try[T]) IsFailure[_ Phantom[T]]() bool {
	return !r.success
}

func (r Try[T]) Get[_ Phantom[T]]() T {
	if r.success {
		return r.v
	}
	if r.err == nil {
		panic(Error(http.StatusNotAcceptable, "Try not initialized correctly"))
	}
	panic(ErrTryNotFailed)
}

func (r Try[T]) Unapply() (T, error) {
	if r.success {
		return r.v, nil
	} else if r.err == nil {
		var zero T
		return zero, Error(http.StatusNotAcceptable, "Try not initialized correctly")
	} else {
		var zero T
		return zero, r.err
	}
}

func (r Try[T]) MapError[_ Phantom[T]](mf func(error) error) Try[T] {
	if !r.success {
		r.err = mf(r.err)
	}
	return r
}

func (r Try[T]) Foreach[_ Phantom[T]](f func(v T)) {
	if r.success {
		f(r.v)
	}
}
func (r Try[T]) Failed[_ Phantom[T]]() Try[error] {
	if r.success {
		return Failure[error](ErrTryNotFailed)
	}
	if r.err == nil {
		return Failure[error](Error(http.StatusNotAcceptable, "Try not initialized correctly"))
	}
	return Success(r.err)
}

func (r Try[T]) OrElse[_ Phantom[T]](t T) T {
	if r.success {
		return r.v
	}
	return t
}

func (r Try[T]) OrZero[_ Phantom[T]]() T {
	if r.success {
		return r.v
	}
	var zero T
	return zero
}

func (r Try[T]) OrElseGet[_ Phantom[T]](f func() T) T {
	if r.success {
		return r.v
	}
	return f()
}

func (r Try[T]) Or[_ Phantom[T]](f func() Try[T]) Try[T] {
	if r.success {
		return r
	}
	return f()
}

func (r Try[T]) OrTry[_ Phantom[T]](v Try[T]) Try[T] {
	if r.success {
		return r
	}
	return v
}

func (r Try[T]) Recover[_ Phantom[T]](f func(err error) T) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}
	return Success(f(err))

}

func (r Try[T]) RecoverCase[_ Phantom[T]](isDefinedAt func(error) bool, then func(error) T) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}

	if isDefinedAt(err) {
		return Success(then(err))
	}

	return r
}

func (r Try[T]) RecoverCaseWith[_ Phantom[T]](isDefinedAt func(error) bool, then func(error) Try[T]) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}

	if isDefinedAt(err) {
		return then(err)
	}

	return r
}

func (r Try[T]) RecoverWith[_ Phantom[T]](f func(err error) Try[T]) Try[T] {
	_, err := r.Unapply()
	if err == nil {
		return r
	}
	return f(err)
}

func (r Try[T]) ToSeq[_ Phantom[T]]() []T {
	if r.success {
		return []T{r.v}
	}
	return nil
}

func (r Try[T]) Map[R any](mf func(T) R) Try[R] {
	if r.success {
		return Success(mf(r.v))
	}
	return Failure[R](r.err)
}

func (r Try[T]) FlatMap[R any](mf func(T) Try[R]) Try[R] {
	if r.success {
		return mf(r.v)
	}
	return Failure[R](r.err)
}

func (r Try[T]) Replace[R any](o R) Try[R] {
	_, err := r.Unapply()
	if err == nil {
		return Success(o)
	}
	return Failure[R](err)
}

func (r Try[T]) ReplaceS[R any](f func() R) Try[R] {
	return r.Map(func(t T) R {
		return f()
	})
}

func (r Try[T]) Void[R any]() Try[Unit] {
	return r.Replace(Unit{})
}

func (r Try[T]) Map2[U, R any](other Try[U], f func(T, U) R) Try[R] {
	return r.FlatMap(func(t T) Try[R] {
		return other.Map(func(u U) R {
			return f(t, u)
		})
	})
}

func (r Try[T]) IntoOption[_ Phantom[T]]() Option[T] {
	if r.success {
		return Some(r.v)
	}
	return None[T]()
}

func (r Try[T]) IntoFuture[_ Phantom[T]]() Future[T] {
	p := NewPromise[T]()
	if r.success {
		p.Success(r.v)
		return p.Future()
	}
	_, err := r.Unapply()
	p.Failure(err)
	return p.Future()
}

func (r Try[T]) TraverseF[R any](f func(T) Future[R]) Future[R] {
	if r.success {
		return f(r.v)
	}
	p := NewPromise[R]()
	_, err := r.Unapply()
	p.Failure(err)
	return p.Future()
}

func (r Future[T]) Map[R any](mf func(T) R, ctx ...Executor) Future[R] {
	np := NewPromise[R]()

	r.OnComplete(func(t Try[T]) {
		v, err := t.Unapply()
		if err == nil {
			np.Success(mf(v))
		} else {
			np.Failure(err)
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) FlatMap[R any](mf func(T) Future[R], ctx ...Executor) Future[R] {
	np := NewPromise[R]()

	r.OnComplete(func(t Try[T]) {
		v, err := t.Unapply()

		if err == nil {
			mf(v).OnComplete(func(t Try[R]) {
				np.Complete(t)
			}, ctx...)
		} else {
			np.Failure(err)
		}
	}, ctx...)

	return np.Future()
}

func (r Future[T]) Replace[R any](o R) Future[R] {
	return r.Map(Const[T](o))
}

func (r Future[T]) ReplaceS[R any](f func() R) Future[R] {
	return r.Map(func(t T) R {
		return f()
	})
}

func (r Future[T]) Void[R any]() Future[Unit] {
	return r.Replace(Unit{})
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

func (r Iterator[T]) All[_ Phantom[T]]() func(func(T) bool) {
	return func(f func(T) bool) {
		for r.hasNext() {
			if !f(r.Next()) {
				return
			}
		}
	}
}

func (r Iterator[T]) ToSeq[_ Phantom[T]]() []T {
	ret := []T{}
	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func (r Iterator[T]) Count[_ Phantom[T]]() int {
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

func (r Iterator[T]) MakeString[_ Phantom[T]](sep string) string {
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

func (r Iterator[T]) HasNext[_ Phantom[T]]() bool {
	if r.hasNext == nil {
		return false
	}

	return r.hasNext()
}

func (r Iterator[T]) Next[_ Phantom[T]]() T {
	return r.next()
}

func (r Iterator[T]) NextOption[_ Phantom[T]]() Option[T] {
	if r.HasNext() {
		v := r.next()
		return Some(v)
	}
	return None[T]()
}

func (r Iterator[T]) Take[_ Phantom[T]](n int) Iterator[T] {

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

func (r Iterator[T]) nextOnEmpty[_ Phantom[T]]() T {
	panic(ErrIteratorEmpty)
}

func (r Iterator[T]) TakeWhile[_ Phantom[T]](p func(T) bool) Iterator[T] {

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

func (r Iterator[T]) Drop[_ Phantom[T]](n int) Iterator[T] {

	for i := 0; i < n && r.HasNext(); i++ {
		r.Next()
	}

	return r
}

func (r Iterator[T]) DropWhile[_ Phantom[T]](p func(T) bool) Iterator[T] {

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

func (r Iterator[T]) Filter[_ Phantom[T]](p func(T) bool) Iterator[T] {

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

func (r Iterator[T]) FilterNot[_ Phantom[T]](p func(T) bool) Iterator[T] {
	return r.Filter(func(t T) bool {
		return !p(t)
	})
}

func (r Iterator[T]) Find[_ Phantom[T]](p func(T) bool) Option[T] {
	for r.HasNext() {
		v := r.Next()
		if p(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (r Iterator[T]) Foreach[_ Phantom[T]](p func(T)) {
	for r.HasNext() {
		v := r.Next()
		p(v)
	}
}

func (r Iterator[T]) TapEach[_ Phantom[T]](p func(T)) Iterator[T] {
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

func (r Iterator[T]) Appended[_ Phantom[T]](elem T) Iterator[T] {
	return r.Concat(IteratorOfSeq([]T{elem}))
}

func (r Iterator[T]) Concat[_ Phantom[T]](tail Iterator[T]) Iterator[T] {

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

func (r Iterator[T]) Map[R any](mf func(T) R) Iterator[R] {
	return MakeIterator(
		func() bool {
			return r.HasNext()
		},
		func() R {
			return mf(r.Next())
		},
	)
}

func (r Iterator[T]) FlatMap[R any](mf func(T) Iterator[R]) Iterator[R] {
	current := None[Iterator[R]]()

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
		func() R {
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

func (r Iterator[T]) Exists[_ Phantom[T]](p func(v T) bool) bool {
	for r.HasNext() {
		if p(r.Next()) {
			return true
		}
	}

	return false
}

func (r Iterator[T]) ForAll[_ Phantom[T]](p func(v T) bool) bool {
	for r.HasNext() {
		if !p(r.Next()) {
			return false
		}
	}
	return true
}

func (r Iterator[T]) IsEmpty[_ Phantom[T]]() bool {
	return !r.HasNext()
}

func (r Iterator[T]) NonEmpty[_ Phantom[T]]() bool {
	return r.HasNext()
}
