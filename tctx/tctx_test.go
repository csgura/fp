package tctx_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/state"
	"github.com/csgura/fp/statet"
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

func fifthFunc(arg1 int, arg3 int, arg2 string) fp.Try[string] {
	return try.Success("hello")
}

func TestTCtx(t *testing.T) {
	r := parseString("12").Eval(context.Background())
	assert.Equal(r.Get(), 12)

	str := tctx.Pure("12")
	r = tctx.MapT2(str, try.Curried3(strconv.ParseInt), 10, 64).
		Eval(context.Background())
	assert.Equal(r.Get(), 12)

}

type something struct {
}

func (r something) Do1(ctx context.Context) string {
	return "10"
}

func (r something) DontCare1() string {
	return "10"
}

func (r something) DontCare2(arg int) string {
	return "10"
}

func (r something) Do2(ctx context.Context, arg string) string {
	return "10"
}

func TestStart(t *testing.T) {
	s1 := tctx.FromFunc3(firstFunc, "hello", 10)
	tctx.MapWithT1(s1, as.Curried3(secondFunc), 10)
	//tctx.FlatMapFunc2(s1, curried.Ap, 10)

	//s2 := tctx.Ap(tctx.Curried4(thirdFunc), 10)

	//s3 := tctx.Compose2(s1, s2)

	s1 = tctx.MapWithT1(s1, as.Curried3(firstFunc), 10)

	s1 = tctx.MapWithT2(s1, tctx.SlipL3(as.Curried4(thirdFunc)), 10, 20)

	s1 = tctx.MapWithT2(s1, tctx.SlipL4(as.Curried4(forthFunc)), 10, 20)
	s1 = tctx.MapT2(s1, curried.SlipL3(as.Curried3(fifthFunc)), 10, 20)

	p1 := tctx.Pure(something{})
	s1 = tctx.MapMethodWith(p1, something.Do1)
	s1 = tctx.MapMethodWith1(p1, something.Do2, "arg")

	s1 = tctx.Map(p1, something.DontCare1)
	s1 = tctx.MapT1(p1, try.CurriedPure2(something.DontCare2), 10)

}

func formatInt(ctx context.Context, fmtstr string, v int) string {
	return fmt.Sprintf(fmtstr, v)
}

func joinPort(ctx context.Context, v string, port int) string {
	return fmt.Sprintf("%s:%d", v, port)
}

func validate(ctx context.Context, v string) string {
	return v
}

func indep(v string) string {
	return v
}

func indeparg(v string, a1 int, a2 int) string {
	return v
}

func TestCompose(t *testing.T) {
	start := tctx.Pure(10)

	second := tctx.MapWithT1(start, tctx.SlipL3(try.CurriedPure3(formatInt)), "hello%d")
	tctx.MapWithT1(second, try.CurriedPure3(joinPort), 8080)

	cf1 := tctx.SlipL3(try.CurriedPure3(formatInt))

	ff := tctx.Compose5(
		tctx.AsWithFunc1(cf1, "hello%d"),
		tctx.AsWithFunc1(try.CurriedPure3(joinPort), 8080),
		try.Pure2(validate).Widen(),
		tctx.AsWithFunc(tctx.Const(try.Pure1(indep))),
		tctx.AsWithFunc2(tctx.Const(try.CurriedPure3(indeparg)), 10, 10),
	)
	tctx.MapWithT(start, ff)

}

var nextToken fp.State[string, string] = func(s string) fp.Tuple2[string, string] {
	s = strings.TrimSpace(s)
	for i, c := range s {
		if unicode.IsSpace(c) {
			return as.Tuple2(s[:i], s[i+1:])
		}
	}
	return as.Tuple(s, "")
}

func TestNextToken(t *testing.T) {
	token, s := nextToken.Run("hello world hi there")
	assert.Equal(token, "hello")

	token, s = nextToken.Run(s)
	assert.Equal(token, "world")

	token, s = nextToken.Run(s)
	assert.Equal(token, "hi")

	token = nextToken.Eval(s)
	assert.Equal(token, "there")

	tokens, s := state.FlatMap(nextToken, func(t1 string) fp.State[string, []string] {
		return state.Map(nextToken, func(t2 string) []string {
			return []string{t1, t2}
		})
	}).Run("hello world hi there")
	assert.Equal(tokens[0], "hello")
	assert.Equal(tokens[1], "world")
	assert.Equal(s, "hi there")

}

func TestParseInt(t *testing.T) {
	nextT := as.StateT(func(s string) fp.Try[fp.Tuple2[string, string]] {
		return try.Success(as.Tuple2(nextToken.Run(s)))
	})

	nextInt := statet.MapT(nextT, func(a string) fp.Try[int] {
		v, err := strconv.ParseInt(a, 10, 64)
		return try.Apply(int(v), err)
	})

	i, s := nextInt.Run("1 2 3 4 5")
	assert.Equal(i, try.Success(1))
	assert.Equal(s, try.Success("2 3 4 5"))

	i, s = nextInt.Run(s.Get())
	assert.Equal(i, try.Success(2))
	assert.Equal(s, try.Success("3 4 5"))

}

func push[T any](v T) fp.State[[]T, fp.Unit] {
	return state.Modify(func(s []T) []T {
		return append(s, v)
	})
}

func pop[T any](s []T) (fp.Option[T], []T) {
	l := len(s)
	if l == 0 {
		return option.None[T](), s
	}
	return option.Some(s[l-1]), s[:l-1]
}

func popState[T any]() fp.State[[]T, fp.Option[T]] {
	return func(t []T) fp.Tuple2[fp.Option[T], []T] {
		l := as.Seq(t).Last()
		return as.Tuple(l, seq.Init(t).Widen())
	}
}

func TestStack(t *testing.T) {
	stack := []int{}
	stack = push(10).Exec(stack)
	stack = push(20).Exec(stack)

	v, stack := popState[int]().Run(stack)
	assert.Equal(v, option.Some(20))
	v = popState[int]().Eval(stack)
	assert.Equal(v, option.Some(10))

}

func TestStack2(t *testing.T) {

	s := push(10)
	s = state.FlatMapConst(s, push(20))
	s2 := state.FlatMapConst(s, popState[int]())

	v := s2.Eval(nil)
	assert.Equal(v, option.Some(20))

}

func TestPop(t *testing.T) {

	s := []int{10, 20, 30}
	v1, s := pop(s)
	v2, s := pop(s)
	fmt.Printf("v1 = %s, v2 = %s\n", v1, v2)
	fmt.Printf("remain = %v\n", s)

	var doPop = state.New(pop[int])

	popTwice := state.FlatMap(doPop, /* 첫번째 pop 실행 */
		func(v1 fp.Option[int]) fp.State[[]int, fp.Option[fp.Tuple2[int, int]]] {
			return state.Map(doPop, /* 두번째 pop 실행 */
				func(v2 fp.Option[int]) fp.Option[fp.Tuple2[int, int]] {
					return option.Zip(v1, v2)
				})
		})
	res, remain := popTwice.Run([]int{10, 20, 30})
	fmt.Printf("res= %s, remain = %v\n", res, remain)

	popTwice2 := state.Zip(doPop, doPop)
	res2, remain := popTwice2.Run([]int{10, 20, 30})
	fmt.Printf("res= %s, remain = %v\n", res2, remain)
}
