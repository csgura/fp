//go:generate go run github.com/csgura/fp/internal/generator/lazy_gen
package lazy

import (
	"sync"
)

// https://github.com/onflow/cadence/blob/v0.5.0-beta2/runtime/trampoline/trampoline.go

type Eval[T any] interface {
	sealed()
	Get() T
	Resume() (T, func() Eval[T])
	FlatMap(f func(T) Eval[T]) Eval[T]
	Map(f func(T) T) Eval[T]
}

func Run[T any](t Eval[T]) T {
	for {
		result, continuation := t.Resume()

		if continuation != nil {
			t = continuation()
			continue
		}

		return result
	}
}

func Map2[T any](a, b Eval[T], f func(T, T) T) Eval[T] {
	return a.FlatMap(func(v1 T) Eval[T] {
		return b.Map(func(v2 T) T {
			return f(v1, v2)
		})
	})
}

func Map[T any](t Eval[T], f func(T) T) Eval[T] {
	return t.FlatMap(func(value T) Eval[T] {
		return done[T]{Result: f(value)}
	})
}

func FlatMap[T any](t Eval[T], f func(T) Eval[T]) Eval[T] {
	return t.FlatMap(f)
}

type done[T any] struct {
	Result T
}

func (d done[T]) sealed() {
}

func (d done[T]) Get() T {
	return d.Result
}

func (d done[T]) Resume() (T, func() Eval[T]) {
	return d.Result, nil
}

func (d done[T]) FlatMap(f func(T) Eval[T]) Eval[T] {
	return cont[T]{Subroutine: d, Continuation: f}
}

func (d done[T]) Map(f func(T) T) Eval[T] {
	return Map[T](d, f)
}

type call[T any] func() Eval[T]

func (d call[T]) Get() T {
	return Run[T](d)
}

func (m call[T]) Resume() (T, func() Eval[T]) {
	var zero T
	return zero, (func() Eval[T])(m)
}

func (m call[T]) FlatMap(f func(T) Eval[T]) Eval[T] {
	return cont[T]{Subroutine: m, Continuation: f}
}

func (m call[T]) Map(f func(T) T) Eval[T] {
	return Map[T](m, f)
}

func (m call[T]) Continue() Eval[T] {
	return m()
}

func (d call[T]) sealed() {
}

type cont[T any] struct {
	Subroutine   Eval[T]
	Continuation func(T) Eval[T]
}

func (d cont[T]) sealed() {
}

func (d cont[T]) Get() T {
	return Run[T](d)
}

func (m cont[T]) FlatMap(f func(T) Eval[T]) Eval[T] {

	continuation := m.Continuation
	return cont[T]{
		Subroutine: m.Subroutine,
		Continuation: func(value T) Eval[T] {
			return continuation(value).FlatMap(f)
		},
	}
}

func (m cont[T]) Resume() (T, func() Eval[T]) {
	continuation := m.Continuation

	switch sub := m.Subroutine.(type) {
	case done[T]:
		var zero T
		return zero, func() Eval[T] {
			return continuation(sub.Result)
		}
	case call[T]:
		var zero T
		return zero, func() Eval[T] {
			return sub.Continue().FlatMap(continuation)
		}
	case cont[T]:
		panic("cont[T] is not a valid subroutine. Use the cont[T] function to construct proper cont[T] structures.")
	}

	panic("")
}

func (m cont[T]) Map(f func(T) T) Eval[T] {
	return Map[T](m, f)
}

func Done[T any](t T) Eval[T] {
	return done[T]{t}
}

func TailCall[T any](f func() Eval[T]) Eval[T] {
	mf := Memoize(f)
	return call[T](mf)
}

func Call[T any](f func() T) Eval[T] {
	mf := Memoize(f)
	return call[T](func() Eval[T] {
		return Done(mf())
	})
}

func Memoize[T any](f func() T) func() T {
	once := sync.Once{}
	var ret T
	return func() T {
		once.Do(func() {
			ret = f()
		})
		return ret
	}
}

func Func1[A, R any](f func(A) R) func(A) func() R {
	return func(a A) func() R {
		return func() R {
			return f(a)
		}
	}
}

func Func2[A, B, R any](f func(A, B) R) func(A, B) func() R {
	return func(a A, b B) func() R {
		return func() R {
			return f(a, b)
		}
	}
}

func Func3[A, B, C, R any](f func(A, B, C) R) func(A, B, C) func() R {
	return func(a A, b B, c C) func() R {
		return func() R {
			return f(a, b, c)
		}
	}
}
