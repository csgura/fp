package metafp

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type TypeClass struct {
	Name      string
	Package   *types.Package
	TypeParam fp.Seq[TypeParam]
}

func (r TypeClass) PackagedName(w ImportSet, workingPackage *types.Package) string {
	if r.Package != nil && r.Package.Path() != workingPackage.Path() {
		pk := w.GetImportedName(r.Package)
		return fmt.Sprintf("%s.%s", pk, r.Name)
	}
	return r.Name
}

type TypeClassDerive struct {
	Package              *types.Package
	PrimitiveInstancePkg *types.Package
	TypeClass            TypeClass
	DeriveFor            TaggedStruct
}

func publicName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func (r TypeClassDerive) GeneratedInstanceName() string {
	if r.DeriveFor.Package != nil && r.Package.Path() != r.DeriveFor.Package.Path() {
		return fmt.Sprintf("%s%s%s", r.TypeClass.Name, publicName(r.DeriveFor.Package.Name()), r.DeriveFor.Name)
	}
	return fmt.Sprintf("%s%s", r.TypeClass.Name, r.DeriveFor.Name)
}

type TypeClassDirective struct {
	Package              *types.Package
	PrimitiveInstancePkg *types.Package
	TypeClass            TypeClass
	TypeArgs             fp.Seq[TypeInfo]
}

func findTypeClsssDirective(p []*packages.Package, directive string) fp.Seq[TypeClassDirective] {
	return seq.FlatMap(p, func(pk *packages.Package) fp.Seq[TypeClassDirective] {
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

		return seq.FlatMap(s3, func(gd *ast.GenDecl) fp.Seq[TypeClassDirective] {
			gdDoc := option.Of(gd.Doc)

			return seq.FlatMap(gd.Specs, func(v ast.Spec) fp.Seq[TypeClassDirective] {
				if vs, ok := v.(*ast.ValueSpec); ok {

					doc := option.Map(option.Of(vs.Doc).Or(fp.Return(gdDoc)), (*ast.CommentGroup).Text)

					if doc.Filter(as.Func2(strings.Contains).ApplyLast(directive)).IsDefined() {

						info := &types.Info{
							Types: make(map[ast.Expr]types.TypeAndValue),
						}
						types.CheckExpr(pk.Fset, pk.Types, v.Pos(), vs.Type, info)
						ti := info.Types[vs.Type]

						if nt, ok := ti.Type.(*types.Named); ok && nt.TypeArgs().Len() == 1 {
							if tt, ok := nt.TypeArgs().At(0).(*types.Named); ok && tt.TypeArgs().Len() > 0 {
								tcType := typeInfo(tt.Obj().Type())

								return seq.Of(TypeClassDirective{
									Package:              pk.Types,
									PrimitiveInstancePkg: nt.Obj().Pkg(),
									TypeClass: TypeClass{
										Name:      tt.Obj().Name(),
										Package:   tt.Obj().Pkg(),
										TypeParam: tcType.TypeParam,
									},
									TypeArgs: typeArgs(tt.TypeArgs()),
								})
							}
						}
					}
				}
				return seq.Of[TypeClassDirective]()
			})
		})
	})
}

func FindTypeClassDerive(p []*packages.Package) fp.Seq[TypeClassDerive] {
	return seq.FlatMap(findTypeClsssDirective(p, "@fp.Derive"), func(v TypeClassDirective) fp.Seq[TypeClassDerive] {
		if v.TypeArgs.Size() == 1 && v.TypeArgs.Head().Get().Name().IsDefined() {
			deriveFor := v.TypeArgs.Head().Get()

			vt := LookupStruct(deriveFor.Pkg, deriveFor.Name().Get())
			if vt.IsDefined() {
				return seq.Of(TypeClassDerive{
					Package:              v.Package,
					PrimitiveInstancePkg: v.PrimitiveInstancePkg,
					TypeClass:            v.TypeClass,
					DeriveFor:            vt.Get(),
				})
			} else {
				fmt.Printf("can't lookup %s\n", deriveFor.Name().Get())
			}
		}
		return seq.Empty[TypeClassDerive]()
	})
}

func FindTypeClassImport(p []*packages.Package) fp.Seq[TypeClassDirective] {
	return findTypeClsssDirective(p, "@fp.ImportGiven")
}

type TypeClassInstance struct {
	Package  *types.Package
	Name     string
	Static   bool
	Implicit bool
	Type     TypeInfo
	Result   TypeInfo
}

type TypeClassInstanceIndex struct {
	Package     *types.Package
	TypeClass   TypeClass
	FixedByType fp.Map[string, TypeClassInstance]
	All         fp.Seq[TypeClassInstance]
}

func (r TypeClassInstanceIndex) Summon(t TypeInfo) fp.Seq[TypeClassInstance] {
	// ret := r.FixedByType.Get(t.Type.String())
	// if ret.IsDefined() {
	// 	return ret
	// }

	fmt.Printf("t type = %s\n", t.Type.String())
	return r.All.Filter(func(v TypeClassInstance) bool {
		argType := v.Result.TypeArgs.Head().Get()
		if argType.IsTypeParam() {
			ctx := types.NewContext()

			_, err := types.Instantiate(ctx, v.Result.Type, []types.Type{t.Type}, true)
			if err == nil {
				return true
			}
			return false
		}

		if argType.TypeParam.Size() > 0 {
			return t.IsInstantiatedOf(argType)
		}

		fmt.Printf("%s result Type = %s\n", v.Name, argType.String())
		return argType.String() == t.Type.String()
	})
}

func ImportTypeClassInstance(pk *types.Package, tc TypeClass) TypeClassInstanceIndex {
	ret := TypeClassInstanceIndex{
		Package:     pk,
		TypeClass:   tc,
		FixedByType: mutable.EmptyMap[string, TypeClassInstance](),

		All: seq.Empty[TypeClassInstance](),
	}
	for _, name := range pk.Scope().Names() {
		ins := pk.Scope().Lookup(name)
		insType := typeInfo(ins.Type())
		rType := insType.ResultType()

		if rType.IsInstanceOf(tc) {
			if insType.IsFunc() {

				if insType.NumArgs() == 0 {
					tins := TypeClassInstance{
						Package:  pk,
						Name:     name,
						Static:   false,
						Implicit: true,
						Type:     insType,
						Result:   rType,
					}
					ret.All = ret.All.Append(tins)
				} else {
					tins := TypeClassInstance{
						Package:  pk,
						Name:     name,
						Static:   false,
						Implicit: false,
						Type:     insType,
						Result:   rType,
					}
					ret.All = ret.All.Append(tins)
				}
			} else {
				tins := TypeClassInstance{
					Package:  pk,
					Name:     name,
					Static:   true,
					Implicit: false,
					Type:     insType,
					Result:   insType,
				}
				ret.FixedByType = ret.FixedByType.Updated(rType.Type.String(), tins)
				ret.All = ret.All.Append(tins)

			}
		}

	}
	return ret
}
