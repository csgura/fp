package optiont

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func IsSomeCase[T comparable](v fp.OptionT[T]) fp.OptionT[T] {
	return try.IsSuccessCaseAnd(v, option.IsSomeCase)
}

func IsNoneCase[T comparable](v fp.OptionT[T]) fp.OptionT[T] {
	return try.IsSuccessCaseAnd(v, option.IsNoneCase)
}

func IsSomeCaseAnd[T comparable](v fp.OptionT[T], nested ...fp.Endo[T]) fp.OptionT[T] {
	return try.IsSuccessCaseAnd(v, option.NestedIsSomeCase(nested...))
}

func NestedIsSomeCase[T comparable](nested ...fp.Endo[T]) fp.Endo[fp.OptionT[T]] {
	return func(o fp.OptionT[T]) fp.OptionT[T] {
		return IsSomeCaseAnd(o, nested...)
	}
}
