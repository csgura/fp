package fp

import (
	"fmt"
)

type Try[T any] struct {
	success bool
	v       T
	err     error
}

func Success[T any](t T) Try[T] {
	return Try[T]{true, t, nil}
}

func Failure[T any](err error) Try[T] {
	if err == nil {
		panic("Failure error is nil")
	}
	var zero T
	return Try[T]{false, zero, err}
}

func (r Try[T]) String() string {
	if r.success {
		return fmt.Sprintf("Success(%v)", r.v)
	}
	return fmt.Sprintf("Failure(%v)", r.err)
}
