package reflectfp

import (
	"errors"
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/try"
)

func IsOptionType(tpe reflect.Type) bool {
	m, exists := tpe.MethodByName("Get")

	if !exists || m.Type.NumIn() != 1 || m.Type.NumOut() != 1 {
		return false
	}

	m, exists = tpe.MethodByName("Recover")
	if !exists || m.Type.NumIn() != 2 || m.Type.NumOut() != 1 {
		return false
	}

	return exists
}

func None(optType reflect.Type) fp.Try[reflect.Value] {
	return try.Success(reflect.Zero(optType))

}

var ErrInvalidType = errors.New("invalid type")

func Some(optType reflect.Type, value reflect.Value) fp.Try[reflect.Value] {
	ret := reflect.Zero(optType)
	get, exists := optType.MethodByName("Get")
	if !exists {
		return try.Failure[reflect.Value](ErrInvalidType)
	}

	getm := ret.Method(get.Index)

	recover, exists := optType.MethodByName("Recover")

	if !exists {
		return try.Failure[reflect.Value](ErrInvalidType)
	}

	m := ret.Method(recover.Index)
	cbtype := reflect.FuncOf(nil, []reflect.Type{getm.Type().Out(0)}, false)
	cbf := reflect.MakeFunc(cbtype, func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{value}
	})

	some := m.Call([]reflect.Value{cbf})
	return try.Success(some[0])

}
