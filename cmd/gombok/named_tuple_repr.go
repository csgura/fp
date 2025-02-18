package main

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) labelledTupleReprType(ctx SummonContext, sf structFunctions) func() string {
	return func() string {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

		names := seq.Map(sf.fields, func(v metafp.StructField) string {
			return r.structFieldNameTypeStr(ctx, sf, v)
		}).MakeString(",")

		return fmt.Sprintf("%s.Labelled%d[%s]", fppk, sf.fields.Size(), names)
	}

}

func (r *TypeClassSummonContext) intoLabelledTupleRepr(ctx SummonContext, sf structFunctions) func() string {
	if sf.namedGenerated {
		return func() string {
			return fmt.Sprintf("%s.AsLabelled", sf.typeStr(ctx.working))
		}
	}

	type fieldName = fp.Entry[string]

	return func() string {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		namedTypeArgs := seq.Zip(sf.names, sf.typeArgs)

		labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
			return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
		}).MakeString(",")

		varlist := iterator.Map(iterator.Range(0, sf.typeArgs.Size()), func(v int) string {
			return fmt.Sprintf("i%d", v)
		}).MakeString(",")

		hlistExpr := seq.Map(seq.ZipWithIndex(namedTypeArgs), func(t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
			idx, t2 := t3.Unapply()
			name, _ := t2.Unapply()
			return fmt.Sprintf(`%s.NamedWithTag("%s", i%d , %s)`, aspk, name.I1, idx, "`"+name.I2+"`")

		}).MakeString(",")

		return fmt.Sprintf(`func(v %s) %s.Labelled%d[%s] {
							%s := %s
							return %s.Labelled%d(%s)
						}`, sf.typeStr(ctx.working), fppk, sf.fields.Size(), labelledtp,
			varlist, sf.unapply("v"),
			aspk, sf.fields.Size(), hlistExpr)
	}

}

func (r TypeClassSummonContext) fromLabelledTupleRepr(ctx SummonContext, sf structFunctions) func() string {
	if sf.namedGenerated {
		return func() string {
			builderreceiver := sf.builderTypeStr(ctx.working)
			fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromLabelled)(%s{}),
					%s.Build,
					)`,
				fppk,
				aspk, builderreceiver, builderreceiver, builderreceiver,
			)
		}
	}

	type fieldName = fp.Entry[string]

	return func() string {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		namedTypeArgs := seq.Zip(sf.names, sf.typeArgs)

		labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
			return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
		}).MakeString(",")

		args := seq.Map(seq.ZipWithIndex(sf.names), func(v fp.Tuple2[int, fieldName]) string {
			return fmt.Sprintf("t.I%d.Value()", v.I1+1)
		})

		return fmt.Sprintf(`func( t %s.Labelled%d[%s] ) %s {
				return %s
			}`, fppk, sf.fields.Size(), labelledtp, sf.typeStr(ctx.working),
			r.structApplyExpr(ctx, sf.tpe.AsNamed(), sf.fields, args...),
		)
	}
}
