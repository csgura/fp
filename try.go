package fp

type Try[T any] interface {
	IsSuccess() bool
	IsFailure() bool
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
