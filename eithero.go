//go:build go1.27

package fp

type EitherO[L, R any] struct {
	left  *L
	right *R
}

func (r EitherO[L, R]) Map[T any](f func(R) T) EitherO[L, T] {
	if r.IsRight() {
		return EitherO[L, T]{
			right: new(f(*r.right)),
		}
	}
	return EitherO[L, T]{
		left: r.left,
	}
}

func (r EitherO[L, R]) FlatMap[T any](f func(R) EitherO[L, T]) EitherO[L, T] {
	if r.IsRight() {
		return f(*r.right)
	}
	return EitherO[L, T]{
		left: r.left,
	}
}

func (r EitherO[L, R]) Replace[T any](o T) EitherO[L, T] {
	return r.Map(Const[R](o))
}

func (r EitherO[L, R]) ReplaceS[T any](f func() T) EitherO[L, T] {
	return r.Map(func(r R) T {
		return f()
	})
}

func (r EitherO[L, R]) IsDefined[_ Phantom[R]]() bool {
	return r.IsLeft() || r.IsRight()
}

func (r EitherO[L, R]) IsEmpty[_ Phantom[R]]() bool {
	return !r.IsDefined()
}

func (r EitherO[L, R]) IsLeft[_ Phantom[R]]() bool {
	if r.left != nil {
		return true
	}
	return false
}

func (r EitherO[L, R]) IsRight[_ Phantom[R]]() bool {
	if r.right != nil {
		return true
	}
	return false
}

func (r EitherO[L, R]) Get[_ Phantom[R]]() R {
	if r.right == nil {
		panic("Either.left")
	}
	return *r.right
}

func (r EitherO[L, R]) Left[_ Phantom[R]]() L {
	if r.left == nil {
		panic("Either.right")
	}
	return *r.left
}

func (r EitherO[L, R]) Unapply() (*L, *R) {
	return r.left, r.right
}

func (r EitherO[L, R]) Reset(left *L, right *R) EitherO[L, R] {
	return EitherO[L, R]{
		left:  left,
		right: right,
	}
}

func (r EitherO[L, R]) Recover[_ Phantom[R]](f func() R) EitherO[L, R] {
	if r.right == nil {
		return EitherO[L, R]{
			right: new(f()),
		}
	}
	return r
}

func (r EitherO[L, R]) Swap[_ Phantom[L]]() EitherO[R, L] {
	return EitherO[R, L]{
		right: r.left,
		left:  r.right,
	}
}

func (r EitherO[L, R]) Fold[T any](lf func(L) T, rf func(R) T) Option[T] {
	if r.IsLeft() {
		return Some(lf(*r.left))
	}
	if r.IsRight() {
		return Some(rf(*r.right))
	}
	return None[T]()
}

func (r EitherO[L, R]) FoldT[T any](lf func(L) Try[T], rf func(R) Try[T]) Try[Option[T]] {
	if r.IsLeft() {
		return lf(*r.left).Map(Some)
	}
	if r.IsRight() {
		return rf(*r.right).Map(Some)
	}
	return Success(None[T]())
}

func (r EitherO[L, R]) FoldF[T any](lf func(L) Future[T], rf func(R) Future[T]) Future[Option[T]] {
	if r.IsLeft() {
		return lf(*r.left).Map(Some)
	}
	if r.IsRight() {
		return rf(*r.right).Map(Some)
	}
	p := NewPromise[Option[T]]()
	p.Success(None[T]())
	return p.Future()
}
