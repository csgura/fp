package tstate

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/xtr"
)

type State[S, A any] func(S) (fp.Try[S], fp.Try[A])

func (r State[S, A]) Run(s S) (fp.Try[S], fp.Try[A]) {
	return r(s)
}

func (r State[S, A]) Exec(s S) fp.Try[S] {
	state, _ := r(s)
	return state
}

func (r State[S, A]) Eval(s S) fp.Try[A] {
	_, result := r(s)
	return result
}

func Pure[S, T any](t T) State[S, T] {
	return func(s S) (fp.Try[S], fp.Try[T]) {
		return try.Success(s), try.Success(t)
	}
}

func FromTry[S, T any](t fp.Try[T]) State[S, T] {
	return func(s S) (fp.Try[S], fp.Try[T]) {
		return try.Success(s), t
	}
}

func MapState[S, A, B any](st State[S, A], f func(S, A) (S, B)) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		b := try.Map2(ns, a, func(s S, a A) fp.Tuple2[S, B] {
			return as.Tuple2(f(s, a))
		})
		return try.Map(b, xtr.Head), try.Map(b, xtr.Tail)
	}
}

func MapStateT[S, A, B any](st State[S, A], f func(S, A) (fp.Try[S], fp.Try[B])) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		b := try.LiftM2(func(s S, a A) fp.Try[fp.Tuple2[S, B]] {
			ts, tb := f(s, a)
			return try.Map2(ts, tb, as.Tuple2)
		})(ns, a)
		return try.Map(b, xtr.Head), try.Map(b, xtr.Tail)
	}
}

func FlatMap[S, A, B any](st State[S, A], f func(A) State[S, B]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		if ns.IsSuccess() && a.IsSuccess() {
			return f(a.Get())(ns.Get())
		}
		if a.IsFailure() {
			return ns, try.Failure[B](a.Failed().Get())
		}
		return ns, try.Failure[B](ns.Failed().Get())
	}
}

func WithState[S, A any](st State[S, A], f func(S) S) State[S, A] {
	return func(s S) (fp.Try[S], fp.Try[A]) {
		ns, a := st(s)
		return try.Map(ns, f), a
	}
}

func WithStateT[S, A any](st State[S, A], f func(S) fp.Try[S]) State[S, A] {
	return func(s S) (fp.Try[S], fp.Try[A]) {
		ns, a := st(s)
		return try.FlatMap(ns, f), a
	}
}

func Ap[S, A, B any](st State[S, fp.Func1[A, B]], a A) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, af := st(s)
		return ns, try.Ap(af, try.Success(a))
	}
}

func ApTry[S, A, B any](st State[S, fp.Func1[A, B]], a fp.Try[A]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, af := st(s)
		return ns, try.Ap(af, a)
	}
}

func ApOption[S, A, B any](st State[S, fp.Func1[A, B]], a fp.Option[A]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, af := st(s)
		return ns, try.Ap(af, try.FromOption(a))
	}
}

func Inspect[S, A, B any](st State[S, A], f func(S) B) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, _ := st(s)
		return ns, try.Map(ns, f)
	}
}

func Map[S, A, B any](st State[S, A], f func(A) B) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		return ns, try.Map(a, f)
	}
}

func MapWithState[S, A, B any](st State[S, A], f func(S, A) B) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		return ns, try.Map2(ns, a, f)
	}
}

func MapT[S, A, B any](st State[S, A], f func(A) fp.Try[B]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		return ns, try.FlatMap(a, f)
	}
}

func MapWithStateT[S, A, B any](st State[S, A], f func(S, A) fp.Try[B]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		return ns, try.LiftM2(f)(ns, a)
	}
}

func PeekState[S, A any](st State[S, A], f func(ctx S)) State[S, A] {
	return func(s S) (fp.Try[S], fp.Try[A]) {
		ns, r := st(s)
		ns.Foreach(f)
		return ns, r
	}
}
