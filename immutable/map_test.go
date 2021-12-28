package immutable_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/immutable"
)

func TestMap(t *testing.T) {
	m := immutable.Map(hash.String,
		as.Tuple("gura", 10),
		as.Tuple("world", 20),
	)
	m = m.Updated("hello", 10)
	fmt.Println(m.Get("hello"))
	fmt.Println(m.Get("world"))

	m.Iterator().Foreach(fp.Println[fp.Tuple2[string, int]])

}
