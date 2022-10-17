//go:build ap
// +build ap

package main_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
)

func Sum(a string, b int) string {
	return fmt.Sprintf("%s:%d", a, b)
}

func TestSome(t *testing.T) {

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
		HListMap(hlist.Rift2(func(addr string, port int) string {
			if port == 80 {
				return "http"
			}
			return "https"
		})).
		Foreach(fp.Println[string])

}

func TestFunc(t *testing.T) {
	o1 := option.Some(1)
	o2 := option.Some(2)

	os := option.Applicative2(monoid.Sum[int]().Combine).
		ApOption(o1).
		ApOption(o2)

	fmt.Println(os)

	oint := option.Some(2)
	plus := func(a, b int) int {
		return a * b
	}
	option.Applicative2(plus).
		ApOption(oint).
		Ap(2)

	otuple := option.Some(as.Tuple(1, 2))
	option.Applicative1(as.Func2(plus).Tupled()).
		ApOption(otuple)

	oreader := option.Applicative3(fp.Id3[int, string, *strings.Reader]).
		ApOption(oint).
		Map(strconv.Itoa).
		Map(strings.NewReader)

	fmt.Println(oreader)
}
