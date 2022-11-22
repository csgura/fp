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

type Executor interface {
	ExecuteUnsafe(runnable Runnable)
}

type Promise[T any] interface {
	Future() Future[T]
	Success(value T) bool
	Failure(err error) bool
	IsCompleted() bool
	Complete(result Try[T]) bool
}

type Future[T any] interface {
	OnFailure(cb func(err error), ctx ...Executor)
	OnSuccess(cb func(success T), ctx ...Executor)
	Foreach(f func(v T), ctx ...Executor)
	OnComplete(cb func(try Try[T]), ctx ...Executor)
	IsCompleted() bool
	Failed() Future[error]
	Recover(f func(err error) T, ctx ...Executor) Future[T]
	RecoverWith(f func(err error) Future[T], ctx ...Executor) Future[T]
}
