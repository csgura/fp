package adaptortest

import (
	"context"
	"io"

	"github.com/csgura/fp/genfp"
)

type SpanContext interface {
	context.Context
	Hello() string
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[SpanContext]{
	File:     "delegate_generated.go",
	Self:     true,
	Delegate: []genfp.Delegate{genfp.DelegatedBy[context.Context]("DefaultContext")},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[SpanContext]{
	File:               "delegate_generated.go",
	Name:               "SpanContextEmbedding",
	Self:               true,
	EmbeddingInterface: []genfp.TypeTag{genfp.TypeOf[context.Context](), genfp.TypeOf[io.Closer]()},
}
