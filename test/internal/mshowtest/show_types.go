package mshowtest

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/test/internal/recursive"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type Person struct {
	Name string
	Age  int
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[Person]]

type NoDerive struct {
	Hello string
}

type Collection struct {
	Index       map[string]Person
	List        []Person
	Description *string
	Set         fp.Set[int]
	Option      fp.Option[Person]
	NoDerive    NoDerive
	Stringer    HasStringMethod
	BoolPtr     *bool
	NoMap       map[string]NoDerive
	Alias       recursive.StringAlias
	StringSeq   fp.Seq[string]
}

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[Collection]]

type HasStringMethod struct {
	There string
}

func (r HasStringMethod) String() string {
	return r.There
}

type DupGenerate struct {
	NoDerive NoDerive
	World    string
}

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[DupGenerate]]

type HasTuple struct {
	Entry fp.Tuple2[string, int]
	HList hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[HasTuple]]

// @fp.Value
type EmbeddedStruct struct {
	hello string
	world struct {
		Level int
		Stage string
	}
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[EmbeddedStruct]]

// @fp.Value
type EmbeddedTypeParamStruct[T any] struct {
	hello string
	world struct {
		Level T
		Stage string
	}
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[EmbeddedTypeParamStruct[any]]]

func UntypedStructFunc(s struct {
	Level   int
	Stage   string
	privacy string
}) {
	fmt.Println("Ok")
}

type EmptyStruct struct {
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[EmptyStruct]]

type TLV = []byte

type HasAliasType struct {
	Data TLV
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[HasAliasType]]
