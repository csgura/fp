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

	opt := option.Some(10)

	isopt := reflectfp.MatchOption(reflect.TypeOf(opt))
	assert.True(isopt.IsDefined())

	res := reflectfp.Some(reflect.TypeOf(opt), reflect.ValueOf(20))

	assert.True(res.IsSuccess())
	s := res.Get().Interface().(fp.Option[int]).Get()
	assert.Equal(s, 20)
}

func TestReflectOptionConvert(t *testing.T) {

	opt := option.Some(10)

	isopt := reflectfp.MatchOption(reflect.TypeOf(opt))
	assert.True(isopt.IsDefined())

	res := reflectfp.Some(reflect.TypeOf(opt), reflect.ValueOf(20.2))

	assert.True(res.IsSuccess())
	s := res.Get().Interface().(fp.Option[int]).Get()
	assert.Equal(s, 20)
}
