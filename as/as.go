//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package as

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

func PartialFunc[T, R any](isDefinedAt func(T) bool, apply func(T) R) fp.PartialFunc[T, R] {
	return fp.PartialFunc[T, R]{IsDefinedAt: isDefinedAt, Apply: apply}
}

func Func0[R any](f func() R) fp.Func1[fp.Unit, R] {
	return func(u fp.Unit) R {
		return f()
	}
}

func Seq[T any](s []T) fp.Seq[T] {
	return fp.Seq[T](s)
}

func SeqNonNil[T any](s []*T) fp.Seq[T] {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		if v != nil {
			ret = append(ret, *v)
		}
	}
	return ret
}

func Ptr[T any](v T) *T {
	return &v
}

func Interface[T, I any](v T) I {
	var a any = v
	return a.(I)
}

func Any[T any](v T) any {
	return v
}

func InstanceOf[T any](v any) T {
	return v.(T)
}

func Tuple[K, V any](k K, v V) fp.Tuple2[K, V] {
	return fp.Tuple2[K, V]{I1: k, I2: v}
}

func Named[V any](name string, v V) fp.RuntimeNamed[V] {
	return fp.RuntimeNamed[V]{I1: name, I2: v}
}

func Dual[T any](t T) fp.Dual[T] {
	return fp.Dual[T]{GetDual: t}
}

func Endo[T any](f func(T) T) fp.Endo[T] {
	return fp.Endo[T](f)
}

func MapEntry[K, V any](xtrKey func(V) K) func(V) fp.Tuple2[K, V] {
	return func(v V) fp.Tuple2[K, V] {
		return Tuple2(xtrKey(v), v)
	}
}

func Left[E fp.Either[L, R], L, R any](l L) fp.Either[L, R] {
	ret := fp.Left[L, R](l)
	return ret
}

func Right[E fp.Either[L, R], L, R any](r R) fp.Either[L, R] {
	ret := fp.Right[L, R](r)
	return ret
}

func Generic[T, Repr any](tpe string, kind string, to func(T) Repr, from func(Repr) T) fp.Generic[T, Repr] {
	return fp.Generic[T, Repr]{
		Type: tpe,
		Kind: kind,
		To:   to,
		From: from,
	}
}

func Ord[T any](less fp.LessFunc[T]) fp.Ord[T] {
	return less
}

func Supplier[T any](v T) fp.Supplier[T] {
	return func() T {
		return v
	}
}

func Predicate[T any](f func(T) bool) fp.Predicate[T] {
	return f
}

func Tupled2[A1, A2, R any](fn fp.Func2[A1, A2, R]) func(fp.Tuple2[A1, A2]) R {
	return func(t fp.Tuple2[A1, A2]) R {
		return fn(t.Unapply())
	}
}

func State[S, A any](f func(S) fp.Tuple2[A, S]) fp.State[S, A] {
	return f
}

func StateT[S, A any](f func(S) fp.Try[fp.Tuple2[A, S]]) fp.StateT[S, A] {
	return f
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  1,
	Until: genfp.MaxFunc,
	Template: `

func Func{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) fp.Func{{.N}}[{{TypeArgs 1 .N}}, R] {
	return fp.Func{{.N}}[{{TypeArgs 1 .N}}, R](f)
}

func Supplier{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R, {{DeclArgs 1 .N}}) func() R {
	return func() R {
		return f({{CallArgs 1 .N}})
	}
}	
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  2,
	Until: 3,
	Template: `
func Curried2[A1, A2, R any](f func(A1,A2) R) fp.Func1[A1, fp.Func1[A2, R]] {
	return func(a1 A1) fp.Func1[A2, R] {
		return func(a2 A2) R {
			return f(a1, a2)
		}
	}
}	
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func Curried{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) {{CurriedFunc 1 .N "R"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "R"}} {
		return Curried{{dec .N}}(func({{DeclArgs 2 .N}}) R {
			return f({{CallArgs 1 .N}})
		})
	}
}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
func UnTupled{{.N}}[{{TypeArgs 1 .N}}, R any](f func(fp.{{TupleType .N}}) R) func({{TypeArgs 1 .N}}) R {
	return func({{DeclArgs 1 .N}}) R {
		return f(Tuple{{.N}}({{CallArgs 1 .N}}))
	}
}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  1,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]({{DeclArgs 1 .N}}) fp.{{TupleType .N}} {
	return fp.{{TupleType .N}}{
	{{- range $idx := Range 1 .N}}
		I{{$idx}}: a{{$idx}},
	{{- end}}
	}
}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  1,
	Until: 2,
	Template: `
		func HList1[A1 any](tuple fp.Tuple1[A1]) hlist.Cons[A1, hlist.Nil] {
			return hlist.Concat(tuple.Head(), hlist.Empty())
		}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func HList{{.N}}[{{TypeArgs 1 .N}} any](tuple fp.{{TupleType .N}}) {{ConsType 1 .N "hlist.Nil"}} {
	return hlist.Concat(tuple.Head(), hlist.Of{{dec .N}}(tuple.Tail()))
}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "labelled_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  1,
	Until: genfp.MaxProduct,
	Template: `
func Labelled{{.N}}[{{TypeArgs 1 .N}} fp.Named]({{DeclArgs 1 .N}}) fp.Labelled{{.N}}[{{TypeArgs 1 .N}}] {
	return fp.Labelled{{.N}}[{{TypeArgs 1 .N}}]{
	{{- range $idx := Range 1 .N}}
		I{{$idx}}: a{{$idx}},
	{{- end}}
	}
}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "labelled_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  1,
	Until: 2,
	Template: `
		func HList1Labelled[A1 fp.Named](tuple fp.Labelled1[A1]) hlist.Cons[A1, hlist.Nil] {
			return hlist.Concat(tuple.Head(), hlist.Empty())
		}
`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "labelled_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
func HList{{.N}}Labelled[{{TypeArgs 1 .N}} fp.Named](tuple fp.Labelled{{.N}}[{{TypeArgs 1 .N}}]) {{ConsType 1 .N "hlist.Nil"}} {
	return hlist.Of{{.N}}(tuple.Unapply())
}
`,
}
