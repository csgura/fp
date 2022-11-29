package atomic_test

import (
	"testing"

	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/internal/atomic"
)

type Promise struct {
	value atomic.Value
}

func TestCopy(t *testing.T) {

	p := Promise{}

	zero := p.value.Load()
	assert.True(zero == nil)

	p.value.Store(10)

	p2 := p

	v2 := p2.value.Load().(int)

	assert.Equal(v2, 10)

	p2.value.Store(20)

	v1 := p.value.Load().(int)
	assert.Equal(v1, 10)

}

type PromiseRef struct {
	value atomic.Reference
}

func TestCopy2(t *testing.T) {

	p := PromiseRef{atomic.New()}

	zero := p.value.Load()
	assert.True(zero == nil)

	p.value.Store(10)

	p2 := p

	v2 := p2.value.Load().(int)

	assert.Equal(v2, 10)

	p2.value.Store(20)

	v1 := p.value.Load().(int)
	assert.Equal(v1, 20)

}
