package reflectfp_test

import (
	"reflect"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/either"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/reflectfp"
)

func TestReflectOption(t *testing.T) {

	opt := option.Some(10)

	isopt := reflectfp.MatchOption(reflect.TypeOf(opt))
	assert.True(isopt.IsDefined())

	islazy := reflectfp.MatchLazyEval(reflect.TypeOf(opt))
	assert.False(islazy.IsDefined())

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

func TestReflectLazy(t *testing.T) {
	l := lazy.Done(10)

	islazy := reflectfp.MatchLazyEval(reflect.TypeOf(l))
	assert.True(islazy.IsDefined())
}

func TestReflectLazyConvert(t *testing.T) {

	lv := lazy.Done(10)

	isopt := reflectfp.MatchOption(reflect.TypeOf(lv))
	assert.False(isopt.IsDefined())

	ialazy := reflectfp.MatchLazyEval(reflect.TypeOf(lv))
	assert.True(ialazy.IsDefined())

	res := reflectfp.LazyCall(reflect.TypeOf(lv), func() reflect.Value {
		return reflect.ValueOf(20.2)
	})

	assert.True(res.IsSuccess())
	s := res.Get().Interface().(lazy.Eval[int]).Get()
	assert.Equal(s, 20)
}

func TestEither(t *testing.T) {

	l := either.Left[int, string](10)
	res := reflectfp.MatchEither(reflect.TypeOf(l))

	assert.True(res.IsDefined())
	assert.True(res.Get().Left == reflectfp.TypeOf[int]())
	assert.True(res.Get().Right == reflectfp.TypeOf[string]())

	res = reflectfp.MatchEither(reflectfp.TypeOf[fp.Either[int, string]]())

	assert.True(res.IsDefined())
	assert.True(res.Get().Left == reflectfp.TypeOf[int]())
	assert.True(res.Get().Right == reflectfp.TypeOf[string]())
}
