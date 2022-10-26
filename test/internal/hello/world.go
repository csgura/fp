package hello

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type World struct {
	message   string
	timestamp time.Time
}

// @fp.Derive
var _ eq.Derives[fp.Eq[World]]
