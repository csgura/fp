package value

import (
	"fmt"
	"os"
	rf "reflect"
	"sync/atomic"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/hello"
	"github.com/csgura/fp/test/internal/js"
)

//go:generate go run github.com/csgura/fp/cmd/gombok
type (
	// Hello is hello
	// @fp.Value
	// @fp.JsonTag
	Hello struct { // Hello
		world string
		hi    int `bson:"hi"`
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

// @fp.Value
type Entry[A interface{ String() string }, B any] struct {
	name  string
	value A
	tuple fp.Tuple2[A, B]
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Entry[interface{ String() string }, any]]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[Entry[interface{ String() string }, any]]]

// @fp.Value
type Key struct {
	a int
	b float32
	c []byte
}

func (r Key) Hash() uint32 {
	return HashableKey.Hash(r)
}

// @fp.Derive
var _ hash.Derives[fp.Hashable[Key]]

// @fp.Value
type Point struct {
	x int
	y int
	z fp.Tuple2[int, int]
}

func (r Point) String() string {
	return fmt.Sprintf("(%d,%d)", r.x, r.y)
}

var MonoidInt = monoid.Sum[int]()

// @fp.Derive
var _ monoid.Derives[fp.Monoid[Point]]

// @fp.Value
// @fp.Json
type Greeting struct {
	hello    hello.World
	language string
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Greeting]]

// @fp.Derive
var _ js.Derives[js.Encoder[Greeting]]
