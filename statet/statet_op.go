package statet

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/unit"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
func _[S, A any]() genfp.GenerateMonadFunctions[fp.StateT[S, A]] {
	return genfp.GenerateMonadFunctions[fp.StateT[S, A]]{
		File:     "state_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @fp.Generate
func _[S, A any]() genfp.GenerateTraverseFunctions[fp.StateT[S, A]] {
	return genfp.GenerateTraverseFunctions[fp.StateT[S, A]]{
		File:     "state_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

func Run[S, A any](f func(S) (A, S)) fp.StateT[S, A] {
	return func(s S) (fp.Try[A], S) {
		ret, ns := f(s)
		return try.Success(ret), ns
	}
}

func Put[S any](s S) fp.StateT[S, fp.Unit] {
	return func(s S) (fp.Try[fp.Unit], S) {
		return unit.Success, s
	}
}

func PutWith[S, V any](withf func(S, V) S) func(v V) fp.StateT[S, fp.Unit] {
	return func(v V) fp.StateT[S, fp.Unit] {
		return func(s S) (fp.Try[fp.Unit], S) {
			return unit.Success, withf(s, v)
		}
	}
}

func Get[S any]() fp.StateT[S, S] {
	return func(s S) (fp.Try[S], S) {
		return try.Success(s), s
	}
}

func Modify[S any](f func(S) S) fp.StateT[S, fp.Unit] {
	return func(s S) (fp.Try[fp.Unit], S) {
		return unit.Success, f(s)
	}
}

func ModifyT[S any](f func(S) fp.Try[S]) fp.StateT[S, fp.Unit] {
	return func(s S) (fp.Try[fp.Unit], S) {
		nst := f(s)
		if nst.IsSuccess() {
			return unit.Success, nst.Get()
		}
		return try.Failure[fp.Unit](nst.Failed().Get()), s
	}
}

func GetS[S, A any](f func(S) A) fp.StateT[S, A] {
	return func(s S) (fp.Try[A], S) {
		return try.Success(f(s)), s
	}
}

func GetST[S, A any](f func(S) fp.Try[A]) fp.StateT[S, A] {
	return func(s S) (fp.Try[A], S) {
		return f(s), s
	}
}

func Pure[S, A any](a A) fp.StateT[S, A] {
	return func(s S) (fp.Try[A], S) {
		return try.Success(a), s
	}
}

func FromTry[S, A any](t fp.Try[A]) fp.StateT[S, A] {
	return func(s S) (fp.Try[A], S) {
		return t, s
	}
}

func FlatMap[S, A, B any](st fp.StateT[S, A], f func(A) fp.StateT[S, B]) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		ret, ns := st(s)
		if ret.IsSuccess() {
			return f(ret.Get())(ns)
		}
		return try.Failure[B](ret.Failed().Get()), ns
	}
}

func FlatMapConst[S, A, B any](st fp.StateT[S, A], next fp.StateT[S, B]) fp.StateT[S, B] {
	return FlatMap(st, fp.Const[A](next))
}

func WithState[S, A any](f func(S) fp.StateT[S, A]) fp.StateT[S, A] {
	return FlatMap(Get[S](), f)
}

func ApTry[S, A, B any](st fp.StateT[S, fp.Func1[A, B]], a fp.Try[A]) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		af, ns := st.Run(s)
		return try.Ap(af, a), ns
	}
}

func ApOption[S, A, B any](st fp.StateT[S, fp.Func1[A, B]], a fp.Option[A]) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		af, ns := st.Run(s)
		return try.Ap(af, try.FromOption(a)), ns
	}
}

func Transform[S, A, B any](st fp.StateT[S, A], f func(S, A) (S, fp.Try[B])) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		a, ns := st.Run(s)
		if a.IsSuccess() {
			nns, b := f(ns, a.Get())
			return b, nns
		}
		return try.Failure[B](a.Failed().Get()), ns
	}
}

func MapWithState[S, A, B any](st fp.StateT[S, A], f func(S, A) B) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		a, ns := st.Run(s)
		return try.Map2(try.Success(ns), a, f), ns
	}
}

func MapT[S, A, B any](st fp.StateT[S, A], f func(A) fp.Try[B]) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		a, ns := st.Run(s)
		return try.FlatMap(a, f), ns
	}
}

func MapWithStateT[S, A, B any](st fp.StateT[S, A], f func(S, A) fp.Try[B]) fp.StateT[S, B] {
	return func(s S) (fp.Try[B], S) {
		a, ns := st.Run(s)
		return try.LiftM2(f)(try.Success(ns), a), ns
	}
}

func PeekState[S, A any](st fp.StateT[S, A], f func(ctx S)) fp.StateT[S, A] {
	return func(s S) (fp.Try[A], S) {
		r, ns := st.Run(s)
		f(ns)
		return r, ns
	}
}

// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldM[S, A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.StateT[S, B]) fp.StateT[S, B] {
	sum := Pure[S](zero)
	for s.HasNext() {
		na := s.Next()
		sum = FlatMap(sum, func(b B) fp.StateT[S, B] {
			return f(b, na)
		})
	}
	return sum
}
