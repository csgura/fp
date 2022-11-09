package testjson_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/test/internal/js"
	"github.com/csgura/fp/test/internal/testjson"
)

func TestEncode(t *testing.T) {
	root := testjson.RootMutable{
		A: 10,
		B: "b",
		C: 2.5,
		D: true,
		E: nil,
		F: []int{1, 2, 3, 4},
		G: map[string]int{
			"k1": 1,
			"k2": 2,
		},
	}.AsImmutable()

	str := testjson.EncoderRoot.Encode(root)
	fmt.Println(str.Get())

	rev := testjson.DecoderRoot.Decode(js.DecoderContext{}, str.Get())
	rev.Failed().Foreach(fp.Println[error])
	assert.True(rev.IsSuccess())
	fmt.Println(rev.Get())
}
