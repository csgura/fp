package atomic

import (
	"sync/atomic"
	"unsafe"
)

type Reference interface {
	Get() *ValuePtr
	Load() any
	Store(v any)
	CompareAndSwap(original *ValuePtr, newval any) bool
}

var _ Reference = &Value{}

type Value struct {
	value unsafe.Pointer
}

// Get returns AtomicValuePtr
func (r *Value) Get() *ValuePtr {
	data := atomic.LoadPointer(&r.value)
	return (*ValuePtr)(data)
}

// Load returns stored value
func (r *Value) Load() any {
	data := atomic.LoadPointer(&r.value)
	ret := (*ValuePtr)(data)
	return ret.Value()
}

// Store stores value
func (r *Value) Store(v any) {
	tostore := &ValuePtr{v}
	atomic.StorePointer(&r.value, unsafe.Pointer(tostore))
}

// CompareAndSwap comapre current value with original and if it is same , set newvalue
func (r *Value) CompareAndSwap(original *ValuePtr, newval any) bool {
	tostore := &ValuePtr{newval}
	return atomic.CompareAndSwapPointer(&r.value, unsafe.Pointer(original), unsafe.Pointer(tostore))
}

type ValuePtr struct {
	v any
}

// Value returns value
func (r *ValuePtr) Value() any {
	if r == nil {
		return nil
	}
	return r.v
}

func New() Reference {
	return &Value{}
}
