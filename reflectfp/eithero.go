//go:build go1.27

package reflectfp

import (
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/optiont"
	"github.com/csgura/fp/try"
)

var eitherOTypeName = eraseTypeParam(TypeOf[fp.EitherO[fp.Unit, fp.Unit]]().Name())

type EitherOType struct {
	EitherType reflect.Type
	LeftType   reflect.Type
	RightType  reflect.Type
}

func MatchEitherO(tpe reflect.Type) fp.Option[EitherOType] {
	none := option.None[EitherOType]()

	if tpe.PkgPath() == eitherIntf.PkgPath() {
		tname := eraseTypeParam(tpe.Name())
		if eitherOTypeName == tname {

			unapplym, exists := tpe.MethodByName("Unapply")

			if !exists || unapplym.Type.NumIn() != 1 || unapplym.Type.NumOut() != 2 {
				return none
			}

			resetm, exists := tpe.MethodByName("Apply")

			if !exists || resetm.Type.NumIn() != 3 || resetm.Type.NumOut() != 1 {

				return none
			}
			return option.Some(EitherOType{
				tpe,
				unapplym.Type.Out(0).Elem(),
				unapplym.Type.Out(1).Elem(),
			})
		}
	}

	return none

}

// func Left(tpe reflect.Type, value reflect.Value) fp.Try[reflect.Value] {

// }

func GetLeft(eitherValue reflect.Value) fp.OptionT[reflect.Value] {
	return MatchEitherO(eitherValue.Type()).IntoTry(func() error {
		return fp.Error(404, "type %s is not either type", eitherValue.Type())
	}).FlatMap(func(et EitherOType) fp.OptionT[reflect.Value] {
		recoverm := eitherValue.MethodByName("Unapply")

		lp := recoverm.Call([]reflect.Value{})[0]

		if lp.Kind() != reflect.Pointer {
			return optiont.Failure[reflect.Value](fp.Error(404, "type %s is not either type", eitherValue.Type()))
		}

		if !lp.IsNil() {
			return optiont.Some(lp.Elem())
		}
		return optiont.None[reflect.Value]()

	})
}

func GetRight(eitherValue reflect.Value) fp.OptionT[reflect.Value] {
	return MatchEitherO(eitherValue.Type()).IntoTry(func() error {
		return fp.Error(404, "type %s is not either type", eitherValue.Type())
	}).FlatMap(func(et EitherOType) fp.OptionT[reflect.Value] {
		recoverm := eitherValue.MethodByName("Unapply")

		rp := recoverm.Call([]reflect.Value{})[1]

		if rp.Kind() != reflect.Pointer {
			return optiont.Failure[reflect.Value](fp.Error(404, "type %s is not either type", eitherValue.Type()))
		}

		if !rp.IsNil() {
			return optiont.Some(rp.Elem())
		}
		return optiont.None[reflect.Value]()

	})
}

func (r EitherOType) Left(value reflect.Value) fp.Try[reflect.Value] {

	vtype := r.LeftType

	return try.Of(func() reflect.Value {

		value = value.Convert(vtype)

		ret := reflect.Zero(r.EitherType)

		recoverm := ret.MethodByName("Apply")

		ptrv := func() reflect.Value {
			if !value.CanAddr() {
				ptr := reflect.New(vtype)
				ptr.Elem().Set(value)
				return ptr
			}
			return value.Addr()
		}()

		some := recoverm.Call([]reflect.Value{ptrv, reflect.Zero(reflect.PointerTo(r.RightType))})
		return some[0]
	})

}

func (r EitherOType) Right(value reflect.Value) fp.Try[reflect.Value] {

	vtype := r.RightType

	return try.Of(func() reflect.Value {

		value = value.Convert(vtype)

		ret := reflect.Zero(r.EitherType)

		recoverm := ret.MethodByName("Apply")

		ptrv := func() reflect.Value {
			if !value.CanAddr() {
				ptr := reflect.New(vtype)
				ptr.Elem().Set(value)
				return ptr
			}
			return value.Addr()
		}()

		some := recoverm.Call([]reflect.Value{reflect.Zero(reflect.PointerTo(r.LeftType)), ptrv})
		return some[0]
	})
}
