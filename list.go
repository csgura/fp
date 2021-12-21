package fp

type List[T any] interface {
	IsEmpty() bool
	NonEmpty() bool
	Head() Option[T]
	Tail() List[T]
	Iterator() Iterator[T]
}

type ListAdaptor[T any] struct {
	GetHead func() Option[T]
	GetTail func() List[T]
}

func (r ListAdaptor[T]) IsEmpty() bool {
	return r.Head().IsEmpty()
}
func (r ListAdaptor[T]) NonEmpty() bool {
	return r.Head().IsDefined()
}
func (r ListAdaptor[T]) Head() Option[T] {
	return r.GetHead()
}

func (r ListAdaptor[T]) Tail() List[T] {
	return r.GetTail()
}

func (r ListAdaptor[T]) Iterator() Iterator[T] {
	var current List[T] = r

	return IteratorAdaptor[T]{
		IsHasNext: func() bool {
			return current.Head().IsDefined()
		},
		GetNext: func() T {
			ret := current.Head().Get()
			current = current.Tail()
			return ret
		},
	}
}
