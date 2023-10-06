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

type Tracer interface {
	Trace(message string)
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

type TracerImpl struct {
}

// Trace implements Tracer.
func (*TracerImpl) Trace(message string) {
	panic("unimplemented")
}

var _ Tracer = &TracerImpl{}

// @fp.Generate
var _ = genfp.GenerateAdaptor[SpanContext]{
	File: "delegate_generated.go",
	Name: "SpanTrace",
	Self: true,
	ExtendsWith: map[string]genfp.TypeTag{
		"TracerImpl": genfp.TypeOf[TracerImpl](),
	},
	Delegate:           []genfp.Delegate{genfp.DelegatedBy[Tracer]("TracerImpl")},
	EmbeddingInterface: []genfp.TypeTag{genfp.TypeOf[context.Context]()},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[SpanContext]{
	File:     "delegate_generated.go",
	Name:     "SpanContextExtends",
	Extends:  true,
	Self:     true,
	Delegate: []genfp.Delegate{genfp.DelegatedBy[context.Context]("DefaultContext")},
}
