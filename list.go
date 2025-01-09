package fp

type Cons[H, T any] interface {
	Head() H
	Tail() T
}

type List[T any] interface {
	IsEmpty() bool
	NonEmpty() bool
	Head() T
	Tail() List[T]
	Unapply() (T, List[T])
	Foreach(f func(v T))
	ToSeq() []T
}

type ListAdaptor[T any] struct {
	getHead Func0[Option[T]]
	getTail Func0[List[T]]
}

func (r ListAdaptor[T]) IsEmpty() bool {
	if r.getHead == nil {
		return true
	}
	return r.getHead(Unit{}).IsEmpty()
}
func (r ListAdaptor[T]) NonEmpty() bool {
	return !r.IsEmpty()
}
func (r ListAdaptor[T]) Head() T {
	if r.getHead == nil {
		panic("List.empty")
	}
	opt := r.getHead(Unit{})
	if opt.IsEmpty() {
		panic("List.empty")
	}
	return opt.Get()
}

func (r ListAdaptor[T]) Tail() List[T] {
	if r.getTail == nil {
		return MakeList(func() Option[T] {
			return None[T]()
		}, func() List[T] {
			return r.Tail()
		})
	}
	return r.getTail(Unit{})
}

func (r ListAdaptor[T]) Unapply() (T, List[T]) {
	return r.Head(), r.Tail()
}

func (r ListAdaptor[T]) Foreach(f func(v T)) {
	var cursor List[T] = r
	for cursor.NonEmpty() {
		f(cursor.Head())
		cursor = cursor.Tail()
	}
}

func (r ListAdaptor[T]) ToSeq() []T {
	ret := []T{}
	r.Foreach(func(v T) {
		ret = append(ret, v)
	})
	return ret
}

// func (r ListAdaptor[T]) Iterator() Iterator[T] {
// 	var current List[T] = r

// 	return MakeIterator(
// 		func() bool {
// 			return current.NonEmpty()
// 		},
// 		func() T {
// 			ret := current.Head()
// 			current = current.Tail()
// 			return ret
// 		},
// 	)
// }

func MakeList[T any](head func() Option[T], tail func() List[T]) List[T] {
	return ListAdaptor[T]{Memoize(head), Memoize(tail)}
}
