package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/slice"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type SliceST[C, V any] = fp.StateT[C, fp.Slice[V]]

func Pure[C, A any](v A) fp.StateT[C, A] {
	panic("")
}

func Map[C, A, B any](v fp.StateT[C, A], f func(A) B) fp.StateT[C, B] {
	panic("")
}

func FlatMap[C, A, B any](v fp.StateT[C, A], f func(A) fp.StateT[C, B]) fp.StateT[C, B] {
	panic("")
}

// @fp.Generate
func _[C, A, B, V any, K comparable]() genfp.GenerateMonadTransformer[SliceST[C, A]] {
	return genfp.GenerateMonadTransformer[SliceST[C, A]]{
		File:     "gendebug_generated.go",
		TypeParm: genfp.TypeOf[A](),
		// GivenMonad: genfp.MonadFunctions{
		// 	Pure:    statet.Pure[C, A],
		// 	FlatMap: statet.FlatMap[C, A, B],
		// },
		ExposureMonad: genfp.MonadFunctions{
			Pure:    slice.Pure[A],
			FlatMap: slice.FlatMap[A, B],
		},
		Sequence: func(v fp.Slice[fp.StateT[C, A]]) SliceST[C, A] {
			panic("")
		},
		Name: "SliceST",
		Transform: []any{
			slice.Add[A],
		},
	}
}
