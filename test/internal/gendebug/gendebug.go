package gendebug

import (
	"io"

	"github.com/csgura/fp"
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	world []Container
}

func (r *Hello) Close() error {
	return nil
}

type Container struct {
	message string
}

func (r *Container) Close() error {
	return nil
}

func Struct1Mine[A io.Closer](ashow mshow.Show[A]) mshow.Show[fp.Tuple1[[]A]] {
	panic("")
}

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[Hello]]
