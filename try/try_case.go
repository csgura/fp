package try

import "github.com/csgura/fp"

func IsSuccessCase[T comparable](v fp.Try[T]) fp.Try[T] {
	return IsSuccessCaseAnd(v)
}

func IsFailureCase[T comparable](v fp.Try[T]) fp.Try[T] {
	return IsFailureCaseAnd(v)
}

func IsSuccessCaseAnd[T comparable](v fp.Try[T], nested ...fp.Endo[T]) fp.Try[T] {
	return NestedIsSuccessCase(nested...)(v)
}

func NestedIsSuccessCase[T comparable](nested ...fp.Endo[T]) fp.Endo[fp.Try[T]] {
	return func(o fp.Try[T]) fp.Try[T] {
		if o.IsSuccess() {
			return Map(o, fp.ComposeEndo(nested))
		}
		var zero T
		return Success(zero)
	}
}

func IsFailureCaseAnd[T comparable](v fp.Try[T], nested ...fp.Endo[error]) fp.Try[T] {
	return NestedIsFailureCase[T](nested...)(v)
}

func NestedIsFailureCase[T comparable](nested ...fp.Endo[error]) fp.Endo[fp.Try[T]] {
	return func(o fp.Try[T]) fp.Try[T] {
		if o.IsFailure() {
			return Flatten(Map(o.Failed(), fp.Compose(fp.ComposeEndo(nested), func(err error) fp.Try[T] {
				return Failure[T](err)
			})))
		}
		return Failure[T](fp.ErrTryNotFailed)
	}
}
