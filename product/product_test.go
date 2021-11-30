package product_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/product"
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

func TestGeneric(t *testing.T) {

	tp := product.Tuple3(10, "hello", map[string]any{})
	hl := tp.ToHList()
	tp2 := hlist.Case2(hl, product.Tuple2[int, string])
	fmt.Printf("%v\n", tp2)
	fmt.Printf("%s\n", hl)

	fmt.Printf("hl hasTail = %t\n", hl.HasTail())
	fmt.Printf("hl.Tail hasTail = %t\n", hl.Tail().HasTail())
	fmt.Printf("hl.Tail.Tail hasTail = %t\n", hl.Tail().Tail().HasTail())
	//fmt.Printf("hl.Tail.Tail.Tail hasTail = %t", hl.Tail().Tail().T.HasTail())

	hl = hlist.Of3(11, "hello", map[string]any{})
	fmt.Printf("%s\n", hlist.Reverse3(hl))

	h10 := hlist.Concact(10, hlist.Empty())
	hhello10 := hlist.Concact("hello", h10)
	fmt.Printf("%s\n", hhello10)
}
