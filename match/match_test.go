package match_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/match"
	"github.com/csgura/fp/option"
)

func TestMatch(t *testing.T) {

	v := option.Some(11)

	r := match.Of(v,
		match.Case(match.SomeAnd(match.Equal(10)), func(v int) string {
			return "10"
		}),
		match.Case(match.Some[int], func(v int) string {
			return "some"
		}),
		match.Case(match.None[int], func(t fp.Unit) string {
			return "none"
		}),
	)

	fmt.Printf("r = %s\n", r)

	t2 := as.Tuple2(option.Some("hello"), option.None[int]())

	r = match.Of(t2,
		match.Case(match.Tuple2(match.Some[string], match.Some[int]), func(t fp.Tuple2[string, int]) string {
			return "some,some"
		}),
		match.Case(match.Tuple2(match.Some[string], match.None[int]), func(t fp.Tuple2[string, fp.Unit]) string {
			return "some,none"

		}),
		match.Case(match.Tuple2(match.None[string], match.Some[int]), func(t fp.Tuple2[fp.Unit, int]) string {
			return "none,some"
		}),
	)

	fmt.Printf("r = %s\n", r)

}
