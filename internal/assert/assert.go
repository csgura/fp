package assert

import "fmt"

func True(b bool) {
	if !b {
		panic("assert fail")
	}
}

func False(b bool) {
	if b {
		panic("assert fail")
	}
}

func Equal[T comparable](a, b T) {
	if a != b {
		panic(fmt.Sprintf("expected [%v], actual [%v]", b, a))
	}
}

func IsNil(a any) {
	if a != nil {
		panic(fmt.Sprintf("expected nil , actual %v", a))
	}
}

func NotNil(a any) {
	if a == nil {
		panic(fmt.Sprintf("expected not nil , actual %v", a))
	}
}

func Success(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(err error) {
	if err == nil {
		panic("expected error")
	}
}
