package list_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/list"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
)

func Fibonacci(n1 int, n2 int) fp.List[int] {
	return fp.ListAdaptor[int]{
		GetHead: func() fp.Option[int] {
			return option.Some(n1)
		},
		GetTail: func() fp.List[int] {
			return Fibonacci(n2, n1+n2)
		},
	}
}

func TestFibonacci(t *testing.T) {

	l := Fibonacci(1, 1)

	itr := l
	for i := 0; i < 10 && itr.NonEmpty(); i++ {
		fmt.Println(itr.Head())
		itr = itr.Tail()
	}

	l2 := list.Map(l, as.Curried2(monoid.Product[int]().Combine)(2))
	fmt.Println("multiply 2")
	itr = l2
	for i := 0; i < 10; i++ {
		fmt.Println(itr.Head())
		itr = itr.Tail()
	}

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
	itr = l3
	for i := 0; i < 20; i++ {
		fmt.Println(itr.Head())
		itr = itr.Tail()
	}

}
