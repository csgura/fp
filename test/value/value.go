package value

import (
	"os"
	rf "reflect"
	"sync/atomic"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/ord"
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
	m2  map[string]atomic.Bool
	mm  fp.Map[string, int]
}

type NoValue struct {
}

// @fp.Derive
var _ ord.Derives[fp.Ord[Hello]]

// @fp.Derive
var _ eq.Derives[fp.Eq[Hello]]
