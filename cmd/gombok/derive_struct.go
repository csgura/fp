package main

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) exprTupleWithName(ctx SummonContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typePkg *types.Package, structName string, names fp.Seq[fp.NameTag], typeArgs fp.Seq[metafp.TypeInfoExpr], genLabelled bool) SummonExpr {
	if len(typeArgs) > 0 {
		list := collectSummonExpr(seq.Map(seq.Zip(typeArgs, names), func(t fp.Tuple2[metafp.TypeInfoExpr, fp.NameTag]) SummonExpr {
			return r.summonRequired(ctx, metafp.RequiredInstance{
				TypeClass: ctx.typeClass,
				Type:      t.I1.Type,
			})
		}))

		retExpr := func() string {
			aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))
			fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

			names := seq.Map(names, func(v fp.NameTag) string {
				return fmt.Sprintf("%s.NameTag(`%s`,`%s`)", aspk, v.I1, v.I2)
			}).MakeString(",")

			return fmt.Sprintf("%s([]%s.Named{%s}, %s)", lt.PackagedName(r.w, ctx.working), fppk, names, list)
		}

		return newSummonExpr(retExpr, list.paramInstance)
	}

	return newSummonExpr(func() string {
		return lt.PackagedName(r.w, ctx.working)
	})
}

func (r *TypeClassSummonContext) summonTupleWithNameGenericRepr(ctx SummonContext, tc metafp.TypeClass, sf structFunctions) fp.Option[GenericRepr] {

	fields := sf.fields
	names := seq.Map(fields, func(v metafp.StructField) fp.NameTag {
		return as.NameTag(v.Name, v.Tag)
	})

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfoExpr {
		return v.TypeInfoExpr(ctx.working)
	})

	result := r.lookupTupleLikeTypeClassFunc(ctx, tc, fmt.Sprintf("Struct%d", typeArgs.Size()), names, sf.typeArgs)

	verbose("lookupTupleLikeTypeClassFunc Struct%d -> %t", typeArgs.Size(), result.IsDefined())
	return option.FlatMap(result, func(tm metafp.TypeClassInstance) fp.Option[GenericRepr] {

		requiredAllTypeClass := tm.RequiredInstance.ForAll(func(v metafp.RequiredInstance) bool {
			return v.TypeClass.Id() == tc.Id()
		})

		verbose("%s requiredAllTypeClass %t", tm.Name, requiredAllTypeClass)
		if requiredAllTypeClass {
			fnamed := option.TraverseSeq(seq.Zip(names, typeArgs), func(a fp.Tuple2[fp.NameTag, metafp.TypeInfoExpr]) fp.Option[metafp.TypeClassInstance] {
				return r.lookupExplicitNamedFunc(ctx, tc, a.I1, a.I2.Type)
			})
			if fnamed.IsDefined() {
				return option.Some(GenericRepr{
					Kind:         fp.GenericKindStruct,
					Type:         as.Supplier1(sf.typeStr, ctx.working),
					ReprType:     r.tupleReprType(ctx, sf, tm.Result.TypeArgs.Head()),
					ToReprExpr:   r.intoTupleRepr(ctx, sf, tm.Result.TypeArgs.Head()),
					FromReprExpr: r.fromTupleRepr(ctx, sf, tm.Result.TypeArgs.Head()),
					ReprExpr: func() SummonExpr {
						return r.exprTypeClassInstanceWithRequiredFound(ctx, tm, false, fnamed.Get())

						// return r.exprTupleWithName(ctx, tc, tm, sf.pack, sf.name, names, typeArgs, sf.namedGenerated)
					},
				})
			}
		}
		return option.Some(GenericRepr{
			Kind:         fp.GenericKindStruct,
			Type:         as.Supplier1(sf.typeStr, ctx.working),
			ReprType:     r.tupleReprType(ctx, sf, tm.Result.TypeArgs.Head()),
			ToReprExpr:   r.intoTupleRepr(ctx, sf, tm.Result.TypeArgs.Head()),
			FromReprExpr: r.fromTupleRepr(ctx, sf, tm.Result.TypeArgs.Head()),
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassInstance(ctx, tm, false)

				// return r.exprTupleWithName(ctx, tc, tm, sf.pack, sf.name, names, typeArgs, sf.namedGenerated)
			},
		})
	}).Or(func() fp.Option[GenericRepr] {
		for scons := range r.lookupTypeClassFunc(ctx, tc, "StructHCons").All() {
			fnamed := option.TraverseSeq(seq.Zip(names, typeArgs), func(a fp.Tuple2[fp.NameTag, metafp.TypeInfoExpr]) fp.Option[metafp.TypeClassInstance] {
				return r.lookupExplicitNamedFunc(ctx, tc, a.I1, a.I2.Type)
			})

			for named := range fnamed.All() {

				return option.Some(GenericRepr{
					Kind:         fp.GenericKindStruct,
					Type:         as.Supplier1(sf.typeStr, ctx.working),
					ReprType:     r.hlistReprType(ctx, sf, scons.Result.TypeArgs.Head()),
					ToReprExpr:   r.toHlistRepr(ctx, sf, scons.Result.TypeArgs.Head()),
					FromReprExpr: r.fromHlistRepr(ctx, sf, scons.Result.TypeArgs.Head()),
					ReprExpr: func() SummonExpr {
						arity := typeArgs.Size()

						hnil := r.lookupTypeClassFunc(ctx, tc, "StructHNil").OrElseGet(as.Supplier2(r.lookupHNilMust, ctx, tc))

						zipped := named.Take(arity).Reverse()
						hlist := seq.Fold(zipped, newSummonExpr(func() string { return hnil.PackagedName(r.w, ctx.working) }), func(tail SummonExpr, ti metafp.TypeClassInstance) SummonExpr {

							instance := r.exprTypeClassInstance(ctx, ti, false)

							return newSummonExpr(func() string {
								return fmt.Sprintf(`%s(
										%s,
										%s,
									)`, scons.PackagedName(r.w, ctx.working), instance, tail,
								)
							}, instance.paramInstance, tail.paramInstance)
						})
						return hlist
					},
				})
			}

		}
		return option.None[GenericRepr]()

	})
}

