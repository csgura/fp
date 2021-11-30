package try_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/try"
)

func print[T any](v T) {
	fmt.Println(v)
}

func TestTry(t *testing.T) {
	v := try.Success(10)

	v.Foreach(print[int])

	f2 := try.Success(try.Success(20))
	v = try.Flatten(f2)
	v.Foreach(print[int])

	e := try.Failure[string](fmt.Errorf("bad request"))
	e.Failed().Foreach(print[error])

	e.Recover(func(err error) string {
		return "recover"
	}).Foreach(print[string])

	e.RecoverWith(func(err error) fp.Try[string] {
		return try.Success("recoverWith")
	}).Foreach(print[string])

	v.ToOption().Foreach(print[int])
}
