package metafp

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type TypeClass struct {
	Name    string
	Package genfp.PackageId
}

func (r TypeClass) IsLazy() bool {
	return r.Name == "Eval" && r.Package.Path() == "github.com/csgura/fp/lazy"
}

func (r TypeClass) Id() string {
	if r.Package != nil {
		return fmt.Sprintf("%s.%s", r.Package.Path(), r.Name)
	}
	return r.Name
}

func (r TypeClass) PackagedName(w genfp.ImportSet, workingPackage genfp.WorkingPackage) string {
	if r.Package != nil && r.Package.Path() != workingPackage.Path() {
		pk := w.GetImportedName(r.Package)
		return fmt.Sprintf("%s.%s", pk, r.Name)
	}
	return r.Name
}

type TypeClassDerive struct {
	Package              genfp.WorkingPackage
	PrimitiveInstancePkg *types.Package
	TypeClass            TypeClass
	TypeClassType        TypeInfo
	DeriveFor            NamedTypeInfo
	StructInfo           fp.Option[TaggedStruct]
	Tags                 fp.Map[string, Annotation]
}

// int  =>  Eq[int]
// Option[T any] =>   (T) -> Eq[Option[T]]
func (r TypeClassDerive) InstantiatedType(t TypeInfo) TypeInfo {

	ret := r.TypeClassType

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
	Package              genfp.WorkingPackage
	PrimitiveInstancePkg *types.Package
	TypeClass            TypeClass
	TypeClassType        TypeInfo
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

						tags := option.Map(doc, extractTag).OrZero()
						if tags.Contains(directive) {
							info := &types.Info{
								Types: make(map[ast.Expr]types.TypeAndValue),
							}
							//fmt.Printf("check expr = %s\n", types.ExprString(vs.Type))
							types.CheckExpr(pk.Fset, pk.Types, v.End(), vs.Type, info)
							ti := info.Types[vs.Type]

							if nt, ok := ti.Type.(*types.Named); ok && nt.TypeArgs().Len() == 1 {
								if tt, ok := nt.TypeArgs().At(0).(*types.Named); ok && tt.TypeArgs().Len() > 0 {

									// types.Named.Obj() 는  generic type 을 리턴해 준다.
									tcType := typeInfo(tt.Obj().Type())
									//fmt.Printf("tcType = %s\n", tcType)

									return seq.Of(TypeClassDirective{
										Package:              genfp.NewWorkingPackage(pk.Types, pk.Fset, pk.Syntax),
										PrimitiveInstancePkg: nt.Obj().Pkg(),
										TypeClass: TypeClass{
											Name:    tt.Obj().Name(),
											Package: genfp.FromTypesPackage(tt.Obj().Pkg()),
										},
										TypeClassType: tcType,
										TypeArgs:      typeArgs(tt.TypeArgs()),
										Tags:          option.Map(doc, extractTag).OrZero(),
									})
								}
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
				TypeClassType:        v.TypeClassType,
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

	// 타입 클래스 타입
	TypeClassType TypeInfo

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

func (r TypeClassInstance) PackagedName(importSet genfp.ImportSet, working genfp.WorkingPackage) string {
	if r.Package.Path() == working.Path() {
		return r.Name
	}
	pk := importSet.GetImportedName(genfp.FromTypesPackage(r.Package))
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
	//fmt.Printf("find %s, from [%s]\n", name, r.ByName.Keys().MakeString(","))
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
	Ok bool

	CheckConstrainedOf fp.Set[fp.Tuple2[string, string]]
	ParamMapping       fp.Map[string, TypeInfo]
	Error              error
}

func (r ConstraintCheckResult) IsConstraintChecked(t TypeInfo, constraint TypeInfo) bool {
	return r.CheckConstrainedOf.Contains(as.Tuple2(t.ID, constraint.ID))
}

func (r ConstraintCheckResult) ConstraintChecked(t TypeInfo, constraint TypeInfo) ConstraintCheckResult {
	r.CheckConstrainedOf = r.CheckConstrainedOf.Incl(as.Tuple2(t.ID, constraint.ID))
	return r
}

func (r ConstraintCheckResult) Failed(err error) ConstraintCheckResult {
	r.Ok = false
	r.Error = err
	return r
}
func replaceTypeArgs(t types.Type, mapping fp.Map[string, TypeInfo]) types.Type {
	switch tv := t.(type) {
	case *types.Interface:
		embeds := iterate(tv.NumEmbeddeds(), tv.EmbeddedType, func(i int, t types.Type) types.Type {
			return replaceTypeArgs(t, mapping)
		})

		method := iterate(tv.NumExplicitMethods(), tv.ExplicitMethod, func(i int, t *types.Func) *types.Func {
			ns := replaceTypeArgs(t.Signature(), mapping)
			return types.NewFunc(token.NoPos, t.Pkg(), t.Name(), ns.(*types.Signature))
		})

		return types.NewInterfaceType(method, embeds)
	case *types.Signature:
		params := tv.Params()
		nparams := iterate(params.Len(), params.At, func(i int, t *types.Var) *types.Var {
			rt := replaceTypeArgs(t.Type(), mapping)
			return types.NewVar(token.NoPos, t.Pkg(), t.Name(), rt)
		})

		result := tv.Results()
		nresult := iterate(result.Len(), result.At, func(i int, t *types.Var) *types.Var {
			rt := replaceTypeArgs(t.Type(), mapping)
			return types.NewVar(token.NoPos, t.Pkg(), t.Name(), rt)
		})

		rp := iterate(tv.RecvTypeParams().Len(), tv.RecvTypeParams().At, func(i int, t *types.TypeParam) *types.TypeParam {
			return t
		})

		tp := iterate(tv.TypeParams().Len(), tv.TypeParams().At, func(i int, t *types.TypeParam) *types.TypeParam {
			return t
		})

		return types.NewSignatureType(tv.Recv(), rp, tp, types.NewTuple(nparams...), types.NewTuple(nresult...), tv.Variadic())

	case *types.TypeParam:
		ti := mapping.Get(tv.Obj().Name())
		if ti.IsDefined() {
			return ti.Get().Type
		}
	case *types.Named:
		//fmt.Printf("replaceTypeParam %s, params = %s, args = %s\n", t, tv.TypeParams(), tv.TypeArgs())

		if tv.TypeArgs().Len() > 0 {
			newargs := atLenToSeq(tv.TypeArgs()).Map(func(t types.Type) types.Type {
				return replaceTypeArgs(t, mapping)
			})
			ctx := types.NewContext()
			nt, err := types.Instantiate(ctx, tv.Origin(), newargs, true)
			if err != nil {
				fmt.Printf("instantiate error :%s\n", err)
				return t
			}
			return nt
		}
	case *types.Map:
		kt := replaceTypeArgs(tv.Key(), mapping)
		vt := replaceTypeArgs(tv.Elem(), mapping)
		return types.NewMap(kt, vt)
	case *types.Slice:
		nt := replaceTypeArgs(tv.Elem(), mapping)
		return types.NewSlice(nt)
	case *types.Pointer:
		nt := replaceTypeArgs(tv.Elem(), mapping)
		return types.NewPointer(nt)
	}
	return t
}

func CompareTypeAndInferParam(ctx ConstraintCheckResult, param fp.Seq[TypeParam], a TypeInfo, b TypeInfo) ConstraintCheckResult {
	//fmt.Printf("compare %s <-> %s\n", a, b)
	if a.IsTypeParam() {
		if b.IsTypeParam() {
			// TODO : check constraint is same
			aisp := param.Find(func(v TypeParam) bool {
				return v.Name == a.TypeName
			})
			if aisp.IsDefined() {
				ctx.ParamMapping = ctx.ParamMapping.Updated(a.TypeName, b)
			}

			bisp := param.Find(func(v TypeParam) bool {
				return v.Name == b.TypeName
			})
			if bisp.IsDefined() {
				ctx.ParamMapping = ctx.ParamMapping.Updated(b.TypeName, a)
			}
			return ctx
		}

		if !ctx.ParamMapping.Contains(a.TypeName) {
			p := param.Find(func(v TypeParam) bool {
				return v.Name == a.TypeName
			})

			ccheck := b.IsConstrainedOf(ctx, param, GetTypeInfo(p.Get().Constraint))
			if ccheck.Ok {
				ccheck.ParamMapping = ccheck.ParamMapping.Updated(b.TypeName, b)
			}
			return ccheck
		}
		return ctx
	}

	if b.IsTypeParam() {
		if !ctx.ParamMapping.Contains(b.TypeName) {

			p := param.Find(func(v TypeParam) bool {
				return v.Name == b.TypeName
			})
			if p.IsEmpty() {
				return ConstraintCheckResult{
					Error: fp.Error(400, "param %s not found", b.TypeName),
				}
			}

			ccheck := a.IsConstrainedOf(ctx, param, GetTypeInfo(p.Get().Constraint))
			if ccheck.Ok {
				ccheck.ParamMapping = ccheck.ParamMapping.Updated(b.TypeName, a)
			}
			return ccheck
		}
		return ctx
	}

	// a : fp.Option[int]
	// b : fp.Option[T]

	if a.TypeArgs.Size() > 0 && a.TypeArgs.Size() == b.TypeArgs.Size() {
		argChecked := seq.Fold(seq.Zip(a.TypeArgs, b.TypeArgs), ctx, func(prev ConstraintCheckResult, t fp.Tuple2[TypeInfo, TypeInfo]) ConstraintCheckResult {
			return CompareTypeAndInferParam(prev, param, t.I1, t.I2)
		})
		return argChecked
	}

	// TODO : T[int] <-> fp.Option[int] ??

	if a.ID == b.ID {
		return ctx
	}

	return ctx.Failed(fp.Error(400, "type not equal %s with %s", a, b))

}

//	func[T constraint]() Eq[T]  에서
//
// instanceType 이 T 자리에 들어갈 수 있는지 체크하는 함수
//
// func[T constraint]() Eq[Seq[T]]  같은 경우는 해당 사항 없음 .

// param : [T []A, A any] 같은 것
// genericType :  Eq[T] 같은 것, 여기서 비교대상은 Eq[T] 가 아니고 T 임
// typeArgs  :  []int  같은 것
// 이경우  결과로  T : []int , A : int 가 나와야 함.
// param 개수는 더 많을 수 있고,  genericType의  param 개수와 typeArgs 의 개수는 같아야 함.
func ConstraintCheck(ctx ConstraintCheckResult, param fp.Seq[TypeParam], genericType TypeInfo, typeArgs fp.Seq[TypeInfo]) ConstraintCheckResult {

	//fmt.Printf("param = %v, genericType =%v, typeArgs = %v\n", param, genericType, typeArgs)
	// size 가 동일하지 않은 경우
	if genericType.TypeArgs.Size() != typeArgs.Size() {
		return ctx.Failed(fp.Error(400, "type args size not same %s <-> %s", genericType, typeArgs))
	}

	// genericType 아규먼트만 비교
	zipped := iterator.Map(iterator.Zip(seq.Iterator(genericType.TypeArgs), seq.Iterator(typeArgs)), func(t fp.Tuple2[TypeInfo, TypeInfo]) typeCompare {
		return typeCompare{
			genericType: t.I1,
			actualType:  t.I2,
		}
	})

	// Eq[T] 가 아니고,  Eq[Seq[T]]  같은 경우는  체크 불가능
	// Eq[T] 와 Eq[int] 케이스 분리
	paramArgsIt, actualArgs := iterator.Partition(zipped, func(t typeCompare) bool {
		return t.genericType.IsTypeParam()
	})

	// Eq[T] 처럼 param 인것
	// T => int
	paramArgs := paramArgsIt.ToSeq()
	ctx = seq.Fold(paramArgs, ctx, func(c ConstraintCheckResult, v typeCompare) ConstraintCheckResult {
		paramName := v.genericType.Name().Get()

		paramCons := param.Filter(func(p TypeParam) bool {
			return p.Name == paramName
		}).Head()

		if paramCons.IsDefined() {

			if paramCons.Get().IsAny() {
				return ctx
			}
			consType := typeInfo(paramCons.Get().Constraint)
			//fmt.Printf("actual = %s , generic = %s\n", v.actualType, consType)

			// [T []A, A] 인데  typeArgs 가 []int 라면
			// T 가 []A 를 만족하는지 확인해야 함.
			ret := v.actualType.IsConstrainedOf(c, param, consType)
			// if !ret.Ok {
			// 	fmt.Printf("not constraint of actual = %s , generic = %s\n", v.actualType, consType)

			// }
			return ret
		}
		return c.Failed(fp.Error(400, "type param %s not exists", paramName))
	})

	ctx = iterator.Fold(actualArgs, ctx, func(c ConstraintCheckResult, v typeCompare) ConstraintCheckResult {
		return v.actualType.IsInstantiatedOf(c, param, v.genericType)
	})

	if !ctx.Ok {
		return ctx
	}

	for _, p := range paramArgs {
		ctx.ParamMapping = ctx.ParamMapping.Updated(p.genericType.TypeName, p.actualType)
	}

	//fmt.Printf("merge = %s\n", merge)
	paramFound := iterator.Map(iterator.FromSlice(paramArgs), func(v typeCompare) fp.Option[paramVar] {
		paramName := v.genericType.Name().Get()

		// func[T constraint]() Eq[A] 혹은 func() Eq[A] 처럼. type parameter 목록이 잘못된 경우도 불가능
		paramCons := option.Map(param.Filter(func(p TypeParam) bool {
			return p.Name == paramName
		}).Head(), func(p TypeParam) paramVar {
			//fmt.Printf("param %s -> %s, constraint = %T(%s)\n", p.TypeName, v.actualType, p.Constraint, p.Constraint)
			return paramVar{
				typeParam:  types.NewTypeParam(p.TypeName, replaceTypeArgs(p.Constraint, ctx.ParamMapping)),
				actualType: v.actualType,
			}
		})

		return paramCons
	}).ToSeq()

	if len(paramFound) == 0 {
		return ctx
	}

	if !as.Seq(paramFound).ForAll(fp.Option[paramVar].IsDefined) {
		return ctx.Failed(fp.Error(400, "all param not found"))
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

	//fmt.Printf("sig = %s, paramCons = %s, paramIns = %s\n", sig, paramCons, paramIns)
	tctx := types.NewContext()

	_, err := types.Instantiate(tctx, sig, paramIns, true)
	if err == nil {
		mapping := seq.ToGoMap(seq.Map(seq.Map(paramFound, fp.Option[paramVar].Get), func(v paramVar) fp.Tuple2[string, TypeInfo] {
			return as.Tuple2(v.typeParam.Obj().Name(), v.actualType)
		}))

		ctx.ParamMapping = ctx.ParamMapping.Concat(mutable.MapOf(mapping))
		return ctx
	}

	return ctx.Failed(err)
}

func (r TypeClassInstance) Check(t TypeInfo) fp.Option[TypeClassInstance] {

	argType := r.Result.TypeArgs.Head().Get()
	//fmt.Printf("check %s.%s : %t(%s), %d\n", r.Package.Name(), r.Name, argType.IsTypeParam(), argType, argType.TypeArgs.Size())

	// if r.Name == "TupleHCons" {
	// 	fmt.Printf("TupleHCons\n")
	// }
	if argType.IsTypeParam() {

		// func[T any]() Eq[T] 인 경우
		// t 가 T constraint 인지 체크해야 함
		check := ConstraintCheck(ConstraintCheckResult{Ok: true}, r.Type.TypeParam, r.Result, seq.Of(t))
		if check.Ok {

			r.RequiredInstance = seq.Map(r.RequiredInstance, func(v RequiredInstance) RequiredInstance {
				res := v.Type.ReplaceTypeParam(check.ParamMapping)
				r.UsedParam = r.UsedParam.Concat(res.I1)
				v.Type = res.I2
				return v
			})
			r.ParamMapping = check.ParamMapping
			//fmt.Printf("typeparam check.Ok\n")
			return option.Some(r)
		}
		return option.None[TypeClassInstance]()
	}

	// func[T any]() Eq[Tuple[T]] 인경우
	// Tuple[T] 와  t 를 비교해야 함
	if argType.TypeArgs.Size() > 0 {

		check := t.IsInstantiatedOf(ConstraintCheckResult{Ok: true}, r.Type.TypeParam, argType)
		if check.Ok {
			r.RequiredInstance = seq.Map(r.RequiredInstance, func(v RequiredInstance) RequiredInstance {
				res := v.Type.ReplaceTypeParam(check.ParamMapping)
				r.UsedParam = r.UsedParam.Concat(res.I1)
				v.Type = res.I2
				return v
			})
			r.ParamMapping = check.ParamMapping
			//fmt.Printf("check %s.%s : %v\n", r.Package.Name(), r.Name, check.ParamMapping)

			//fmt.Printf("typeargs check.Ok\n")
			return option.Some(r)
		}
		return option.None[TypeClassInstance]()

	}

	if argType.Type.String() == t.Type.String() {
		//fmt.Printf("type check.Ok\n")
		return option.Some(r)
	}
	return option.None[TypeClassInstance]()

}

func (r TypeClassInstancesOfPackage) FindByNamePrefix(namePrefix string, t TypeInfo) fp.Option[TypeClassInstance] {
	return r.Find(t).Filter(func(v TypeClassInstance) bool {
		return strings.HasPrefix(v.Name, namePrefix)
	}).Head()
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
	//fmt.Printf("%s will generated\n", tc.GeneratedInstanceName())
	list := r.tcMap.Get(tc.TypeClass.Id()).OrZero()

	pkPred := func(v TypeClassInstancesOfPackage) bool {
		return v.Package.Path() == tc.Package.Path()
	}

	found := list.Find(pkPred).OrElseGet(func() TypeClassInstancesOfPackage {
		return LoadTypeClassInstance(tc.Package.Package(), tc.TypeClass)
	})

	t := tc.InstantiatedType(tc.DeriveFor.Info)
	//fmt.Printf("will generate %s, %v, %v\n", tc.GeneratedInstanceName(), tc.DeriveFor.Info.TypeParam, tc.DeriveFor.Info.TypeArgs)

	// if t.TypeArgs.Head().IsDefined() {
	// 	fmt.Printf("will generate arg type = %s\n", t.TypeArgs.Head().Get().TypeArgs)
	// }
	ins := TypeClassInstance{
		Package:         tc.Package.Package(),
		Name:            tc.GeneratedInstanceName(),
		Static:          false,
		Implicit:        false,
		Type:            t,
		TypeClassType:   tc.TypeClassType,
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

	// if name == "ShowHlistHCons" {
	// 	fmt.Printf("find ShowHlistHCons\n")
	// }
	ret := iterator.Map(seq.Iterator(r.List), func(p TypeClassInstancesOfPackage) fp.Option[TypeClassInstance] {
		return p.FindByName(name, t)
	}).Filter(fp.Option[TypeClassInstance].IsDefined).NextOption()

	return option.Flatten(ret)
}

func (r TypeClassScope) FindByNamePrefix(namePrefix string, t TypeInfo) fp.Option[TypeClassInstance] {

	// if name == "ShowHlistHCons" {
	// 	fmt.Printf("find ShowHlistHCons\n")
	// }
	ret := iterator.Map(seq.Iterator(r.List), func(p TypeClassInstancesOfPackage) fp.Option[TypeClassInstance] {
		return p.FindByNamePrefix(namePrefix, t)
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
		Name:    v.Name().Get(),
		Package: genfp.FromTypesPackage(v.Pkg),
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

func AsRequiredInstance(v TypeInfo) fp.Option[RequiredInstance] {
	if v.TypeArgs.IsEmpty() {
		return option.None[RequiredInstance]()
	}

	return option.Some(asRequired(v))
}

func AsTypeClassInstance(tc TypeClass, ins types.Object) fp.Option[TypeClassInstance] {
	insType := typeInfo(ins.Type())
	rType := insType.ResultType()
	name := ins.Name()

	if _, ok := ins.(*types.TypeName); ok {
		return option.None[TypeClassInstance]()
	}

	if rType.IsInstanceOf(tc) && rType.TypeArgs.Size() > 0 {

		under := rType.TypeArgs.Head().Get()

		if insType.IsFunc() {

			if insType.NumArgs() == 0 && insType.TypeParam.Size() == 1 {
				return option.Some(TypeClassInstance{
					Package:       ins.Pkg(),
					Name:          name,
					Static:        false,
					Implicit:      true,
					Type:          insType,
					Result:        rType,
					TypeClassType: rType.GenericType(),
					Under:         under,
					Instance:      ins,
					TypeParam:     insType.TypeParam,
				})
				// ret.ByName = ret.ByName.Updated(name, tins)
				// ret.All = ret.All.Append(tins)
			} else {

				fargs := insType.FuncArgs()
				allArgTypeClass := fargs.ForAll(func(v TypeInfo) bool {
					if v.Name().IsDefined() && v.TypeArgs.Size() == 1 {
						if v.Name().Get() == "Eval" && v.Pkg != nil && v.Pkg.Path() == "github.com/csgura/fp/lazy" {
							realv := v.TypeArgs.Head().Get()
							return realv.Name().IsDefined() && realv.TypeArgs.Size() == 1
						}
						return true
					}
					return false
				})

				if allArgTypeClass {

					required := seq.Map(fargs, asRequired)

					return option.Some(TypeClassInstance{
						Package:          ins.Pkg(),
						Name:             name,
						Static:           false,
						Implicit:         false,
						Type:             insType,
						Result:           rType,
						TypeClassType:    rType.GenericType(),
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
						TypeClassType: rType.GenericType(),
						TypeParam:     insType.TypeParam,
						Under:         under,
						Instance:      ins,
					})
					//ret.OtherFuncs = ret.OtherFuncs.Updated(name, tins)
				}
			}
		} else {
			return option.Some(TypeClassInstance{
				Package:       ins.Pkg(),
				Name:          name,
				Static:        true,
				Implicit:      false,
				Type:          insType,
				Result:        insType,
				TypeClassType: insType.GenericType(),
				Under:         under,
				Instance:      ins,
			})
			// ret.ByName = ret.ByName.Updated(name, tins)
			// ret.FixedByType = ret.FixedByType.Updated(under.Type.String(), tins)
		}
	}
	return option.None[TypeClassInstance]()

}

func loadPackage(path string) genfp.WorkingPackage {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	//fmt.Printf("load package %s\n", path)
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return nil
	}
	if len(pkgs) == 0 {
		return nil
	}

	return genfp.NewWorkingPackage(pkgs[0].Types, pkgs[0].Fset, pkgs[0].Syntax)
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
	names := pk.Scope().Names()
	if len(names) == 0 {
		working := loadPackage(pk.Path())
		if working == nil {
			return ret
		}
		pk = working.Package()
		ret.Package = pk
	}

	for _, name := range pk.Scope().Names() {
		ins := pk.Scope().Lookup(name)
		AsTypeClassInstance(tc, ins).Foreach(func(tins TypeClassInstance) {
			//fmt.Printf("tc found %s, type = %T, underlying = %T\n", tins.Name, tins.Type.Type, tins.Type.Type.Underlying())
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
