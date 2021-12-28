package immutable_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/immutable"
)

func TestMap(t *testing.T) {
	m := immutable.NewMap[string, int](hash.String)
	m = m.Set("hello", 10)
	fmt.Println(m.Get("hello"))
	fmt.Println(m.Get("world"))

}
