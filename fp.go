//go:generate go run github.com/csgura/fp/internal/generator/fp_gen
package fp

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/csgura/fp/hlist"
)

type Unit struct {
}

func (r Unit) String() string {
	return "()"
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

func (r Tuple1[T1]) ToHList() hlist.Cons[T1, hlist.Nil] {
	return hlist.Concact(r.Head(), hlist.Empty())
}

type Func0[R any] func() R

type Func1[A1, R any] func(a1 A1) R

func (r Func1[A1, R]) Curried() Func1[A1, R] {
	return r
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

type Eq[T any] interface {
	Eqv(a T, b T) bool
}

type Ord[T any] interface {
	Eq[T]
	Less(a T, b T) bool
}

type LessFunc[T any] func(a, b T) bool

func (r LessFunc[T]) Eqv(a, b T) bool {
	return r(a, b) == false && r(b, a) == false
}

func (r LessFunc[T]) Less(a, b T) bool {
	return r(a, b)
}

func Less[T ImplicitOrd]() Ord[T] {
	return LessFunc[T](func(a, b T) bool {
		return a < b
	})
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
