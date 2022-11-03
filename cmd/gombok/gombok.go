package main

import (
	"fmt"
	"go/types"
	"os"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"

	// . "github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

type typeClassMember struct {
	pack   *types.Package
	name   string
	object types.Object
}

func (r typeClassMember) Name() string {
	return r.name
}

func (r typeClassMember) PackagedName(importSet metafp.ImportSet, workingPackage *types.Package) string {
	if r.pack.Path() == workingPackage.Path() {
		return r.name
	}
	pk := importSet.GetImportedName(r.pack)
	return fmt.Sprintf("%s.%s", pk, r.name)

}

func (r typeClassMember) Type() types.Type {
	return r.object.Type()
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

func genValue() {
	pack := os.Getenv("GOPACKAGE")

	metafp.Generate(pack, pack+"_value_generated.go", func(w metafp.Writer) {

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

		if st.Size() == 0 {
			return
		}

		keyTags := mutable.EmptySet[string]()

		st.Foreach(func(ts metafp.TaggedStruct) {

			privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

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

			mutableFields := iterator.Map(ts.Fields.Iterator(), func(v metafp.StructField) string {
				tag := v.Tag

				if ts.Tags.Contains("@fp.JsonTag") || ts.Tags.Contains("@fp.Json") {
					if !strings.Contains(tag, "json") {
						if tag != "" {
							tag = tag + " "
						}
						if v.Type.IsNilable() {
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

				printable := privateFields.Filter(func(v metafp.StructField) bool {
					return v.Type.IsPrintable()
				})
				fm := seq.Map(printable, func(f metafp.StructField) string {
					return fmt.Sprintf("%s=%%v", f.Name)
				}).Iterator().MakeString(", ")

				fields := seq.Map(printable, func(f metafp.StructField) string {
					return fmt.Sprintf("r.%s", f.Name)
				}).Iterator().MakeString(",")

				fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, valuereceiver,
					fmtalias, ts.Name, fm, fields,
				)
			}

			if ts.Info.Method.Get("AsTuple").IsEmpty() {
				fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

				tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v metafp.StructField) string {
					return w.TypeName(workingPackage, v.Type.Type)
				}).MakeString(",")

				fields := seq.Map(privateFields, func(f metafp.StructField) string {
					return fmt.Sprintf("r.%s", f.Name)
				}).Iterator().Take(max.Product).MakeString(",")

				fmt.Fprintf(w, `
					func(r %s) AsTuple() %s.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, fppkg, privateFields.Size(), tp,
					asalias, privateFields.Size(), fields,
				)
			}

			if ts.Info.Method.Get("AsMutable").IsEmpty() {

				fields := seq.Map(privateFields, func(f metafp.StructField) string {
					return fmt.Sprintf(`%s : r.%s`, publicName(f.Name), f.Name)
				}).Iterator().Take(max.Product).MakeString(",\n")

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

				fields := seq.Map(ts.Fields, func(f metafp.StructField) string {
					return fmt.Sprintf(`%s : r.%s`, f.Name, publicName(f.Name))
				}).Iterator().Take(max.Product).MakeString(",\n")

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

			if !isMethodDefined(workingPackage, builderTypeName, "FromTuple") {
				fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

				tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v metafp.StructField) string {
					return w.TypeName(workingPackage, v.Type.Type)
				}).MakeString(",")

				fields := iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, metafp.StructField]) string {
					return fmt.Sprintf("r.%s = t.I%d", f.I2.Name, f.I1+1)
				}).Take(max.Product).MakeString("\n")

				fmt.Fprintf(w, `
					func (r %s) FromTuple(t %s.Tuple%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, fppkg, privateFields.Size(), tp, builderreceiver,
					fields,
				)
			}

			if ts.Info.Method.Get("AsMap").IsEmpty() {

				fields := seq.Map(privateFields, func(f metafp.StructField) string {
					return fmt.Sprintf(`"%s" : r.%s`, f.Name, f.Name)
				}).Iterator().Take(max.Product).MakeString(",\n")

				fmt.Fprintf(w, `
					func(r %s) AsMap() map[string]any {
						return map[string]any {
							%s,
						}
					}

				`, valuereceiver,
					fields,
				)
			}

			if !isMethodDefined(workingPackage, builderTypeName, "FromMap") {

				fields := iterator.Map(privateFields.Iterator(), func(f metafp.StructField) string {
					return fmt.Sprintf(`if v , ok := m["%s"].(%s); ok {
							r.%s = v
						}
							`, f.Name, w.TypeName(workingPackage, f.Type.Type),
						f.Name,
					)
				}).Take(max.Product).MakeString("\n")

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
				tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v metafp.StructField) string {
					return fmt.Sprintf("NameIs%s[%s]", publicName(v.Name), w.TypeName(workingPackage, v.Type.Type))
				}).MakeString(",")

				fields := seq.Map(privateFields, func(f metafp.StructField) string {
					return fmt.Sprintf(`NameIs%s[%s]{r.%s}`, publicName(f.Name), w.TypeName(workingPackage, f.Type.Type), f.Name)
				}).Iterator().Take(max.Product).MakeString(",")

				privateFields.Foreach(func(v metafp.StructField) {
					keyTags = keyTags.Incl(v.Name)
				})

				if ts.Info.Method.Get("AsLabelled").IsEmpty() {
					fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

					fmt.Fprintf(w, `
					func(r %s) AsLabelled() %s.Labelled%d[%s] {
						return %s.Labelled%d(%s)
					}

				`, valuereceiver, fppkg, privateFields.Size(), tp,
						asalias, privateFields.Size(), fields,
					)

				}
				if !isMethodDefined(workingPackage, builderTypeName, "FromLabelled") {
					fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

					fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, metafp.StructField]) string {
						return fmt.Sprintf("r.%s = t.I%d.Value()", f.I2.Name, f.I1+1)
					}).Take(max.Product).MakeString("\n")

					fmt.Fprintf(w, `
					func (r %s) FromLabelled(t %s.Labelled%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, fppkg, privateFields.Size(), tp, builderreceiver,
						fields,
					)
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
			fmt.Fprintf(w, `type NameIs%s[T any] fp.Tuple1[T]
			`, publicName(name))

			fmt.Fprintf(w, `func (r NameIs%s[T]) Name() string {
				return "%s"
			}
			`, publicName(name), name)

			fmt.Fprintf(w, `func (r NameIs%s[T]) Value() T {
				return r.I1
			}
			`, publicName(name))

			fmt.Fprintf(w, `func (r NameIs%s[T]) WithValue(v T) NameIs%s[T] {
				r.I1 = v
				return r
			}
			`, publicName(name), publicName(name))
		})
	})
}

type lookupTarget struct {
	instanceOf metafp.TypeInfo
	pk         *types.Package
	name       string
	required   fp.Seq[metafp.RequiredInstance]
	typeParam  bool
	instance   types.Object
	// tc       *TypeClass
}

func (r lookupTarget) isFunc() bool {
	if r.instance != nil {
		switch r.instance.Type().(type) {
		case *types.Signature:
			return true
		}
	}
	return false
}
func (r lookupTarget) lookup() lookupTarget {
	r.instance = r.pk.Scope().Lookup(r.name)
	return r
}

func (r lookupTarget) available(genSet mutable.Set[string]) bool {
	if r.typeParam {
		return true
	}
	if r.instance != nil {
		return true
	}

	if genSet.Contains(r.name) {
		return true
	}
	return false
}

func (r lookupTarget) instanceExpr(w metafp.Writer, workingPkg *types.Package) string {
	if r.pk == nil || r.pk.Path() == workingPkg.Path() {
		return r.name
	}

	// if r.typeParam {
	// 	return fmt.Sprintf("%s%s", privateName(r.tc.Name), r.name)
	// }

	pk := w.GetImportedName(r.pk)

	return fmt.Sprintf("%s.%s", pk, r.name)

}

func (r TypeClassSummonContext) typeclassInstanceMust(f metafp.TypeInfo, name string) lookupTarget {

	return lookupTarget{
		instanceOf: f,
		pk:         r.tc.Package,
		name:       r.tc.TypeClass.Name + publicName(name),
		required: seq.Map(f.TypeArgs, func(v metafp.TypeInfo) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: r.tc.TypeClass,
				Type:      v,
			}
		}),
	}
}

// f 는 Eq 쌓이지 않은 타입
// Eq[T] 같은거 아님
func (r TypeClassSummonContext) lookupTypeClassInstanceLocalDeclared(f metafp.TypeInfo, name ...string) fp.Option[lookupTarget] {

	scope := r.tcCache.Get(r.tc)

	itr := seq.FlatMap(name, func(v string) fp.Seq[string] {
		if f.Pkg != nil && r.tc.Package.Path() != f.Pkg.Path() {
			return []string{
				r.tc.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(v),
				r.tc.TypeClass.Name + publicName(v),
			}

		}
		return []string{r.tc.TypeClass.Name + publicName(v)}
	}).Iterator()

	ins := iterator.FlatMap(itr, func(v string) fp.Iterator[metafp.TypeClassInstance] {
		ret := option.Iterator(scope.WorkingScope.FindByName(v, f))

		if f.TypeArgs.Size() > 0 {
			tnames := seq.Map(f.TypeArgs, func(v metafp.TypeInfo) fp.Option[string] {
				return v.Name()
			})

			if tnames.ForAll(fp.Option[string].IsDefined) {
				//n := iterator.Map(tnames.Iterator(), fp.Option[string].Get).MakeString("")
				r := scope.WorkingScope.Find(f)
				//fmt.Printf("find %s , type = %s , result = %v\n", v+n, f.Type.String(), r)
				//fixed := option.Iterator(r)
				return r.Iterator().Concat(ret)
			}
		}
		return ret
	})

	return iterator.Map(ins, func(v metafp.TypeClassInstance) lookupTarget {
		return lookupTarget{
			instanceOf: f,
			pk:         r.tc.Package,
			name:       v.Name,

			// 함수의 아규먼트는 Eq 가 포함 되어 있음.
			required: v.RequiredInstance,
		}

	}).Head()

}

func (r TypeClassSummonContext) lookupHNilMust() typeClassMember {
	ret := r.lookupTypeClassMember("HNil")
	if ret.IsDefined() {
		return ret.Get()
	}

	ret = r.lookupTypeClassMember("HlistNil")
	if ret.IsDefined() {
		return ret.Get()
	}
	nameWithTc := r.tc.TypeClass.Name + "HNil"

	return typeClassMember{r.tc.Package, nameWithTc, nil}
}

func (r TypeClassSummonContext) lookupTypeClassMember(name string) fp.Option[typeClassMember] {
	nameWithTc := r.tc.TypeClass.Name + name

	ins := r.tc.Package.Scope().Lookup(nameWithTc)
	if ins != nil {
		return fp.Some(typeClassMember{r.tc.Package, nameWithTc, ins})
	}

	ins = r.tc.PrimitiveInstancePkg.Scope().Lookup(nameWithTc)
	if ins != nil {
		return fp.Some(typeClassMember{r.tc.PrimitiveInstancePkg, nameWithTc, ins})
	}

	ins = r.tc.PrimitiveInstancePkg.Scope().Lookup(name)
	if ins != nil {
		return fp.Some((typeClassMember{r.tc.PrimitiveInstancePkg, name, ins}))
	}
	return fp.None[typeClassMember]()
}

func (r TypeClassSummonContext) lookupTypeClassMemberMust(name string) typeClassMember {
	ret := r.lookupTypeClassMember(name)
	if ret.IsDefined() {
		return ret.Get()
	}

	nameWithTc := r.tc.TypeClass.Name + name

	return typeClassMember{r.tc.Package, nameWithTc, nil}
}

func (r TypeClassSummonContext) lookupTypeClassInstancePrimitivePkgLazy(f metafp.TypeInfo, name ...string) func() fp.Option[lookupTarget] {
	return func() fp.Option[lookupTarget] {
		return r.lookupTypeClassInstancePrimitivePkg(f, name...)
	}
}

func (r TypeClassSummonContext) lookupTypeClassInstancePrimitivePkg(f metafp.TypeInfo, name ...string) fp.Option[lookupTarget] {

	itr := seq.FlatMap(name, func(v string) fp.Seq[string] {
		ret := seq.Of(
			r.tc.TypeClass.Name+publicName(v),
			publicName(v),
		)
		if f.Pkg != nil {
			return seq.Of(
				r.tc.TypeClass.Name+publicName(f.Pkg.Name())+publicName(v),
				publicName(f.Pkg.Name())+publicName(v),
			).Concat(ret)
		}
		return ret
	}).Iterator()

	return iterator.Map(itr, func(v string) lookupTarget {
		return lookupTarget{
			instanceOf: f,
			pk:         r.tc.PrimitiveInstancePkg,
			name:       v,
			required: seq.Map(f.TypeArgs, func(v metafp.TypeInfo) metafp.RequiredInstance {
				return metafp.RequiredInstance{
					TypeClass: r.tc.TypeClass,
					Type:      v,
				}
			}),
		}.lookup()
	}).Filter(func(lt lookupTarget) bool {
		return lt.available(r.genSet)
	}).Head()

}

func (r TypeClassSummonContext) lookupTypeClassInstanceTypePkg(f metafp.TypeInfo, name string) fp.Option[lookupTarget] {

	if f.Pkg != nil && f.Pkg.Path() != r.tc.Package.Path() {
		ret := lookupTarget{
			instanceOf: f,
			pk:         f.Pkg,
			name:       r.tc.TypeClass.Name + publicName(name),
			required: seq.Map(f.TypeArgs, func(v metafp.TypeInfo) metafp.RequiredInstance {
				return metafp.RequiredInstance{
					TypeClass: r.tc.TypeClass,
					Type:      v,
				}
			}),
		}.lookup()

		if ret.available(r.genSet) {
			return option.Some(ret)
		}
	}

	return option.None[lookupTarget]()
}

func (r TypeClassSummonContext) namedLookup(f metafp.TypeInfo, name string) typeClassInstance {
	ret := r.lookupTypeClassInstanceLocalDeclared(f, name).Or(lazy.Func2(r.lookupTypeClassInstanceTypePkg)(f, name)).Or(r.lookupTypeClassInstancePrimitivePkgLazy(f, name))

	return typeClassInstance{
		ret,
		r.typeclassInstanceMust(f, name),
	}
}

func (r TypeClassSummonContext) lookupPrimitiveTypeClassInstance(f metafp.TypeInfo, name ...string) typeClassInstance {
	ret := r.lookupTypeClassInstanceLocalDeclared(f, name...).Or(r.lookupTypeClassInstancePrimitivePkgLazy(f, name...))

	return typeClassInstance{
		ret,
		r.typeclassInstanceMust(f, name[0]),
	}
}

func (r TypeClassSummonContext) exprTypeClassInstance(lt lookupTarget) string {
	if len(lt.required) > 0 {
		list := seq.Map(lt.required, func(t metafp.RequiredInstance) string {
			return r.summon(t.Type)
		}).MakeString(",")
		return fmt.Sprintf("%s(%s)", lt.instanceExpr(r.w, r.tc.Package), list)
	}

	if lt.isFunc() && len(lt.required) == 0 {
		return fmt.Sprintf("%s[%s]()", lt.instanceExpr(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, lt.instanceOf.Type))
	}

	return lt.instanceExpr(r.w, r.tc.Package)

}

func (r TypeClassSummonContext) exprTypeClassMember(lt typeClassMember, typeArgs fp.Seq[metafp.TypeInfo]) string {
	if len(typeArgs) > 0 {
		list := seq.Map(typeArgs, func(t metafp.TypeInfo) string {
			return r.summon(t)
		}).MakeString(",")
		return fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, r.tc.Package), list)
	}

	return lt.PackagedName(r.w, r.tc.Package)

}

func (r TypeClassSummonContext) exprTypeClassMemberLabelled(lt typeClassMember, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfo]) string {
	if len(typeArgs) > 0 {
		list := seq.Map(seq.Zip(typeArgs, names), func(t fp.Tuple2[metafp.TypeInfo, string]) string {
			return r.summonNamed(t.I2, t.I1)
		}).MakeString(",")
		return fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, r.tc.Package), list)
	}

	return lt.PackagedName(r.w, r.tc.Package)

}

type typeClassInstance struct {
	available fp.Option[lookupTarget]
	must      lookupTarget
}

func newTypeClassInstance(t lookupTarget) typeClassInstance {
	return typeClassInstance{
		available: option.Some(t),
		must:      t,
	}
}

func (r TypeClassSummonContext) lookupTypeClassInstance(f metafp.TypeInfo) typeClassInstance {
	switch at := f.Type.(type) {
	case *types.TypeParam:
		return newTypeClassInstance(lookupTarget{
			instanceOf: f,
			name:       privateName(r.tc.TypeClass.Name) + at.Obj().Name(),
			typeParam:  true,
		})
	case *types.Named:
		if at.Obj().Pkg().Path() == "github.com/csgura/fp/hlist" {
			if at.Obj().Name() == "Nil" {
				return typeClassInstance{r.lookupTypeClassInstanceLocalDeclared(f, "HNil", "HListNil").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(f, "HNil", "HListNil")),
					r.typeclassInstanceMust(f, "HNil"),
				}

			} else if at.Obj().Name() == "Cons" {
				return typeClassInstance{
					r.lookupTypeClassInstanceLocalDeclared(f, "HCons", "HListCons").
						Or(r.lookupTypeClassInstancePrimitivePkgLazy(f, "HCons", "HListCons")),

					r.typeclassInstanceMust(f, "HCons"),
				}
			}
		}
		return r.namedLookup(f, at.Obj().Name())
	case *types.Array:
		panic(fmt.Sprintf("can't summon array type, while deriving %s[%s]", r.tc.TypeClass.Name, r.tc.DeriveFor.Name))
		//return r.namedLookup(f, "Array")
	case *types.Slice:
		if at.Elem().String() == "byte" {
			bytesInstance := r.namedLookup(metafp.TypeInfo{
				Pkg:      f.Pkg,
				Type:     f.Type,
				TypeArgs: nil,
			}, "Bytes")
			if bytesInstance.available.IsDefined() {
				return bytesInstance
			}
			return r.namedLookup(f, "Slice")
		}
		return r.namedLookup(f, "Slice")
	case *types.Map:
		return r.namedLookup(f, "GoMap")
	case *types.Pointer:
		return r.namedLookup(f, "Ptr")
	case *types.Basic:
		return r.namedLookup(f, at.Name())
	case *types.Struct:
		panic(fmt.Sprintf("can't summon unnamed struct type, while deriving %s[%s]", r.tc.TypeClass.Name, r.tc.DeriveFor.Name))
	case *types.Interface:
		panic(fmt.Sprintf("can't summon unnamed interface type, while deriving %s[%s]", r.tc.TypeClass.Name, r.tc.DeriveFor.Name))
	case *types.Chan:
		panic(fmt.Sprintf("can't summon unnamed chan type, while deriving %s[%s]", r.tc.TypeClass.Name, r.tc.DeriveFor.Name))

	}
	return r.namedLookup(f, f.Type.String())
}

type TypeClassSummonContext struct {
	w       metafp.Writer
	tc      metafp.TypeClassDerive
	genSet  mutable.Set[string]
	tcCache *metafp.TypeClassInstanceCache
}

type GenericRepr struct {
	//	ReprType     func() string
	ToReprExpr   func() string
	FromReprExpr func() string
	ReprExpr     func() string
}

func (r TypeClassSummonContext) summonLabelledGenericRepr(receiver string, receiverType string, builderreceiver string, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfo]) fp.Option[GenericRepr] {
	result := r.lookupTypeClassMember(fmt.Sprintf("Labelled%d", typeArgs.Size()))

	return option.Map(result, func(tm typeClassMember) GenericRepr {
		return GenericRepr{
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr: func() string {
				return fmt.Sprintf("%s.AsLabelled", receiver)
			},
			FromReprExpr: func() string {
				fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

				return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromLabelled)(%s{}),
					%s.Build,
					)`,
					fppk,
					aspk, builderreceiver, builderreceiver, builderreceiver,
				)
			},
			ReprExpr: func() string {
				return r.exprTypeClassMemberLabelled(tm, names, typeArgs)
			},
		}
	}).Or(func() fp.Option[GenericRepr] {
		return option.Map(r.lookupTypeClassMember("HConsLabelled"), func(hcons typeClassMember) GenericRepr {
			return GenericRepr{
				// ReprType: func() string {
				// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
				// },
				ToReprExpr: func() string {
					fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
					aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

					namedTypeArgs := seq.Zip(names, typeArgs)

					tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
						return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
					}).MakeString(",")

					return fmt.Sprintf(`%s.Compose(
						%s.AsLabelled,
						%s.HList%dLabelled[%s],
					)`, fppk,
						receiver,
						aspk, typeArgs.Size(), tp,
					)

				},
				FromReprExpr: func() string {
					fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
					aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
					hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

					namedTypeArgs := seq.Zip(names, typeArgs)

					tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
						return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
					}).MakeString(",")

					hlisttp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
						return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
					}).MakeString(",")

					hlistToTuple := fmt.Sprintf(`%s.Func2(
						%s.Case%d[%s,%s.Nil,fp.Labelled%d[%s]],
					).ApplyLast( 
						%s.Labelled%d[%s] ,
					)`,
						aspk,
						hlistpk, typeArgs.Size(), hlisttp, hlistpk, typeArgs.Size(), tp,
						aspk, typeArgs.Size(), tp,
					)

					tupleToStruct := fmt.Sprintf(`%s.Compose(
						%s.Curried2(%s.FromLabelled)(%s{}),
						%s.Build,
						)`,
						fppk,
						aspk, builderreceiver, builderreceiver, builderreceiver,
					)
					return fmt.Sprintf(`
						fp.Compose(
							%s, 
							%s ,
						)`, hlistToTuple, tupleToStruct)
				},
				ReprExpr: func() string {
					hnil := r.lookupHNilMust()
					namedTypeArgs := seq.Zip(names, typeArgs)
					hlist := seq.Fold(namedTypeArgs.Reverse(), hnil.PackagedName(r.w, r.tc.Package), func(tail string, ti fp.Tuple2[string, metafp.TypeInfo]) string {
						instance := r.summonNamed(ti.Unapply())
						return fmt.Sprintf(`%s(
							%s,
						%s,
						)`, hcons.PackagedName(r.w, r.tc.Package), instance, tail)
					})

					return hlist
				},
			}
		})
	})
}

