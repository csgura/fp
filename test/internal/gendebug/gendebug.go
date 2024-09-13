package gendebug

import (
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/test/internal/testpk1"
	"github.com/csgura/fp/test/internal/testpk2"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type AliasTest struct {
	ctx    testpk1.Pk1Context
	other  Pk1Context
	pk2ctx testpk2.Pk1Context
}

type Pk1Context = testpk1.Pk1Context

type AliasIntf interface {
	Ctx(arg any) testpk1.Pk1Context
	Other() Pk1Context
	Pk2Ctx() testpk2.Pk1Context
	Some(ctxs ...testpk2.Pk1Context) (string, error)
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[AliasIntf]{
	File: "adaptor.go",
	Self: true,
}
