// Code generated by template_gen, DO NOT EDIT.
package tctx

import (
	"context"
	"github.com/csgura/fp"
)

func EvalT2[A1, R any](f func(context.Context, A1) fp.Try[R], a1 A1) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1)
	})
}

func EvalT3[A1, A2, R any](f func(context.Context, A1, A2) fp.Try[R], a1 A1, a2 A2) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2)
	})
}

func EvalT4[A1, A2, A3, R any](f func(context.Context, A1, A2, A3) fp.Try[R], a1 A1, a2 A2, a3 A3) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2, a3)
	})
}

func EvalT5[A1, A2, A3, A4, R any](f func(context.Context, A1, A2, A3, A4) fp.Try[R], a1 A1, a2 A2, a3 A3, a4 A4) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2, a3, a4)
	})
}

func EvalT6[A1, A2, A3, A4, A5, R any](f func(context.Context, A1, A2, A3, A4, A5) fp.Try[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2, a3, a4, a5)
	})
}

func EvalT7[A1, A2, A3, A4, A5, A6, R any](f func(context.Context, A1, A2, A3, A4, A5, A6) fp.Try[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2, a3, a4, a5, a6)
	})
}

func EvalT8[A1, A2, A3, A4, A5, A6, A7, R any](f func(context.Context, A1, A2, A3, A4, A5, A6, A7) fp.Try[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2, a3, a4, a5, a6, a7)
	})
}

func EvalT9[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) fp.Try[R], a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
	})
}
