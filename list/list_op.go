package list

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func Generate[T any](generator func(index int) T) fp.List[T] {
	return GenerateFrom(generator, 0)
}

func GenerateFrom[T any](generator func(index int) T, startIndex int) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(generator(startIndex))
		},
		func() fp.List[T] {
			return GenerateFrom(generator, startIndex+1)
		},
	)
}

func Recurrence1[T any](a1 T, relation func(n1 T) T) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(a1)
		},
		func() fp.List[T] {
			return Recurrence1(relation(a1), relation)
		},
	)
}

func Recurrence2[T any](a1 T, a2 T, relation func(n1, n2 T) T) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(a1)
		},
		func() fp.List[T] {
			return Recurrence2(a2, relation(a1, a2), relation)
		},
	)
}

func Map[T, U any](opt fp.List[T], fn func(v T) U) fp.List[U] {
	return fp.MakeList(
		func() fp.Option[U] {
			return option.Map(opt.Head(), fn)
		},
		func() fp.List[U] {
			return Map(opt.Tail(), fn)
		},
	)
}

func FlatMap[T, U any](opt fp.List[T], fn func(v T) fp.List[U]) fp.List[U] {

	if opt.IsEmpty() {
		return Of[U]()
	}

	mappedHeadLazy := lazy.Func0(fp.Compose(as.Func0(opt.Head().Get), fn))

	tail := opt.Tail()

	return fp.MakeList(
		func() fp.Option[U] {
			headList := mappedHeadLazy.Get()

			if headList.IsEmpty() {
				return FlatMap(tail, fn).Head()
			}

			return headList.Head()
		},
		func() fp.List[U] {
			headList := mappedHeadLazy.Get()

			if headList.IsEmpty() {
				return FlatMap(tail, fn).Tail()
			}

			return Concat(headList.Tail(), FlatMap(tail, fn))
		},
	)

}

func Apply[T any](head T, tail fp.List[T]) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(head)
		},
		func() fp.List[T] {
			return tail
		},
	)
}

func Of[T any](e ...T) fp.List[T] {
	return seq.Of(e...)
}

func Concat[T any](l1 fp.List[T], l2 fp.List[T]) fp.List[T] {

	if l1.IsEmpty() {
		return l2
	}

	return fp.MakeList(
		func() fp.Option[T] {
			return l1.Head()
		},
		func() fp.List[T] {
			l1Tail := l1.Tail()
			if l1Tail.NonEmpty() {
				return Concat(l1Tail, l2)
			}
			return l2
		},
	)
}

func Ap[T, U any](t fp.List[fp.Func1[T, U]], a fp.List[T]) fp.List[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.List[U] {
		return Map(a, f)
	})
}

func Lift[T, U any](f func(v T) U) fp.Func1[fp.List[T], fp.List[U]] {
	return func(opt fp.List[T]) fp.List[U] {
		return Map(opt, f)
	}
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.List[B]], f2 fp.Func1[B, fp.List[C]]) fp.Func1[A, fp.List[C]] {
	return func(a A) fp.List[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.List[B]], f2 fp.Func1[B, C]) fp.Func1[A, fp.List[C]] {
	return func(a A) fp.List[C] {
		return Map(f1(a), f2)
	}
}

func Flatten[T any](opt fp.List[fp.List[T]]) fp.List[T] {
	return FlatMap(opt, func(v fp.List[T]) fp.List[T] {
		return v
	})
}

func Zip[T, U any](a fp.List[T], b fp.List[U]) fp.List[fp.Tuple2[T, U]] {
	return fp.MakeList(
		func() fp.Option[fp.Tuple2[T, U]] {
			return option.Applicative2(as.Tuple[T, U]).
				ApOption(a.Head()).
				ApOption(b.Head())
		},
		func() fp.List[fp.Tuple2[T, U]] {
			return Zip(a.Tail(), b.Tail())
		},
	)
}

// func Fold[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
// 	sum := zero
// 	for s.HasNext() {
// 		sum = f(sum, s.Next())
// 	}
// 	return sum
// }

func FoldRight[A, B any](s fp.List[A], zero B, f func(A, fp.Lazy[B]) B) B {
	if s.IsEmpty() {
		return zero
	}

	v := lazy.Eval(func() B {
		return FoldRight(s.Tail(), zero, f)
	})
	return f(s.Head().Get(), v)
}

func Scan[A, B any](s fp.List[A], zero B, f func(B, A) B) fp.List[B] {

	c := lazy.Eval(func() fp.Option[B] {
		return option.Map(s.Head(), as.Curried2(f)(zero))
	})

	return fp.MakeList(
		func() fp.Option[B] {
			return option.Some(zero)
		},
		func() fp.List[B] {
			z := c.Get()
			if z.IsDefined() {
				return Scan(s.Tail(), z.Get(), f)
			}
			return Of[B]()
		},
	)
}
