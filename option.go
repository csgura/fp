package fp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ref[T any] struct {
	ref *T
}

var ErrInvalidRef = Error(http.StatusBadRequest, "fp.Ref not initialized correctly")

func (r Ref[T]) Get() T {
	if r.ref == nil {
		panic(ErrInvalidRef)
	}
	return *r.ref
}

func RefOf[T any](v *T) Ref[T] {
	return Ref[T]{
		ref: v,
	}
}

type Ptr[T any] = *T

type PtrT[T any] = Try[Ptr[T]]

type OptionT[T any] = Try[Option[T]]

type Option[T any] struct {
	present bool
	v       T
}

type GoIter[V any] = func(yield func(V) bool)

func Some[T any](v T) Option[T] {
	return Option[T]{true, v}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (r *Option[T]) UnmarshalJSON(b []byte) error {
	if r == nil {
		return Error(http.StatusBadRequest, "target ptr is nil")
	}
	if len(b) > 0 {
		if b[0] != 'n' {
			var t T
			err := json.Unmarshal(b, &t)
			if err == nil {
				*r = Some(t)
			}
			return err
		}
	}
	*r = None[T]()

	return nil
}

func (r Option[T]) MarshalJSON() ([]byte, error) {
	if r.present {
		return json.Marshal(r.v)
	}

	return []byte("null"), nil
}

func (r Option[T]) Unapply() (T, bool) {
	if r.present {
		return r.v, true
	} else {
		var zero T
		return zero, false
	}
}

func (r Option[T]) Recover(f func() T) Option[T] {
	if r.present {
		return r
	}
	t := f()
	return Option[T]{true, t}
}

func (r Option[T]) String() string {
	if r.present {
		return fmt.Sprintf("Some(%v)", r.v)
	} else {
		return "None"
	}
}

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

func (r Option[T]) Void() Option[Unit] {
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

func (r Option[T]) All() GoIter[T] {
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

func (r Option[T]) IsDefined() bool {
	return r.present
}

func (r Option[T]) IsEmpty[_ Phantom[T]]() bool {
	return !r.IsDefined()
}

func (r Option[T]) Get() T {
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
func (r Option[T]) OrElse(t T) T {
	if r.IsDefined() {
		return r.Get()
	}
	return t
}

func (r Option[T]) OrZero() T {
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

func (r Option[T]) Fold[A any](zero A, f func(A, T) A) A {
	if r.present {
		return f(zero, r.v)
	}
	return zero
}

func (r Option[T]) Either[R any](noncase func() R, somecase func(T) R) R {
	if r.present {
		return somecase(r.v)
	}
	return noncase()
}
