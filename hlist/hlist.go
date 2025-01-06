//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package hlist

import (
	"github.com/csgura/fp/genfp"
)

// HList is contraints interface type to force some argument type to be one of Cons[_,_] | Nil
// but go does not support existential or wildcard type
type HList interface {
	// Cons[_,_] | Nil
}

// Header is constrains interface type,  enforce Head type of Cons is HT
type Header[HT any] interface {
	HList
	Head() HT
}

type Nil struct {
}

func (r Nil) Head() Nil {
	return r
}

type Cons[H any, T HList] struct {
	head H
	tail T
}

func IsNil[T HList](v T) bool {
	switch any(v).(type) {
	case Nil:
		return true
	}
	return false
}

func Head[H any, T HList](r Cons[H, T]) H {
	return r.head
}

func Tail[H any, T HList](r Cons[H, T]) T {
	return r.tail
}

func (r Cons[H, T]) Head() H {
	return r.head
}

func Concat[H any, T HList](h H, t T) Cons[H, T] {
	return Cons[H, T]{h, t}
}

func Of1[H any](h H) Cons[H, Nil] {
	return Concat(h, Nil(Nil{}))
}

func Empty() Nil {
	return Nil{}
}

func Lift1[A, R any](f func(a A) R) func(v Cons[A, Nil]) R {
	return func(v Cons[A, Nil]) R {
		return f(v.Head())
	}
}

func Rift1[A, R any](f func(a A) R) func(v Cons[A, Nil]) R {
	return func(v Cons[A, Nil]) R {
		return f(v.Head())
	}
}

func Case1[A1 any, T HList, R any](hl Cons[A1, T], f func(a1 A1) R) R {
	return f(hl.Head())
}

// func Reverse1[A1 any](hl Cons[A1, Nil]) Cons[A1, Nil] {
// 	return hl
// }

// func Reverse2[A1, A2 any](hl Cons[A1, Cons[A2, Nil]]) Cons[A2, Cons[A1, Nil]] {
// 	panic("")
// 	//return Concat(Reverse1(hl.Tail()), hl.Head())
// }

// func Reverse3[A1, A2, A3 any](hl Cons[A1, Cons[A2, Cons[A3, Nil]]]) Cons[A3, Cons[A2, Cons[A1, Nil]]] {
// 	//panic("")

// }

func Unapply[H any, T HList](list Cons[H, T]) (H, T) {
	return list.Head(), Tail(list)
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "lift_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    2,
	Until:   genfp.MaxFunc,
	Template: `
func Lift{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) func({{RecursiveType "Cons" 1 .N "Nil"}}) R {
	return func(v {{RecursiveType "Cons" 1 .N "Nil"}}) R {
		rf := Lift{{dec .N}}(func({{DeclArgs 2 .N}}) R {
			return f(v.Head(),{{CallArgs 2 .N}})
		})

		return rf(Tail(v))
	}
}
func Rift{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) func({{RecursiveType "Cons" .N 1 "Nil"}}) R {
	return func(v {{RecursiveType "Cons" .N 1 "Nil"}}) R {
		rf := Rift{{dec .N}}(func({{DeclArgs 1 (dec .N)}}) R {
			return f({{CallArgs 1 (dec .N)}}, v.Head())
		})

		return rf(Tail(v))
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "case_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    2,
	Until:   genfp.MaxProduct,
	Template: `
func Case{{.N}}[{{TypeArgs 1 .N}} any, T HList, R any](hl {{RecursiveType "Cons" 1 .N "T"}}, f func({{TypeArgs 1 .N}}) R) R {
	return Case{{dec .N}}(Tail(hl), func({{DeclArgs 2 .N}}) R {
		return f(hl.Head(), {{CallArgs 2 .N}})
	})
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "of_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    2,
	Until:   genfp.MaxProduct,
	Template: `
func Of{{.N}}[{{TypeArgs 1 .N}} any]({{DeclArgs 1 .N}}) {{RecursiveType "Cons" 1 .N "Nil"}} {
	return Concat(a1, Of{{dec .N}}({{CallArgs 2 .N}}))
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "reverse_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    2,
	Until:   genfp.MaxFunc,
	Template: `
func Reverse{{.N}}[{{TypeArgs 1 .N}} any](hl {{RecursiveType "Cons" 1 .N "Nil"}}) {{RecursiveType "Cons" .N 1 "Nil"}} {
	return Case{{.N}}(hl, func({{DeclArgs 1 .N}}) {{RecursiveType "Cons" .N 1 "Nil"}} {
		return Of{{.N}}({{ReverseCallArgs 1 .N}})
	})
}
	`,
}
