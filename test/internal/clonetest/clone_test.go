package clonetest_test

import (
	"testing"

	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/test/internal/clonetest"
)

func TestClone(t *testing.T) {
	r := clonetest.HasReference{
		A: as.Ptr("hello"),
		S: []int{1, 2, 3},
	}

	r2 := clonetest.CloneHasReference().Clone(r)
	r3 := r
	r.S[0] = 2

	assert.True(r.S[0] == r3.S[0])
	assert.True(r.A == r3.A)

	assert.False(r.S[0] == r2.S[0])
	assert.False(r.A == r2.A)
}
