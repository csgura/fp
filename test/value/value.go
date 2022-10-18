package value

import "github.com/csgura/fp"

//go:generate go run github.com/csgura/fp/cmd/gombok
type (
	// Hello is hello
	// @fp.Value
	Hello struct { // Hello
		world string
		hi    int
	}
)

// @fp.Value
type MyMy struct { // what the
	// hihi
	hi fp.Option[int]
}

type NoValue struct {
}
