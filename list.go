package fp

type List[T any] interface {
	IsEmpty() bool
	NonEmpty() bool
	Head() Option[T]
	Tail() List[T]
	Unapply() (Option[T], List[T])
	Iterator() Iterator[T]
	Foreach(f func(v T))
}

type ListAdaptor[T any] struct {
	getHead Func0[Option[T]]
	getTail Func0[List[T]]
}

func (r ListAdaptor[T]) IsEmpty() bool {
	return r.Head().IsEmpty()
}
func (r ListAdaptor[T]) NonEmpty() bool {
	return r.Head().IsDefined()
}
func (r ListAdaptor[T]) Head() Option[T] {
	return r.getHead.Apply()
}

func (r ListAdaptor[T]) Tail() List[T] {
	return r.getTail.Apply()
}

func (r ListAdaptor[T]) Unapply() (Option[T], List[T]) {
	return r.Head(), r.Tail()
}

func (r ListAdaptor[T]) Foreach(f func(v T)) {
	var cursor List[T] = r
	for cursor.NonEmpty() {
		f(cursor.Head().Get())
		cursor = cursor.Tail()
	}
}

func (r ListAdaptor[T]) Iterator() Iterator[T] {
	var current List[T] = r

	return MakeIterator(
		func() bool {
			return current.Head().IsDefined()
		},
		func() T {
			ret := current.Head().Get()
			current = current.Tail()
			return ret
		},
	)
}

func MakeList[T any](head func() Option[T], tail func() List[T]) List[T] {
	return ListAdaptor[T]{Memoize(head), Memoize(tail)}
}
