// Code generated by gombok, DO NOT EDIT.
package gendebug

import (
	"fmt"
)

type ContainerAdaptor[T fmt.Stringer] struct {
	DoGet func() T
}

func (r *ContainerAdaptor[T]) Get() T {

	if r.DoGet != nil {
		return r.DoGet()
	}

	panic("ContainerAdaptor[T].Get not implemented")
}
