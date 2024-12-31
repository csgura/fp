package should

import (
	"testing"

	"github.com/csgura/fp"
)

func BeTrue(t testing.TB, b bool) {

	if !b {
		t.Helper()
		t.Fatalf("expected true")
	}
}

func BeFalse(t testing.TB, b bool) {

	if !b {
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
		t.Fatalf("expected not equal, actual %v", a)
	}
}

func BeNil(t testing.TB, a any) {
	if a != nil {
		t.Helper()
		t.Fatalf("expected nil, actual %v", a)
	}
}

func NotBeNil(t testing.TB, a any) {
	if a == nil {
		t.Helper()
		t.Fatalf("expected not nil, actual %v", a)
	}
}

func BeSuccess[T any](t testing.TB, tt fp.Try[T]) {
	if tt.IsFailure() {
		t.Helper()
		t.Fatalf("expected success, actual %s", tt.Failed().Get())
	}
}

func BeFailure[T any](t testing.TB, tt fp.Try[T]) {
	if tt.IsSuccess() {
		t.Helper()
		t.Fatalf("expected error, actual %v", tt.Get())
	}
}

func BeSome[T any](t testing.TB, tt fp.Option[T]) {
	if tt.IsEmpty() {
		t.Helper()
		t.Fatalf("expected some, but none")
	}
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