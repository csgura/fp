package reflectfp

import (
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/either"
	"github.com/csgura/fp/option"
)

var eitherIntf = TypeOf[fp.Either[fp.Unit, fp.Unit]]()

var zeroLeft = reflect.TypeOf(either.Left[fp.Unit, fp.Unit](fp.Unit{}))
var zeroRight = reflect.TypeOf(either.Right[fp.Unit, fp.Unit](fp.Unit{}))

var leftTypeName = eraseTypeParam(zeroLeft.Name())
var rightTypeName = eraseTypeParam(zeroRight.Name())
var eitherTypeName = eraseTypeParam(eitherIntf.Name())

type EitherType struct {
	Left  reflect.Type
	Right reflect.Type
}

func MatchEither(tpe reflect.Type) fp.Option[EitherType] {
	none := option.None[EitherType]()

	if tpe.PkgPath() == zeroLeft.PkgPath() {
		tname := eraseTypeParam(tpe.Name())
		if leftTypeName == tname || rightTypeName == tname {

			getm, exists := tpe.MethodByName("Get")

			if !exists || getm.Type.NumIn() != 1 || getm.Type.NumOut() != 1 {
				return none
			}

			leftm, exists := tpe.MethodByName("Left")

			if !exists || getm.Type.NumIn() != 1 || getm.Type.NumOut() != 1 {
				return none
			}
			return option.Some(EitherType{leftm.Type.Out(0), getm.Type.Out(0)})

		}
	}

	if tpe.PkgPath() == eitherIntf.PkgPath() {
		tname := eraseTypeParam(tpe.Name())

		if eitherTypeName == tname {
			getm, exists := tpe.MethodByName("Get")

			if !exists || getm.Type.NumIn() != 0 || getm.Type.NumOut() != 1 {
				return none
			}

			leftm, exists := tpe.MethodByName("Left")

			if !exists || getm.Type.NumIn() != 0 || getm.Type.NumOut() != 1 {
				return none
			}
			return option.Some(EitherType{leftm.Type.Out(0), getm.Type.Out(0)})
		}
	}

	return none

}

// left[L,R]  로 right[L,R] 을 만들 수 있고, 반대도 가능한데
// fp.Either[L,R] 타입을 이용해서   left[L,R] 이나 right[L,R] 을 만드는 건 불가능
