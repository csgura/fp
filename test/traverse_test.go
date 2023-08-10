package main_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func Swap[T any](a fp.Option[fp.Try[T]]) fp.Try[fp.Option[T]] {
	if a.IsEmpty() {
		return try.Success(option.None[T]())
	}
	return try.Map(a.Get(), option.Some)
}

func TestTraverse(t *testing.T) {

	optTry := option.Some(try.Success(10))
	Swap(optTry)

}
