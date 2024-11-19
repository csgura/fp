//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package lazy

import (
	"sync"

	"github.com/csgura/fp/genfp"
)

// https://github.com/onflow/cadence/blob/v0.5.0-beta2/runtime/trampoline/trampoline.go

type Eval[T any] struct {
	// 첫번째 함수
	firstFunc func() T

	// 첫번째 함수의 리턴 결과로, 다음에 실행할 함수를 가져오는 함수
	getNextFunc func(T) Eval[T]
}

// Resume 의 두번째 결과가 함수 타입인 것은 nil 체크를 하기 위함
// fp 패키지를 참조하지 못하기 때문에  Either[T, EvalT]] 형태로 쓸 수 있는 것을
// (T, func() Eval[T]) 로 사용함
func (r Eval[T]) Resume() (T, func() Eval[T]) {

	firstFunc := r.firstFunc
	if firstFunc == nil {
		// 첫번째 함수가 없으면 zero 를 리턴하는 함수를 첫번째 함수로 함
		firstFunc = func() T {
			var zero T
			return zero
		}
	}

	// 다음에 호출할 함수가 없으면
	// 첫번째 함수의 결과를 리턴함
	if r.getNextFunc == nil {
		return firstFunc(), nil
	}

	var zero T

	// 두번째 결과가 Eval[T] 일 경우에는  첫번째 결과는 사용되지 않음.
	return zero, func() Eval[T] {
		// 다음에 호출할 함수가 있으면, 첫번째 함수의 결과를 이용하여
		// 다음에 호출할 함수를 구한 후에, 그 함수를 리턴함
		return r.getNextFunc(firstFunc())
	}

}

func (r Eval[T]) FlatMap(f func(T) Eval[T]) Eval[T] {

	// f는 다음에 실행될 함수
	// 만약 다음에 실행될 함수가 없으면  f 가 바로 getNextFunc가 됨
	if r.getNextFunc == nil {
		return Eval[T]{
			// 첫번째 실행 함수는 그대로임.
			firstFunc:   r.firstFunc,
			getNextFunc: f,
		}
	} else {
		// 만약 다음에 실행될 함수가 있었으면    r.getNextFunc -> f 순서로 실행되게 chain을 만듬
		// 일단 기존 실행 함수를 capture 해 두고
		getNextFunc := r.getNextFunc

		return Eval[T]{
			// 첫번째 실행 함수는 그대로임.
			firstFunc: r.firstFunc,
			getNextFunc: func(value T) Eval[T] {
				// capture 해둔 함수를 실행 후에,  FlatMap 으로 다음에 f가 실행되게 함.
				// getNextFunc의 리턴결과에 또 getNextFunc 가 있었으면
				// 안에서 FlatMap 이 또 호출되서,  FlatMap 호출 횟수가 어마무시하게 늘어날 수 있음.
				// flatMap 을 한번만 쓴 chain 길이가 n 이라면
				// 총 flatMap 호출 횟수는 sum(1...n) => n * (n+1) / 2
				return getNextFunc(value).FlatMap(f)
			},
		}
	}

	// 아무리 FlatMap 을 여려번 해도, 첫번째 실행 함수를 계속 유지하고 있기 때문에
	// stack overflow 없이 첫번째 함수 부터 실행 가능함.
}

func (r Eval[T]) Map(f func(T) T) Eval[T] {
	return r.FlatMap(func(value T) Eval[T] {
		return Done(f(value))
	})
}

func (r Eval[T]) Get() T {
	return Run(r)
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
	return t.Map(f)
}

func FlatMap[T any](t Eval[T], f func(T) Eval[T]) Eval[T] {
	return t.FlatMap(f)
}

// type done[T any] struct {
// 	Result T
// }

// func (d done[T]) sealed() {
// }

// func (d done[T]) Get() T {
// 	return d.Result
// }

// func (d done[T]) Resume() (T, func() Eval[T]) {
// 	return d.Result, nil
// }

