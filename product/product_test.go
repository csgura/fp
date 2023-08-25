package product_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/unit"
)

type clImpl struct {
}

func (r clImpl) isNil() bool {
	return true
}

func (r clImpl) String() string {
	return "Nil"
}

func (r clImpl) HasTail() bool {
	return false
}

type IceCream struct {
	Name  string
	price int16
	Maker string
}

func hello(a string, b int) {
	fmt.Printf("%s:%d\n", a, b)
}

func format(a string, b int) string {
	return fmt.Sprintf("%s:%d\n", a, b)
}

func returnError(a string, b int) (string, error) {
	return fmt.Sprintf("%s:%d", a, b), nil
}
func TestGeneric(t *testing.T) {

	tp := product.Tuple3(10, "hello", map[string]any{})

	tp = as.Tuple3(10, "hello", map[string]any{})

	age, name, attr := tp.Unapply()
	fmt.Printf("age = %d, name = %s, attr = %v", age, name, attr)

	hl := hlist.Of3(tp.Unapply())
	tp2 := hlist.Case2(hl, product.Tuple2[int, string])
	fmt.Printf("%v\n", tp2)
	fmt.Printf("%s\n", hl)

	fmt.Printf("hl hasTail = %t\n", hl.Tail().IsNil())
	fmt.Printf("hl.Tail hasTail = %t\n", hl.Tail().Tail().IsNil())
	fmt.Printf("hl.Tail.Tail hasTail = %t\n", hl.Tail().Tail().Tail().IsNil())
	//fmt.Printf("hl.Tail.Tail.Tail hasTail = %t", hl.Tail().Tail().T.HasTail())

	hl = hlist.Of3(11, "hello", map[string]any{})
	fmt.Printf("%s\n", hlist.Reverse3(hl))

	h10 := hlist.Concat(10, hlist.Empty())
	hhello10 := hlist.Concat("hello", h10)
	fmt.Printf("%s\n", hhello10)

	ice := IceCream{"hello", 100, "lotte"}
	iceTup := *(*fp.Tuple3[string, int16, string])(unsafe.Pointer(&ice))
	fmt.Printf("%v\n", iceTup)

	ice = *(*IceCream)(unsafe.Pointer(&iceTup))

	hello(product.Tuple2("hello", 10).Unapply())

	s := try.Func2(returnError)("hello", 20)
	s.Foreach(fp.Println[string])

	curried.Func2(try.Func2(returnError))("hello")(30).Foreach(fp.Println[string])

	a, b, _ := tp.Unapply()
	fmt.Printf("a = %d , b = %s\n", a, b)

	as.Tupled2(unit.Func2(hello))(product.Tuple2("a", 20))

	fn := as.Tupled2(format)

	println(fn(product.Tuple2("hello", 10)))

}
