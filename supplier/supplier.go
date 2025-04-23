package supplier

import (
	"sync"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

func Pure[A any](v A) fp.Supplier[A] {
	return func() A {
		return v
	}
}

func Map[A, B any](m fp.Supplier[A], fn fp.Func1[A, B]) fp.Supplier[B] {
	return func() B {
		return fn(m())
	}
}

func FlatMap[A, B any](m fp.Supplier[A], fn fp.Func1[A, fp.Supplier[B]]) fp.Supplier[B] {
	return func() B {
		return fn(m())()
	}
}

func Flatten[A any](m fp.Supplier[fp.Supplier[A]]) fp.Supplier[A] {
	return func() A {
		return m()()
	}
}

func Memoize[A any](f fp.Supplier[A]) fp.Supplier[A] {
	once := sync.Once{}
	var ret A
	return func() A {
		once.Do(func() {
			ret = f()
		})
		return ret
	}
}

func Get[A any](f fp.Supplier[A]) A {
	return f()
}

func IntoFunc0[A any](f fp.Supplier[A]) fp.Func0[A] {
	return func(a1 fp.Unit) A {
		return f()
	}
}

func FromFunc0[A any](f fp.Func0[A]) fp.Supplier[A] {
	return func() A {
		return f(fp.Unit{})
	}
}

//go:generate go run github.com/csgura/fp/internal/generator/template_gen

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "supplier_gen.go",

	From:  1,
	Until: genfp.MaxFunc,
	Template: `

func FromFunc{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R, {{DeclArgs 1 .N}}) func() R {
	return func() R {
		return f({{CallArgs 1 .N}})
	}
}	
`,
}
