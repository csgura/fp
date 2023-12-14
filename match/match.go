package match

import (
	"github.com/csgura/fp"
)

type PartialFunction[T, R any] func(T) fp.Option[R]

func (r PartialFunction[T, R]) Apply(v T) fp.Option[R] {
	return r(v)
}

func Of[V, R any](v V, c ...PartialFunction[V, R]) R {
	for _, m := range c {
		opt := m.Apply(v)
		if opt.IsDefined() {
			return opt.Get()
		}
	}
	panic("case not matched")
}
