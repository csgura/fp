package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"strings"
	"unicode"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/generator/common"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"

	// . "github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

type StructField struct {
	Name string
	Type TypeInfo
	Tag  string
}

func (r StructField) Public() bool {
	return unicode.IsUpper([]rune(r.Name)[0])
}

type TypeParam struct {
	Name       string
	Constraint types.Type
}
type TypeInfo struct {
	Pkg       *types.Package
	Type      types.Type
	TypeArgs  fp.Seq[TypeInfo]
	TypeParam fp.Seq[TypeParam]
	Method    fp.Map[string, *types.Func]
}

func (r TypeInfo) IsTypeParam() bool {
	switch r.Type.(type) {
	case *types.TypeParam:
		return true
	}
	return false
}

func (r TypeInfo) IsBasic() bool {
	switch r.Type.(type) {
	case *types.Basic:
		return true
	}
	return false
}

func (r TypeInfo) IsSlice() bool {
	switch r.Type.(type) {
	case *types.Slice:
		return true
	}
	return false
}

func (r TypeInfo) IsArray() bool {
	switch r.Type.(type) {
	case *types.Array:
		return true
	}
	return false
}

func (r TypeInfo) IsMap() bool {
	switch r.Type.(type) {
	case *types.Map:
		return true
	}
	return false
}

func (r TypeInfo) IsPtr() bool {
	switch r.Type.(type) {
	case *types.Pointer:
		return true
	}
	return false
}

func (r TypeInfo) IsTuple() bool {
	switch nt := r.Type.(type) {
	case *types.Named:
		if nt.Obj().Pkg().Path() == "github.com/csgura/fp" && strings.HasPrefix(nt.Obj().Name(), "Tuple") {
			return true
		}
	}
	return false
}

type ValueType struct {
	Name      string
	Scope     *types.Scope
	Package   *types.Package
	Struct    *types.Struct
	Fields    fp.Seq[StructField]
	TypeParam fp.Seq[TypeParam]
	Method    fp.Map[string, *types.Func]
}

type TypeClass struct {
	Name    string
	Package *types.Package
}

func (r TypeClass) expr(w common.Writer, pk *types.Package) string {
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
	DeriveFor            ValueType
}

func lookupValueType(pk *types.Package, name string) fp.Option[ValueType] {
	l := pk.Scope().Lookup(name)
	if st, ok := l.Type().Underlying().(*types.Struct); ok {
		fl := iterator.Map(iterator.Range(0, st.NumFields()), func(i int) StructField {
			f := st.Field(i)
			tn := typeInfo(l.Pkg(), f.Type())
			return StructField{
				Name: f.Name(),
				Type: tn,
				Tag:  st.Tag(i),
			}
		}).ToSeq()

		info := typeInfo(l.Pkg(), l.Type())

		return option.Some(ValueType{
			Name:      name,
			Scope:     l.Parent(),
			Package:   l.Pkg(),
			Struct:    st,
			Fields:    fl,
			TypeParam: info.TypeParam,
			Method:    info.Method,
		})
	}
	return option.None[ValueType]()
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
									vt := lookupValueType(deriveFor.Obj().Pkg(), deriveFor.Obj().Name())
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
func findValueStruct(p []*packages.Package) fp.Seq[ValueType] {

	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[ValueType] {
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

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[ValueType] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[ValueType] {

				if ts, ok := v.(*ast.TypeSpec); ok {
					doc := option.Map(option.Of(ts.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)

					if doc.Filter(as.Func2(strings.Contains).ApplyLast("@fp.Value")).IsDefined() {

						l := pk.Types.Scope().Lookup(ts.Name.Name)
						if st, ok := l.Type().Underlying().(*types.Struct); ok {
							fl := iterator.Map(iterator.Range(0, st.NumFields()), func(i int) StructField {
								f := st.Field(i)
								tn := typeInfo(l.Pkg(), f.Type())

								return StructField{
									Name: f.Name(),
									Type: tn,
									Tag:  st.Tag(i),
								}
							}).ToSeq()

							info := typeInfo(l.Pkg(), l.Type())

							return seq.Of(ValueType{
								Name:      ts.Name.Name,
								Scope:     l.Parent(),
								Package:   l.Pkg(),
								Struct:    st,
								Fields:    fl,
								TypeParam: info.TypeParam,
								Method:    info.Method,
							})
						}
					}

				}
				return seq.Of[ValueType]()
			})
		})
	})

}

