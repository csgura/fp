package future

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)

type MonadChain2[H hlist.Header[HT], HT, A1, A2, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r MonadChain2[H, HT, A1, A2, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain2[H, HT, A1, A2, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain2[H, HT, A1, A2, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApFuture(a fp.Future[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain1[hlist.Cons[A1, H], A1, A2, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain2[H, HT, A1, A2, R]) ApTry(a fp.Try[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain2[H, HT, A1, A2, R]) ApOption(a fp.Option[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain2[H, HT, A1, A2, R]) Ap(a A1) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain2[H, HT, A1, A2, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain2[A1, A2, R any](fn fp.Func2[A1, A2, R]) MonadChain2[hlist.Nil, hlist.Nil, A1, A2, R] {
	return MonadChain2[hlist.Nil, hlist.Nil, A1, A2, R]{Successful(hlist.Empty()), Successful(curried.Func2(fn))}
}

type MonadChain3[H hlist.Header[HT], HT, A1, A2, A3, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r MonadChain3[H, HT, A1, A2, A3, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApFuture(a fp.Future[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApTry(a fp.Try[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApOption(a fp.Option[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) Ap(a A1) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) MonadChain3[hlist.Nil, hlist.Nil, A1, A2, A3, R] {
	return MonadChain3[hlist.Nil, hlist.Nil, A1, A2, A3, R]{Successful(hlist.Empty()), Successful(curried.Func3(fn))}
}

type MonadChain4[H hlist.Header[HT], HT, A1, A2, A3, A4, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApFuture(a fp.Future[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApTry(a fp.Try[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) Ap(a A1) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) MonadChain4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R] {
	return MonadChain4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R]{Successful(hlist.Empty()), Successful(curried.Func4(fn))}
}

type MonadChain5[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApFuture(a fp.Future[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApTry(a fp.Try[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) Ap(a A1) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) MonadChain5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R] {
	return MonadChain5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R]{Successful(hlist.Empty()), Successful(curried.Func5(fn))}
}

type MonadChain6[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFuture(a fp.Future[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApTry(a fp.Try[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) Ap(a A1) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) MonadChain6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R] {
	return MonadChain6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R]{Successful(hlist.Empty()), Successful(curried.Func6(fn))}
}

type MonadChain7[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFuture(a fp.Future[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApTry(a fp.Try[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) MonadChain7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R] {
	return MonadChain7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R]{Successful(hlist.Empty()), Successful(curried.Func7(fn))}
}

type MonadChain8[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFuture(a fp.Future[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTry(a fp.Try[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) MonadChain8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return MonadChain8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R]{Successful(hlist.Empty()), Successful(curried.Func8(fn))}
}

type MonadChain9[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	h  fp.Future[H]
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Map(a func(HT) A1, ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListMap(a func(H) A1, ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFuture(a fp.Future[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTry(a fp.Try[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(FromTry(a))
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(FromOption(a))
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(Successful(a))

}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFunc(a func() A1, ctx ...fp.Executor) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Chain9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) MonadChain9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return MonadChain9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{Successful(hlist.Empty()), Successful(curried.Func9(fn))}
}

type ApplicativeFunctor2[A1, A2, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[A1, A2, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor1[A2, R] {

	return ApplicativeFunctor1[A2, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor2[A1, A2, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor1[A2, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor2[A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[A2, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor2[A1, A2, R]) Ap(a A1) ApplicativeFunctor1[A2, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor2[A1, A2, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor1[A2, R] {

	return ApplicativeFunctor1[A2, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor2[A1, A2, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor1[A2, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor2[A1, A2, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor1[A2, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor2[A1, A2, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor1[A2, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[A1, A2, R] {
	return ApplicativeFunctor2[A1, A2, R]{Successful(curried.Func2(fn))}
}

type ApplicativeFunctor3[A1, A2, A3, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor2[A2, A3, R] {

	return ApplicativeFunctor2[A2, A3, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor3[A1, A2, A3, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor2[A2, A3, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor3[A1, A2, A3, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor2[A2, A3, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor3[A1, A2, A3, R]) Ap(a A1) ApplicativeFunctor2[A2, A3, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor3[A1, A2, A3, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor2[A2, A3, R] {

	return ApplicativeFunctor2[A2, A3, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor3[A1, A2, A3, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor2[A2, A3, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor3[A1, A2, A3, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor2[A2, A3, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor3[A1, A2, A3, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor2[A2, A3, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) ApplicativeFunctor3[A1, A2, A3, R] {
	return ApplicativeFunctor3[A1, A2, A3, R]{Successful(curried.Func3(fn))}
}

type ApplicativeFunctor4[A1, A2, A3, A4, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor3[A2, A3, A4, R] {

	return ApplicativeFunctor3[A2, A3, A4, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor3[A2, A3, A4, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor3[A2, A3, A4, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) Ap(a A1) ApplicativeFunctor3[A2, A3, A4, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor3[A2, A3, A4, R] {

	return ApplicativeFunctor3[A2, A3, A4, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor3[A2, A3, A4, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor3[A2, A3, A4, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor3[A2, A3, A4, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) ApplicativeFunctor4[A1, A2, A3, A4, R] {
	return ApplicativeFunctor4[A1, A2, A3, A4, R]{Successful(curried.Func4(fn))}
}

type ApplicativeFunctor5[A1, A2, A3, A4, A5, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return ApplicativeFunctor4[A2, A3, A4, A5, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) Ap(a A1) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return ApplicativeFunctor4[A2, A3, A4, A5, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor4[A2, A3, A4, A5, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) ApplicativeFunctor5[A1, A2, A3, A4, A5, R] {
	return ApplicativeFunctor5[A1, A2, A3, A4, A5, R]{Successful(curried.Func5(fn))}
}

type ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return ApplicativeFunctor5[A2, A3, A4, A5, A6, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) Ap(a A1) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return ApplicativeFunctor5[A2, A3, A4, A5, A6, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]{Successful(curried.Func6(fn))}
}

type ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]{Successful(curried.Func7(fn))}
}

type ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]{Successful(curried.Func8(fn))}
}

type ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	fn fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R]{Ap(r.fn, a)}
}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(FromTry(a))
}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(FromOption(a))
}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFuture(Successful(a))

}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.Executor) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R]{ApFunc(r.fn, a)}

}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.Executor) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.Executor) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
}
func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFunc(a func() A1, ctx ...fp.Executor) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApFutureFunc(func() fp.Future[A1] {
		return Successful(a())
	}, ctx...)
}
func Applicative9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{Successful(curried.Func9(fn))}
}