func (r TypeClassSummonContext) summonGenericRepr(receiver string, receiverType string, builderreceiver string, typeArgs fp.Seq[metafp.TypeInfo]) GenericRepr {
	result := r.lookupTypeClassMember(fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {

		// tp := iterator.Map(typeArgs.Iterator(), func(v metafp.TypeInfo) string {
		// 	return r.w.TypeName(r.tc.Package, v.Type)
		// }).MakeString(",")
		return GenericRepr{
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr: func() string {
				return fmt.Sprintf("%s.AsTuple", receiver)
			},
			FromReprExpr: func() string {
				fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

				return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromTuple)(%s{}),
					%s.Build,
					)`,
					fppk,
					aspk, builderreceiver, builderreceiver, builderreceiver,
				)
			},
			ReprExpr: func() string {
				return r.exprTypeClassMember(result.Get(), typeArgs)
			},
		}
	}

	tupleGeneric := r.summonTupleGenericRepr(typeArgs)

	return GenericRepr{
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {
			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

			return fmt.Sprintf(`%s.Compose(
				%s.AsTuple,
				%s, 
			)`, fppk,
				receiver,
				tupleGeneric.ToReprExpr(),
			)

		},
		FromReprExpr: func() string {
			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			tupleToStruct := fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromTuple)(%s{}),
					%s.Build,
					)`,
				fppk,
				aspk, builderreceiver, builderreceiver, builderreceiver,
			)
			return fmt.Sprintf(`
				fp.Compose(
					%s, 
					%s ,
				)`, tupleGeneric.FromReprExpr(), tupleToStruct)
		},
		ReprExpr: func() string {
			return tupleGeneric.ReprExpr()
		},
	}
}
func (r TypeClassSummonContext) summonTupleGenericRepr(typeArgs fp.Seq[metafp.TypeInfo]) GenericRepr {
	return GenericRepr{
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(r.tc.Package, f.Type)
			}).MakeString(",")

			return fmt.Sprintf(`%s.HList%d[%s]`,
				aspk, typeArgs.Size(), tp,
			)

		},
		FromReprExpr: func() string {
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
			hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

			tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(r.tc.Package, f.Type)
			}).MakeString(",")

			hlistToTuple := fmt.Sprintf(`%s.Func2(
					%s.Case%d[%s,%s.Nil,fp.Tuple%d[%s]],
				).ApplyLast( 
					%s.Tuple%d[%s] ,
				)`,
				aspk, hlistpk, typeArgs.Size(), tp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
			)

			return hlistToTuple
		},
		ReprExpr: func() string {
			hcons := r.lookupTypeClassMemberMust("HCons")

			hnil := r.lookupHNilMust()

			hlist := seq.Fold(typeArgs.Reverse(), hnil.PackagedName(r.w, r.tc.Package), func(tail string, ti metafp.TypeInfo) string {
				instance := r.summon(ti)
				return fmt.Sprintf(`%s(
					%s,
					%s,
				)`, hcons.PackagedName(r.w, r.tc.Package), instance, tail)
			})
			return hlist
		},
	}
}

