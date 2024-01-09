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
	"github.com/csgura/fp/tctx"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/tstate"
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

type State[S, A any] func(S) (S, A)

func (r State[S, A]) Run(s S) (S, A) {
	return r(s)
}

var nextToken State[string, string] = func(s string) (string, string) {
	s = strings.TrimSpace(s)
	for i, c := range s {
		if unicode.IsSpace(c) {
			return s[i+1:], s[:i]
		}
	}
	return "", s
}

func TestNextToken(t *testing.T) {
	s, token := nextToken.Run("hello world hi there")
	assert.Equal(token, "hello")

	s, token = nextToken.Run(s)
	assert.Equal(token, "world")

	s, token = nextToken.Run(s)
	assert.Equal(token, "hi")

	_, token = nextToken.Run(s)
	assert.Equal(token, "there")

	s, tokens := flatMap(nextToken, func(t1 string) State[string, []string] {
		return Map(nextToken, func(t2 string) []string {
			return []string{t1, t2}
		})
	}).Run("hello world hi there")
	assert.Equal(tokens[0], "hello")
	assert.Equal(tokens[1], "world")
	assert.Equal(s, "hi there")

}

func TestParseInt(t *testing.T) {
	nextT := tstate.Of(func(s string) fp.Try[fp.Tuple2[string, string]] {
		return try.Success(as.Tuple2(nextToken.Run(s)))
	})

	nextInt := tstate.MapT(nextT, func(a string) fp.Try[int] {
		v, err := strconv.ParseInt(a, 10, 64)
		return try.Apply(int(v), err)
	})

	s, i := nextInt.Run("1 2 3 4 5")
	assert.Equal(i, try.Success(1))
	assert.Equal(s, try.Success("2 3 4 5"))

	s, i = nextInt.Run(s.Get())
	assert.Equal(i, try.Success(2))
	assert.Equal(s, try.Success("3 4 5"))

}

func push[T any](v T) State[[]T, T] {
	return func(t []T) ([]T, T) {
		return append(t, v), v
	}
}

func pop[T any]() State[[]T, fp.Option[T]] {
	return func(t []T) ([]T, fp.Option[T]) {
		l := as.Seq(t).Last()
		return as.Seq(t).Init(), l
	}
}

func TestStack(t *testing.T) {
	stack := []int{}
	stack, _ = push(10).Run(stack)
	stack, _ = push(20).Run(stack)

	stack, v := pop[int]().Run(stack)
	assert.Equal(v, option.Some(20))
	_, v = pop[int]().Run(stack)
	assert.Equal(v, option.Some(10))

}

func Map[S, A, B any](st State[S, A], f func(A) B) State[S, B] {
	return func(s S) (S, B) {
		s, a := st.Run(s)
		return s, f(a)
	}
}

func flatMap[S, A, B any](st State[S, A], f func(A) State[S, B]) State[S, B] {
	return func(s S) (S, B) {
		s, a := st.Run(s)
		return f(a).Run(s)
	}
}

func flatMap2[S, A, B any](st State[S, A], f State[S, B]) State[S, B] {
	return flatMap(st, fp.Const[A](f))
}

func TestStack2(t *testing.T) {

	s := push(10)
	s = flatMap2(s, push(20))
	s2 := flatMap2(s, pop[int]())

	_, v := s2.Run(nil)
	assert.Equal(v, option.Some(20))
}
