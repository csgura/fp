package match_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/match"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func TestMatch(t *testing.T) {

	v := option.Some(11)

	r := match.Of(v,
		match.Case(match.SomeAnd(match.Equal(10)), func(v int) string {
			return "10"
		}),
		match.Case(match.Some, func(v int) string {
			return "some"
		}),
		match.CaseNone[int](func() string {
			return "none"
		}),
	)

	fmt.Printf("r = %s\n", r)

	t2 := as.Tuple2(option.Some("hello"), option.None[int]())

	r = match.Of(t2,
		match.Case(match.Tuple2(match.Some[string], match.Some[int]), as.Tupled2(func(v1 string, v2 int) string {
			return "some,some"
		})),
		match.Case(match.Tuple2(match.Some[string], match.None[int]), func(t fp.Tuple2[string, fp.Unit]) string {
			return "some,none"

		}),
		match.CaseTuple2(match.None[string], match.Some, func(u fp.Unit, v int) string {
			return "none,some"
		}),
	)

	fmt.Printf("r = %s\n", r)

	s := seq.Of(1, 2)

	r = match.Of(s,
		match.CaseCons[fp.Seq[int]](match.Some[int], match.Head[fp.Seq[int]](match.SomeAnd(match.Equal(2))), func(h int, h2 int) string {
			return "int,2"
		}),
		match.CaseCons[fp.Seq[int]](match.Some[int], match.Any[fp.Seq[int]], func(h int, tail fp.Seq[int]) string {
			return "head"
		}),
		match.Case(match.Any, func(v fp.Seq[int]) string {
			return "list"
		}),
	)

	fmt.Printf("r = %s\n", r)

}
