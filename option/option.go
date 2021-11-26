//go:generate go run github.com/csgura/fp/internal/generator/option_gen
package option

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
)

type option[T any] struct {
	v *T
}

func (r option[T]) Foreach(f func(v T)) {
	if r.v != nil {
		f(*r.v)
	}
}

func (r option[T]) IsDefined() bool {
	return r.v != nil
}

func (r option[T]) Get() T {
	return *r.v
}

func (r option[T]) Filter(p func(v T) bool) fp.Option[T] {
	return FlatMap[T](r, func(v T) fp.Option[T] {
		if p(v) {
			return r
		}
		return None[T]()
	})
}

func Some[T any](v T) fp.Option[T] {
	return &option[T]{&v}
}

func None[T any]() fp.Option[T] {
	return &option[T]{}
}

func Ap[T, U any](t fp.Option[fp.Func1[T, U]], a fp.Option[T]) fp.Option[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Option[U] {
		return Map(a.(fp.Option[T]), func(a T) U {
			return f(a)
		})
	})
}

func Map[T, U any](opt fp.Option[T], f func(v T) U) fp.Option[U] {
	return FlatMap(opt, func(v T) fp.Option[U] {
		return Some(f(v))
	})
}

func FlatMap[T, U any](opt fp.Option[T], fn func(v T) fp.Option[U]) fp.Option[U] {
	if opt.IsDefined() {
		return fn(opt.Get()).(fp.Option[U])
	}
	return None[U]()
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
// 			return hlist.Concact(av, hv)
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
// 			return hlist.Concact(av, hv)
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
// 			return hlist.Concact(av, hv)
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
