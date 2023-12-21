package match

import "github.com/csgura/fp"

func Of[V, R any](v V, c ...fp.PartialFunc[V, R]) R {
	for _, m := range c {
		if m.IsDefinedAt(v) {
			return m.Apply(v)
		}
	}
	panic("case not matched")
}

func Error[R any](head fp.PartialFunc[error, R], c ...fp.PartialFunc[error, R]) (func(error) bool, func(error) R) {
	ret := head

	for _, m := range c {
		ret = ret.OrElse(m)
	}
	return ret.Unapply()
}
