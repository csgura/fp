package try

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)

type ApplicativeFunctor2[H hlist.Header[HT], HT, A1, A2, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[H, HT, A1, A2, R]) FlatMap(a func(HT) fp.Try[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Map(a func(HT) A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListMap(a func(H) A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListFlatMap(a func(H) fp.Try[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	nh := FlatMap(r.h, func(hv H) fp.Try[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concact(av, hv)
		})
	})

	return ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApTry(FromOption(a))
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Ap(a A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApTry(Success(a))

}
func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R] {
	return ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R]{Success(hlist.Empty()), Success(curried.Func2(fn))}
}

type ApplicativeFunctor3[H hlist.Header[HT], HT, A1, A2, A3, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) FlatMap(a func(HT) fp.Try[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Map(a func(HT) A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListMap(a func(H) A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListFlatMap(a func(H) fp.Try[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	nh := FlatMap(r.h, func(hv H) fp.Try[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concact(av, hv)
		})
	})

	return ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApTry(FromOption(a))
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Ap(a A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApTry(Success(a))

}
func Applicative3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) ApplicativeFunctor3[hlist.Nil, hlist.Nil, A1, A2, A3, R] {
	return ApplicativeFunctor3[hlist.Nil, hlist.Nil, A1, A2, A3, R]{Success(hlist.Empty()), Success(curried.Func3(fn))}
}

type ApplicativeFunctor4[H hlist.Header[HT], HT, A1, A2, A3, A4, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) FlatMap(a func(HT) fp.Try[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Map(a func(HT) A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListMap(a func(H) A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListFlatMap(a func(H) fp.Try[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	nh := FlatMap(r.h, func(hv H) fp.Try[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concact(av, hv)
		})
	})

	return ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApTry(FromOption(a))
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Ap(a A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApTry(Success(a))

}
func Applicative4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) ApplicativeFunctor4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R] {
	return ApplicativeFunctor4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R]{Success(hlist.Empty()), Success(curried.Func4(fn))}
}