func (r *TypeClassSummonContext) lookupExplicitNamedFunc(ctx SummonContext, tc metafp.TypeClass, name fp.NameTag, argType metafp.TypeInfo) fp.Option[metafp.TypeClassInstance] {
	workingScope := ctx.workingScope(r.tcCache, tc)
	primScope := ctx.primScope(r.tcCache, tc)

	ins := workingScope.FindFuncHasNameArg(name, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindFuncHasNameArg(name, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	return option.None[metafp.TypeClassInstance]()
}

func (r *TypeClassSummonContext) structFieldNameTypeStr(ctx SummonContext, sf structFunctions, field metafp.StructField) string {

	name := field.Name
	valueType := field.TypeInfoExpr(ctx.working).TypeName(r.w, ctx.working)
	if sf.namedGenerated {
		ret := publicName(name)
		if ret == name {
			ret = fmt.Sprintf("PubNamed%sOf%s", ret, sf.name)
		} else {
			ret = fmt.Sprintf("Named%sOf%s", ret, sf.name)
		}

		if isSamePkg(ctx.working, genfp.FromTypesPackage(sf.pack)) {
			return ret
		} else {
			return fmt.Sprintf("%s.%s", r.w.GetImportedName(genfp.FromTypesPackage(sf.pack)), ret)
		}
	} else {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

		return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, valueType)

	}
}

func (r *TypeClassSummonContext) labelledHlistReprType(ctx SummonContext, sf structFunctions) func() string {
	return func() string {
		if sf.fields.Size() == 0 {
			hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

			return fmt.Sprintf(`%s.Nil`, hlistpk)
		}

		hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

		hlisttp := seq.Fold(sf.fields.Reverse(), hlistpk+".Nil", func(b string, f metafp.StructField) string {
			return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, r.structFieldNameTypeStr(ctx, sf, f), b)
		})

		return hlisttp
	}
}
