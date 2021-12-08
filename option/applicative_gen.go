package option

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
)

type ApplicativeFunctor2[H hlist.Header[HT], HT, A1, A2, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[H, HT, A1, A2, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Map(a func(HT) A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListMap(a func(H) A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Ap(a A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApOption(Some(a))

}
func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R] {
	return ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R]{Some(hlist.Empty()), Some(curried.Func2(fn))}
}

type ApplicativeFunctor3[H hlist.Header[HT], HT, A1, A2, A3, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Map(a func(HT) A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListMap(a func(H) A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor3[H, HT, A1, A2, A3, R]) Ap(a A1) ApplicativeFunctor2[hlist.Cons[A1, H], A1, A2, A3, R] {

	return r.ApOption(Some(a))

}
func Applicative3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) ApplicativeFunctor3[hlist.Nil, hlist.Nil, A1, A2, A3, R] {
	return ApplicativeFunctor3[hlist.Nil, hlist.Nil, A1, A2, A3, R]{Some(hlist.Empty()), Some(curried.Func3(fn))}
}

type ApplicativeFunctor4[H hlist.Header[HT], HT, A1, A2, A3, A4, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Map(a func(HT) A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListMap(a func(H) A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor4[H, HT, A1, A2, A3, A4, R]) Ap(a A1) ApplicativeFunctor3[hlist.Cons[A1, H], A1, A2, A3, A4, R] {

	return r.ApOption(Some(a))

}
func Applicative4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) ApplicativeFunctor4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R] {
	return ApplicativeFunctor4[hlist.Nil, hlist.Nil, A1, A2, A3, A4, R]{Some(hlist.Empty()), Some(curried.Func4(fn))}
}

type ApplicativeFunctor5[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) Map(a func(HT) A1) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) HListMap(a func(H) A1) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor5[H, HT, A1, A2, A3, A4, A5, R]) Ap(a A1) ApplicativeFunctor4[hlist.Cons[A1, H], A1, A2, A3, A4, A5, R] {

	return r.ApOption(Some(a))

}
func Applicative5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) ApplicativeFunctor5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R] {
	return ApplicativeFunctor5[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, R]{Some(hlist.Empty()), Some(curried.Func5(fn))}
}

type ApplicativeFunctor6[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) Map(a func(HT) A1) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListMap(a func(H) A1) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor6[H, HT, A1, A2, A3, A4, A5, A6, R]) Ap(a A1) ApplicativeFunctor5[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, R] {

	return r.ApOption(Some(a))

}
func Applicative6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) ApplicativeFunctor6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor6[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, R]{Some(hlist.Empty()), Some(curried.Func6(fn))}
}

type ApplicativeFunctor7[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Map(a func(HT) A1) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListMap(a func(H) A1) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor7[H, HT, A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) ApplicativeFunctor6[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApOption(Some(a))

}
func Applicative7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplicativeFunctor7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor7[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, R]{Some(hlist.Empty()), Some(curried.Func7(fn))}
}

type ApplicativeFunctor8[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Map(a func(HT) A1) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListMap(a func(H) A1) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor8[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) ApplicativeFunctor7[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApOption(Some(a))

}
func Applicative8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplicativeFunctor8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor8[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, R]{Some(hlist.Empty()), Some(curried.Func8(fn))}
}

type ApplicativeFunctor9[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Map(a func(HT) A1) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListMap(a func(H) A1) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor9[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) ApplicativeFunctor8[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApOption(Some(a))

}
func Applicative9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplicativeFunctor9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor9[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{Some(hlist.Empty()), Some(curried.Func9(fn))}
}

type ApplicativeFunctor10[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, R]]]]]]]]]]]
}

func (r ApplicativeFunctor10[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor10[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Map(a func(HT) A1) ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor10[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) HListMap(a func(H) A1) ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor10[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor10[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor10[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Ap(a A1) ApplicativeFunctor9[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	return r.ApOption(Some(a))

}
func Applicative10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](fn fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) ApplicativeFunctor10[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return ApplicativeFunctor10[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]{Some(hlist.Empty()), Some(curried.Func10(fn))}
}

type ApplicativeFunctor11[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, R]]]]]]]]]]]]
}

func (r ApplicativeFunctor11[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor11[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Map(a func(HT) A1) ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor11[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) HListMap(a func(H) A1) ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor11[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor11[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor11[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Ap(a A1) ApplicativeFunctor10[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	return r.ApOption(Some(a))

}
func Applicative11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](fn fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) ApplicativeFunctor11[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return ApplicativeFunctor11[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]{Some(hlist.Empty()), Some(curried.Func11(fn))}
}

type ApplicativeFunctor12[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, R]]]]]]]]]]]]]
}

