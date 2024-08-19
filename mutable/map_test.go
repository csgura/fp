package mutable_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/iterator"
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

func TestMapForAll(t *testing.T) {
	m := map[string]bool{
		"1": false,
		"2": true,
	}

	allTrue := mutable.MapOf(m).Values().ForAll(fp.Id)
	assert.False(allTrue)

	anyTrue := mutable.MapOf(m).Values().Exists(fp.Id)
	assert.True(anyTrue)

	allTrue = iterator.FromMapValue(m).ForAll(eq.GivenValue(true))
	assert.False(allTrue)

	m = map[string]bool{
		"1": true,
		"2": true,
	}

	allTrue = mutable.MapOf(m).Values().Exists(fp.Id)
	assert.True(allTrue)

	allTrue = iterator.FromMapValue(m).ForAll(eq.GivenValue(true))
	assert.True(allTrue)

	m = map[string]bool{}

	allTrue = mutable.MapOf(m).Values().ForAll(fp.Id)
	assert.True(allTrue)

	anyTrue = mutable.MapOf(m).Values().Exists(fp.Id)
	assert.False(anyTrue)

}

func TestMapIterator(t *testing.T) {

	m := map[string]bool{
		"1": false,
		"2": true,
		"3": false,
		"4": true,
	}

	itr := fp.IteratorOfGoMap(m).Filter(func(t fp.Tuple2[string, bool]) bool {
		return t.Last()
	})

	v := itr.Next()

	fmt.Printf("v = %v\n", v)

	runtime.GC()
	fmt.Println("after gc1")

	fmt.Printf("v = %v\n", itr.Next())
	runtime.GC()
	fmt.Println("after gc2")

}
