package list_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
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
}
