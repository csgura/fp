//go:generate go run github.com/csgura/fp/internal/generator/fp_gen
package fp

import "github.com/csgura/fp/hlist"

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
}

type Try[T any] interface {
	Get() T
	IsSuccess() bool
	Failed() Try[error]
	Foreach(f func(v T))
}

type Future[T any] interface {
	OnFailure(cb func(err error), ctx ...ExecContext)
	OnSuccess(cb func(success T), ctx ...ExecContext)
	Foreach(f func(v T), ctx ...ExecContext)
	OnComplete(cb func(try Try[T]), ctx ...ExecContext)
	IsCompleted() bool
	Value() Try[T]
}
type Unit struct {
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
