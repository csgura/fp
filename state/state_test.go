package state_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/state"
)

func popPtr[T any](stack *[]T) fp.Option[T] {
	l := len(*stack)
	if l == 0 {
		return option.None[T]()
	}
	ret := option.Some((*stack)[l-1])

	*stack = (*stack)[0 : l-1]
	return ret
}

func pop[T any](stack []T) (fp.Option[T], []T) {
	l := len(stack)
	if l == 0 {
		return option.None[T](), nil
	}
	return option.Some(stack[l-1]), stack[:l-1]
}

func pushState[T any](v T) fp.State[[]T, fp.Unit] {
	return func(stack []T) fp.Tuple2[fp.Unit, []T] {
		return as.Tuple(fp.Unit{}, append(stack, v))
	}
}

func popState[T any]() fp.State[[]T, fp.Option[T]] {
	return func(stack []T) fp.Tuple2[fp.Option[T], []T] {
		l := len(stack)
		if l == 0 {
			return as.Tuple(option.None[T](), stack)
		}
		return as.Tuple(option.Some(stack[l-1]), stack[:l-1])
	}
}

func pushOption[T any](v fp.Option[T]) fp.State[[]T, fp.Unit] {
	return func(stack []T) fp.Tuple2[fp.Unit, []T] {
		return as.Tuple(fp.Unit{}, append(stack, v.ToSeq()...))
	}
}

func calcPlus[T any](m fp.Monoid[T]) fp.State[[]T, fp.Unit] {
	v1 := popState[T]()
	v2 := popState[T]()

	sum := state.Map2(v1, v2, monoid.Option(m).Combine)
	return state.FlatMap(sum, pushOption)
}

func peekHighOrder[T any]() fp.State[[]T, fp.Option[T]] {
	return func(t []T) fp.Tuple2[fp.Option[T], []T] {
		return as.Tuple(seq.Last(t), t)
	}
}

func peek[T any]() fp.State[[]T, fp.Option[T]] {
	return state.Map(state.Get[[]T](), seq.Last)
}

func peekS[T any]() fp.State[[]T, fp.Option[T]] {
	return state.GetS[[]T](seq.Last)
}

func size[T any]() fp.State[[]T, int] {
	return state.Map(state.Get[[]T](), seq.Size)
}

func calcRecursive[T any](m fp.Monoid[T]) fp.State[[]T, fp.Unit] {
	res := state.FlatMapConst(calcPlus(m), size[T]())

	return state.FlatMap(res, func(a int) fp.State[[]T, fp.Unit] {
		if a > 1 {
			return calcRecursive(m)
		}
		return state.Replace(state.Get[[]T](), fp.Unit{})
	})
}

func TestPop(t *testing.T) {
	s := calcPlus(monoid.Sum[int]())

	p := state.FlatMapConst(s, peek[int]())

	fmt.Printf("1 + 2 = %v\n", p.Eval(seq.Of(1, 2)))
	fmt.Printf("1 + None = %v\n", p.Eval(seq.Of(1)))

	a := calcRecursive(monoid.Sum[int]())

	p = state.FlatMapConst(a, peek[int]())

	fmt.Printf("1 + 2 + 3 + 4 + 5 = %v\n", p.Eval(seq.Of(1, 2, 3, 4, 5)))

}

func TestCalcRecursive(t *testing.T) {

	a := calcRecursive(monoid.Sum[int]())

	p := state.FlatMapConst(a, peek[int]())

	fmt.Printf("1 + 2 + 3 + 4 + 5 = %v\n", p.Eval(seq.Of(1, 2, 3, 4, 5)))

}
