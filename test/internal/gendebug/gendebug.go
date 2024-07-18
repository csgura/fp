package gendebug

import (
	"context"

	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type SimpleIntf interface {
	Hello(msg string) string
}
type ComplexIntf interface {
	Hello(ctx context.Context, msg string) string
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[ComplexIntf]{
	File: "gendebug_generated.go",
	Self: true,
	ExtendsWith: map[string]genfp.TypeTag{
		"Extends": genfp.TypeOf[SimpleIntf](),
	},
	Options: []genfp.ImplOption{
		{
			Method:   ComplexIntf.Hello,
			Delegate: genfp.DelegatedBy[SimpleIntf]("Extends"),
		},
	},
}
