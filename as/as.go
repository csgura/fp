//go:generate go run github.com/csgura/fp/internal/generator/as_gen
package as

import "github.com/csgura/fp"

func Func0[R any](f func() R) fp.Func0[R] {
	return fp.Func0[R](f)
}