func typeInfo(pk *types.Package, tpe types.Type) TypeInfo {
	switch realtp := tpe.(type) {
	case *types.TypeParam:
		return TypeInfo{
			Type: tpe,
		}
	case *types.Named:
		args := fp.Seq[TypeInfo]{}
		params := fp.Seq[TypeParam]{}

		if realtp.TypeArgs() != nil {
			args = iterator.Map(iterator.Range(0, realtp.TypeArgs().Len()), func(i int) TypeInfo {
				return typeInfo(pk, realtp.TypeArgs().At(i))
			}).ToSeq()

		}

		if realtp.TypeParams() != nil {
			params = iterator.Map(iterator.Range(0, realtp.TypeParams().Len()), func(i int) TypeParam {
				return TypeParam{
					Name:       realtp.TypeParams().At(i).Obj().Name(),
					Constraint: realtp.TypeParams().At(i).Constraint(),
				}
			}).ToSeq()
		}

		method := iterator.Map(iterator.Range(0, realtp.NumMethods()), func(v int) fp.Tuple2[string, *types.Func] {
			m := realtp.Method(v)
			return as.Tuple2(m.Name(), m)
		})

		return TypeInfo{
			Pkg:       realtp.Obj().Pkg(),
			Type:      tpe,
			TypeArgs:  args,
			TypeParam: params,
			Method:    mutable.MapOf(iterator.ToGoMap(method)),
		}
	case *types.Array:
		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(pk, realtp.Elem())},
		}
	case *types.Map:

		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(pk, realtp.Key()), typeInfo(pk, realtp.Elem())},
		}
	case *types.Slice:
		return TypeInfo{
			Type:     tpe,
			TypeArgs: []TypeInfo{typeInfo(pk, realtp.Elem())},
		}
	}

	return TypeInfo{
		Type: tpe,
	}
}

