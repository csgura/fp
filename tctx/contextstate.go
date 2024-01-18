package tctx

import (
	"context"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/state"
	"github.com/csgura/fp/try"
)

// State[A]
// A 의 메소드 호출 ( Context  포함 )   func (r A) Method( ctx context.Context,  b B, c C ) R
// A 의 메소드 호출 ( Context  미포함 )   func (r A) Method( b B, c C ) R

// A 아규먼트 받는 함수 호출
// Context 있는 두번째 아규먼트  func Func(ctx contex.Context, a A , b B , c C ) R  => easy
// Context 있는 다른 위치 아규먼트  func Func(ctx contex.Context, b B, a A , c C ) R

// Context 없는 첫번째 아규먼트  func Func( a A , b B , c C ) R  => easy
// Context 없는 다른 위치 아규먼트  func Func( b B, a A , c C ) R

type Context struct {
	ctx  context.Context
	cncl func()
}

func (r Context) WithContext(nc context.Context) Context {
	r.ctx = nc
	return r
}

type State[A any] fp.State[Context, fp.Try[A]]

func (r State[A]) Run(ctx context.Context) (fp.Try[A], Context) {
	return r(Context{ctx: ctx}).Unapply()
}

func (r State[A]) Exec(ctx context.Context) Context {
	_, state := r.Run(ctx)
	return state
}

func (r State[A]) Eval(ctx context.Context) fp.Try[A] {
	result, _ := r.Run(ctx)
	return result
}

func Run[A any](f func(context.Context) (context.Context, A)) State[A] {
	return func(ctx Context) fp.Tuple2[fp.Try[A], Context] {
		nc, a := f(ctx.ctx)
		return as.Tuple(try.Success(a), ctx.WithContext(nc))
	}
}

func RunT[A any](f func(context.Context) (context.Context, fp.Try[A])) State[A] {
	return func(ctx Context) fp.Tuple2[fp.Try[A], Context] {
		nc, ta := f(ctx.ctx)
		return as.Tuple(ta, ctx.WithContext(nc))
	}
}

func Eval[A any](f func(context.Context) A) State[A] {
	return func(ctx Context) fp.Tuple2[fp.Try[A], Context] {
		return as.Tuple(try.Success(f(ctx.ctx)), ctx)
	}
}

func EvalT[A any](f func(context.Context) fp.Try[A]) State[A] {
	return func(ctx Context) fp.Tuple2[fp.Try[A], Context] {
		return as.Tuple(f(ctx.ctx), ctx)
	}
}

func Modify(f func(context.Context) context.Context) State[fp.Unit] {
	return func(ctx Context) fp.Tuple2[fp.Try[fp.Unit], Context] {
		return as.Tuple(try.Unit, ctx.WithContext(f(ctx.ctx)))
	}
}

// narrow 의 의미는  A -> B which extends A  ( sub type )
// 더 상세한 타입으로 변경하는 것을 의미
func Narrow[A any](s fp.State[Context, fp.Try[A]]) State[A] {
	return State[A](s)
}

func Pure[T any](t T) State[T] {
	return Narrow(state.Pure[Context](try.Success(t)))
}

func FlatMap[A, B any](sa State[A], f func(A) State[B]) State[B] {
	return func(ctx Context) fp.Tuple2[fp.Try[B], Context] {
		ta, nctx := sa(ctx).Unapply()
		if ta.IsFailure() {
			return as.Tuple(try.Failure[B](ta.Failed().Get()), nctx)
		}
		return f(ta.Get())(nctx)
	}
}

func FromTry[T any](t fp.Try[T]) State[T] {
	return Narrow(state.Pure[Context](t))
}

func Of[T any](f func(ctx context.Context) fp.Try[T]) State[T] {

	return func(ctx Context) fp.Tuple2[fp.Try[T], Context] {
		rt := f(ctx.ctx)
		return as.Tuple(rt, ctx)
	}
}

type WithFunc[A, R any] func(context.Context, A) R

