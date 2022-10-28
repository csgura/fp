package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
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

type TypeClass struct {
	Name    string
	Package *types.Package
}

func (r TypeClass) expr(w metafp.Writer, pk *types.Package) string {
	if r.Package != nil && r.Package.Path() != pk.Path() {
		pk := w.GetImportedName(r.Package)
		return fmt.Sprintf("%s.%s", pk, r.Name)
	}
	return r.Name
}

type TypeClassDerive struct {
	Package              *types.Package
	PrimitiveInstancePkg *types.Package
	TypeClass            TypeClass
	DeriveFor            metafp.TaggedStruct
}

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

func findTypeClassDerive(p []*packages.Package) fp.Seq[TypeClassDerive] {
	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[TypeClassDerive] {
		s2 := seq.FlatMap(pk.Syntax, func(v *ast.File) fp.Seq[ast.Decl] {

			return v.Decls
		})

		s3 := seq.FlatMap(s2, func(v ast.Decl) fp.Seq[*ast.GenDecl] {
			switch r := v.(type) {
			case *ast.GenDecl:
				return seq.Of(r)
			}
			return seq.Of[*ast.GenDecl]()
		})

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[TypeClassDerive] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[TypeClassDerive] {
				if vs, ok := v.(*ast.ValueSpec); ok {
					doc := option.Map(option.Of(vs.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)
					if doc.Filter(as.Func2(strings.Contains).ApplyLast("@fp.Derive")).IsDefined() {

						info := &types.Info{
							Types: make(map[ast.Expr]types.TypeAndValue),
						}
						types.CheckExpr(pk.Fset, pk.Types, v.Pos(), vs.Type, info)
						ti := info.Types[vs.Type]
						if nt, ok := ti.Type.(*types.Named); ok && nt.TypeArgs().Len() == 1 {
							if tt, ok := nt.TypeArgs().At(0).(*types.Named); ok && tt.TypeArgs().Len() == 1 {
								if deriveFor, ok := tt.TypeArgs().At(0).(*types.Named); ok {

									//fmt.Printf("lookup %s from %s\n", deriveFor.Obj().Name(), deriveFor.Obj().Pkg())
									vt := metafp.LookupStruct(deriveFor.Obj().Pkg(), deriveFor.Obj().Name())
									if vt.IsDefined() {
										return seq.Of(TypeClassDerive{
											Package:              pk.Types,
											PrimitiveInstancePkg: nt.Obj().Pkg(),
											TypeClass: TypeClass{
												Name:    tt.Obj().Name(),
												Package: tt.Obj().Pkg(),
											},
											DeriveFor: vt.Get(),
										})
									} else {
										fmt.Println("can't lookup")
									}
								}

							}

						}
					}
				}
				return seq.Of[TypeClassDerive]()
			})
		})
	})
}

