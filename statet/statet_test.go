package statet_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/statet"
	"github.com/csgura/fp/try"
)

func push[T any](v T) fp.StateT[fp.Seq[T], fp.Unit] {
	return statet.PutWith(fp.Seq[T].Add)(v)
}

func pop[T any]() fp.StateT[fp.Seq[T], T] {
	return statet.MapT(statet.Merge(fp.Seq[T].Init, fp.Seq[T].Last), try.FromOption)
}

func calc[T any](op fp.Monoid[T]) fp.StateT[fp.Seq[T], fp.Unit] {
	v1 := pop[T]()
	v2 := pop[T]()

	sum := statet.Map2(v1, v2, op.Combine)
	return statet.FlatMap(sum, push)

}

func TestStateT(t *testing.T) {

	sum := statet.FlatMapConst(calc(monoid.Sum[int]()), pop[int]()).Eval(seq.Of(1, 2, 3, 4))
	assert.True(sum.IsSuccess())
	assert.Equal(sum.Get(), 7)
}

func TestRecover(t *testing.T) {
	rv := pop[int]().RecoverWith(func(err error) fp.StateT[fp.Seq[int], int] {
		return func(s fp.Seq[int]) (fp.Try[int], fp.Seq[int]) {
			return try.Success(10), nil
		}
	}).Eval(seq.Of[int]())
	assert.True(rv.IsSuccess())
	assert.Equal(rv.Get(), 10)
}
