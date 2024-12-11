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
	ctx := fp.WithContextValue(context.Background(), stringKey{}, "hello")

	v := fp.GetContextValue(ctx, stringKey{})
	assert.Equal(v, option.Some("hello"))

	v2 := fp.GetContextValue(ctx, intKey{})
	assert.True(v2.IsEmpty())

	v = fp.GetContextValue(ctx, otherStringKey{})
	assert.True(v.IsEmpty())

}
