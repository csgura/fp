//go:build go1.27

package should

import (
	"testing"

	"github.com/csgura/fp"
)

type The struct {
	testing.TB
}

type TheTry[T any] struct {
	The
	t fp.Try[T]
}

func (r TheTry[T]) ShouldBeSuccess() T {
	r.Helper()
	return BeSuccess[T](r.TB, r.t)
}

type TheOption[T any] struct {
	The
	t fp.Option[T]
}

func (r TheOption[T]) ShouldBeSome() T {
	r.Helper()
	return BeSome[T](r.TB, r.t)
}
func (r The) Option[T any](v fp.Option[T]) TheOption[T] {
	return TheOption[T]{
		r, v,
	}
}

func (r The) Try[T any](v fp.Try[T]) TheTry[T] {
	return TheTry[T]{
		r, v,
	}
}

type TheOptionT[T any] struct {
	The
	t fp.OptionT[T]
}

func (r TheOptionT[T]) ShouldBeSome() T {
	r.Helper()
	return BeSome[T](r.TB, BeSuccess[fp.Option[T]](r.TB, r.t))
}

func (r The) OptionT[T any](v fp.OptionT[T]) TheOptionT[T] {
	return TheOptionT[T]{
		r, v,
	}
}

func Test(t testing.TB) The {
	return The{t}
}

type TheValue[T comparable] struct {
	The
	t T
}

func (r TheValue[T]) ShouldEqual(other T) {
	r.Helper()
	Equal[T](r.TB, r.t, other)
}

func (r The) Value[T comparable](v T) TheValue[T] {
	return TheValue[T]{
		r, v,
	}
}
