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

	result := r.lookupTypeClassFunc(ctx, tc, fmt.Sprintf("Struct%d", typeArgs.Size()))

	return option.Map(result, func(tm metafp.TypeClassInstance) GenericRepr {
		return GenericRepr{
			Kind:         fp.GenericKindStruct,
			ToReprExpr:   sf.asMinTuple,
			FromReprExpr: sf.fromMinTuple,
			ReprExpr: func() SummonExpr {
				return r.exprTupleWithName(ctx, tc, tm, sf.pack, sf.name, names, typeArgs, sf.namedGenerated)
			},
		}
	}).Or(func() fp.Option[GenericRepr] {
		fnamed := option.TraverseSeq(typeArgs, func(a metafp.TypeInfoExpr) fp.Option[metafp.TypeClassInstance] {
			return r.lookupTypeClassFuncCheckType(ctx, tc, "Named", a.Type)
		})
		if fnamed.IsDefined() {
			panic("all find")
		}
		return option.None[GenericRepr]()

	})
}
