package fn1

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

// first :: a b c -> a (b, d) (c, d)
func First[D, B, C any](f func(B) C) func(B, D) (C, D) {
	return func(b B, d D) (C, D) {
		return f(b), d
	}
}

// second :: a b c -> a (d, b) (d, c)
func Second[D, B, C any](f func(B) C) func(D, B) (D, C) {
	return func(d D, b B) (D, C) {
		return d, f(b)
	}
}

// (***) :: a b c -> a b' c' -> a (b, b') (c, c')
// haskell 에서는 기호외에 따로 이름은 없는데
// scala cats 에서 split 이라는 이름으로 정의함
func Split[A, B, C, D any](f1 func(A) B, f2 func(C) D) func(A, C) (B, D) {
	return func(a A, c C) (B, D) {
		return f1(a), f2(c)
	}
}

// (&&&) :: a b c -> a b c' -> a b (c, c')
// haskell 에서는 기호외에 따로 이름은 없는데
// scala cats 에서 merge 라는 이름으로 정의함.
func Merge[A, B, C any](f1 func(A) B, f2 func(A) C) func(A) (B, C) {
	return func(a A) (B, C) {
		return f1(a), f2(a)
	}
}

func Merge2[A, B, C any](f1 func(A) B, f2 func(A) C) func(A) fp.Tuple2[B, C] {
	return func(a A) fp.Tuple2[B, C] {
		return fp.Tuple2[B, C]{
			I1: f1(a),
			I2: f2(a),
		}
	}
}

//go:generate go run github.com/csgura/fp/internal/generator/template_gen

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "arrow_func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
	{{define "fargs"}}
		{{- range $idx := Range 1 .N -}}
			f{{$idx}} func(A) {{TypeArg $idx}},
		{{- end -}}
	{{end}}

	func Merge{{.N}}[A, {{TypeArgs 1 .N}} any]({{template "fargs" .}}) func(A) fp.Tuple{{.N}}[{{TypeArgs 1 .N}}] {
		return func(a A) fp.Tuple{{.N}}[{{TypeArgs 1 .N}}] {
			return fp.Tuple{{.N}}[{{TypeArgs 1 .N}}] {
				{{- range $idx := Range 1 .N}}
					I{{$idx}}: f{{$idx}}(a),
				{{- end}}
			}
		}
	}
	`,
}
