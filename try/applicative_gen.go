// Code generated by template_gen, DO NOT EDIT.
package try

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)

// generated by try_op.go:465

type MonadChain2[H hlist.Header[HT], HT, A1, A2, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r MonadChain2[H, HT, A1, A2, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) Map(a func(HT) A1) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain2[H, HT, A1, A2, R]) HListMap(a func(H) A1) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain2[H, HT, A1, A2, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApTry(a fp.Try[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain1[hlist.Cons[A1, H], A1, A2, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain2[H, HT, A1, A2, R]) ApOption(a fp.Option[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain2[H, HT, A1, A2, R]) Ap(a A1) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain2[H, HT, A1, A2, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain2[H, HT, A1, A2, R]) ApFunc(a func() A1) MonadChain1[hlist.Cons[A1, H], A1, A2, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain2[A1, A2, R any](fn fp.Func2[A1, A2, R]) MonadChain2[hlist.Nil, hlist.Nil, A1, A2, R] {
	return MonadChain2[hlist.Nil, hlist.Nil, A1, A2, R]{Success(hlist.Empty()), Success(curried.Func2(fn))}
}

type MonadChain3[H hlist.Header[HT], HT, A1, A2, A3, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r MonadChain3[H, HT, A1, A2, A3, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) Map(a func(HT) A1) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) HListMap(a func(H) A1) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApTry(a fp.Try[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApOption(a fp.Option[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) Ap(a A1) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain3[H, HT, A1, A2, A3, R]) ApFunc(a func() A1) MonadChain2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) MonadChain3[hlist.Nil, hlist.Nil, A1, A2, A3, R] {
	return MonadChain3[hlist.Nil, hlist.Nil, A1, A2, A3, R]{Success(hlist.Empty()), Success(curried.Func3(fn))}
}

type MonadChain4[H hlist.Header[HT], HT, A1, A2, A3, A4, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) Map(a func(HT) A1) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) HListMap(a func(H) A1) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApTry(a fp.Try[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) Ap(a A1) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain4[H, HT, A1, A2, A3, A4, R]) ApFunc(a func() A1) MonadChain3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) MonadChain4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R] {
	return MonadChain4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R]{Success(hlist.Empty()), Success(curried.Func4(fn))}
}

type MonadChain5[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) Map(a func(HT) A1) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) HListMap(a func(H) A1) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApTry(a fp.Try[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) Ap(a A1) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain5[H, HT, A1, A2, A3, A4, A5, R]) ApFunc(a func() A1) MonadChain4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) MonadChain5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R] {
	return MonadChain5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R]{Success(hlist.Empty()), Success(curried.Func5(fn))}
}

type MonadChain6[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) Map(a func(HT) A1) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListMap(a func(H) A1) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApTry(a fp.Try[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) Ap(a A1) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApFunc(a func() A1) MonadChain5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) MonadChain6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R] {
	return MonadChain6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R]{Success(hlist.Empty()), Success(curried.Func6(fn))}
}

type MonadChain7[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Map(a func(HT) A1) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListMap(a func(H) A1) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApTry(a fp.Try[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApFunc(a func() A1) MonadChain6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) MonadChain7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R] {
	return MonadChain7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R]{Success(hlist.Empty()), Success(curried.Func7(fn))}
}

type MonadChain8[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Map(a func(HT) A1) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListMap(a func(H) A1) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTry(a fp.Try[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFunc(a func() A1) MonadChain7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) MonadChain8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return MonadChain8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R]{Success(hlist.Empty()), Success(curried.Func8(fn))}
}

