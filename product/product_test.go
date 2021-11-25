package product_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/product"
)

func TestGeneric(t *testing.T) {

	gen2 := product.Concact(100, product.HNil{})
	fmt.Printf("%v\n", gen2)

}
