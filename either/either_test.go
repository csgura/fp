package either_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/either"
	"github.com/csgura/fp/option"
)

func TestEither(t *testing.T) {
	l := either.NotRight[float64](10)

	either.Fold(l, option.Some, option.ConstNone).Foreach(fp.Println[int])

	s := either.Swap(l)
	either.Foreach(s, fp.Println[int])

	l = either.Right[int](10.2)
	either.Foreach(l, fp.Println)

}
