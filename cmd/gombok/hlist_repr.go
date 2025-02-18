package main

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) hlistReprType(ctx SummonContext, sf structFunctions, constp fp.Option[metafp.TypeInfo]) func() string {
	return func() string {
		hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

		conspkid := option.Map(constp, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
		conspk := r.w.GetImportedName(conspkid)

		fields := sf.fields

		typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfoExpr {
			return v.TypeInfoExpr(ctx.working)
		})

		if typeArgs.Size() == 0 {
			return fmt.Sprintf(`%s.Nil`, hlistpk)
		}

		hlisttp := seq.Fold(typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfoExpr) string {
			return fmt.Sprintf("%s.Cons[%s,%s]", conspk, a.TypeName(r.w, ctx.working), b)
		})

		return hlisttp
	}
}

func (r *TypeClassSummonContext) toHlistRepr(ctx SummonContext, sf structFunctions, constp fp.Option[metafp.TypeInfo]) func() string {
	return func() string {
		hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

		if sf.typeArgs.Size() == 0 {
			return fmt.Sprintf(`func (%s) %s.Nil {
							return %s.Empty()
						}`, sf.typeStr(ctx.working), hlistpk, hlistpk)
		}

		conspkid := option.Map(constp, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
		conspk := r.w.GetImportedName(conspkid)

		hlisttp := seq.Fold(sf.typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfoExpr) string {
			return fmt.Sprintf("%s.Cons[%s,%s]", conspk, a.TypeName(r.w, ctx.working), b)
		})

		varlist := iterator.Map(iterator.Range(0, sf.typeArgs.Size()), func(v int) string {
			return fmt.Sprintf("i%d", v)
		}).MakeString(",")

		hlast := fmt.Sprintf("h%d := %s.Empty()", sf.typeArgs.Size(), hlistpk)
		hlistExpr := seq.Fold(as.Seq(iterator.Range(0, sf.typeArgs.Size()).ToSeq()).Reverse(), seq.Of(hlast), func(expr fp.Seq[string], v int) fp.Seq[string] {
			return append(expr, fmt.Sprintf(`h%d := %s.Concat(i%d, h%d)`, v, conspk, v, v+1))
		}).MakeString("\n")

		return fmt.Sprintf(`func(v %s) %s {
					%s := %s
					%s
					return h0
				}`, sf.typeStr(ctx.working), hlisttp,
			varlist, sf.unapply("v"),
			hlistExpr)
	}
}

func (r *TypeClassSummonContext) fromHlistRepr(ctx SummonContext, sf structFunctions, constp fp.Option[metafp.TypeInfo]) func() string {
	return func() string {
		hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

		if sf.typeArgs.Size() == 0 {
			valuereceiver := sf.typeStr(ctx.working)
			return fmt.Sprintf(`func (%s.Nil) %s{
							return %s{}
						}`, hlistpk, valuereceiver, valuereceiver)
		}

		conspkid := option.Map(constp, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
		conspk := r.w.GetImportedName(conspkid)

		headfield, tailfield := func() (string, string) {

			if constp.IsDefined() {
				ct := constp.Get()
				if ct.Fields().Size() == 2 {
					if ct.Fields()[0].Public() && ct.Fields()[1].Public() {
						return ct.Fields()[0].Name, ct.Fields()[1].Name
					}
				}
			}
			return "", ""
		}()

		hlisttp := seq.Fold(sf.typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfoExpr) string {
			return fmt.Sprintf("%s.Cons[%s,%s]", conspk, a.TypeName(r.w, ctx.working), b)
		})

		expr := seq.Map(iterator.Range(0, sf.typeArgs.Size()).ToSeq(), func(idx int) string {
			if headfield != "" {
				if idx == sf.typeArgs.Size()-1 {
					return fmt.Sprintf(`i%d := hl%d.%s`, idx, idx, headfield)
				}
				return fmt.Sprintf(`i%d , hl%d := hl%d.%s, hl%d.%s`, idx, idx+1, idx, headfield, idx, tailfield)
			}
			if idx == sf.typeArgs.Size()-1 {
				return fmt.Sprintf(`i%d := %s.Head(hl%d)`, idx, conspk, idx)
			}
			return fmt.Sprintf(`i%d , hl%d := %s.Unapply(hl%d)`, idx, idx+1, conspk, idx)
		}).MakeString("\n")

		arglist := seq.Map(iterator.Range(0, sf.typeArgs.Size()).ToSeq(), func(idx int) string {
			return fmt.Sprintf("i%d", idx)
		})
		return fmt.Sprintf(`func(hl0 %s) %s {
		%s
		return %s
	}`, hlisttp, sf.typeStr(ctx.working),
			expr,
			sf.apply(arglist))
	}
}

func (r *TypeClassSummonContext) summonStructHlistGenericRepr(ctx SummonContext, tc metafp.TypeClass, sf structFunctions, consf string, nilf string, force bool) fp.Option[GenericRepr] {
	shcons := r.lookupTypeClassFunc(ctx, tc, consf)
	if shcons.IsDefined() || force {
		constp := option.FlatMap(shcons, func(tci metafp.TypeClassInstance) fp.Option[metafp.TypeInfo] {
			return tci.Result.TypeArgs.Head()
		})
		return option.Some(GenericRepr{
			Kind:         fp.GenericKindStruct,
			Type:         as.Supplier1(sf.typeStr, ctx.working),
			ReprType:     r.hlistReprType(ctx, sf, constp),
			ToReprExpr:   r.toHlistRepr(ctx, sf, constp),
			FromReprExpr: r.fromHlistRepr(ctx, sf, constp),
			ReprExpr: func() SummonExpr {
				arity := sf.typeArgs.Size()
				hnil := r.lookupTypeClassFunc(ctx, tc, nilf).OrElseGet(as.Supplier2(r.lookupHNilMust, ctx, tc))
				hlist := seq.Fold(sf.typeArgs.Take(arity).Reverse(), newSummonExpr(func() string { return hnil.PackagedName(r.w, ctx.working) }), func(tail SummonExpr, ti metafp.TypeInfoExpr) SummonExpr {

					consname := option.Map(shcons, func(tci metafp.TypeClassInstance) string {
						return tci.PackagedName(r.w, ctx.working)
					}).OrElse("HCons")

					instance := r.summonRequired(ctx, metafp.RequiredInstance{
						TypeClass: ctx.typeClass,
						Type:      ti.Type,
					})
					return newSummonExpr(func() string {
						return fmt.Sprintf(`%s(
								%s,
								%s,
							)`, consname, instance, tail,
						)
					}, instance.paramInstance, tail.paramInstance)
				})
				return hlist
			},
		})
	}
	return option.None[GenericRepr]()
}
