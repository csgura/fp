//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package fp

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/csgura/fp/genfp"
)

type Unit struct {
}

func (r Unit) String() string {
	return "()"
}

func (r Unit) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (r *Unit) UnmarshalJSON(data []byte) error {
	return nil
}

type Tuple1[T1 any] struct {
	I1 T1
}

func (r Tuple1[T1]) Head() T1 {
	return r.I1
}

func (r Tuple1[T1]) Tail() Unit {
	return Unit{}
}

type Named interface {
	Name() string
}

type NamedField[T any] interface {
	Named
	Value() T
	Tag() string
}

type RuntimeNamed[T any] Tuple3[string, T, string]

func (r RuntimeNamed[T]) Name() string {
	return r.I1
}

func (r RuntimeNamed[T]) Value() T {
	return r.I2
}

func (r RuntimeNamed[T]) WithValue(v T) RuntimeNamed[T] {
	return RuntimeNamed[T]{r.I1, v, r.I3}
}

func (r RuntimeNamed[T]) Tag() string {
	return r.I3
}

func (r RuntimeNamed[T]) WithTag(v string) RuntimeNamed[T] {
	return RuntimeNamed[T]{r.I1, r.I2, v}
}

type Labelled1[T1 Named] struct {
	I1 T1
}

func (r Labelled1[T1]) Head() T1 {
	return r.I1
}

func (r Labelled1[T1]) Tail() Unit {
	return Unit{}
}

type Supplier[R any] func() R

type Predicate[T any] func(T) bool

func (r Predicate[T]) Negate() Predicate[T] {
	return func(t T) bool {
		return !r(t)
	}
}

func (r Predicate[T]) And(and Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return r(t) && and(t)
	}
}

func (r Predicate[T]) Or(or Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return r(t) || or(t)
	}
}

func Not[T any](f Predicate[T]) Predicate[T] {
	return func(v T) bool {
		return !f(v)
	}
}

func And[T any](flist ...func(v T) bool) Predicate[T] {
	return func(v T) bool {
		for _, f := range flist {
			if !f(v) {
				return false
			}
		}
		return true
	}
}

func Or[T any](flist ...func(v T) bool) Predicate[T] {
	return func(v T) bool {
		for _, f := range flist {
			if f(v) {
				return true
			}
		}
		return false
	}
}

type PartialFunc[T, R any] struct {
	IsDefinedAt func(T) bool
	Apply       func(T) R
}

func (r PartialFunc[T, R]) Unapply() (func(T) bool, func(T) R) {
	return r.IsDefinedAt, r.Apply
}

func (r PartialFunc[T, R]) OrElse(other PartialFunc[T, R]) PartialFunc[T, R] {
	return PartialFunc[T, R]{
		IsDefinedAt: func(t T) bool {
			return r.IsDefinedAt(t) || other.IsDefinedAt(t)
		},
		Apply: func(t T) R {
			if r.IsDefinedAt(t) {
				return r.Apply(t)
			}
			return other.Apply(t)
		},
	}
}

type Func0[R any] Func1[Unit, R]

func (r Func0[R]) Apply() R {
	return r(Unit{})
}

type Func1[A1, R any] func(a1 A1) R

type Func2[A1, A2, R any] func(a1 A1, a2 A2) R

func (r Func2[A1, A2, R]) Widen() func(a1 A1, a2 A2) R {
	return r
}

// func (r Func2[A1, A2, R]) Curried() Func1[A1, Func1[A2, R]] {
// 	return func(a1 A1) Func1[A2, R] {
// 		return Func1[A2, R](func(a2 A2) R {
// 			return r(a1, a2)
// 		})
// 	}
// }

func (r Func2[A1, A2, R]) ApplyFirst(a1 A1) func(A2) R {
	return func(a2 A2) R {
		return r(a1, a2)
	}
}

func (r Func2[A1, A2, R]) ApplyLast(a2 A2) func(A1) R {
	return func(a1 A1) R {
		return r(a1, a2)
	}
}

func Println[T any](v T) {
	fmt.Println(v)
}