func Compose2[A, B, R any](f1 WithFunc[A, fp.Try[B]], f2 WithFunc[B, fp.Try[R]]) WithFunc[A, fp.Try[R]] {
	return func(ctx context.Context, a A) fp.Try[R] {
		return try.FlatMap(f1(ctx, a), as.Func2(f2).ApplyFirst(ctx))
	}
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[S, A any]() genfp.GenerateMonadFunctions[State[A]] {
	return genfp.GenerateMonadFunctions[State[A]]{
		File:     "tctx_monad.go",
		TypeParm: genfp.TypeOf[A](),
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

func ApTry[A, B any](s State[fp.Func1[A, B]], a fp.Try[A]) State[B] {
	return Ap(s, FromTry(a))
}

func ApOption[A, B any](s State[fp.Func1[A, B]], a fp.Option[A]) State[B] {
	return Ap(s, FromTry(try.FromOption(a)))
}

// widen 의 의미는 B which extends A -> A ( super type )
// 더 일반적인 타입으로 변환하는 것을 의미
func Widen[A any](s State[A]) fp.State[Context, fp.Try[A]] {
	return fp.State[Context, fp.Try[A]](s)
}

func WithContext[A any](s State[A], f func(context.Context) context.Context) State[A] {
	return Narrow(state.WithState(Widen(s), func(ctx Context) Context {
		return ctx.WithContext(f(ctx.ctx))
	}))
}

func WithValue[A any](s State[A], k any, v any) State[A] {
	return WithContext(s, func(ctx context.Context) context.Context {
		return context.WithValue(ctx, k, v)
	})
}

func MapWith[A, B any](s State[A], f func(context.Context, A) B) State[B] {
	wsa := Widen(s)
	wsb := state.MapWithState(wsa, func(s Context, a fp.Try[A]) fp.Try[B] {
		return try.Map(a, as.Curried2(f)(s.ctx))
	})

	return Narrow(wsb)
}

func MapWithT[A, B any](s State[A], f func(context.Context, A) fp.Try[B]) State[B] {
	wsa := Widen(s)
	wsb := state.MapWithState(wsa, func(s Context, a fp.Try[A]) fp.Try[B] {
		return try.FlatMap(a, as.Curried2(f)(s.ctx))
	})

	return Narrow(wsb)
}

func MapMethodWith[A, B any](s State[A], f func(A, context.Context) B) State[B] {
	wsa := Widen(s)
	wsb := state.MapWithState(wsa, func(s Context, a fp.Try[A]) fp.Try[B] {
		return try.Map(a, as.Func2(f).ApplyLast(s.ctx))
	})

	return Narrow(wsb)
}

func MapMethodWithT[A, B any](s State[A], f func(A, context.Context) fp.Try[B]) State[B] {
	wsa := Widen(s)
	wsb := state.MapWithState(wsa, func(s Context, a fp.Try[A]) fp.Try[B] {
		return try.FlatMap(a, as.Func2(f).ApplyLast(s.ctx))
	})

	return Narrow(wsb)
}

func JoinT[A, B any](s State[fp.Try[A]]) State[A] {
	return MapT(s, fp.Id)
}

func MapT[A, B any](s State[A], f func(A) fp.Try[B]) State[B] {
	wsa := Widen(s)
	wsb := state.Map(wsa, func(a fp.Try[A]) fp.Try[B] {
		return try.FlatMap(a, f)
	})

	return Narrow(wsb)
}

func PeekContext[A any](s State[A], f func(ctx context.Context)) State[A] {
	return Narrow(state.PeekState(Widen(s), func(c Context) {
		f(c.ctx)
	}))
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

		{Package: "context", Name: "context"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
func EvalT{{.N}}[{{TypeArgs 1 (dec .N)}}, R any](f func(context.Context, {{TypeArgs 1 (dec .N)}}) fp.Try[R], {{DeclArgs 1 (dec .N)}}) State[R] {
	return EvalT(func(ctx context.Context) fp.Try[R] {
		return f(ctx, {{CallArgs 1 (dec .N)}})
	})
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
		{Package: "context", Name: "context"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `



func MapT{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f {{CurriedFunc 1 .N "fp.Try[R]"}}, {{DeclArgs 2 .N}}) State[R] {
	return MapT(s, func(a1 A1) fp.Try[R] {
		return f{{CurriedCallArgs 1 .N}}
	})
}


func AsWithFunc{{dec .N}}[{{TypeArgs 1 .N}}, R any](f fp.Func1[context.Context, {{CurriedFunc 1 .N "R"}}], {{DeclArgs 2 .N}}) func(context.Context, A1) R {
	return func(s context.Context, a1 A1) R {
		return f(s){{CurriedCallArgs 1 .N}}
	}
}

func MapWithT{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f fp.Func1[context.Context, {{CurriedFunc 1 .N "fp.Try[R]"}}], {{DeclArgs 2 .N}}) State[R] {
	return MapWithT(s, func(s context.Context, a1 A1) fp.Try[R] {
		return f(s){{CurriedCallArgs 1 .N}}
	})
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "method_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "context", Name: "context"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `


func MapMethodWith{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f func(a1 A1, ctx context.Context, {{DeclArgs 2 .N}}) R,  {{DeclArgs 2 .N}}) State[R] {
	return MapMethodWith(s , func(a1 A1, s context.Context) R {
		return f(a1 , s , {{CallArgs 2 .N}})
	})
}

func MapMethodWithT{{dec .N}}[{{TypeArgs 1 .N}}, R any](s State[A1], f func(a1 A1, ctx context.Context, {{DeclArgs 2 .N}}) fp.Try[R],  {{DeclArgs 2 .N}}) State[R] {
	return MapMethodWithT(s, func(a1 A1, s context.Context) fp.Try[R] {
		return f(a1 , s , {{CallArgs 2 .N}})
	})
}
	`,
}
