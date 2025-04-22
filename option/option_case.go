package option

import "github.com/csgura/fp"

func IsSomeCase[T comparable](v fp.Option[T]) fp.Option[T] {
	return IsSomeCaseAnd(v)
}

func IsNoneCase[T comparable](v fp.Option[T]) fp.Option[T] {
	if v.IsEmpty() {
		return v
	}

	return None[T]()
}

func IsSomeCaseAnd[T comparable](v fp.Option[T], nested ...fp.Endo[T]) fp.Option[T] {
	return NestedIsSomeCase(nested...)(v)
}

func NestedIsSomeCase[T comparable](nested ...fp.Endo[T]) fp.Endo[fp.Option[T]] {
	return func(o fp.Option[T]) fp.Option[T] {
		if o.IsDefined() {
			return Map(o, fp.ComposeEndo(nested))
		}
		var zero T
		return Some(zero)
	}
}
