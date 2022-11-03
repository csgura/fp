package metafp

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
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

func (r TypeClass) Id() string {
	if r.Package != nil {
		return fmt.Sprintf("%s.%s", r.Package.Path(), r.Name)
	}
	return r.Name
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

type TypeClassInstancesOfPackage struct {
	Package     *types.Package
	TypeClass   TypeClass
	ByName      fp.Map[string, TypeClassInstance]
	FixedByType fp.Map[string, TypeClassInstance]
	All         fp.Seq[TypeClassInstance]
}

type typeCompare struct {
	genericType TypeInfo
	actualType  TypeInfo
}

type paramVar struct {
	typeParam  *types.TypeParam
	actualType TypeInfo
}

//	func[T constraint]() Eq[T]  에서
//
// instanceType 이 T 자리에 들어갈 수 있는지 체크하는 함수
//
// func[T constraint]() Eq[Seq[T]]  같은 경우는 해당 사항 없음 .
func ConstraintCheck(param fp.Seq[TypeParam], genericType TypeInfo, typeArgs fp.Seq[TypeInfo]) bool {

	// size 가 동일하지 않은 경우
	if genericType.TypeArgs.Size() != typeArgs.Size() {
		return false
	}

	zipped := iterator.Map(iterator.Zip(genericType.TypeArgs.Iterator(), typeArgs.Iterator()), func(t fp.Tuple2[TypeInfo, TypeInfo]) typeCompare {
		return typeCompare{
			genericType: t.I1,
			actualType:  t.I2,
		}
	})

	// Eq[T] 가 아니고,  Eq[Seq[T]]  같은 경우는  체크 불가능
	paramArgs, actualArgs := zipped.Partition(func(t typeCompare) bool {
		return t.genericType.IsTypeParam()
	})

	actualAllMatch := actualArgs.ForAll(func(v typeCompare) bool {
		return v.actualType.IsInstantiatedOf(param, v.genericType)
	})

	if !actualAllMatch {
		return false
	}

	paramFound := iterator.Map(paramArgs, func(v typeCompare) fp.Option[paramVar] {
		paramName := v.genericType.Name().Get()

		// func[T constraint]() Eq[A] 혹은 func() Eq[A] 처럼. type parameter 목록이 잘못된 경우도 불가능
		paramCons := option.Map(param.Filter(func(p TypeParam) bool {
			return p.Name == paramName
		}).Head(), func(p TypeParam) paramVar {
			//fmt.Printf("param %s -> %s\n", p.TypeName, v.actualType)
			return paramVar{
				typeParam:  types.NewTypeParam(p.TypeName, p.Constraint),
				actualType: v.actualType,
			}
		})

		return paramCons
	}).ToSeq()

	if paramFound.IsEmpty() {
		return true
	}

	if !paramFound.ForAll(fp.Option[paramVar].IsDefined) {
		return false
	}

	paramCons := seq.Map(paramFound, func(v fp.Option[paramVar]) *types.TypeParam {
		return v.Get().typeParam
	})

	paramIns := seq.Map(paramFound, func(v fp.Option[paramVar]) types.Type {
		return v.Get().actualType.Type
	})

	sig := types.NewSignatureType(nil,
		nil,
		paramCons,
		types.NewTuple(),
		types.NewTuple(types.NewVar(0, genericType.Pkg, "ret", genericType.Type)),
		false,
	)

	ctx := types.NewContext()

	_, err := types.Instantiate(ctx, sig, paramIns, true)
	if err == nil {
		return true
	}

	return false
}

func (r TypeClassInstancesOfPackage) Find(t TypeInfo) fp.Seq[TypeClassInstance] {
	// ret := r.FixedByType.Get(t.Type.String())
	// if ret.IsDefined() {
	// 	return ret
	// }

	fmt.Printf("t type = %s\n", t.Type.String())
	return r.All.Filter(func(v TypeClassInstance) bool {
		argType := v.Result.TypeArgs.Head().Get()
		if argType.IsTypeParam() {

			// func[T any]() Eq[T] 인 경우
			// t 가 T constraint 인지 체크해야 함
			return ConstraintCheck(v.Type.TypeParam, v.Result, seq.Of(t))
		}

		// func[T any]() Eq[Tuple[T]] 인경우
		// Tuple[T] 와  t 를 비교해야 함
		if argType.TypeParam.Size() > 0 {
			return t.IsInstantiatedOf(v.Type.TypeParam, argType)
		}

		return argType.String() == t.Type.String()
	})
}

type TypeClassInstanceCache struct {
	tcMap   fp.Map[string, fp.Seq[TypeClassInstancesOfPackage]]
	willGen fp.Map[string, fp.Seq[TypeClassInstancesOfPackage]]
}

func (r *TypeClassInstanceCache) Load(pk *types.Package, tc TypeClass) TypeClassInstancesOfPackage {

	list := r.tcMap.Get(tc.Id()).OrElseGet(seq.Empty[TypeClassInstancesOfPackage])

	found := list.Find(func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == pk.Path()
	})

	if found.IsEmpty() {
		loaded := LoadTypeClassInstance(pk, tc)

		list = list.Append(loaded)
		r.tcMap = r.tcMap.Updated(tc.Id(), list)
		return loaded
	}

	return found.Get()
}

