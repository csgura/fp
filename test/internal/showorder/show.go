package showorder

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/show"
)

type World struct {
	Loc string
}

type Hello struct {
	A     int
	B     string
	AN    any
	World World
}

func ShowAny[T any]() fp.Show[T] {
	return show.New(func(t T) string {
		return fmt.Sprint(t)
	})
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive(recursive=true)
var _ show.Derives[fp.Show[Hello]]
