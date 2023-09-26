package adaptortest_test

import (
	"context"
	"testing"

	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/test/internal/adaptortest"
)

func TestSpanContext(t *testing.T) {
	a := adaptortest.SpanContextAdaptor{
		DefaultContext: context.Background(),
	}

	assert.Equal(a.Err(), nil)
}
