package tctx_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/tctx"
	"github.com/csgura/fp/try"
)

func parseString(v string) tctx.State[int64] {
	return tctx.Of(func(ctx context.Context) fp.Try[int64] {
		return try.Apply(strconv.ParseInt(v, 10, 64))
	})
}

func TestTCtx(t *testing.T) {
	r := parseString("12").Eval(context.Background())
	assert.Equal(r.Get(), 12)

	str := tctx.Pure("12")
	r = tctx.MapNonContextLegacy3(str, strconv.ParseInt, 10, 64).
		Eval(context.Background())
	assert.Equal(r.Get(), 12)

}
