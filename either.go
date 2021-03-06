package fp

type Either[L, R any] interface {
	IsLeft() bool
	IsRight() bool
	Left() Option[L]
	Right() Option[R]
	Swap() Either[R, L]
	Get() R
	Foreach(f func(v R))
	OrElse(t R) R
	OrElseGet(func() R) R
	Exists(p func(v R) bool) bool
	ForAll(p func(v R) bool) bool
	// ToSeq() Seq[T]

}
