package testpk1

import (
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/read"
	"github.com/csgura/fp/test/internal/show"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Value
// @fp.Json
// @fp.GenLabelled
type World struct {
	message   string
	timestamp time.Time
}

// @fp.Derive
var _ eq.Derives[fp.Eq[World]]

// @fp.Derive
var _ js.Derives[js.Encoder[World]]

// @fp.Derive
var _ js.Derives[js.Decoder[World]]

// @fp.Derive
var _ show.Derives[fp.Show[World]]

// @fp.Value
// @fp.GenLabelled
type HasOption struct {
	message  string
	addr     fp.Option[string]
	phone    []string
	emptySeq []int
}

// @fp.Derive
var _ js.Derives[js.Encoder[HasOption]]

// @fp.Value
type CustomValue struct {
	a string
	b int
}

func (r CustomValue) A() string {
	return "hello" + r.a
}

func (r CustomValue) WithB(v int) CustomValue {
	if v > 0 {
		r.b = v
	}
	return r
}

type CustomValueBuilder CustomValue

func (r CustomValueBuilder) B(v int) CustomValueBuilder {
	if v > 0 {
		r.b = v
	}
	return r
}

// // @fp.Value
// type NotDerivable struct {
// 	a int
// 	b interface {
// 		Hello() string
// 	}
// }

// // @fp.Derive
// var _ monoid.Derives[fp.Monoid[NotDerivable]]

// var MonoidInt = monoid.Sum[int]()

// @fp.Value
type AliasedStruct World

// @fp.Derive
var _ eq.Derives[fp.Eq[AliasedStruct]]

// @fp.Value
type HListInsideHList struct {
	tp    fp.Tuple2[string, int]
	value string
	hello World
}

// @fp.Derive
var _ show.Derives[fp.Show[HListInsideHList]]

// @fp.Derive
var _ read.Derives[read.Read[HListInsideHList]]

// @fp.Derive
var _ read.Derives[read.Read[World]]
