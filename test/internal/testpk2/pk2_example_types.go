package testpk2

import (
	"fmt"
	"io"
	"os"
	rf "reflect"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
	"github.com/csgura/fp/test/internal/testpk1"
	"github.com/csgura/fp/try"
	ftry "github.com/csgura/fp/try"
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

// @fp.GenerateTest
var _ = genfp.GenerateFromUntil{
	File: "show_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{"github.com/csgura/fp/seq", "seq"},
	},
	From:     3,
	Until:    genfp.MaxProduct,
	Template: "hello world",
}

type ApiContext struct {
	Map map[string]any
}

type AdaptorAPI interface {
	Context() ApiContext
	TTL() time.Duration
	Timeout() time.Duration
	TestZero() (complex64, time.Time, *string, []int, [3]byte, map[string]any)
	Hello() string
	Tell(target string) fp.Try[string]
	Send(target string) fp.Try[string]
	Active() bool
	IsOk() bool
	Receive(msg string)
	Write(w io.Writer, b []byte) (int, error)
	Create(a string, b int) (int, error)
	Update(a string, b int) (int, error)
	VarArgs(fmtstr string, args ...string)
	IsZero(ptr unsafe.Pointer) bool
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[AdaptorAPI]{
	File:         "adaptor_generated.go",
	Name:         "APIAdaptor",
	Extends:      false,
	Self:         false,
	Getter:       []any{AdaptorAPI.Active, AdaptorAPI.IsOk},
	EventHandler: []any{AdaptorAPI.Receive},
	ValOverride:  []any{AdaptorAPI.Hello, AdaptorAPI.TTL, AdaptorAPI.Context},
	ZeroReturn:   []any{AdaptorAPI.TestZero},
	Options: genfp.AdaptorMethods{
		{

			Method:                  AdaptorAPI.Timeout,
			Prefix:                  "Get",
			ValOverride:             true,
			OmitGetterIfValOverride: false,
		},
		{
			Method:      AdaptorAPI.Receive,
			Prefix:      "On",
			DefaultImpl: genfp.ZeroReturn,
		},
		{
			Method:      AdaptorAPI.Write,
			DefaultImpl: defaultWrite,
		},
		{
			Method: AdaptorAPI.Send,
			DefaultImpl: func() fp.Try[string] {
				return ftry.Success("hello")
			},
		},
		{
			Method: AdaptorAPI.Create,
			DefaultImpl: func(v int) (int, error) {
				return v, nil
			},
		},
		{
			Method:      AdaptorAPI.Update,
			DefaultImpl: 1,
		},
		{
			Method:  AdaptorAPI.Tell,
			Private: true,
			DefaultImpl: func(self AdaptorAPI, target string) fp.Try[string] {
				return self.Send(target)
			},
		},
	},
}

func defaultWrite(self AdaptorAPI, w io.Writer, b []byte) (int, error) {
	return 0, nil
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[AdaptorAPI]{
	File:         "adaptor_generated.go",
	Name:         "APIAdaptorExtends",
	Extends:      true,
	Self:         true,
	Getter:       []any{AdaptorAPI.Hello, AdaptorAPI.Active, AdaptorAPI.IsOk},
	EventHandler: []any{AdaptorAPI.Receive},
	ValOverride:  []any{AdaptorAPI.Hello, AdaptorAPI.TTL},
	Options: genfp.AdaptorMethods{
		{
			Method:      AdaptorAPI.Receive,
			Prefix:      "On",
			DefaultImpl: genfp.ZeroReturn,
		},
		{
			Method:      AdaptorAPI.Write,
			DefaultImpl: defaultWrite,
		},
		{
			Method:      AdaptorAPI.Send,
			DefaultImpl: try.Success("ok"),
		},
		{
			Method:  AdaptorAPI.Tell,
			Private: true,
			DefaultImpl: func(self AdaptorAPI, target string) fp.Try[string] {
				return self.Send(target)
			},
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[AdaptorAPI]{
	File:         "adaptor_generated.go",
	Name:         "APIAdaptorExtendsNotSelf",
	Extends:      true,
	Self:         false,
	Getter:       []any{AdaptorAPI.Hello, AdaptorAPI.Active, AdaptorAPI.IsOk},
	EventHandler: []any{AdaptorAPI.Receive},
	ValOverride:  []any{AdaptorAPI.Hello},
	Options: genfp.AdaptorMethods{
		{
			Method:      AdaptorAPI.Receive,
			Prefix:      "On",
			DefaultImpl: genfp.ZeroReturn,
		},
		{
			Method:      AdaptorAPI.Write,
			DefaultImpl: defaultWrite,
		},
		{
			Method:      AdaptorAPI.Hello,
			ValOverride: true,
		},
	},
}

// @fp.GenerateTest
var _ = genfp.GenerateAdaptor[Tree]{
	File:         "hello",
	Extends:      false,
	Self:         false,
	Getter:       []any{AdaptorAPI.Hello, AdaptorAPI.Active},
	EventHandler: []any{AdaptorAPI.Receive},
	ValOverride:  []any{AdaptorAPI.Hello},
}

// @fp.GenerateTest
var _ = genfp.GenerateAdaptor[io.Closer]{
	File:         "hello",
	Extends:      false,
	Self:         false,
	Getter:       []any{io.Closer.Close, AdaptorAPI.Active},
	EventHandler: []any{AdaptorAPI.Receive},
	ValOverride:  []any{AdaptorAPI.Hello},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[testpk1.AdTester]{
	File:    "adaptor_generated.go",
	Extends: true,
	Self:    true,
	Options: genfp.AdaptorMethods{
		{
			Method:      testpk1.AdTester.Write,
			DefaultImpl: testpk1.DefaultWrite,
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[AdaptorAPI]{
	File:         "adaptor_generated.go",
	Name:         "APIAdaptorNotExtendsWithSelf",
	Extends:      false,
	Self:         true,
	Getter:       []any{AdaptorAPI.Hello, AdaptorAPI.Active, AdaptorAPI.IsOk},
	EventHandler: []any{AdaptorAPI.Receive},
	ValOverride:  []any{AdaptorAPI.Hello},
	Options: genfp.AdaptorMethods{
		{
			Method:      AdaptorAPI.Receive,
			Prefix:      "On",
			DefaultImpl: genfp.ZeroReturn,
		},
		{
			Method:      AdaptorAPI.Write,
			DefaultImpl: defaultWrite,
		},
		{
			Method:      AdaptorAPI.Hello,
			ValOverride: true,
		},
	},
}
