package fp

import "encoding/json"

// Option 의 zero value 는 None
// Try의 zero value 는 Failure()
// Either 의 zero value 는?  Left(zero)??
// zero value를 정의 할 수 없으면, struct 로 선언할 수 없음.
type Either[L, R any] interface {
	IsLeft() bool
	IsRight() bool
	Left() L
	Get() R
	Recover(f func() R) Either[L, R]
}

func Left[L, R any](l L) Either[L, R] {
	return left[L, R]{l}
}

func Right[L, R any](r R) Either[L, R] {
	return right[L, R]{r}
}

type left[L, R any] struct {
	v L
}

func (r left[L, R]) IsLeft() bool {
	return true
}
func (r left[L, R]) IsRight() bool {
	return false
}
func (r left[L, R]) Left() L {
	return r.v
}

func (r left[L, R]) Get() R {
	panic("Either.left")
}

func (r left[L, R]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.v)
}

func (r left[L, R]) Recover(f func() R) Either[L, R] {
	return Right[L, R](f())
}

type right[L, R any] struct {
	v R
}

func (r right[L, R]) IsLeft() bool {
	return false
}
func (r right[L, R]) IsRight() bool {
	return true
}
func (r right[L, R]) Left() L {
	panic("Either.right")
}

func (r right[L, R]) Get() R {
	return r.v
}

func (r right[L, R]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.v)
}

func (r right[L, R]) Recover(f func() R) Either[L, R] {
	return r
}
