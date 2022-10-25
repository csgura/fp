package value

import (
	"os"
	rf "reflect"
	"sync/atomic"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
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

// @fp.Value
type Person struct {
	name   string
	age    int
	height float64
	phone  fp.Option[string]
	addr   []string
	list   hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
	seq    fp.Seq[float64]
	blob   []byte
}

func EqFpSeq[T any](e fp.Eq[T]) fp.Eq[fp.Seq[T]] {
	return eq.Seq(e)
}

var EqFloat64 = eq.Given[float64]()

// @fp.Derive
var _ eq.Derives[fp.Eq[Person]]

// @fp.Value
type Wallet struct {
	owner  Person
	amount int64
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Wallet]]
