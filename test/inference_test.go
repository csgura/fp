package main_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
)

type hello struct {
}

func (r hello) String() string {
	return "hello"
}

func ToStringer[F any]() func(v F) fmt.Stringer {
	return func(v F) fmt.Stringer {
		var a any = v
		return a.(fmt.Stringer)
	}
}

func TestInference(t *testing.T) {
	v := option.Some(hello{})

	v2 := option.Map(v, as.Ptr[hello])
	v3 := option.Map(v2, as.Interface[*hello, fmt.Stringer])

	v4 := option.Map(v3, option.Some[fmt.Stringer])
	v5 := option.Map(v4, fp.Option[fmt.Stringer].ToSeq)
	fmt.Println(v5)
}
