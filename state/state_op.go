package state

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/try"
)

func Run[S, A any](f func(S) (A, S)) fp.State[S, A] {
	return func(s S) fp.Tuple2[A, S] {
		return as.Tuple(f(s))
	}
}

func Put[S any](s S) fp.State[S, fp.Unit] {
	return func(s S) fp.Tuple2[fp.Unit, S] {
		return as.Tuple(fp.Unit{}, s)
	}
}

func PutWith[C, V any](withf func(C, V) C) func(v V) fp.State[C, fp.Unit] {
	return func(v V) fp.State[C, fp.Unit] {
		return func(c C) fp.Tuple2[fp.Unit, C] {
			return as.Tuple(fp.Unit{}, withf(c, v))
		}
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

func GetS[S, B any](f func(S) B) fp.State[S, B] {
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

// func WithState[S, A any](st fp.State[S, A], f func(S) S) fp.State[S, A] {
// 	return func(s S) fp.Tuple2[A, S] {
// 		v := st(s)
// 		return product.MapValue(v, f)
// 	}
// }

func WithState[S, A any](f func(S) fp.State[S, A]) fp.State[S, A] {
	return FlatMap(Get[S](), f)
}

func MapWithState[S, A, B any](st fp.State[S, A], f func(S, A) B) fp.State[S, B] {
	return func(s S) fp.Tuple2[B, S] {
		a, ns := st.Run(s)
		return as.Tuple2(f(ns, a), ns)
	}
}

// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldM[S, A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.State[S, B]) fp.State[S, B] {
	sum := Pure[S](zero)
	for s.HasNext() {
		na := s.Next()
		sum = FlatMap(sum, func(b B) fp.State[S, B] {
			return f(b, na)
		})
	}
	return sum
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
func _[S, A any]() genfp.GenerateMonadFunctions[fp.State[S, A]] {
	return genfp.GenerateMonadFunctions[fp.State[S, A]]{
		File:     "state_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @fp.Generate
func _[S, A any]() genfp.GenerateTraverseFunctions[fp.State[S, A]] {
	return genfp.GenerateTraverseFunctions[fp.State[S, A]]{
		File:     "state_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

func PeekState[S, A any](st fp.State[S, A], f func(ctx S)) fp.State[S, A] {
	return func(s S) fp.Tuple2[A, S] {
		r, ns := st.Run(s)
		f(ns)
		return as.Tuple(r, ns)
	}
}
