package metafp

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
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

func (r TypeClass) IsLazy() bool {
	return r.Name == "Eval" && r.Package.Path() == "github.com/csgura/fp/lazy"
}

// int  =>  Eq[int]
// Option[T any] =>   (T) -> Eq[Option[T]]
func (r TypeClass) InstantiatedType(t TypeInfo) TypeInfo {

	ret := r.Type

	ret.TypeParam = t.TypeParam

	t.TypeParam = seq.Empty[TypeParam]()
	t.TypeArgs = seq.Map(ret.TypeParam, func(v TypeParam) TypeInfo {
		p := types.NewTypeParam(v.TypeName, v.Constraint)
		return typeInfo(p)
	})
	ret.TypeArgs = seq.Of(t)

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

func (r TypeClass) PackagedName(w genfp.ImportSet, workingPackage *types.Package) string {
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
	DeriveFor            NamedTypeInfo
	StructInfo           fp.Option[TaggedStruct]
	Tags                 fp.Map[string, Annotation]
}

func (r TypeClassDerive) IsRecursive() bool {
	return r.StructInfo.IsDefined() && r.StructInfo.Get().IsRecursive()
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
	Tags                 fp.Map[string, Annotation]
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

					doc := option.Map(option.Of(vs.Doc).Or(as.Supplier(gdDoc)), (*ast.CommentGroup).Text)

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
									Tags:     option.Map(doc, extractTag).OrZero(),
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
		if v.TypeArgs.Size() == 1 && v.TypeArgs.Head().Get().IsNamed() {
			deriveFor := v.TypeArgs.Head().Get()

			obj := deriveFor.Pkg.Scope().Lookup(deriveFor.Name().Get())
			vt := LookupStruct(deriveFor.Pkg, deriveFor.Name().Get())

			return seq.Of(TypeClassDerive{
				Package:              v.Package,
				PrimitiveInstancePkg: v.PrimitiveInstancePkg,
				TypeClass:            v.TypeClass,
				DeriveFor:            typeInfo(obj.Type()).AsNamed().Get(),
				StructInfo:           vt,
				Tags:                 v.Tags,
			})

		} else if v.TypeArgs.Size() == 1 {
			fmt.Printf("can't derive not named type %s\n", v.TypeArgs[0])
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
	Lazy      bool
	FieldOf   fp.Option[TypeInfo]
}

func (r RequiredInstance) String() string {
	return fmt.Sprintf("%s[%s]", r.TypeClass.Id(), r.Type.String())
}

type TypeClassInstance struct {
	Package *types.Package
	Name    string

	// var 인지 func 인지 여부
	Static bool

	// Given[comparable]() 형태인지 여부
	Implicit bool

	// ContraMap 처럼,  아규먼트가 타입클래스 인스턴스로만 이루어 진것이 아니라, 다른 아규먼트를 가지고 있음.
	HasExplictArg bool

	// lookup한 instance
	Instance types.Object

	// instance 의 타입
	Type TypeInfo

	// func 인 경우에 return 타입,  var 인 경우는 Type과 동일
	Result TypeInfo

	// type A int  와 같은 경우에 underlying 타입
	Under TypeInfo

	// func[A, B any]()  형태인 경우에  타입 파라미터 목록
	TypeParam fp.Seq[TypeParam]

	// func( a Eq[int], b Ord[int] ) 형태인 경우에 , Eq[int], Ord[int] 정보
	RequiredInstance fp.Seq[RequiredInstance]

	// func[A,B any]( b B ) 형태인 경우 , B 만 들어 있음
	UsedParam fp.Set[string]

	// Eq[int] 를 찾는데  func Given[T any]() Eq[T]  를 찾았으면  T -> int  정보가 저장됨
	ParamMapping fp.Map[string, TypeInfo]

	// 생성될 instance 여서  RequiredInstance 가 정확하지 않을 수 있음.
	WillGeneratedBy fp.Option[TypeClassDerive]
}

func (r TypeClassInstance) String() string {
	if r.Package != nil {
		return fmt.Sprintf("%s.%s : (%s) -> %s", r.Package.Path(), r.Name, r.RequiredInstance.MakeString(","), r.Result)
	}
	return r.Name
}

func (r TypeClassInstance) PackagedName(importSet genfp.ImportSet, working *types.Package) string {
	if r.Package.Path() == working.Path() {
		return r.Name
	}
	pk := importSet.GetImportedName(r.Package)
	return fmt.Sprintf("%s.%s", pk, r.Name)
}

func (r TypeClassInstance) IsFunc() bool {
	return r.Type.IsFunc()
}

func (r TypeClassInstance) IsGivenAny() bool {
	return r.Implicit && r.TypeParam.Size() == 1 && r.TypeParam[0].IsAny()
}

type TypeClassInstancesOfPackage struct {
	Package     *types.Package
	TypeClass   TypeClass
	ByName      fp.Map[string, TypeClassInstance]
	FixedByType fp.Map[string, TypeClassInstance]
	OtherFuncs  fp.Map[string, TypeClassInstance]
	All         fp.Seq[TypeClassInstance]
}

func (r TypeClassInstancesOfPackage) FindFunc(name string) fp.Option[TypeClassInstance] {

	ret := r.ByName.Get(name)
	if ret.IsDefined() {
		return ret
	}
	return r.OtherFuncs.Get(name)
}

func (r TypeClassInstancesOfPackage) FindByName(name string, t TypeInfo) fp.Option[TypeClassInstance] {
	//	fmt.Printf("find %s\n", name)
	ret := option.FlatMap(r.ByName.Get(name), as.Func2(TypeClassInstance.Check).ApplyLast(t))
	return ret
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

	zipped := iterator.Map(iterator.Zip(seq.Iterator(genericType.TypeArgs), seq.Iterator(typeArgs)), func(t fp.Tuple2[TypeInfo, TypeInfo]) typeCompare {
		return typeCompare{
			genericType: t.I1,
			actualType:  t.I2,
		}
	})

	// Eq[T] 가 아니고,  Eq[Seq[T]]  같은 경우는  체크 불가능
	paramArgs, actualArgs := iterator.Partition(zipped, func(t typeCompare) bool {
		return t.genericType.IsTypeParam()
	})

	actualCheck := iterator.Map(actualArgs, func(v typeCompare) ConstraintCheckResult {
		return v.actualType.IsInstantiatedOf(param, v.genericType)
	}).ToSeq()

	actualAllMatch := as.Seq(actualCheck).ForAll(func(v ConstraintCheckResult) bool {
		return v.Ok
	})

	if !actualAllMatch {
		return ConstraintCheckResult{
			Ok: false,
		}
	}

	merge := seq.Reduce(seq.Map(actualCheck, func(v ConstraintCheckResult) fp.Map[string, TypeInfo] {
		return v.ParamMapping
	}), monoid.MergeMap[string, TypeInfo]())

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

	if len(paramFound) == 0 {
		return ConstraintCheckResult{
			Ok: true,
		}
	}

	if !as.Seq(paramFound).ForAll(fp.Option[paramVar].IsDefined) {
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
	//fmt.Printf("check %s.%s : %t, %d\n", r.Package.Name(), r.Name, argType.IsTypeParam(), argType.TypeArgs.Size())

	if argType.IsTypeParam() {

		// func[T any]() Eq[T] 인 경우
		// t 가 T constraint 인지 체크해야 함
		check := ConstraintCheck(r.Type.TypeParam, r.Result, seq.Of(t))
		if check.Ok {

			r.RequiredInstance = seq.Map(r.RequiredInstance, func(v RequiredInstance) RequiredInstance {
				res := v.Type.ReplaceTypeParam(check.ParamMapping)
				r.UsedParam = r.UsedParam.Concat(res.I1)
				v.Type = res.I2
				return v
			})
			r.ParamMapping = check.ParamMapping
			return option.Some(r)
		}
		return option.None[TypeClassInstance]()
	}

	// func[T any]() Eq[Tuple[T]] 인경우
	// Tuple[T] 와  t 를 비교해야 함
	if argType.TypeArgs.Size() > 0 {

		check := t.IsInstantiatedOf(r.Type.TypeParam, argType)
		if check.Ok {
			r.RequiredInstance = seq.Map(r.RequiredInstance, func(v RequiredInstance) RequiredInstance {
				res := v.Type.ReplaceTypeParam(check.ParamMapping)
				r.UsedParam = r.UsedParam.Concat(res.I1)
				v.Type = res.I2
				return v
			})
			r.ParamMapping = check.ParamMapping
			//fmt.Printf("check %s.%s : %v\n", r.Package.Name(), r.Name, check.ParamMapping)

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

func (r *TypeClassInstanceCache) IsWillGenerated(tc TypeClassDerive) bool {
	list := r.tcMap.Get(tc.TypeClass.Id()).OrZero()

	pkPred := func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == tc.Package.Path()
	}

	found := list.Find(pkPred)
	if found.IsDefined() {
		return found.Get().ByName.Contains(tc.GeneratedInstanceName())
	}
	return false
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
	//fmt.Printf("will generate %s, %v, %v\n", tc.GeneratedInstanceName(), tc.DeriveFor.Info.TypeParam, tc.DeriveFor.Info.TypeArgs)

	// if t.TypeArgs.Head().IsDefined() {
	// 	fmt.Printf("will generate arg type = %s\n", t.TypeArgs.Head().Get().TypeArgs)
	// }
	ins := TypeClassInstance{
		Package:         tc.Package,
		Name:            tc.GeneratedInstanceName(),
		Static:          false,
		Implicit:        false,
		Type:            t,
		Result:          t,
		WillGeneratedBy: option.Some(tc),
	}

	// if tc.IsRecursive() {
	// 	ins.Static = false
	// }

	if tc.DeriveFor.Info.TypeParam.Size() > 0 {
		ins.Static = false
		ins.TypeParam = tc.DeriveFor.Info.TypeParam
		ins.RequiredInstance = seq.Map(tc.DeriveFor.Info.TypeParam, func(v TypeParam) RequiredInstance {
			p := types.NewTypeParam(v.TypeName, v.Constraint)
			return RequiredInstance{
				TypeClass: tc.TypeClass,
				Type:      typeInfo(p),
			}
		})
	}

	found.ByName = found.ByName.Updated(ins.Name, ins)

	newList := list.FilterNot(pkPred).Append(found)
	r.tcMap = r.tcMap.Updated(tc.TypeClass.Id(), newList)

	return found
}

type TypeClassScope struct {
	Cache     *TypeClassInstanceCache
	Typeclass TypeClass
	List      fp.Seq[TypeClassInstancesOfPackage]
}

func (r TypeClassScope) FindByName(name string, t TypeInfo) fp.Option[TypeClassInstance] {

	ret := iterator.Map(seq.Iterator(r.List), func(p TypeClassInstancesOfPackage) fp.Option[TypeClassInstance] {
		return p.FindByName(name, t)
	}).Filter(fp.Option[TypeClassInstance].IsDefined).NextOption()

	return option.Flatten(ret)
}

func (r TypeClassScope) FindFunc(name string) fp.Option[TypeClassInstance] {

	ret := iterator.Map(seq.Iterator(r.List), func(p TypeClassInstancesOfPackage) fp.Option[TypeClassInstance] {
		return p.FindFunc(name)
	}).Filter(fp.Option[TypeClassInstance].IsDefined).NextOption()

	return option.Flatten(ret)
}

func (r TypeClassScope) Find(t TypeInfo) fp.Seq[TypeClassInstance] {
	ret := seq.FlatMap(r.List, func(p TypeClassInstancesOfPackage) fp.Seq[TypeClassInstance] {
		return p.Find(t)
	})
	return ret
}

func (r *TypeClassInstanceCache) GetImported(tc TypeClass) TypeClassScope {
	return TypeClassScope{
		Cache:     r,
		Typeclass: tc,
		List:      r.tcMap.Get(tc.Id()).OrZero(),
	}
}

func (r *TypeClassInstanceCache) GetLocal(pk *types.Package, tc TypeClass) TypeClassScope {

	working := r.Load(pk, tc)

	return TypeClassScope{
		Cache:     r,
		Typeclass: tc,
		List:      seq.Of(working),
	}

}

func (r *TypeClassInstanceCache) Get(pk *types.Package, tc TypeClass) TypeClassScope {

	working := r.Load(pk, tc)

	others := r.tcMap.Get(tc.Id()).OrZero().FilterNot(func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == working.Package.Path()
	})

	return TypeClassScope{
		Cache:     r,
		Typeclass: tc,
		List:      seq.Concat(working, others),
	}

}

func asRequired(v TypeInfo) RequiredInstance {
	tc := TypeClass{
		Name:      v.Name().Get(),
		Package:   v.Pkg,
		Type:      v,
		TypeParam: v.TypeParam,
	}

	if tc.IsLazy() {
		ret := asRequired(v.TypeArgs.Head().Get())
		ret.Lazy = true
		return ret
	}

	return RequiredInstance{
		TypeClass: tc,
		Type:      v.TypeArgs.Head().Get(),
	}
}

func AsTypeClassInstance(tc TypeClass, ins types.Object) fp.Option[TypeClassInstance] {
	insType := typeInfo(ins.Type())
	rType := insType.ResultType()
	name := ins.Name()

	if rType.IsInstanceOf(tc) && rType.TypeArgs.Size() > 0 {
		under := rType.TypeArgs.Head().Get()

		if insType.IsFunc() {

			if insType.NumArgs() == 0 && insType.TypeParam.Size() == 1 {
				return option.Some(TypeClassInstance{
					Package:   ins.Pkg(),
					Name:      name,
					Static:    false,
					Implicit:  true,
					Type:      insType,
					Result:    rType,
					Under:     under,
					Instance:  ins,
					TypeParam: insType.TypeParam,
				})
				// ret.ByName = ret.ByName.Updated(name, tins)
				// ret.All = ret.All.Append(tins)
			} else {

				fargs := insType.FuncArgs()
				allArgTypeClass := fargs.ForAll(func(v TypeInfo) bool {
					if v.Name().IsDefined() && v.TypeArgs.Size() == 1 {
						if v.Name().Get() == "Evel" && v.Pkg != nil && v.Pkg.Path() == "github.com/csgura/fp/lazy" {
							realv := v.TypeArgs.Head().Get()
							return realv.Name().IsDefined() && realv.TypeArgs.Size() == 1
						}
						return true
					}
					return false
				})

				if allArgTypeClass == true {

					required := seq.Map(fargs, asRequired)

					return option.Some(TypeClassInstance{
						Package:          ins.Pkg(),
						Name:             name,
						Static:           false,
						Implicit:         false,
						Type:             insType,
						Result:           rType,
						TypeParam:        insType.TypeParam,
						RequiredInstance: required,
						Under:            under,
						Instance:         ins,
					})
					// ret.ByName = ret.ByName.Updated(name, tins)
					// ret.All = ret.All.Append(tins)
				} else {
					return option.Some(TypeClassInstance{
						Package:       ins.Pkg(),
						Name:          name,
						Static:        false,
						Implicit:      false,
						HasExplictArg: true,
						Type:          insType,
						Result:        rType,
						TypeParam:     insType.TypeParam,
						Under:         under,
						Instance:      ins,
					})
					//ret.OtherFuncs = ret.OtherFuncs.Updated(name, tins)
				}
			}
		} else {
			return option.Some(TypeClassInstance{
				Package:  ins.Pkg(),
				Name:     name,
				Static:   true,
				Implicit: false,
				Type:     insType,
				Result:   insType,
				Under:    under,
				Instance: ins,
			})
			// ret.ByName = ret.ByName.Updated(name, tins)
			// ret.FixedByType = ret.FixedByType.Updated(under.Type.String(), tins)
		}
	}
	return option.None[TypeClassInstance]()

}

func LoadTypeClassInstance(pk *types.Package, tc TypeClass) TypeClassInstancesOfPackage {
	ret := TypeClassInstancesOfPackage{
		Package:     pk,
		TypeClass:   tc,
		FixedByType: mutable.EmptyMap[string, TypeClassInstance](),
		OtherFuncs:  mutable.EmptyMap[string, TypeClassInstance](),
		ByName:      mutable.EmptyMap[string, TypeClassInstance](),
		All:         seq.Empty[TypeClassInstance](),
	}
	// fmt.Printf("Searching instances of %s from %s\n", tc.Name, pk.Path())
	for _, name := range pk.Scope().Names() {
		ins := pk.Scope().Lookup(name)
		AsTypeClassInstance(tc, ins).Foreach(func(tins TypeClassInstance) {
			if tins.HasExplictArg {
				ret.OtherFuncs = ret.OtherFuncs.Updated(name, tins)
			} else {
				ret.All = ret.All.Append(tins)
				ret.ByName = ret.ByName.Updated(name, tins)
			}
			if tins.Static {
				ret.FixedByType = ret.FixedByType.Updated(tins.Under.Type.String(), tins)
			}
		})

	}
	ret.All = seq.Sort(ret.All, as.Ord(func(a, b TypeClassInstance) bool {
		if !a.Implicit && b.Implicit {
			return true
		}

		if a.Implicit && !b.Implicit {
			return false
		}

		if a.Implicit && b.Implicit {
			consA := a.Type.TypeParam.Head().Get().Constraint.Underlying()
			consB := b.Type.TypeParam.Head().Get().Constraint.Underlying()

			return types.Implements(consA, consB.(*types.Interface))

		}
		return a.RequiredInstance.Size() < b.RequiredInstance.Size()

	}))
	// ord := seq.Map(ret.All, func(v TypeClassInstance) string {
	// 	return v.Name
	// }).MakeString(",")
	// fmt.Printf("%s sorted =%s\n", tc.Name, ord)
	return ret
}
