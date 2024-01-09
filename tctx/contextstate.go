package tctx

import (
	"context"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/tstate"
)

// State[A]
// A 의 메소드 호출 ( Context  포함 )   func (r A) Method( ctx context.Context,  b B, c C ) R
// A 의 메소드 호출 ( Context  미포함 )   func (r A) Method( b B, c C ) R

// A 아규먼트 받는 함수 호출
// Context 있는 두번째 아규먼트  func Func(ctx contex.Context, a A , b B , c C ) R  => easy
// Context 있는 다른 위치 아규먼트  func Func(ctx contex.Context, b B, a A , c C ) R

// Context 없는 첫번째 아규먼트  func Func( a A , b B , c C ) R  => easy
// Context 없는 다른 위치 아규먼트  func Func( b B, a A , c C ) R

type State[A any] tstate.State[context.Context, A]

func (r State[A]) Run(ctx context.Context) (fp.Try[context.Context], fp.Try[A]) {
	return r(ctx)
}

func (r State[A]) Exec(ctx context.Context) fp.Try[context.Context] {
	state, _ := r(ctx)
	return state
}

func (r State[A]) Eval(ctx context.Context) fp.Try[A] {
	_, result := r(ctx)
	return result
}

func Pure[T any](t T) State[T] {
	return Narrow(tstate.Pure[context.Context](t))
}

func FromTry[T any](t fp.Try[T]) State[T] {
	return Narrow(tstate.FromTry[context.Context](t))
}

func Of[T any](f func(ctx context.Context) fp.Try[T]) State[T] {
	return func(ctx context.Context) (fp.Try[context.Context], fp.Try[T]) {
		rt := f(ctx)
		return try.Success(ctx), rt
	}
}

type WithFunc[A, R any] func(context.Context, A) R

func Compose2[A, B, R any](f1 WithFunc[A, fp.Try[B]], f2 WithFunc[B, fp.Try[R]]) WithFunc[A, fp.Try[R]] {
	return func(ctx context.Context, a A) fp.Try[R] {
		return try.FlatMap(f1(ctx, a), as.Func2(f2).ApplyFirst(ctx))
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxCompose,
	Template: `
func Compose{{.N}}[{{TypeArgs 1 .N}}, R any]({{(Monad "fp.Try").FuncChain 1 .N "WithFunc"}}) WithFunc[A1,fp.Try[R]] {
	return Compose2(f1, Compose{{dec .N}}({{CallArgs 2 .N "f"}}))
}
	`,
}

func Ap[A, B any](s State[fp.Func1[A, B]], a A) State[B] {
	return Narrow(tstate.Ap(Widen(s), a))
}

func ApTry[A, B any](s State[fp.Func1[A, B]], a fp.Try[A]) State[B] {
	return Narrow(tstate.ApTry(Widen(s), a))
}

func ApOption[A, B any](s State[fp.Func1[A, B]], a fp.Option[A]) State[B] {
	return Narrow(tstate.ApOption(Widen(s), a))
}

// widen 의 의미는 B which extends A -> A ( super type )
// 더 일반적인 타입으로 변환하는 것을 의미
func Widen[A any](s State[A]) tstate.State[context.Context, A] {
	return tstate.State[context.Context, A](s)
}

// narrow 의 의미는  A -> B which extends A  ( sub type )
// 더 상세한 타입으로 변경하는 것을 의미
func Narrow[A any](s tstate.State[context.Context, A]) State[A] {
	return State[A](s)
}

func WithContext[A any](s State[A], f func(context.Context) context.Context) State[A] {
	return Narrow(tstate.WithState(Widen(s), f))
}

func WithValue[A any](s State[A], k any, v any) State[A] {
	return WithContext(s, func(ctx context.Context) context.Context {
		return context.WithValue(ctx, k, v)
	})
}

func Map[A, B any](s State[A], f func(A) B) State[B] {
	return Narrow(tstate.Map(Widen(s), f))
}

func MapWith[A, B any](s State[A], f func(context.Context, A) B) State[B] {
	return Narrow(tstate.MapWithState(Widen(s), f))
}

func MapMethodWith[A, B any](s State[A], f func(A, context.Context) B) State[B] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a A) B {
		return f(a, s)
	}))
}

func MapMethodWithT[A, B any](s State[A], f func(A, context.Context) fp.Try[B]) State[B] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a A) fp.Try[B] {
		return f(a, s)
	}))
}

func Inspect[A, B any](s State[A], f func(context.Context) B) State[B] {
	return Narrow(tstate.Inspect(Widen(s), f))
}

