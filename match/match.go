package match

import "github.com/csgura/fp"

func Of[V, R any](v V, c ...fp.PartialFunc[V, R]) R {
	for _, m := range c {
		if m.IsDefined(v) {
			return m.Apply(v)
		}
	}
	panic("case not matched")
}
