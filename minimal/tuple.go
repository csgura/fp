package minimal

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/internal/generator/template_gen

type Unit = fp.Unit

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "tuple_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    1,
	Until:   genfp.MaxProduct,
	Template: `
{{define "Receiver"}}func(r Tuple{{.N}}[{{TypeArgs 1 .N "T"}}]){{end}}

type Tuple{{.N}}[{{TypeArgs 1 .N "T"}} any] struct {
	{{- range $idx := Range 1 .N}}
		I{{$idx}} T{{$idx}}
	{{- end}}
}

func AsTuple{{.N}}[{{TypeArgs 1 .N}} any]({{DeclArgs 1 .N}}) {{TupleType .N}} {
	return {{TupleType .N}}{
	{{- range $idx := Range 1 .N}}
		I{{$idx}}: a{{$idx}},
	{{- end}}
	}
}
	`,
}