func (r TypeClassSummonContext) summonTuple(typeArgs fp.Seq[metafp.TypeInfo]) string {

	result := r.lookupTypeClassMember(fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {
		return r.exprTypeClassMember(result.Get(), typeArgs)
	}

	tupleGeneric := r.summonTupleGenericRepr(typeArgs)

	if generic := r.lookupTypeClassMember("Generic"); generic.IsDefined() {
		aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		return fmt.Sprintf(`%s(%s.Generic("", %s,%s), %s)`, generic.Get().PackagedName(r.w, r.tc.Package), aspk, tupleGeneric.ToReprExpr(), tupleGeneric.FromReprExpr(), tupleGeneric.ReprExpr())
	}

	if imap := r.lookupTypeClassMember("IMap"); imap.IsDefined() {
		return fmt.Sprintf(`%s(
			%s , 
			%s, 
			%s,
			)`,
			imap.Get().PackagedName(r.w, r.tc.Package), tupleGeneric.ReprExpr(), tupleGeneric.FromReprExpr(), tupleGeneric.ToReprExpr())
	}

	if functor := r.lookupTypeClassMember("Map"); functor.IsDefined() {
		return fmt.Sprintf(`%s(
			%s , 
			%s,
			)`,
			functor.Get().PackagedName(r.w, r.tc.Package), tupleGeneric.ReprExpr(), tupleGeneric.FromReprExpr(),
		)
	}

	contrmap := r.lookupTypeClassMemberMust("ContraMap")
	return fmt.Sprintf(`%s( 
		%s, 
		%s,
		)`, contrmap.PackagedName(r.w, r.tc.Package), tupleGeneric.ReprExpr(), tupleGeneric.ToReprExpr())
}

func (r TypeClassSummonContext) summonNamed(name string, t metafp.TypeInfo) string {

	instance := r.lookupTypeClassMemberMust("Named")

	return fmt.Sprintf("%s[NameIs%s[%s]](%s)", instance.PackagedName(r.w, r.tc.Package), publicName(name),
		r.w.TypeName(r.tc.Package, t.Type), r.summon(t))

	// pk := r.w.GetImportedName(r.tc.Package)
	// return fmt.Sprintf("%s.Named(%s)", pk, r.summon(t))
}

func (r TypeClassSummonContext) summon(t metafp.TypeInfo) string {

	if t.IsTuple() {
		return r.summonTuple(t.TypeArgs)
	}

	result := r.lookupTypeClassInstance(t)

	if result.available.IsDefined() {
		return r.exprTypeClassInstance(result.available.Get())
	}

	instance := r.lookupTypeClassMember("UInt")
	if instance.IsDefined() {
		if _, ok := instance.Get().Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
			if err == nil {
				return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	instance = r.lookupTypeClassMember("Int")
	if instance.IsDefined() {
		if _, ok := instance.Get().Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
			if err == nil {
				return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	instance = r.lookupTypeClassMember("Float")
	if instance.IsDefined() {
		if _, ok := instance.Get().Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
			if err == nil {
				return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	instance = r.lookupTypeClassMember("Number")
	if instance.IsDefined() {
		if _, ok := instance.Get().Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
			if err == nil {
				return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	instance = r.lookupTypeClassMember("Given")
	if instance.IsDefined() {
		if _, ok := instance.Get().Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
			if err == nil {
				return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	return r.exprTypeClassInstance(result.must)

}

func genDerive() {
	pack := os.Getenv("GOPACKAGE")

	metafp.Generate(pack, pack+"_derive_generated.go", func(w metafp.Writer) {

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

		// fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))
		// asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		d := metafp.FindTypeClassDerive(pkgs)

		if d.Size() == 0 {
			return
		}

		tccache := metafp.TypeClassInstanceCache{}

		genSet := iterator.ToGoSet(iterator.Map(d.Iterator(), func(v metafp.TypeClassDerive) string {
			tccache.WillGenerated(v)
			return fmt.Sprintf("%s", v.GeneratedInstanceName())
		}))

		d.Foreach(func(v metafp.TypeClassDerive) {

			workingPackage := v.Package

			// fmt.Printf("lookup %s.Option = %v\n", v.Generator.Name(), l)
			//fmt.Printf("derive %v for %v\n", v.TypeClass, v.DeriveFor)
			privateFields := v.DeriveFor.Fields.FilterNot(metafp.StructField.Public)

			names := seq.Map(privateFields, func(v metafp.StructField) string {
				return v.Name
			})

			typeArgs := seq.Map(privateFields, func(v metafp.StructField) metafp.TypeInfo {
				return v.Type
			})

			summonCtx := TypeClassSummonContext{
				w:       w,
				tc:      v,
				genSet:  genSet,
				tcCache: &tccache,
			}

			valuetpdec := ""
			valuetp := ""
			if v.DeriveFor.Info.TypeParam.Size() > 0 {
				valuetpdec = "[" + iterator.Map(v.DeriveFor.Info.TypeParam.Iterator(), func(v metafp.TypeParam) string {
					tn := w.TypeName(workingPackage, v.Constraint)
					return fmt.Sprintf("%s %s", v.Name, tn)
				}).MakeString(",") + "]"

				valuetp = "[" + iterator.Map(v.DeriveFor.Info.TypeParam.Iterator(), func(v metafp.TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			builderreceiver := fmt.Sprintf("%sBuilder%s", v.DeriveFor.PackagedName(w, workingPackage), valuetp)
			valuereceiver := fmt.Sprintf("%s%s", v.DeriveFor.PackagedName(w, workingPackage), valuetp)

			labelledExpr := summonCtx.summonLabelledGenericRepr(valuereceiver, valuetp, builderreceiver, names, typeArgs)
			summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
				return summonCtx.summonGenericRepr(valuereceiver, valuetp, builderreceiver, typeArgs)
			})

			mapExpr := option.Map(summonCtx.lookupTypeClassMember("Generic"), func(generic typeClassMember) string {

				aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
				return fmt.Sprintf(`%s(
					%s.Generic(
							"%s.%s",
							%s,
							%s,
						), 
						%s, 
					)`, generic.PackagedName(w, workingPackage), aspk,
					pack, v.DeriveFor.Name,
					summonExpr.ToReprExpr(),
					summonExpr.FromReprExpr(),
					summonExpr.ReprExpr())
			}).Or(func() fp.Option[string] {
				return option.Map(summonCtx.lookupTypeClassMember("IMap"), func(imapfunc typeClassMember) string {

					return fmt.Sprintf(`%s( 
						%s, 
						%s , 
						%s,
						)`,
						imapfunc.PackagedName(w, workingPackage), summonExpr.ReprExpr(), summonExpr.FromReprExpr(), summonExpr.ToReprExpr())
				})
			}).Or(func() fp.Option[string] {
				functor := summonCtx.lookupTypeClassMember("Map")
				return option.Map(functor, func(v typeClassMember) string {

					return fmt.Sprintf(`%s( 
						%s, 
						%s,
						)`,
						v.PackagedName(w, workingPackage), summonExpr.ReprExpr(), summonExpr.FromReprExpr(),
					)
				})

			}).OrElseGet(func() string {
				contrmap := summonCtx.lookupTypeClassMemberMust("ContraMap")

				return fmt.Sprintf(`%s( 
					%s , 
					%s,
					)`,
					contrmap.PackagedName(w, workingPackage), summonExpr.ReprExpr(), summonExpr.ToReprExpr(),
				)
			})

			if v.DeriveFor.Info.TypeParam.Size() > 0 {

				tcname := v.TypeClass.PackagedName(w, workingPackage)
				fargs := seq.Map(v.DeriveFor.Info.TypeParam, func(p metafp.TypeParam) string {
					return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
				}).MakeString(",")

				fmt.Fprintf(w, `
					func %s%s( %s ) %s[%s%s] {
						return %s
					}
					`, v.GeneratedInstanceName(), valuetpdec, fargs, tcname, v.DeriveFor.PackagedName(w, workingPackage), valuetp,
					mapExpr)

			} else {
				fmt.Fprintf(w, `
				var %s = %s
			`, v.GeneratedInstanceName(), mapExpr)
			}

		})
	})
}
func main() {
	pack := os.Getenv("GOPACKAGE")
	if pack == "" {
		fmt.Println("invalid package. please run gombok using go generate command")
		return
	}
	genValue()
	genDerive()

}
