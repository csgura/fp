package fp

type Semigroup[T any] interface {
	Combine(a T, b T) T
}

type SemigroupFunc[T any] func(a T, b T) T

func (r SemigroupFunc[T]) Empty() T {
	var zero T
	return zero
}

func (r SemigroupFunc[T]) Combine(a T, b T) T {
	return r(a, b)
}

func (r SemigroupFunc[T]) AsMonoid(emptyFunc EmptyFunc[T]) Monoid[T] {
	return monoid[T]{
		emptyFunc, r,
	}
}

type Monoid[T any] interface {
	Semigroup[T]
	Empty() T
}

type EmptyFunc[T any] func() T

func (r EmptyFunc[T]) Empty() T {
	return r()
}

func Sum[T ImplicitOrd]() Monoid[T] {
	return SemigroupFunc[T](func(a, b T) T {
		return a + b
	})
}

type monoid[T any] struct {
	zero    EmptyFunc[T]
	combine SemigroupFunc[T]
}

func (r monoid[T]) Empty() T {
	return r.zero()
}

func (r monoid[T]) Combine(a, b T) T {
	return r.combine(a, b)
}

func Product[T ImplicitNum]() Monoid[T] {

	return SemigroupFunc[T](func(a, b T) T {
		return a * b
	}).AsMonoid(func() T {
		return 1
	})

}
