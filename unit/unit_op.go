//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package unit

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/genfp"
)

func Func0(f func()) fp.Func1[fp.Unit, fp.Unit] {
	return func(fp.Unit) fp.Unit {
		f()
		return fp.Unit{}
	}
}

var Success = fp.Success(fp.Unit{})
var Some = fp.Some(fp.Unit{})
var None = fp.None[fp.Unit]()
var Completed = future.Successful(fp.Unit{})

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  1,
	Until: genfp.MaxFunc,
	Template: `
func Func{{.N}}[{{TypeArgs 1 .N}} any]( f func({{TypeArgs 1 .N}}) ) fp.Func{{.N}}[{{TypeArgs 1 .N}},fp.Unit] {
	return func({{DeclArgs 1 .N}}) fp.Unit {
		f({{CallArgs 1 .N}})
		return fp.Unit{}
	}
}
	`,
}
