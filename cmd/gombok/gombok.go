package main

import (
	"fmt"
	"go/types"
	"os"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"

	"golang.org/x/tools/go/packages"
)

func isSamePkg(p1 *types.Package, p2 *types.Package) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}

	return p1.Path() == p2.Path()
}

func namedName(w genfp.Writer, working *types.Package, typePkg *types.Package, name string) string {

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

func isTypeDefined(pk *types.Package, name string) bool {
	return pk.Scope().Lookup(name) != nil
}

func isMethodDefined(pk *types.Package, tpeName string, method string) bool {
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

func processDeref(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	if _, ok := ts.Tags.Get("@fp.Deref").Unapply(); ok {
		//fmt.Printf("rhs type = %s\n", ts.RhsType)

		if rhs, ok := ts.RhsType.Unapply(); ok {

			//valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
			valuetp := ts.Info.TypeParamIns(w, workingPackage)
			typetp := ts.Info.TypeParamDecl(w, workingPackage)

			valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

			rhsTypeName := w.TypeName(workingPackage, rhs.Type)

			unwrapFunc := "Deref"
			if ts.Info.Method.Get(unwrapFunc).IsEmpty() {

				fmt.Fprintf(w, `
					func (r %s) %s() %s {
						return %s(r)
					}
				`, valuereceiver, unwrapFunc, rhsTypeName, rhsTypeName)
				genMethod = genMethod.Incl(unwrapFunc)
			}

			castFunc := fmt.Sprintf("Into%s%s", ts.Name, typetp)
			if workingPackage.Scope().Lookup(castFunc) == nil {
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

func processGetter(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	if _, ok := ts.Tags.Get("@fp.Getter").Unapply(); ok {
		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

		genMethod = genPrivateGetters(w, workingPackage, ts, privateFields, genMethod)
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

				ftp := w.TypeName(workingPackage, f.Type.Type)

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

func genStringMethod(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, allFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {
	if ts.Info.Method.Get("String").IsEmpty() && !genMethod.Contains("String") {

		valuereceiver := ts.Info.TypeStr(w, workingPackage)
		fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))

		printable := allFields.Filter(func(v metafp.StructField) bool {
			return v.Type.IsPrintable()
		})
		fm := iterator.Map(iterator.FromSeq(printable), func(f metafp.StructField) string {
			return fmt.Sprintf("%s=%%v", f.Name)
		}).MakeString(", ")

		fields := iterator.Map(iterator.FromSeq(printable), func(f metafp.StructField) string {
			return fmt.Sprintf("r.%s", f.Name)
		}).MakeString(",")

		fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, valuereceiver,
			fmtalias, ts.Name, fm, fields,
		)
		genMethod = genMethod.Incl("String")

	}

	return genMethod
}

func genUnapply(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, allFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {
	valuereceiver := ts.Info.TypeStr(w, workingPackage)
	if allFields.Size() < max.Product {
		asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		if ts.Info.Method.Get("AsTuple").IsEmpty() && !genMethod.Contains("AsTuple") {
			fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

			arity := fp.Min(allFields.Size(), max.Product-1)
			tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
				return w.TypeName(workingPackage, v.Type.Type)
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
			return w.TypeName(workingPackage, v.Type.Type)
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
func processString(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	if _, ok := ts.Tags.Get("@fp.String").Unapply(); ok {
		allFields := applyFields(ts)

		genMethod = genStringMethod(w, workingPackage, ts, allFields, genMethod)
	}

	return genMethod
}

func applyFields(ts metafp.TaggedStruct) fp.Seq[metafp.StructField] {
	return ts.Fields.FilterNot(func(v metafp.StructField) bool {
		// field 가 아무 것도 없는 embedded struct 는 생성에서 제외
		return strings.HasPrefix(v.Name, "_") || (v.Embedded && v.Type.Underlying().IsStruct() && v.Type.Fields().Size() == 0)
	})
}

func genBuilder(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

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
			ftp := w.TypeName(workingPackage, f.Type.Type)

			fmt.Fprintf(w, `
						func (r %s) %s( v %s) %s {
							r.%s = v
							return r
						}
					`, builderreceiver, uname, ftp, builderreceiver, f.Name)
		}

		if f.Type.IsOption() {
			optiont := w.TypeName(workingPackage, f.Type.TypeArgs.Head().Get().Type)
			optionpk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/option", "option"))

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
			fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

			arity := fp.Min(allFields.Size(), max.Product-1)

			tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
				return w.TypeName(workingPackage, v.Type.Type)
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
			return fmt.Sprintf("%s %s", v.Name, w.TypeName(workingPackage, v.Type.Type))
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
			if f.Type.IsOption() {
				optionpk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/option", "option"))

				return fmt.Sprintf(`if v , ok := m["%s"].(%s); ok {
							r.%s = v
						} else if v, ok := m["%s"].(%s); ok {
							r.%s = %s.Some(v)
						}
						`, f.Name, w.TypeName(workingPackage, f.Type.Type),
					f.Name,
					f.Name, w.TypeName(workingPackage, f.Type.TypeArgs.Head().Get().Type),
					f.Name, optionpk,
				)
			} else {
				return fmt.Sprintf(`if v , ok := m["%s"].(%s); ok {
							r.%s = v
						}
							`, f.Name, w.TypeName(workingPackage, f.Type.Type),
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

				fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

				tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
					return fmt.Sprintf("%s[%s]", namedName(w, workingPackage, workingPackage, v.Name), w.TypeName(workingPackage, v.Type.Type))
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

func processBuilder(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	if _, ok := ts.Tags.Get("@fp.Builder").Unapply(); ok {

		genMethod = genBuilder(w, workingPackage, ts, genMethod)
	}

	return genMethod
}

func genPrivateWiths(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, privateFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {

	valuereceiver := ts.Info.TypeStr(w, workingPackage)

	privateFields.Foreach(func(f metafp.StructField) {

		uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]

		fnName := "With" + uname

		if ts.Info.Method.Get(fnName).IsEmpty() && !genMethod.Contains(fnName) {
			ftp := w.TypeName(workingPackage, f.Type.Type)

			fmt.Fprintf(w, `
						func (r %s) %s(v %s) %s {
							r.%s = v
							return r
						}
					`, valuereceiver, fnName, ftp, valuereceiver, f.Name)
			genMethod = genMethod.Incl(fnName)

		}

		if f.Type.IsOption() {
			optiont := w.TypeName(workingPackage, f.Type.TypeArgs.Head().Get().Type)
			optionpk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/option", "option"))

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

func processWith(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	if _, ok := ts.Tags.Get("@fp.With").Unapply(); ok {
		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

		genMethod = genPrivateWiths(w, workingPackage, ts, privateFields, genMethod)
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

				ftp := w.TypeName(workingPackage, f.Type.Type)

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

func genAllArgsCons(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	fnName := fmt.Sprintf("New%s", ts.Name)

	if !genMethod.Contains(fnName) {

		allFields := applyFields(ts)

		tp := iterator.Map(seq.Iterator(allFields), func(v metafp.StructField) string {
			return fmt.Sprintf("%s %s", v.Name, w.TypeName(workingPackage, v.Type.Type))
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

func processAllArgsCons(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {
	if _, ok := ts.Tags.Get("@fp.AllArgsConstructor").Unapply(); ok {

		genMethod = genAllArgsCons(w, workingPackage, ts, genMethod)
	}
	return genMethod
}
func genTaggedStruct(w genfp.Writer, workingPackage *types.Package, st fp.Seq[metafp.TaggedStruct], derives fp.Seq[metafp.TypeClassDerive]) {
	keyTags := mutable.EmptySet[string]()

	st.Foreach(func(ts metafp.TaggedStruct) {
		genMethod := fp.Set[string]{}
		stDerives := derives.Filter(func(v metafp.TypeClassDerive) bool {
			return v.DeriveFor.Name == ts.Name
		})
		genMethod = processAllArgsCons(w, workingPackage, ts, genMethod)

		genMethod, keyTags = processValue(w, workingPackage, ts, stDerives, genMethod, keyTags)

		genMethod = processGetter(w, workingPackage, ts, genMethod)
		genMethod = processWith(w, workingPackage, ts, genMethod)
		genMethod = processDeref(w, workingPackage, ts, genMethod)
		genMethod = processString(w, workingPackage, ts, genMethod)

		genMethod = processBuilder(w, workingPackage, ts, genMethod)

	})

	klist := keyTags.Iterator().ToSeq()
	seq.Sort(klist, ord.Given[string]()).Foreach(func(name string) {
		fmt.Fprintf(w, `type %s[T any] fp.Tuple1[T]
			`, namedName(w, workingPackage, workingPackage, name))

		fmt.Fprintf(w, `func (r %s[T]) Name() string {
				return "%s"
			}
			`, namedName(w, workingPackage, workingPackage, name), name)

		fmt.Fprintf(w, `func (r %s[T]) Value() T {
				return r.I1
			}
			`, namedName(w, workingPackage, workingPackage, name))

		fmt.Fprintf(w, `func (r %s[T]) WithValue(v T) %s[T] {
				r.I1 = v
				return r
			}
			`, namedName(w, workingPackage, workingPackage, name), namedName(w, workingPackage, workingPackage, name))
	})
}

func genPrivateGetters(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, privateFields fp.Seq[metafp.StructField], genMethod fp.Set[string]) fp.Set[string] {

	valuetp := ts.Info.TypeParamIns(w, workingPackage)

	valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

	privateFields.Foreach(func(f metafp.StructField) {
		uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]

		fnName := uname
		if ts.Info.Method.Get(fnName).IsEmpty() && !genMethod.Contains(fnName) {
			ftp := w.TypeName(workingPackage, f.Type.Type)

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

func genTypeClassMethod(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, derives fp.Seq[metafp.TypeClassDerive], genMethod fp.Set[string]) fp.Set[string] {

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
		fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

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

func genMutable(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	allFields := applyFields(ts)

	valuetpdec := ts.Info.TypeParamDecl(w, workingPackage)
	valuetp := ts.Info.TypeParamIns(w, workingPackage)

	valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

	mutablereceiver := fmt.Sprintf("%sMutable%s", ts.Name, valuetp)
	mutableTypeName := fmt.Sprintf("%sMutable", ts.Name)

	mutableType := mutableTypeName + valuetpdec

	mutableFields := iterator.Map(seq.Iterator(ts.Fields), func(v metafp.StructField) string {

		tag := v.Tag

		if !strings.HasPrefix(v.Name, "_") {
			if ts.Tags.Contains("@fp.JsonTag") || ts.Tags.Contains("@fp.Json") {
				if !strings.Contains(tag, "json") {
					if tag != "" {
						tag = tag + " "
					}
					if v.Type.IsNilable() || v.Type.IsOption() {
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

		name = name.Append(w.TypeName(workingPackage, v.Type.Type))

		if tag != "" {
			name = name.Append(fmt.Sprintf("`%s`", tag))
		}
		return name.MakeString(" ")
	}).MakeString("\n")

	if !isTypeDefined(workingPackage, mutableTypeName) {
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

func processValue(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, derives fp.Seq[metafp.TypeClassDerive], genMethod fp.Set[string], keyTags fp.Set[string]) (fp.Set[string], fp.Set[string]) {

	if _, ok := ts.Tags.Get("@fp.Value").Unapply(); ok {

		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)
		allFields := applyFields(ts)

		if allFields.Size() == 0 {
			return genMethod, keyTags
		}

		asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		valuetp := ts.Info.TypeParamIns(w, workingPackage)

		//valueType := fmt.Sprintf("%s%s", v.Name, valuetpdec)

		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

		genMethod = genPrivateGetters(w, workingPackage, ts, privateFields, genMethod)
		genMethod = genPrivateWiths(w, workingPackage, ts, privateFields, genMethod)
		genMethod = genStringMethod(w, workingPackage, ts, allFields, genMethod)
		genMethod = genUnapply(w, workingPackage, ts, allFields, genMethod)

		if ts.Info.Method.Get("AsMap").IsEmpty() {

			fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
				if f.Type.IsOption() {
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

				arity := fp.Min(allFields.Size(), max.Product-1)

				tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
					return fmt.Sprintf("%s[%s]", namedName(w, workingPackage, workingPackage, v.Name), w.TypeName(workingPackage, v.Type.Type))
				}).MakeString(",")

				fields := iterator.Map(iterator.FromSeq(allFields), func(f metafp.StructField) string {
					return fmt.Sprintf(`%s[%s]{r.%s}`, namedName(w, workingPackage, workingPackage, f.Name), w.TypeName(workingPackage, f.Type.Type), f.Name)
				}).Take(arity).MakeString(",")

				if ts.Info.Method.Get("AsLabelled").IsEmpty() {
					fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

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
				jsonpk := w.GetImportedName(types.NewPackage("encoding/json", "json"))

				fmt.Fprintf(w, `
					func(r %s) MarshalJSON() ([]byte, error) {
						m := r.AsMutable()
						return %s.Marshal(m)
					}
				`, valuereceiver, jsonpk)
				genMethod = genMethod.Incl("MarshalJSON")

			}

			if ts.Info.Method.Get("UnmarshalJSON").IsEmpty() {
				httppk := w.GetImportedName(types.NewPackage("net/http", "http"))
				fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				jsonpk := w.GetImportedName(types.NewPackage("encoding/json", "json"))

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

		genMethod = genBuilder(w, workingPackage, ts, genMethod)
		genMethod = genMutable(w, workingPackage, ts, genMethod)

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

		workingPackage := pkgs[0].Types

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
		)

		if st.Size() == 0 {
			return
		}

		derives := metafp.FindTypeClassDerive(pkgs)

		genTaggedStruct(w, workingPackage, st, derives)

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
	genGenerate()

	//fmt.Printf("GOPACKAGE = %s\n", pack)
	genValueAndGetter()
	genDerive()

}
