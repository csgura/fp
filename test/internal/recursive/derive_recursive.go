package recursive

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/show"
	"github.com/csgura/fp/test/internal/testpk1"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type NormalStruct struct {
	Name    string
	Age     int
	Address string
}

func (r NormalStruct) Print() {

}

// @fp.Derive
var _ eq.Derives[fp.Eq[NormalStruct]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[NormalStruct]]

// @fp.Derive
var _ show.Derives[fp.Show[NormalStruct]]

// @fp.Derive
var _ js.Derives[js.Encoder[NormalStruct]]

// @fp.Derive
var _ js.Derives[js.Decoder[NormalStruct]]

type Tuple2Struct struct {
	Name string
	Age  int
}

// @fp.Derive
var _ show.Derives[fp.Show[Tuple2Struct]]

// @fp.Derive
var _ js.Derives[js.Encoder[Tuple2Struct]]

// @fp.Derive
var _ js.Derives[js.Decoder[Tuple2Struct]]

type Over21[T any] struct {
	I1  T
	I2  int
	I3  int
	I4  int
	I5  int
	I6  int
	I7  int
	I8  int
	I9  int
	I10 int
	I11 int
	I12 int
	I13 int
	I14 int
	I15 int
	I16 int
	I17 int
	I18 int
	I19 int
	I20 int
	I21 int
	I22 int
	I23 int
	I24 int
	I25 int
	I26 int
	I27 int
	I28 int
	I29 int
	I30 int
}

// @fp.Derive
var _ eq.Derives[fp.Eq[Over21[any]]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[Over21[any]]]

// @fp.Derive
var _ show.Derives[fp.Show[Over21[any]]]

// @fp.Derive
var _ js.Derives[js.Encoder[Over21[any]]]

// @fp.Derive
var _ js.Derives[js.Decoder[Over21[any]]]

// @fp.Derive
var _ eq.Derives[fp.Eq[testpk1.LegacyStruct]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[testpk1.LegacyStruct]]

// @fp.Derive
var _ show.Derives[fp.Show[testpk1.LegacyStruct]]

// @fp.Derive
var _ js.Derives[js.Encoder[testpk1.LegacyStruct]]

// @fp.Derive
var _ js.Derives[js.Decoder[testpk1.LegacyStruct]]

// @fp.Derive
var _ monoid.Derives[fp.Monoid[testpk1.LegacyStructCompose]]

// @fp.Derive
var _ eq.Derives[fp.Eq[testpk1.LegacyPhoneBook]]

// @fp.Derive(recursive=true)
var _ monoid.Derives[fp.Monoid[testpk1.LegacyPhoneBook]]
