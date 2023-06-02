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

func (r SemigroupFunc[T]) Curried() func(T) func(T) T {
	return func(a1 T) func(T) T {
		return func(a2 T) T {
			return r.Combine(a1, a2)
		}
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

func (r monoid[T]) ToMonoid(emptyFunc EmptyFunc[T]) Monoid[T] {
	return monoid[T]{emptyFunc, r.combine}
}

func (r monoid[T]) Curried() func(T) func(T) T {
	return r.combine.Curried()
}

func Product[T ImplicitNum]() Monoid[T] {
	return monoid[T]{
		zero: func() T {
			return 1
		},
		combine: SemigroupFunc[T](func(a, b T) T {
			return a * b
		}),
	}
}

type Endo[T any] func(T) T

func (r Endo[T]) AsFunc() func(T) T {
	return r
}

// semigroup 이나 monoid 의 combine 순서를 반대로 하고 싶을 때는 Dual 로 감싸면 된다.
type Dual[T any] struct {
	GetDual T
}
