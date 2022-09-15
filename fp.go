//go:generate go run github.com/csgura/fp/internal/generator/fp_gen
package fp

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"
	"sync"

	"github.com/csgura/fp/hlist"
)

type Unit struct {
}

func (r Unit) String() string {
	return "()"
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

func (r Tuple1[T1]) ToHList() hlist.Cons[T1, hlist.Nil] {
	return hlist.Concat(r.Head(), hlist.Empty())
}

type Func0[R any] Func1[Unit, R]

func (r Func0[R]) Apply() R {
	return r(Unit{})
}

type Func1[A1, R any] func(a1 A1) R

type Func2[A1, A2, R any] func(a1 A1, a2 A2) R

func (r Func2[A1, A2, R]) Tupled() Func1[Tuple2[A1, A2], R] {
	return func(t Tuple2[A1, A2]) R {
		return r(t.Unapply())
	}
}

func (r Func2[A1, A2, R]) Curried() Func1[A1, Func1[A2, R]] {
	return func(a1 A1) Func1[A2, R] {
		return Func1[A2, R](func(a2 A2) R {
			return r(a1, a2)
		})
	}
}

func (r Func2[A1, A2, R]) ApplyFirst(a1 A1) Func1[A2, R] {
	return func(a2 A2) R {
		return r(a1, a2)
	}
}

func (r Func2[A1, A2, R]) ApplyLast(a2 A2) Func1[A1, R] {
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

func Compose[A, B, C any](f1 Func1[A, B], f2 Func1[B, C]) Func1[A, C] {
	return func(a A) C {
		return f2(f1(a))
	}
}

func Compose2[A, B, C any](f1 Func1[A, B], f2 Func1[B, C]) Func1[A, C] {
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

func Flip[A, B, R any](f Func1[A, Func1[B, R]]) Func1[B, Func1[A, R]] {
	return func(b B) Func1[A, R] {
		return func(a A) R {
			return f(a)(b)
		}
	}
}
