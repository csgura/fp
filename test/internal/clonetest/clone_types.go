package clonetest

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/clone"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
type ValueStruct struct {
	hello string
	world int
}

// @fp.Value
type CloneStruct struct {
	hello string
	world int
}

// @fp.Derive
var _ clone.Derives[fp.Clone[CloneStruct]]

type RecursiveDerive struct {
	S []string
}

type MySeq []string

type HasReference struct {
	A  *string
	S  []int
	M  map[string]int
	RD RecursiveDerive
	T  time.Time
	MS MySeq
	VS ValueStruct
	CS CloneStruct
}

// @fp.Derive(recursive=true)
var _ clone.Derives[fp.Clone[HasReference]]
