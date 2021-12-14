package hlist_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func plus(a string, b string, c int) string {
	return fmt.Sprintf("%s-%s:%d", a, b, c)
}

func TestHList(t *testing.T) {
	list := product.Tuple3("hello", "world", 10).ToHList()
	tuple3 := hlist.Case3(list, product.Tuple3[string, string, int])
	fp.Println(tuple3)

	tuple2 := hlist.Case2(list, product.Tuple2[string, string])
	fp.Println(tuple2)

	lifted := hlist.Lift3(plus)
	ret := lifted(list)
	println(ret)

	head2 := hlist.Case2(list, hlist.Of2[string, string])
	fp.Println(head2)

	tail2 := hlist.Reverse2(hlist.Case2(hlist.Reverse3(list), hlist.Of2[int, string]))
	fp.Println(tail2)

	var h1 hlist.Header[string] = list
	println(h1)

	// var h2 hlist.Cons[string, hlist.Sealed] = list
	// println(h2)

	a, t1 := hlist.Unapply(list)
	println(a)

	b, t2 := hlist.Unapply(t1)
	println(b)

	c, _ := hlist.Unapply(t2)
	println(c)

	// hlist.Unapply(t3)

	list.Foreach(func(v any) {
		fmt.Println("v = ", v)
	})

	slist := hlist.Fold(list, seq.Of[any](), func(s fp.Seq[any], v any) fp.Seq[any] {
		return s.Append(v)
	})

	fmt.Println(slist)

	println(String(list))

	println(hlist.Fold(list, "", func(s string, v any) string {
		return s + " :: " + fmt.Sprint(v)
	}))

	println(ShowCons(Sprint[string](), ShowCons(Sprint[string](), ShowCons(Sprint[int](), ShowNil))).Show(list))
}

// func String(list hlist.HList) string {
// 	switch v := list.(type) {
// 	case hlist.Nil:
// 		return "Nil"
// 	case hlist.Cons[any, hlist.HList]:
// 		return fmt.Sprintf("%v :: %s", v.Head(), String(v.Tail()))
// 	}
// 	panic("not possible")
// }

func String(list hlist.HList) string {
	if list.IsNil() {
		return "Nil"
	}

	h, t := list.Unapply()
	return fmt.Sprint(h) + " :: " + String(t)
}

type Show[T any] interface {
	Show(t T) string
}

type ShowFunc[T any] func(T) string

func Sprint[T any]() Show[T] {
	return ShowFunc[T](func(v T) string {
		return fmt.Sprint(v)
	})
}

var ShowNil Show[hlist.Nil] = ShowFunc[hlist.Nil](func(v hlist.Nil) string {
	return "Nil"
})

func (r ShowFunc[T]) Show(t T) string {
	return r(t)
}

func ShowCons[H any, T hlist.HList](headShow Show[H], tailShow Show[T]) Show[hlist.Cons[H, T]] {
	return ShowFunc[hlist.Cons[H, T]](func(list hlist.Cons[H, T]) string {
		return headShow.Show(list.Head()) + " :: " + tailShow.Show(list.Tail())
	})
}

func TestHListEq(t *testing.T) {
	list := product.Tuple3("hello", "world", 10).ToHList()
	list2 := product.Tuple3("hello", "world", 10).ToHList()
	list3 := product.Tuple3("hello", "world", 11).ToHList()

	assert.True(eq.HCons(eq.Given[string](), eq.HCons(eq.Given[string](), eq.HCons(eq.Given[int](), eq.HNil))).Eqv(list, list2))
	assert.True(!eq.HCons(eq.Given[string](), eq.HCons(eq.Given[string](), eq.HCons(eq.Given[int](), eq.HNil))).Eqv(list, list3))

}
