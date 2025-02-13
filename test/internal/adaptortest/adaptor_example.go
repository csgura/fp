package adaptortest

//lint:file-ignore ST1019 test code

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
	"unsafe"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/test/internal/testpk1"
	"github.com/csgura/fp/try"
	ftry "github.com/csgura/fp/try"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type ApiContext struct {
	Map map[string]any
}

type AdaptorAPI interface {
	Context() ApiContext
	TTL() time.Duration
	Timeout() time.Duration
	MaxConn() int
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
	ValOverride:  []any{AdaptorAPI.Hello, AdaptorAPI.TTL, AdaptorAPI.Context, AdaptorAPI.MaxConn},
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

type HTTP interface {
	Send(msg string) (int, error)
}

type HTTP2 interface {
	HTTP
	KeepAlive(v bool)
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[HTTP]{
	File:    "adaptor_generated.go",
	Self:    true,
	Extends: true,
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[HTTP2]{
	File:           "adaptor_generated.go",
	ImplementsWith: seq.Of(genfp.TypeOf[io.Closer]()),
	ExtendsWith: map[string]genfp.TypeTag{
		"Closer": genfp.TypeOf[io.Closer](),
	},
	Embedding:          seq.Of(genfp.TypeOf[HTTPAdaptor]()),
	Delegate:           seq.Of(genfp.DelegatedBy[HTTP]("HTTPAdaptor")),
	Self:               true,
	ExtendsByEmbedding: true,
}

type SpanContextAdaptorToBe struct {
	DefaultContext context.Context
}

type ConnHandler interface {
	NewReader(conn net.Conn)
}

// TODO : 아규먼트 공변
type ConnHandlerAdaptor struct {
	ReaderMaker func(conn io.Reader)
}

// net.Conn 은 io.Reader 기 때문에 , ReaderMaker 호출 가능
func (r *ConnHandlerAdaptor) NewReader(conn net.Conn) {
	r.ReaderMaker(conn)
}

type Sender interface {
	Send(msg string) (int, error)
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File: "example_adaptor.go",
	Name: "SenderAdaptor", // 생성될 adaptor 의 이름. 지정하지 않으면 Sender + Adaptor
}

type Status interface {
	Active() bool
	DisplayName() string
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Status]{
	File:       "example_adaptor.go",
	Getter:     []any{Status.Active, Status.DisplayName},
	ZeroReturn: []any{Status.Active},
}

type Handler interface {
	ReceiveEvent(event string)
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Handler]{
	File:         "example_adaptor.go",
	EventHandler: []any{Handler.ReceiveEvent},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Status]{
	File:       "example_adaptor.go",
	Name:       "StatusAdaptorZero",
	Getter:     []any{Status.Active, Status.DisplayName},
	ZeroReturn: []any{Status.Active},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Status]{
	File:        "example_adaptor.go",
	Name:        "StatusAdaptorVal",
	ValOverride: []any{Status.Active, Status.DisplayName},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Status]{
	File:        "example_adaptor.go",
	Name:        "StatusAdaptorValGetter",
	Getter:      []any{Status.Active, Status.DisplayName},
	ValOverride: []any{Status.Active, Status.DisplayName},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Status]{
	File:        "example_adaptor.go",
	Name:        "StatusAdaptorCustom",
	Getter:      []any{Status.Active},
	ValOverride: []any{Status.DisplayName},
	Options: []genfp.ImplOption{
		{
			Method: Status.DisplayName,
			DefaultImpl: func() string {
				return "Inactive"
			},
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Handler]{
	File:         "example_adaptor.go",
	Name:         "HandlerCustom",
	EventHandler: []any{Handler.ReceiveEvent},
	Options: []genfp.ImplOption{
		{
			Method: Handler.ReceiveEvent,
			DefaultImpl: func(v string) {
				fmt.Printf("receive event : %s\n", v)
			},
		},
	},
}

func sendStdout(msg string) (int, error) {
	fmt.Printf("msg = %s\n", msg)
	return len(msg), nil
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File: "example_adaptor.go",
	Name: "SenderCustom",
	Options: []genfp.ImplOption{
		{
			Method:      Sender.Send,
			DefaultImpl: sendStdout,
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File: "example_adaptor.go",
	Name: "Sender42",
	Options: []genfp.ImplOption{
		{
			Method:      Sender.Send,
			DefaultImpl: 42,
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File:    "example_adaptor.go",
	Name:    "SenderExtends",
	Extends: true,
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File:    "example_adaptor.go",
	Name:    "SenderExtendsSelf",
	Extends: true,
	Self:    true,
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File: "example_adaptor.go",
	Name: "SenderSelfArg",
	Options: []genfp.ImplOption{
		{
			Method: Sender.Send,
			DefaultImpl: func(self Sender, msg string) (int, error) {
				return 0, nil
			},
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	File:    "example_adaptor.go",
	Name:    "SenderSelfSelfArg",
	Extends: true,
	Self:    true,
	Options: []genfp.ImplOption{
		{
			Method: Sender.Send,
			DefaultImpl: func(self Sender, msg string) (int, error) {
				return 0, nil
			},
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{
	Extends: true,
	Self:    true,
	Options: []genfp.ImplOption{
		{
			Method: Sender.Send,
			DefaultImpl: func(self Sender, msg string) (int, error) {
				return 0, nil
			},
		},
	},
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Sender]{

	Extends: true,
	Self:    true,
	Name:    "SenderError",
	Options: []genfp.ImplOption{
		{
			Method: Sender.Send,
			DefaultImpl: func(msg int) (int, error) {
				return 0, nil
			},
		},
	},
}

type Invoker interface {
	Invoke(interface{})
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Invoker]{
	File: "example_adaptor.go",
}

// @fp.Generate
var _ = genfp.GenerateAdaptor[Invoker]{
	File:             "example_adaptor.go",
	Name:             "InvokerCheckSelf",
	Extends:          true,
	ExtendsSelfCheck: true,
}

type StringMaker = fmt.Stringer

// @fp.Generate
var _ = genfp.GenerateAdaptor[StringMaker]{
	File:             "example_adaptor.go",
	Extends:          true,
	ExtendsSelfCheck: true,
	Options: []genfp.ImplOption{
		{
			Method: StringMaker.String,
			DefaultImpl: func(self StringMaker) string {
				return "hello world"
			},
		},
	},
}
