package hlist_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/product"
)

func plus(a string, b string, c int) string {
	return fmt.Sprintf("%s-%s:%d", a, b, c)
}

func TestHList(t *testing.T) {
	list := hlist.Of3("hello", "world", 10)
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
		return headShow.Show(list.Head()) + " :: " + tailShow.Show(hlist.Tail(list))
	})
}

func TestHListEq(t *testing.T) {
	list := hlist.Of3("hello", "world", 10)
	list2 := hlist.Of3("hello", "world", 10)
	list3 := hlist.Of3("hello", "world", 11)

	assert.True(eq.HCons(eq.Given[string](), eq.HCons(eq.Given[string](), eq.HCons(eq.Given[int](), eq.HNil))).Eqv(list, list2))
	assert.True(!eq.HCons(eq.Given[string](), eq.HCons(eq.Given[string](), eq.HCons(eq.Given[int](), eq.HNil))).Eqv(list, list3))

}

type Table[H hlist.HList] struct {
	records []H
}

func (r *Table[H]) Insert(record H) {
	r.records = append(r.records, record)
}

func (r *Table[H]) Find(p func(H) bool) *H {
	for _, v := range r.records {
		if p(v) {
			return &v
		}
	}
	return nil
}

func TestTable(t *testing.T) {
	tbl := Table[hlist.Cons[string, hlist.Cons[int, hlist.Nil]]]{}
	tbl.Insert(hlist.Of2("hello", 10))

	tbl.Find(hlist.Lift2(func(a string, b int) bool {
		return a == "hello"
	}))
}
