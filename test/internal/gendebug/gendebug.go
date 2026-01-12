package gendebug

import (
	"fmt"
	"io"

	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
var _ = genfp.GenerateAdaptor[fmt.Stringer]{
	ExtendsWith: map[string]genfp.TypeTag{
		"Closer": genfp.TypeOf[io.Closer](),
	},
	Options: []genfp.ImplOption{
		{
			Method: fmt.Stringer.String,
			DefaultImpl: func(closer io.Closer) string {
				return "hello"
			},
		},
	},
}
