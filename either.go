package fp

type Either[L, R any] interface {
	IsLeft() bool
	IsRight() bool
	Left() L
	Get() R
	Recover(f func() R) Either[L, R]
}
