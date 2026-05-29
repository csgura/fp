//go:build go1.27

package should

import (
	"testing"

	"github.com/csgura/fp"
)

type The struct {
	testing.TB
}

func (t The) BeSome[T any](v fp.Option[T]) T {
	t.Helper()
	return BeSome(t, v)
}

func (t The) BeSuccess[T any](v fp.Try[T]) T {
	t.Helper()
	return BeSuccess(t, v)
}

type TheOptionT[T any] struct {
	The
	t fp.OptionT[T]
}

func (r TheOptionT[T]) BeSome() T {
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

func (r TheValue[T]) Equal(other T) {
	r.Helper()
	Equal[T](r.TB, r.t, other)
}

func (r The) Value[T comparable](v T) TheValue[T] {
	return TheValue[T]{
		r, v,
	}
}
