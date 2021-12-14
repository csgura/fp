package ord_test

import (
	"testing"

	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
)

func assertTrue(b bool) {
	if !b {
		panic("assert fail")
	}
}

func assertFalse(b bool) {
	if b {
		panic("assert fail")
	}
}

func TestOrd(t *testing.T) {
	c := ord.Tuple3(ord.Given[string](), ord.Given[int](), ord.Given[int]())
	assertFalse(c.Less(product.Tuple3("hello", 20, 20), product.Tuple3("hello", 10, 30)))

	assertTrue(c.Less(product.Tuple3("hello", 10, 20), product.Tuple3("world", 20, 30)))
	assertTrue(c.Less(product.Tuple3("hello", 10, 20), product.Tuple3("hello", 20, 30)))
	assertTrue(c.Less(product.Tuple3("hello", 10, 20), product.Tuple3("hello", 10, 30)))

	ic := ord.Option(ord.Given[int]())

	assertTrue(ic.Less(option.Some(10), option.Some(20)))
	assertFalse(ic.Less(option.Some(30), option.Some(20)))
	assertFalse(ic.Less(option.Some(30), option.Some(30)))

	assertTrue(ic.Less(option.None[int](), option.Some(20)))
	assertFalse(ic.Less(option.Some(20), option.None[int]()))
	assertFalse(ic.Less(option.None[int](), option.None[int]()))

}
