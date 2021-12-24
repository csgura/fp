package lazy

import "sync"

type Eval[T any] interface {
	Get() T
	Resume() (T, func() Eval[T])
	FlatMap(f func(T) Eval[T]) Eval[T]
	Map(f func(T) T) Eval[T]
}

// Run runs one Eval at a time, until there is no more continuation.
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

func Map[T any](t Eval[T], f func(T) T) Eval[T] {
	return t.FlatMap(func(value T) Eval[T] {
		return done[T]{Result: f(value)}
	})
}

// func ThenEval(t Eval, f func(interface{})) Eval {
// 	return t.Map(func(value interface{}) interface{} {
// 		f(value)
// 		return value
// 	})
// }

// done is a Eval, which has an executed result.

type done[T any] struct {
	Result T
}

func (d done[T]) Get() T {
	return d.Result
}

func (d done[T]) Resume() (T, func() Eval[T]) {
	return d.Result, nil
}

func (d done[T]) FlatMap(f func(T) Eval[T]) Eval[T] {
	return flatMap[T]{Subroutine: d, Continuation: f}
}

func (d done[T]) Map(f func(T) T) Eval[T] {
	return Map[T](d, f)
}

// func (d done) Then(f func(interface{})) Eval {
// 	return ThenEval(d, f)
// }

type Continuation[T any] interface {
	Continue() Eval[T]
}

// more is a Eval that returns a Eval as more work.

type more[T any] func() Eval[T]

func (d more[T]) Get() T {
	return Run[T](d)
}

func (m more[T]) Resume() (T, func() Eval[T]) {
	var zero T
	return zero, (func() Eval[T])(m)
}

func (m more[T]) FlatMap(f func(T) Eval[T]) Eval[T] {
	return flatMap[T]{Subroutine: m, Continuation: f}
}

func (m more[T]) Map(f func(T) T) Eval[T] {
	return Map[T](m, f)
}

func (m more[T]) Continue() Eval[T] {
	return m()
}

// FlatMap is a struct that contains the current computation and the continuation computation
type flatMap[T any] struct {
	Subroutine   Eval[T]
	Continuation func(T) Eval[T]
}

func (d flatMap[T]) Get() T {
	return Run[T](d)
}

func (m flatMap[T]) FlatMap(f func(T) Eval[T]) Eval[T] {

	continuation := m.Continuation
	return flatMap[T]{
		Subroutine: m.Subroutine,
		Continuation: func(value T) Eval[T] {
			return continuation(value).FlatMap(f)
		},
	}
}

func (m flatMap[T]) Resume() (T, func() Eval[T]) {
	continuation := m.Continuation

	switch sub := m.Subroutine.(type) {
	case done[T]:
		var zero T
		// if the subroutine is done, then the result is ready to be used as input for the continuation
		return zero, func() Eval[T] {
			return continuation(sub.Result)
		}
	case Continuation[T]:
		var zero T
		// if the subroutine is a continuation, then the result is not available yet, it has to call
		// sub.Continue() and use flatMap[T] to wait until the result is ready and be given the the
		// current continuation.
		return zero, func() Eval[T] {
			return sub.Continue().FlatMap(continuation)
		}
	case flatMap[T]:
		panic("flatMap[T] is not a valid subroutine. Use the flatMap[T] function to construct proper flatMap[T] structures.")
	}

	panic("")
}

func (m flatMap[T]) Map(f func(T) T) Eval[T] {
	return Map[T](m, f)
}

// func (m flatMap[T]) Then(f func(interface{})) Eval[T] {
// 	return ThenEval[T](m, f)
// }

func Value[T any](t T) Eval[T] {
	return done[T]{t}
}

func Defer[T any](f func() Eval[T]) Eval[T] {
	mf := Memoize(f)
	return more[T](mf)
}

func Call[T any](f func() T) Eval[T] {
	mf := Memoize(f)
	return more[T](func() Eval[T] {
		return Value(mf())
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
