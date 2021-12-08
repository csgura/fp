//go:generate go run github.com/csgura/fp/internal/generator/ord_gen
package ord

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
)

func Tuple1[A any](a fp.Ord[A]) fp.Ord[fp.Tuple1[A]] {
	return fp.LessFunc[fp.Tuple1[A]](func(t1 fp.Tuple1[A], t2 fp.Tuple1[A]) bool {
		if a.Less(t1.I1, t2.I1) {
			return true
		}
		return false
	})
}

func Option[T any](m fp.Ord[T]) fp.Ord[fp.Option[T]] {
	return fp.LessFunc[fp.Option[T]](func(t1 fp.Option[T], t2 fp.Option[T]) bool {
		return option.Applicative2(m.Less).ApOption(t1).ApOption(t2).OrElse(false)
	})
}
