package main

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) tupleReprType(ctx SummonContext, sf structFunctions, tptypeOpt fp.Option[metafp.TypeInfo]) func() string {
	return func() string {
		tuplepkid := option.Map(tptypeOpt, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		tuplepk := r.w.GetImportedName(tuplepkid)

		fields := sf.fields

		p := seq.Map(sf.typeArgs, func(f metafp.TypeInfoExpr) string {
			return f.TypeName(r.w, ctx.working)
		}).MakeString(",")

		if sf.typeArgs.Size() == 0 {
			return fmt.Sprintf(`%s.Unit`, tuplepk)
		}

		return fmt.Sprintf("%s.Tuple%d[%s]", tuplepk, fields.Size(), p)
	}
}

func (r *TypeClassSummonContext) intoTupleRepr(ctx SummonContext, sf structFunctions, tptypeOpt fp.Option[metafp.TypeInfo]) func() string {

	return func() string {
		if sf.valueGenerated {
			for tpm := range sf.tpe.Method.Get("AsTuple").All() {
				rt := tpm.Signature().Results()

				possible := func() bool {
					if tptypeOpt.IsDefined() {
						rett := metafp.GetTypeInfo(rt.At(0).Type())
						if rt.Len() == 1 && rett.PkgName() == tptypeOpt.Get().PkgName() {
							return true
						}
						return false
					}
					return true
				}()
				if possible {
					return fmt.Sprintf("%s.AsTuple", sf.typeStr(ctx.working))
				}
			}
		}
		type fieldName = fp.Entry[string]

		p := seq.Map(sf.typeArgs, func(f metafp.TypeInfoExpr) string {
			return f.TypeName(r.w, ctx.working)
		}).MakeString(",")

		tppk := r.w.GetImportedName(option.Map(tptypeOpt, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp", "fp")))
		args := func() string {
			for tptype := range tptypeOpt.All() {
				if tptype.Fields().Size() == sf.fields.Size() {
					return seq.Map(seq.ZipWithIndex(sf.names), func(v fp.WithIndex[fieldName]) string {
						fi := tptype.Fields().Get(v.I1).Get()
						return fmt.Sprintf("%s : v.%s", fi.Name, v.I2.I1)
					}).MakeString(",\n")
				} else {
					return seq.Map(seq.ZipWithIndex(sf.names), func(v fp.WithIndex[fieldName]) string {
						return fmt.Sprintf("v.%s", v.I2.I1)
					}).MakeString(",\n")
				}
			}
			return seq.Map(seq.ZipWithIndex(sf.names), func(v fp.WithIndex[fieldName]) string {
				return fmt.Sprintf("I%d: v.%s", v.I1+1, v.I2.I1)
			}).MakeString(",\n")
		}()

		return fmt.Sprintf(`func( v %s) %s.Tuple%d[%s] {
				return %s.Tuple%d[%s]{
					%s,
				}
			}`, sf.typeStr(ctx.working), tppk, sf.fields.Size(), p,
			tppk, sf.fields.Size(), p, args,
		)

	}
}

func (r TypeClassSummonContext) fromTupleRepr(ctx SummonContext, sf structFunctions, tptypeOpt fp.Option[metafp.TypeInfo]) func() string {
	return func() string {
		if sf.valueGenerated {
			for tpm := range sf.tpe.Method.Get("AsTuple").All() {
				rt := tpm.Signature().Results()

				possible := func() bool {
					if tptypeOpt.IsDefined() {
						rett := metafp.GetTypeInfo(rt.At(0).Type())
						if rt.Len() == 1 && rett.PkgName() == tptypeOpt.Get().PkgName() {
							return true
						}
						return false
					} else {
						return true
					}
				}()
				if possible {
					fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
					aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

					builderreceiver := sf.builderTypeStr(ctx.working)
					return fmt.Sprintf(`%s.Compose(
							%s.Curried2(%s.FromTuple)(%s{}),
							%s.Build,
							)`,
						fppk,
						aspk, builderreceiver, builderreceiver, builderreceiver,
					)
				}
			}
		}

		p := seq.Map(sf.typeArgs, func(f metafp.TypeInfoExpr) string {
			return f.TypeName(r.w, ctx.working)
		}).MakeString(",")

		tppk := r.w.GetImportedName(option.Map(tptypeOpt, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp", "fp")))

		assign := func() string {
			if tptypeOpt.IsDefined() && tptypeOpt.Get().Fields().Size() == sf.fields.Size() {
				tptype := tptypeOpt.Get()
				return seq.Map(seq.ZipWithIndex(sf.names), func(v fp.WithIndex[fp.Entry[string]]) string {
					fi := tptype.Fields().Get(v.I1).Get()
					return fmt.Sprintf("%s : t.%s", v.I2.I1, fi.Name)
				}).MakeString(",\n")
			} else {
				return seq.Map(seq.ZipWithIndex(sf.names), func(v fp.WithIndex[fp.Entry[string]]) string {
					return fmt.Sprintf("%s : t.I%d", v.I2.I1, v.I1+1)
				}).MakeString(",\n")
			}
		}()

		valuereceiver := sf.typeStr(ctx.working)
		return fmt.Sprintf(`func(t %s.Tuple%d[%s]) %s {
					return %s{
						%s,
					}
				}`, tppk, sf.fields.Size(), p, valuereceiver,
			valuereceiver,
			assign,
		)
	}
}
