package main

import (
	"fmt"
	"go/types"
	"os"
	"reflect"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"

	"golang.org/x/tools/go/packages"
)

func isSamePkg(p1 genfp.WorkingPackage, p2 genfp.PackageId) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}

	return p1.Path() == p2.Path()
}

func namedName(w genfp.Writer, working genfp.WorkingPackage, typePkg genfp.WorkingPackage, name string) string {

	ret := publicName(name)
	if ret == name {
		ret = fmt.Sprintf("PubNamed%s", ret)
	} else {
		ret = fmt.Sprintf("Named%s", ret)
	}

	if isSamePkg(working, typePkg) {
		return ret
	} else {
		return fmt.Sprintf("%s.%s", w.GetImportedName(typePkg), ret)
	}

}

func publicName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func privateName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func isTypeDefined(pk genfp.WorkingPackage, name string) bool {
	return pk.Scope().Lookup(name) != nil
}

func isMethodDefined(pk genfp.WorkingPackage, tpeName string, method string) bool {
	obj := pk.Scope().Lookup(tpeName)
	if obj == nil {
		return false
	}

	if atp, ok := obj.Type().(*types.Named); ok {
		for i := 0; i < atp.NumMethods(); i++ {
			if atp.Method(i).Name() == method {
				return true
			}
		}
	}
	return false

}

func iterate[T, R any](len int, getter func(idx int) T, fn func(int, T) R) fp.Seq[R] {
	ret := []R{}
	for i := 0; i < len; i++ {
		ret = append(ret, fn(i, getter(i)))
	}
	return ret
}

func processDeref(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {

	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	if _, ok := ts.Tags.Get("@fp.Deref").Unapply(); ok {
		//fmt.Printf("rhs type = %s\n", ts.RhsType)

		if rhs, ok := ts.RhsType.Unapply(); ok {

			//valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
			valuetp := ts.Info.TypeParamIns(w, workingPackage)
			typetp := ts.Info.TypeParamDecl(w, workingPackage)

			valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

			unwrapFunc := "Deref"
			if ts.Info.Method.Get(unwrapFunc).IsEmpty() {

				rhsTypeName := w.TypeName(workingPackage, rhs.Type)

				fmt.Fprintf(w, `
					func (r %s) %s() %s {
						return %s(r)
					}
				`, valuereceiver, unwrapFunc, rhsTypeName, rhsTypeName)
				genMethod = genMethod.Incl(unwrapFunc)
			}

			castFunc := fmt.Sprintf("Into%s%s", ts.Name, typetp)
			if workingPackage.Scope().Lookup(castFunc) == nil {
				rhsTypeName := w.TypeName(workingPackage, rhs.Type)

				fmt.Fprintf(w, `
					func %s(v %s) %s {
						return %s%s(v)
					}
				`, castFunc, rhsTypeName, valuereceiver,
					ts.Name, valuetp,
				)
			}

			sorted := seq.Sort(rhs.Method.Iterator().ToSeq(), ord.ContraMap(ord.Given[string](), fp.Tuple2[string, *types.Func].Head))

			sorted.Foreach(func(t fp.Tuple2[string, *types.Func]) {
				name, f := t.Unapply()
				if ts.Info.Method.Get(name).IsEmpty() && !genMethod.Contains(name) {

					if sig, ok := f.Type().(*types.Signature); ok {

						rhsTypeName := w.TypeName(workingPackage, rhs.Type)

						isPtrReceiver := false
						if sig.Recv() != nil && sig.Recv().Type() != nil {
							if _, ok := sig.Recv().Type().(*types.Pointer); ok {
								isPtrReceiver = true
							}
						}
						argTypeStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
							return fmt.Sprintf("%s %s", t.Name(), w.TypeName(workingPackage, t.Type()))
						}).MakeString(",")

						argstr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
							return t.Name()
						}).MakeString(",")

						resstr := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
							return w.TypeName(workingPackage, t.Type())
						}).MakeString(",")

						if resstr != "" {
							if isPtrReceiver {
								fmt.Fprintf(w, `
									func (r *%s) %s(%s) (%s) {
										return (*%s)(r).%s(%s)
									}
								`, valuereceiver, name, argTypeStr, resstr, rhsTypeName, name, argstr)
							} else {
								fmt.Fprintf(w, `
									func (r %s) %s(%s) (%s) {
										return %s(r).%s(%s)
									}
								`, valuereceiver, name, argTypeStr, resstr, rhsTypeName, name, argstr)
							}
						} else {
							if isPtrReceiver {

								fmt.Fprintf(w, `
									func (r *%s) %s(%s) {
										(*%s)(r).%s(%s)
									}
								`, valuereceiver, name, argTypeStr, rhsTypeName, name, argstr)
							} else {
								fmt.Fprintf(w, `
									func (r %s) %s(%s) {
										%s(r).%s(%s)
									}
								`, valuereceiver, name, argTypeStr, rhsTypeName, name, argstr)
							}
						}
						genMethod = genMethod.Incl(name)

					}
				}

			})
		}
	}
	return genMethod
}

