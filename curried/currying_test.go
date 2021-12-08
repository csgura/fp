package curried_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/curried"
)

func Concat(prefix string, suffix string) string {
	return fmt.Sprintf("%s-%s", prefix, suffix)
}

func Map(s string, fn func(string) string) string {
	return fn(s)
}

func TestCurried(t *testing.T) {

	prefix := "hello"
	Map("world", func(s string) string {
		return Concat(prefix, s)
	})

	Map("world", curried.Func2(Concat)(prefix))

	fn := curried.Func3(func(a string, b int, c string) string {
		return fmt.Sprint(a, b, c)
	})("hello")

	println(curried.Revert2(fn)(10, "world"))
}
