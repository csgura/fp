package optiont_test

import (
	"testing"

	"github.com/csgura/fp/optiont"
	"github.com/csgura/fp/should"
)

func TestOptionT(t *testing.T) {
	v := optiont.Some(10)

	res := optiont.Map(v, func(a int) int {
		return a + 1
	})

	should.BeTrue(t, res.IsSuccess())
}
