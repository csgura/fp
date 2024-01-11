package state

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/try"
)

func New[S, A any](f func(S) (A, S)) fp.State[S, A] {
	return func(s S) fp.Tuple2[A, S] {
		return as.Tuple(f(s))
	}
}

func Put[S any](s S) fp.State[S, fp.Unit] {
	return func(s S) fp.Tuple2[fp.Unit, S] {
		return as.Tuple(fp.Unit{}, s)
	}
}

func Get[S any]() fp.State[S, S] {
	return func(s S) fp.Tuple2[S, S] {
		return as.Tuple(s, s)
	}
}

func Modify[S any](f func(S) S) fp.State[S, fp.Unit] {
	return func(s S) fp.Tuple2[fp.Unit, S] {
		return as.Tuple(fp.Unit{}, f(s))
	}
}

func Inspect[S, B any](f func(S) B) fp.State[S, B] {
	return func(s S) fp.Tuple2[B, S] {
		return as.Tuple(f(s), s)
	}
}

func Pure[S, A any](a A) fp.State[S, A] {
	return func(s S) fp.Tuple2[A, S] {
		return as.Tuple2(a, s)
	}
}

func LiftT[S, A any](st fp.State[S, A]) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Success(st(s))
	}
}

// FlatMap 과 동일
func MapState[S, A, B any](st fp.State[S, A], f func(A) fp.State[S, B]) fp.State[S, B] {
	return func(s S) fp.Tuple2[B, S] {
		v := st(s)
		return f(v.I1)(v.I2)
	}
}

func MapStateT[S, A, B any](st fp.State[S, A], f func(A) fp.StateT[S, B]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		v := st(s)
		return f(v.I1)(v.I2)
	}
}

func FlatMap[S, A, B any](st fp.State[S, A], f func(A) fp.State[S, B]) fp.State[S, B] {
	return MapState(st, f)
}

func FlatMapConst[S, A, B any](st fp.State[S, A], next fp.State[S, B]) fp.State[S, B] {
	return FlatMap(st, fp.Const[A](next))
}

func WithState[S, A any](st fp.State[S, A], f func(S) S) fp.State[S, A] {
	return func(s S) fp.Tuple2[A, S] {
		v := st(s)
		return product.MapValue(v, f)
	}
}

func Ap[S, A, B any](st fp.State[S, fp.Func1[A, B]], a A) fp.State[S, B] {
	return func(s S) fp.Tuple2[B, S] {
		v, ns := st.Run(s)
		return as.Tuple2(v(a), ns)
	}
}

func Map[S, A, B any](st fp.State[S, A], f func(A) B) fp.State[S, B] {
	return func(s S) fp.Tuple2[B, S] {
		a, ns := st.Run(s)
		return as.Tuple2(f(a), ns)
	}
}

func Replace[S, A, B any](s fp.State[S, A], b B) fp.State[S, B] {
	return Map(s, fp.Const[A](b))
}

func MapWithState[S, A, B any](st fp.State[S, A], f func(S, A) B) fp.State[S, B] {
	return func(s S) fp.Tuple2[B, S] {
		a, ns := st.Run(s)
		return as.Tuple2(f(ns, a), ns)
	}
}

func Map2[S, A, B, R any](first fp.State[S, A], second fp.State[S, B], fab func(A, B) R) fp.State[S, R] {
	return FlatMap(first, func(a A) fp.State[S, R] {
		return Map(second, func(b B) R {
			return fab(a, b)
		})
	})
}

func Zip[S, A, B any](first fp.State[S, A], second fp.State[S, B]) fp.State[S, fp.Tuple2[A, B]] {
	return Map2(first, second, product.Tuple2)
}
