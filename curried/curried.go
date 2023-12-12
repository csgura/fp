//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package curried

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

func Func1[A, R any](f func(A) R) fp.Func1[A, R] {
	return fp.Func1[A, R](f)
}

func Flip[A, B, R any](f fp.Func1[A, fp.Func1[B, R]]) fp.Func1[B, fp.Func1[A, R]] {
	return func(b B) fp.Func1[A, R] {
		return func(a A) R {
			return f(a)(b)
		}
	}
}

func FlipApply[A, B, R any](f fp.Func1[A, fp.Func1[B, R]], b B) fp.Func1[A, R] {
	return func(a A) R {
		return f(a)(b)
	}
}

func Compose2[A, B, GA, GR any](f fp.Func1[A, fp.Func1[B, GA]], g fp.Func1[GA, GR]) fp.Func1[A, fp.Func1[B, GR]] {

	return func(a A) fp.Func1[B, GR] {
		return func(b B) GR {
			return g(f(a)(b))
		}
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "curried_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
func Func{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) {{CurriedFunc 1 .N "R"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "R"}} {
		return Func{{dec .N}}(func({{DeclArgs 2 .N}}) R {
			return f({{CallArgs 1 .N}})
		})
	}
}
func Revert{{.N}}[{{TypeArgs 1 .N}}, R any](f {{CurriedFunc 1 .N "R"}}) func({{TypeArgs 1 .N}}) R {
	return func({{DeclArgs 1 .N}}) R {
		return f{{CurriedCallArgs 1 .N}}
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "curried_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func Flip{{dec .N}}[{{TypeArgs 1 .N}}, R any](f {{CurriedFunc 1 .N "R"}}) {{CurriedFunc 2 .N "fp.Func1[A1, R]"}} {
	return Func{{.N}}(
		func({{DeclArgs 2 .N}}, a1 A1) R {
			return f{{CurriedCallArgs 1 .N}}
		},
	)
}

func Compose{{.N}}[{{TypeArgs 1 .N}}, GA, GR any](f {{CurriedFunc 1 .N "GA"}}, g fp.Func1[GA, GR]) {{CurriedFunc 1 .N "GR"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "GR"}} {
		return Compose{{dec .N}}(f(a1), g)
	}
}

	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "curried_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func FlipApply{{dec .N}}[{{TypeArgs 1 .N}}, R any](f {{CurriedFunc 1 .N "R"}}, {{DeclArgs 2 .N}} ) fp.Func1[A1, R] {
		return func(a1 A1) R {
			return f{{CurriedCallArgs 1 .N}}
		}
	
}
	`,
}
