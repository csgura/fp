package immutable_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/mutable"
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

	s := immutable.Set(hash.String, "hello", "world")
	s2 := s.Incl("kkk")
	s.Iterator().Foreach(fp.Println[string])
	s2.Iterator().Foreach(fp.Println[string])

	m2 := mutable.Map[string, int]{"gura": 100, "hello": 200}
	m2["world"] = 200

	m2.Iterator().Foreach(fp.Println[fp.Tuple2[string, int]])

	m3 := immutable.MapBuilder[string, int](hash.String).Add("hello", 10).Add("world", 20).Build()
	m3.Iterator().Foreach(fp.Println[fp.Tuple2[string, int]])

}
