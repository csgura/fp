//go:build go1.27

package should

import (
	"reflect"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
)

type The struct {
	testing.TB
}

type TheOptionT[T any] struct {
	The
	t fp.OptionT[T]
}

func (r TheOptionT[T]) BeSome() T {
	r.Helper()
	return BeSome(r.TB, BeSuccess(r.TB, r.t))
}

func (t The) OptionT[T any](v fp.OptionT[T]) TheOptionT[T] {
	return TheOptionT[T]{
		t, v,
	}
}

func Test(t testing.TB, description ...string) The {
	return The{t}
}

type TheValue[T comparable] struct {
	The
	t T
}

func (r TheValue[T]) Equal(other T) {
	r.Helper()
	Equal(r.TB, r.t, other)
}

func (t The) Value[T comparable](v T) TheValue[T] {
	return TheValue[T]{
		t, v,
	}
}

func (t The) BeTrue(b bool) {

	if !b {
		t.Helper()
		t.Fatalf("expected true")
	}
}

func (t The) BeFalse(b bool) {

	if b {
		t.Helper()
		t.Fatalf("expected false")
	}
}

func (t The) Equal[T comparable](a, b T) {

	if a != b {
		t.Helper()
		t.Fatalf("expected [%v], actual [%v]", b, a)
	}
}

func (t The) NotEqual[T comparable](a, b T) {

	if a == b {
		t.Helper()
		t.Fatalf("expected not equal, actual [%v]", a)
	}
}

func (t The) BeZero[T comparable](a T) {
	var zero T
	if a != zero {
		t.Helper()
		t.Fatalf("expected zero, actual [%v]", a)
	}
}

func (t The) NotBeZero[T comparable](a T) {
	var zero T
	if a == zero {
		t.Helper()
		t.Fatalf("expected not zero")
	}
}

func (t The) BeNil(a any) {
	if a != nil {
		rv := reflect.ValueOf(a)
		switch rv.Kind() {
		case reflect.Chan, reflect.Func, reflect.Pointer, reflect.UnsafePointer, reflect.Slice:
			if !rv.IsNil() {
				t.Helper()
				t.Fatalf("expected nil, actual %v", a)
			}
		default:
			t.Helper()
			t.Fatalf("expected nil, actual %v", a)
		}
	}
}

func (t The) NotBeNil(a any) {
	if a == nil {
		t.Helper()
		t.Fatalf("expected not nil, actual [%v]", a)
	}

	rv := reflect.ValueOf(a)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Pointer, reflect.UnsafePointer, reflect.Slice:
		if rv.IsNil() {
			t.Helper()
			t.Fatalf("expected not nil, actual [%v]", a)
		}
	}
}

func (t The) BeSuccess[T any](tt fp.Try[T]) T {
	if tt.IsFailure() {
		t.Helper()
		t.Fatalf("expected success, actual %s", tt.Failed().Get())
	}
	return tt.Get()
}

func (t The) BeFailure[T any](tt fp.Try[T]) error {
	if tt.IsSuccess() {
		t.Helper()
		t.Fatalf("expected error, actual %v", tt.Get())
	}
	return tt.Failed().Get()
}

func (t The) BeSome[T any](tt fp.Option[T]) T {
	if tt.IsEmpty() {
		t.Helper()
		t.Fatalf("expected some, but none")
	}
	return tt.Get()
}

func (t The) BeNone[T any](tt fp.Option[T]) {
	if tt.IsDefined() {
		t.Helper()
		t.Fatalf("expected none, actual %v", tt.Get())
	}
}

func (t The) BeError(err error) {
	if err == nil {
		t.Helper()
		t.Fatal("expected error")
	}
}

func (t The) BeSuccessful[T any](f fp.Future[T], timeout time.Duration) T {
	t.Helper()
	tt := future.Await(f, timeout)
	return BeSuccess(t, tt)
}

func (t The) BeFailed[T any](f fp.Future[T], timeout time.Duration) error {
	t.Helper()
	tt := future.Await(f, timeout)
	return BeFailure(t, tt)
}
