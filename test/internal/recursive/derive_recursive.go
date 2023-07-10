package recursive

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/test/internal/show"
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
var _ show.Derives[fp.Show[NormalStruct]]

// // @fp.Derive
// var _ js.Derives[js.Encoder[NormalStruct]]
