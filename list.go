package fp

import "sync"

type List[T any] interface {
	IsEmpty() bool
	NonEmpty() bool
	Head() Option[T]
	Tail() List[T]
	Unapply() (Option[T], List[T])
	Iterator() Iterator[T]
}

type ListAdaptor[T any] struct {
	GetHead Lazy[Option[T]]
	GetTail Lazy[List[T]]
}

func (r ListAdaptor[T]) IsEmpty() bool {
	return r.Head().IsEmpty()
}
func (r ListAdaptor[T]) NonEmpty() bool {
	return r.Head().IsDefined()
}
func (r ListAdaptor[T]) Head() Option[T] {
	return r.GetHead.Apply()
}

func (r ListAdaptor[T]) Tail() List[T] {
	return r.GetTail.Apply()
}

func (r ListAdaptor[T]) Unapply() (Option[T], List[T]) {
	return r.Head(), r.Tail()
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

func MakeList[T any](head func() Option[T], tail func() List[T]) List[T] {
	return ListAdaptor[T]{LazyFunc(head), LazyFunc(tail)}
}

type Lazy[T any] Func0[T]

func (r Lazy[T]) Apply() T {
	return r(Unit{})
}

func (r Lazy[T]) Get() T {
	return r(Unit{})
}

func LazyFunc[T any](f func() T) Lazy[T] {
	once := sync.Once{}
	var ret T
	return func(Unit) T {
		once.Do(func() {
			ret = f()
		})
		return ret
	}
}
