package fp

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

type Func1[A, R any] func(a A) R

type Func2[A, B, R any] func(a A, b B) R

func (r Func2[A, B, R]) Tupled() func(t Tuple2[A, B]) R {
	return func(t Tuple2[A, B]) R {
		return r(t.I1, t.I2)
	}
}

type Tuple1[A any] struct {
	I1 A
}

type Tuple2[A, B any] struct {
	I1 A
	I2 B
}

func (r Tuple2[A, B]) Head() A {
	return r.I1
}

func (r Tuple2[A, B]) Tail() B {
	return r.I2
}

type Tuple3[A, B, C any] struct {
	I1 A
	I2 B
	I3 C
}

func (r Tuple3[A, B, C]) Head() A {
	return r.I1
}

func (r Tuple3[A, B, C]) Tail() Tuple2[B, C] {
	return Tuple2[B, C]{
		I1: r.I2,
		I2: r.I3,
	}
}

type Tuple4[A, B, C, D any] struct {
	I1 A
	I2 B
	I3 C
	I4 D
}

func (r Tuple4[A, B, C, D]) Head() A {
	return r.I1
}

func (r Tuple4[A, B, C, D]) Tail() Tuple3[B, C, D] {
	return Tuple3[B, C, D]{
		I1: r.I2,
		I2: r.I3,
		I3: r.I4,
	}
}

type Unit struct {
}

type Func3[A, B, C, R any] func(a A, b B, c C) R
