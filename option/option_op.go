//go:generate go run github.com/csgura/fp/internal/generator/option_gen
package option

import (
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/product"
)

func Some[T any](v T) fp.Option[T] {
	return fp.None[T]().Recover(func() T {
		return v
	})
}

func None[T any]() fp.Option[T] {
	return fp.None[T]()
}

func Of[T any](v T) fp.Option[T] {
	var i any = v
	if i == nil {
		return None[T]()
	}

	rv := reflect.ValueOf(i)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return None[T]()
	}
	return Some(v)
}

func Ptr[T any](v *T) fp.Option[T] {
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func Ap[T, U any](t fp.Option[fp.Func1[T, U]], a fp.Option[T]) fp.Option[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Option[U] {
		return Map(a, f)
	})
}

func Map[T, U any](opt fp.Option[T], f func(v T) U) fp.Option[U] {
	return FlatMap(opt, func(v T) fp.Option[U] {
		return Some(f(v))
	})
}

func Map2[A, B, U any](a fp.Option[A], b fp.Option[B], f func(A, B) U) fp.Option[U] {
	return FlatMap(a, func(v1 A) fp.Option[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

func Lift[T, U any](f func(v T) U) fp.Func1[fp.Option[T], fp.Option[U]] {
	return func(opt fp.Option[T]) fp.Option[U] {
		return Map(opt, f)
	}
}

func LiftA2[A1, A2, R any](f fp.Func2[A1, A2, R]) fp.Func2[fp.Option[A1], fp.Option[A2], fp.Option[R]] {
	return func(a1 fp.Option[A1], a2 fp.Option[A2]) fp.Option[R] {
		return Ap(Ap(Some(f.Curried()), a1), a2)
	}
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.Option[B]], f2 fp.Func1[B, fp.Option[C]]) fp.Func1[A, fp.Option[C]] {
	return func(a A) fp.Option[C] {
		return FlatMap(f1(a), f2)
	}
}

func Compose2[A, B, C any](f1 fp.Func1[A, fp.Option[B]], f2 fp.Func1[B, fp.Option[C]]) fp.Func1[A, fp.Option[C]] {
	return func(a A) fp.Option[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.Option[B]], f2 fp.Func1[B, C]) fp.Func1[A, fp.Option[C]] {
	return func(a A) fp.Option[C] {
		return Map(f1(a), f2)
	}
}

func FlatMap[T, U any](opt fp.Option[T], fn func(v T) fp.Option[U]) fp.Option[U] {
	if opt.IsDefined() {
		return fn(opt.Get())
	}
	return None[U]()
}

func Flatten[T any](opt fp.Option[fp.Option[T]]) fp.Option[T] {
	return FlatMap(opt, func(v fp.Option[T]) fp.Option[T] {
		return v
	})
}

func Zip[A, B any](c1 fp.Option[A], c2 fp.Option[B]) fp.Option[fp.Tuple2[A, B]] {
	return FlatMap(c1, func(v1 A) fp.Option[fp.Tuple2[A, B]] {
		return Map(c2, func(v2 B) fp.Tuple2[A, B] {
			return product.Tuple2(v1, v2)
		})
	})
}

func Zip3[A, B, C any](c1 fp.Option[A], c2 fp.Option[B], c3 fp.Option[C]) fp.Option[fp.Tuple3[A, B, C]] {
	return Applicative3(as.Tuple3[A, B, C]).
		ApOption(c1).
		ApOption(c2).
		ApOption(c3)
}

func SequenceIterator[T any](optItr fp.Iterator[fp.Option[T]]) fp.Option[fp.Iterator[T]] {

	return iterator.Fold(optItr, Some(iterator.Empty[T]()), func(list fp.Option[fp.Iterator[T]], v fp.Option[T]) fp.Option[fp.Iterator[T]] {
		return Map2(list, v, func(l fp.Iterator[T], e T) fp.Iterator[T] {
			return l.Concat(iterator.Of(e))
		})
	})
}

func Traverse[T, U any](itr fp.Iterator[T], fn func(T) fp.Option[U]) fp.Option[fp.Iterator[U]] {
	return iterator.Fold(itr, Some(iterator.Empty[U]()), func(tryItr fp.Option[fp.Iterator[U]], v T) fp.Option[fp.Iterator[U]] {
		return FlatMap(tryItr, func(acc fp.Iterator[U]) fp.Option[fp.Iterator[U]] {
			return Map(fn(v), func(v U) fp.Iterator[U] {
				return acc.Concat(iterator.Of(v))
			})
		})
	})
}

func TraverseSeq[T, U any](seq fp.Seq[T], fn func(T) fp.Option[U]) fp.Option[fp.Seq[U]] {
	return Map(Traverse(seq.Iterator(), fn), fp.Iterator[U].ToSeq)
}

func Sequence[T any](optSeq fp.Seq[fp.Option[T]]) fp.Option[fp.Seq[T]] {
	return Map(SequenceIterator(optSeq.Iterator()), fp.Iterator[T].ToSeq)
}

func Fold[A, B any](s fp.Option[A], zero B, f func(B, A) B) B {
	if s.IsEmpty() {
		return zero
	}

	return f(zero, s.Get())
}

func FoldRight[A, B any](s fp.Option[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	return f(s.Get(), lazy.Done(zero))
}

func ToSeq[T any](r fp.Option[T]) fp.Seq[T] {
	if r.IsDefined() {
		return fp.Seq[T]{r.Get()}
	}
	return nil
}

func Iterator[T any](r fp.Option[T]) fp.Iterator[T] {
	return ToSeq(r).Iterator()
}

type ApplicativeFunctor1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[H, HT, A, R]) Map(a func(HT) A) fp.Option[R] {
	return r.FlatMap(func(h HT) fp.Option[A] {
		return Some(a(h))
	})
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListMap(a func(H) A) fp.Option[R] {
	return r.HListFlatMap(func(h H) fp.Option[A] {
		return Some(a(h))
	})
}

func (r ApplicativeFunctor1[H, HT, A, R]) HListFlatMap(a func(H) fp.Option[A]) fp.Option[R] {
	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a(v)
	})

	return r.ApOption(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) FlatMap(a func(HT) fp.Option[A]) fp.Option[R] {
	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a(v.Head())
	})

	return r.ApOption(av)
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Option[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[H, HT, A, R]) Ap(a A) fp.Option[R] {
	return r.ApOption(Some(a))
}

func (r ApplicativeFunctor1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Option[R] {

	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a()
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor1[H, HT, A, R]) ApFunc(a func() A) fp.Option[R] {

	av := Map(r.h, func(v H) A {
		return a()
	})
	return r.ApOption(av)
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R] {
	return ApplicativeFunctor1[hlist.Nil, hlist.Nil, A, R]{Some(hlist.Empty()), Some(fn)}
}

// type ApplicativeFunctor2[H hlist.Header[HT], HT, A, B, R any] struct {
// 	h  fp.Option[H]
// 	fn fp.Option[fp.Func1[A, fp.Func1[B, R]]]
// }

// func (r ApplicativeFunctor2[H, HT, A, B, R]) FlatMap(a func(HT) fp.Option[A]) ApplicativeFunctor1[hlist.Cons[A, H], A, B, R] {

// 	av := FlatMap(r.h, func(v H) fp.Option[A] {
// 		return a(v.Head())
// 	})
// 	return r.ApOption(av)
// }

// func (r ApplicativeFunctor2[H, HT, A, B, R]) Map(a func(HT) A) ApplicativeFunctor1[hlist.Cons[A, H], A, B, R] {
// 	return r.FlatMap(func(h HT) fp.Option[A] {
// 		return Some(a(h))
// 	})
// }

// func (r ApplicativeFunctor2[H, HT, A, B, R]) HListMap(a func(H) A) ApplicativeFunctor1[hlist.Cons[A, H], A, B, R] {
// 	return r.HListFlatMap(func(h H) fp.Option[A] {
// 		return Some(a(h))
// 	})
// }

// func (r ApplicativeFunctor2[H, HT, A, B, R]) HListFlatMap(a func(H) fp.Option[A]) ApplicativeFunctor1[hlist.Cons[A, H], A, B, R] {
// 	av := FlatMap(r.h, func(v H) fp.Option[A] {
// 		return a(v)
// 	})

// 	return r.ApOption(av)
// }

// func (r ApplicativeFunctor2[H, HT, A, B, R]) ApOption(a fp.Option[A]) ApplicativeFunctor1[hlist.Cons[A, H], A, B, R] {
// 	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A, H]] {
// 		return Map(a, func(av A) hlist.Cons[A, H] {
// 			return hlist.Concat(av, hv)
// 		})
// 	})

// 	return ApplicativeFunctor1[hlist.Cons[A, H], A, B, R]{nh, Ap(r.fn, a)}
// }

// func (r ApplicativeFunctor2[H, HT, A, B, R]) Ap(a A) ApplicativeFunctor1[hlist.Cons[A, H], A, B, R] {
// 	return r.ApOption(Some(a))
// }

// func Applicative2[A, B, R any](fn fp.Func2[A, B, R]) ApplicativeFunctor2[hlist.Nil, hlist.Nil, A, B, R] {
// 	return ApplicativeFunctor2[hlist.Nil, hlist.Nil, A, B, R]{Some(hlist.Empty()), Some(curried.Func2(fn))}
// }

// type ApplicativeFunctor3[H hlist.Header[HT], HT, A, B, C, R any] struct {
// 	h  fp.Option[H]
// 	fn fp.Option[fp.Func1[A, fp.Func1[B, fp.Func1[C, R]]]]
// }

// func (r ApplicativeFunctor3[H, HT, A, B, C, R]) ApOption(a fp.Option[A]) ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R] {

// 	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A, H]] {
// 		return Map(a, func(av A) hlist.Cons[A, H] {
// 			return hlist.Concat(av, hv)
// 		})
// 	})

// 	return ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R]{nh, Ap(r.fn, a)}
// }

// func (r ApplicativeFunctor3[H, HT, A, B, C, R]) Ap(a A) ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R] {
// 	return r.ApOption(Some(a))
// }

// func (r ApplicativeFunctor3[H, HT, A, B, C, R]) FlatMap(a func(HT) fp.Option[A]) ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R] {

// 	av := FlatMap(r.h, func(v H) fp.Option[A] {
// 		return a(v.Head())
// 	})
// 	return r.ApOption(av)
// }

// func (r ApplicativeFunctor3[H, HT, A, B, C, R]) Map(a func(HT) A) ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R] {
// 	return r.FlatMap(func(h HT) fp.Option[A] {
// 		return Some(a(h))
// 	})
// }

// func (r ApplicativeFunctor3[H, HT, A, B, C, R]) HListMap(a func(H) A) ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R] {
// 	return r.HListFlatMap(func(h H) fp.Option[A] {
// 		return Some(a(h))
// 	})
// }

// func (r ApplicativeFunctor3[H, HT, A, B, C, R]) HListFlatMap(a func(H) fp.Option[A]) ApplicativeFunctor2[hlist.Cons[A, H], A, B, C, R] {
// 	av := FlatMap(r.h, func(v H) fp.Option[A] {
// 		return a(v)
// 	})

// 	return r.ApOption(av)
// }

// func Applicative3[A, B, C, R any](fn fp.Func3[A, B, C, R]) ApplicativeFunctor3[hlist.Nil, hlist.Nil, A, B, C, R] {
// 	return ApplicativeFunctor3[hlist.Nil, hlist.Nil, A, B, C, R]{Some(hlist.Empty()), Some(curried.Func3(fn))}
// }

// type ApplicativeFunctor4[H hlist.Header[HT], HT, A, B, C, D, R any] struct {
// 	h  fp.Option[H]
// 	fn fp.Option[fp.Func1[A, fp.Func1[B, fp.Func1[C, fp.Func1[D, R]]]]]
// }

// func (r ApplicativeFunctor4[H, HT, A, B, C, D, R]) ApOption(a fp.Option[A]) ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R] {

// 	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A, H]] {
// 		return Map(a, func(av A) hlist.Cons[A, H] {
// 			return hlist.Concat(av, hv)
// 		})
// 	})

// 	return ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R]{nh, Ap(r.fn, a)}
// }

// func (r ApplicativeFunctor4[H, HT, A, B, C, D, R]) Ap(a A) ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R] {
// 	return r.ApOption(Some(a))
// }

// func (r ApplicativeFunctor4[H, HT, A, B, C, D, R]) FlatMap(a func(HT) fp.Option[A]) ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R] {

// 	av := FlatMap(r.h, func(v H) fp.Option[A] {
// 		return a(v.Head())
// 	})
// 	return r.ApOption(av)
// }

// func (r ApplicativeFunctor4[H, HT, A, B, C, D, R]) Map(a func(HT) A) ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R] {
// 	return r.FlatMap(func(h HT) fp.Option[A] {
// 		return Some(a(h))
// 	})
// }

// func (r ApplicativeFunctor4[H, HT, A, B, C, D, R]) HListMap(a func(H) A) ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R] {
// 	return r.HListFlatMap(func(h H) fp.Option[A] {
// 		return Some(a(h))
// 	})
// }

// func (r ApplicativeFunctor4[H, HT, A, B, C, D, R]) HListFlatMap(a func(H) fp.Option[A]) ApplicativeFunctor3[hlist.Cons[A, H], A, B, C, D, R] {
// 	av := FlatMap(r.h, func(v H) fp.Option[A] {
// 		return a(v)
// 	})

// 	return r.ApOption(av)
// }

// func Applicative4[A, B, C, D, R any](fn fp.Func4[A, B, C, D, R]) ApplicativeFunctor4[hlist.Nil, hlist.Nil, A, B, C, D, R] {
// 	return ApplicativeFunctor4[hlist.Nil, hlist.Nil, A, B, C, D, R]{Some(hlist.Empty()), Some(curried.Func4(fn))}
// }
