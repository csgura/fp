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

func iterate[T any](len int, getter func(idx int) T, fn func(int, T) string) fp.Seq[string] {
	ret := []string{}
	for i := 0; i < len; i++ {
		ret = append(ret, fn(i, getter(i)))
	}
	return ret
}

func genAlias(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct, genMethod fp.Set[string]) fp.Set[string] {

	if rhs, ok := ts.RhsType.Unapply(); ok {
		valuetp := ""
		if len(ts.Info.TypeParam) > 0 {
			//valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
			valuetp = "[" + ts.Info.TypeParamIns(w, workingPackage) + "]"
		}
		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

		rhsTypeName := w.TypeName(workingPackage, rhs.Type)

		unwrapFunc := "Unwrap"
		if ts.Info.Method.Get(unwrapFunc).IsEmpty() {

			fmt.Fprintf(w, `
					func (r %s) %s() %s {
						return %s(r)
					}
				`, valuereceiver, unwrapFunc, rhsTypeName, rhsTypeName)
			genMethod = genMethod.Incl(unwrapFunc)
		}

		castFunc := fmt.Sprintf("Cast%s", ts.Name)
		if workingPackage.Scope().Lookup(castFunc) == nil {
			fmt.Fprintf(w, `
					func %s(v %s) %s {
						return %s(v)
					}
				`, castFunc, rhsTypeName, valuereceiver,
				ts.Name,
			)
		}

		rhs.Method.Iterator().Foreach(func(t fp.Tuple2[string, *types.Func]) {
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
	return genMethod
}
func genGetter(w genfp.Writer, workingPackage *types.Package, ts metafp.TaggedStruct) fp.Set[string] {
	ret := fp.Set[string]{}

	if anno, ok := ts.Tags.Get("@fp.GetterPubField").Unapply(); ok {

		publicFields := ts.Fields.Filter(metafp.StructField.Public)

		if publicFields.Size() == 0 {
			return ret
		}

		//valuetpdec := ""
		valuetp := ""
		if len(ts.Info.TypeParam) > 0 {
			//valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
			valuetp = "[" + ts.Info.TypeParamIns(w, workingPackage) + "]"
		}

		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)

		publicFields.Foreach(func(f metafp.StructField) {

			getterName := fmt.Sprintf("Get%s", f.Name)

			if ts.Info.Method.Get(getterName).IsEmpty() {
				if anno.Params().Get("override").OrElse("false") != "true" && ts.Tags.Contains("@fp.Alias") && ts.RhsType.IsDefined() {
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

				ret = ret.Incl(getterName)
			}
		})
	}
	return ret
}
func genGetterAndAlias(w genfp.Writer, workingPackage *types.Package, st fp.Seq[metafp.TaggedStruct]) {
	st.Foreach(func(ts metafp.TaggedStruct) {
		genMethod := genGetter(w, workingPackage, ts)

		genMethod = genAlias(w, workingPackage, ts, genMethod)

	})
}

func genValue(w genfp.Writer, workingPackage *types.Package, st fp.Seq[metafp.TaggedStruct]) {
	keyTags := mutable.EmptySet[string]()

	st.Foreach(func(ts metafp.TaggedStruct) {

		privateFields := ts.Fields.FilterNot(metafp.StructField.Public)
		allFields := ts.Fields.FilterNot(func(v metafp.StructField) bool {
			return strings.HasPrefix(v.Name, "_")
		})

		if privateFields.Size() == 0 {
			return
		}

		asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		valuetpdec := ""
		valuetp := ""
		if len(ts.Info.TypeParam) > 0 {
			valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
			valuetp = "[" + ts.Info.TypeParamIns(w, workingPackage) + "]"
		}

		builderTypeName := fmt.Sprintf("%sBuilder", ts.Name)
		builderType := builderTypeName + valuetpdec

		mutableTypeName := fmt.Sprintf("%sMutable", ts.Name)
		mutableType := mutableTypeName + valuetpdec

		//valueType := fmt.Sprintf("%s%s", v.Name, valuetpdec)

		valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)
		builderreceiver := fmt.Sprintf("%sBuilder%s", ts.Name, valuetp)
		mutablereceiver := fmt.Sprintf("%sMutable%s", ts.Name, valuetp)

		if !isTypeDefined(workingPackage, builderTypeName) {
			fmt.Fprintf(w, `
					type %s %s
				`, builderType, valuereceiver)
		}

		mutableFields := iterator.Map(seq.Iterator(ts.Fields), func(v metafp.StructField) string {
			tag := v.Tag

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
		}

		privateFields.Foreach(func(f metafp.StructField) {

			uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]

			if ts.Info.Method.Get(uname).IsEmpty() {
				ftp := w.TypeName(workingPackage, f.Type.Type)

				fmt.Fprintf(w, `
						func (r %s) %s() %s {
							return r.%s
						}
					`, valuereceiver, uname, ftp, f.Name)

			}

			if ts.Info.Method.Get("With" + uname).IsEmpty() {
				ftp := w.TypeName(workingPackage, f.Type.Type)

				fmt.Fprintf(w, `
						func (r %s) With%s(v %s) %s {
							r.%s = v
							return r
						}
					`, valuereceiver, uname, ftp, valuereceiver, f.Name)
			}

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

				if ts.Info.Method.Get("WithSome" + uname).IsEmpty() {

					fmt.Fprintf(w, `
							func (r %s) WithSome%s(v %s) %s {
								r.%s = %s.Some(v)
								return r
							}
						`, valuereceiver, uname, optiont, valuereceiver, f.Name, optionpk)
				}

				if ts.Info.Method.Get("WithNone" + uname).IsEmpty() {

					fmt.Fprintf(w, `
							func (r %s) WithNone%s() %s {
								r.%s = %s.None[%s]()
								return r
							}
						`, valuereceiver, uname, valuereceiver, f.Name, optionpk, optiont)
				}

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

		if ts.Info.Method.Get("String").IsEmpty() {
			fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))

			printable := allFields.Filter(func(v metafp.StructField) bool {
				return v.Type.IsPrintable()
			})
			fm := seq.Iterator(seq.Map(printable, func(f metafp.StructField) string {
				return fmt.Sprintf("%s=%%v", f.Name)
			})).MakeString(", ")

			fields := seq.Iterator(seq.Map(printable, func(f metafp.StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			})).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, valuereceiver,
				fmtalias, ts.Name, fm, fields,
			)
		}

		if privateFields.Size() < max.Product {
			if ts.Info.Method.Get("AsTuple").IsEmpty() {
				fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

				arity := fp.Min(allFields.Size(), max.Product-1)
				tp := iterator.Map(seq.Iterator(allFields).Take(arity), func(v metafp.StructField) string {
					return w.TypeName(workingPackage, v.Type.Type)
				}).MakeString(",")

				fields := seq.Iterator(seq.Map(allFields, func(f metafp.StructField) string {
					return fmt.Sprintf("r.%s", f.Name)
				})).Take(arity).MakeString(",")

				fmt.Fprintf(w, `
					func(r %s) AsTuple() %s.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, fppkg, arity, tp,
					asalias, arity, fields,
				)
			}
		}
		if ts.Info.Method.Get("Unapply").IsEmpty() {

			tp := seq.Map(allFields, func(v metafp.StructField) string {
				return w.TypeName(workingPackage, v.Type.Type)
			}).MakeString(",")

			fields := seq.Map(allFields, func(f metafp.StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) Unapply() (%s) {
						return %s
					}

				`, valuereceiver, tp,
				fields,
			)
		}

		if ts.Info.Method.Get("AsMutable").IsEmpty() {

			fields := seq.Iterator(seq.Map(allFields, func(f metafp.StructField) string {
				return fmt.Sprintf(`%s : r.%s`, publicName(f.Name), f.Name)
			})).MakeString(",\n")

			fmt.Fprintf(w, `
					func(r %s) AsMutable() %s {
						return %s{
							%s,
						}
					}

				`, valuereceiver, mutablereceiver,
				mutablereceiver, fields,
			)
		}

		if !isMethodDefined(workingPackage, mutableTypeName, "AsImmutable") {

			fields := seq.Iterator(seq.Map(allFields, func(f metafp.StructField) string {
				return fmt.Sprintf(`%s : r.%s`, f.Name, publicName(f.Name))
			})).MakeString(",\n")

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

		if privateFields.Size() < max.Product {

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

		if ts.Info.Method.Get("AsMap").IsEmpty() {

			fields := seq.Iterator(seq.Map(allFields, func(f metafp.StructField) string {
				if f.Type.IsOption() {
					return fmt.Sprintf(`if r.%s.IsDefined() {
							m["%s"] = r.%s.Get()
						}`, f.Name, f.Name, f.Name)
				} else {
					return fmt.Sprintf(`m["%s"] = r.%s`, f.Name, f.Name)
				}
			})).MakeString("\n")

			fmt.Fprintf(w, `
					func(r %s) AsMap() map[string]any {
						m := map[string]any{}
						%s
						return m
					}

				`, valuereceiver,
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

		if ts.Tags.Contains("@fp.GenLabelled") {
			privateFields.Foreach(func(v metafp.StructField) {
				keyTags = keyTags.Incl(v.Name)
			})

			if privateFields.Size() < max.Product {

				arity := fp.Min(privateFields.Size(), max.Product-1)

				tp := iterator.Map(seq.Iterator(privateFields).Take(arity), func(v metafp.StructField) string {
					return fmt.Sprintf("Named%s[%s]", publicName(v.Name), w.TypeName(workingPackage, v.Type.Type))
				}).MakeString(",")

				fields := seq.Iterator(seq.Map(privateFields, func(f metafp.StructField) string {
					return fmt.Sprintf(`Named%s[%s]{r.%s}`, publicName(f.Name), w.TypeName(workingPackage, f.Type.Type), f.Name)
				})).Take(arity).MakeString(",")

				if ts.Info.Method.Get("AsLabelled").IsEmpty() {
					fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

					fmt.Fprintf(w, `
					func(r %s) AsLabelled() %s.Labelled%d[%s] {
						return %s.Labelled%d(%s)
					}

				`, valuereceiver, fppkg, arity, tp,
						asalias, arity, fields,
					)

				}
				if !isMethodDefined(workingPackage, builderTypeName, "FromLabelled") {
					arity := fp.Min(privateFields.Size(), max.Product-1)

					fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

					fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), seq.Iterator(privateFields)), func(f fp.Tuple2[int, metafp.StructField]) string {
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

		if ts.Tags.Contains("@fp.Json") {

			if ts.Info.Method.Get("MarshalJSON").IsEmpty() {
				jsonpk := w.GetImportedName(types.NewPackage("encoding/json", "json"))

				fmt.Fprintf(w, `
					func(r %s) MarshalJSON() ([]byte, error) {
						m := r.AsMutable()
						return %s.Marshal(m)
					}
				`, valuereceiver, jsonpk)
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
			}
		}

	})

	klist := keyTags.Iterator().ToSeq()
	seq.Sort(klist, ord.Given[string]()).Foreach(func(name string) {
		fmt.Fprintf(w, `type Named%s[T any] fp.Tuple1[T]
			`, publicName(name))

		fmt.Fprintf(w, `func (r Named%s[T]) Name() string {
				return "%s"
			}
			`, publicName(name), name)

		fmt.Fprintf(w, `func (r Named%s[T]) Value() T {
				return r.I1
			}
			`, publicName(name))

		fmt.Fprintf(w, `func (r Named%s[T]) WithValue(v T) Named%s[T] {
				r.I1 = v
				return r
			}
			`, publicName(name), publicName(name))
	})
}
func genValueAndGetter() {
	pack := os.Getenv("GOPACKAGE")

	genfp.Generate(pack, pack+"_value_generated.go", func(w genfp.Writer) {

		cwd, _ := os.Getwd()

		//	fmt.Printf("cwd = %s , pack = %s file = %s, line = %s\n", try.Apply(os.Getwd()), pack, file, line)

		//packages.LoadFiles()

		cfg := &packages.Config{
			Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax,
		}

		pkgs, err := packages.Load(cfg, cwd)
		if err != nil {
			fmt.Println(err)
			return
		}

		workingPackage := pkgs[0].Types

		st := metafp.FindTaggedStruct(pkgs, "@fp.Value")
		gt := metafp.FindTaggedStruct(pkgs, "@fp.GetterPubField", "@fp.Alias")

		if st.Size() == 0 && gt.Size() == 0 {
			return
		}

		genValue(w, workingPackage, st)
		genGetterAndAlias(w, workingPackage, gt)

	})
}

func main() {
	pack := os.Getenv("GOPACKAGE")
	if pack == "" {
		fmt.Println("invalid package. please run gombok using go generate command")
		return
	}
	//fmt.Printf("GOPACKAGE = %s\n", pack)
	genValueAndGetter()
	genDerive()

}
