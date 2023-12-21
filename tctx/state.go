package tctx

import (
	"context"

	"github.com/csgura/fp"
	"github.com/csgura/fp/tstate"
)

type State[A any] tstate.State[context.Context, A]

func (r State[A]) Run(ctx context.Context) fp.Try[A] {
	_, result := r(ctx)
	return result
}

func Pure[T any](t T) State[T] {
	return Narrow(tstate.Pure[context.Context](t))
}

func Widen[A any](s State[A]) tstate.State[context.Context, A] {
	return tstate.State[context.Context, A](s)
}

func Narrow[A any](s tstate.State[context.Context, A]) State[A] {
	return State[A](s)
}

func ModifyContext[A any](s State[A], f func(context.Context) context.Context) State[A] {
	return Narrow(tstate.MapState(Widen(s), f))
}

func WithValue[A any](s State[A], k any, v any) State[A] {
	return ModifyContext(s, func(ctx context.Context) context.Context {
		return context.WithValue(ctx, k, v)
	})
}

func Map[A, B any](s State[A], f func(A) B) State[B] {
	return Narrow(tstate.MapValue(Widen(s), f))

}

func FlatMap[A, B any](s State[A], f func(A) fp.Try[B]) State[B] {
	return Narrow(tstate.FlatMapValue(Widen(s), f))

}

func PeekContext[A any](s State[A], f func(ctx context.Context)) State[A] {
	return Narrow(tstate.PeekState(Widen(s), f))
}
