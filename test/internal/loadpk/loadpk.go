package loadpk

import "github.com/csgura/fp/genfp"

const Hello = 10
const File = "gen.go"
const start = 1

// @ConstTest
var _ = genfp.GenerateFromUntil{
	File: File,
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/tstate", Name: "tstate"},
		{Package: "context", Name: "context"},
	},
	From:     start,
	Until:    genfp.MaxFunc,
	Template: ``,
}
