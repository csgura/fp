//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package option

import (
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
)

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.Option[A]] {
	return genfp.GenerateMonadFunctions[fp.Option[A]]{
		File:     "option_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateTraverseFunctions[fp.Option[A]] {
	return genfp.GenerateTraverseFunctions[fp.Option[A]]{
		File:     "option_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

func Pure[T any](v T) fp.Option[T] {
	return Some(v)
}

func Some[T any](v T) fp.Option[T] {
	return fp.None[T]().Recover(func() T {
		return v
	})
}

func None[T any]() fp.Option[T] {
	return fp.None[T]()
}

// 아규먼트를 무시하고 항상 None 을 리턴
func ConstNone[A, B any](a A) fp.Option[B] {
	return fp.None[B]()
}

func isNil(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func Of[T any](v T) fp.Option[T] {
	var i any = v
	if i == nil {
		return None[T]()
	}

	rv := reflect.ValueOf(i)
	if isNil(rv) {
		return None[T]()
	}
	return Some(v)
}

func Ptr[T any](v *T) fp.Option[T] {
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func String(v string) fp.Option[string] {
	return NonZero(v)
}

func NonZero[T comparable](t T) fp.Option[T] {
	if t == fp.Zero[T]() {
		return None[T]()
	}
	return Some(t)
}

func NonEmptySlice[T ~[]E, E any](t T) fp.Option[T] {
	if t == nil {
		return None[T]()
	}
	return Some(t)
}

func FromTry[T any](t fp.Try[T]) fp.Option[T] {
	if t.IsSuccess() {
		return Some(t.Get())
	}
	return None[T]()
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Option[B] {
	return fp.Compose(fab, Some)
}

func FlatMap[T, U any](opt fp.Option[T], fn func(v T) fp.Option[U]) fp.Option[U] {
	if opt.IsDefined() {
		return fn(opt.Get())
	}
	return None[U]()
}

func FlatPtr[T any](opt fp.Option[*T]) fp.Option[T] {
	return FlatMap(opt, func(v *T) fp.Option[T] {
		return Ptr(v)
	})
}

func Fold[A, B any](s fp.Option[A], zero B, f func(B, A) B) B {
	if s.IsEmpty() {
		return zero
	}

	return f(zero, s.Get())
}

func FoldRight[A, B any](s fp.Option[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	return f(s.Get(), lazy.Done(zero))
}

// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.Option[B]) fp.Option[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsDefined() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return Pure(sum)
}

func ToSeq[T any](r fp.Option[T]) fp.Seq[T] {
	if r.IsDefined() {
		return fp.Seq[T]{r.Get()}
	}
	return nil
}

func Iterator[T any](r fp.Option[T]) fp.Iterator[T] {
	return fp.IteratorOfSeq(ToSeq(r))
}

func Deref[R any, T fp.Deref[R]](opt fp.Option[T]) fp.Option[R] {
	return Map(opt, T.Deref)
}

type MonadChain1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A, R]]
}

func (r MonadChain1[H, HT, A, R]) Map(a func(HT) A) fp.Option[R] {
	return r.FlatMap(func(h HT) fp.Option[A] {
		return Some(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListMap(a func(H) A) fp.Option[R] {
	return r.HListFlatMap(func(h H) fp.Option[A] {
		return Some(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListFlatMap(a func(H) fp.Option[A]) fp.Option[R] {
	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a(v)
	})

	return r.ApOption(av)
}

func (r MonadChain1[H, HT, A, R]) FlatMap(a func(HT) fp.Option[A]) fp.Option[R] {
	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a(v.Head())
	})

	return r.ApOption(av)
}

func (r MonadChain1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Option[R] {
	return Ap(r.fn, a)
}

func (r MonadChain1[H, HT, A, R]) Ap(a A) fp.Option[R] {
	return r.ApOption(Some(a))
}

func (r MonadChain1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Option[R] {

	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a()
	})
	return r.ApOption(av)
}
func (r MonadChain1[H, HT, A, R]) ApFunc(a func() A) fp.Option[R] {

	av := Map(r.h, func(v H) A {
		return a()
	})
	return r.ApOption(av)
}

func Chain1[A, R any](fn fp.Func1[A, R]) MonadChain1[hlist.Nil, hlist.Nil, A, R] {
	return MonadChain1[hlist.Nil, hlist.Nil, A, R]{Some(hlist.Empty()), Some(fn)}
}

type ApplicativeFunctor1[A, R any] struct {
	fn fp.Option[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[A, R]) ApOption(a fp.Option[A]) fp.Option[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[A, R]) Ap(a A) fp.Option[R] {
	return r.ApOption(Some(a))
}

func (r ApplicativeFunctor1[A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Option[R] {

	return ApFunc(r.fn, a)
}
func (r ApplicativeFunctor1[A, R]) ApFunc(a func() A) fp.Option[R] {
	return ApFunc(r.fn, func() fp.Option[A] {
		return Some(a())
	})
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[A, R] {
	return ApplicativeFunctor1[A, R]{fn: Some(fn)}
}

func Pure0[R any](f func() R) fp.Func1[fp.Unit, fp.Option[R]] {
	return func(fp.Unit) fp.Option[R] {
		return Some(f())
	}
}

func Pure1[A, R any](f func(A) R) fp.Func1[A, fp.Option[R]] {
	return func(a A) fp.Option[R] {
		return Some(f(a))
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/curried", Name: "curried"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r MonadChain{{.N}}[H, HT, {{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}MonadChain{{dec .N}}[hlist.Cons[A1, H], {{TypeArgs 1 .N}}, R]{{end}}


type MonadChain{{.N}}[H hlist.Header[HT], HT, {{TypeArgs 1 .N}}, R any] struct {
	h  fp.Option[H]
	fn fp.Option[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} FlatMap(a func(HT) fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
{{template "Receiver" .}} Map(a func(HT) A1) {{template "Next" .}} {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
{{template "Receiver" .}} HListMap(a func(H) A1) {{template "Next" .}} {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
{{template "Receiver" .}} HListFlatMap(a func(H) fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}

{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return {{template "Next" .}}{nh, Ap(r.fn, a)}
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApOption(Some(a))

}

{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a()
	})
	return r.ApOption(av)
}

{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApOption(av)
}

func Chain{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R] {
	return MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R]{Some(hlist.Empty()), Some(curried.Func{{.N}}(fn))}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/curried", Name: "curried"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}ApplicativeFunctor{{dec .N}}[{{TypeArgs 2 .N}}, R]{{end}}

type ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R any] struct {
	fn fp.Option[{{CurriedFunc 1 .N "R"}}]
}


{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	return {{template "Next" .}}{Ap(r.fn, a)}
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApOption(Some(a))

}

{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {

	return {{template "Next" .}}{ApFunc(r.fn, a)}

}

{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {

	return r.ApOptionFunc(func() fp.Option[A1] {
		return Some(a())
	})
}
func Applicative{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R] {
	return ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]{Some(curried.Func{{.N}}(fn))}
}
	`,
}
