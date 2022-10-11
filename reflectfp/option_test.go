package reflectfp_test

import (
	"reflect"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/reflectfp"
)

func TestReflectOption(t *testing.T) {

	opt := option.Some("10")

	isopt := reflectfp.IsOptionType(reflect.TypeOf(opt))
	assert.True(isopt)

	res := reflectfp.Some(reflect.TypeOf(opt), reflect.ValueOf("20"))

	s := res.Get().Interface().(fp.Option[string]).Get()
	assert.Equal(s, "20")
}
