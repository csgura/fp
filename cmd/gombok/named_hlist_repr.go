package main

import (
	"fmt"

	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func (r *TypeClassSummonContext) toLabelledHlistRepr(ctx SummonContext, sf structFunctions, constp fp.Option[metafp.TypeInfo]) func() string {
	type fieldName = fp.Entry[string]

	return func() string {
		if sf.typeArgs.Size() == 0 {
			hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

			return fmt.Sprintf(`func (%s) %s.Nil {
							return %s.Empty()
						}`, sf.typeStr(ctx.working), hlistpk, hlistpk)
			// } else if sf.typeArgs.Size() < max.Product {
			// 	arity := fp.Min(sf.typeArgs.Size(), max.Product-1)
			// 	//arity := typeArgs.Size()

			// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			// 	aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			// 	namedTypeArgs := seq.Zip(sf.names, sf.typeArgs)

			// 	if r.implicitTypeInference {
			// 		return fmt.Sprintf(`%s.Compose(
			// 						%s,
			// 						%s.HList%dLabelled,
			// 					)`, fppk,
			// 			r.intoLabelledTupleRepr(ctx, sf)(),
			// 			aspk, arity,
			// 		)
			// 	} else {
			// 		tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
			// 			return namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, f.I1.I1, sf.namedGenerated, f.I2.TypeName(r.w, ctx.working))
			// 		}).Take(arity).MakeString(",")

			// 		return fmt.Sprintf(`%s.Compose(
			// 						%s,
			// 						%s.HList%dLabelled[%s],
			// 					)`, fppk,
			// 			r.intoLabelledTupleRepr(ctx, sf)(),
			// 			aspk, arity, tp,
			// 		)
			// 	}

		} else {
			nilpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

			conspkid := option.Map(constp, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
			conspk := r.w.GetImportedName(conspkid)

			namedTypeArgs := seq.Zip(sf.names, sf.typeArgs)

			hlisttp := seq.Fold(namedTypeArgs.Reverse(), nilpk+".Nil", func(b string, f fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
				name, a := f.Unapply()
				return fmt.Sprintf("%s.Cons[%s,%s]", conspk, r.namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, name.I1, sf.namedGenerated, a), b)
			})

			varlist := iterator.Map(iterator.Range(0, sf.typeArgs.Size()), func(v int) string {
				return fmt.Sprintf("i%d", v)
			}).MakeString(",")

			hlistExpr := option.Map(option.Of(sf.namedGenerated).Filter(eq.GivenValue(true)), func(bool) string {
				return seq.Fold(seq.ZipWithIndex(namedTypeArgs).Reverse(), conspk+".Empty()", func(expr string, t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
					idx, t2 := t3.Unapply()
					name, tp := t2.Unapply()
					return fmt.Sprintf(`%s.Concat(%s{i%d}, 
									%s,
								)`, conspk, r.namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, name.I1, sf.namedGenerated, tp), idx, expr)
				})
			}).OrElseGet(func() string {
				aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

				return seq.Fold(seq.ZipWithIndex(namedTypeArgs).Reverse(), conspk+".Empty()", func(expr string, t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
					idx, t2 := t3.Unapply()
					name, _ := t2.Unapply()
					return fmt.Sprintf(`%s.Concat(%s.NamedWithTag("%s", i%d, %s), 
									%s,
								)`, conspk, aspk, name.I1, idx, "`"+name.I2+"`", expr)
				})
			})
			//hlistExpr :=

			unapplyexpr := sf.unapply("v")
			return fmt.Sprintf(`func(v %s) %s {
							%s := %s
							return %s
						}`, sf.typeStr(ctx.working), hlisttp,
				varlist, unapplyexpr,
				hlistExpr)
		}
	}
}

func (r *TypeClassSummonContext) fromLabelledHlistRepr(ctx SummonContext, sf structFunctions, constp fp.Option[metafp.TypeInfo]) func() string {
	return func() string {
		hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

		conspkid := option.Map(constp, metafp.TypeInfo.PkgId).OrElse(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
		conspk := r.w.GetImportedName(conspkid)

		if sf.typeArgs.Size() == 0 {
			valuereceiver := sf.typeStr(ctx.working)
			return fmt.Sprintf(`func (%s.Nil) %s{
							return %s{}
						}`, hlistpk, valuereceiver, valuereceiver)
		}

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
