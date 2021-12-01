package future_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/promise"
)

func TestFuture(t *testing.T) {

	p := promise.New[fp.Option[string]]()

	p.Success(option.Some("hello"))
	fp.Println(p.Future())
}
