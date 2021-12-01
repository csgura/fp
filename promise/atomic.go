package promise

import (
	"sync/atomic"
	"unsafe"
)

type atomicValue struct {
	value unsafe.Pointer
}

// Get returns AtomicValuePtr
func (r *atomicValue) Get() *atomicValuePtr {
	data := atomic.LoadPointer(&r.value)
	return (*atomicValuePtr)(data)
}

// Load returns stored value
func (r *atomicValue) Load() any {
	data := atomic.LoadPointer(&r.value)
	ret := (*atomicValuePtr)(data)
	return ret.Value()
}

// Store stores value
func (r *atomicValue) Store(v any) {
	tostore := &atomicValuePtr{v}
	atomic.StorePointer(&r.value, unsafe.Pointer(tostore))
}

// CompareAndSwap comapre current value with original and if it is same , set newvalue
func (r *atomicValue) CompareAndSwap(original *atomicValuePtr, newval any) bool {
	tostore := &atomicValuePtr{newval}
	return atomic.CompareAndSwapPointer(&r.value, unsafe.Pointer(original), unsafe.Pointer(tostore))
}

type atomicValuePtr struct {
	v any
}

// Value returns value
func (r *atomicValuePtr) Value() any {
	if r == nil {
		return nil
	}
	return r.v
}
