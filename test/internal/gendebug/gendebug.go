package gendebug

import (
	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Tracer interface {
	Trace(message string)
}

type Closer interface {
	Close() error
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Tracer]{
	File: "gendebug_generated.go",
	Name: "TracerWith",
	Self: true,
	ExtendsWith: map[string]genfp.TypeTag{
		"Closer": genfp.TypeOf[Closer](),
	},
	Options: []genfp.ImplOption{
		{
			Method: Closer.Close,
			Delegate: genfp.Delegate{
				Field: "Closer",
			},
			DefaultImpl: func() error {
				return nil
			},
		},
	},
}
