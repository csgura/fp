package future

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)

type ApplicativeFunctor2[H hlist.Header[HT], HT, A1, A2, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Flip() ApplicativeFunctor2[H, HT, A2, A1, R] {

	return ApplicativeFunctor2[H, HT, A2, A1, R]{
		r.h,
		Map(r.fn, curried.Flip[A1, A2, R]),
	}

}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Ap(a A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R] {
	return ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R]{Successful(hlist.Empty()), Successful(curried.Func2(fn))}
}

type ApplicativeFunctor3[H hlist.Header[HT], HT, A1, A2, A3, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Ap(a A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) ApplicativeFunctor3[hlist.Nil, hlist.Nil, A1, A2, A3, R] {
	return ApplicativeFunctor3[hlist.Nil, hlist.Nil, A1, A2, A3, R]{Successful(hlist.Empty()), Successful(curried.Func3(fn))}
}

type ApplicativeFunctor4[H hlist.Header[HT], HT, A1, A2, A3, A4, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Ap(a A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) ApplicativeFunctor4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R] {
	return ApplicativeFunctor4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R]{Successful(hlist.Empty()), Successful(curried.Func4(fn))}
}

type ApplicativeFunctor5[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) Ap(a A1) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) ApplicativeFunctor5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R] {
	return ApplicativeFunctor5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R]{Successful(hlist.Empty()), Successful(curried.Func5(fn))}
}

type ApplicativeFunctor6[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) Ap(a A1) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) ApplicativeFunctor6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R]{Successful(hlist.Empty()), Successful(curried.Func6(fn))}
}

type ApplicativeFunctor7[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplicativeFunctor7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R]{Successful(hlist.Empty()), Successful(curried.Func7(fn))}
}

type ApplicativeFunctor8[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplicativeFunctor8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R]{Successful(hlist.Empty()), Successful(curried.Func8(fn))}
}

type ApplicativeFunctor9[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Map(a func(HT) A1, ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListMap(a func(H) A1, ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplicativeFunctor9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{Successful(hlist.Empty()), Successful(curried.Func9(fn))}
}
