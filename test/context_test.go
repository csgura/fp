package main_test

import (
	"context"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
)

type stringKey struct {
	fp.ValueType[string]
}

type intKey struct {
	fp.ValueType[int]
}

type otherStringKey struct {
	fp.ValueType[string]
}

func TestContextKey(t *testing.T) {
	ctx := fp.WithContextValue[stringKey](context.Background(), "hello")

	v := fp.GetContextValue[stringKey](ctx)
	assert.Equal(v, option.Some("hello"))

	v2 := fp.GetContextValue[intKey](ctx)
	assert.True(v2.IsEmpty())

	v = fp.GetContextValue[otherStringKey](ctx)
	assert.True(v.IsEmpty())

}
