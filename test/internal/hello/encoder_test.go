package hello_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/option"
	"github.com/csgura/fp/test/internal/hello"
)

func TestEncoderOption(t *testing.T) {
	str := hello.EncoderHasOption.Encode(hello.HasOptionMutable{
		Message: "hello",
		Addr:    option.None[string](),
		Phone:   []string{"1234", "2345"},
	}.AsImmutable())

	fmt.Println(str)
}
