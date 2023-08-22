package either_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/either"
	"github.com/csgura/fp/option"
)

func TestEither(t *testing.T) {
	l := either.Left[int, float64](10)
	either.Fold(l, option.Some, option.ConstNone).Foreach(fp.Println[int])

	s := either.Swap(l)
	s.Foreach(fp.Println[int])

}
