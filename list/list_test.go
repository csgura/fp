package list_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/list"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
)

func Fibonacci(n1 int, n2 int) fp.List[int] {
	return fp.MakeList(
		func() fp.Option[int] {
			return option.Some(n1)
		},
		func() fp.List[int] {
			return Fibonacci(n2, n1+n2)
		},
	)
}

func printFirst10[T any](l fp.List[T]) {
	itr := l
	for i := 0; i < 10 && itr.NonEmpty(); i++ {
		fmt.Println(itr.Head())
		itr = itr.Tail()
	}
}
func TestFibonacci(t *testing.T) {

	l := Fibonacci(1, 1)

	printFirst10(l)

	l2 := list.Map(l, as.Curried2(monoid.Product[int]().Combine)(2))
	fmt.Println("multiply 2")
	printFirst10(l2)

	fmt.Println("Test flatmap")
	l3 := list.FlatMap(l, func(v int) fp.List[int] {
		switch v % 3 {
		case 0:
			return list.Of[int]()
		case 1:
			return list.Of(v)
		case 2:
			return list.Of(v, v)
		}
		panic(fmt.Sprintf("not posssible v : %d", v))
	})
	printFirst10(l3)

	fmt.Println("Test Drop While")
	l4 := l.Iterator().DropWhile(func(v int) bool {
		return v < 100
	}).Take(50).ToList()

	l4 = list.Map(l4, func(v int) int {
		fmt.Printf("check lazy map v : %d\n", v)
		return v
	})

	printFirst10(l4)

	s := seq.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s2 := seq.Map(s, as.Curried2(monoid.Product[int]().Combine)(5))
	f := s2.Find(as.Curried2(ord.Given[int]().Less)(20))
	fmt.Println(f)

	fmt.Println(l.Iterator().Take(20).ToSeq())

	count := list.FoldRight(l, 0, func(v int, sum fp.Lazy[int]) int {
		if v < 100 {
			return sum.Get() + 1
		}
		return 0
	})

	println(count)

}

func TestSum(t *testing.T) {
	l := list.GenerateFrom(1, func(i int) float64 {
		return 1.0 / (float64(i) * float64(i))
	})

	printFirst10(l)

	sum := list.FoldRight(l, 0, func(v float64, sum fp.Lazy[float64]) float64 {

		if v < 0.0000000001 {
			return 0
		}
		return v + sum.Get()

	})

	println(sum)
	println(math.Pi * math.Pi / 6)

	fmt.Println("print list scan")
	l2 := list.Scan(l, 0.0, monoid.Sum[float64]().Combine)

	zip := list.Zip(l2.Iterator().Drop(100).ToList(), l2)
	printFirst10(zip)

	sumOpt := zip.Iterator().Find(as.Func2(func(a float64, b float64) bool {
		return a-b < 0.0000000001
	}).Tupled())

	fmt.Println("sum 1/n^2 = ", sumOpt)

	l = list.GenerateFrom(1, func(i int) float64 {
		return 1.0 / float64(i)
	})

	printFirst10(l)

	fmt.Println("print list scan")
	l2 = list.Scan(l, 0.0, monoid.Sum[float64]().Combine)

	zip = list.Zip(l2.Tail(), l2)
	printFirst10(zip)
	fmt.Println("print list at 100000")

	sumOpt = zip.Iterator().Take(100000).Find(as.Func2(func(a float64, b float64) bool {
		return a-b < 0.0000000001
	}).Tupled())

	fmt.Println("sum 1/n = ", sumOpt)

	iterator.Scan(l.Iterator(), 0.0, monoid.Sum[float64]().Combine).Take(10).Foreach(fp.Println[float64])

}

func TestScan(t *testing.T) {

	l := list.GenerateFrom(1, fp.Id[int])
	list.Scan(l, 0, monoid.Sum[int]().Combine).
		Iterator().Take(10).Foreach(fp.Println[int])

	l = list.Of(1, 2, 3, 4, 5)
	l2 := list.Scan(l, 0, monoid.Sum[int]().Combine)
	printFirst10(l2)

	s := seq.Of(1, 2, 3, 4, 5)
	fmt.Println(seq.Scan(s, 0, monoid.Sum[int]().Combine))

	iterator.Scan(l.Iterator(), 0, monoid.Sum[int]().Combine).Foreach(fp.Println[int])
}
