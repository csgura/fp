package list

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

type Nil[T any] struct {
}

func (r Nil[T]) IsEmpty() bool {
	return true
}

func (r Nil[T]) NonEmpty() bool {
	return false
}

func (r Nil[T]) Head() fp.Option[T] {
	return option.None[T]()
}

func (r Nil[T]) Tail() fp.List[T] {
	return r
}

func (r Nil[T]) Unapply() (fp.Option[T], fp.List[T]) {
	return r.Head(), r
}

func (r Nil[T]) Iterator() fp.Iterator[T] {
	return fp.MakeIterator(func() bool {
		return false
	}, func() T {
		panic("next on empty iterator")
	})
}

type Cons[T any] struct {
	head T
	tail fp.List[T]
}

func (r Cons[T]) IsEmpty() bool {
	return false
}

func (r Cons[T]) NonEmpty() bool {
	return true
}

func (r Cons[T]) Head() fp.Option[T] {
	return option.Some(r.head)
}

func (r Cons[T]) Tail() fp.List[T] {
	return r.tail
}

func (r Cons[T]) Unapply() (fp.Option[T], fp.List[T]) {
	return r.Head(), r.Tail()
}

func (r Cons[T]) Iterator() fp.Iterator[T] {
	var cursor fp.List[T] = r

	hasNext := func() bool {
		return cursor.NonEmpty()
	}

	return fp.MakeIterator(hasNext, func() T {
		if hasNext() {
			ret := cursor.Head().Get()
			cursor = r.Tail()
			return ret
		}
		panic("next on empty iterator")
	})
}

func Empty[T any]() fp.List[T] {
	return Nil[T]{}
}

func Generate[T any](generator func(index int) T) fp.List[T] {
	return GenerateFrom(0, generator)
}

func GenerateFrom[T any](startIndex int, generator func(index int) T) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(generator(startIndex))
		},
		func() fp.List[T] {
			return GenerateFrom(startIndex+1, generator)
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

func Map2[T, U any](a, b fp.List[T], f func(T, T) U) fp.List[U] {
	return FlatMap(a, func(v1 T) fp.List[U] {
		return Map(b, func(v2 T) U {
			return f(v1, v2)
		})
	})
}

func FlatMap[T, U any](opt fp.List[T], fn func(v T) fp.List[U]) fp.List[U] {

	if opt.IsEmpty() {
		return Empty[U]()
	}

	mappedHeadLazy := lazy.Call(func() fp.List[U] {
		return fn(opt.Head().Get())
	})

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

			return Combine(headList.Tail(), FlatMap(tail, fn))
		},
	)

}

func Apply[T any](head T, tail fp.List[T]) fp.List[T] {
	return Cons[T]{head, tail}
}

func Of[T any](e ...T) fp.List[T] {
	return seq.Of(e...)
}

func Concat[T any](head T, tail fp.List[T]) fp.List[T] {
	return Apply(head, tail)
}

func Combine[T any](l1 fp.List[T], l2 fp.List[T]) fp.List[T] {

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
				return Combine(l1Tail, l2)
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

func Zip3[A, B, C any](a fp.List[A], b fp.List[B], c fp.List[C]) fp.List[fp.Tuple3[A, B, C]] {
	return fp.MakeList(
		func() fp.Option[fp.Tuple3[A, B, C]] {
			return option.Applicative3(as.Tuple3[A, B, C]).
				ApOption(a.Head()).
				ApOption(b.Head()).
				ApOption(c.Head())
		},
		func() fp.List[fp.Tuple3[A, B, C]] {
			return Zip3(a.Tail(), b.Tail(), c.Tail())
		},
	)
}

func Reduce[A any](s fp.List[A], m fp.Monoid[A]) A {
	return FoldRight(s, m.Empty(), func(a A, b lazy.Eval[A]) lazy.Eval[A] {
		return b.Map(func(v A) A {
			return m.Combine(a, v)
		})
	}).Get()
}

func Fold[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
	sum := zero

	cursor := s

	for !cursor.IsEmpty() {
		sum = f(sum, cursor.Head().Get())
		cursor = cursor.Tail()
	}
	return sum
}

func FoldLeft[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
	cf := as.Func2(f).Shift().Curried()
	ret := FoldRight[A, fp.Endo[B]](s, fp.Id[B], func(a A, endo lazy.Eval[fp.Endo[B]]) lazy.Eval[fp.Endo[B]] {
		ef := endo.Get().AsFunc()
		return lazy.Done(fp.Endo[B](fp.Compose(cf(a), ef)))
	})
	return ret.Get()(zero)
}

func FoldMap[A, B any](s fp.List[A], f func(A) B, m fp.Monoid[B]) B {
	ret := FoldRight(s, m.Empty(), func(a A, b lazy.Eval[B]) lazy.Eval[B] {
		ab := f(a)

		return b.Map(func(t B) B {
			return m.Combine(ab, t)
		})

	})

	return ret.Get()
}

func FoldRight[A, B any](s fp.List[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	v := lazy.TailCall(func() lazy.Eval[B] {
		return FoldRight(s.Tail(), zero, f)
	})
	return f(s.Head().Get(), v)
}

func FoldLeftUsingMap[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
	cf := as.Func2(f).Shift().Curried()
	m := monoid.Dual(monoid.Endo[B]())

	f2 := func(a A) fp.Dual[fp.Endo[B]] {
		return as.Dual(as.Endo(cf(a)))
	}

	ret := FoldMap(s, f2, m)
	return ret.GetDual(zero)
}

func FoldRightUsingMap[A, B any](s fp.List[A], zero B, f func(A, B) B) B {
	cf := as.Func2(f).Curried()
	m := monoid.Endo[B]()

	f2 := func(a A) fp.Endo[B] {
		return as.Endo(cf(a))
	}

	ret := FoldMap(s, f2, m)
	return ret(zero)
}

func Scan[A, B any](s fp.List[A], zero B, f func(B, A) B) fp.List[B] {

	return fp.MakeList(
		func() fp.Option[B] {
			return option.Some(zero)
		},
		func() fp.List[B] {
			z := option.Map(s.Head(), as.Curried2(f)(zero))
			if z.IsDefined() {
				return Scan(s.Tail(), z.Get(), f)
			}
			return Empty[B]()
		},
	)
}
