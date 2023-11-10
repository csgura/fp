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
