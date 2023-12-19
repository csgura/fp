package match

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

// Case 로 시작하는 함수들은  then 함수를 인자로 받음
// And 로 끝나는 함수는,  추가로 적용될 guard 을 인자로 받음.

// 가장 기본인 Case 는  guard 를 인자로 받지만,  And 로 끝나지 않음.
func Case[V, T, R any](guard fp.PartialFunc[V, T], then func(T) R) fp.PartialFunc[V, R] {
	return func(v V) fp.Option[R] {
		return option.Map(guard(v), then)
	}
}

func CaseTuple2[V1, V2, T1, T2, R any](guard1 fp.PartialFunc[V1, T1], guard2 fp.PartialFunc[V2, T2], then func(T1, T2) R) fp.PartialFunc[fp.Tuple2[V1, V2], R] {
	return Case(Tuple2(guard1, guard2), as.Tupled2(then))
}

func CaseConsAnd[C fp.Cons[H1, T1], H1, H2, T1, T2, R any](hguard fp.PartialFunc[H1, H2], tguard fp.PartialFunc[T1, T2], then func(H2, T2) R) fp.PartialFunc[C, R] {
	return func(c C) fp.Option[R] {
		return option.Map2(hguard(c.Head()), tguard(c.Tail()), then)
	}
}

func CaseSeqCons[T, R any](then func(T, fp.Seq[T]) R) fp.PartialFunc[fp.Seq[T], R] {
	return func(s fp.Seq[T]) fp.Option[R] {
		return option.Map2(s.Head(), option.Of(s.Tail()), then)
	}
}

func CaseSeqConsAnd[T, HT, TT, R any](hguard fp.PartialFunc[T, HT], tguard fp.PartialFunc[fp.Seq[T], TT], then func(HT, TT) R) fp.PartialFunc[fp.Seq[T], R] {
	return func(s fp.Seq[T]) fp.Option[R] {
		return option.Map2(option.FlatMap(s.Head(), hguard), tguard(s.Tail()), then)
	}
}

func CaseNone[T, R any](then func() R) fp.PartialFunc[fp.Option[T], R] {
	return func(v fp.Option[T]) fp.Option[R] {
		return option.Map(None[T](v), as.Func0(then))
	}
}

func CaseAny[T, R any](then func(T) R) fp.PartialFunc[T, R] {
	return func(v T) fp.Option[R] {
		return option.Some(then(v))
	}
}

func CaseSeqEmpty[T, R any](then func() R) fp.PartialFunc[fp.Seq[T], R] {
	return func(v fp.Seq[T]) fp.Option[R] {
		if v.IsEmpty() {
			return option.Some(then())
		}
		return option.None[R]()
	}
}

func CaseTrue[T, R any](predicate fp.Predicate[T], then func() R) fp.PartialFunc[T, R] {
	return func(v T) fp.Option[R] {
		if predicate(v) {
			return option.Some(then())
		}
		return option.None[R]()
	}
}

func CaseStatusCode[T interface {
	StatusCode() int
}, R any](guard int, then func() R) fp.PartialFunc[T, R] {
	return func(v T) fp.Option[R] {
		if v.StatusCode() == guard {
			return option.Some(then())
		}
		return option.None[R]()
	}
}

func CaseEmpty[T interface {
	IsEmpty() bool
}, R any](then func() R) fp.PartialFunc[T, R] {
	return func(v T) fp.Option[R] {
		if v.IsEmpty() {
			return option.Some(then())
		}
		return option.None[R]()
	}
}
