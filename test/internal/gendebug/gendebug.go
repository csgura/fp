package gendebug

import (
	"time"

	"github.com/csgura/fp/test/internal/namedptr"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Value struct {
	Present int
	Value   *time.Time
}

// @fp.Derive
var _ namedptr.Derives[namedptr.Validator[Value]]
