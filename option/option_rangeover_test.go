//go:build go1.23

package option_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/option"
)

func TestSomeRangeOver(t *testing.T) {
	opt := option.Some("10")

	for v := range opt.All() {
		fmt.Printf("value = %s\n", v)
	}
}
