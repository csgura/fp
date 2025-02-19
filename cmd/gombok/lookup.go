package main

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) lookupTupleLikeTypeClassFunc(ctx SummonContext, tc metafp.TypeClass, name string, fieldNames fp.Seq[fp.NameTag], tupleArgs fp.Seq[metafp.TypeInfoExpr]) fp.Option[DefinedInstance] {

	workingScope := ctx.workingScope(r.tcCache, tc)
	primScope := ctx.primScope(r.tcCache, tc)

	argType := seq.Map(tupleArgs, func(v metafp.TypeInfoExpr) metafp.TypeInfo {
		return v.Type
	})

	insLocal := workingScope.FindTupleLikeByNamePrefix(name, fieldNames, argType)
	if insLocal.IsDefined() && r.checkRequired(ctx, insLocal.Get(), insLocal.Get().RequiredInstance) {
		return option.Some(DefinedInstance{
			instance:   insLocal.Get(),
			searchName: name,
			checked:    true,
		})
	}

	ins := primScope.FindTupleLikeByNamePrefix(name, fieldNames, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return option.Some(DefinedInstance{
			instance:   ins.Get(),
			searchName: name,
			checked:    true,
		})
	}

	return option.Map(insLocal.OrOption(ins), func(a metafp.TypeClassInstance) DefinedInstance {
		return DefinedInstance{
			instance:   a,
			searchName: name,
			checked:    true,
		}
	})
}
