//go:build ap
// +build ap

package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
)

func GetScheme() fp.Future[string] {
	return future.Successful("https")
}

func GetHost() fp.Future[string] {
	return future.Successful("localhost")
}

func GetPort() fp.Future[int] {
	return future.Successful(8080)
}

func MakeURLWithPort(scheme, addr string, port int) string {
	return fmt.Sprintf("%s://%s:%d", scheme, addr, port)
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

func TestApCircuitBreaking(t *testing.T) {
	res := future.Applicative3(MakeURLWithPort).
		ApFutureFunc(GetScheme).
		ApFutureFunc(GetHost).
		ApFutureFunc(GetPort)

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
		Ap("localhost").
		Ap(8080)

	fmt.Println(future.Await(res, time.Second))

	res = future.Chain3(MakeURLWithPort).
		ApOption(option.Some("https")).
		Ap("localhost").
		Map(func(scheme string) int {
			switch scheme {
			case "https":
				return 8443
			default:
				return 8080
			}
		})

	fmt.Println(future.Await(res, time.Second))

	// res = future.Applicative3(MakeURLWithPort).
	// 	ApOption(option.Some("https")).
	// 	Flip().
	// 	Map(func(scheme string) int {
	// 		switch scheme {
	// 		case "https":
	// 			return 8443
	// 		default:
	// 			return 8080
	// 		}
	// 	}).
	// 	HListMap(hlist.Rift2(func(scheme string, port int) string {
	// 		switch port {
	// 		case 8443:
	// 			return "localhost.uangel.com"
	// 		}
	// 		return "localhost"
	// 	}))

	// fmt.Println(future.Await(res, time.Second))

	// calcPort := func(scheme string) int {
	// 	switch scheme {
	// 	case "https":
	// 		return 8443
	// 	default:
	// 		return 8080
	// 	}
	// }

	// calcHost := func(scheme string, port int) string {
	// 	switch port {
	// 	case 8443:
	// 		return "localhost.uangel.com"
	// 	}
	// 	return "localhost"
	// }

	// res = future.Applicative3(MakeURLWithPort).
	// 	ApFuture(GetScheme()).
	// 	Flip().
	// 	Map(calcPort).
	// 	HListMap(hlist.Rift2(calcHost))

	// fmt.Println(future.Await(res, time.Second))

}
