package future_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/hlist"

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

	ch := make(chan fp.Try[fp.Seq[int]])

	futureSeq.OnComplete(func(result fp.Try[fp.Seq[int]]) {
		ch <- result
	})

	<-ch
	fp.Println(futureSeq)

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

func MakeURL(scheme, addr string) string {
	return fmt.Sprintf("%s://%s", scheme, addr)
}

func MakeURLWithPort(scheme, addr string, port int) string {
	return fmt.Sprintf("%s://%s:%d", scheme, addr, port)
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

func TestAp2(t *testing.T) {
	scheme := GetScheme()
	host := GetHost()
	port := GetPort()

	futureFunc3 := future.Successful(curried.Func3(MakeURLWithPort))
	futureFunc2 := future.Ap(futureFunc3, scheme)
	futureFunc1 := future.Ap(futureFunc2, host)

	res := future.Ap(futureFunc1, port)

	fmt.Println(future.Await(res, time.Second))

}

func TestApChain(t *testing.T) {

	res := future.Applicative3(MakeURLWithPort).
		ApFuture(GetScheme()).
		ApFuture(GetHost()).
		ApFuture(GetPort())

	fmt.Println(future.Await(res, time.Second))

	res = future.Applicative3(MakeURLWithPort).
		ApOption(option.Some("http")).
		ApTry(try.Success("localhost")).
		Ap(8080)

	fmt.Println(future.Await(res, time.Second))

	res = future.Applicative3(MakeURLWithPort).
		ApOption(option.Some("http")).
		Shift().
		Ap(8080).
		Ap("localhost")

	fmt.Println(future.Await(res, time.Second))

	res = future.Applicative3(MakeURLWithPort).
		ApOption(option.Some("https")).
		Shift().
		Map(func(scheme string) int {
			switch scheme {
			case "https":
				return 8443
			default:
				return 8080
			}
		}).
		Ap("localhost")
	fmt.Println(future.Await(res, time.Second))

	res = future.Applicative3(MakeURLWithPort).
		ApOption(option.Some("https")).
		Shift().
		Map(func(scheme string) int {
			switch scheme {
			case "https":
				return 8443
			default:
				return 8080
			}
		}).
		HListMap(hlist.Rift2( func( scheme string, port int ) string {
			switch port {
			case 8443:
				return "localhost.uangel.com"
			}
			return "localhost"
		}))
		
	fmt.Println(future.Await(res, time.Second))

	
	calcPort := func(scheme string) int {
		switch scheme {
		case "https":
			return 8443
		default:
			return 8080
		}
	}

	calcHost := func( scheme string, port int ) string {
		switch port {
		case 8443:
			return "localhost.uangel.com"
		}
		return "localhost"
	}

	res = future.Applicative3(MakeURLWithPort).
		ApFuture(GetScheme()).
		Shift().
		Map(calcPort).
		HListMap(hlist.Rift2(calcHost))
		
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
