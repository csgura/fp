package monoid_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

func TestMonoid(t *testing.T) {
	s := seq.Of(1, 2, 3, 4, 5)
	assert.Equal(s.Reduce(monoid.Sum[int]()), 15)

	assert.Equal(s.Reduce(monoid.Product[int]()), 120)

	s2 := seq.Of("hello", " ", "world")
	assert.Equal(s2.Reduce(monoid.Sum[string]()), "hello world")

	seqseq := seq.Of(seq.Of(1, 2), seq.Of(3, 4), seq.Of(5, 6))
	flatten := seqseq.Reduce(monoid.MergeSeq[int]())
	fmt.Println(flatten)

	optseq := seq.Of(option.Some(1), option.Some(2))
	assert.Equal(optseq.Reduce(monoid.Option(monoid.Sum[int]())).Get(), 3)

	optseq = seq.Of(option.Some(1), option.None[int]())
	assert.True(optseq.Reduce(monoid.Option(monoid.Sum[int]())).IsEmpty())

	vecseq := seq.Of(product.Tuple2(1, 2), product.Tuple2(5, 6))
	assert.Equal(vecseq.Reduce(monoid.Tuple2(monoid.Sum[int](), monoid.Sum[int]())), product.Tuple2(6, 8))

	hlistMonoid := monoid.HCons(monoid.Sum[int](), monoid.HCons(monoid.Sum[string](), monoid.HNil))
	fmt.Println(hlistMonoid.Empty())
	h2 := hlistMonoid.Combine(hlistMonoid.Empty(), hlist.Of2(1, "hello"))
	fmt.Println(hlistMonoid.Combine(h2, hlist.Of2(2, " world")))
}
