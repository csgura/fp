package match

import (
	"errors"
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

func andthen[A, B, R any](pf fp.PartialFunc[A, B], then func(B) R) fp.PartialFunc[A, R] {
	return as.PartialFunc(pf.IsDefinedAt, func(a A) R {
		return then(pf.Apply(a))
	})
}

func combine[A, B, R any](pf fp.PartialFunc[A, B], then fp.PartialFunc[B, R]) fp.PartialFunc[A, R] {
	return as.PartialFunc(func(a A) bool {
		if pf.IsDefinedAt(a) {
			return then.IsDefinedAt(pf.Apply(a))
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
		return hguard.IsDefinedAt(c.Head()) && tguard.IsDefinedAt(c.Tail())
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
		return option.Map(c.Head(), hguard.IsDefinedAt).OrZero() && tguard.IsDefinedAt(c.Tail())
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

func CaseStatusCode[T, R any](guard int, then func(T) R) fp.PartialFunc[T, R] {
	return as.PartialFunc(func(t T) bool {
		if sc, ok := any(t).(interface {
			StatusCode() int
		}); ok {
			return sc.StatusCode() == guard
		}
		return false

	}, func(c T) R {
		return then(c)
	})
}

func CaseErrorType[E error, R any](then func(E) R) fp.PartialFunc[error, R] {
	return as.PartialFunc(func(err error) bool {

		var target E
		return errors.As(err, &target)

	}, func(err error) R {
		var target E
		if errors.As(err, &target) {
			return then(target)
		}
		panic(fmt.Sprintf("error is not %T", target))
	})
}

func CaseErrorCode[R any](guard int, then func(error) R) fp.PartialFunc[error, R] {
	return as.PartialFunc(func(t error) bool {
		if sc, ok := t.(interface {
			StatusCode() int
		}); ok {
			return sc.StatusCode() == guard
		}
		return false

	}, func(c error) R {
		return then(c)
	})
}

func CaseErrorIs[R any](guard error, then func(error) R) fp.PartialFunc[error, R] {
	return as.PartialFunc(func(err error) bool {
		return errors.Is(err, guard)
	}, func(c error) R {
		return then(c)
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
