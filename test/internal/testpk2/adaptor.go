// Code generated by gombok, DO NOT EDIT.
package testpk2

import (
	"github.com/csgura/fp/test/internal/testpk1"
)

type AliasIntfAdaptor struct {
	DoCtx   func(self AliasIntf) testpk1.Pk1Context
	DoOther func(self AliasIntf) Pk1Context
}

func (r *AliasIntfAdaptor) Ctx() testpk1.Pk1Context {
	return r.CtxImpl(r)
}

func (r *AliasIntfAdaptor) CtxImpl(self AliasIntf) testpk1.Pk1Context {

	if r.DoCtx != nil {
		return r.DoCtx(self)
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