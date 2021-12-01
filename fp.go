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
