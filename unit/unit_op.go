//go:generate go run github.com/csgura/fp/internal/generator/unit_gen
package unit

import "github.com/csgura/fp"

func Func0(f func()) fp.Func1[fp.Unit, fp.Unit] {
	return func(fp.Unit) fp.Unit {
		f()
		return fp.Unit{}
	}
}
