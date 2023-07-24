package ord_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func TestOrd(t *testing.T) {
	c := ord.Tuple3(ord.Given[string](), ord.Given[int](), ord.Given[int]())
	assert.False(c.Less(product.Tuple3("hello", 20, 20), product.Tuple3("hello", 10, 30)))

	assert.True(c.Less(product.Tuple3("hello", 10, 20), product.Tuple3("world", 20, 30)))
	assert.True(c.Less(product.Tuple3("hello", 10, 20), product.Tuple3("hello", 20, 30)))
	assert.True(c.Less(product.Tuple3("hello", 10, 20), product.Tuple3("hello", 10, 30)))

	ic := ord.Option(ord.Given[int]())

	assert.True(ic.Less(option.Some(10), option.Some(20)))
	assert.False(ic.Less(option.Some(30), option.Some(20)))
	assert.False(ic.Less(option.Some(30), option.Some(30)))

	assert.True(ic.Less(option.None[int](), option.Some(20)))
	assert.False(ic.Less(option.Some(20), option.None[int]()))
	assert.False(ic.Less(option.None[int](), option.None[int]()))

	c2 := ord.Tuple2(ord.Given[int](), ord.Given[int]())
	assert.True(c2.Less(product.Tuple2(0, 10), product.Tuple2(1, 0)))
	assert.False(c2.Less(product.Tuple2(1, 10), product.Tuple2(0, 20000)))

	sc := ord.Seq(ord.Given[int]())
	assert.True(sc.Less(seq.Of(1, 2, 3), seq.Of(1, 3)))
	assert.False(sc.Less(seq.Of(1, 3, 3), seq.Of(1, 3)))
	assert.False(sc.Less(seq.Of(1, 3), seq.Of(1, 3)))

	list1 := hlist.Of3(5, 8, 9)
	list2 := hlist.Of3(5, 7, 10)

	hlistc := ord.HCons(ord.Given[int](), ord.HCons(ord.Given[int](), ord.HCons(ord.Given[int](), ord.HNil)))
	assert.False(hlistc.Less(list1, list2))
	assert.False(hlistc.Less(list1, list1))
	assert.True(hlistc.Eqv(list1, list1))

	assert.True(hlistc.Less(list2, list1))

	type Pair struct {
		A int
		B string
	}

	ord.New(eq.Given[Pair](), func(a, b Pair) bool {
		if a.A < b.A {
			return true
		}
		if a.A == b.A {
			return false
		}

		return a.B < b.B
	})

}

type data struct {
	a int
	b int
	c int
}

func (r data) A() int {
	return r.a
}

func (r data) B() int {
	return r.b
}

func (r data) C() int {
	return r.c
}

func TestCombinator(t *testing.T) {
	s := []data{
		{1, 2, 3},
		{1, 1, 2},
		{2, 3, 4},
		{2, 4, 1},
		{2, 2, 5},
		{0, 3, 1},
	}

	res := seq.Sort(s, ord.GivenField(data.A).ThenComparing(ord.GivenField(data.B)))
	//fmt.Println(res)
	assert.Equal(fmt.Sprint(res), "[{0 3 1} {1 1 2} {1 2 3} {2 2 5} {2 3 4} {2 4 1}]")

	res = seq.Sort(s, ord.GivenField(data.A).ThenComparing(ord.GivenField(data.B).Reversed()))
	assert.Equal(fmt.Sprint(res), "[{0 3 1} {1 2 3} {1 1 2} {2 4 1} {2 3 4} {2 2 5}]")

}
