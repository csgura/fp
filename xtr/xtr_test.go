package xtr_test

import (
	"testing"

	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/xtr"
)

func TestExtr(t *testing.T) {
	s := seq.Of(1, 2, 3)
	h := xtr.Head(s)
	assert.Equal(h, option.Some(1))

	i := xtr.Init(s)

	assert.True(eq.Seq(eq.Given[int]()).Eqv(i, seq.Of(1, 2)))

	tl := xtr.Tail(s)

	assert.True(eq.Seq(eq.Given[int]()).Eqv(tl, seq.Of(2, 3)))
}
