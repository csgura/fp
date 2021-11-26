package product_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/product"
)

func TestGeneric(t *testing.T) {

	tp := product.Tuple3(10, "hello", true)
	hl := tp.ToHList()
	tp2 := hlist.Case2(hl, product.Tuple2[int, string])
	fmt.Printf("%v\n", tp2)
	fmt.Printf("%s\n", hl)

}
