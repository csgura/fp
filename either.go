package fp

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
