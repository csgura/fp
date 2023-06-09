//go:generate go run github.com/csgura/fp/internal/generator/unit_gen
package unit

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
)

func Func0(f func()) fp.Func1[fp.Unit, fp.Unit] {
	return func(fp.Unit) fp.Unit {
		f()
		return fp.Unit{}
	}
}

var Success = fp.Success(fp.Unit{})
var Some = fp.Some(fp.Unit{})
var None = fp.None[fp.Unit]()
var Completed = future.Successful(fp.Unit{})
