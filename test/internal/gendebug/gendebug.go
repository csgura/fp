package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.ImportGiven
var _ show.Derives[fp.Show[any]]

// @fp.Summon
var showMap fp.Show[map[string]string]
