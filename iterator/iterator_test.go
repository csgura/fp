package iterator_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
)

func plus(a int, b int) int {
	return a + b
}

func TestIterator(t *testing.T) {
	s := seq.Of(1, 2, 3, 4, 5, 6, 7)
	iterator.Map(seq.Iterator(s), curried.Func2(plus)(2)).TakeWhile(func(v int) bool {
		return v < 7
	}).Foreach(fp.Println[int])

	iterator.FlatMap(seq.Iterator(s), func(v int) fp.Iterator[int] {
		println("v = ", v)
		switch v % 3 {
		case 0:
			return seq.Iterator(seq.Of[int]())
		case 1:
			return seq.Iterator(seq.Of(v))
		case 2:
			return seq.Iterator(seq.Of(v, v, v))
		}
		panic("not possible")
	}).TakeWhile(func(v int) bool {
		return v < 8
	}).Foreach(fp.Println[int])

	k := seq.Of("a", "b", "c")
	v := seq.Of(10, 20, 30, 40, 50)

	fmt.Println(iterator.ToMap(iterator.Zip(seq.Iterator(k), seq.Iterator(v)), hash.String))

	p1, p2 := iterator.Partition(seq.Iterator(s), func(v int) bool {
		return v%2 == 0
	})

	fmt.Println(seq.Collect(p1))
	fmt.Println(seq.Collect(p2))

	p1, p2 = iterator.Span(iterator.Map(seq.Iterator(s), func(v int) int {
		println("before span v= ", v)
		return v
	}), func(v int) bool {
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

func TestToList(t *testing.T) {

	list := iterator.ToList(iterator.Range(0, 10))

	assert.Equal(list.Tail().Head(), 1)
	assert.Equal(list.Tail().Head(), 1)

	list.Foreach(fp.Println[int])
	list.Foreach(fp.Println[int])

}

func TestMax(t *testing.T) {
	i := iterator.Range(0, 10)
	max := iterator.Max(i, ord.Given[int]())
	assert.Equal(max, option.Some(9))

	i = iterator.Range(0, 10)
	min := iterator.Min(i, ord.Given[int]())
	assert.Equal(min, option.Some(0))

}

func TestReverse(t *testing.T) {
	i := iterator.Range(0, 10).ToSeq()

	i = iterator.ReverseSeq(i).ToSeq()
	assert.Equal(i[0], 9)
	assert.Equal(len(i), 10)
}

func TestZipLeft(t *testing.T) {
	l := iterator.Range(0, 10)
	r := iterator.Range(0, 5)

	optr := iterator.Map(r, option.Some).Concat(iterator.Generate(option.None[int]))
	zl := iterator.Zip(l, optr).ToSeq()
	assert.Equal(len(zl), 10)
}

// range over func
// func TestForLoop(t *testing.T) {
// 	for v := range iterator.Range(0, 100).All() {
// 		fmt.Printf("v = %d\n", v)
// 		if v > 50 {
// 			break
// 		}
// 	}
// }