func (r ApplicativeFunctor12[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor12[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Map(a func(HT) A1) ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor12[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) HListMap(a func(H) A1) ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor12[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor12[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor12[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Ap(a A1) ApplicativeFunctor11[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	return r.ApOption(Some(a))

}
func Applicative12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](fn fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) ApplicativeFunctor12[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return ApplicativeFunctor12[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]{Some(hlist.Empty()), Some(curried.Func12(fn))}
}

type ApplicativeFunctor13[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, R]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor13[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor13[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Map(a func(HT) A1) ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor13[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) HListMap(a func(H) A1) ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor13[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor13[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor13[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Ap(a A1) ApplicativeFunctor12[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	return r.ApOption(Some(a))

}
func Applicative13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](fn fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) ApplicativeFunctor13[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return ApplicativeFunctor13[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]{Some(hlist.Empty()), Some(curried.Func13(fn))}
}

type ApplicativeFunctor14[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, R]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor14[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor14[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Map(a func(HT) A1) ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor14[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) HListMap(a func(H) A1) ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor14[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor14[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor14[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Ap(a A1) ApplicativeFunctor13[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	return r.ApOption(Some(a))

}
func Applicative14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](fn fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) ApplicativeFunctor14[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return ApplicativeFunctor14[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]{Some(hlist.Empty()), Some(curried.Func14(fn))}
}

type ApplicativeFunctor15[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, R]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor15[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor15[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Map(a func(HT) A1) ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor15[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) HListMap(a func(H) A1) ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor15[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor15[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor15[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Ap(a A1) ApplicativeFunctor14[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {

	return r.ApOption(Some(a))

}
func Applicative15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any](fn fp.Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) ApplicativeFunctor15[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return ApplicativeFunctor15[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]{Some(hlist.Empty()), Some(curried.Func15(fn))}
}

type ApplicativeFunctor16[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, R]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor16[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor16[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Map(a func(HT) A1) ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor16[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) HListMap(a func(H) A1) ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor16[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor16[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor16[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Ap(a A1) ApplicativeFunctor15[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {

	return r.ApOption(Some(a))

}
func Applicative16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any](fn fp.Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) ApplicativeFunctor16[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return ApplicativeFunctor16[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]{Some(hlist.Empty()), Some(curried.Func16(fn))}
}

type ApplicativeFunctor17[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, fp.Func1[A17, R]]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor17[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor17[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) Map(a func(HT) A1) ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor17[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) HListMap(a func(H) A1) ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor17[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor17[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor17[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) Ap(a A1) ApplicativeFunctor16[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {

	return r.ApOption(Some(a))

}
func Applicative17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any](fn fp.Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]) ApplicativeFunctor17[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {
	return ApplicativeFunctor17[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R]{Some(hlist.Empty()), Some(curried.Func17(fn))}
}

type ApplicativeFunctor18[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, fp.Func1[A17, fp.Func1[A18, R]]]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor18[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor18[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) Map(a func(HT) A1) ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor18[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) HListMap(a func(H) A1) ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor18[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor18[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor18[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) Ap(a A1) ApplicativeFunctor17[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {

	return r.ApOption(Some(a))

}
func Applicative18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any](fn fp.Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]) ApplicativeFunctor18[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {
	return ApplicativeFunctor18[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R]{Some(hlist.Empty()), Some(curried.Func18(fn))}
}

type ApplicativeFunctor19[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, fp.Func1[A17, fp.Func1[A18, fp.Func1[A19, R]]]]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor19[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor19[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) Map(a func(HT) A1) ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor19[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) HListMap(a func(H) A1) ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor19[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor19[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor19[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) Ap(a A1) ApplicativeFunctor18[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {

	return r.ApOption(Some(a))

}
func Applicative19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any](fn fp.Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]) ApplicativeFunctor19[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {
	return ApplicativeFunctor19[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R]{Some(hlist.Empty()), Some(curried.Func19(fn))}
}

type ApplicativeFunctor20[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, fp.Func1[A17, fp.Func1[A18, fp.Func1[A19, fp.Func1[A20, R]]]]]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor20[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor20[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) Map(a func(HT) A1) ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor20[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) HListMap(a func(H) A1) ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor20[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor20[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor20[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) Ap(a A1) ApplicativeFunctor19[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {

	return r.ApOption(Some(a))

}
func Applicative20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any](fn fp.Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]) ApplicativeFunctor20[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {
	return ApplicativeFunctor20[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R]{Some(hlist.Empty()), Some(curried.Func20(fn))}
}

type ApplicativeFunctor21[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, fp.Func1[A17, fp.Func1[A18, fp.Func1[A19, fp.Func1[A20, fp.Func1[A21, R]]]]]]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor21[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor21[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) Map(a func(HT) A1) ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor21[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) HListMap(a func(H) A1) ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor21[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor21[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor21[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) Ap(a A1) ApplicativeFunctor20[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {

	return r.ApOption(Some(a))

}
func Applicative21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any](fn fp.Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]) ApplicativeFunctor21[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {
	return ApplicativeFunctor21[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R]{Some(hlist.Empty()), Some(curried.Func21(fn))}
}

type ApplicativeFunctor22[H hlist.Header[HT], HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, fp.Func1[A15, fp.Func1[A16, fp.Func1[A17, fp.Func1[A18, fp.Func1[A19, fp.Func1[A20, fp.Func1[A21, fp.Func1[A22, R]]]]]]]]]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor22[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) FlatMap(a func(HT) fp.Option[A1]) ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
func (r ApplicativeFunctor22[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) Map(a func(HT) A1) ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor22[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) HListMap(a func(H) A1) ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
func (r ApplicativeFunctor22[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) HListFlatMap(a func(H) fp.Option[A1]) ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}
func (r ApplicativeFunctor22[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor22[H, HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) Ap(a A1) ApplicativeFunctor21[hlist.Cons[A1, H], A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {

	return r.ApOption(Some(a))

}
func Applicative22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any](fn fp.Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]) ApplicativeFunctor22[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {
	return ApplicativeFunctor22[hlist.Nil, hlist.Nil, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R]{Some(hlist.Empty()), Some(curried.Func22(fn))}
}
