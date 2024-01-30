package reflectfp

import "reflect"

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(&t).Elem()
}