func JoinT[A, B any](s State[fp.Try[A]]) State[A] {
	return Narrow(tstate.MapT(Widen(s), fp.Id))
}

func MapT[A, B any](s State[A], f func(A) fp.Try[B]) State[B] {
	return Narrow(tstate.MapT(Widen(s), f))
}

func MapWithT[A, B any](s State[A], f func(context.Context, A) fp.Try[B]) State[B] {
	return Narrow(tstate.MapWithStateT(Widen(s), f))
}

func PeekContext[A any](s State[A], f func(ctx context.Context)) State[A] {
	return Narrow(tstate.PeekState(Widen(s), f))
}

func Const[A any](a A) fp.Func1[context.Context, A] {
	return func(a1 context.Context) A {
		return a
	}
}

func AsWithFunc[A1, R any](f fp.Func1[context.Context, fp.Func1[A1, R]]) func(context.Context, A1) R {
	return func(s context.Context, a1 A1) R {
		return f(s)(a1)
	}
}

// func Curried3[A1, A2, R any](f func(context.Context, A1, A2) R) State[fp.Func1[A1, fp.Func1[A2, R]]] {
// 	return func(ctx context.Context) (fp.Try[context.Context], fp.Try[fp.Func1[A1, fp.Func1[A2, R]]]) {
// 		ret := as.Curried3(f)(ctx)
// 		return try.Success(ctx), try.Success(ret)
// 	}
// }

// func Curried4[A1, A2, A3, R any](f func(context.Context, A1, A2, A3) R) State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]] {
// 	return func(ctx context.Context) (fp.Try[context.Context], fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]) {
// 		ret := as.Curried4(f)(ctx)
// 		return try.Success(ctx), try.Success(ret)
// 	}
// }

//go:generate go run github.com/csgura/fp/internal/generator/template_gen

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "from_func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/try", Name: "try"},

		{Package: "context", Name: "context"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func FromFunc{{.N}}[{{TypeArgs 1 (dec .N)}}, R any](f func(context.Context, {{TypeArgs 1 (dec .N)}}) fp.Try[R], {{DeclArgs 1 (dec .N)}}) State[R] {
	return func(ctx context.Context) (fp.Try[context.Context], fp.Try[R]) {
		r := f(ctx, {{CallArgs 1 (dec .N)}})
		return try.Success(ctx), r
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "curried_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
		{Package: "context", Name: "context"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func SlipL{{.N}}[{{TypeArgs 1 (dec .N)}}, R any](f fp.Func1[context.Context,{{CurriedFunc 1 (dec .N) "R"}}]) fp.Func1[context.Context, fp.Func1[A{{dec .N}}, {{CurriedFunc 1 (dec ( dec .N )) "R"}}]] {
	return as.Curried{{.N}}(func(ctx context.Context, a{{dec .N}} A{{dec .N}}, {{DeclArgs 1 (dec (dec .N))}}) R {
		return f(ctx){{CurriedCallArgs 1 (dec .N)}}

	})
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "curried_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/tstate", Name: "tstate"},
		{Package: "context", Name: "context"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `



func MapT{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f {{CurriedFunc 1 .N "fp.Try[R]"}}, {{DeclArgs 2 .N}}) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f{{CurriedCallArgs 1 .N}}
	}))
}


func AsWithFunc{{dec .N}}[{{TypeArgs 1 .N}}, R any](f fp.Func1[context.Context, {{CurriedFunc 1 .N "R"}}], {{DeclArgs 2 .N}}) func(context.Context, A1) R {
	return func(s context.Context, a1 A1) R {
		return f(s){{CurriedCallArgs 1 .N}}
	}
}

func MapWithT{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f fp.Func1[context.Context, {{CurriedFunc 1 .N "fp.Try[R]"}}], {{DeclArgs 2 .N}}) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(s){{CurriedCallArgs 1 .N}}
	}))
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "method_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/tstate", Name: "tstate"},
		{Package: "context", Name: "context"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `


func MapMethodWith{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f func(a1 A1, ctx context.Context, {{DeclArgs 2 .N}}) R,  {{DeclArgs 2 .N}}) State[R] {
	return Narrow(tstate.MapWithState(Widen(s), func(s context.Context, a1 A1) R {
		return f(a1 , s , {{CallArgs 2 .N}})
	}))
}

func MapMethodWithT{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f func(a1 A1, ctx context.Context, {{DeclArgs 2 .N}}) fp.Try[R],  {{DeclArgs 2 .N}}) State[R] {
	return Narrow(tstate.MapWithStateT(Widen(s), func(s context.Context, a1 A1) fp.Try[R] {
		return f(a1 , s , {{CallArgs 2 .N}})
	}))
}
	`,
}
