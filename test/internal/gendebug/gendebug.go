package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/seq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Hello interface {
	World(address string, count int) fp.Try[string]
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
var _ = genfp.GenerateFromInterfaces{
	File: "intf_generated.go",
	List: seq.Of(genfp.TypeOf[Hello]()),
	Template: `
		type Message struct {
		}
	`,
}
