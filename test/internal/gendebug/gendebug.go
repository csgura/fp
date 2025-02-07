package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/testpk1"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type LegacyPerson struct {
	Name    string
	Age     int
	privacy string
}

type LegacyPhoneBook struct {
	Person LegacyPerson
	Phone  string
}

// @fp.Derive(recursive=true)
var _ monoid.Derives[fp.Monoid[testpk1.LegacyPhoneBook]]