func ToString[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

func TypeName[T any]() string {
	var zero *T
	return reflect.TypeOf(zero).Elem().String()
}

type ImplicitUInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type ImplicitInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type ImplicitFloat interface {
	~float32 | ~float64
}

type ImplicitNum interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type ImplicitOrd interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func ParseInt(v string) Try[int] {
	ret, err := strconv.Atoi(v)
	if err != nil {
		return Failure[int](err)
	}
	return Success(ret)
}

func Min[T ImplicitOrd](a1 T, a2 T) T {
	if a1 < a2 {
		return a1
	}
	return a2
}

func Max[T ImplicitOrd](a1 T, a2 T) T {
	if a1 > a2 {
		return a1
	}
	return a2
}

var ErrOptionEmpty = Error(http.StatusNotFound, "Option.empty")
var ErrTryNotFailed = Error(http.StatusBadRequest, "Success.failed not supported")
var ErrFutureNotFailed = Error(http.StatusBadRequest, "Future.Failed not completed with a error")

type errorCode struct {
	code    int
	message string
	cause   error
}

func (r *errorCode) Error() string {
	return r.message
}

func (r *errorCode) Unwrap() error {
	return r.cause
}

func (e *errorCode) StatusCode() int {
	return e.code
}

func (e *errorCode) ErrorTitle() string {
	return http.StatusText(e.code)
}

func Error(code int, fmtStr string, args ...any) error {

	var cause error = nil

	for _, v := range args {
		if e, ok := v.(error); ok {
			cause = e
		}
	}

	return &errorCode{
		code:    code,
		message: fmt.Sprintf(fmtStr, args...),
		cause:   cause,
	}
}

type panicError struct {
	panicMessage any
	stack        []byte
}

func (r *panicError) Error() string {
	return fmt.Sprintf("%v %v", r.panicMessage, string(r.stack))
}

func (r *panicError) Panic() any {
	return r.panicMessage
}

func (r *panicError) Stack() []byte {
	return r.stack
}

func (e *panicError) StatusCode() int {
	return http.StatusInternalServerError
}

func PanicError(message any) error {
	return &panicError{
		panicMessage: message,
		stack:        debug.Stack(),
	}
}

// . 하고는  아규먼트 순서가 반대
//
//	g(f(_)) == g . f   ==  Compose f g  ==  f AndThen g
//
// go 는 AndThen 메소드를 Func1 타입에 정의할 수 없음.
func Compose[A, B, C any](f1 func(A) B, f2 func(B) C) func(A) C {
	return func(a A) C {
		return f2(f1(a))
	}
}

func Compose2[A, B, C any](f1 func(A) B, f2 func(B) C) func(A) C {
	return func(a A) C {
		return f2(f1(a))
	}
}

func Id[T any](t T) T {
	return t
}

func ConvertNumber[From, To ImplicitNum](f From) To {
	return To(f)
}

func Memoize[T any](f func() T) Func0[T] {
	once := sync.Once{}
	var ret T
	return func(Unit) T {
		once.Do(func() {
			ret = f()
		})
		return ret
	}
}

// ( A -> B -> R  ) -> B -> A -> R
func Flip[A, B, R any](f Func1[A, Func1[B, R]]) Func1[B, Func1[A, R]] {
	return func(b B) Func1[A, R] {
		return func(a A) R {
			return f(a)(b)
		}
	}
}

// flip 의 tupled 버젼
// ( ( A , B ) -> R ) -> B -> A -> R
func Flip2[A, B, R any](f func(A, B) R) Func1[B, Func1[A, R]] {
	return func(b B) Func1[A, R] {
		return func(a A) R {
			return f(a, b)
		}
	}
}

func IsInstanceOf[T, I any](v I) bool {
	if _, ok := any(v).(T); ok {
		return true
	}
	return false
}

func New[F, T any](nf func(F) T) T {
	var zero F
	return nf(zero)
}

func Builder[T interface{ Builder() B }, B any]() B {
	var zero T
	return zero.Builder()
}

const GenericKindStruct = "Struct"
const GenericKindTuple = "Tuple"
const GenericKindNewType = "NewType"

// Generic 은 shapeless generic trait 같은 것
//
// T : Struct 나 Tuple 등의 타입
// type NewType SomeType  같은 경우에는 NewType 이 T 에 해당
//
// Repr :Generic representation ,  보통은 HList 형태
// 다만 NewType 패턴의 경우,  기존타입임
type Generic[T, Repr any] struct {
	// Type 은,  T 의 타입이름.  앞에 패키지 이름 붙어 있음.
	Type string

	// Kind 는 T 타입의 종류를 표시함. GenericKindStruct, GenericKindTuple , GenericKindNewType 중의 하나
	Kind string

	// To 는  T를  Generic representation 으로 변환하는 함수
	To func(T) Repr

	// From 은 Generic represantation 을 T로 변환하는 함수
	From func(Repr) T
}

func Zero[T any]() T {
	var zero T
	return zero
}

// https://hackage.haskell.org/package/base/docs/Prelude.html#v:const
// const a b always evaluates to a, ignoring its second argument.
// A -> B -> A
func Const[B, A any](a A) func(b B) A {
	return func(b B) A {
		return a
	}
}

func ConstS[B, A any](f func() A) func(b B) A {
	return func(b B) A {
		return f()
	}
}

// (  A -> B -> A  ) -> B -> A -> A  인 함수인데
// flip 함수가 ( A -> B -> C ) -> B -> A -> C  이니까..
// 사실 같은거고 , 중복 정의할 필요가 없는데
// option.Map(optA ,  fp.With(A.WithField, "2"))
// 형태로 코딩하기 위해 정의
func With[A, B any](withf func(A, B) A, v B) func(A) A {
	return Flip2(withf)(v)
}

// With 와 마찬가지로 Flip2 와 동일한 함수 인데
// option.Filter(optA, fp.Test(A.Contains, "2"))
// 형태로 코딩하기 위해 정의
func Test[A, B any](testf func(A, B) bool, v B) Predicate[A] {
	return Predicate[A](Flip2(testf)(v))
}

// option.Filter(optA, fp.TestWith(A.GetField)(eq.GivenValue(17)))
// pf아규먼트가 커링되어 있는데, 이유는,  커링되어 있으면 gopls 에서  타입추론되어서 자동완성이 되지만
// 커링 안되어 있으면 getter 로 부터 pf 타입을 추론하지 못하기 때문
// 가독성은 커링안되어 있는게 보기가 더 좋은데 어쩔 수 없지.
func TestWith[A, B any](getter func(A) B) func(pf Predicate[B]) Predicate[A] {
	return func(pf Predicate[B]) Predicate[A] {
		return func(a A) bool {
			return pf(getter(a))
		}
	}
}

type Deref[T any] interface {
	Deref() T
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "func_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    3,
	Until:   genfp.MaxFunc,
	Template: `
type Func{{.N}}[{{TypeArgs 1 .N}}, R any] func({{TypeArgs 1 .N}}) R
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "func_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    3,
	Until:   genfp.MaxFunc,
	Template: `
func (r Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplyFirst{{dec .N}}({{DeclArgs 1 (dec .N)}}) func(A{{.N}}) R {
	return func(a{{.N}} A{{.N}}) R {
		return r({{CallArgs 1 .N}})
	}
}

func (r Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplyLast{{dec .N}}({{DeclArgs 2 .N}}) func(A1) R {
	return func(a1 A1) R {
		return r({{CallArgs 1 .N}})
	}
}

func (r Func{{.N}}[{{TypeArgs 1 .N}}, R]) Widen() func({{TypeArgs 1 .N}}) R{
	return r
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "func_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    3,
	Until:   genfp.MaxCompose,
	Template: `
func Compose{{.N}}[{{TypeArgs 1 .N}}, R any]({{FuncChain 1 .N}}) Func1[A1, R] {
	return Compose2(f1, Compose{{dec .N}}({{CallArgs 2 .N "f"}}))
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "func_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    2,
	Until:   genfp.MaxFunc,
	Template: `
func Id{{.N}}[{{TypeArgs 1 (dec .N)}}, R any]({{DeclArgs 1 (dec .N)}}, r R) R {
	return r
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "tuple_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "fmt", Name: "fmt"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
{{define "Receiver"}}func(r Tuple{{.N}}[{{TypeArgs 1 .N "T"}}]){{end}}

type Tuple{{.N}}[{{TypeArgs 1 .N "T"}} any] struct {
	{{- range $idx := Range 1 .N}}
		I{{$idx}} T{{$idx}}
	{{- end}}
}

{{template "Receiver" .}} Head() T1 {
	return r.I1
}

{{template "Receiver" .}} Last() T{{.N}} {
	return r.I{{.N}}
}

{{template "Receiver" .}} Init() ({{TypeArgs 1 (dec .N) "T"}}) {
	return {{CallArgs 1 (dec .N) "r.I"}}
}

{{template "Receiver" .}} Tail() ({{TypeArgs 2 .N "T"}}) {
	return {{CallArgs 2 .N "r.I"}}
}

{{template "Receiver" .}} String() string {
	return fmt.Sprintf("({{FormatStr 1 .N}})", {{CallArgs 1 .N "r.I"}})
}

{{template "Receiver" .}} Unapply() ({{TypeArgs 1 .N "T"}}) {
	return {{CallArgs 1 .N "r.I"}}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "labelled_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "fmt", Name: "fmt"},
	},
	From:  2,
	Until: genfp.MaxProduct,
	Template: `
{{define "Receiver"}}func(r Labelled{{.N}}[{{TypeArgs 1 .N "T"}}]){{end}}

type Labelled{{.N}}[{{TypeArgs 1 .N "T"}} Named] struct {
	{{- range $idx := Range 1 .N}}
		I{{$idx}} T{{$idx}}
	{{- end}}
}

{{template "Receiver" .}} Head() T1 {
	return r.I1
}

{{template "Receiver" .}} Last() T{{.N}} {
	return r.I{{.N}}
}

{{template "Receiver" .}} Init() ({{TypeArgs 1 (dec .N) "T"}}) {
	return {{CallArgs 1 (dec .N) "r.I"}}
}

{{template "Receiver" .}} Tail() ({{TypeArgs 2 .N "T"}}) {
	return {{CallArgs 2 .N "r.I"}}
}

{{template "Receiver" .}} String() string {
	return fmt.Sprintf("({{FormatStr 1 .N}})", {{CallArgs 1 .N "r.I"}})
}

{{template "Receiver" .}} Unapply() ({{TypeArgs 1 .N "T"}}) {
	return {{CallArgs 1 .N "r.I"}}
}
	`,
}
