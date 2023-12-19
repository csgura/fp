package match

import "github.com/csgura/fp"

func Of[V, R any](v V, c ...fp.PartialFunc[V, R]) R {
	for _, m := range c {
		opt := m(v)
		if opt.IsDefined() {
			return opt.Get()
		}
	}
	panic("case not matched")
}
