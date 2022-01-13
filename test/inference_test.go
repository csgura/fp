package main_test

import (
	"fmt"
	"testing"
)

func ToStringer[F any]() func(v F) fmt.Stringer {
	return func(v F) fmt.Stringer {
		var a any = v
		return a.(fmt.Stringer)
	}
}

func TestInference(t *testing.T) {
	// v := option.Some(hello{})
	// _ = v
	//fmt.Println(v)
	// 	v2 := option.Map(v, as.Ptr[hello])
	// 	v3 := option.Map(v2, as.Interface[*hello, fmt.Stringer])

	// 	v4 := option.Map(v3, option.Some[fmt.Stringer])
	// 	v5 := option.Map(v4, fp.Option[fmt.Stringer].ToSeq)
	// 	fmt.Println(v5)

	// 	// option.Applicative3(as.UnTupled3(fp.Println[fp.Tuple3[int, string, int]])).
	// 	// 	Ap(1).
	// 	// 	Ap("hello").
	// 	// 	Ap(10)
}

func TestCompileError(t *testing.T) {
	// res := option.Applicative2(func(a int, b int) int {
	// 	fmt.Println(a, b)
	// 	return 10
	// }).
	// 	Ap(1).
	// 	Ap(20)
	// fmt.Println(res)

}

type Wrapper[B any] interface {
	Unwrap() B
}

func Unwrap[B any](a Wrapper[B]) B {
	return a.Unwrap()
}

func Unwrap2[B any, A Wrapper[B]](a A) B {
	return a.Unwrap()
}

func Unwrap3[A Wrapper[B], B any](a A) B {
	return a.Unwrap()
}

type wrapperImpl[B any] struct {
	v B
}

func (r *wrapperImpl[B]) Unwrap() B {
	return r.v
}

func TestDependent(t *testing.T) {
	var a Wrapper[int]

	Unwrap(a)
	Unwrap2[int](a)

	b := &wrapperImpl[int]{10}

	Unwrap[int](b)
	Unwrap2[int](b)
	// Unwrap3[Wrapper[int]](b)
}