type MonadChain9[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) FlatMap(a func(HT) fp.Try[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Map(a func(HT) A1) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListMap(a func(H) A1) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListFlatMap(a func(H) fp.Try[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTry(a fp.Try[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{nh, Ap(r.fn, a)}
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApTry(FromOption(a))
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApTry(Success(a))

}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTryFunc(a func() fp.Try[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOptionFunc(a func() fp.Option[A1]) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFunc(a func() A1) MonadChain8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) MonadChain9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return MonadChain9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{Success(hlist.Empty()), Success(curried.Func9(fn))}
}

// generated by try_op.go:555

type ApplicativeFunctor2[A1, A2, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[A1, A2, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor1[A2, R] {
	return ApplicativeFunctor1[A2, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor2[A1, A2, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2)
}

func (r ApplicativeFunctor2[A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[A2, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor2[A1, A2, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2)
}

func (r ApplicativeFunctor2[A1, A2, R]) Ap(a A1) ApplicativeFunctor1[A2, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor2[A1, A2, R]) ApAll(a1 A1, a2 A2) fp.Try[R] {
	return r.Ap(a1).Ap(a2)
}

func (r ApplicativeFunctor2[A1, A2, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor1[A2, R] {
	return ApplicativeFunctor1[A2, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor2[A1, A2, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor1[A2, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor2[A1, A2, R]) ApFunc(a func() A1) ApplicativeFunctor1[A2, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[A1, A2, R] {
	return ApplicativeFunctor2[A1, A2, R]{Success(curried.Func2(fn))}
}

type ApplicativeFunctor3[A1, A2, A3, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor2[A2, A3, R] {
	return ApplicativeFunctor2[A2, A3, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3)
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor2[A2, A3, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3)
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) Ap(a A1) ApplicativeFunctor2[A2, A3, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApAll(a1 A1, a2 A2, a3 A3) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3)
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor2[A2, A3, R] {
	return ApplicativeFunctor2[A2, A3, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor2[A2, A3, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor3[A1, A2, A3, R]) ApFunc(a func() A1) ApplicativeFunctor2[A2, A3, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) ApplicativeFunctor3[A1, A2, A3, R] {
	return ApplicativeFunctor3[A1, A2, A3, R]{Success(curried.Func3(fn))}
}

type ApplicativeFunctor4[A1, A2, A3, A4, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor3[A2, A3, A4, R] {
	return ApplicativeFunctor3[A2, A3, A4, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3).ApTry(ins4)
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor3[A2, A3, A4, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3).ApOption(ins4)
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) Ap(a A1) ApplicativeFunctor3[A2, A3, A4, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApAll(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3).Ap(a4)
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor3[A2, A3, A4, R] {
	return ApplicativeFunctor3[A2, A3, A4, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor3[A2, A3, A4, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor4[A1, A2, A3, A4, R]) ApFunc(a func() A1) ApplicativeFunctor3[A2, A3, A4, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) ApplicativeFunctor4[A1, A2, A3, A4, R] {
	return ApplicativeFunctor4[A1, A2, A3, A4, R]{Success(curried.Func4(fn))}
}

type ApplicativeFunctor5[A1, A2, A3, A4, A5, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {
	return ApplicativeFunctor4[A2, A3, A4, A5, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3).ApTry(ins4).ApTry(ins5)
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3).ApOption(ins4).ApOption(ins5)
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) Ap(a A1) ApplicativeFunctor4[A2, A3, A4, A5, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApAll(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3).Ap(a4).Ap(a5)
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {
	return ApplicativeFunctor4[A2, A3, A4, A5, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor4[A2, A3, A4, A5, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor5[A1, A2, A3, A4, A5, R]) ApFunc(a func() A1) ApplicativeFunctor4[A2, A3, A4, A5, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) ApplicativeFunctor5[A1, A2, A3, A4, A5, R] {
	return ApplicativeFunctor5[A1, A2, A3, A4, A5, R]{Success(curried.Func5(fn))}
}

type ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor5[A2, A3, A4, A5, A6, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3).ApTry(ins4).ApTry(ins5).ApTry(ins6)
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3).ApOption(ins4).ApOption(ins5).ApOption(ins6)
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) Ap(a A1) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApAll(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3).Ap(a4).Ap(a5).Ap(a6)
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor5[A2, A3, A4, A5, A6, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]) ApFunc(a func() A1) ApplicativeFunctor5[A2, A3, A4, A5, A6, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, R]{Success(curried.Func6(fn))}
}

type ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3).ApTry(ins4).ApTry(ins5).ApTry(ins6).ApTry(ins7)
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6], ins7 fp.Option[A7]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3).ApOption(ins4).ApOption(ins5).ApOption(ins6).ApOption(ins7)
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApAll(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3).Ap(a4).Ap(a5).Ap(a6).Ap(a7)
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]) ApFunc(a func() A1) ApplicativeFunctor6[A2, A3, A4, A5, A6, A7, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, R]{Success(curried.Func7(fn))}
}

type ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7], ins8 fp.Try[A8]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3).ApTry(ins4).ApTry(ins5).ApTry(ins6).ApTry(ins7).ApTry(ins8)
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6], ins7 fp.Option[A7], ins8 fp.Option[A8]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3).ApOption(ins4).ApOption(ins5).ApOption(ins6).ApOption(ins7).ApOption(ins8)
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApAll(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3).Ap(a4).Ap(a5).Ap(a6).Ap(a7).Ap(a8)
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApFunc(a func() A1) ApplicativeFunctor7[A2, A3, A4, A5, A6, A7, A8, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, R]{Success(curried.Func8(fn))}
}

type ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	fn fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTry(a fp.Try[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R]{Ap(r.fn, a)}
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTryAll(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7], ins8 fp.Try[A8], ins9 fp.Try[A9]) fp.Try[R] {
	return r.ApTry(ins1).ApTry(ins2).ApTry(ins3).ApTry(ins4).ApTry(ins5).ApTry(ins6).ApTry(ins7).ApTry(ins8).ApTry(ins9)
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOptionAll(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6], ins7 fp.Option[A7], ins8 fp.Option[A8], ins9 fp.Option[A9]) fp.Try[R] {
	return r.ApOption(ins1).ApOption(ins2).ApOption(ins3).ApOption(ins4).ApOption(ins5).ApOption(ins6).ApOption(ins7).ApOption(ins8).ApOption(ins9)
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApAll(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
	return r.Ap(a1).Ap(a2).Ap(a3).Ap(a4).Ap(a5).Ap(a6).Ap(a7).Ap(a8).Ap(a9)
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApTryFunc(a func() fp.Try[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R]{ApFunc(r.fn, a)}
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOptionFunc(a func() fp.Option[A1]) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}

func (r ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApFunc(a func() A1) ApplicativeFunctor8[A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}

func Applicative9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{Success(curried.Func9(fn))}
}