func publicName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func privateName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func genValue() {
	pack := os.Getenv("GOPACKAGE")

	common.Generate(pack, "value_generated.go", func(w common.Writer) {

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

		st := findValueStruct(pkgs)

		st.Foreach(func(v ValueType) {

			valuetpdec := ""
			valuetp := ""
			if len(v.TypeParam) > 0 {
				valuetpdec = "[" + iterator.Map(v.TypeParam.Iterator(), func(v TypeParam) string {
					tn := w.TypeName(workingPackage, v.Constraint)
					return fmt.Sprintf("%s %s", v.Name, tn)
				}).MakeString(",") + "]"

				valuetp = "[" + iterator.Map(v.TypeParam.Iterator(), func(v TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			builderType := fmt.Sprintf("%sBuilder%s", v.Name, valuetpdec)
			mutableType := fmt.Sprintf("%sMutable%s", v.Name, valuetpdec)

			//valueType := fmt.Sprintf("%s%s", v.Name, valuetpdec)

			valuereceiver := fmt.Sprintf("%s%s", v.Name, valuetp)
			builderreceiver := fmt.Sprintf("%sBuilder%s", v.Name, valuetp)
			mutablereceiver := fmt.Sprintf("%sMutable%s", v.Name, valuetp)

			fmt.Fprintf(w, `
				type %s %s
			`, builderType, valuereceiver)

			mutableFields := iterator.Map(v.Fields.Iterator(), func(v StructField) string {
				if v.Tag != "" {
					return fmt.Sprintf("%s %s `%s`", publicName(v.Name), w.TypeName(workingPackage, v.Type.Type), v.Tag)
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

			privateFields := v.Fields.FilterNot(StructField.Public)
			privateFields.Foreach(func(f StructField) {

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
			})

			fm := seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf("%s=%%v", f.Name)
			}).Iterator().MakeString(", ")

			fields := seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Iterator().MakeString(",")

			m := v.Method.Get("String")

			if m.IsEmpty() {
				fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))

				fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, valuereceiver,
					fmtalias, v.Name, fm, fields,
				)
			}

			tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v StructField) string {
				return w.TypeName(workingPackage, v.Type.Type)
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f StructField) string {
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

			fields = seq.Map(privateFields, func(f StructField) string {
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

			fields = seq.Map(v.Fields, func(f StructField) string {
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

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, StructField]) string {
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

			fields = seq.Map(privateFields, func(f StructField) string {
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

			fields = iterator.Map(privateFields.Iterator(), func(f StructField) string {
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

			tp = iterator.Map(privateFields.Iterator().Take(max.Product), func(v StructField) string {
				return fmt.Sprintf("%s.Tuple2[string,%s]", fppkg, w.TypeName(workingPackage, v.Type.Type))
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf(`%s.Tuple2("%s", r.%s)`, asalias, f.Name, f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsLabelled() %s.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, fppkg, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, StructField]) string {
				return fmt.Sprintf("r.%s = t.I%d.I2", f.I2.Name, f.I1+1)
			}).Take(max.Product).MakeString("\n")

			fmt.Fprintf(w, `
					func (r %s) FromLabelled(t %s.Tuple%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, fppkg, privateFields.Size(), tp, builderreceiver,
				fields,
			)

		})

	})
}

type TypeClassInstance struct {
	name     string
	instance types.Object
}

type lookupTarget struct {
	name     string
	scope    *types.Scope
	typeArgs fp.Seq[TypeInfo]
	genPk    *types.Package
	tc       *TypeClass
}

func (r lookupTarget) available(genSet mutable.Set[string]) bool {
	if r.tc != nil {
		return true
	}
	ins := r.scope.Lookup(r.name)
	if ins != nil {
		return true
	}

	if r.genPk == nil && genSet.Contains(r.name) {
		return true
	}
	return false
}

func (r lookupTarget) instanceExpr(w common.Writer) string {
	if r.genPk != nil {

		if r.tc != nil {
			//pk := w.GetImportedName(r.tc.Package)

			return fmt.Sprintf("%s%s", privateName(r.tc.Name), r.name)
		}

		pk := w.GetImportedName(r.genPk)

		return fmt.Sprintf("%s.%s", pk, r.name)
	}
	return r.name
}

func (r TypeClassSummonContext) lookupLocal(f TypeInfo, name string) fp.Seq[lookupTarget] {

	var ret []lookupTarget

	if f.Pkg != nil && r.tc.Package.Path() != f.Pkg.Path() {
		ret = append(ret, lookupTarget{
			r.tc.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(name),
			r.tc.Package.Scope(),
			f.TypeArgs,
			nil,
			nil,
		})
	}

	ret = append(ret, lookupTarget{
		r.tc.TypeClass.Name + publicName(name),
		r.tc.Package.Scope(),
		f.TypeArgs,
		nil,
		nil,
	})

	return ret
}

func (r TypeClassSummonContext) lookupPrimitiveInstancePkg(f TypeInfo, name string) fp.Seq[lookupTarget] {

	var ret []lookupTarget

	if f.Pkg != nil {
		ret = append(ret, lookupTarget{
			r.tc.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(name),
			r.tc.PrimitiveInstancePkg.Scope(),
			f.TypeArgs,
			r.tc.PrimitiveInstancePkg,
			nil,
		})
	}

	ret = append(ret, lookupTarget{
		r.tc.TypeClass.Name + publicName(name),
		r.tc.PrimitiveInstancePkg.Scope(),
		f.TypeArgs,
		r.tc.PrimitiveInstancePkg,
		nil,
	})

	ret = append(ret, lookupTarget{
		publicName(name),
		r.tc.PrimitiveInstancePkg.Scope(),
		f.TypeArgs,
		r.tc.PrimitiveInstancePkg,
		nil,
	})

	return ret
}

func (r TypeClassSummonContext) lookupDeclaredPkg(f TypeInfo, name string) fp.Seq[lookupTarget] {

	var ret []lookupTarget

	if f.Pkg != nil && f.Pkg.Path() != r.tc.Package.Path() {
		ret = append(ret, lookupTarget{
			r.tc.TypeClass.Name + publicName(name),
			f.Pkg.Scope(),
			f.TypeArgs,
			f.Pkg,
			nil,
		})
	}

	return ret
}

func (r TypeClassSummonContext) namedLookup(f TypeInfo, name string) fp.Seq[lookupTarget] {
	return r.lookupLocal(f, name).
		Concat(r.lookupDeclaredPkg(f, name)).
		Concat(r.lookupPrimitiveInstancePkg(f, name))
}

func (r TypeClassSummonContext) expr(lt lookupTarget) string {
	if len(lt.typeArgs) > 0 {
		list := seq.Map(lt.typeArgs, func(t TypeInfo) string {
			return r.summon(t)
		}).MakeString(",")
		return fmt.Sprintf("%s(%s)", lt.instanceExpr(r.w), list)
	}

	return lt.instanceExpr(r.w)

}

func (r TypeClassSummonContext) implicitTypeClassInstanceName(f TypeInfo) fp.Seq[lookupTarget] {
	switch at := f.Type.(type) {
	case *types.TypeParam:
		return []lookupTarget{
			{
				name:  at.Obj().Name(),
				genPk: r.tc.TypeClass.Package,
				tc:    &r.tc.TypeClass,
			},
		}
	case *types.Named:
		if at.Obj().Pkg().Path() == "github.com/csgura/fp/hlist" {
			if at.Obj().Name() == "Nil" {
				return r.lookupLocal(f, "HNil").
					Concat(r.lookupLocal(f, "HListNil")).
					Concat(r.lookupPrimitiveInstancePkg(f, "HNil")).
					Concat(r.lookupPrimitiveInstancePkg(f, "HListNil"))

			} else if at.Obj().Name() == "Cons" {
				return r.lookupLocal(f, "HCons").
					Concat(r.lookupLocal(f, "HListCons")).
					Concat(r.lookupPrimitiveInstancePkg(f, "HCons")).
					Concat(r.lookupPrimitiveInstancePkg(f, "HListCons"))
			}
		}
		return r.namedLookup(f, at.Obj().Name())
	case *types.Array:
		return r.namedLookup(f, "Array")
	case *types.Slice:
		if at.Elem().String() == "byte" {
			return r.namedLookup(TypeInfo{
				Pkg:      f.Pkg,
				Type:     f.Type,
				TypeArgs: nil,
			}, "Bytes").Concat(r.namedLookup(f, "Slice"))
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
	w      common.Writer
	tc     TypeClassDerive
	genSet mutable.Set[string]
}

func (r TypeClassSummonContext) summonTuple(typeArgs fp.Seq[TypeInfo]) string {

	list := fp.Seq[lookupTarget]{
		{
			name:     fmt.Sprintf("%sTuple%d", r.tc.TypeClass.Name, typeArgs.Size()),
			scope:    r.tc.Package.Scope(),
			typeArgs: typeArgs,
		},
		{
			name:     fmt.Sprintf("%sTuple%d", r.tc.TypeClass.Name, typeArgs.Size()),
			scope:    r.tc.PrimitiveInstancePkg.Scope(),
			genPk:    r.tc.PrimitiveInstancePkg,
			typeArgs: typeArgs,
		},
		{
			name:     fmt.Sprintf("Tuple%d", typeArgs.Size()),
			scope:    r.tc.PrimitiveInstancePkg.Scope(),
			genPk:    r.tc.PrimitiveInstancePkg,
			typeArgs: typeArgs,
		},
	}

	result := list.Iterator().Filter(as.Func2(lookupTarget.available).ApplyLast(r.genSet)).First()

	if result.IsDefined() {
		return r.expr(result.Get())
	}

	aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

	gpk := r.w.GetImportedName(r.tc.PrimitiveInstancePkg)

	hlist := seq.Fold(typeArgs.Reverse(), fmt.Sprintf("%s.HNil", gpk), func(tail string, ti TypeInfo) string {
		instance := r.summon(ti)
		return fmt.Sprintf(`%s.HCons(%s,
			%s)`, gpk, instance, tail)
	})

	tp := seq.Map(typeArgs, func(f TypeInfo) string {
		return r.w.TypeName(r.tc.Package, f.Type)
	}).MakeString(",")

	if imap := r.tc.PrimitiveInstancePkg.Scope().Lookup("IMap"); imap != nil {
		hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

		return fmt.Sprintf(`%s.IMap(%s , 
			%s.Func2(%s.Case%d[%s,%s.Nil,fp.Tuple%d[%s]]).ApplyLast( %s.Tuple%d[%s] ),
			%s.HList%d[%s])`,
			gpk, hlist,
			aspk, hlistpk, typeArgs.Size(), tp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
			aspk, typeArgs.Size(), tp)
	}

	return fmt.Sprintf("%s.ContraMap( %s,  %s.HList%d[%s])", gpk, hlist, aspk, typeArgs.Size(), tp)
}

func (r TypeClassSummonContext) summon(t TypeInfo) string {

	if t.IsTuple() {
		return r.summonTuple(t.TypeArgs)
	}

	list := r.implicitTypeClassInstanceName(t)

	result := list.Iterator().Filter(as.Func2(lookupTarget.available).ApplyLast(r.genSet)).First()

	if result.IsDefined() {
		return r.expr(result.Get())
	}

	instance := r.tc.PrimitiveInstancePkg.Scope().Lookup("Number")
	if instance != nil {
		if _, ok := instance.Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Type(), []types.Type{t.Type}, true)
			if err == nil {
				gpk := r.w.GetImportedName(r.tc.PrimitiveInstancePkg)
				return fmt.Sprintf("%s.Number[%s]()", gpk, r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	instance = r.tc.PrimitiveInstancePkg.Scope().Lookup("Given")
	if instance != nil {
		if _, ok := instance.Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Type(), []types.Type{t.Type}, true)
			if err == nil {
				gpk := r.w.GetImportedName(r.tc.PrimitiveInstancePkg)
				return fmt.Sprintf("%s.Given[%s]()", gpk, r.w.TypeName(r.tc.Package, t.Type))
			}
		}
	}

	return list.Head().Get().instanceExpr(r.w)

}

func genDerive() {
	pack := os.Getenv("GOPACKAGE")

	common.Generate(pack, "derive_generated.go", func(w common.Writer) {

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
			privateFields := v.DeriveFor.Fields.FilterNot(StructField.Public)

			gpk := w.GetImportedName(v.PrimitiveInstancePkg)

			typeArgs := seq.Map(privateFields, func(v StructField) TypeInfo {
				return v.Type
			})

			summonCtx := TypeClassSummonContext{
				w:      w,
				tc:     v,
				genSet: genSet,
			}

			summonExpr := summonCtx.summonTuple(typeArgs)

			if v.DeriveFor.TypeParam.Size() > 0 {
				valuetpdec := "[" + iterator.Map(v.DeriveFor.TypeParam.Iterator(), func(v TypeParam) string {
					tn := w.TypeName(workingPackage, v.Constraint)
					return fmt.Sprintf("%s %s", v.Name, tn)
				}).MakeString(",") + "]"

				valuetp := "[" + iterator.Map(v.DeriveFor.TypeParam.Iterator(), func(v TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"

				tcname := v.TypeClass.expr(w, workingPackage)
				fargs := seq.Map(v.DeriveFor.TypeParam, func(p TypeParam) string {
					return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
				}).MakeString(",")

				if imap := v.PrimitiveInstancePkg.Scope().Lookup("IMap"); imap != nil {
					fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
					aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
					fmt.Fprintf(w, `
					func %s%s%s( %s ) %s[%s%s] {
						return %s.IMap( %s, %s.Compose(
							%s.Curried2(%sBuilder.FromTuple)(%sBuilder{}), %sBuilder.Build),  
							%s.AsTuple )
				}
			`, v.TypeClass.Name, v.DeriveFor.Name, valuetpdec, fargs, tcname, v.DeriveFor.Name, valuetp,
						gpk, summonExpr, fppk,
						aspk, v.DeriveFor.Name, v.DeriveFor.Name, v.DeriveFor.Name,
						v.DeriveFor.Name)
				} else {
					fmt.Fprintf(w, `
						func %s%s%s( %s ) %s[%s%s] {
							return %s.ContraMap( %s , %s%s.AsTuple )
						}
					`, v.TypeClass.Name, v.DeriveFor.Name, valuetpdec, fargs, tcname, v.DeriveFor.Name, valuetp,
						gpk, summonExpr, v.DeriveFor.Name, valuetp,
					)
				}
			} else {
				if imap := v.PrimitiveInstancePkg.Scope().Lookup("IMap"); imap != nil {
					fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
					aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
					fmt.Fprintf(w, `
				var %s%s = %s.IMap( %s, %s.Compose(
					%s.Curried2(%sBuilder.FromTuple)(%sBuilder{}), %sBuilder.Build),  
					%s.AsTuple )
			`, v.TypeClass.Name, v.DeriveFor.Name, gpk, summonExpr, fppk,
						aspk, v.DeriveFor.Name, v.DeriveFor.Name, v.DeriveFor.Name,
						v.DeriveFor.Name)
				} else {
					fmt.Fprintf(w, `
				var %s%s = %s.ContraMap( %s, %s.AsTuple )
			`, v.TypeClass.Name, v.DeriveFor.Name, gpk, summonExpr, v.DeriveFor.Name)
				}
			}

		})
	})
}
func main() {
	genValue()
	genDerive()

}
