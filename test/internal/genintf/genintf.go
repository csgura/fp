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

type Alias = Todo

type Todo interface {
	Today() fp.Try[string]
}

type handler struct {
	ref     string
	timeout time.Duration
}

//go:generate go run github.com/csgura/fp/cmd/gombok

const generateTemplate = `
	{{$receiver := .receiver}}
	{{$timeout := .timeout}}
	{{$actorRef := .actorRef}}

	{{range .N.Methods}}
		// @fp.Getter
		// @fp.AllArgsConstructor
		type Message{{.Name}} struct {
			{{- range .Args}}
				{{.Name}} {{.Type | TypeDecl}}
			{{- end}}
			ResponseType[{{(.ReturnAt 0).Type.TypeArgAt 0 | TypeDecl}}]
		}

		func ({{$receiver}}) {{.Name}}( {{.Args | VarDecl}} ) {{.Returns | TypeDecl}} {
			return NewMessage{{.Name}}({{.ArgsCall}}).SendRequest({{$actorRef}},{{$timeout}})
		}
	{{end}}
		
		
`

type World struct {
}

func (r World) Size() int {
	return 0
}

// @fp.Generate
var _ = genfp.GenerateFromInterfaces{
	File: "intf_generated.go",
	List: seq.Of(
		genfp.TypeOf[Hello](),
		genfp.TypeOf[Alias](),
	),
	Parameters: map[string]string{
		"receiver": "r handler",
		"actorRef": "r.ref",
		"timeout":  "r.timeout",
	},
	Template: generateTemplate,
}

// @fp.Generate
var _ = genfp.GenerateFromStructs{
	File: "intf_generated.go",
	List: seq.Of(
		genfp.TypeOf[World](),
	),
	Template: `
		{{$st := .N}}
		// {{.N}} HasMethod Size :{{.N.HasMethod "Size"}} 
		// {{.N}} HasMethod Head :{{.N.HasMethod "Head"}} 

		{{range .N.Methods}}
			// {{$st}} has method {{.Name}}
		{{end}}
	`,
}
