package seq_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func TestSeq(t *testing.T) {
	s := seq.Of(10, 2, 23, 15, 9, 99)
	s = s.Sort(ord.Given[int]())

	s.Foreach(fp.Println[int])

	sum := s.Reduce(fp.Sum[int]())
	println("sum = ", sum)
	println("product = ", s.Reduce(fp.Product[int]()))

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

	matrix := seq.Of(product.Tuple2(1, 2), product.Tuple2(2, 3))
	fp.Println(matrix.Reduce(monoid.Tuple2(fp.Sum[int](), fp.Sum[int]())))

	opts := seq.Of(option.Some(1), option.Some(2))
	fp.Println(opts.Reduce(monoid.Option(fp.Sum[int]())))

}
