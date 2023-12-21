package tstate

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/try"
)

type State[S, A any] func(S) (fp.Try[S], fp.Try[A])

func (r State[S, A]) Run(ctx S) fp.Try[A] {
	_, result := r(ctx)
	return result
}

func Pure[S, T any](t T) State[S, T] {
	return func(s S) (fp.Try[S], fp.Try[T]) {
		return try.Success(s), try.Success(t)
	}
}

func MapState[S, A any](st State[S, A], f func(S) S) State[S, A] {
	return func(s S) (fp.Try[S], fp.Try[A]) {
		ns, a := st(s)
		return try.Map(ns, f), a
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

func MapValue[S, A, B any](st State[S, A], f func(A) B) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		return ns, try.Map(a, f)
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

func FlatMapValue[S, A, B any](st State[S, A], f func(A) fp.Try[B]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		return ns, try.FlatMap(a, f)
	}
}

func PeekState[S, A any](st State[S, A], f func(ctx S)) State[S, A] {
	return func(s S) (fp.Try[S], fp.Try[A]) {
		ns, r := st(s)
		ns.Foreach(f)
		return ns, r
	}
}

func MapWithState[S, A, B any](st State[S, A], f func(S, A) B) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		b := try.Map2(ns, a, f)
		return ns, b
	}
}

func FlatMapWithState[S, A, B any](st State[S, A], f func(S, A) fp.Try[B]) State[S, B] {
	return func(s S) (fp.Try[S], fp.Try[B]) {
		ns, a := st(s)
		b := try.Map2(ns, a, f)
		return ns, try.Flatten(b)
	}
}
