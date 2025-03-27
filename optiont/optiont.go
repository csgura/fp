package optiont

import (
	"iter"
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func Some[A any](v A) fp.OptionT[A] {
	return Pure(v)
}

func None[A any]() fp.OptionT[A] {
	return try.Success(option.None[A]())
}

func Failure[T any](err error) fp.OptionT[T] {
	return try.Failure[fp.Option[T]](err)
}

func Get[T any](v fp.OptionT[T]) fp.Try[T] {
	return try.FromOptionT(v)
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

func Of[T any](v T) fp.OptionT[T] {
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

func NonZero[T comparable](t T) fp.OptionT[T] {
	if t == fp.Zero[T]() {
		return None[T]()
	}
	return Some(t)
}

func NonEmptySlice[T ~[]E, E any](t T) fp.OptionT[T] {
	if len(t) == 0 {
		return None[T]()
	}
	return Some(t)
}

func FromTry[A any](v fp.Try[A]) fp.OptionT[A] {
	return try.Map(v, option.Some)
}

func Fold[T any, U any](optionT fp.OptionT[T], zero U, f func(U, T) U) U {
	if optionT.IsFailure() {
		return zero
	}
	return option.Fold(optionT.Get(), zero, f)
}

func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.OptionT[B]) fp.OptionT[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsSuccess() && t.Get().IsDefined() {
			sum = t.Get().Get()
		} else {
			return t
		}
	}
	return Pure(sum)
}

func Iterator[T any](optionT fp.OptionT[T]) fp.Iterator[T] {
	return iterator.FlatMap(try.Iterator(optionT), option.Iterator)
}

func All[T any](optionT fp.OptionT[T]) iter.Seq[T] {
	return Iterator(optionT).All()
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen
//go:generate go run github.com/csgura/fp/internal/generator/template_gen

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.OptionT[T]] {
	return genfp.GenerateMonadTransformer[fp.OptionT[T]]{
		File:     "optiont_op.go",
		TypeParm: genfp.TypeOf[T](),
		GivenMonad: genfp.MonadFunctions{
			Pure: try.Success[T],
		},
		ExposureMonad: genfp.MonadFunctions{
			Pure:    option.Pure[T],
			FlatMap: option.FlatMap[T, U],
		},
		Sequence: func(v fp.Option[fp.Try[T]]) fp.OptionT[T] {
			if v.IsDefined() {
				return try.Map(v.Get(), option.Some)
			}
			return try.Success(fp.Option[T]{})
		},
		Transform: []any{
			fp.Option[T].IsDefined,
			fp.Option[T].IsEmpty,
			fp.Option[T].Filter,
			fp.Option[T].OrElse,
			fp.Option[T].OrZero,
			fp.Option[T].OrElseGet,
			fp.Option[T].Or,
			fp.Option[T].OrOption,
			fp.Option[T].OrPtr,
			fp.Option[T].Recover,
			fp.Option[T].Foreach,
		},
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.OptionT[A]] {
	return genfp.GenerateMonadFunctions[fp.OptionT[A]]{
		File:     "optiont_monad.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// @internal.Generate
func _[A any]() genfp.GenerateTraverseFunctions[fp.OptionT[A]] {
	return genfp.GenerateTraverseFunctions[fp.OptionT[A]]{
		File:     "optiont_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/curried", Name: "curried"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}ApplicativeFunctor{{dec .N}}[{{TypeArgs 2 .N}}, R]{{end}}

type ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R any] struct {
	fn fp.OptionT[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} ApOptionT(a fp.OptionT[A1]) {{template "Next" .}} {
	return {{template "Next" .}}{Ap(r.fn, a)}
}

{{template "Receiver" .}} ApTry(a fp.Try[A1]) {{template "Next" .}} {
	return {{template "Next" .}}{Ap(r.fn, FromTry(a))}
}

{{template "Receiver" .}} ApTryAll({{DeclTypeClassArgs 1 .N "fp.Try"}}) fp.OptionT[R] {
	return r.
	{{- range (dec .N) -}}
		ApTry(ins{{inc .}}).
	{{- end -}}
		ApTry(ins{{.N}})
}

{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {
	return r.ApOptionT(fp.Success(a))
}

{{template "Receiver" .}} ApOptionAll({{DeclTypeClassArgs 1 .N "fp.Option"}}) fp.OptionT[R] {
	return r.
	{{- range (dec .N) -}}
		ApOption(ins{{inc .}}).
	{{- end -}}
		ApOption(ins{{.N}})
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {
	return r.ApOptionT(Some(a))
}

{{template "Receiver" .}} ApAll({{DeclArgs 1 .N}}) fp.OptionT[R] {
	return r.
	{{- range (dec .N) -}}
		Ap(a{{inc .}}).
	{{- end -}}
		Ap(a{{.N}})
}

{{template "Receiver" .}} ApOptionTFunc(a func() fp.OptionT[A1]) {{template "Next" .}} {
	return {{template "Next" .}}{ApFunc(r.fn, a)}
}

{{template "Receiver" .}} ApTryFunc(a func() fp.Try[A1]) {{template "Next" .}} {
	return r.ApOptionTFunc(func() fp.OptionT[A1] {
		return FromTry(a())
	})}

{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {
	return r.ApOptionTFunc(func() fp.OptionT[A1] {
		return fp.Success(a())
	})
}

{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {
	return r.ApOptionTFunc(func() fp.OptionT[A1] {
		return Some(a())
	})
}

func Applicative{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R] {
	return ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]{Some(curried.Func{{.N}}(fn))}
}
	`,
}

// @internal.Generate
func _[T any]() genfp.GenerateApplicative[fp.OptionT[T]] {
	return genfp.GenerateApplicative[fp.OptionT[T]]{
		File:     "applicative_gen.go",
		TypeParm: genfp.TypeOf[T](),
		Mapper: []genfp.Mapping{
			{
				Prefix: "Try",
				Mapper: FromTry[T],
			},
			{
				Prefix: "Option",
				Mapper: fp.Success[fp.Option[T]],
			},
			{
				Prefix: "Pure",
				Mapper: Some[T],
			},
		},
	}
}
