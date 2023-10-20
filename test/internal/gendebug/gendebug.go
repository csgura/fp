package gendebug

import (
	"context"

	"github.com/csgura/fp/genfp"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.String(useShow=true)
// type ShowConstraintExplicit[T fmt.Stringer] struct {
// 	hello   string
// 	world   int
// 	message T
// }

// func ShowStringer[T fmt.Stringer]() fp.Show[T] {
// 	return show.New[T](func(t T) string {
// 		return t.String()
// 	})
// }

// @fp.Derive
//var _ show.Derives[fp.Show[ShowConstraintExplicit[fmt.Stringer]]]

// func ShowShowConstraintExplicit[T fmt.Stringer]() fp.Show[ShowConstraintExplicit[T]] {
// 	return show.New(func(t ShowConstraintExplicit[T]) string {
// 		return fmt.Sprintf("ShowConstraintExplicit(message=%s)", t.message)
// 	})
// }

// @fp.Generate
var _ = genfp.GenerateAdaptor[context.Context]{
	File:               "gendebug_delegate.go",
	Name:               "ContextWrapper",
	EmbeddingInterface: []genfp.TypeTag{genfp.TypeOf[context.Context]()},
	Options: []genfp.ImplOption{
		{
			Prefix: "Get",
			Method: context.Context.Value,
		},
	},
}
