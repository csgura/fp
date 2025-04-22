package main_test

import (
	"errors"
	"testing"

	"github.com/csgura/fp/option"
	"github.com/csgura/fp/optiont"
	"github.com/csgura/fp/try"
)

func TestTryCase(t *testing.T) {
	v := try.Success(10)

	switch v {
	case try.IsFailureCase(v):
		t.Fatal("fail")

	case try.IsSuccessCase(v):
		println("success")
	default:
		t.Fatal("fail")
	}

	ov := option.Some(v)
	switch ov {
	case option.IsSomeCaseAnd(ov, try.IsSuccessCase):
		println("some success")
	default:
		t.Fatal("fail")
	}

	tov := try.Success(ov)
	switch tov {
	case try.IsSuccessCaseAnd(tov, option.NestedIsSomeCase(try.IsSuccessCase[int])):
		println("success some success")
	default:
		t.Fatal("fail")
	}

	switch tov {
	case optiont.IsSomeCaseAnd(tov, try.IsSuccessCase):
		println("success some success")
	default:
		t.Fatal("fail")
	}

	v = try.Failure[int](errors.New("too bad"))

	switch v {
	case try.IsSuccessCase(v):
		t.Fatal("success")
	case try.IsFailureCase(v):
		println("fail")
	default:
		t.Fatal("success")
	}

	ov = option.Some(v)
	switch ov {
	case option.IsSomeCaseAnd(ov, try.IsSuccessCase):
		t.Fatal("some success")
	case option.IsSomeCaseAnd(ov, try.IsFailureCase):
		println("some faliure")
	default:
		t.Fatal("fail")
	}

}
