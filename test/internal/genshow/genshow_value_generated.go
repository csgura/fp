// Code generated by gombok, DO NOT EDIT.
package genshow

import (
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

func (r EmbeddedStruct) Hello() string {
	return r.hello
}

func (r EmbeddedStruct) World() struct {
	Level int
	Stage string
} {
	return r.world
}

func (r EmbeddedStruct) WithHello(v string) EmbeddedStruct {
	r.hello = v
	return r
}

func (r EmbeddedStruct) WithWorld(v struct {
	Level int
	Stage string
}) EmbeddedStruct {
	r.world = v
	return r
}

func (r EmbeddedStruct) String() string {
	return fmt.Sprintf("genshow.EmbeddedStruct{hello:%v, world:%v}", r.hello, r.world)
}

func (r EmbeddedStruct) AsTuple() fp.Tuple2[string, struct {
	Level int
	Stage string
}] {
	return as.Tuple2(r.hello, r.world)
}

func (r EmbeddedStruct) Unapply() (string, struct {
	Level int
	Stage string
}) {
	return r.hello, r.world
}

func (r EmbeddedStruct) AsMap() map[string]any {
	m := map[string]any{}
	m["hello"] = r.hello
	m["world"] = r.world
	return m
}

type EmbeddedStructBuilder EmbeddedStruct

func (r EmbeddedStructBuilder) Build() EmbeddedStruct {
	return EmbeddedStruct(r)
}

func (r EmbeddedStruct) Builder() EmbeddedStructBuilder {
	return EmbeddedStructBuilder(r)
}

func (r EmbeddedStructBuilder) Hello(v string) EmbeddedStructBuilder {
	r.hello = v
	return r
}

func (r EmbeddedStructBuilder) World(v struct {
	Level int
	Stage string
}) EmbeddedStructBuilder {
	r.world = v
	return r
}

func (r EmbeddedStructBuilder) FromTuple(t fp.Tuple2[string, struct {
	Level int
	Stage string
}]) EmbeddedStructBuilder {
	r.hello = t.I1
	r.world = t.I2
	return r
}

func (r EmbeddedStructBuilder) Apply(hello string, world struct {
	Level int
	Stage string
}) EmbeddedStructBuilder {
	r.hello = hello
	r.world = world
	return r
}

func (r EmbeddedStructBuilder) FromMap(m map[string]any) EmbeddedStructBuilder {

	if v, ok := m["hello"].(string); ok {
		r.hello = v
	}

	if v, ok := m["world"].(struct {
		Level int
		Stage string
	}); ok {
		r.world = v
	}

	return r
}

type EmbeddedStructMutable struct {
	Hello string
	World struct {
		Level int
		Stage string
	}
}

func (r EmbeddedStruct) AsMutable() EmbeddedStructMutable {
	return EmbeddedStructMutable{
		Hello: r.hello,
		World: r.world,
	}
}

func (r EmbeddedStructMutable) AsImmutable() EmbeddedStruct {
	return EmbeddedStruct{
		hello: r.Hello,
		world: r.World,
	}
}

func (r EmbeddedTypeParamStruct[T]) Hello() string {
	return r.hello
}

func (r EmbeddedTypeParamStruct[T]) World() struct {
	Level T
	Stage string
} {
	return r.world
}

func (r EmbeddedTypeParamStruct[T]) WithHello(v string) EmbeddedTypeParamStruct[T] {
	r.hello = v
	return r
}

func (r EmbeddedTypeParamStruct[T]) WithWorld(v struct {
	Level T
	Stage string
}) EmbeddedTypeParamStruct[T] {
	r.world = v
	return r
}

func (r EmbeddedTypeParamStruct[T]) String() string {
	return fmt.Sprintf("genshow.EmbeddedTypeParamStruct{hello:%v, world:%v}", r.hello, r.world)
}

func (r EmbeddedTypeParamStruct[T]) AsTuple() fp.Tuple2[string, struct {
	Level T
	Stage string
}] {
	return as.Tuple2(r.hello, r.world)
}

func (r EmbeddedTypeParamStruct[T]) Unapply() (string, struct {
	Level T
	Stage string
}) {
	return r.hello, r.world
}

func (r EmbeddedTypeParamStruct[T]) AsMap() map[string]any {
	m := map[string]any{}
	m["hello"] = r.hello
	m["world"] = r.world
	return m
}

type EmbeddedTypeParamStructBuilder[T any] EmbeddedTypeParamStruct[T]

func (r EmbeddedTypeParamStructBuilder[T]) Build() EmbeddedTypeParamStruct[T] {
	return EmbeddedTypeParamStruct[T](r)
}

func (r EmbeddedTypeParamStruct[T]) Builder() EmbeddedTypeParamStructBuilder[T] {
	return EmbeddedTypeParamStructBuilder[T](r)
}

func (r EmbeddedTypeParamStructBuilder[T]) Hello(v string) EmbeddedTypeParamStructBuilder[T] {
	r.hello = v
	return r
}

func (r EmbeddedTypeParamStructBuilder[T]) World(v struct {
	Level T
	Stage string
}) EmbeddedTypeParamStructBuilder[T] {
	r.world = v
	return r
}

func (r EmbeddedTypeParamStructBuilder[T]) FromTuple(t fp.Tuple2[string, struct {
	Level T
	Stage string
}]) EmbeddedTypeParamStructBuilder[T] {
	r.hello = t.I1
	r.world = t.I2
	return r
}

func (r EmbeddedTypeParamStructBuilder[T]) Apply(hello string, world struct {
	Level T
	Stage string
}) EmbeddedTypeParamStructBuilder[T] {
	r.hello = hello
	r.world = world
	return r
}

func (r EmbeddedTypeParamStructBuilder[T]) FromMap(m map[string]any) EmbeddedTypeParamStructBuilder[T] {

	if v, ok := m["hello"].(string); ok {
		r.hello = v
	}

	if v, ok := m["world"].(struct {
		Level T
		Stage string
	}); ok {
		r.world = v
	}

	return r
}

type EmbeddedTypeParamStructMutable[T any] struct {
	Hello string
	World struct {
		Level T
		Stage string
	}
}

func (r EmbeddedTypeParamStruct[T]) AsMutable() EmbeddedTypeParamStructMutable[T] {
	return EmbeddedTypeParamStructMutable[T]{
		Hello: r.hello,
		World: r.world,
	}
}

func (r EmbeddedTypeParamStructMutable[T]) AsImmutable() EmbeddedTypeParamStruct[T] {
	return EmbeddedTypeParamStruct[T]{
		hello: r.Hello,
		world: r.World,
	}
}
