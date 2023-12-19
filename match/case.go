package match

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

func andthen[A, B, R any](pf fp.PartialFunc[A, B], then func(B) R) fp.PartialFunc[A, R] {
	return as.PartialFunc(pf.IsDefined, func(a A) R {
		return then(pf.Apply(a))
	})
}

func combine[A, B, R any](pf fp.PartialFunc[A, B], then fp.PartialFunc[B, R]) fp.PartialFunc[A, R] {
	return as.PartialFunc(func(a A) bool {
		if pf.IsDefined(a) {
			return then.IsDefined(pf.Apply(a))
		}
		return false
	}, func(a A) R {
		return then.Apply(pf.Apply(a))
	})
}

// Case 로 시작하는 함수들은  then 함수를 인자로 받음
// And 로 끝나는 함수는,  추가로 적용될 guard 을 인자로 받음.

// 가장 기본인 Case 는  guard 를 인자로 받지만,  And 로 끝나지 않음.
func Case[V, T, R any](guard fp.PartialFunc[V, T], then func(T) R) fp.PartialFunc[V, R] {
	return andthen(guard, then)
}

func CaseTuple2[V1, V2, T1, T2, R any](guard1 fp.PartialFunc[V1, T1], guard2 fp.PartialFunc[V2, T2], then func(T1, T2) R) fp.PartialFunc[fp.Tuple2[V1, V2], R] {
	return Case(Tuple2(guard1, guard2), as.Tupled2(then))
}

func CaseConsAnd[C fp.Cons[H1, T1], H1, H2, T1, T2, R any](hguard fp.PartialFunc[H1, H2], tguard fp.PartialFunc[T1, T2], then func(H2, T2) R) fp.PartialFunc[C, R] {
	return as.PartialFunc(func(c C) bool {
		return hguard.IsDefined(c.Head()) && tguard.IsDefined(c.Tail())
	}, func(c C) R {
		return then(hguard.Apply(c.Head()), tguard.Apply(c.Tail()))
	})
}

func CaseSeqCons[T, R any](then func(T, fp.Seq[T]) R) fp.PartialFunc[fp.Seq[T], R] {
	return as.PartialFunc(func(c fp.Seq[T]) bool {
		return c.Head().IsDefined()
	}, func(c fp.Seq[T]) R {
		return then(c.Head().Get(), c.Tail())
	})
}

func CaseSeqConsAnd[T, HT, TT, R any](hguard fp.PartialFunc[T, HT], tguard fp.PartialFunc[fp.Seq[T], TT], then func(HT, TT) R) fp.PartialFunc[fp.Seq[T], R] {
	return as.PartialFunc(func(c fp.Seq[T]) bool {
		return option.Map(c.Head(), hguard.IsDefined).OrZero() && tguard.IsDefined(c.Tail())
	}, func(c fp.Seq[T]) R {
		return then(hguard.Apply(c.Head().Get()), tguard.Apply(c.Tail()))
	})
}

func CaseNone[T, R any](then func() R) fp.PartialFunc[fp.Option[T], R] {
	return as.PartialFunc(func(t fp.Option[T]) bool {
		return t.IsEmpty()
	}, func(t fp.Option[T]) R {
		return then()
	})

}

func CaseAny[T, R any](then func(T) R) fp.PartialFunc[T, R] {
	return as.PartialFunc(func(t T) bool {
		return true
	}, func(t T) R {
		return then(t)
	})
}

func CaseSeqEmpty[T, R any](then func() R) fp.PartialFunc[fp.Seq[T], R] {
	return as.PartialFunc(func(c fp.Seq[T]) bool {
		return c.IsEmpty()
	}, func(c fp.Seq[T]) R {
		return then()
	})
}

func CaseTrue[T, R any](predicate fp.Predicate[T], then func() R) fp.PartialFunc[T, R] {
	return as.PartialFunc(predicate, func(c T) R {
		return then()
	})
}

func CaseStatusCode[T interface {
	StatusCode() int
}, R any](guard int, then func() R) fp.PartialFunc[T, R] {
	return as.PartialFunc(func(t T) bool {
		return t.StatusCode() == guard
	}, func(c T) R {
		return then()
	})
}

func CaseEmpty[T interface {
	IsEmpty() bool
}, R any](then func() R) fp.PartialFunc[T, R] {
	return as.PartialFunc(func(c T) bool {
		return c.IsEmpty()
	}, func(c T) R {
		return then()
	})
}
