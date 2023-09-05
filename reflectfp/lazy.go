package reflectfp

import (
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func MatchLazyEval(tpe reflect.Type) fp.Option[reflect.Type] {
	none := option.None[reflect.Type]()

	getm, exists := tpe.MethodByName("Get")

	if !exists || getm.Type.NumIn() != 1 || getm.Type.NumOut() != 1 {
		return none
	}

	mapm, exists := tpe.MethodByName("Map")
	if !exists || mapm.Type.NumIn() != 2 || mapm.Type.NumOut() != 1 {
		return none
	}

	cbtype := mapm.Type.In(1)

	if cbtype.Kind() != reflect.Func {
		return none
	}

	if cbtype.NumIn() != 1 || cbtype.NumOut() != 1 {
		return none
	}

	if cbtype.In(0) != getm.Type.Out(0) {
		return none
	}

	return option.Some(getm.Type.Out(0))
}

func LazyCall(lazyType reflect.Type, f func() reflect.Value) fp.Try[reflect.Value] {
	o := MatchLazyEval(lazyType)
	if o.IsEmpty() {
		return try.Failure[reflect.Value](ErrInvalidType)
	}

	vtype := o.Get()

	return try.Of(func() reflect.Value {
		//value = value.Convert(vtype)

		ret := reflect.Zero(lazyType)

		mapm := ret.MethodByName("Map")

		cbtype := reflect.FuncOf([]reflect.Type{vtype}, []reflect.Type{vtype}, false)
		cbf := reflect.MakeFunc(cbtype, func(args []reflect.Value) (results []reflect.Value) {
			value := f().Convert(vtype)
			return []reflect.Value{value}
		})

		lazyeval := mapm.Call([]reflect.Value{cbf})
		return lazyeval[0]
	})
}
