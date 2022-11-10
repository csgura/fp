package testjson

import "github.com/csgura/fp/test/internal/js"

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.GenLabelled
type Root struct {
	a int
	b string
	c float64
	d bool
	e *int
	f []int
	g map[string]int
	h Child
}

// @fp.Value
// @fp.GenLabelled
type Child struct {
	a map[string]any
	b any
}

// @fp.Derive
var _ js.Derives[js.Encoder[Root]]

// @fp.Derive
var _ js.Derives[js.Decoder[Root]]

// @fp.Derive
var _ js.Derives[js.Encoder[Child]]

// @fp.Derive
var _ js.Derives[js.Decoder[Child]]

// @fp.Value
// @fp.GenLabelled
type Node struct {
	name  string
	left  *Node
	right *Node
}

// @fp.Derive
var _ js.Derives[js.Encoder[Node]]

// @fp.Value
// @fp.GenLabelled
type Tree struct {
	root *Node
}

// @fp.Derive
var _ js.Derives[js.Encoder[Tree]]

// @fp.Value
// @fp.GenLabelled
type Entry[V any] struct {
	name  string
	value V
}

// @fp.Derive
var _ js.Derives[js.Encoder[Entry[any]]]

// @fp.Value
// @fp.GenLabelled
type NotUsedParam[K, V any] struct {
	param string
	value V
}

// @fp.Derive
var _ js.Derives[js.Encoder[NotUsedParam[any, any]]]

// @fp.Value
// @fp.GenLabelled
type Movie struct {
	name    string
	casting Entry[string]
	//notUsed NotUsedParam[int, string]
}

// @fp.Derive
var _ js.Derives[js.Encoder[Movie]]
