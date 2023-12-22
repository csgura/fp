package try_test

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/tctx"
	"github.com/csgura/fp/try"
)

func print[T any](v T) {
	fmt.Println(v)
}

func TestTry(t *testing.T) {
	v := try.Success(10)

	assert.True(v.IsSuccess())
	assert.Equal(v.Get(), 10)
	assert.False(v.IsFailure())

	f2 := try.Success(try.Success(20))
	v = try.Flatten(f2)

	assert.True(v.IsSuccess())
	assert.Equal(v.Get(), 20)

	e := try.Failure[string](fmt.Errorf("bad request"))
	assert.True(e.IsFailure())
	assert.False(e.IsSuccess())
	assert.Equal(e.Failed().Get().Error(), "bad request")

	assert.Equal(e.Recover(func(err error) string {
		return "recover"
	}).Get(), "recover")

	e.RecoverWith(func(err error) fp.Try[string] {
		return try.Success("recoverWith")
	}).Foreach(print[string])

	assert.True(option.FromTry(v).IsDefined())

	// fp.Try[*url.URL]
	var u fp.Try[*url.URL] = try.Func1(url.Parse)("http://[abc")

	assert.True(u.IsFailure())

	var p fp.Try[string] = try.Map(u, (*url.URL).Port)
	assert.True(p.IsFailure())

	var intPort fp.Try[int] = try.Flatten(try.Map(p, try.Func1(strconv.Atoi)))
	fmt.Println(intPort)
	assert.True(intPort.IsFailure())

	try.FlatMap(p, try.Func1(strconv.Atoi)).Foreach(fp.Println[int])

}

func TestFlatMap(t *testing.T) {

	// fp.Try[*url.URL]
	var u fp.Try[*url.URL] = try.Func1(url.Parse)("http://localhost:8080/abcd")
	assert.True(u.IsSuccess())
	fmt.Println(u)

	var p fp.Try[string] = try.Map(u, (*url.URL).Port)
	assert.True(p.IsSuccess())
	assert.Equal(p.Get(), "8080")

	var intPort fp.Try[int] = try.FlatMap(p, try.Func1(strconv.Atoi))
	fmt.Println(intPort)
	assert.True(intPort.IsSuccess())
	assert.Equal(intPort.Get(), 8080)

}

func TestCompose(t *testing.T) {

	var intPort fp.Try[int] = try.Compose(
		try.Func1(url.Parse),
		fp.Compose((*url.URL).Port, try.Func1(strconv.Atoi)),
	)("http://localhost:8080/abcd")

	assert.True(intPort.IsSuccess())
	assert.Equal(intPort.Get(), 8080)
}

func ParsePort() (int, error) {
	u, err := url.Parse("http://localhost:8080/abcd")
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(u.Port())
}

func ParsePortFn() fp.Try[int] {
	return try.Compose(
		try.Func1(url.Parse),
		fp.Compose((*url.URL).Port, try.Func1(strconv.Atoi)),
	)("http://localhost:8080/abcd")
}

func TestSequence(t *testing.T) {

	successItr := iterator.Of(
		try.Success(10),
		try.Success(20),
		try.Success(30),
	)

	tryItr := try.SequenceIterator(successItr)
	assert.True(tryItr.IsSuccess())

	seq := tryItr.Get().ToSeq()
	assert.True(seq[0] == 10)
	assert.True(len(seq) == 3)

	failureItr := iterator.Of(
		try.Success(10),
		try.Failure[int](errors.New("hey")),
		try.Success(30),
	)

	tryItr = try.SequenceIterator(failureItr)
	assert.True(tryItr.IsFailure())

}

func TestTraverse(t *testing.T) {

	res := try.Traverse(iterator.Range(0, 10), func(v int) fp.Try[int] {
		return try.Success(v)
	})
	assert.True(res.IsSuccess())

	assert.Equal(len(res.Get().ToSeq()), 10)

	cnt := 0
	res = try.Traverse(iterator.Range(0, 10), func(v int) fp.Try[int] {
		cnt++
		return try.Failure[int](errors.New("error"))
	})

	assert.True(res.IsFailure())
	assert.Equal(cnt, 1)

}

func FunctorCompose(t *testing.T) {

	s := option.Of(1)

	ts := try.Success(s)

	tm := try.Map(ts, option.Lift(fp.Id[int]))
	assert.Equal(tm, ts)

	tm = try.Lift(option.Lift(fp.Id[int]))(ts)
	assert.Equal(tm, ts)

	tm = try.Map(ts, option.LiftM(func(v int) fp.Option[int] {
		return option.Some(v)
	}))

	assert.Equal(tm, ts)

}

func parseString(v string) tctx.State[int] {
	return tctx.Pure(10)
}

func TestState(t *testing.T) {
	//s, a := parseString("10")(context.Background())

	p := parseString("10")
	p = tctx.WithValue(p, "hello", "world")
	p = tctx.PeekContext(p, func(ctx context.Context) {
		fmt.Println("hello = ", ctx.Value("hello"))
	})
	a := tctx.Map(p, func(a int) string {
		return "hello"

	}).Eval(context.Background())
	fmt.Println(a)
}
