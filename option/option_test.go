package option_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/option"
)

func Sum(a string, b int) string {
	return fmt.Sprintf("%s:%d", a, b)
}

func TestSome(t *testing.T) {
	opt := option.Some("10")
	opt.Foreach(func(v string) {
		println(v)
	})

	opt2 := option.FlatMap(opt, func(v string) fp.Option[int] {
		return option.Some(20)
	})

	opt2.Foreach(fp.Println[int])

	ap := option.Applicative2(Sum)
	result := ap.Ap("hello").Ap(20)

	result.Foreach(fp.Println[string])

	result = ap.Ap("hello").ApOption(option.None[int]())

	fmt.Printf("result is defined = %t\n", result.IsDefined())

	result = ap.ApOption(option.None[string]()).Ap(10)

	fmt.Printf("result is defined = %t\n", result.IsDefined())

	ap = option.Applicative2(Sum)
	r2 := ap.Ap("ap flatmap").FlatMap(func(h string) fp.Option[int] {
		fmt.Printf("v = %v", h)
		return option.Some(10)
	})

	r2.Foreach(func(v string) {
		println(v)
	})

	r2 = ap.ApOption(option.None[string]()).FlatMap(func(h string) fp.Option[int] {
		return option.Some(10)
	})

	r2.Foreach(fp.Println[string])
	fmt.Printf("result is defined = %t\n", r2.IsDefined())

	option.Applicative3(func(addr string, port int, scheme string) string {
		return fmt.Sprintf("connect to %s://%s:%d", scheme, addr, port)
	}).
		Ap("hello.world.com").
		FlatMap(func(addr string) fp.Option[int] {
			return option.Some(80)
		}).
		HListMap(hlist.Ap2(func(addr string, port int) string {
			if port == 80 {
				return "http"
			}
			return "https"
		})).
		Foreach(fp.Println[string])

	// hl := hlist.Concact(true, hlist.Concact("hello", hlist.Concact(10, hlist.Empty())))
	// hlist.Case2(hl, func(a bool, b string) string {
	// 	fmt.Printf("a = %v , b = %v\n", a, b)
	// 	return "good"
	// })

}
