package clone

import (
	"github.com/csgura/fp"

	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Derives[T any] interface {
}

func New[T any](f func(T) T) fp.Clone[T] {
	return fp.CloneFunc[T](f)
}

func Ptr[T any](tshow lazy.Eval[fp.Clone[T]]) fp.Clone[*T] {
	return New(func(pt *T) *T {
		if pt == nil {
			return nil
		}
		var t = *pt
		return &t
	})
}

func Given[T any]() fp.Clone[T] {
	return New(func(t T) T {
		return t
	})
}

var HNil = New(func(hlist.Nil) hlist.Nil {
	return hlist.Empty()
})

func Seq[T any](tclone fp.Clone[T]) fp.Clone[fp.Seq[T]] {
	return New(func(s fp.Seq[T]) fp.Seq[T] {
		return seq.Map(s, tclone.Clone)
	})
}

func GoMap[K comparable, V any](clonek fp.Clone[K], clonev fp.Clone[V]) fp.Clone[map[K]V] {
	return New(func(s map[K]V) map[K]V {
		ret := map[K]V{}
		for k, v := range s {
			ret[clonek.Clone(k)] = clonev.Clone(v)
		}
		return ret
	})
}

func Slice[T any](tclone fp.Clone[T]) fp.Clone[[]T] {
	return New(func(s []T) []T {
		return seq.Map(s, tclone.Clone)
	})
}

func Option[T any](tclone fp.Clone[T]) fp.Clone[fp.Option[T]] {
	return New(func(s fp.Option[T]) fp.Option[T] {
		return option.Map(s, tclone.Clone)
	})
}

func HCons[H any, T hlist.HList](hclone fp.Clone[H], tclone fp.Clone[T]) fp.Clone[hlist.Cons[H, T]] {
	return New(func(list hlist.Cons[H, T]) hlist.Cons[H, T] {

		h := hclone.Clone(list.Head())
		t := tclone.Clone(list.Tail())

		return hlist.Concat(h, t)

	})
}

func Tuple2[A1, A2 any](ins1 fp.Clone[A1], ins2 fp.Clone[A2]) fp.Clone[fp.Tuple2[A1, A2]] {
	return New(func(t fp.Tuple2[A1, A2]) fp.Tuple2[A1, A2] {
		return as.Tuple2(
			ins1.Clone(t.I1),
			ins2.Clone(t.I2),
		)
	})
}

// @fp.Generate
var GenClone = genfp.GenerateFromUntil{
	File: "clone_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  3,
	Until: genfp.MaxProduct,
	Template: `
func Tuple{{.N}}[{{TypeArgs 1 .N}} any]({{DeclTypeClassArgs 1 .N "fp.Clone"}}) fp.Clone[fp.{{TupleType .N}}] {
	return New(func(t fp.{{TupleType .N}}) fp.{{TupleType .N}} {
		return as.Tuple{{.N}}(
			{{- range $idx := Range 1 .N}}
			ins{{$idx}}.Clone(t.I{{$idx}}),
			{{- end}}
		)
	})
}
	`,
}

func Generic[A, Repr any](gen fp.Generic[A, Repr], reprClone fp.Clone[Repr]) fp.Clone[A] {
	return New(func(a A) A {
		return gen.From(reprClone.Clone(gen.To(a)))
	})
}
