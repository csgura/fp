package fptest

import (
	"testing"
)

func True(t *testing.T, b bool) {

	if !b {
		t.Helper()
		t.Fatalf("assert not true")
	}
}

func False(t *testing.T, b bool) {

	if !b {
		t.Helper()
		t.Fatalf("assert not false")
	}
}

func Equal[T comparable](t *testing.T, a, b T) {

	if a != b {
		t.Helper()
		t.Fatalf("expected [%v], actual [%v]", b, a)
	}
}

func IsNil(t *testing.T, a any) {
	if a != nil {
		t.Helper()
		t.Fatalf("expected nil , actual %v", a)
	}
}

func NotNil(t *testing.T, a any) {
	if a == nil {
		t.Helper()
		t.Fatalf("expected not nil , actual %v", a)
	}
}

func Success(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Fatalf("error = %s", err)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Helper()
		t.Fatal("expected error")
	}
}
