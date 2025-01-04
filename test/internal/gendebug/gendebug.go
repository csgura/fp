package gendebug

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.Json
// @fp.GenLabelled
type World struct {
	message    string `hello:"message"`
	timestamp  time.Time
	Pub        string
	_notExport string
}

// @fp.Derive
var _ eq.Derives[fp.Eq[World]]

// @fp.Summon
var eqPtrWorld fp.Eq[*World]
