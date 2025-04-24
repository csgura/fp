package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/slice"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type StateT[C, A any] struct {
}

func Pure[C, A any](a A) StateT[C, A] {
	panic("")
}

func Map[C, A, B any](s StateT[C, A], f func(A) B) StateT[C, B] {
	panic("")
}

func FlatMap[C, A, B any](s StateT[C, A], f func(A) StateT[C, B]) StateT[C, B] {
	panic("")
}

// @fp.Generate
func _[C, A, B any]() genfp.GenerateMonadTransformer[StateT[C, fp.Slice[A]]] {
	return genfp.GenerateMonadTransformer[StateT[C, fp.Slice[A]]]{
		File:     "slicest_generate.go",
		TypeParm: genfp.TypeOf[A](),
		ExposureMonad: genfp.MonadFunctions{
			Pure:    slice.Pure[A],
			FlatMap: slice.FlatMap[A, B],
		},
		Sequence: func(v fp.Slice[StateT[C, A]]) StateT[C, fp.Slice[A]] {
			panic("not implemented")
		},
		Transform: []any{
			slice.Filter[A],
			slice.Add[A],
			slice.Append[A],
			slice.Concat[A],
			slice.Drop[A],
			slice.Exists[A],
			slice.FilterNot[A],
			slice.Find[A],
			slice.ForAll[A],
			slice.Foreach[A],
			slice.Get[A],
			slice.Head[A],
			slice.Tail[A],
			slice.Init[A],
			slice.IsEmpty[A],
			slice.Last[A],
			slice.MakeString[A],
			slice.NonEmpty[A],
			slice.Reverse[A],
			slice.Size[A],
			slice.Take[A],
			slice.Fold[A, B],
			slice.Scan[A, B],
			slice.Sort[A],
			slice.Min[A],
			slice.Max[A],
			slice.FilterMap[A, B],

			// TODO: SPAN (multi value return)
			// ToGoMAP :  comparable constraints
			// FoldTry : -> flatten
		},
	}
}
