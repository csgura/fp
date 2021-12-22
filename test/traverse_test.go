package main_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func Swap[T any](a fp.Option[fp.Try[T]]) fp.Try[fp.Option[T]] {
	ret := iterator.FoldRight(a.Iterator(), try.Success(option.None[T]()), func(fb fp.Try[T], acc fp.Lazy[fp.Try[fp.Option[T]]]) fp.Try[fp.Option[T]] {
		panic("")
	})

	// ret := option.FoldRight(a, try.Success(option.None[T]()), func(fb fp.Try[T], acc fp.Lazy[fp.Try[fp.Option[T]]]) fp.Try[fp.Option[T]] {
	// 	panic("")
	// })
	return ret
}
func TestTraverse(t *testing.T) {

	optTry := option.Some(try.Success(10))
	Swap(optTry)

}
