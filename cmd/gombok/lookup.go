package main

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) lookupTupleLikeTypeClassFunc(ctx SummonContext, tc metafp.TypeClass, name string, fieldNames fp.Seq[fp.NameTag], tupleArgs fp.Seq[metafp.TypeInfoExpr]) fp.Option[metafp.TypeClassInstance] {

	workingScope := ctx.workingScope(r.tcCache, tc)
	primScope := ctx.primScope(r.tcCache, tc)

	argType := seq.Map(tupleArgs, func(v metafp.TypeInfoExpr) metafp.TypeInfo {
		return v.Type
	})

	ins := workingScope.FindTupleLikeByNamePrefix(name, fieldNames, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindTupleLikeByNamePrefix(name, fieldNames, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	return option.None[metafp.TypeClassInstance]()
}
