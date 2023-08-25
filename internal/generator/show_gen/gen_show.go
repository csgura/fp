package main

import (
	"go/types"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {

	genfp.Generate("show", "show_gen.go", func(f genfp.Writer) {

		f.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		f.GetImportedName(types.NewPackage("github.com/csgura/fp/iterator", "iterator"))

		f.Iteration(3, max.Product).Write(`
func Labelled{{.N}}[{{TypeArgs 1 .N}} fp.Named]({{DeclTypeClassArgs 1 .N "fp.Show"}}) fp.Show[fp.Labelled{{.N}}[{{TypeArgs 1 .N}}]] {
	return NewAppend(func(buf []string, t fp.Labelled{{.N}}[{{TypeArgs 1 .N}}], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			{{- range $idx := Range 1 .N}}
			ins{{$idx}}.Append(nil, t.I{{$idx}}, opt),
			{{- end}}
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}
		`, map[string]any{})

	})
}