func processGetter(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage
	if _, ok := ts.Tags.Get("@fp.Getter").Unapply(); ok {
		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

		genMethod = genPrivateGetters(ctx, privateFields, genMethod)
	}

	if anno, ok := ts.Tags.Get("@fp.GetterPubField").Unapply(); ok {

		publicFields := ts.Fields.Filter(metafp.StructField.Public)

		if publicFields.Size() == 0 {
			return genMethod
		}

		//valuetpdec := ""

		//valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
		valuetp := ts.Info.TypeParamIns(w, workingPackage)

		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

		publicFields.Foreach(func(f metafp.StructField) {

			getterName := fmt.Sprintf("Get%s", f.Name)

			if ts.Info.Method.Get(getterName).IsEmpty() && !genMethod.Contains(getterName) {
				if anno.Params().Get("override").OrElse("false") != "true" && ts.Tags.Contains("@fp.Deref") && ts.RhsType.IsDefined() {
					if ts.RhsType.Get().Method.Contains(getterName) {
						return
					}
				}

				ftp := f.TypeName(w, workingPackage)

				fmt.Fprintf(w, `
						func (r %s) %s() %s {
							return r.%s
						}
					`, valuereceiver, getterName, ftp, f.Name)

				genMethod = genMethod.Incl(getterName)
			}
		})
	}

	return genMethod
}

type ParsedTag struct {
	NotLabeled []string
	Labels     fp.Map[string, string]
}

func parseGombokTag(tags reflect.StructTag) ParsedTag {
	ret := ParsedTag{}
	tag, ok := tags.Lookup("fp")
	if ok {
		values := strings.Split(tag, ",")
		for _, v := range values {
			attrs := strings.Split(v, ";")
			for _, a := range attrs {
				kv := strings.Split(a, "=")
				if len(kv) == 2 {

					ret.Labels = ret.Labels.Updated(kv[0], kv[1])
				} else {
					ret.NotLabeled = append(ret.NotLabeled, kv[0])
				}
			}
		}
	}
	return ret
}

