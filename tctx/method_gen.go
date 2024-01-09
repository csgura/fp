// Code generated by template_gen, DO NOT EDIT.
package tctx

import (
	"context"
	"github.com/csgura/fp"
	"github.com/csgura/fp/tstate"
)

func MapMethodWith1[A1, A2, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2) R, a2 A2) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2)
	}))
}

func MapMethodWithT1[A1, A2, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2) fp.Try[R], a2 A2) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2)
	}))
}

func MapMethodWith2[A1, A2, A3, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3) R, a2 A2, a3 A3) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3)
	}))
}

func MapMethodWithT2[A1, A2, A3, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3) fp.Try[R], a2 A2, a3 A3) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3)
	}))
}

func MapMethodWith3[A1, A2, A3, A4, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4) R, a2 A2, a3 A3, a4 A4) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3, a4)
	}))
}

func MapMethodWithT3[A1, A2, A3, A4, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4) fp.Try[R], a2 A2, a3 A3, a4 A4) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3, a4)
	}))
}

func MapMethodWith4[A1, A2, A3, A4, A5, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5) R, a2 A2, a3 A3, a4 A4, a5 A5) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3, a4, a5)
	}))
}

func MapMethodWithT4[A1, A2, A3, A4, A5, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R], a2 A2, a3 A3, a4 A4, a5 A5) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3, a4, a5)
	}))
}

func MapMethodWith5[A1, A2, A3, A4, A5, A6, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3, a4, a5, a6)
	}))
}

func MapMethodWithT5[A1, A2, A3, A4, A5, A6, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R], a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3, a4, a5, a6)
	}))
}

func MapMethodWith6[A1, A2, A3, A4, A5, A6, A7, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3, a4, a5, a6, a7)
	}))
}

func MapMethodWithT6[A1, A2, A3, A4, A5, A6, A7, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R], a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3, a4, a5, a6, a7)
	}))
}

func MapMethodWith7[A1, A2, A3, A4, A5, A6, A7, A8, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3, a4, a5, a6, a7, a8)
	}))
}

func MapMethodWithT7[A1, A2, A3, A4, A5, A6, A7, A8, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R], a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3, a4, a5, a6, a7, a8)
	}))
}

func MapMethodWith8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1, s, a2, a3, a4, a5, a6, a7, a8, a9)
	}))
}

func MapMethodWithT8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](s State[A1], f func(a1 A1, ctx context.Context, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R], a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1, s, a2, a3, a4, a5, a6, a7, a8, a9)
	}))
}