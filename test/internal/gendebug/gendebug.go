package gendebug

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/test/internal/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello struct {
	AN any
}

func ShowAny[T any]() fp.Show[T] {
	return show.New(func(t T) string {
		return fmt.Sprint(t)
	})
}

// @fp.Derive(recursive=true)
var _ show.Derives[fp.Show[Hello]]
