package seq_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func TestSeq(t *testing.T) {
	s := seq.Of(10, 2, 23, 15, 9, 99)
	s = seq.Sort(s, ord.Given[int]())

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

	seqf := seq.Of(as.Func1(func(v int) int {
		return v + 2
	}), as.Func1(func(v int) int {
		return v * 3
	}))

	apres := seq.Ap(seqf, seq.Of(1, 2, 3))
	fmt.Println(apres)
}

func TestCompileError(t *testing.T) {
	res := option.Applicative2(func(a int, b int) int {
		println(a, b)
		return 10
	}).
		Ap(1).
		Ap(20)
	fmt.Println(res)

}

func TestAp(t *testing.T) {

	plus := func(a int, b int) int {
		return a + b
	}
	s := seq.Of(1, 2, 3, 4)

	s2 := seq.Map(s, as.Func2(plus).Curried()(2))

	s3 := seq.Map(s2, as.Func2(plus).Curried()(3))
	s3.Foreach(fp.Println[int])

	plus3 := func(a, b, c int) int {
		return a + b + c
	}

	f1 := seq.Of(as.Func3(plus3).Curried())
	f2 := seq.Ap(f1, seq.Of(1, 2))
	f3 := seq.Ap(f2, seq.Of(2, 3))
	f4 := seq.Ap(f3, seq.Of(3, 4))
	f4.Foreach(fp.Println[int])
}
