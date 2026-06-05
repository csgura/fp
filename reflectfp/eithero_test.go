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
	should := should.Test(t)

	e := new(fp.EitherO[string, int]).Apply(nil, new(10))
	should.BeSome(reflectfp.MatchEitherO(reflect.TypeOf(e)))
	v := should.OptionT(reflectfp.GetRight(reflect.ValueOf(e))).BeSome()
	should.Value(fmt.Sprint(v.Interface())).Equal("10")

	s := e.Swap()
	v = should.OptionT(reflectfp.GetLeft(reflect.ValueOf(s))).BeSome()
	should.Value(fmt.Sprint(v.Interface())).Equal("10")

	eto := reflectfp.MatchEitherO(reflectfp.TypeOf[fp.EitherO[string, int]]())
	et := should.BeSome(eto)
	vt := et.Left(reflect.ValueOf("hello"))
	lv := should.BeSuccess(vt).Interface().(fp.EitherO[string, int])
	should.Value(lv.Left()).Equal("hello")

	vt = et.Right(reflect.ValueOf(10))
	rightv := should.BeSuccess(vt).Interface().(fp.EitherO[string, int])
	should.Value(rightv.Get()).Equal(10)

}
