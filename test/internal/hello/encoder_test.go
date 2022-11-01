package hello_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
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

func TestHListInsideHList(t *testing.T) {
	v := hello.HListInsideHListMutable{
		Tp:    as.Tuple("10", 10),
		Value: "20",
		Hello: hello.WorldMutable{
			Message:   "message is sparta",
			Timestamp: time.Now(),
		}.AsImmutable(),
	}.AsImmutable()

	str := hello.ShowHListInsideHList.Show(v)

	fmt.Println(str)
	res := hello.ReadHListInsideHList.Read(str)

	res.Failed().Foreach(fp.Println[error])
	assert.True(res.IsSuccess())
	fmt.Println(res)
}
