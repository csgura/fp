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

func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Shift() ApplicativeFunctor2[H, HT, A2, A1, R] {

	nf := fp.Compose(curried.Revert2[A1, A2, R], fp.Compose(fp.Func2[A1, A2, R].Shift, fp.Func2[A2, A1, R].Curried))
	return ApplicativeFunctor2[H, HT, A2, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Map(a func(HT) A1, ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListMap(a func(H) A1, ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	nh := FlatMap(r.h, func(hv H) fp.Future[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

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
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApFunc(a func() A1, ctx ...fp.ExecContext) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

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

func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Shift() ApplicativeFunctor3[H, HT, A2, A3, A1, R] {

	nf := fp.Compose(curried.Revert3[A1, A2, A3, R], fp.Compose(fp.Func3[A1, A2, A3, R].Shift, fp.Func3[A2, A3, A1, R].Curried))
	return ApplicativeFunctor3[H, HT, A2, A3, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Map(a func(HT) A1, ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListMap(a func(H) A1, ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	nh := FlatMap(r.h, func(hv H) fp.Future[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

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
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApFunc(a func() A1, ctx ...fp.ExecContext) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

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

func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Shift() ApplicativeFunctor4[H, HT, A2, A3, A4, A1, R] {

	nf := fp.Compose(curried.Revert4[A1, A2, A3, A4, R], fp.Compose(fp.Func4[A1, A2, A3, A4, R].Shift, fp.Func4[A2, A3, A4, A1, R].Curried))
	return ApplicativeFunctor4[H, HT, A2, A3, A4, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Map(a func(HT) A1, ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListMap(a func(H) A1, ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	nh := FlatMap(r.h, func(hv H) fp.Future[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

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
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApFunc(a func() A1, ctx ...fp.ExecContext) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

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

func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) Shift() ApplicativeFunctor5[H, HT, A2, A3, A4, A5, A1, R] {

	nf := fp.Compose(curried.Revert5[A1, A2, A3, A4, A5, R], fp.Compose(fp.Func5[A1, A2, A3, A4, A5, R].Shift, fp.Func5[A2, A3, A4, A5, A1, R].Curried))
	return ApplicativeFunctor5[H, HT, A2, A3, A4, A5, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) FlatMap(a func(HT) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v.Head())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) Map(a func(HT) A1, ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.FlatMap(func(h HT) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) HListMap(a func(H) A1, ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.HListFlatMap(func(h H) fp.Future[A1] {
		return Successful(a(h))
	}, ctx...)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) HListFlatMap(a func(H) fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a(v)
	}, ctx...)

	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApFuture(a fp.Future[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	nh := FlatMap(r.h, func(hv H) fp.Future[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

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
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApFutureFunc(a func() fp.Future[A1], ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApTryFunc(a func() fp.Try[A1], ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromTry(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApOptionFunc(a func() fp.Option[A1], ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Future[A1] {
		return FromOption(a())
	}, ctx...)
	return r.ApFuture(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApFunc(a func() A1, ctx ...fp.ExecContext) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	}, ctx...)
	return r.ApFuture(av)
}
func Applicative5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) ApplicativeFunctor5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R] {
	return ApplicativeFunctor5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R]{Successful(hlist.Empty()), Successful(curried.Func5(fn))}
}
