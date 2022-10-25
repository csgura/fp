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
	Package   *types.Package
	Struct    *types.Struct
	Fields    fp.Seq[StructField]
	TypeParam fp.Seq[TypeParam]
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
	Package   *types.Package
	Generator *types.Package
	TypeClass TypeClass
	DeriveFor ValueType
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
			}
		}).ToSeq()

		info := typeInfo(l.Pkg(), l.Type())

		return option.Some(ValueType{
			Name:      name,
			Package:   l.Pkg(),
			Struct:    st,
			Fields:    fl,
			TypeParam: info.TypeParam,
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
											Package:   pk.Types,
											Generator: nt.Obj().Pkg(),
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
								}
							}).ToSeq()

							info := typeInfo(l.Pkg(), l.Type())

							return seq.Of(ValueType{
								Name:      ts.Name.Name,
								Package:   l.Pkg(),
								Struct:    st,
								Fields:    fl,
								TypeParam: info.TypeParam,
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
		return TypeInfo{
			Pkg:       realtp.Obj().Pkg(),
			Type:      tpe,
			TypeArgs:  args,
			TypeParam: params,
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

	common.Generate(pack, "value_generated.go", func(w common.Writer) {

		fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))
		asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		st := findValueStruct(pkgs)

		st.Foreach(func(v ValueType) {

			valuetpdec := ""
			valuetp := ""
			if len(v.TypeParam) > 0 {
				valuetpdec = "[" + iterator.Map(v.TypeParam.Iterator(), func(v TypeParam) string {
					tn := w.TypeName(pkgs[0].Types, v.Constraint)
					return fmt.Sprintf("%s %s", v.Name, tn)
				}).MakeString(",") + "]"

				valuetp = "[" + iterator.Map(v.TypeParam.Iterator(), func(v TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			builderType := fmt.Sprintf("%sBuilder%s", v.Name, valuetpdec)
			//valueType := fmt.Sprintf("%s%s", v.Name, valuetpdec)

			valuereceiver := fmt.Sprintf("%s%s", v.Name, valuetp)
			builderreceiver := fmt.Sprintf("%sBuilder%s", v.Name, valuetp)

			fmt.Fprintf(w, `
				type %s %s
			`, builderType, valuereceiver)

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
				ftp := w.TypeName(pkgs[0].Types, f.Type.Type)

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

			fmt.Fprintf(w, `
					func(r %s) String() string {
						return %s.Sprintf("%s(%s)", %s)
					}
				`, valuereceiver,
				fmtalias, v.Name, fm, fields,
			)

			tp := iterator.Map(privateFields.Iterator().Take(max.Product), func(v StructField) string {
				return w.TypeName(pkgs[0].Types, v.Type.Type)
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf("r.%s", f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsTuple() fp.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, StructField]) string {
				return fmt.Sprintf("r.%s = t.I%d", f.I2.Name, f.I1+1)
			}).Take(max.Product).MakeString("\n")

			fmt.Fprintf(w, `
					func (r %s) FromTuple(t fp.Tuple%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, privateFields.Size(), tp, builderreceiver,
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
						`, f.Name, w.TypeName(pkgs[0].Types, f.Type.Type),
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
				return fmt.Sprintf("fp.Tuple2[string,%s]", w.TypeName(pkgs[0].Types, v.Type.Type))
			}).MakeString(",")

			fields = seq.Map(privateFields, func(f StructField) string {
				return fmt.Sprintf(`as.Tuple2("%s", r.%s)`, f.Name, f.Name)
			}).Iterator().Take(max.Product).MakeString(",")

			fmt.Fprintf(w, `
					func(r %s) AsLabelled() fp.Tuple%d[%s] {
						return %s.Tuple%d(%s)
					}

				`, valuereceiver, privateFields.Size(), tp,
				asalias, privateFields.Size(), fields,
			)

			fields = iterator.Map(iterator.Zip(iterator.Range(0, privateFields.Size()), privateFields.Iterator()), func(f fp.Tuple2[int, StructField]) string {
				return fmt.Sprintf("r.%s = t.I%d.I2", f.I2.Name, f.I1+1)
			}).Take(max.Product).MakeString("\n")

			fmt.Fprintf(w, `
					func (r %s) FromLabelled(t fp.Tuple%d[%s] ) %s {
						%s
						return r
					}
				`, builderreceiver, privateFields.Size(), tp, builderreceiver,
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

func lookupLocal(pk *types.Package, tc TypeClass, f TypeInfo, name string) fp.Seq[lookupTarget] {

	var ret []lookupTarget

	if f.Pkg != nil && pk.Path() != f.Pkg.Path() {
		ret = append(ret, lookupTarget{
			tc.Name + publicName(f.Pkg.Name()) + publicName(name),
			pk.Scope(),
			f.TypeArgs,
			nil,
			nil,
		})
	}

	ret = append(ret, lookupTarget{
		tc.Name + publicName(name),
		pk.Scope(),
		f.TypeArgs,
		nil,
		nil,
	})

	return ret
}

func lookupGenerator(tc TypeClass, tcgen *types.Package, f TypeInfo, name string) fp.Seq[lookupTarget] {

	var ret []lookupTarget

	if f.Pkg != nil {
		ret = append(ret, lookupTarget{
			tc.Name + publicName(f.Pkg.Name()) + publicName(name),
			tcgen.Scope(),
			f.TypeArgs,
			tcgen,
			nil,
		})
	}

	ret = append(ret, lookupTarget{
		tc.Name + publicName(name),
		tcgen.Scope(),
		f.TypeArgs,
		tcgen,
		nil,
	})

	ret = append(ret, lookupTarget{
		publicName(name),
		tcgen.Scope(),
		f.TypeArgs,
		tcgen,
		nil,
	})

	return ret
}

func namedLookup(pk *types.Package, tc TypeClass, tcgen *types.Package, f TypeInfo, name string) fp.Seq[lookupTarget] {
	return lookupLocal(pk, tc, f, name).
		Concat(lookupGenerator(tc, tcgen, f, name))
}

func implicitTypeClassInstanceName(pk *types.Package, tc TypeClass, tcgen *types.Package, f TypeInfo) fp.Seq[lookupTarget] {
	switch at := f.Type.(type) {
	case *types.TypeParam:
		return []lookupTarget{
			{
				name:  at.Obj().Name(),
				genPk: tc.Package,
				tc:    &tc,
			},
		}
	case *types.Named:
		if at.Obj().Pkg().Path() == "github.com/csgura/fp/hlist" {
			if at.Obj().Name() == "Nil" {
				return lookupLocal(pk, tc, f, "HNil").
					Concat(lookupLocal(pk, tc, f, "HListNil")).
					Concat(lookupGenerator(tc, tcgen, f, "HNil")).
					Concat(lookupGenerator(tc, tcgen, f, "HListNil"))

			} else if at.Obj().Name() == "Cons" {
				return lookupLocal(pk, tc, f, "HCons").
					Concat(lookupLocal(pk, tc, f, "HListCons")).
					Concat(lookupGenerator(tc, tcgen, f, "HCons")).
					Concat(lookupGenerator(tc, tcgen, f, "HListCons"))
			}
		}
		return namedLookup(pk, tc, tcgen, f, at.Obj().Name())
	case *types.Array:
		return namedLookup(pk, tc, tcgen, f, "Array")
	case *types.Slice:
		if at.Elem().String() == "byte" {
			return namedLookup(pk, tc, tcgen, TypeInfo{
				Pkg:      f.Pkg,
				Type:     f.Type,
				TypeArgs: nil,
			}, "Bytes").Concat(namedLookup(pk, tc, tcgen, f, "Slice"))
		}
		return namedLookup(pk, tc, tcgen, f, "Slice")
	case *types.Map:
		return namedLookup(pk, tc, tcgen, f, "GoMap")
	case *types.Pointer:
		return namedLookup(pk, tc, tcgen, f, "Ptr")
	case *types.Basic:
		return namedLookup(pk, tc, tcgen, f, at.Name())
	}
	return namedLookup(pk, tc, tcgen, f, f.Type.String())
}

func summon(w common.Writer, pk *types.Package, tc TypeClass, tcgen *types.Package, t TypeInfo, genSet mutable.Set[string]) string {

	list := implicitTypeClassInstanceName(pk, tc, tcgen, t)

	result := list.Iterator().Filter(func(v lookupTarget) bool {
		if v.tc != nil {
			return true
		}
		ins := v.scope.Lookup(v.name)
		if ins != nil {
			return true
		}

		if v.genPk == nil && genSet.Contains(v.name) {
			return true
		}
		return false
	}).First()

	if result.IsDefined() {
		if len(result.Get().typeArgs) > 0 {
			list := seq.Map(t.TypeArgs, func(t TypeInfo) string {
				return summon(w, pk, tc, tcgen, t, genSet)
			}).MakeString(",")
			return fmt.Sprintf("%s(%s)", result.Get().instanceExpr(w), list)
		}

		return result.Get().instanceExpr(w)
	}

	if t.IsTuple() {

		aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		gpk := w.GetImportedName(tcgen)

		hlist := seq.Fold(t.TypeArgs.Reverse(), fmt.Sprintf("%s.HNil", gpk), func(tail string, ti TypeInfo) string {
			instance := summon(w, pk, tc, tcgen, ti, genSet)
			return fmt.Sprintf("%s.HCons(%s,%s)", gpk, instance, tail)
		})

		tp := seq.Map(t.TypeArgs, func(f TypeInfo) string {
			return w.TypeName(pk, f.Type)
		}).MakeString(",")

		return fmt.Sprintf("%s.ContraMap( %s,  %s.HList%d[%s])", gpk, hlist, aspk, t.TypeArgs.Size(), tp)
	}

	instance := tcgen.Scope().Lookup("Given")
	if instance != nil {
		if _, ok := instance.Type().(*types.Signature); ok {
			ctx := types.NewContext()
			_, err := types.Instantiate(ctx, instance.Type(), []types.Type{t.Type}, true)
			if err == nil {
				gpk := w.GetImportedName(tcgen)
				return fmt.Sprintf("%s.Given[%s]()", gpk, w.TypeName(pk, t.Type))
			}
		}
	}

	return list.Head().Get().instanceExpr(w)

}

func genDerive() {
	pack := os.Getenv("GOPACKAGE")

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

	common.Generate(pack, "derive_generated.go", func(w common.Writer) {

		// fmtalias := w.GetImportedName(types.NewPackage("fmt", "fmt"))
		// asalias := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		d := findTypeClassDerive(pkgs)

		genSet := iterator.ToGoSet(iterator.Map(d.Iterator(), func(v TypeClassDerive) string {
			return fmt.Sprintf("%s%s", v.TypeClass.Name, v.DeriveFor.Name)
		}))

		d.Foreach(func(v TypeClassDerive) {
			// fmt.Printf("lookup %s.Option = %v\n", v.Generator.Name(), l)
			//fmt.Printf("derive %v for %v\n", v.TypeClass, v.DeriveFor)
			privateFields := v.DeriveFor.Fields.FilterNot(StructField.Public)

			gpk := w.GetImportedName(v.Generator)

			args := seq.Map(privateFields, func(f StructField) string {
				return summon(w, pkgs[0].Types, v.TypeClass, v.Generator, f.Type, genSet)
			})

			tupleTc := v.Generator.Scope().Lookup(fmt.Sprintf("Tuple%d", privateFields.Size()))

			fmt.Printf("for %s, num type param = %d\n", v.DeriveFor.Name, v.DeriveFor.TypeParam.Size())
			if v.DeriveFor.TypeParam.Size() > 0 {

				valuetpdec := "[" + iterator.Map(v.DeriveFor.TypeParam.Iterator(), func(v TypeParam) string {
					tn := w.TypeName(pkgs[0].Types, v.Constraint)
					return fmt.Sprintf("%s %s", v.Name, tn)
				}).MakeString(",") + "]"

				valuetp := "[" + iterator.Map(v.DeriveFor.TypeParam.Iterator(), func(v TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"

				tcname := v.TypeClass.expr(w, pkgs[0].Types)
				fargs := seq.Map(v.DeriveFor.TypeParam, func(p TypeParam) string {
					return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
				}).MakeString(",")

				if tupleTc != nil {

					fmt.Fprintf(w, `
						func %s%s%s( %s ) %s[%s%s] {
							return %s.ContraMap( %s.Tuple%d(%s), %s%s.AsTuple )
						}
					`, v.TypeClass.Name, v.DeriveFor.Name, valuetpdec, fargs, tcname, v.DeriveFor.Name, valuetp,
						gpk, gpk, privateFields.Size(), args.MakeString(","), v.DeriveFor.Name, valuetp,
					)

				} else {
					fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

					aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

					hlist := seq.Fold(args.Reverse(), fmt.Sprintf("%s.HNil", gpk), func(tail string, instance string) string {
						return fmt.Sprintf("%s.HCons(%s,%s)", gpk, instance, tail)
					})

					tp := seq.Map(privateFields, func(f StructField) string {
						return w.TypeName(pkgs[0].Types, f.Type.Type)
					}).MakeString(",")

					fmt.Fprintf(w, `
						func %s%s%s( %s ) %s[%s%s] {
							return %s.ContraMap( %s, %s.Compose( %s%s.AsTuple , %s.HList%d[%s]))
						}
					`, v.TypeClass.Name, v.DeriveFor.Name, valuetpdec, fargs, tcname, v.DeriveFor.Name, valuetp,
						gpk, hlist, fppk, v.DeriveFor.Name, valuetp, aspk, privateFields.Size(), tp)
				}

			} else {
				if tupleTc != nil {
					fmt.Fprintf(w, `
				var %s%s = %s.ContraMap( %s.Tuple%d(%s), %s.AsTuple )
			`, v.TypeClass.Name, v.DeriveFor.Name, gpk, gpk, privateFields.Size(), args.MakeString(","), v.DeriveFor.Name)

				} else {
					fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

					aspk := w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

					hlist := seq.Fold(args.Reverse(), fmt.Sprintf("%s.HNil", gpk), func(tail string, instance string) string {
						return fmt.Sprintf("%s.HCons(%s,%s)", gpk, instance, tail)
					})

					tp := seq.Map(privateFields, func(f StructField) string {
						return w.TypeName(pkgs[0].Types, f.Type.Type)
					}).MakeString(",")

					fmt.Fprintf(w, `
				var %s%s = %s.ContraMap( %s, %s.Compose( %s.AsTuple , %s.HList%d[%s]))
			`, v.TypeClass.Name, v.DeriveFor.Name, gpk, hlist, fppk, v.DeriveFor.Name, aspk, privateFields.Size(), tp)
					//panic("tuple type class instance not found")

				}
			}

		})
	})
}
func main() {
	genValue()
	genDerive()

}
