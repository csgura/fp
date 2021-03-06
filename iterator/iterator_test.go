package iterator_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/seq"
)

func plus(a int, b int) int {
	return a + b
}

func TestIterator(t *testing.T) {
	s := seq.Of(1, 2, 3, 4, 5, 6, 7)
	iterator.Map(s.Iterator(), curried.Func2(plus)(2)).TakeWhile(func(v int) bool {
		return v < 7
	}).Foreach(fp.Println[int])

	iterator.FlatMap(s.Iterator(), func(v int) fp.Iterator[int] {
		println("v = ", v)
		switch v % 3 {
		case 0:
			return seq.Of[int]().Iterator()
		case 1:
			return seq.Of(v).Iterator()
		case 2:
			return seq.Of(v, v, v).Iterator()
		}
		panic("not possible")
	}).TakeWhile(func(v int) bool {
		return v < 8
	}).Foreach(fp.Println[int])

	k := seq.Of("a", "b", "c")
	v := seq.Of(10, 20, 30, 40, 50)

	fmt.Println(iterator.ToMap(iterator.Zip(k.Iterator(), v.Iterator()), hash.String))

	p1, p2 := s.Iterator().Partition(func(v int) bool {
		return v%2 == 0
	})

	fmt.Println(seq.Collect(p1))
	fmt.Println(seq.Collect(p2))

	p1, p2 = iterator.Map(s.Iterator(), func(v int) int {
		println("before span v= ", v)
		return v
	}).Span(func(v int) bool {
		return v < 4
	})

	fmt.Println(seq.Collect(p1))
	fmt.Println(seq.Collect(p2))

}

func TestRange(t *testing.T) {
	iterator.Range(0, 10).Foreach(fp.Println[int])
}

func TestToSet(t *testing.T) {
	iterator.ToSet(iterator.Of("hello", "world", "hello", "merong"), hash.String).Foreach(fp.Println[string])
}
