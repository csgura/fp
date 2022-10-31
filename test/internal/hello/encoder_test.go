package hello_test

import (
	"fmt"
	"testing"
	"time"

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

func TestShow(t *testing.T) {
	v := hello.WorldMutable{
		Message:   "hello",
		Timestamp: time.Now(),
	}.AsImmutable()

	fmt.Println(hello.ShowWorld.Show(v))
}
