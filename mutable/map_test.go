package mutable_test

import (
	"testing"

	"github.com/csgura/fp"
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

func TestMapZeroValue(t *testing.T) {
	var m fp.Map[string, int]

	assert.False(m.Get("hello").IsDefined())

	var m2 = m.Updated("hello", 10)
	assert.False(m.Get("hello").IsDefined())
	assert.True(m2.Get("hello").IsDefined())

	var s fp.Set[string]

	assert.False(s.Contains("hello"))

	s2 := s.Incl("hello")
	assert.False(s.Contains("hello"))
	assert.True(s2.Contains("hello"))

}
