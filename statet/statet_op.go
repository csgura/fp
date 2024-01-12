package statet

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/try"
)

func Run[S, A any](f func(S) (A, S)) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Success(as.Tuple(f(s)))
	}
}

func RunT[S, A any](f func(S) (fp.Try[A], fp.Try[S])) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Zip(f(s))
	}
}

func Put[S any](s S) fp.StateT[S, fp.Unit] {
	return func(s S) fp.Try[fp.Tuple2[fp.Unit, S]] {
		return try.Success(as.Tuple(fp.Unit{}, s))
	}
}

func Get[S any]() fp.StateT[S, S] {
	return func(s S) fp.Try[fp.Tuple2[S, S]] {
		return try.Success(as.Tuple(s, s))
	}
}

func Modify[S any](f func(S) S) fp.StateT[S, fp.Unit] {
	return func(s S) fp.Try[fp.Tuple2[fp.Unit, S]] {
		return try.Success(as.Tuple(fp.Unit{}, f(s)))
	}
}

func ModifyT[S any](f func(S) fp.Try[S]) fp.StateT[S, fp.Unit] {
	return func(s S) fp.Try[fp.Tuple2[fp.Unit, S]] {
		return try.Zip(try.Success(fp.Unit{}), f(s))
	}
}

func Eval[S, A any](f func(S) A) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Success(as.Tuple(f(s), s))
	}
}

func EvalT[S, A any](f func(S) fp.Try[A]) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Zip(f(s), try.Success(s))
	}
}

func Pure[S, A any](a A) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Success(as.Tuple2(a, s))
	}
}

func FromTry[S, A any](t fp.Try[A]) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Map(t, as.Func2(as.Tuple2[A, S]).ApplyLast(s))
	}
}

func MapState[S, A, B any](st fp.StateT[S, A], f func(A) fp.State[S, B]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		ns := st(s)
		return try.Map(ns, func(v fp.Tuple2[A, S]) fp.Tuple2[B, S] {
			return f(v.I1)(v.I2)
		})
	}
}

// FlatMap 과 동일
func MapStateT[S, A, B any](st fp.StateT[S, A], f func(A) fp.StateT[S, B]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		ns := st(s)
		return try.FlatMap(ns, func(v fp.Tuple2[A, S]) fp.Try[fp.Tuple2[B, S]] {
			return f(v.I1)(v.I2)
		})
	}
}

func FlatMap[S, A, B any](st fp.StateT[S, A], f func(A) fp.StateT[S, B]) fp.StateT[S, B] {
	return MapStateT(st, f)
}

func FlatMapConst[S, A, B any](st fp.StateT[S, A], next fp.StateT[S, B]) fp.StateT[S, B] {
	return FlatMap(st, fp.Const[A](next))
}

func WithState[S, A any](st fp.StateT[S, A], f func(S) S) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		return try.Map(st(s), as.Func2(product.MapValue[A, S, S]).ApplyLast(f))
	}
}

func WithStateT[S, A any](st fp.StateT[S, A], f func(S) fp.Try[S]) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		a, ns := st.Run(s)
		return try.Zip(a, try.FlatMap(ns, f))
	}
}

func Ap[S, A, B any](st fp.StateT[S, fp.Func1[A, B]], a A) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		af, ns := st.Run(s)
		return try.Zip(try.Ap(af, try.Success(a)), ns)
	}
}

func ApTry[S, A, B any](st fp.StateT[S, fp.Func1[A, B]], a fp.Try[A]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		af, ns := st.Run(s)
		return try.Zip(try.Ap(af, a), ns)
	}
}

func ApOption[S, A, B any](st fp.StateT[S, fp.Func1[A, B]], a fp.Option[A]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		af, ns := st.Run(s)
		return try.Zip(try.Ap(af, try.FromOption(a)), ns)
	}
}

func Map[S, A, B any](st fp.StateT[S, A], f func(A) B) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		a, ns := st.Run(s)
		return try.Zip(try.Map(a, f), ns)
	}
}

func Replace[S, A, B any](s fp.StateT[S, A], b B) fp.StateT[S, B] {
	return Map(s, fp.Const[A](b))
}

func MapWithState[S, A, B any](st fp.StateT[S, A], f func(S, A) B) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		a, ns := st.Run(s)
		return try.Zip(try.Map2(ns, a, f), ns)
	}
}

func MapT[S, A, B any](st fp.StateT[S, A], f func(A) fp.Try[B]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		a, ns := st.Run(s)
		return try.Zip(try.FlatMap(a, f), ns)
	}
}

func MapWithStateT[S, A, B any](st fp.StateT[S, A], f func(S, A) fp.Try[B]) fp.StateT[S, B] {
	return func(s S) fp.Try[fp.Tuple2[B, S]] {
		a, ns := st.Run(s)
		return try.Zip(try.LiftM2(f)(ns, a), ns)
	}
}

func PeekState[S, A any](st fp.StateT[S, A], f func(ctx S)) fp.StateT[S, A] {
	return func(s S) fp.Try[fp.Tuple2[A, S]] {
		r, ns := st.Run(s)
		ns.Foreach(f)
		return try.Zip(r, ns)
	}
}

func Map2[S, A, B, R any](first fp.StateT[S, A], second fp.StateT[S, B], fab func(A, B) R) fp.StateT[S, R] {
	return FlatMap(first, func(a A) fp.StateT[S, R] {
		return Map(second, func(b B) R {
			return fab(a, b)
		})
	})
}

func Zip[S, A, B any](first fp.StateT[S, A], second fp.StateT[S, B]) fp.StateT[S, fp.Tuple2[A, B]] {
	return Map2(first, second, product.Tuple2)
}
