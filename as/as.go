//go:generate go run github.com/csgura/fp/internal/generator/as_gen
package as

import "github.com/csgura/fp"

func Func0[R any](f func() R) fp.Func1[fp.Unit, R] {
	return func(u fp.Unit) R {
		return f()
	}
}

func Seq[T any](s []T) fp.Seq[T] {
	return fp.Seq[T](s)
}

func Ptr[T any](v T) *T {
	return &v
}

func Interface[T, I any](v T) I {
	var a any = v
	return a.(I)
}

func InstanceOf[T any](v any) T {
	return v.(T)
}

func Tuple[K, V any](k K, v V) fp.Tuple2[K, V] {
	return fp.Tuple2[K, V]{k, v}
}

func Dual[T any](t T) fp.Dual[T] {
	return fp.Dual[T]{GetDual: t}
}

func Endo[T any](f fp.Func1[T, T]) fp.Endo[T] {
	return fp.Endo[T](f)
}
