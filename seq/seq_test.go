package seq_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func TestSeq(t *testing.T) {
	s := seq.Of(10, 2, 23, 15, 9, 99)
	s = s.Sort(fp.Less[int])

	s.Foreach(fp.Println[int])

	s2 := seq.Of("A", "B", "C", "D", "E")

	zipped := seq.Zip(s, s2)
	zipped.Foreach(fp.Println[fp.Tuple2[int, string]])

	l := seq.Of(product.Tuple2("A", 10), product.Tuple2("A", 12), product.Tuple2("B", 20), product.Tuple2("B", 30))
	m := seq.GroupBy(l, func(t fp.Tuple2[string, int]) string {
		return t.I1
	})
	for k, v := range m {
		fmt.Printf("key = %s, v = %v\n", k, v)
	}

}
