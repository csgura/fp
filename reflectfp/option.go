package reflectfp

import (
	"errors"
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

var zeroOption = reflect.TypeOf(option.None[fp.Unit]())
var optionTypeName = eraseTypeParam(zeroOption.Name())

func MatchOption(tpe reflect.Type) fp.Option[reflect.Type] {

	none := option.None[reflect.Type]()

	if tpe.PkgPath() != zeroOption.PkgPath() {
		return none
	}

	if optionTypeName != eraseTypeParam(tpe.Name()) {
		return none
	}

	getm, exists := tpe.MethodByName("Get")

	if !exists || getm.Type.NumIn() != 1 || getm.Type.NumOut() != 1 {
		return none
	}

	return option.Some(getm.Type.Out(0))
}

func None(optType reflect.Type) fp.Try[reflect.Value] {
	return try.Success(reflect.Zero(optType))

}

var ErrInvalidType = errors.New("invalid type")

func Some(optType reflect.Type, value reflect.Value) fp.Try[reflect.Value] {

	o := MatchOption(optType)
	if o.IsEmpty() {
		return try.Failure[reflect.Value](ErrInvalidType)
	}

	vtype := o.Get()

	return try.Of(func() reflect.Value {

		value = value.Convert(vtype)

		ret := reflect.Zero(optType)

		recoverm := ret.MethodByName("Recover")

		cbtype := reflect.FuncOf(nil, []reflect.Type{vtype}, false)
		cbf := reflect.MakeFunc(cbtype, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{value}
		})

		some := recoverm.Call([]reflect.Value{cbf})
		return some[0]
	})

}
