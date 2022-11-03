package metafp

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type TypeClass struct {
	Name      string
	Package   *types.Package
	Type      TypeInfo
	TypeParam fp.Seq[TypeParam]
}

func (r TypeClass) InstantiatedType(t TypeInfo) TypeInfo {

	ret := r.Type

	ret.TypeArgs = seq.Of(t)
	ret.TypeParam = seq.Empty[TypeParam]()

	ctx := types.NewContext()
	ins, _ := types.Instantiate(ctx, ret.Type, []types.Type{t.Type}, false)
	ret.Type = ins

	return ret

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
										Type:      tcType,
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

type RequiredInstance struct {
	TypeClass TypeClass
	Type      TypeInfo
}
type TypeClassInstance struct {
	Package  *types.Package
	Name     string
	Static   bool
	Implicit bool
	Type     TypeInfo
	Result   TypeInfo
	Under    TypeInfo

	TypeParam        fp.Seq[TypeParam]
	RequiredInstance fp.Seq[RequiredInstance]
}

type TypeClassInstancesOfPackage struct {
	Package     *types.Package
	TypeClass   TypeClass
	ByName      fp.Map[string, TypeClassInstance]
	FixedByType fp.Map[string, TypeClassInstance]
	All         fp.Seq[TypeClassInstance]
}

func (r TypeClassInstancesOfPackage) FindByName(name string, t TypeInfo) fp.Option[TypeClassInstance] {
	return option.FlatMap(r.ByName.Get(name), as.Func2(TypeClassInstance.Check).ApplyLast(t))
}

type typeCompare struct {
	genericType TypeInfo
	actualType  TypeInfo
}

type paramVar struct {
	typeParam  *types.TypeParam
	actualType TypeInfo
}

type ConstraintCheckResult struct {
	Ok           bool
	ParamMapping fp.Map[string, TypeInfo]
}

//	func[T constraint]() Eq[T]  에서
//
// instanceType 이 T 자리에 들어갈 수 있는지 체크하는 함수
//
// func[T constraint]() Eq[Seq[T]]  같은 경우는 해당 사항 없음 .
func ConstraintCheck(param fp.Seq[TypeParam], genericType TypeInfo, typeArgs fp.Seq[TypeInfo]) ConstraintCheckResult {

	// size 가 동일하지 않은 경우
	if genericType.TypeArgs.Size() != typeArgs.Size() {
		return ConstraintCheckResult{
			Ok: false,
		}
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

	actualCheck := iterator.Map(actualArgs, func(v typeCompare) ConstraintCheckResult {
		return v.actualType.IsInstantiatedOf(param, v.genericType)
	}).ToSeq()

	actualAllMatch := actualCheck.ForAll(func(v ConstraintCheckResult) bool {
		return v.Ok
	})

	if !actualAllMatch {
		return ConstraintCheckResult{
			Ok: false,
		}
	}

	merge := seq.Map(actualCheck, func(v ConstraintCheckResult) fp.Map[string, TypeInfo] {
		return v.ParamMapping
	}).Reduce(monoid.MergeMap[string, TypeInfo]())

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
		return ConstraintCheckResult{
			Ok: true,
		}
	}

	if !paramFound.ForAll(fp.Option[paramVar].IsDefined) {
		return ConstraintCheckResult{
			Ok: false,
		}
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
		mapping := seq.ToGoMap(seq.Map(seq.Map(paramFound, fp.Option[paramVar].Get), func(v paramVar) fp.Tuple2[string, TypeInfo] {
			return as.Tuple2(v.typeParam.Obj().Name(), v.actualType)
		}))

		return ConstraintCheckResult{
			Ok:           true,
			ParamMapping: merge.Concat(mutable.MapOf(mapping)),
		}
	}

	return ConstraintCheckResult{
		Ok: false,
	}
}

func (r TypeClassInstance) Check(t TypeInfo) fp.Option[TypeClassInstance] {

	argType := r.Result.TypeArgs.Head().Get()

	if argType.IsTypeParam() {

		// func[T any]() Eq[T] 인 경우
		// t 가 T constraint 인지 체크해야 함
		check := ConstraintCheck(r.Type.TypeParam, r.Result, seq.Of(t))
		if check.Ok {

			r.RequiredInstance = seq.Map(r.RequiredInstance, func(v RequiredInstance) RequiredInstance {
				v.Type = v.Type.ReplaceTypeParam(check.ParamMapping)
				return v
			})
			return option.Some(r)
		}
		return option.None[TypeClassInstance]()
	}

	// func[T any]() Eq[Tuple[T]] 인경우
	// Tuple[T] 와  t 를 비교해야 함
	if argType.TypeParam.Size() > 0 {
		check := t.IsInstantiatedOf(r.Type.TypeParam, argType)
		if check.Ok {
			r.RequiredInstance = seq.Map(r.RequiredInstance, func(v RequiredInstance) RequiredInstance {
				v.Type = v.Type.ReplaceTypeParam(check.ParamMapping)
				return v
			})
			return option.Some(r)
		}
		return option.None[TypeClassInstance]()

	}

	if argType.Type.String() == t.Type.String() {
		return option.Some(r)
	}
	return option.None[TypeClassInstance]()

}

// t 는 Eq 쌓이지 않은 타입
// Eq[T] 여서는 안됨
func (r TypeClassInstancesOfPackage) Find(t TypeInfo) fp.Seq[TypeClassInstance] {
	ret := option.FlatMap(
		r.FixedByType.Get(t.Type.String()),
		as.Func2(TypeClassInstance.Check).ApplyLast(t),
	)

	if ret.IsDefined() {
		return ret.ToSeq()
	}

	return seq.FlatMap(r.All, func(v TypeClassInstance) fp.Seq[TypeClassInstance] {
		return v.Check(t).ToSeq()
	})
}

type TypeClassInstanceCache struct {
	tcMap fp.Map[string, fp.Seq[TypeClassInstancesOfPackage]]
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
	list := r.tcMap.Get(tc.TypeClass.Id()).OrZero()

	pkPred := func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == tc.Package.Path()
	}

	found := list.Find(pkPred).OrElseGet(func() TypeClassInstancesOfPackage {
		return LoadTypeClassInstance(tc.Package, tc.TypeClass)
	})

	t := tc.TypeClass.InstantiatedType(tc.DeriveFor.Info)
	ins := TypeClassInstance{
		Package:  tc.Package,
		Name:     tc.GeneratedInstanceName(),
		Static:   true,
		Implicit: false,
		Type:     t,
		Result:   t,
	}

	found.ByName = found.ByName.Updated(ins.Name, ins)

	r.tcMap = r.tcMap.Updated(tc.TypeClass.Id(),
		list.FilterNot(pkPred).Append(found),
	)

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

		if rType.IsInstanceOf(tc) && rType.TypeArgs.Size() > 0 {
			under := rType.TypeArgs.Head().Get()

			if insType.IsFunc() {

				if insType.NumArgs() == 0 {
					tins := TypeClassInstance{
						Package:  pk,
						Name:     name,
						Static:   false,
						Implicit: true,
						Type:     insType,
						Result:   rType,
						Under:    under,
					}
					ret.ByName = ret.ByName.Updated(name, tins)
					ret.All = ret.All.Append(tins)
				} else {

					fargs := insType.FuncArgs()
					allArgTypeClass := fargs.ForAll(func(v TypeInfo) bool {
						return v.Name().IsDefined() && v.TypeArgs.Size() == 1
					})

					if allArgTypeClass == true {

						required := seq.Map(fargs, func(v TypeInfo) RequiredInstance {
							return RequiredInstance{
								TypeClass: TypeClass{
									Name:      v.Name().Get(),
									Package:   v.Pkg,
									Type:      v,
									TypeParam: v.TypeParam,
								},
								Type: v.TypeArgs.Head().Get(),
							}
						})

						tins := TypeClassInstance{
							Package:          pk,
							Name:             name,
							Static:           false,
							Implicit:         false,
							Type:             insType,
							Result:           rType,
							TypeParam:        insType.TypeParam,
							RequiredInstance: required,
							Under:            under,
						}
						ret.ByName = ret.ByName.Updated(name, tins)
						ret.All = ret.All.Append(tins)
					}
				}
			} else {
				tins := TypeClassInstance{
					Package:  pk,
					Name:     name,
					Static:   true,
					Implicit: false,
					Type:     insType,
					Result:   insType,
					Under:    under,
				}
				ret.ByName = ret.ByName.Updated(name, tins)
				//fmt.Printf("add fixed type %s -> %s\n", name, under.Type.String())
				ret.FixedByType = ret.FixedByType.Updated(under.Type.String(), tins)
				//ret.All = ret.All.Append(tins)
			}
		}

	}
	return ret
}
