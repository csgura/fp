package genintf

import (
	"context"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/seq"
)

type ResponseType[V any] struct {
}

func (r ResponseType[V]) SendRequest(ref string, timeout time.Duration) fp.Try[V] {
	panic("")
}

type Hello interface {
	World(address string, count int) fp.Try[string]
	Universe(ctx context.Context, galaxy string) fp.Try[string]
}

type handler struct {
	ref     string
	timeout time.Duration
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
var _ = genfp.GenerateFromInterfaces{
	File: "intf_generated.go",
	Imports: genfp.Imports(
		"context",
		"github.com/csgura/fp",
	),

	List: seq.Of(genfp.TypeOf[Hello]()),
	Variables: map[string]string{
		"receiver": "r handler",
		"actorRef": "r.ref",
		"timeout":  "r.timeout",
	},
	Template: `
		{{$receiver := .receiver}}
		{{$timeout := .timeout}}
		{{$actorRef := .actorRef}}

		{{range .N.Methods}}
			// @fp.Getter
			// @fp.AllArgsConstructor
			type Message{{.Name}} struct {
				{{- range .Args}}
					{{.Name}} {{.Type}}
				{{- end}}
				ResponseType[{{(.ReturnAt 0).Type.TypeArgAt 0}}]
			}

			func ({{$receiver}}) {{.Name}}( {{.ArgsDef}} ) {{.ReturnsDef}} {
				return NewMessage{{.Name}}({{.ArgsCall}}).SendRequest({{$actorRef}},{{$timeout}})
			}
		{{end}}
		
		
	`,
}
