package hello

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/test/internal/js"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.Json
type World struct {
	message   string
	timestamp time.Time
}

// @fp.Derive
var _ eq.Derives[fp.Eq[World]]

// @fp.Derive
var _ js.Derives[js.Encoder[World]]
