package main_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/unit"
)

type hello struct {
}

func (r hello) String() string {
	return "hello"
}

func TestFunc(t *testing.T) {
	ovalue := option.Some(hello{})

	ovalue.Map(hello.String)
	ovalue.Map(func(v hello) string {
		return v.String()
	})

	optr := option.Some(&hello{})
	optr.Map((*hello).String)

	product := func(a, b int) int {
		return a * b
	}

	oint := option.Some(2)
	oint.Map(as.Curried2(product)(2))
	oint.Map(curried.Func2(
		monoid.Product[int]().Combine)(2),
	)

	oint.Map(func(a int) int {
		return a * 2
	})

	otuple := option.Some(as.Tuple(1, 2))
	otuple.Map(as.Tupled2(product))

	otuple.Map(as.Tupled2(monoid.Product[int]().Combine))

	otuple.Map(func(tuple fp.Tuple2[int, int]) int {
		return tuple.I1 + tuple.I2
	})

	// f := as.Func2(strconv.FormatInt).Shift()
	// ostr := option.Map(oint, as.InstanceOf(v any))

	ostr := oint.Map(strconv.Itoa)
	oreader := ostr.Map(strings.NewReader)
	fmt.Println(oreader)

	oreader = oint.Map(fp.Compose2(
		strconv.Itoa,
		strings.NewReader,
	))
	fmt.Println(oreader)

	tstr := try.Success("25380")
	tint := try.FlatMap(tstr, try.Func1(strconv.Atoi))
	tprocess := try.FlatMap(tint, try.Func1(os.FindProcess))
	killResult := try.FlatMap(tprocess, try.Unit1((*os.Process).Kill))
	fmt.Println(killResult)

	killResult = try.FlatMap(tstr, try.Compose3(
		try.Func1(strconv.Atoi),
		try.Func1(os.FindProcess),
		try.Unit1((*os.Process).Kill),
	))
	fmt.Println(killResult)

	o1 := option.Some(1)
	o2 := option.Some(2)

	os2 := option.FlatMap(o1, func(a int) fp.Option[int] {
		return o2.Map(func(b int) int {
			return a + b
		})
	})

	fmt.Println(os2)

	f3 := func(a, b, c int) int {
		return a + b + c
	}

	f2 := curried.Revert2(as.Curried3(f3)(1))
	f2(2, 3)

	format16 := curried.Flip(as.Curried2(strconv.FormatInt))(16)
	format16(123456)

	f := fp.Compose(strconv.Itoa, option.Some)
	fmt.Println(f(20))

	tf := try.Func1(strconv.Atoi)
	tf("20")

	tu := try.Unit3(os.Chown)
	tu("a.txt", 1, 2)

	unit.Func2(func(int, int) {
	})

	as.Func0(os.Environ)

	a := func() {}
	b := func() {}
	c := func() {}

	fp.Compose3(unit.Func0(a), unit.Func0(b), unit.Func0(c))
}
