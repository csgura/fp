package optiont_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/optiont"
	"github.com/csgura/fp/should"
	"github.com/csgura/fp/try"
)

func TestOptionT(t *testing.T) {
	v := optiont.Some(10)

	res := optiont.Map(v, func(a int) int {
		return a + 1
	})

	should.BeTrue(t, res.IsSuccess())

	res = try.Map(v, func(t fp.Option[int]) fp.Option[int] {
		return option.Map(t, func(a int) int {
			return a + 1
		})
	})
}
