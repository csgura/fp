//go:generate go run github.com/csgura/fp/internal/generator/fp_gen
package fp

import (
	"fmt"
	"reflect"

	"github.com/csgura/fp/hlist"
)

type Runnable interface {
	Run()
}

// RunnableFunc is converter which converts function to Runnable interface
type RunnableFunc func()

// Run is Runnable.Run
func (r RunnableFunc) Run() {
	r()
}

type ExecContext interface {
	Execute(runnable Runnable)
}

type Option[T any] interface {
	IsDefined() bool
	Get() T
	Foreach(f func(v T))
	Filter(p func(v T) bool) Option[T]
	OrElse(t T) T
	OrElseGet(func() T) T
	Recover(func() T) Option[T]
	Or(func() Option[T]) Option[T]
	String() string
}

type Try[T any] interface {
	IsSuccess() bool
	Get() T
	Foreach(f func(v T))
	Failed() Try[error]
	OrElse(t T) T
	OrElseGet(func() T) T
	Or(func() Try[T]) Try[T]
	Recover(func(err error) T) Try[T]
	RecoverWith(func(err error) Try[T]) Try[T]
	ToOption() Option[T]
	Unapply() (T, error)
	String() string
}

type Promise[T any] interface {
	Future() Future[T]
	Success(value T) bool
	Failure(err error) bool
	IsCompleted() bool
	Complete(result Try[T]) bool
}

type Future[T any] interface {
	OnFailure(cb func(err error), ctx ...ExecContext)
	OnSuccess(cb func(success T), ctx ...ExecContext)
	Foreach(f func(v T), ctx ...ExecContext)
	OnComplete(cb func(try Try[T]), ctx ...ExecContext)
	IsCompleted() bool
	Value() Option[Try[T]]
	Failed() Future[error]
	Recover(f func(err error) T, ctx ...ExecContext) Future[T]
	RecoverWith(f func(err error) Future[T], ctx ...ExecContext) Future[T]
}
type Unit struct {
}

func (r Unit) String() string {
	return "()"
}

type Tuple1[T1 any] struct {
	I1 T1
}

func (r Tuple1[T1]) Head() T1 {
	return r.I1
}

func (r Tuple1[T1]) Tail() Unit {
	return Unit{}
}

func (r Tuple1[T1]) ToHList() hlist.Cons[T1, hlist.Nil] {
	return hlist.Concact(r.Head(), hlist.Empty())
}

type Func0[R any] func() R

func Println[T any](v T) {
	fmt.Println(v)
}

func ToString[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

func TypeName[T any]() string {
	var zero *T
	return reflect.TypeOf(zero).Elem().String()
}

type Seq[T any] []T

func (r Seq[T]) Size() int {
	return len(r)
}

func (r Seq[T]) IsEmpty() bool {
	return r.Size() == 0
}

func (r Seq[T]) NonEmpty() bool {
	return r.Size() > 0
}

func (r Seq[T]) Head() Option[T] {
	if r.Size() > 0 {
		return Some[T]{r[0]}
	} else {
		return None[T]{}
	}
}

func (r Seq[T]) Tail() Seq[T] {
	if r.Size() > 0 {
		return r[1:]
	} else {
		return nil
	}
}

func (r Seq[T]) Take(n int) Seq[T] {
	if r.Size() > n {
		return r[0:n]
	} else {
		return r
	}
}

func (r Seq[T]) Drop(n int) Seq[T] {
	if r.Size() > n {
		return r[r.Size()-n : r.Size()]
	} else {
		return nil
	}
}

func (r Seq[T]) Foreach(f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) Filter(p func(v T) bool) Seq[T] {
	ret := make(Seq[T], 0, r.Size())
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
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
			return Some[T]{v}
		}
	}
	return None[T]{}
}

func (r Seq[T]) Append(items ...T) Seq[T] {
	tail := Seq[T](items)
	ret := make(Seq[T], r.Size()+tail.Size())

	for i := range r {
		ret[i] = r[i]
	}

	for i := range tail {
		ret[i+r.Size()] = tail[i]
	}

	return ret
}

func (r Seq[T]) Concact(tail Seq[T]) Seq[T] {
	ret := make(Seq[T], r.Size()+tail.Size())

	for i := range r {
		ret[i] = r[i]
	}

	for i := range tail {
		ret[i+r.Size()] = tail[i]
	}

	return ret
}
