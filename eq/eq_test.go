package eq_test

import (
	"testing"

	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func TestEq(t *testing.T) {

	type Pair struct {
		a int
		b string
	}

	assert.False(eq.Given[int]().Eqv(10, 20))
	assert.True(eq.Given[Pair]().Eqv(Pair{1, "h"}, Pair{1, "h"}))

	assert.False(eq.Option(eq.Given[int]()).Eqv(option.Some(10), option.None[int]()))

	assert.True(eq.Seq(eq.Given[int]()).Eqv(seq.Of(1, 2), seq.Of(1, 2)))

	assert.False(eq.Tuple2(eq.Given[int](), eq.Given[string]()).
		Eqv(product.Tuple2(1, "hello"), product.Tuple2(1, "world")))

	hlist1 := product.Tuple3(1, "2", 3.0).ToHList()
	hlist2 := product.Tuple3(1, "2", 3.2).ToHList()
	//hlist3 := product.Tuple4(1, "2", 3.2, "4").ToHList()

	hlistEq :=
		eq.HCons(eq.Given[int](),
			eq.HCons(eq.Given[string](),
				eq.HCons(eq.Given[float64](), eq.HNil),
			),
		)

	assert.False(hlistEq.Eqv(hlist1, hlist2))

}
