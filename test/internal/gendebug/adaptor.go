// Code generated by gombok, DO NOT EDIT.
package gendebug

import (
	"github.com/csgura/fp/test/internal/testpk1"
	"github.com/csgura/fp/test/internal/testpk2"
)

type AliasIntfAdaptor struct {
	DoCtx    func(self AliasIntf, arg any) testpk1.Pk1Context
	DoOther  func(self AliasIntf) Pk1Context
	DoPk2Ctx func(self AliasIntf) testpk2.Pk1Context
	DoSome   func(self AliasIntf, ctxs ...testpk2.Pk1Context) (string, error)
}

func (r *AliasIntfAdaptor) Ctx(arg any) testpk1.Pk1Context {
	return r.CtxImpl(r, arg)
}

func (r *AliasIntfAdaptor) CtxImpl(self AliasIntf, arg any) testpk1.Pk1Context {

	if r.DoCtx != nil {
		return r.DoCtx(self, arg)
	}

	panic("AliasIntfAdaptor.Ctx not implemented")
}

func (r *AliasIntfAdaptor) Other() Pk1Context {
	return r.OtherImpl(r)
}

func (r *AliasIntfAdaptor) OtherImpl(self AliasIntf) Pk1Context {

	if r.DoOther != nil {
		return r.DoOther(self)
	}

	panic("AliasIntfAdaptor.Other not implemented")
}

func (r *AliasIntfAdaptor) Pk2Ctx() testpk2.Pk1Context {
	return r.Pk2CtxImpl(r)
}

func (r *AliasIntfAdaptor) Pk2CtxImpl(self AliasIntf) testpk2.Pk1Context {

	if r.DoPk2Ctx != nil {
		return r.DoPk2Ctx(self)
	}

	panic("AliasIntfAdaptor.Pk2Ctx not implemented")
}

func (r *AliasIntfAdaptor) Some(ctxs ...testpk2.Pk1Context) (string, error) {
	return r.SomeImpl(r, ctxs...)
}

func (r *AliasIntfAdaptor) SomeImpl(self AliasIntf, ctxs ...testpk2.Pk1Context) (string, error) {

	if r.DoSome != nil {
		return r.DoSome(self, ctxs...)
	}

	panic("AliasIntfAdaptor.Some not implemented")
}
