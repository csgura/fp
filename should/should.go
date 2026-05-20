package should

import (
	"reflect"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
)

func BeTrue(t testing.TB, b bool) {

	if !b {
		t.Helper()
		t.Fatalf("expected true")
	}
}

func BeFalse(t testing.TB, b bool) {

	if b {
		t.Helper()
		t.Fatalf("expected false")
	}
}

func Equal[T comparable](t testing.TB, a, b T) {

	if a != b {
		t.Helper()
		t.Fatalf("expected [%v], actual [%v]", b, a)
	}
}

func NotEqual[T comparable](t testing.TB, a, b T) {

	if a == b {
		t.Helper()
		t.Fatalf("expected not equal, actual [%v]", a)
	}
}

func BeZero[T comparable](t testing.TB, a T) {
	var zero T
	if a != zero {
		t.Helper()
		t.Fatalf("expected zero, actual [%v]", a)
	}
}

func NotBeZero[T comparable](t testing.TB, a T) {
	var zero T
	if a == zero {
		t.Helper()
		t.Fatalf("expected not zero")
	}
}

func BeNil(t testing.TB, a any) {
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

func NotBeNil(t testing.TB, a any) {
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

func BeSuccess[T any](t testing.TB, tt fp.Try[T]) T {
	if tt.IsFailure() {
		t.Helper()
		t.Fatalf("expected success, actual %s", tt.Failed().Get())
	}
	return tt.Get()
}

func BeFailure[T any](t testing.TB, tt fp.Try[T]) error {
	if tt.IsSuccess() {
		t.Helper()
		t.Fatalf("expected error, actual %v", tt.Get())
	}
	return tt.Failed().Get()
}

func BeSome[T any](t testing.TB, tt fp.Option[T]) T {
	if tt.IsEmpty() {
		t.Helper()
		t.Fatalf("expected some, but none")
	}
	return tt.Get()
}

func BeNone[T any](t testing.TB, tt fp.Option[T]) {
	if tt.IsDefined() {
		t.Helper()
		t.Fatalf("expected none, actual %v", tt.Get())
	}
}

func BeError(t testing.TB, err error) {
	if err == nil {
		t.Helper()
		t.Fatal("expected error")
	}
}

func BeSuccessful[T any](t testing.TB, f fp.Future[T], timeout time.Duration) T {
	t.Helper()
	tt := future.Await(f, timeout)
	return BeSuccess(t, tt)
}

func BeFailed[T any](t testing.TB, f fp.Future[T], timeout time.Duration) error {
	t.Helper()
	tt := future.Await(f, timeout)
	return BeFailure(t, tt)
}
