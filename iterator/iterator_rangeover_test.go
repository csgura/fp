//go:build go1.23

package iterator_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/iterator"
)

func TestRangeOver(t *testing.T) {
	itr := iterator.Of(1, 2, 3, 4)

	for v := range itr.All() {
		fmt.Printf("v = %d\n", v)
		if v > 1 {
			break
		}
	}
}
