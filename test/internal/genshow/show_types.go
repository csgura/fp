package genshow

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/show"
	"github.com/csgura/fp/test/internal/recursive"
)

type Person struct {
	Name string
	Age  int
}

type NoDerive struct {
	Hello string
}

type Collection struct {
	Index       map[string]Person
	List        []Person
	Description *string
	Set         fp.Set[int]
	Option      fp.Option[Person]
	NoDerive    NoDerive
	Stringer    HasStringMethod
	BoolPtr     *bool
	NoMap       map[string]NoDerive
	Alias       recursive.StringAlias
	StringSeq   fp.Seq[string]
}

type HasStringMethod struct {
	There string
}

func (r HasStringMethod) String() string {
	return r.There
}

type DupGenerate struct {
	NoDerive NoDerive
	World    string
}

type HasTuple struct {
	Entry fp.Tuple2[string, int]
	HList hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
}

// @fp.Value
type EmbeddedStruct struct {
	hello string
	world struct {
		Level int
		Stage string
	}
}

// @fp.Value
type EmbeddedTypeParamStruct[T any] struct {
	hello string
	world struct {
		Level T
		Stage string
	}
}

func UntypedStructFunc(s struct {
	Level   int
	Stage   string
	privacy string
}) {
	fmt.Println("Ok")
}

type EmptyStruct struct {
}

type TLV = []byte

type HasAliasType struct {
	Data TLV
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.ImportGiven
var _ show.Derives[fp.Show[any]]

// @fp.Summon
var showMap fp.Show[map[string]string]

// @fp.Generate
var _ = genfp.GenerateFromStructs{
	File: "show_gen.go",
	Imports: genfp.Imports(
		"github.com/csgura/fp",
		"github.com/csgura/fp/as",
		"github.com/csgura/fp/show",
	//	"github.com/csgura/fp/lazy",
	),
	List: seq.Of(
		genfp.TypeOf[Person](),
	//	genfp.TypeOf[Collection](),
	),
	Recursive: true,
	Template: `
		func Show{{.N}}(v {{.N.Type}}, opt fp.ShowOption) string {
			return show.FormatStruct("{{.N.Package}}.{{.N}}", opt, 
			{{- range .N.Fields}}
				{{- $name := .Name -}}
				{{- with Summon "fp.Show" .Type}}
					as.Tuple2("{{$name}}", show.AsAppender({{.}}, v.{{$name}})),
				{{- else }}
					{{- if .Type.IsNumber}}
						as.Tuple2("{{.Name}}", show.AsAppender(show.Number[{{.Type}}](), v.{{.Name}})),
					{{- else}}
						as.Tuple2("{{.Name}}", show.AsAppender(nil, v.{{.Name}})),
					{{- end}}	
				{{- end}}
			{{- end}}
			)
		}
	`,
}

var _ mshow.Derives[mshow.Show[HasAliasType]]

var _ mshow.Derives[mshow.Show[Person]]

var _ mshow.Derives[mshow.Show[Collection]]

var _ mshow.Derives[mshow.Show[DupGenerate]]

var _ mshow.Derives[mshow.Show[HasTuple]]

var _ mshow.Derives[mshow.Show[EmbeddedStruct]]

var _ mshow.Derives[mshow.Show[EmbeddedTypeParamStruct[any]]]

var _ mshow.Derives[mshow.Show[EmptyStruct]]

// @fp.Generate
var _ = genfp.GenerateFromStructs{
	File:    "firstfield.go",
	Imports: genfp.Imports(),
	List: seq.Of(
		genfp.TypeOf[EmptyStruct](),
		genfp.TypeOf[Person](),

	//	genfp.TypeOf[Collection](),
	),
	Recursive: true,
	Template: `
		func FirstField{{.N}}() string {
			{{- if .N.FieldAt 0}}
				return "{{(.N.FieldAt 0).Name}}"
			{{- else}}
				return ""
			{{- end}}
		}
	`,
}