func (r *TypeClassInstanceCache) WillGenerated(tc TypeClassDerive) TypeClassInstancesOfPackage {
	list := r.willGen.Get(tc.TypeClass.Id()).OrZero()

	found := list.Find(func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == tc.Package.Path()
	}).OrElseGet(func() TypeClassInstancesOfPackage {

		return TypeClassInstancesOfPackage{
			Package:   tc.Package,
			TypeClass: tc.TypeClass,
		}
	})

	ins := TypeClassInstance{
		Package:  tc.Package,
		Name:     tc.GeneratedInstanceName(),
		Static:   true,
		Implicit: false,
	}

	found.ByName = found.ByName.Updated(ins.Name, ins)

	return found
}

type TypeClassScope struct {
	Cache          *TypeClassInstanceCache
	Target         TypeClassDerive
	WorkingScope   TypeClassInstancesOfPackage
	PrimitiveScope TypeClassInstancesOfPackage
	TargetScope    fp.Option[TypeClassInstancesOfPackage]
	Others         fp.Seq[TypeClassInstancesOfPackage]
}

func (r *TypeClassInstanceCache) Get(tc TypeClassDerive) TypeClassScope {

	working := r.Load(tc.Package, tc.TypeClass)
	prim := r.Load(tc.PrimitiveInstancePkg, tc.TypeClass)
	target := option.Map(option.Of(tc.DeriveFor.Package), as.Func2(r.Load).ApplyLast(tc.TypeClass))

	others := r.tcMap.Get(tc.TypeClass.Id()).OrZero().FilterNot(func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == working.Package.Path()
	}).FilterNot(func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == prim.Package.Path()
	}).FilterNot(func(v TypeClassInstancesOfPackage) bool {
		return target.IsDefined() && target.Get().Package.Path() == v.Package.Path()
	})

	return TypeClassScope{
		Cache:          r,
		Target:         tc,
		WorkingScope:   working,
		PrimitiveScope: prim,
		TargetScope:    target,
		Others:         others,
	}

}

func LoadTypeClassInstance(pk *types.Package, tc TypeClass) TypeClassInstancesOfPackage {
	ret := TypeClassInstancesOfPackage{
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
					ret.ByName = ret.ByName.Updated(name, tins)
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
					ret.ByName = ret.ByName.Updated(name, tins)
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
				ret.ByName = ret.ByName.Updated(name, tins)
				ret.FixedByType = ret.FixedByType.Updated(rType.Type.String(), tins)
				//ret.All = ret.All.Append(tins)
			}
		}

	}
	return ret
}
