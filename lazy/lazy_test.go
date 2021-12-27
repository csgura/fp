package lazy_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/lazy"
)

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

	return lazy.TailCall(func() lazy.Eval[int] {
		return fiboEval(n-1, curr, curr+prev)
	})
}

func TestFibo(t *testing.T) {
	fmt.Println(fibo(20, 0, 1))

	fmt.Println(fiboEval(20, 0, 1).Get())
}
