package fn1_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/fn1"
)

func TestWith(t *testing.T) {

	a := fn1.Pure[int]("world")

	a2 := fn1.Map(a, func(a1 string) string {
		return "hello " + a1
	})

	a3 := fn1.FlatMap(a, func(a1 string) fp.Func1[int, string] {
		return func(x int) string {
			return fmt.Sprintf("hello %s x %d", a1, x)
		}
	})

	a2(30)
	w := a3(40)
	fmt.Println(w)

}
