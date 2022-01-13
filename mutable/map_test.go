package mutable_test

import (
	"testing"

	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/mutable"
)

func TestCopyOnWriteMap(t *testing.T) {
	m := mutable.CopyOnWriteMap[string, int]{}

	v := m.ComputeIfAbsent("hello", func() int {
		return 10
	})

	assert.Equal(v, 10)

	assert.Equal(m.Get("hello").Get(), 10)
	assert.True(m.Get("hello2").IsEmpty())

}
