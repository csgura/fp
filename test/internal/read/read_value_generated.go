// Code generated by gombok, DO NOT EDIT.
package read

import (
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

type ResultBuilder[T any] Result[T]

type ResultMutable[T any] struct {
	Value   T
	Remains string
}

func (r ResultBuilder[T]) Build() Result[T] {
	return Result[T](r)
}

func (r Result[T]) Builder() ResultBuilder[T] {
	return ResultBuilder[T](r)
}

func (r Result[T]) Value() T {
	return r.value
}

func (r Result[T]) WithValue(v T) Result[T] {
	r.value = v
	return r
}

func (r ResultBuilder[T]) Value(v T) ResultBuilder[T] {
	r.value = v
	return r
}

func (r Result[T]) Remains() string {
	return r.remains
}

func (r Result[T]) WithRemains(v string) Result[T] {
	r.remains = v
	return r
}

func (r ResultBuilder[T]) Remains(v string) ResultBuilder[T] {
	r.remains = v
	return r
}

func (r Result[T]) String() string {
	return fmt.Sprintf("Result(value=%v, remains=%v)", r.value, r.remains)
}

func (r Result[T]) AsTuple() fp.Tuple2[T, string] {
	return as.Tuple2(r.value, r.remains)
}

func (r Result[T]) AsMutable() ResultMutable[T] {
	return ResultMutable[T]{
		Value:   r.value,
		Remains: r.remains,
	}
}

func (r ResultMutable[T]) AsImmutable() Result[T] {
	return Result[T]{
		value:   r.Value,
		remains: r.Remains,
	}
}

func (r ResultBuilder[T]) FromTuple(t fp.Tuple2[T, string]) ResultBuilder[T] {
	r.value = t.I1
	r.remains = t.I2
	return r
}

func (r Result[T]) AsMap() map[string]any {
	return map[string]any{
		"value":   r.value,
		"remains": r.remains,
	}
}

func (r ResultBuilder[T]) FromMap(m map[string]any) ResultBuilder[T] {

	if v, ok := m["value"].(T); ok {
		r.value = v
	}

	if v, ok := m["remains"].(string); ok {
		r.remains = v
	}

	return r
}