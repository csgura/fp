package reflectfp

import (
	"reflect"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func eraseTypeParam(v string) string {
	if sidx := strings.Index(v, "["); sidx > 0 {
		return v[0:sidx]
	}
	return v
}

var zeroLazyEval = reflect.TypeOf(lazy.Eval[fp.Unit]{})
var lazyEvalTypeName = eraseTypeParam(zeroLazyEval.Name())

func MatchLazyEval(tpe reflect.Type) fp.Option[reflect.Type] {

	none := option.None[reflect.Type]()

	if tpe.PkgPath() != zeroLazyEval.PkgPath() {
		return none
	}

	if lazyEvalTypeName != eraseTypeParam(tpe.Name()) {
		return none
	}

	getm, exists := tpe.MethodByName("Get")

	if !exists || getm.Type.NumIn() != 1 || getm.Type.NumOut() != 1 {
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
