package cs

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

type GuardType bool

const Guard GuardType = true

func If[T ~bool](f func() T) GuardType {
	return GuardType(try.Of(f).OrZero())
}

func NoMatter[T comparable](v T) T {
	return v
}

func Is[T ~bool](v T) T {
	return true
}

func No[T ~bool](v T) T {
	return false
}

func Has[T ~bool](v T) T {
	return true
}

func Not[T ~bool](v T) T {
	return false
}

func IsSome[T comparable](v fp.Option[T]) fp.Option[T] {
	return option.IsSomeCase(v)
}

func IsSomeAnd[T comparable](v fp.Option[T], nested ...fp.Endo[T]) fp.Option[T] {
	return option.IsSomeCaseAnd(v, nested...)
}

func NestedIsSome[T comparable](nested ...fp.Endo[T]) fp.Endo[fp.Option[T]] {
	return option.NestedIsSomeCase(nested...)
}

func IsNone[T comparable](v fp.Option[T]) fp.Option[T] {
	return option.IsNoneCase(v)
}

func IsSuccess[T comparable](v fp.Try[T]) fp.Try[T] {
	return try.IsSuccessCaseAnd(v)
}

func IsFailure[T comparable](v fp.Try[T]) fp.Try[T] {
	return try.IsFailureCaseAnd(v)
}

func IsSuccessAnd[T comparable](v fp.Try[T], nested ...fp.Endo[T]) fp.Try[T] {
	return try.IsSuccessCaseAnd(v, nested...)
}

func NestedIsSuccess[T comparable](nested ...fp.Endo[T]) fp.Endo[fp.Try[T]] {
	return try.NestedIsSuccessCase(nested...)
}

func IsFailureCaseAnd[T comparable](v fp.Try[T], nested ...fp.Endo[error]) fp.Try[T] {
	return try.NestedIsFailureCase[T](nested...)(v)
}

func NestedIsFailure[T comparable](nested ...fp.Endo[error]) fp.Endo[fp.Try[T]] {
	return try.NestedIsFailureCase[T](nested...)
}

//go:generate go run github.com/csgura/fp/internal/generator/template_gen

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  1,
	Until: genfp.MaxProduct,
	Template: `
func Match{{.N}}[{{TypeArgs 1 .N}} any]({{DeclArgs 1 .N}}) fp.{{TupleType .N}} {
	return fp.{{TupleType .N}}{
	{{- range $idx := Range 1 .N}}
		I{{$idx}}: a{{$idx}},
	{{- end}}
	}
}
`,
}