func publicName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func privateName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func genValue() {
	pack := os.Getenv("GOPACKAGE")

	metafp.Generate(pack, "value_generated.go", func(w metafp.Writer) {

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

		asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		st := metafp.FindTaggedStruct(pkgs, "@fp.Value")

		keyTags := mutable.EmptySet[string]()

		st.Foreach(func(ts metafp.TaggedStruct) {

			valuetpdec := ""
			valuetp := ""
			if len(ts.Info.TypeParam) > 0 {
				valuetpdec = "[" + ts.Info.TypeParamDecl(w, workingPackage) + "]"
				valuetp = "[" + ts.Info.TypeParamIns(w, workingPackage) + "]"
			}

			builderType := fmt.Sprintf("%sBuilder%s", ts.Name, valuetpdec)
			mutableType := fmt.Sprintf("%sMutable%s", ts.Name, valuetpdec)

			//valueType := fmt.Sprintf("%s%s", v.Name, valuetpdec)

			valuereceiver := fmt.Sprintf("%s%s", ts.Name, valuetp)
			builderreceiver := fmt.Sprintf("%sBuilder%s", ts.Name, valuetp)
			mutablereceiver := fmt.Sprintf("%sMutable%s", ts.Name, valuetp)

			fmt.Fprintf(w, `
				type %s %s
			`, builderType, valuereceiver)

			mutableFields := iterator.Map(ts.Fields.Iterator(), func(v metafp.StructField) string {
				tag := v.Tag

				if ts.Tags.Contains("@fp.JsonTag") || ts.Tags.Contains("@fp.Json") {
					if !strings.Contains(tag, "json") {
						if tag != "" {
							tag = tag + " "
						}
						tag = tag + fmt.Sprintf(`json:"%s"`, v.Name)
					}
				}
				if tag != "" {
					return fmt.Sprintf("%s %s `%s`", publicName(v.Name), w.TypeName(workingPackage, v.Type.Type), tag)
				} else {
					return fmt.Sprintf("%s %s", publicName(v.Name), w.TypeName(workingPackage, v.Type.Type))

				}
			}).MakeString("\n")
			fmt.Fprintf(w, `
				type %s struct {
					%s
				}
			`, mutableType, mutableFields)

			fmt.Fprintf(w, `
				func(r %s) Build() %s {
					return %s(r)
				}
			`, builderreceiver, valuereceiver, valuereceiver)

			fmt.Fprintf(w, `
				func(r %s) Builder() %s {
					return %s(r)
				}
			`, valuereceiver, builderreceiver, builderreceiver)

			privateFields := ts.Fields.FilterNot(metafp.StructField.Public)

			privateFields.Foreach(func(v metafp.StructField) {
				keyTags = keyTags.Incl(v.Name)
			})

			privateFields.Foreach(func(f metafp.StructField) {

				uname := strings.ToUpper(f.Name[:1]) + f.Name[1:]
				ftp := w.TypeName(workingPackage, f.Type.Type)

				fmt.Fprintf(w, `
						func (r %s) %s() %s {
							return r.%s
						}
					`, valuereceiver, uname, ftp, f.Name)

				fmt.Fprintf(w, `
						func (r %s) With%s(v %s) %s {
							r.%s = v
							return r
						}
					`, valuereceiver, uname, ftp, valuereceiver, f.Name)

				fmt.Fprintf(w, `
						func (r %s) %s( v %s) %s {
							r.%s = v
							return r
						}
					`, builderreceiver, uname, ftp, builderreceiver, f.Name)

				if f.Type.IsOption() {
					optiont := w.TypeName(workingPackage, f.Type.TypeArgs.Head().Get().Type)
					optionpk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/option", "option"))
					fmt.Fprintf(w, `
						func (r %s) Some%s(v %s) %s {
							r.%s = %s.Some(v)
							return r
						}
					`, builderreceiver, uname, optiont, builderreceiver,
						f.Name, optionpk)

					fmt.Fprintf(w, `
						func (r %s) None%s() %s {
							r.%s = %s.None[%s]()
							return r
						}
					`, builderreceiver, uname, builderreceiver,
						f.Name, optionpk, optiont)
				}

			})

			fm := seq.Map(privateFields, func(f metafp.StructField) string {
				return fmt.Sprintf("%s=%%v", f.Name)
			}).Iterator().MakeString(", ")

			fields := seq.Map(privateFields, func(f metafp.StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Iterator().MakeString(",")

			m := ts.Info.Method.Get("String")

			if m.IsEmpty() {
				fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))

				fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, valuereceiver,
					fmtalias, ts.Name, fm, fields,
				)
			}

			tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v metafp.StructField) string {
				return w.TypeName(workingPackage, v.Type.Type)
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f metafp.StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fppkg := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

			fmt.Fprintf(w, `
					func(r %s) AsTuple() %s.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, fppkg, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

			fields = seq.Map(privateFields, func(f metafp.StructField) string {
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

			fields = seq.Map(ts.Fields, func(f metafp.StructField) string {
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

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, metafp.StructField]) string {
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

			fields = seq.Map(privateFields, func(f metafp.StructField) string {
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

			fields = iterator.Map(privateFields.Iterator(), func(f metafp.StructField) string {
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

			tp = iterator.Map(privateFields.Iterator().Take(max.Product), func(v metafp.StructField) string {
				return fmt.Sprintf("NameIs%s[%s]", publicName(v.Name), w.TypeName(workingPackage, v.Type.Type))
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f metafp.StructField) string {
				return fmt.Sprintf(`NameIs%s[%s]{r.%s}`, publicName(f.Name), w.TypeName(workingPackage, f.Type.Type), f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsLabelled() %s.Labelled%d[%s] {
						return %s.Labelled%d(%s)
					}

				`, valuereceiver, fppkg, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

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

			if ts.Tags.Contains("@fp.Json") {
				jsonpk := w.GetImportedName(types.NewPackage("encoding/json", "json"))
				httppk := w.GetImportedName(types.NewPackage("net/http", "http"))
				fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

				fmt.Fprintf(w, `
					func(r %s) MarshalJSON() ([]byte, error) {
						m := r.AsMutable()
						return %s.Marshal(m)
					}
				`, valuereceiver, jsonpk)

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
	pk        *types.Package
	name      string
	typeArgs  fp.Seq[metafp.TypeInfo]
	typeParam bool
	// tc       *TypeClass
}

func (r lookupTarget) available(genSet mutable.Set[string]) bool {
	if r.typeParam {
		return true
	}
	ins := r.pk.Scope().Lookup(r.name)
	if ins != nil {
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
		pk:       r.tc.Package,
		name:     name,
		typeArgs: f.TypeArgs,
	}
}

func (r TypeClassSummonContext) lookupTypeClassInstanceLocalDeclared(f metafp.TypeInfo, name ...string) fp.Option[lookupTarget] {

	itr := seq.FlatMap(name, func(v string) fp.Seq[string] {
		if f.Pkg != nil && r.tc.Package.Path() != f.Pkg.Path() {
			return []string{
				r.tc.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(v),
				r.tc.TypeClass.Name + publicName(v),
			}

		}
		return []string{r.tc.TypeClass.Name + publicName(v)}
	}).Iterator()

	return iterator.Map(itr, func(v string) lookupTarget {
		return lookupTarget{
			pk:       r.tc.Package,
			name:     v,
			typeArgs: f.TypeArgs,
		}
	}).Filter(func(lt lookupTarget) bool {
		return lt.available(r.genSet)
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
			return seq.Of(r.tc.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(v)).
				Concat(ret)
		}
		return ret
	}).Iterator()

	return iterator.Map(itr, func(v string) lookupTarget {
		return lookupTarget{
			pk:       r.tc.PrimitiveInstancePkg,
			name:     v,
			typeArgs: f.TypeArgs,
		}
	}).Filter(func(lt lookupTarget) bool {
		return lt.available(r.genSet)
	}).Head()

}

func (r TypeClassSummonContext) lookupTypeClassInstanceTypePkg(f metafp.TypeInfo, name string) fp.Option[lookupTarget] {

	if f.Pkg != nil && f.Pkg.Path() != r.tc.Package.Path() {
		ret := lookupTarget{
			pk:       f.Pkg,
			name:     r.tc.TypeClass.Name + publicName(name),
			typeArgs: f.TypeArgs,
		}

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
	if len(lt.typeArgs) > 0 {
		list := seq.Map(lt.typeArgs, func(t metafp.TypeInfo) string {
			return r.summon(t)
		}).MakeString(",")
		return fmt.Sprintf("%s(%s)", lt.instanceExpr(r.w, r.tc.Package), list)
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
			name:      privateName(r.tc.TypeClass.Name) + at.Obj().Name(),
			typeParam: true,
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
		return r.namedLookup(f, "Array")
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
	}
	return r.namedLookup(f, f.Type.String())
}

type TypeClassSummonContext struct {
	w      metafp.Writer
	tc     TypeClassDerive
	genSet mutable.Set[string]
}

func (r TypeClassSummonContext) summonLabelled(names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfo]) fp.Option[string] {
	result := r.lookupTypeClassMember(fmt.Sprintf("Labelled%d", typeArgs.Size()))

	res := option.Map(result, as.Func3(r.exprTypeClassMemberLabelled).ApplyLast2(names, typeArgs))

	if res.IsDefined() {
		return res
	}

	result = r.lookupTypeClassMember("HConsLabelled")

	if result.IsDefined() {
		aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
		//fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		hnil := r.lookupHNilMust()
		namedTypeArgs := seq.Zip(names, typeArgs)
		hlist := seq.Fold(namedTypeArgs.Reverse(), hnil.PackagedName(r.w, r.tc.Package), func(tail string, ti fp.Tuple2[string, metafp.TypeInfo]) string {
			instance := r.summonNamed(ti.Unapply())
			return fmt.Sprintf(`%s(%s,
			%s)`, result.Get().PackagedName(r.w, r.tc.Package), instance, tail)
		})

		tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
			return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
		}).MakeString(",")

		hlisttp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
			return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
		}).MakeString(",")

		if imap := r.lookupTypeClassMember("IMap"); imap.IsDefined() {
			hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

			ret := fmt.Sprintf(`%s(%s , 
			%s.Func2(%s.Case%d[%s,%s.Nil,fp.Labelled%d[%s]]).ApplyLast( %s.Labelled%d[%s] ),
			%s.HList%dLabelled[%s])`,
				imap.Get().PackagedName(r.w, r.tc.Package), hlist,
				aspk, hlistpk, typeArgs.Size(), hlisttp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
				aspk, typeArgs.Size(), tp)

			return option.Some(ret)
		}

		if functor := r.lookupTypeClassMember("Map"); functor.IsDefined() {
			hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

			ret := fmt.Sprintf(`%s(%s , 
			%s.Func2(%s.Case%d[%s,%s.Nil,fp.Labelled%d[%s]]).ApplyLast( %s.Labelled%d[%s] ))`,
				functor.Get().PackagedName(r.w, r.tc.Package), hlist,
				aspk, hlistpk, typeArgs.Size(), hlisttp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
			)

			return option.Some(ret)
		}

		contrmap := r.lookupTypeClassMemberMust("ContraMap")

		ret := fmt.Sprintf("%s( %s,  %s.HList%dLabelled[%s])", contrmap.PackagedName(r.w, r.tc.Package), hlist, aspk, typeArgs.Size(), tp)
		return option.Some(ret)
	}
	return option.None[string]()
}

