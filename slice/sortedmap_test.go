package slice_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/should"
	"github.com/csgura/fp/slice"
)

func TestSortedMap(t *testing.T) {
	s := slice.Of(
		as.Tuple(1, 1),
		as.Tuple(3, 3),
		as.Tuple(5, 5),
		as.Tuple(7, 7),
		as.Tuple(9, 9),
		as.Tuple(11, 11),
	)
	sm := slice.ToSortedMap(s, ord.Given[int]())
	should.Equal(t, sm.TailMap(5).Iterator().NextOption(), option.Some(as.Tuple(5, 5)))
	should.Equal(t, sm.TailMap(6).Iterator().NextOption(), option.Some(as.Tuple(7, 7)))
	should.Equal(t, sm.TailMap(0).Iterator().NextOption(), option.Some(as.Tuple(1, 1)))
	should.Equal(t, sm.TailMap(20).Iterator().NextOption(), option.None[fp.Tuple2[int, int]]())

}
