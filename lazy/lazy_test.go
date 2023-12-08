package lazy_test

import (
	"fmt"
	"testing"

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