func genStringMethod(ctx TaggedStructContext, allFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {

	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage
	tccache := ctx.summCtx.tcCache
	derives := ctx.derives

	if ts.Info.Method.Get("String").IsEmpty() && !genMethod.Contains("String") {

		valuereceiver := ts.Info.TypeStr(w, workingPackage)

		useShow := option.Map(ts.Tags.Get("@fp.String"), func(v metafp.Annotation) bool {
			return v.Params().Get("useShow").OrElse("false") == "true"
		}).OrElse(false)

		if useShow {

			tc := metafp.TypeClass{
				Name:    "Show",
				Package: genfp.NewImportPackage("github.com/csgura/fp", "fp"),
			}

			// ctx.summCtx.summon(ctx CurrentContext, req metafp.RequiredInstance)
			scope := tccache.GetLocal(workingPackage.Package(), tc)

			showDerive := derives.Find(func(v metafp.TypeClassDerive) bool {
				return v.TypeClass.Name == "Show" && v.TypeClass.Package.Path() == "github.com/csgura/fp"
			})

			insOpt := option.FlatMap(showDerive, func(v metafp.TypeClassDerive) fp.Option[metafp.TypeClassInstance] {
				//fmt.Printf("find by name %s\n", v.GeneratedInstanceName())
				return scope.FindByName(v.GeneratedInstanceName(), ts.Info)
			}).Or(func() fp.Option[metafp.TypeClassInstance] {
				return scope.Find(ts.Info).Head()
			})
			if insOpt.IsDefined() {
				ins := insOpt.Get()
				//fmt.Printf("insOpt isDefined, tc= %s, static %t, required %d\n", ins.TypeClassType, ins.Static, ins.RequiredInstance.Size())

				if ins.Static {
					fmt.Fprintf(w, `
						func(r %s) String() string {
							return %s.Show(r)
						}
					`, valuereceiver,
						ins.Name,
					)

					genMethod = genMethod.Incl("String")

					return genMethod
				} else if ins.RequiredInstance.Size() == 0 {

					valuetp := ""
					if !ins.TypeParam.IsEmpty() {
						valuetp = "[" + seq.Map(ins.TypeParam, func(v metafp.TypeParam) string {
							return option.Map(ins.ParamMapping.Get(v.Name), func(v metafp.TypeInfo) string {
								return w.TypeName(workingPackage, v.Type)
							}).OrElse(v.Name)
						}).MakeString(",") + "]"
					}

					fmt.Fprintf(w, `
						func(r %s) String() string {
							return %s%s().Show(r)
						}
					`, valuereceiver,
						ins.Name, valuetp,
					)

					genMethod = genMethod.Incl("String")

					return genMethod
				}
			}

			valuetp := ts.Info.TypeParamIns(w, workingPackage)

			if showDerive.IsDefined() && valuetp == "" {

				fmt.Fprintf(w, `
						func(r %s) String() string {
							return %s().Show(r)
						}
					`, valuereceiver,
					showDerive.Get().GeneratedInstanceName(),
				)

				genMethod = genMethod.Incl("String")

				return genMethod
			}
		}

		fmtalias := w.GetImportedName(genfp.NewImportPackage("fmt", "fmt"))

		printable := allFields.Filter(func(v metafp.StructField) bool {
			return v.FieldType.IsPrintable()
		}).FilterNot(func(v metafp.StructField) bool {
			t := parseGombokTag(reflect.StructTag(v.Tag))
			return as.Seq(t.NotLabeled).Exists(eq.GivenValue("String.Exclude"))
		})

		fm := iterator.Map(iterator.FromSeq(printable), func(f metafp.StructField) string {
			return fmt.Sprintf("%s:%%v", f.Name)
		}).MakeString(", ")

		fields := iterator.Map(iterator.FromSeq(printable), func(f metafp.StructField) string {
			return fmt.Sprintf("r.%s", f.Name)
		}).MakeString(",")

		fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s.%s{%s}", %s)
					}
				`, valuereceiver,
			fmtalias, ts.Package.Name(), ts.Name, fm, fields,
		)
		genMethod = genMethod.Incl("String")

	}

	return genMethod
}

func genUnapply(ctx TaggedStructContext, allFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	valuereceiver := ts.Info.TypeStr(w, workingPackage)
	if allFields.Size() < max.Product {

		if ts.Info.Method.Get("AsTuple").IsEmpty() && !genMethod.Contains("AsTuple") {
			asalias := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			fppkg := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

			arity := fp.Min(allFields.Size(), max.Product-1)
			tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
				return v.TypeName(w, workingPackage)
			}).MakeString(",")

			fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Take(arity).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsTuple() %s.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, fppkg, arity, tp,
				asalias, arity, fields,
			)
			genMethod = genMethod.Incl("AsTuple")

		}
	}

	if ts.Info.Method.Get("Unapply").IsEmpty() && !genMethod.Contains("Unapply") {

		tp := iterator.Map(iterator.FromSeq(allFields), func(v metafp.StructField) string {
			return v.TypeName(w, workingPackage)
		}).MakeString(",")

		fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
			return fmt.Sprintf("r.%s", f.Name)
		}).MakeString(",")

		fmt.Fprintf(w, `
					func(r %s) Unapply() (%s) {
						return %s
					}

				`, valuereceiver, tp,
			fields,
		)
		genMethod = genMethod.Incl("Unapply")

	}
	return genMethod
}
func processString(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts

	if _, ok := ts.Tags.Get("@fp.String").Unapply(); ok {
		allFields := applyFields(ts)

		genMethod = genStringMethod(ctx, allFields, genMethod)
	}

	return genMethod
}

func applyFields(ts metafp.TaggedStruct) fp.Seq[metafp.StructField] {
	return ts.Fields.FilterNot(func(v metafp.StructField) bool {
		// field 가 아무 것도 없는 embedded struct 는 생성에서 제외
		return strings.HasPrefix(v.Name, "_") || (v.Embedded && v.FieldType.Underlying().IsStruct() && v.FieldType.Fields().Size() == 0)
	})
}

func genBuilder(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	valuetpdec := ts.Info.TypeParamDecl(w, workingPackage)
	valuetp := ts.Info.TypeParamIns(w, workingPackage)

	valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

	builderTypeName := fmt.Sprintf("%sBuilder", ts.Name)
	builderType := builderTypeName + valuetpdec
	builderreceiver := fmt.Sprintf("%sBuilder%s", ts.Name, valuetp)

	privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

	allFields := applyFields(ts)

	if !isTypeDefined(workingPackage, builderTypeName) {
		fmt.Fprintf(w, `
					type %s %s
				`, builderType, valuereceiver)
	}

	if !isMethodDefined(workingPackage, builderTypeName, "Build") {
		fmt.Fprintf(w, `
				func(r %s) Build() %s {
					return %s(r)
				}
			`, builderreceiver, valuereceiver, valuereceiver)
	}

	if ts.Info.Method.Get("Builder").IsEmpty() {

		fmt.Fprintf(w, `
				func(r %s) Builder() %s {
					return %s(r)
				}
			`, valuereceiver, builderreceiver, builderreceiver)
		genMethod = genMethod.Incl("Builder")
	}

	privateFields.Foreach(func(f metafp.StructField) {

		uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]

		if !isMethodDefined(workingPackage, builderTypeName, uname) {
			ftp := f.TypeName(w, workingPackage)

			fmt.Fprintf(w, `
						func (r %s) %s( v %s) %s {
							r.%s = v
							return r
						}
					`, builderreceiver, uname, ftp, builderreceiver, f.Name)
		}

		if f.FieldType.IsOption() {
			optiont := w.TypeName(workingPackage, f.FieldType.TypeArgs.Head().Get().Type)
			optionpk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/option", "option"))

			if !isMethodDefined(workingPackage, builderTypeName, "Some"+uname) {

				fmt.Fprintf(w, `
						func (r %s) Some%s(v %s) %s {
							r.%s = %s.Some(v)
							return r
						}
					`, builderreceiver, uname, optiont, builderreceiver,
					f.Name, optionpk)
			}
			if !isMethodDefined(workingPackage, builderTypeName, "None"+uname) {

				fmt.Fprintf(w, `
						func (r %s) None%s() %s {
							r.%s = %s.None[%s]()
							return r
						}
					`, builderreceiver, uname, builderreceiver,
					f.Name, optionpk, optiont)
			}
		}

	})

	if allFields.Size() < max.Product {

		if !isMethodDefined(workingPackage, builderTypeName, "FromTuple") {
			fppkg := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

			arity := fp.Min(allFields.Size(), max.Product-1)

			tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
				return v.TypeName(w, workingPackage)
			}).MakeString(",")

			fields := iterator.Map(iterator.Zip(iterator.Range(0, allFields.Size()), seq.Iterator(allFields)), func(f fp.Tuple2[int, metafp.StructField]) string {
				return fmt.Sprintf("r.%s = t.I%d", f.I2.Name, f.I1+1)
			}).Take(arity).MakeString("\n")

			fmt.Fprintf(w, `
					func (r %s) FromTuple(t %s.Tuple%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, fppkg, arity, tp, builderreceiver,
				fields,
			)
		}
	}

	if !isMethodDefined(workingPackage, builderTypeName, "Apply") {

		tp := iterator.Map(seq.Iterator(allFields), func(v metafp.StructField) string {
			return fmt.Sprintf("%s %s", v.Name, v.TypeName(w, workingPackage))
		}).MakeString(",")

		fields := iterator.Map(iterator.Zip(iterator.Range(0, allFields.Size()), seq.Iterator(allFields)), func(f fp.Tuple2[int, metafp.StructField]) string {
			return fmt.Sprintf("r.%s = %s", f.I2.Name, f.I2.Name)
		}).MakeString("\n")

		fmt.Fprintf(w, `
					func (r %s) Apply( %s ) %s {
						%s
						return r
					}
				`, builderreceiver, tp, builderreceiver,
			fields,
		)
	}

	if !isMethodDefined(workingPackage, builderTypeName, "FromMap") {

		fields := iterator.Map(seq.Iterator(allFields), func(f metafp.StructField) string {
			if f.FieldType.IsOption() {
				optionpk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/option", "option"))

				return fmt.Sprintf(`if v , ok := m["%s"].(%s); ok {
							r.%s = v
						} else if v, ok := m["%s"].(%s); ok {
							r.%s = %s.Some(v)
						}
						`, f.Name, f.TypeName(w, workingPackage),
					f.Name,
					f.Name, w.TypeName(workingPackage, f.FieldType.TypeArgs.Head().Get().Type),
					f.Name, optionpk,
				)
			} else {
				return fmt.Sprintf(`if v , ok := m["%s"].(%s); ok {
							r.%s = v
						}
							`, f.Name, f.TypeName(w, workingPackage),
					f.Name,
				)
			}
		}).MakeString("\n")

		fmt.Fprintf(w, `
					func(r %s) FromMap(m map[string]any) %s {

						%s
						
						return r
					}

				`, builderreceiver, builderreceiver,
			fields,
		)
	}

	if allFields.Size() < max.Product {
		if ts.Tags.Contains("@fp.GenLabelled") {

			if !isMethodDefined(workingPackage, builderTypeName, "FromLabelled") {

				arity := fp.Min(allFields.Size(), max.Product-1)

				fppkg := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

				tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
					return fmt.Sprintf("%s[%s]", namedName(w, workingPackage, workingPackage, v.Name), v.TypeName(w, workingPackage))
				}).MakeString(",")

				fields := iterator.Map(iterator.Zip(iterator.Range(0, allFields.Size()), seq.Iterator(allFields)), func(f fp.Tuple2[int, metafp.StructField]) string {
					return fmt.Sprintf("r.%s = t.I%d.Value()", f.I2.Name, f.I1+1)
				}).Take(arity).MakeString("\n")

				fmt.Fprintf(w, `
					func (r %s) FromLabelled(t %s.Labelled%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, fppkg, arity, tp, builderreceiver,
					fields,
				)
			}
		}
	}

	return genMethod
}

func processBuilder(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts

	if _, ok := ts.Tags.Get("@fp.Builder").Unapply(); ok {

		genMethod = genBuilder(ctx, genMethod)
	}

	return genMethod
}

func genPrivateWiths(ctx TaggedStructContext, privateFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	valuereceiver := ts.Info.TypeStr(w, workingPackage)

	privateFields.Foreach(func(f metafp.StructField) {

		uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]

		fnName := "With" + uname

		if ts.Info.Method.Get(fnName).IsEmpty() && !genMethod.Contains(fnName) {
			ftp := f.TypeName(w, workingPackage)

			fmt.Fprintf(w, `
						func (r %s) %s(v %s) %s {
							r.%s = v
							return r
						}
					`, valuereceiver, fnName, ftp, valuereceiver, f.Name)
			genMethod = genMethod.Incl(fnName)

		}

		if f.FieldType.IsOption() {
			optiont := w.TypeName(workingPackage, f.FieldType.TypeArgs.Head().Get().Type)
			optionpk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/option", "option"))

			fnName := "WithSome" + uname
			if ts.Info.Method.Get(fnName).IsEmpty() && !genMethod.Contains(fnName) {

				fmt.Fprintf(w, `
							func (r %s) %s(v %s) %s {
								r.%s = %s.Some(v)
								return r
							}
						`, valuereceiver, fnName, optiont, valuereceiver, f.Name, optionpk)
				genMethod = genMethod.Incl(fnName)

			}

			fnName = "WithNone" + uname
			if ts.Info.Method.Get(fnName).IsEmpty() && !genMethod.Contains(fnName) {

				fmt.Fprintf(w, `
							func (r %s) %s() %s {
								r.%s = %s.None[%s]()
								return r
							}
						`, valuereceiver, fnName, valuereceiver, f.Name, optionpk, optiont)
				genMethod = genMethod.Incl(fnName)

			}
		}

	})
	return genMethod
}

func processWith(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	if _, ok := ts.Tags.Get("@fp.With").Unapply(); ok {
		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

		genMethod = genPrivateWiths(ctx, privateFields, genMethod)
	}

	if anno, ok := ts.Tags.Get("@fp.WithPubField").Unapply(); ok {

		publicFields := ts.Fields.Filter(metafp.StructField.Public)

		if publicFields.Size() == 0 {
			return genMethod
		}

		//valuetpdec := ""

		//valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
		valuetp := ts.Info.TypeParamIns(w, workingPackage)

		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

		publicFields.Foreach(func(f metafp.StructField) {

			funcName := fmt.Sprintf("With%s", f.Name)

			if ts.Info.Method.Get(funcName).IsEmpty() && !genMethod.Contains(funcName) {
				if anno.Params().Get("override").OrElse("false") != "true" && ts.Tags.Contains("@fp.Deref") && ts.RhsType.IsDefined() {
					if ts.RhsType.Get().Method.Contains(funcName) {
						return
					}
				}

				ftp := f.TypeName(w, workingPackage)

				fmt.Fprintf(w, `
						func (r %s) %s(v %s) %s {
							r.%s = v
							return r
						}
					`, valuereceiver, funcName, ftp, valuereceiver,
					f.Name,
				)

				genMethod = genMethod.Incl(funcName)
			}
		})
	}
	return genMethod
}

func genAllArgsCons(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {

	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	fnName := fmt.Sprintf("New%s", ts.Name)

	if !genMethod.Contains(fnName) {

		allFields := applyFields(ts)

		tp := iterator.Map(seq.Iterator(allFields), func(v metafp.StructField) string {
			return fmt.Sprintf("%s %s", v.Name, v.TypeName(w, workingPackage))
		}).MakeString(",")

		fields := iterator.Map(iterator.Zip(iterator.Range(0, allFields.Size()), seq.Iterator(allFields)), func(f fp.Tuple2[int, metafp.StructField]) string {
			return fmt.Sprintf("%s : %s", f.I2.Name, f.I2.Name)
		}).MakeString(",\n")

		valueType := ts.Info.TypeStr(w, workingPackage)
		fmt.Fprintf(w, `
			func %s%s(%s) %s {
				return %s {
					%s,
				}
			}
		`, fnName, ts.Info.TypeParamDecl(w, workingPackage), tp, valueType,
			valueType,
			fields,
		)
	}

	return genMethod
}

func genRequiredArgsCons(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	fnName := fmt.Sprintf("New%s", ts.Name)

	if !genMethod.Contains(fnName) {

		allFields := applyFields(ts).FilterNot(func(v metafp.StructField) bool {
			return v.FieldType.IsPtr() || v.FieldType.IsOption()
		})

		tp := iterator.Map(seq.Iterator(allFields), func(v metafp.StructField) string {
			return fmt.Sprintf("%s %s", v.Name, v.TypeName(w, workingPackage))
		}).MakeString(",")

		fields := iterator.Map(iterator.Zip(iterator.Range(0, allFields.Size()), seq.Iterator(allFields)), func(f fp.Tuple2[int, metafp.StructField]) string {
			return fmt.Sprintf("%s : %s", f.I2.Name, f.I2.Name)
		}).MakeString(",\n")

		valueType := ts.Info.TypeStr(w, workingPackage)
		fmt.Fprintf(w, `
			func %s%s(%s) %s {
				return %s {
					%s,
				}
			}
		`, fnName, ts.Info.TypeParamDecl(w, workingPackage), tp, valueType,
			valueType,
			fields,
		)
	}

	return genMethod
}

func processAllArgsCons(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts

	if _, ok := ts.Tags.Get("@fp.AllArgsConstructor").Unapply(); ok {
		genMethod = genAllArgsCons(ctx, genMethod)
	} else if _, ok := ts.Tags.Get("@fp.RequiredArgsConstructor").Unapply(); ok {
		genMethod = genRequiredArgsCons(ctx, genMethod)
	}
	return genMethod
}

type TaggedStructContext struct {
	w              genfp.Writer
	workingPackage genfp.WorkingPackage
	ts             metafp.TaggedStruct
	summCtx        *TypeClassSummonContext
	derives        fp.Seq[metafp.TypeClassDerive]
}

func genTaggedStruct(w genfp.Writer, workingPackage genfp.WorkingPackage, st fp.Seq[metafp.TaggedStruct], summonCtx *TypeClassSummonContext) {
	keyTags := mutable.EmptySet[string]()

	st.Foreach(func(ts metafp.TaggedStruct) {
		genMethod := fp.Set[string]{}
		stDerives := summonCtx.recursiveGen.Filter(func(v metafp.TypeClassDerive) bool {
			return v.DeriveFor.Name == ts.Name
		})

		ctx := TaggedStructContext{
			w:              w,
			workingPackage: workingPackage,
			ts:             ts,
			derives:        stDerives,
			summCtx:        summonCtx,
		}
		genMethod = processAllArgsCons(ctx, genMethod)

		genMethod, keyTags = processValue(ctx, genMethod, keyTags)

		genMethod = processGetter(ctx, genMethod)
		genMethod = processWith(ctx, genMethod)
		genMethod = processDeref(ctx, genMethod)
		genMethod = processString(ctx, genMethod)

		//lint:ignore SA4006 for future
		genMethod = processBuilder(ctx, genMethod)

	})

	klist := keyTags.Iterator().ToSeq()
	seq.Sort(klist, ord.Given[string]()).Foreach(func(name string) {
		fmt.Fprintf(w, `type %s[T any] fp.Tuple2[T,string]
			`, namedName(w, workingPackage, workingPackage, name))

		fmt.Fprintf(w, `func (r %s[T]) Name() string {
				return "%s"
			}
			`, namedName(w, workingPackage, workingPackage, name), name)

		fmt.Fprintf(w, `func (r %s[T]) Value() T {
				return r.I1
			}
			`, namedName(w, workingPackage, workingPackage, name))

		fmt.Fprintf(w, `func (r %s[T]) Tag() string {
				return r.I2
			}
			`, namedName(w, workingPackage, workingPackage, name))

		fmt.Fprintf(w, `func (r %s[T]) WithValue(v T) %s[T] {
				r.I1 = v
				return r
			}
			`, namedName(w, workingPackage, workingPackage, name), namedName(w, workingPackage, workingPackage, name))

		fmt.Fprintf(w, `func (r %s[T]) WithTag(v string) %s[T] {
				r.I2 = v
				return r
			}
			`, namedName(w, workingPackage, workingPackage, name), namedName(w, workingPackage, workingPackage, name))
	})
}

func genPrivateGetters(ctx TaggedStructContext, privateFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	valuetp := ts.Info.TypeParamIns(w, workingPackage)

	valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

	privateFields.Foreach(func(f metafp.StructField) {
		uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]

		fnName := uname
		if ts.Info.Method.Get(fnName).IsEmpty() && !genMethod.Contains(fnName) {
			ftp := f.TypeName(w, workingPackage)

			fmt.Fprintf(w, `
						func (r %s) %s() %s {
							return r.%s
						}
					`, valuereceiver, fnName, ftp, f.Name)
			genMethod = genMethod.Incl(fnName)

		}
	})
	return genMethod
}

//lint:ignore U1000 for future
func genTypeClassMethod(ctx TaggedStructContext, derives fp.Seq[metafp.TypeClassDerive], genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	valuetp := ts.Info.TypeParamIns(w, workingPackage)

	valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

	showDerive := derives.Find(func(v metafp.TypeClassDerive) bool {
		return v.TypeClass.Name == "Show" && v.TypeClass.Package.Path() == "github.com/csgura/fp"
	})

	if ts.Info.Method.Get("String").IsEmpty() && !genMethod.Contains("String") {
		if showDerive.IsDefined() && valuetp == "" {
			if showDerive.Get().IsRecursive() {
				fmt.Fprintf(w, `
				func(r %s) String() string {
					return %s().Show(r)
				}
			`, valuereceiver,
					showDerive.Get().GeneratedInstanceName(),
				)
			} else {
				fmt.Fprintf(w, `
				func(r %s) String() string {
					return %s.Show(r)
				}
			`, valuereceiver,
					showDerive.Get().GeneratedInstanceName(),
				)
			}
			genMethod = genMethod.Incl("String")
		}
	}

	if showDerive.IsDefined() && ts.Info.Method.Get("ShowIndent").IsEmpty() && valuetp == "" {
		fppkg := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

		if showDerive.Get().IsRecursive() {
			fmt.Fprintf(w, `
				func(r %s) ShowIndent(opt %s.ShowOption) string {
					return %s().ShowIndent(r, opt)
				}
			`, valuereceiver, fppkg,
				showDerive.Get().GeneratedInstanceName(),
			)
		} else {
			fmt.Fprintf(w, `
				func(r %s) ShowIndent(opt %s.ShowOption) string {
					return %s.ShowIndent(r, opt)
				}
			`, valuereceiver, fppkg,
				showDerive.Get().GeneratedInstanceName(),
			)
		}

		genMethod = genMethod.Incl("ShowIndent")

	}

	eqDerive := derives.Find(func(v metafp.TypeClassDerive) bool {
		return v.TypeClass.Name == "Eq" && v.TypeClass.Package.Path() == "github.com/csgura/fp"
	})

	if eqDerive.IsDefined() && ts.Info.Method.Get("Eqv").IsEmpty() && valuetp == "" {

		if eqDerive.Get().IsRecursive() {
			fmt.Fprintf(w, `
				func(r %s) Eqv(other %s) bool {
					return %s().Eqv(r, other)
				}
			`, valuereceiver, valuereceiver,
				eqDerive.Get().GeneratedInstanceName(),
			)
		} else {
			fmt.Fprintf(w, `
				func(r %s) Eqv(other %s) bool {
					return %s.Eqv(r, other)
				}
			`, valuereceiver, valuereceiver,
				eqDerive.Get().GeneratedInstanceName(),
			)
		}
		genMethod = genMethod.Incl("Eqv")

	}

	hashDerive := derives.Find(func(v metafp.TypeClassDerive) bool {
		return v.TypeClass.Name == "Hashable" && v.TypeClass.Package.Path() == "github.com/csgura/fp"
	})
	if hashDerive.IsDefined() && ts.Info.Method.Get("Hash").IsEmpty() {

		if hashDerive.Get().IsRecursive() {
			fmt.Fprintf(w, `
				func(r %s) Hash() uint32 {
					return %s().Hash(r)
				}
			`, valuereceiver,
				hashDerive.Get().GeneratedInstanceName(),
			)
		} else {
			fmt.Fprintf(w, `
				func(r %s) Hash() uint32 {
					return %s.Hash(r)
				}
			`, valuereceiver,
				hashDerive.Get().GeneratedInstanceName(),
			)
		}
		genMethod = genMethod.Incl("Hash")

	}

	ordDerive := derives.Find(func(v metafp.TypeClassDerive) bool {
		return v.TypeClass.Name == "Ord" && v.TypeClass.Package.Path() == "github.com/csgura/fp"
	})
	if ordDerive.IsDefined() && ts.Info.Method.Get("Eqv").IsEmpty() && !genMethod.Contains("Eqv") {

		if ordDerive.Get().IsRecursive() {
			fmt.Fprintf(w, `
				func(r %s) Eqv(other %s) bool {
					return %s().Eqv(r, other)
				}
			`, valuereceiver, valuereceiver,
				ordDerive.Get().GeneratedInstanceName(),
			)
		} else {
			fmt.Fprintf(w, `
				func(r %s) Eqv(other %s) bool {
					return %s.Eqv(r, other)
				}
			`, valuereceiver, valuereceiver,
				ordDerive.Get().GeneratedInstanceName(),
			)
		}
		genMethod = genMethod.Incl("Eqv")

	}

	if ordDerive.IsDefined() && ts.Info.Method.Get("Less").IsEmpty() {

		if ordDerive.Get().IsRecursive() {
			fmt.Fprintf(w, `
				func(r %s) Less(other %s) bool {
					return %s().Less(r,other)
				}
			`, valuereceiver, valuereceiver,
				ordDerive.Get().GeneratedInstanceName(),
			)
		} else {
			fmt.Fprintf(w, `
				func(r %s) Less(other %s) bool {
					return %s.Less(r, other)
				}
			`, valuereceiver, valuereceiver,
				ordDerive.Get().GeneratedInstanceName(),
			)
		}
		genMethod = genMethod.Incl("Less")

	}

	cloneDerive := derives.Find(func(v metafp.TypeClassDerive) bool {
		return v.TypeClass.Name == "Cloner" && v.TypeClass.Package.Path() == "github.com/csgura/fp"
	})

	if cloneDerive.IsDefined() && ts.Info.Method.Get("Clone").IsEmpty() && valuetp == "" {

		if cloneDerive.Get().IsRecursive() {
			fmt.Fprintf(w, `
				func(r %s) Clone() %s {
					return %s().Clone(r)
				}
			`, valuereceiver, valuereceiver,
				cloneDerive.Get().GeneratedInstanceName(),
			)
		} else {
			fmt.Fprintf(w, `
				func(r %s) Clone() %s {
					return %s.Clone(r)
				}
			`, valuereceiver, valuereceiver,
				cloneDerive.Get().GeneratedInstanceName(),
			)
		}
		genMethod = genMethod.Incl("Eqv")

	}
	return genMethod
}

func genMutable(ctx TaggedStructContext, genMethod fp.Set[string]) fp.Set[string] {
	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	allFields := applyFields(ts)

	valuetpdec := ts.Info.TypeParamDecl(w, workingPackage)
	valuetp := ts.Info.TypeParamIns(w, workingPackage)

	valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

	mutablereceiver := fmt.Sprintf("%sMutable%s", ts.Name, valuetp)
	mutableTypeName := fmt.Sprintf("%sMutable", ts.Name)

	mutableType := mutableTypeName + valuetpdec

	if !isTypeDefined(workingPackage, mutableTypeName) {
		mutableFields := iterator.Map(seq.Iterator(ts.Fields), func(v metafp.StructField) string {

			tag := v.Tag

			if !strings.HasPrefix(v.Name, "_") {
				if ts.Tags.Contains("@fp.JsonTag") || ts.Tags.Contains("@fp.Json") {
					if !strings.Contains(tag, "json") {
						if tag != "" {
							tag = tag + " "
						}
						if v.FieldType.IsNilable() || v.FieldType.IsOption() {
							tag = tag + fmt.Sprintf(`json:"%s,omitempty"`, v.Name)
						} else {
							tag = tag + fmt.Sprintf(`json:"%s"`, v.Name)
						}
					}
				}
			}

			name := fp.Seq[string]{}
			if !v.Embedded {
				name = name.Append(publicName(v.Name))
			}

			name = name.Append(v.TypeName(w, workingPackage))

			if tag != "" {
				name = name.Append(fmt.Sprintf("`%s`", tag))
			}
			return name.MakeString(" ")
		}).MakeString("\n")

		fmt.Fprintf(w, `
				type %s struct {
					%s
				}
			`, mutableType, mutableFields)
	}

	if ts.Info.Method.Get("AsMutable").IsEmpty() {

		fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
			return fmt.Sprintf(`%s : r.%s`, publicName(f.Name), f.Name)
		}).MakeString(",\n")

		fmt.Fprintf(w, `
					func(r %s) AsMutable() %s {
						return %s{
							%s,
						}
					}

				`, valuereceiver, mutablereceiver,
			mutablereceiver, fields,
		)
		genMethod = genMethod.Incl("AsMutable")

	}

	if !isMethodDefined(workingPackage, mutableTypeName, "AsImmutable") {

		fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
			return fmt.Sprintf(`%s : r.%s`, f.Name, publicName(f.Name))
		}).MakeString(",\n")

		fmt.Fprintf(w, `
					func(r %s) AsImmutable() %s {
						return %s{
							%s,
						}
					}

				`, mutablereceiver, valuereceiver,
			valuereceiver, fields,
		)
	}

	return genMethod
}

func processValue(ctx TaggedStructContext, genMethod fp.Set[string], keyTags fp.Set[string]) (fp.Set[string], fp.Set[string]) {

	ts := ctx.ts
	w := ctx.w
	workingPackage := ctx.workingPackage

	if _, ok := ts.Tags.Get("@fp.Value").Unapply(); ok {

		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)
		allFields := applyFields(ts)

		if allFields.Size() == 0 {
			return genMethod, keyTags
		}

		valuetp := ts.Info.TypeParamIns(w, workingPackage)

		//valueType := fmt.Sprintf("%s%s", v.Name, valuetpdec)

		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

		genMethod = genPrivateGetters(ctx, privateFields, genMethod)
		genMethod = genPrivateWiths(ctx, privateFields, genMethod)
		genMethod = genStringMethod(ctx, allFields, genMethod)
		genMethod = genUnapply(ctx, allFields, genMethod)

		if ts.Info.Method.Get("AsMap").IsEmpty() {

			fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
				if f.FieldType.IsOption() {
					return fmt.Sprintf(`if r.%s.IsDefined() {
							m["%s"] = r.%s.Get()
						}`, f.Name, f.Name, f.Name)
				} else {
					return fmt.Sprintf(`m["%s"] = r.%s`, f.Name, f.Name)
				}
			}).MakeString("\n")

			fmt.Fprintf(w, `
					func(r %s) AsMap() map[string]any {
						m := map[string]any{}
						%s
						return m
					}

				`, valuereceiver,
				fields,
			)
			genMethod = genMethod.Incl("AsMap")

		}

		if ts.Tags.Contains("@fp.GenLabelled") {
			allFields.Foreach(func(v metafp.StructField) {
				keyTags = keyTags.Incl(v.Name)
			})

			if allFields.Size() < max.Product {

				asalias := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

				arity := fp.Min(allFields.Size(), max.Product-1)

				if ts.Info.Method.Get("AsLabelled").IsEmpty() {
					tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
						return fmt.Sprintf("%s[%s]", namedName(w, workingPackage, workingPackage, v.Name), v.TypeName(w, workingPackage))
					}).MakeString(",")

					fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
						return fmt.Sprintf(`%s[%s]{r.%s, %s}`, namedName(w, workingPackage, workingPackage, f.Name), f.TypeName(w, workingPackage), f.Name, "`"+f.Tag+"`")
					}).Take(arity).MakeString(",")

					fppkg := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

					fmt.Fprintf(w, `
					func(r %s) AsLabelled() %s.Labelled%d[%s] {
						return %s.Labelled%d(%s)
					}

				`, valuereceiver, fppkg, arity, tp,
						asalias, arity, fields,
					)
					genMethod = genMethod.Incl("AsLabelled")

				}

			}
		}

		if ts.Tags.Contains("@fp.Json") {

			if ts.Info.Method.Get("MarshalJSON").IsEmpty() {
				jsonpk := w.GetImportedName(genfp.NewImportPackage("encoding/json", "json"))

				fmt.Fprintf(w, `
					func(r %s) MarshalJSON() ([]byte, error) {
						m := r.AsMutable()
						return %s.Marshal(m)
					}
				`, valuereceiver, jsonpk)
				genMethod = genMethod.Incl("MarshalJSON")

			}

			if ts.Info.Method.Get("UnmarshalJSON").IsEmpty() {
				httppk := w.GetImportedName(genfp.NewImportPackage("net/http", "http"))
				fppk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
				jsonpk := w.GetImportedName(genfp.NewImportPackage("encoding/json", "json"))

				fmt.Fprintf(w, `
					func(r *%s) UnmarshalJSON(b []byte) error {
						if r == nil {
							return %s.Error(%s.StatusBadRequest, "target ptr is nil")
						}
						m := r.AsMutable()
						err := %s.Unmarshal(b, &m)
						if err == nil {
							*r = m.AsImmutable()
						}
						return err
					}
				`, valuereceiver,
					fppk, httppk,
					jsonpk,
				)
				genMethod = genMethod.Incl("UnmarshalJSON")

			}
		}

		genMethod = genBuilder(ctx, genMethod)
		genMethod = genMutable(ctx, genMethod)

	}
	return genMethod, keyTags
}

func genValueAndGetter() {
	pack := os.Getenv("GOPACKAGE")

	genfp.Generate(pack, value_generated_file_name(pack), func(w genfp.Writer) {

		cwd, _ := os.Getwd()

		cfg := &packages.Config{
			Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
		}

		pkgs, err := packages.Load(cfg, cwd)
		if err != nil {
			fmt.Println(err)
			return
		}

		workingPackage := genfp.NewWorkingPackage(pkgs[0].Types, pkgs[0].Fset, pkgs[0].Syntax)

		st := metafp.FindTaggedStruct(pkgs,
			"@fp.Value",
			"@fp.GetterPubField",
			"@fp.Deref",
			"@fp.WithPubField",
			"@fp.Getter",
			"@fp.String",
			"@fp.With",
			"@fp.Builder",
			"@fp.AllArgsConstructor",
			"@fp.RequiredArgsConstructor",
		)

		if st.Size() == 0 {
			return
		}

		deriveCtx := NewTypeClassSummonContext(pkgs, genfp.NewImportSet())

		genTaggedStruct(w, workingPackage, st, deriveCtx)

	})
}

func value_generated_file_name(pack string) string {
	return pack + "_value_generated.go"
}

func derive_generated_file_name(pack string) string {
	return pack + "_derive_generated.go"
}

func delete_gen_files(pack string) {
	os.Remove(value_generated_file_name(pack))
	os.Remove(derive_generated_file_name(pack))
}

func main() {

	pack := os.Getenv("GOPACKAGE")
	if pack == "" {
		fmt.Println("invalid package. please run gombok using go generate command")
		return
	}

	delete_gen_files(pack)
	genValueAndGetter()
	genGenerate()
	//fmt.Printf("GOPACKAGE = %s\n", pack)
	genDerive()

}
