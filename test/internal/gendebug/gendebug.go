package gendebug

import "github.com/csgura/fp/test/internal/testpk1"

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type AliasTest struct {
	ctx testpk1.Pk1Context
}
