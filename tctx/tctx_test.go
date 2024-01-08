package tctx_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/tctx"
	"github.com/csgura/fp/try"
)

func parseString(v string) tctx.State[int64] {
	return tctx.Of(func(ctx context.Context) fp.Try[int64] {
		return try.Apply(strconv.ParseInt(v, 10, 64))
	})
}

func firstFunc(ctx context.Context, arg1 string, arg2 int) fp.Try[string] {
	return try.Success("hello")
}

func secondFunc(ctx context.Context, arg1 string, arg2 int) fp.Try[string] {
	return try.Success("hello")
}

func thirdFunc(ctx context.Context, arg1 int, arg2 string, arg3 int) fp.Try[string] {
	return try.Success("hello")
}

func forthFunc(ctx context.Context, arg1 int, arg3 int, arg2 string) fp.Try[string] {
	return try.Success("hello")
}

func TestTCtx(t *testing.T) {
	r := parseString("12").Eval(context.Background())
	assert.Equal(r.Get(), 12)

	str := tctx.Pure("12")
	r = tctx.MapNonContextLegacy3(str, strconv.ParseInt, 10, 64).
		Eval(context.Background())
	assert.Equal(r.Get(), 12)

}

func TestStart(t *testing.T) {
	s1 := tctx.FromFunc3(firstFunc, "hello", 10)
	tctx.FlatMapFunc3(s1, secondFunc, 10)
	//tctx.FlatMapFunc2(s1, curried.Ap, 10)

	//s2 := tctx.Ap(tctx.Curried4(thirdFunc), 10)

	//s3 := tctx.Compose2(s1, s2)

	tctx.MapPureArg1(s1, as.Curried3(firstFunc), 10)

	tctx.MapPureArg2(s1, tctx.Fit3(as.Curried4(thirdFunc)), 10, 20)

	tctx.MapPureArg2(s1, tctx.Fit4(as.Curried4(forthFunc)), 10, 20)

}
