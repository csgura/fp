//go:build go1.27

package reflectfp_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/reflectfp"
	"github.com/csgura/fp/should"
)

func TestEitherO(t *testing.T) {
	the := should.Test(t)

	e := fp.EitherO[string, int]{}.Reset(nil, new(10))
	the.Option(reflectfp.MatchEitherO(reflect.TypeOf(e))).ShouldBeSome()
	v := the.OptionT(reflectfp.GetRight(reflect.ValueOf(e))).ShouldBeSome()
	the.Value(fmt.Sprint(v.Interface())).ShouldEqual("10")

	s := e.Swap()
	v = the.OptionT(reflectfp.GetLeft(reflect.ValueOf(s))).ShouldBeSome()
	the.Value(fmt.Sprint(v.Interface())).ShouldEqual("10")

	eto := reflectfp.MatchEitherO(reflectfp.TypeOf[fp.EitherO[string, int]]())
	et := the.Option(eto).ShouldBeSome()
	vt := et.Left(reflect.ValueOf("hello"))
	lv := the.Try(vt).ShouldBeSuccess().Interface().(fp.EitherO[string, int])
	the.Value(lv.Left()).ShouldEqual("hello")

	vt = et.Right(reflect.ValueOf(10))
	rightv := the.Try(vt).ShouldBeSuccess().Interface().(fp.EitherO[string, int])
	the.Value(rightv.Get()).ShouldEqual(10)

}
