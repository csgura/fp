package gendebug

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/test/internal/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.String(useShow=true)
type ShowConstraintExplicit[T fmt.Stringer] struct {
	hello   string
	world   int
	message T
}

func ShowShowConstraintExplicit[T fmt.Stringer]() fp.Show[ShowConstraintExplicit[T]] {
	return show.New(func(t ShowConstraintExplicit[T]) string {
		return fmt.Sprintf("ShowConstraintExplicit(message=%s)", t.message)
	})
}
