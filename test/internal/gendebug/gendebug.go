package gendebug

import (
	"fmt"

	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Container[T fmt.Stringer] interface {
	Get() T
}

// @fp.Generate
func _[T fmt.Stringer]() genfp.GenerateAdaptor[Container[T]] {
	return genfp.GenerateAdaptor[Container[T]]{
		File: "gendebug_generated.go",
	}
}

type preparable[T any] interface {
	StmtFor(tx string) T
	Close() error
}

type ExecBuilder[T preparable[T]] interface {
	Build(id string) func() T
	Make(id string) T
	WithoutPublish() ExecBuilder[T]
}

// @fp.Generate
func _[T preparable[T]]() genfp.GenerateAdaptor[ExecBuilder[T]] {
	return genfp.GenerateAdaptor[ExecBuilder[T]]{
		File: "gendebug_generated.go",
		Options: []genfp.ImplOption{
			{
				Method: ExecBuilder[T].Build,
				DefaultImpl: func(self ExecBuilder[T], id string) func() T {
					return func() T {
						return self.Make(id)
					}
				},
			},
		},
	}
}
