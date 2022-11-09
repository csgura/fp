package testjson_test

import (
	"fmt"
	"testing"

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

	fmt.Println(testjson.EncoderRoot.Encode(root))
}
