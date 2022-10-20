package value_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/test/value"
)

func TestBuilder(t *testing.T) {
	a := fp.New(value.Hello.Builder).World("world").Hi(0).Build()
	fmt.Println(a)
	fmt.Println(a.WithWorld("No").World())
}
