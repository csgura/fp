package value

import (
	"os"
	rf "reflect"

	"github.com/csgura/fp"
)

//go:generate go run github.com/csgura/fp/cmd/gombok
type (
	// Hello is hello
	// @fp.Value
	Hello struct { // Hello
		world string
		hi    int
	}
)

type Embed struct {
}

type Local interface {
	Local()
}

// @fp.Value
type MyMy struct { // what the
	Embed
	hi fp.Option[int]

	tpe rf.Type
	arr []os.File
	m   map[string]int
	a   any
	p   *int
	l   Local
	t   fp.Try[fp.Option[Local]]
}

type NoValue struct {
}