// func (d done[T]) FlatMap(f func(T) Eval[T]) Eval[T] {
// 	return cont[T]{Subroutine: d, Continuation: f}
// }

// func (d done[T]) Map(f func(T) T) Eval[T] {
// 	return Map[T](d, f)
// }

// type call[T any] func() Eval[T]

// func (d call[T]) Get() T {
// 	return Run[T](d)
// }

// func (m call[T]) Resume() (T, func() Eval[T]) {
// 	var zero T
// 	return zero, (func() Eval[T])(m)
// }

// func (m call[T]) FlatMap(f func(T) Eval[T]) Eval[T] {
// 	return cont[T]{Subroutine: m, Continuation: f}
// }

// func (m call[T]) Map(f func(T) T) Eval[T] {
// 	return Map[T](m, f)
// }

// func (m call[T]) Continue() Eval[T] {
// 	return m()
// }

// func (d call[T]) sealed() {
// }

// type cont[T any] struct {
// 	Subroutine   Eval[T]
// 	Continuation func(T) Eval[T]
// }

// func (d cont[T]) sealed() {
// }

// func (d cont[T]) Get() T {
// 	return Run[T](d)
// }

// func (m cont[T]) FlatMap(f func(T) Eval[T]) Eval[T] {

// 	continuation := m.Continuation
// 	return cont[T]{
// 		Subroutine: m.Subroutine,
// 		Continuation: func(value T) Eval[T] {
// 			return continuation(value).FlatMap(f)
// 		},
// 	}
// }

// func (m cont[T]) Resume() (T, func() Eval[T]) {
// 	continuation := m.Continuation

// 	switch sub := m.Subroutine.(type) {
// 	case done[T]:
// 		var zero T
// 		return zero, func() Eval[T] {
// 			return continuation(sub.Result)
// 		}
// 	case call[T]:
// 		var zero T
// 		return zero, func() Eval[T] {
// 			return sub.Continue().FlatMap(continuation)
// 		}
// 	case cont[T]:
// 		panic("cont[T] is not a valid subroutine. Use the cont[T] function to construct proper cont[T] structures.")
// 	}

// 	panic("")
// }

// func (m cont[T]) Map(f func(T) T) Eval[T] {
// 	return Map[T](m, f)
// }

func Done[T any](t T) Eval[T] {
	return Eval[T]{
		firstFunc: func() T {
			return t
		},
	}
}

func TailCall[T any](f func() Eval[T]) Eval[T] {
	mf := Memoize(f)
	// return call[T](mf)

	return Eval[T]{
		firstFunc: func() T {
			var zero T
			return zero
		},
		getNextFunc: func(T) Eval[T] {
			return mf()
		},
	}
}

func Call[T any](f func() T) Eval[T] {
	mf := Memoize(f)
	// return call[T](func() Eval[T] {
	// 	return Done(mf())
	// })
	return Eval[T]{
		firstFunc: mf,
	}
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

func Func1[A, R any](f func(A) R) func(A) Eval[R] {
	return func(a A) Eval[R] {
		return Call(func() R {
			return f(a)
		})

	}
}

func Func2[A, B, R any](f func(A, B) R) func(A, B) Eval[R] {
	return func(a A, b B) Eval[R] {
		return Call(func() R {
			return f(a, b)
		})
	}
}

func Func3[A, B, C, R any](f func(A, B, C) R) func(A, B, C) Eval[R] {
	return func(a A, b B, c C) Eval[R] {
		return Call(func() R {
			return f(a, b, c)
		})
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File:    "tailcall_gen.go",
	Imports: []genfp.ImportPackage{},
	From:    1,
	Until:   genfp.MaxFunc,
	Template: `
func TailCall{{.N}}[{{TypeArgs 1 .N}}, R any]( f func({{TypeArgs 1 .N}}) Eval[R], {{DeclArgs 1 .N}} ) Eval[R] {
	return TailCall( func() Eval[R] {
		return f({{CallArgs 1 .N}})
	})
}
	`,
}
