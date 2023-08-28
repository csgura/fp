package testpk2

import (
	"fmt"
	"io"
	"os"
	rf "reflect"
	"sync/atomic"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
	"github.com/csgura/fp/test/internal/testpk1"
)

//go:generate go run github.com/csgura/fp/cmd/gombok
type (
	// Hello is hello
	// @fp.Value
	// @fp.JsonTag
	Hello struct { // Hello
		world string
		hi    int `bson:"hi" json:"merong"`
	}
)

type Embed struct {
}

type Local interface {
	Local()
}

// @fp.Value
type AllKindTypes struct { // what the
	Embed
	hi fp.Option[int]

	tpe  rf.Type
	arr  []os.File
	m    map[string]int
	a    any
	p    *int
	l    Local
	t    fp.Try[fp.Option[Local]]
	m2   map[string]atomic.Bool
	mm   fp.Map[string, int]
	intf fp.Future[int]
	ch   chan fp.Try[fp.Either[int, string]]
	ch2  chan<- int
	ch3  <-chan int

	fn3  fp.Func1[int, fp.Try[string]]
	fn   func(a string) fp.Try[int]
	fn2  func(fp.Try[string]) (result int, err error)
	arr2 [2]int
	st   struct {
		Embed
		A int
		B fp.Option[string]
	}
	i2 interface {
		io.Closer
		Hello() fp.Try[int]
	}
}

type NoValue struct {
}

// @fp.Value
type Person struct {
	name       string
	age        int
	height     float64
	phone      fp.Option[string]
	addr       []string
	list       hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
	seq        fp.Seq[float64]
	blob       []byte
	_notExport string
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
type Entry[A comparable, B any, C fmt.Stringer, D interface{ Hello() string }] struct {
	name  string
	value A
	tuple fp.Tuple2[A, B]
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Entry[string, any, fmt.Stringer, interface{ Hello() string }]]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[Entry[string, any, fmt.Stringer, interface{ Hello() string }]]]

// @fp.Value
type Key struct {
	a int
	b float32
	c []byte
}

func (r Key) Hash() uint32 {
	return HashableKey().Hash(r)
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
// @fp.GenLabelled
type Greeting struct {
	hello    testpk1.World
	language string
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Greeting]]

// @fp.Derive
var _ js.Derives[js.Encoder[Greeting]]

// @fp.Derive
var _ js.Derives[js.Decoder[Greeting]]

// @fp.Value
// @fp.GenLabelled
type Three struct {
	one   int
	two   string
	three float64
}

// @fp.Derive
var _ js.Derives[js.Encoder[Three]]

// @fp.Derive
var _ js.Derives[js.Decoder[Three]]

// @fp.Derive
var _ show.Derives[fp.Show[Three]]

// @fp.Derive
var _ read.Derives[read.Read[Three]]

// @fp.Derive
var _ eq.Derives[fp.Eq[testpk1.World]]

// @fp.Derive
var _ eq.Derives[fp.Eq[testpk1.Wrapper[any]]]

// @fp.Value
type Tree struct {
	root testpk1.Node
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Tree]]

// @fp.GetterPubField
// @fp.Deref
// @fp.WithPubField
type AliasedStruct testpk1.DefinedOtherPackage

func (r AliasedStruct) String() string {
	return "AliasedStruct"
}

// @fp.GetterPubField(override=true)
// @fp.Deref
type GetterOverride testpk1.DefinedOtherPackage

// @fp.Derive
var _ js.Derives[js.Encoder[testpk1.World]]

// @fp.Generate
var GenShow = genfp.GenerateDirective{
	File: "show_gen.go",
	Imports: []genfp.ImportName{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{"github.com/csgura/fp/seq", "seq"},
	},
	From:     3,
	Until:    max.Product,
	Template: "hello world",
}