func (r TypeClassSummonContext) summonTuple(typeArgs fp.Seq[metafp.TypeInfo]) string {

	result := r.lookupTypeClassMember(fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {
		return r.exprTypeClassMember(result.Get(), typeArgs)
	}

	aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

	hcons := r.lookupTypeClassMemberMust("HCons")

	hnil := r.lookupHNilMust()

	hlist := seq.Fold(typeArgs.Reverse(), hnil.PackagedName(r.w, r.tc.Package), func(tail string, ti metafp.TypeInfo) string {
		instance := r.summon(ti)
		return fmt.Sprintf(`%s(%s,
			%s)`, hcons.PackagedName(r.w, r.tc.Package), instance, tail)
	})

	tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
		return r.w.TypeName(r.tc.Package, f.Type)
	}).MakeString(",")

	if imap := r.lookupTypeClassMember("IMap"); imap.IsDefined() {
		hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		return fmt.Sprintf(`%s(%s , 
			%s.Func2(%s.Case%d[%s,%s.Nil,fp.Tuple%d[%s]]).ApplyLast( %s.Tuple%d[%s] ),
			%s.HList%d[%s])`,
			imap.Get().PackagedName(r.w, r.tc.Package), hlist,
			aspk, hlistpk, typeArgs.Size(), tp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
			aspk, typeArgs.Size(), tp)
	}

	if functor := r.lookupTypeClassMember("Map"); functor.IsDefined() {
		hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		return fmt.Sprintf(`%s(%s , 
			%s.Func2(%s.Case%d[%s,%s.Nil,fp.Tuple%d[%s]]).ApplyLast( %s.Tuple%d[%s] ))`,
			functor.Get().PackagedName(r.w, r.tc.Package), hlist,
			aspk, hlistpk, typeArgs.Size(), tp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
		)
	}

	contrmap := r.lookupTypeClassMemberMust("ContraMap")

	return fmt.Sprintf("%s( %s,  %s.HList%d[%s])", contrmap.PackagedName(r.w, r.tc.Package), hlist, aspk, typeArgs.Size(), tp)
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

	instance := r.lookupTypeClassMember("Number")
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

	metafp.Generate(pack, "derive_generated.go", func(w metafp.Writer) {

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

		d := findTypeClassDerive(pkgs)

		genSet := iterator.ToGoSet(iterator.Map(d.Iterator(), func(v TypeClassDerive) string {
			return fmt.Sprintf("%s%s", v.TypeClass.Name, v.DeriveFor.Name)
		}))

		d.Foreach(func(v TypeClassDerive) {

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
				w:      w,
				tc:     v,
				genSet: genSet,
			}

			labelledExpr := summonCtx.summonLabelled(names, typeArgs)
			summonExpr := labelledExpr.OrElseGet(func() string {
				return summonCtx.summonTuple(typeArgs)
			})

			convExpr := option.Map(labelledExpr, func(v string) string {
				return "AsLabelled"
			}).OrElse("AsTuple")

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

			builderreceiver := fmt.Sprintf("%sBuilder%s", v.DeriveFor.Name, valuetp)
			valuereceiver := fmt.Sprintf("%s%s", v.DeriveFor.Name, valuetp)

			mapExpr := option.Map(summonCtx.lookupTypeClassMember("IMap"), func(imapfunc typeClassMember) string {
				fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

				revExpr := option.Map(labelledExpr, func(v string) string {
					return "FromLabelled"
				}).OrElse("FromTuple")

				return fmt.Sprintf(`%s( %s, %s.Compose(
							%s.Curried2(%s.%s)(%s{}), %s.Build),  
							%s.%s )`,
					imapfunc.PackagedName(w, workingPackage), summonExpr, fppk,
					aspk, builderreceiver, revExpr, builderreceiver, builderreceiver,
					valuereceiver, convExpr)
			}).Or(func() fp.Option[string] {
				functor := summonCtx.lookupTypeClassMember("Map")

				return option.Map(functor, func(v typeClassMember) string {
					fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
					aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

					revExpr := option.Map(labelledExpr, func(v string) string {
						return "FromLabelled"
					}).OrElse("FromTuple")

					return fmt.Sprintf(`%s( %s, %s.Compose(
							%s.Curried2(%s.%s)(%s{}), %s.Build))`,
						v.PackagedName(w, workingPackage), summonExpr, fppk,
						aspk, builderreceiver, revExpr, builderreceiver, builderreceiver,
					)
				})

			}).OrElseGet(func() string {
				contrmap := summonCtx.lookupTypeClassMemberMust("ContraMap")

				return fmt.Sprintf(`%s( %s , %s.%s )`,
					contrmap.PackagedName(w, workingPackage), summonExpr, valuereceiver, convExpr,
				)
			})

			if v.DeriveFor.Info.TypeParam.Size() > 0 {

				tcname := v.TypeClass.expr(w, workingPackage)
				fargs := seq.Map(v.DeriveFor.Info.TypeParam, func(p metafp.TypeParam) string {
					return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
				}).MakeString(",")

				fmt.Fprintf(w, `
					func %s%s%s( %s ) %s[%s%s] {
						return %s
					}
					`, v.TypeClass.Name, v.DeriveFor.Name, valuetpdec, fargs, tcname, v.DeriveFor.Name, valuetp,
					mapExpr)

			} else {
				fmt.Fprintf(w, `
				var %s%s = %s
			`, v.TypeClass.Name, v.DeriveFor.Name, mapExpr)
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
