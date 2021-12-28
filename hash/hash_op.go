package hash

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

func hashUint64(value uint64) uint32 {
	hash := value
	for value > 0xffffffff {
		value /= 0xffffffff
		hash ^= value
	}
	return uint32(hash)
}

type hasher[T any] struct {
	fp.Eq[T]
	f func(T) uint32
}

func (r hasher[T]) Hash(t T) uint32 {
	return r.f(t)
}

func New[T any](eq fp.Eq[T], f func(T) uint32) fp.Hashable[T] {
	return hasher[T]{eq, f}
}

// intHasher implements Hasher for int keys.
func Number[T fp.ImplicitNum]() fp.Hashable[T] {
	return New(eq.Given[T](), func(key T) uint32 {
		return hashUint64(uint64(key))
	})
}

var String fp.Hashable[string] = New(eq.Given[string](), func(value string) uint32 {
	var hash uint32
	for i, value := 0, value; i < len(value); i++ {
		hash = 31*hash + uint32(value[i])
	}
	return hash
})

var Bytes fp.Hashable[[]byte] = New(eq.Bytes, func(value []byte) uint32 {
	var hash uint32
	for i, value := 0, value; i < len(value); i++ {
		hash = 31*hash + uint32(value[i])
	}
	return hash
})
