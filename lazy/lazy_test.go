package lazy_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/monoid"
)

//lint:file-ignore U1000 test code

func fibo(n int, prev int, curr int) int {
	if n == 0 {
		return prev
	}

	if n == 1 {
		return curr
	}

	return fibo(n-1, curr, curr+prev)
}

type tailFunc func() (int, tailFunc)

func fiboT(n int, prev int, curr int) (int, tailFunc) {
	if n == 0 {
		return prev, nil
	}

	if n == 1 {
		return curr, nil
	}

	return 0, func() (int, tailFunc) {
		return fiboT(n-1, curr, curr+prev)
	}
}

func fiboEval(n int, prev int, curr int) lazy.Eval[int] {
	if n == 0 {
		return lazy.Done(prev)
	}

	if n == 1 {
		return lazy.Done(curr)
	}

	return lazy.TailCall3(fiboEval, n-1, curr, curr+prev)
}

func fiboNoOpt(n int) lazy.Eval[int] {
	if n < 2 {
		return lazy.Done(n)
	}

	x := lazy.TailCall1(fiboNoOpt, (n - 1))

	y := lazy.TailCall1(fiboNoOpt, (n - 2))

	return lazy.Map2(x, y, monoid.Sum[int]().Combine)

}
func TestFibo(t *testing.T) {
	result := fibo(20, 0, 1)
	fmt.Println(result)

	lazyResult := fiboEval(20, 0, 1)
	fmt.Println(lazyResult.Get())

	assert.Equal(result, lazyResult.Get())
	fmt.Println(fiboNoOpt(20).Get())

}

func TestFiboMax(t *testing.T) {
	// max 1GB stack
	// 0.51s
	result := fibo(11184787, 0, 1)
	fmt.Println(result)

}

func TestLazyFiboMax(t *testing.T) {
	// 1.81s  -> 3~4 times slower
	lazyResult := fiboEval(11184787, 0, 1)
	fmt.Println(lazyResult.Get())
}

func TestLazyFiboMore(t *testing.T) {
	lazyResult := fiboEval(20000000, 0, 1)
	fmt.Println(lazyResult.Get())
}

type Result[S, A any] struct {
	s S
	a A
}

type State[S, A any] func(S) lazy.Eval[fp.Tuple2[A, S]]

func Pure[S, A any](v A) State[S, A] {
	return func(s S) lazy.Eval[fp.Tuple2[A, S]] {
		return lazy.Done(as.Tuple(v, s))
	}
}

func (r State[S, A]) FlatMap(f func(A) State[S, A]) State[S, A] {
	return func(s S) lazy.Eval[fp.Tuple2[A, S]] {
		res := r(s)
		return res.FlatMap(func(r fp.Tuple2[A, S]) lazy.Eval[fp.Tuple2[A, S]] {
			return f(r.I1)(r.I2)
		})
	}
}

func (r State[S, A]) Map(f func(A) A) State[S, A] {
	return r.FlatMap(func(a A) State[S, A] {
		return Pure[S](f(a))
	})
}

func runCountResume[T any](t lazy.Eval[T]) (T, int) {
	cnt := 0
	for {
		cnt++
		fmt.Printf("%d resume\n", cnt)
		result, continuation := lazy.Resume(t)

		if continuation != nil {
			t = continuation()
			continue
		}

		return result, cnt
	}
}

func TestLazyState(t *testing.T) {
	inc := func(v int) int {
		fmt.Printf("v = %d\n", v)
		return v + 1
	}
	i := Pure[context.Context](1).
		Map(inc). // 첫번째 resume 에서 flatMap 4회
		Map(inc). // flatMap 3회
		Map(inc). // flatMap 2회
		Map(inc)  // flatMap 1회
	res, resumeCnt := runCountResume(i(context.Background()))
	assert.Equal(res.I1, 5)
	assert.Equal(resumeCnt, 5)

}
