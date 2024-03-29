package future_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/seq"
)

func TestFuture(t *testing.T) {

	p := promise.New[fp.Option[string]]()

	p.Success(option.Some("hello"))
	fp.Println(p.Future())

	s := seq.Of(1, 2, 3, 4)
	seqFuture := seq.Map(s, func(v int) fp.Future[int] {
		return future.Apply(func() int {
			time.Sleep(100 * time.Millisecond)
			return v * v
		})
	})

	futureSeq := future.Sequence(seqFuture)
	fp.Println(futureSeq)

	ch := make(chan fp.Try[[]int])

	futureSeq.OnComplete(func(result fp.Try[[]int]) {
		ch <- result
	})

	<-ch
	fp.Println(futureSeq)

}

func MakeURL(scheme, addr string) string {
	return fmt.Sprintf("%s://%s", scheme, addr)
}

func MakeURLWithPort(scheme, addr string, port int) string {
	return fmt.Sprintf("%s://%s:%d", scheme, addr, port)
}

func GetScheme() fp.Future[string] {
	return future.Successful("https")
}

func GetHost() fp.Future[string] {
	return future.Successful("localhost")
}

func GetPort() fp.Future[int] {
	return future.Successful(8080)
}

func TestApFlat(t *testing.T) {
	scheme := GetScheme()
	host := GetHost()

	res := future.FlatMap(scheme, func(s string) fp.Future[string] {
		return future.Map(host, func(h string) string {
			return MakeURL(s, h)
		})
	})

	fmt.Println(future.Await(res, time.Second))

}

func TestApFlat2(t *testing.T) {
	scheme := GetScheme()
	host := GetHost()
	port := GetPort()

	res := future.FlatMap(scheme, func(s string) fp.Future[string] {
		return future.FlatMap(host, func(h string) fp.Future[string] {
			return future.Map(port, func(p int) string {
				return MakeURLWithPort(s, h, p)
			})
		})
	})

	fmt.Println(future.Await(res, time.Second))

}

func TestCurrying(t *testing.T) {
	scheme := "http"
	host := "localhost"

	func2 := curried.Func2(MakeURL)
	func1 := func2(scheme)
	res := func1(host)

	fmt.Println(res)

}

func TestAp(t *testing.T) {
	scheme := GetScheme()
	host := GetHost()

	futureFunc2 := future.Successful(curried.Func2(MakeURL))
	futureFunc1 := future.Ap(futureFunc2, scheme)
	res := future.Ap(futureFunc1, host)

	fmt.Println(future.Await(res, time.Second))

}

func TestApFail(t *testing.T) {
	host := GetHost()

	futureFunc2 := future.Successful(curried.Func2(MakeURL))
	futureFunc1 := future.Ap(futureFunc2, future.Failed[string](fp.Error(500, "internal server error")))
	res := future.Ap(futureFunc1, host)

	fmt.Println(future.Await(res, time.Second))
}

func TestTraverse(t *testing.T) {
	ft := future.Traverse(iterator.Range(0, 10), func(v int) fp.Future[int] {
		return future.Successful(v)
	})

	res := future.Await(ft, time.Second)
	assert.True(res.IsSuccess())
	assert.Equal(len(res.Get().ToSeq()), 10)

	cnt := 0
	ft = future.Traverse(iterator.Range(0, 10), func(v int) fp.Future[int] {
		cnt++
		return future.Failed[int](errors.New("error"))
	})
	res = future.Await(ft, time.Second)

	assert.True(res.IsFailure())
	assert.Equal(cnt, 1)

}
