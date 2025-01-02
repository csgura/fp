package main

import (
	"fmt"
	"go/types"
	"os"
	"runtime"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/either"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/xtr"
	"golang.org/x/tools/go/packages"
)

type TypeClassInstanceGenerated struct {
	Derive metafp.TypeClassDerive
	Expr   SummonExpr
}

type TypeClassSummonContext struct {
	w                     genfp.ImportSet
	fpPkg                 fp.Option[*types.Package]
	tcCache               *metafp.TypeClassInstanceCache
	summoned              fp.Map[string, TypeClassInstanceGenerated]
	loopCheck             fp.Set[string]
	recursiveGen          fp.Seq[metafp.TypeClassDerive]
	implicitTypeInference bool
}

type CurrentContext struct {
	working      genfp.WorkingPackage
	tc           metafp.TypeClassDerive
	primScope    metafp.TypeClassScope
	workingScope metafp.TypeClassScope
	recursiveGen bool
}

type GenericRepr struct {
	//	ReprType     func() string
	Kind         string
	ToReprExpr   func() string
	FromReprExpr func() string
	ReprExpr     func() SummonExpr
}

type ParamInstance struct {
	ArgName   string
	ParamName string
	TypeClass metafp.TypeClass
}

var EqParamInstance = eq.New(func(a, b ParamInstance) bool {
	if a.ArgName != b.ArgName {
		return false
	}

	if a.TypeClass.Id() != b.TypeClass.Id() {
		return false
	}

	return true
})

func (r ParamInstance) Expr(importSet genfp.ImportSet, working genfp.WorkingPackage) string {
	return fmt.Sprintf("%s %s[%s]", r.ArgName, r.TypeClass.PackagedName(importSet, working), r.ParamName)
}

type SummonExpr struct {
	expr          string
	paramInstance fp.Seq[ParamInstance]
}

func (r SummonExpr) Expr() string {
	return r.expr
}
func (r SummonExpr) String() string {
	return r.expr
}

func (r SummonExpr) ParamInstance() fp.Seq[ParamInstance] {
	return r.paramInstance
}

func collectSummonExpr(list fp.Seq[SummonExpr]) SummonExpr {
	expr := seq.Map(list, SummonExpr.Expr).MakeString(",")
	paramList := seq.Reduce(seq.Map(list, SummonExpr.ParamInstance), MergeSeqDistinct(EqParamInstance))
	return SummonExpr{
		expr:          expr,
		paramInstance: paramList,
	}
}

type ArgumentInstance struct {
	instanceOf metafp.TypeInfo
	name       string
	typeParam  fp.Option[metafp.TypeClass]
}

func (r ArgumentInstance) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	param := option.Map(r.typeParam, func(v metafp.TypeClass) ParamInstance {
		return ParamInstance{
			ArgName:   r.name,
			TypeClass: v,
			ParamName: r.instanceOf.Name().Get(),
		}
	}).ToSeq()

	return SummonExpr{
		expr:          r.name,
		paramInstance: param,
	}
}

func (r ArgumentInstance) Required() fp.Seq[metafp.RequiredInstance] {
	return nil
}

type DefinedInstance struct {
	instance metafp.TypeClassInstance
	local    bool
}

func (r DefinedInstance) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	if r.instance.Package == nil || r.instance.Package.Path() == workingPkg.Path() {
		return SummonExpr{
			expr: r.instance.Name,
		}
	}

	pk := w.GetImportedName(genfp.FromTypesPackage(r.instance.Package))

	return SummonExpr{
		expr: fmt.Sprintf("%s.%s", pk, r.instance.Name),
	}
}

func (r DefinedInstance) Instance() metafp.TypeClassInstance {
	return r.instance
}

func (r DefinedInstance) Required() fp.Seq[metafp.RequiredInstance] {
	return r.instance.RequiredInstance
}

func (r DefinedInstance) IsLocal() bool {
	return r.local
}

type NotDefinedInstance struct {
	instanceOf metafp.TypeInfo
	name       string
	required   fp.Seq[metafp.RequiredInstance]
}

func (r NotDefinedInstance) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	return SummonExpr{
		expr: r.name,
	}
}
func (r NotDefinedInstance) Required() fp.Seq[metafp.RequiredInstance] {
	return r.required
}

type SummonExprInstance struct {
	expr func(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr
}

func (r SummonExprInstance) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	return r.expr(w, workingPkg)
}

func (r SummonExprInstance) Required() fp.Seq[metafp.RequiredInstance] {
	return nil
}

// Argument or Defined or NotDefined or Expr
type lookupTarget struct {
	target fp.Either[NotDefinedInstance, fp.Either[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance]]]
}

func (r lookupTarget) instance() fp.Option[metafp.TypeClassInstance] {
	if r1, ok := either.Fold(r.target, option.ConstNone, option.Some).Unapply(); ok {
		if r2, ok := either.Fold(r1, option.ConstNone, option.Some).Unapply(); ok {
			return option.Map(either.Fold(r2, option.ConstNone, option.Some), DefinedInstance.Instance)
		}
	}
	return option.None[metafp.TypeClassInstance]()
}

func (r lookupTarget) required() fp.Seq[metafp.RequiredInstance] {
	return either.Fold(
		r.target,
		NotDefinedInstance.Required,
		as.Func3(either.Fold[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance], fp.Seq[metafp.RequiredInstance]]).ApplyLast2(
			SummonExprInstance.Required,
			as.Func3(either.Fold[ArgumentInstance, DefinedInstance, fp.Seq[metafp.RequiredInstance]]).ApplyLast2(
				ArgumentInstance.Required,
				DefinedInstance.Required,
			),
		),
	)
}

func (r lookupTarget) isGivenAny() bool {
	return option.Map(r.instance(), metafp.TypeClassInstance.IsGivenAny).OrElse(false)
}

func (r lookupTarget) isLocal() bool {
	return either.Fold(
		r.target,
		fp.Const[NotDefinedInstance](false),
		as.Func3(either.Fold[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance], bool]).ApplyLast2(
			fp.Const[SummonExprInstance](false),
			as.Func3(either.Fold[ArgumentInstance, DefinedInstance, bool]).ApplyLast2(
				fp.Const[ArgumentInstance](false),
				DefinedInstance.IsLocal,
			),
		),
	)
}

func (r lookupTarget) isFunc() bool {

	if r.target.IsLeft() {
		return true
	}

	instance := r.instance()
	if instance.IsDefined() {
		return !instance.Get().Static
	}
	return false
}

func (r lookupTarget) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {

	return either.Fold(
		r.target,
		as.Func3(NotDefinedInstance.instanceExpr).ApplyLast2(w, workingPkg),
		as.Func3(either.Fold[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance], SummonExpr]).ApplyLast2(
			as.Func3(SummonExprInstance.instanceExpr).ApplyLast2(w, workingPkg),
			as.Func3(either.Fold[ArgumentInstance, DefinedInstance, SummonExpr]).ApplyLast2(
				as.Func3(ArgumentInstance.instanceExpr).ApplyLast2(w, workingPkg),
				as.Func3(DefinedInstance.instanceExpr).ApplyLast2(w, workingPkg),
			),
		),
	)
}

func (r *TypeClassSummonContext) typeclassInstanceMust(ctx CurrentContext, req metafp.RequiredInstance, name string) lookupTarget {

	genName := req.TypeClass.Name + publicName(name)

	if req.Type.Pkg != nil && req.Type.Pkg.Path() != "" && !isSamePkg(ctx.working, genfp.FromTypesPackage(req.Type.Pkg)) {
		genName = req.TypeClass.Name + publicName(req.Type.Pkg.Name()) + publicName(name)
	}

	f := req.Type

	ret := NotDefinedInstance{
		instanceOf: f,
		name:       genName,
		required: seq.Map(f.TypeArgs, func(v metafp.TypeInfo) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: req.TypeClass,
				Type:      v,
			}
		}),
	}
	// 에러내기 위해 사용. instance 가 None

	return lookupTarget{
		target: either.NotRight[fp.Either[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance]]](ret),
	}
}

// f 는 Eq 쌓이지 않은 타입
// Eq[T] 같은거 아님
func (r *TypeClassSummonContext) lookupTypeClassInstanceLocalDeclared(ctx CurrentContext, req metafp.RequiredInstance, name ...string) fp.Option[metafp.TypeClassInstance] {

	f := req.Type

	scope := ctx.workingScope
	if req.TypeClass.Id() != ctx.tc.TypeClass.Id() {
		scope = r.tcCache.GetLocal(ctx.working.Package(), req.TypeClass)
	}
	itr := seq.Iterator(seq.FlatMap(name, func(v string) fp.Seq[string] {
		if f.Pkg != nil && ctx.working.Path() != f.Pkg.Path() {
			return []string{
				req.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(v),
				req.TypeClass.Name + publicName(v),
			}

		}
		return []string{req.TypeClass.Name + publicName(v)}
	}))

	ins := iterator.FlatMap(itr, func(v string) fp.Iterator[metafp.TypeClassInstance] {
		res := scope.FindByName(v, f)
		// fmt.Printf("FindByName %s = %s\n", v, res)
		return option.Iterator(res)
	})

	ins = iterator.Map(ins, func(tci metafp.TypeClassInstance) metafp.TypeClassInstance {
		if tci.WillGeneratedBy.IsDefined() {
			// 현재 생성 중인 인스턴스를 참조하는 경우에
			// 만약 그게 함수라면, type param 개수와 , 실제 아규먼트 개수가 일치하지 않을 수 있다.
			// 왜냐하면  Map[K,V] 같은 타입의 monoid 는 K의 monoid 가 필요 없기 때문.

			// 그래서 이런 타입을 만난 경우에,   먼저 생성해 본다.
			// 그런데,   A -> B -> A  순으로  순환 참조가 있다면,  생성을 해볼 수가 없다.
			expr := r.summonVar(tci.WillGeneratedBy.Get())
			if expr.IsDefined() && tci.RequiredInstance.Size() != expr.Get().paramInstance.Size() {

				// paramInstance 에  실제 인스턴스의 아규먼트 목록이 있다.
				// RequiredInstance 를  실제 아규먼트로 변경 해주어야 함.
				tci.RequiredInstance = seq.Map(expr.Get().paramInstance, func(v ParamInstance) metafp.RequiredInstance {
					return metafp.RequiredInstance{
						TypeClass: v.TypeClass,
						Type:      tci.ParamMapping.Get(v.ParamName).Get(),
					}
				})

				// 시용되지 않은 type param 추출
				notused := tci.ParamMapping.Keys().FilterNot(func(name string) bool {
					return expr.Get().paramInstance.Exists(func(v ParamInstance) bool {
						return name == v.ParamName
					})
				}).ToSeq()

				// 사용되지 않았다고 표시
				as.Seq(notused).Foreach(func(v string) {
					tci.UsedParam = tci.UsedParam.Excl(v)
				})
			}
		}
		return tci
	})

	if f.TypeArgs.Size() > 0 {
		ins = seq.Iterator(scope.Find(f)).Concat(ins)
	} else {
		ins = ins.Concat(seq.Iterator(scope.Find(f)))
	}

	ins = ins.Filter(func(tci metafp.TypeClassInstance) bool {

		if tci.IsGivenAny() && ctx.recursiveGen && isRecursiveDerivable(req) {
			return false
			//fmt.Printf("%s is recursive derivable\n", req.Type)
		}

		return r.checkRequired(ctx, tci.RequiredInstance)
	})

	// instance 가 있는 경우 , instance 가 Some
	ret := ins.NextOption()

	return ret
}

func (r *TypeClassSummonContext) lookupHConsMust(ctx CurrentContext, tc metafp.TypeClass) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(ctx, tc, "HCons")
	if ret.IsDefined() {
		return ret.Get()
	}

	ret = r.lookupTypeClassFunc(ctx, tc, "HListCons")
	if ret.IsDefined() {
		return ret.Get()
	}
	nameWithTc := tc.Name + "HCons"

	return metafp.TypeClassInstance{
		Package: ctx.working.Package(),
		Name:    nameWithTc,
		Static:  true,
	}
}

func (r *TypeClassSummonContext) lookupHNilMust(ctx CurrentContext, tc metafp.TypeClass) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(ctx, tc, "HNil")
	if ret.IsDefined() {
		return ret.Get()
	}

	ret = r.lookupTypeClassFunc(ctx, tc, "HListNil")
	if ret.IsDefined() {
		return ret.Get()
	}
	nameWithTc := tc.Name + "HNil"

	return metafp.TypeClassInstance{
		Package: ctx.working.Package(),
		Name:    nameWithTc,
		Static:  true,
	}
}

// TODO: 사용을 줄이고,  typecheck 를 할 필요 있음
// 현재 사용하는 경우
// HCons, HListCons, HNil, HListNil ,
// Tuple%d , Labelled%d, HConsLabelled%d ,
// StructHCons, StructHNil
// TupleHCons, TupleHNil
// Generic , IMap, Map
func (r *TypeClassSummonContext) lookupTypeClassFunc(ctx CurrentContext, tc metafp.TypeClass, name string) fp.Option[metafp.TypeClassInstance] {
	nameWithTc := tc.Name + name

	workingScope := ctx.workingScope
	primScope := ctx.primScope
	if ctx.tc.TypeClass.Id() != tc.Id() {
		primScope = r.tcCache.GetImported(tc)
		workingScope = r.tcCache.GetLocal(ctx.working.Package(), tc)
	}

	ins := workingScope.FindFunc(nameWithTc)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindFunc(nameWithTc)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindFunc(name)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get().RequiredInstance) {
		return ins
	}

	return option.None[metafp.TypeClassInstance]()
}

func (r *TypeClassSummonContext) lookupTupleTypeClassFunc(ctx CurrentContext, tc metafp.TypeClass, name string, tupleArgs fp.Seq[metafp.TypeInfoExpr]) fp.Option[metafp.TypeClassInstance] {
	ret := r.lookupTypeClassFunc(ctx, tc, name)
	if ret.IsDefined() {
		tc := ret.Get()
		tpType := tc.Result.TypeArgs.Head()
		if tpType.IsDefined() {
			tpt := tpType.Get()

			named := tpt.Type.(*types.Named)
			//fmt.Printf("type = %v\n", named.Origin())

			origin := named.Origin()
			//fmt.Printf("tparms = %v\n", origin.TypeParams().At(0))

			ctx := types.NewContext()

			targs := seq.Map(tupleArgs, func(v metafp.TypeInfoExpr) types.Type {
				return v.Type.Type
			})
			it, err := types.Instantiate(ctx, origin, targs, false)
			if err == nil {
				//fmt.Printf("it= %v\n", it)

				checked := tc.Check(metafp.GetTypeInfo(it))
				//fmt.Printf("checked = %v, required = %v\n", checked.Get().ParamMapping, checked.Get().RequiredInstance)
				return checked
			}

		}
	}

	return ret
}

func (r *TypeClassSummonContext) lookupTypeClassFuncMust(ctx CurrentContext, tc metafp.TypeClass, name string) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(ctx, tc, name)
	if ret.IsDefined() {
		return ret.Get()
	}

	nameWithTc := tc.Name + name

	return metafp.TypeClassInstance{
		Package: ctx.working.Package(),
		Name:    nameWithTc,
		Static:  true,
	}
}

func (r *TypeClassSummonContext) lookupTypeClassInstancePrimitivePkgLazy(ctx CurrentContext, req metafp.RequiredInstance, name ...string) func() fp.Option[metafp.TypeClassInstance] {
	return func() fp.Option[metafp.TypeClassInstance] {
		return r.lookupTypeClassInstancePrimitivePkg(ctx, req, name...)
	}
}

func (r *TypeClassSummonContext) checkRequired(ctx CurrentContext, required fp.Seq[metafp.RequiredInstance]) bool {
	for _, v := range required {
		if v.Type.IsTuple() {
			req := seq.Map(v.Type.TypeArgs, func(t metafp.TypeInfo) metafp.RequiredInstance {
				return metafp.RequiredInstance{
					TypeClass: v.TypeClass,
					Type:      t,
				}
			})
			res := r.checkRequired(ctx, req)
			if !res {
				return false
			}

		} else {
			// TODO: summonArgs에서 다시  lookup 하는 코드 있음.
			res := r.lookupTypeClassInstance(ctx, v)
			if res.target.IsLeft() {
				if ctx.recursiveGen && v.TypeClass.Id() == ctx.tc.TypeClass.Id() {
					named := res.target.Left().instanceOf.AsNamed()
					if named.IsDefined() {

						deriveFor := named.Get().Info

						vt := metafp.LookupStruct(deriveFor.Pkg, deriveFor.Name().Get())

						tc := metafp.TypeClassDerive{
							Package:              ctx.tc.Package,
							PrimitiveInstancePkg: ctx.tc.PrimitiveInstancePkg,
							TypeClass:            ctx.tc.TypeClass,
							TypeClassType:        ctx.tc.TypeClassType,
							DeriveFor:            named.Get(),
							StructInfo:           vt,
							Tags:                 ctx.tc.Tags,
						}

						if !r.tcCache.IsWillGenerated(tc) {
							r.recursiveGen = append(r.recursiveGen, tc)
							r.tcCache.WillGenerated(tc)
						}
						continue
					}
				}
				return false
			}
		}
	}
	return true
}

func isValueGeneratedType(t metafp.TypeInfo) bool {

	return t.Method.Contains("Unapply") && t.Method.Contains("Builder")
}

func isRecursiveDerivable(req metafp.RequiredInstance) bool {
	if req.Type.IsNamed() {
		namedType := req.Type.AsNamed().Get()
		if namedType.Underlying.IsStruct() {
			if namedType.Underlying.Fields().Exists(metafp.StructField.Public) || isValueGeneratedType(req.Type) {
				return true
			} else {
				return false
			}
		}
		return true
	}
	return false

}

func (r *TypeClassSummonContext) lookupTypeClassInstancePrimitivePkg(ctx CurrentContext, req metafp.RequiredInstance, name ...string) fp.Option[metafp.TypeClassInstance] {

	scope := ctx.primScope
	if ctx.tc.TypeClass.Id() != req.TypeClass.Id() {
		scope = r.tcCache.GetImported(req.TypeClass)
	}
	f := req.Type
	itr := seq.Iterator(seq.FlatMap(name, func(v string) fp.Seq[string] {
		ret := seq.Of(
			req.TypeClass.Name+publicName(v),
			publicName(v),
		)
		if f.Pkg != nil {
			return seq.Of(
				req.TypeClass.Name+publicName(f.Pkg.Name())+publicName(v),
				publicName(f.Pkg.Name())+publicName(v),
			).Concat(ret)
		}
		return ret
	}))

	ins := iterator.FlatMap(itr, func(v string) fp.Iterator[metafp.TypeClassInstance] {
		return option.Iterator(scope.FindByName(v, f))
	})

	if f.TypeArgs.Size() > 0 {
		ins = seq.Iterator(scope.Find(f)).Concat(ins)
	} else {
		ins = ins.Concat(seq.Iterator(scope.Find(f)))
	}

	ins = ins.Filter(func(tci metafp.TypeClassInstance) bool {
		//fmt.Printf("result for %s[%s] is %s, is given %t\n", req.TypeClass.Name, req.Type, tci.Name, tci.IsGivenAny())

		if tci.IsGivenAny() && ctx.recursiveGen && isRecursiveDerivable(req) {
			return false
			//fmt.Printf("%s is recursive derivable\n", req.Type)
		}
		return r.checkRequired(ctx, tci.RequiredInstance)
	})

	// instance 가 있는 경우 , instance 가 Some
	return ins.NextOption()

}

func (r *TypeClassSummonContext) lookupTypeClassInstanceTypePkg(ctx CurrentContext, req metafp.RequiredInstance, name string) fp.Option[metafp.TypeClassInstance] {

	f := req.Type
	if f.Pkg != nil && f.Pkg.Path() != ctx.working.Path() {

		name := req.TypeClass.Name + publicName(name)
		obj := f.Pkg.Scope().Lookup(name)

		if obj != nil {

			ti := metafp.GetTypeInfo(obj.Type())
			rhsType := ti.ResultType()
			if rhsType.IsInstanceOf(ctx.tc.TypeClass) {
				return metafp.AsTypeClassInstance(req.TypeClass, obj)

			}

		}
	}

	return option.None[metafp.TypeClassInstance]()
}

func (r *TypeClassSummonContext) namedLookup(ctx CurrentContext, req metafp.RequiredInstance, name string) fp.Option[metafp.TypeClassInstance] {

	localInsOpt := r.lookupTypeClassInstanceLocalDeclared(ctx, req, name)
	if localInsOpt.IsDefined() {

		localIns := localInsOpt.Get()
		if !localIns.IsGivenAny() {
			return localInsOpt

		}

	}

	ret := r.lookupTypeClassInstanceTypePkg(ctx, req, name).
		Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, name))

	if localInsOpt.IsDefined() && ret.IsDefined() {
		retIns := ret.Get()
		if retIns.IsGivenAny() {
			return localInsOpt
		}

	}

	return ret.OrOption(localInsOpt)

}

func TypeClassInstanceToLookupTarget(a metafp.TypeClassInstance) lookupTarget {
	return lookupTarget{
		target: either.Right[NotDefinedInstance](either.Right[SummonExprInstance](either.Right[ArgumentInstance](DefinedInstance{
			instance: a,
		}))),
	}
}
func (r *TypeClassSummonContext) orMust(ctx CurrentContext, req metafp.RequiredInstance, name string, ins fp.Option[metafp.TypeClassInstance]) lookupTarget {
	lt := option.Map(ins, TypeClassInstanceToLookupTarget)
	return lt.OrElse(r.typeclassInstanceMust(ctx, req, name))

}
func (r *TypeClassSummonContext) namedLookupMust(ctx CurrentContext, req metafp.RequiredInstance, name string) lookupTarget {
	return r.orMust(ctx, req, name, r.namedLookup(ctx, req, name))
}

// func (r *TypeClassSummonContext) lookupPrimitiveTypeClassInstance(ctx CurrentContext, req metafp.RequiredInstance, name ...string) lookupTarget {
// 	ret := r.lookupTypeClassInstanceLocalDeclared(ctx, req, name...).Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, name...))

// 	return ret.OrElse(r.typeclassInstanceMust(ctx, req, name[0]))

// }

// 타입 추론이 가능한지 따지는 함수
func (r *TypeClassSummonContext) typeParamStringOfLookupTarget(ctx CurrentContext, lt lookupTarget) fp.Option[string] {

	if lt.instance().IsDefined() {
		ins := lt.instance().Get()

		// 타입 추론이 가능하려면,  모든 타입 파라미터가, 아규먼트에서 사용되어야 한다.
		possible := seq.Map(ins.TypeParam, func(v metafp.TypeParam) bool {
			return ins.UsedParam.Contains(v.Name)
		})

		if possible.ForAll(fp.Id) {
			return option.None[string]()
		}
		// 전부 사용되지 않아 타입 추론이 불가능하다면
		// 타입을 명시한다.

		_, notPossible := seq.Span(possible.Reverse(), fp.Id)
		if !r.implicitTypeInference {
			notPossible = possible
		}

		ret := seq.Map(ins.TypeParam.Take(notPossible.Size()), func(v metafp.TypeParam) string {
			return option.Map(ins.ParamMapping.Get(v.Name), func(v metafp.TypeInfo) string {
				return r.w.TypeName(ctx.working, v.Type)
			}).OrElse(v.Name)
		}).MakeString(",")
		return option.Some(ret)

	}

	return option.None[string]()
}

func (r *TypeClassSummonContext) typeParamString(ctx CurrentContext, ins metafp.TypeClassInstance) fp.Option[string] {

	// 타입 추론이 가능하려면,  모든 타입 파라미터가, 아규먼트에서 사용되어야 한다.
	possible := seq.Map(ins.TypeParam, func(v metafp.TypeParam) bool {
		return ins.UsedParam.Contains(v.Name)
	})

	if possible.ForAll(fp.Id) {
		return option.None[string]()
	}

	_, notPossible := seq.Span(possible.Reverse(), fp.Id)
	if !r.implicitTypeInference {
		notPossible = possible
	}

	// 전부 사용되지 않아 타입 추론이 불가능하다면
	// 타입을 명시한다.
	ret := seq.Map(ins.TypeParam.Take(notPossible.Size()), func(v metafp.TypeParam) string {
		return option.Map(ins.ParamMapping.Get(v.Name), func(v metafp.TypeInfo) string {
			return r.w.TypeName(ctx.working, v.Type)
		}).OrElse(v.Name)
	}).MakeString(",")
	return option.Some(ret)

}

func MergeSeqDistinct[T any](eqt fp.Eq[T]) fp.Monoid[fp.Seq[T]] {
	return monoid.New(
		func() fp.Seq[T] {
			return seq.Of[T]()
		},
		func(a fp.Seq[T], b fp.Seq[T]) fp.Seq[T] {
			bf := b.FilterNot(func(bv T) bool {
				return a.Find(func(av T) bool {
					return eqt.Eqv(av, bv)
				}).IsDefined()
			})
			return a.Concat(bf)
		},
	)
}

func (r *TypeClassSummonContext) summonArgs(ctx CurrentContext, args fp.Seq[metafp.RequiredInstance]) SummonExpr {
	list := seq.Map(args, func(t metafp.RequiredInstance) SummonExpr {
		// TODO: checkRequired 에서  lookup 하는 코드 있음. checkRequired 에서 한번 했으면 안하게 할 필요 있음.
		ret := r.summonRequired(ctx, t)
		if t.Lazy {
			lazypk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/lazy", "lazy"))
			expr := fmt.Sprintf(`%s.Call( func() %s[%s] {
				return %s
			})`, lazypk, t.TypeClass.PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type.Type), ret.expr)

			return newSummonExpr(expr, ret.paramInstance)
		}
		return ret
	})

	return collectSummonExpr(list)
}

func newSummonExpr(expr string, params ...fp.Seq[ParamInstance]) SummonExpr {
	return SummonExpr{
		expr:          expr,
		paramInstance: seq.Reduce(as.Seq(params), MergeSeqDistinct(EqParamInstance)),
	}
}

func instanceExprOfTypeClassInstance(r metafp.TypeClassInstance, w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	if r.Package == nil || r.Package.Path() == workingPkg.Path() {
		return SummonExpr{
			expr: r.Name,
		}
	}

	pk := w.GetImportedName(genfp.FromTypesPackage(r.Package))

	return SummonExpr{
		expr: fmt.Sprintf("%s.%s", pk, r.Name),
	}
}

func (r *TypeClassSummonContext) exprTypeClassInstance(ctx CurrentContext, lt metafp.TypeClassInstance) SummonExpr {
	//fmt.Printf("lt : %s, %v\n", lt.instance(), lt.required())

	if len(lt.RequiredInstance) > 0 {
		list := r.summonArgs(ctx, lt.RequiredInstance)

		instanceExpr := instanceExprOfTypeClassInstance(lt, r.w, ctx.working)
		tpstr := r.typeParamString(ctx, lt)
		if tpstr.IsDefined() {
			//fmt.Printf("%s param infer not possible = %s \n", lt.Name, lt.ParamMapping)

			return newSummonExpr(fmt.Sprintf("%s[%s](%s)", instanceExpr.expr, tpstr.Get(), list.expr), instanceExpr.paramInstance, list.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s(%s)", instanceExpr.expr, list.expr), instanceExpr.paramInstance, list.paramInstance)

		}
	}

	if !lt.Static && len(lt.RequiredInstance) == 0 {
		instanceExpr := instanceExprOfTypeClassInstance(lt, r.w, ctx.working)

		tpstr := r.typeParamString(ctx, lt)
		if tpstr.IsDefined() {
			return newSummonExpr(fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get()), instanceExpr.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s()", instanceExpr), instanceExpr.paramInstance)
		}

	}

	return instanceExprOfTypeClassInstance(lt, r.w, ctx.working)

}

func (r *TypeClassSummonContext) exprLookupTarget(ctx CurrentContext, lt lookupTarget) SummonExpr {
	//fmt.Printf("lt : %s, %v\n", lt.instance(), lt.required())

	if len(lt.required()) > 0 {
		list := r.summonArgs(ctx, lt.required())

		instanceExpr := lt.instanceExpr(r.w, ctx.working)
		tpstr := r.typeParamStringOfLookupTarget(ctx, lt)
		if tpstr.IsDefined() {
			//fmt.Printf("%s param infer not possible = %s \n", lt.name, lt.instance.Get().ParamMapping)

			return newSummonExpr(fmt.Sprintf("%s[%s](%s)", instanceExpr.expr, tpstr.Get(), list.expr), instanceExpr.paramInstance, list.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s(%s)", instanceExpr.expr, list.expr), instanceExpr.paramInstance, list.paramInstance)

		}
	}

	if lt.isFunc() && len(lt.required()) == 0 {
		instanceExpr := lt.instanceExpr(r.w, ctx.working)

		tpstr := r.typeParamStringOfLookupTarget(ctx, lt)
		if tpstr.IsDefined() {
			return newSummonExpr(fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get()), instanceExpr.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s()", instanceExpr), instanceExpr.paramInstance)
		}

	}

	return lt.instanceExpr(r.w, ctx.working)

}

func (r *TypeClassSummonContext) exprTypeClassMember(ctx CurrentContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typeArgs fp.Seq[metafp.TypeInfoExpr], fieldOf fp.Option[metafp.TypeInfo]) SummonExpr {
	if len(typeArgs) > 0 {
		list := r.summonArgs(ctx, seq.Map(typeArgs, func(t metafp.TypeInfoExpr) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: tc,
				Type:      t.Type,
				FieldOf:   fieldOf,
			}
		}))

		return newSummonExpr(fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, ctx.working), list), list.paramInstance)
	}

	return newSummonExpr(lt.PackagedName(r.w, ctx.working))

}

func (r *TypeClassSummonContext) exprTypeClassMemberLabelled(ctx CurrentContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typePkg *types.Package, structName string, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfoExpr], genLabelled bool) SummonExpr {
	if len(typeArgs) > 0 {
		list := collectSummonExpr(seq.Map(seq.Zip(typeArgs, names), func(t fp.Tuple2[metafp.TypeInfoExpr, string]) SummonExpr {
			return r.summonFpNamed(ctx, tc, typePkg, structName, t.I2, t.I1, genLabelled)
		}))

		return newSummonExpr(fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, ctx.working), list), list.paramInstance)
	}

	return newSummonExpr(lt.PackagedName(r.w, ctx.working))

}

func (r *TypeClassSummonContext) lookupTypeClassInstance(ctx CurrentContext, req metafp.RequiredInstance) lookupTarget {
	f := req.Type

	switch at := f.Type.(type) {
	case *types.TypeParam:
		ret := ArgumentInstance{
			instanceOf: f,
			name:       privateName(req.TypeClass.Name) + at.Obj().Name(),
			typeParam:  option.Some(req.TypeClass),
		}
		// type parameter 의 instance ,
		// type parameter의 타입이 컴파일 타입에 정해지지 않기 때문에 instance를 아규먼트로 받아야 하는 값
		return lookupTarget{
			target: either.Right[NotDefinedInstance](either.Right[SummonExprInstance](either.NotRight[DefinedInstance](ret))),
		}
	case *types.Named:
		if at.Obj().Pkg().Path() == "github.com/csgura/fp/hlist" {
			//fmt.Printf("lookup named hlist %s\n", req.Type)

			if at.Obj().Name() == "Nil" {
				return option.Map(r.lookupTypeClassInstanceLocalDeclared(ctx, req, "HNil", "HListNil").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, "HNil", "HListNil")), TypeClassInstanceToLookupTarget).OrElse(r.typeclassInstanceMust(ctx, req, "HNil"))

			} else if at.Obj().Name() == "Cons" {
				return option.Map(r.lookupTypeClassInstanceLocalDeclared(ctx, req, "HCons", "HListCons").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, "HCons", "HListCons")), TypeClassInstanceToLookupTarget).OrElse(r.typeclassInstanceMust(ctx, req, "HCons"))
			}
		}
		return r.namedLookupMust(ctx, req, at.Obj().Name())
	case *types.Array:
		panic(fmt.Sprintf("can't summon array type, while deriving %s[%s]", req.TypeClass.Name, ctx.tc.DeriveFor.Name))
		//return r.namedLookup(f, "Array")
	case *types.Slice:
		if at.Elem().String() == "byte" {
			bytesInstance := r.namedLookupMust(ctx,
				metafp.RequiredInstance{
					TypeClass: req.TypeClass,
					Type: metafp.TypeInfo{
						Pkg:      f.Pkg,
						Type:     f.Type,
						TypeArgs: nil,
					}}, "Bytes")

			if bytesInstance.target.IsRight() {
				return bytesInstance
			}
			return r.namedLookupMust(ctx, req, "Slice")
		}
		return r.namedLookupMust(ctx, req, "Slice")
	case *types.Map:
		return r.namedLookupMust(ctx, req, "GoMap")
	case *types.Pointer:
		return r.namedLookupMust(ctx, req, "Ptr")
	case *types.Basic:
		return r.namedLookupMust(ctx, req, at.Name())
	case *types.Struct:
		fields := f.Fields()

		if fields.ForAll(metafp.StructField.Public) || req.FieldOf.Exists(fp.Test(metafp.TypeInfo.IsSamePkg, ctx.working)) {
			ret := func(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
				return r.summonUntypedStruct(ctx, req.TypeClass, f, fields)
			}
			return lookupTarget{
				target: either.Right[NotDefinedInstance](either.NotRight[fp.Either[ArgumentInstance, DefinedInstance]](SummonExprInstance{ret})),
			}
		} else {
			fmt.Printf("fieldOf = %v\n", req.FieldOf)
			panic(fmt.Sprintf("can't summon unnamed struct type %v containing private field, while deriving %s[%s]", f.Type, ctx.tc.TypeClass.Name, ctx.tc.DeriveFor.Name))
		}

	case *types.Interface:
		if f.IsAny() {
			return r.namedLookupMust(ctx, req, "Given")
		}
		panic(fmt.Sprintf("can't summon unnamed interface type %v, while deriving %s[%s]", f.Type, ctx.tc.TypeClass.Name, ctx.tc.DeriveFor.Name))
	case *types.Chan:
		panic(fmt.Sprintf("can't summon unnamed chan type, while deriving %s[%s]", ctx.tc.TypeClass.Name, ctx.tc.DeriveFor.Name))

	}
	return r.namedLookupMust(ctx, req, f.Type.String())
}

// v.A , v.B
func (r *TypeClassSummonContext) structUnapplyExpr(ctx CurrentContext, named fp.Option[metafp.NamedTypeInfo], fields fp.Seq[metafp.StructField], varexpr string) string {
	hasUnapply := option.Map(named, func(v metafp.NamedTypeInfo) bool { return v.Info.Method.Contains("Unapply") }).OrElse(false)

	if hasUnapply {
		return fmt.Sprintf("%s.Unapply()", varexpr)
	}

	//fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})

	return seq.Map(names, func(v string) string { return fmt.Sprintf("%s.%s", varexpr, v) }).MakeString(",")
}

// struct{ A : x , B : y }
func (r *TypeClassSummonContext) structApplyExpr(ctx CurrentContext, named fp.Option[metafp.NamedTypeInfo], fields fp.Seq[metafp.StructField], args ...string) string {
	hasApply := option.Map(named, func(v metafp.NamedTypeInfo) bool { return v.Info.Method.Contains("Builder") }).OrElse(false)

	if hasApply {

		valuetp := ""
		if named.Get().Info.TypeParam.Size() > 0 {
			valuetp = "[" + iterator.Map(seq.Iterator(named.Get().Info.TypeParam), func(v metafp.TypeParam) string {
				return v.Name
			}).MakeString(",") + "]"
		}

		builderreceiver := fmt.Sprintf("%sBuilder%s", named.Get().PackagedName(r.w, ctx.working), valuetp)

		return fmt.Sprintf(`%s{}.Apply(%s).Build()`,
			builderreceiver, as.Seq(args).MakeString(","))
	}

	//fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})
	argslist := seq.Map(seq.Zip(names, args), func(v fp.Tuple2[string, string]) string {
		return fmt.Sprintf("%s: %s", v.I1, v.I2)
	}).MakeString(",")

	valuereceiver := option.Map(named, func(v metafp.NamedTypeInfo) string {
		valuetp := ""
		if v.Info.TypeParam.Size() > 0 {
			valuetp = "[" + iterator.Map(seq.Iterator(named.Get().Info.TypeParam), func(v metafp.TypeParam) string {
				return v.Name
			}).MakeString(",") + "]"
		}
		return fmt.Sprintf("%s%s", v.PackagedName(r.w, ctx.working), valuetp)
	}).OrElseGet(func() string {
		return "struct { " + seq.Map(fields, func(v metafp.StructField) string {
			if v.Embedded {
				return v.TypeName(r.w, ctx.working)
			}
			return fmt.Sprintf("%s %s",
				v.Name,
				v.TypeName(r.w, ctx.working),
			)
		}).MakeString("\n") + "}"
	})

	return fmt.Sprintf(`%s{%s}`, valuereceiver, argslist)
}

func namedOrRuntimeType(fpPkg *types.Package, working genfp.WorkingPackage, typePkg *types.Package, structName string, name string, vtype metafp.TypeInfo, genLabelled bool) metafp.TypeInfo {

	if genLabelled {
		ret := publicName(name)
		if ret == name {
			ret = fmt.Sprintf("PubNamed%sOf%s", ret, structName)
		} else {
			ret = fmt.Sprintf("Named%sOf%s", ret, structName)
		}

		obj := typePkg.Scope().Lookup(ret)
		ti := metafp.GetTypeInfo(obj.Type())
		if ti.TypeParam.Size() > 0 {
			ctx := types.NewContext()
			targs := seq.Map(ti.TypeParam, func(v metafp.TypeParam) types.Type {
				return types.NewTypeParam(v.TypeName, v.Constraint)
			})

			// ti : NamedOptOfCar[T comparable] fp.Tuple1[fp.Option[T]]
			// targs : fp.Option[T]

			it, err := types.Instantiate(ctx, ti.Type, targs, false)
			if err == nil {
				return metafp.GetTypeInfo(it)
			}
		} else {
			return ti
		}
	}

	rtobj := fpPkg.Scope().Lookup("RuntimeNamed")
	if tpe, ok := rtobj.(*types.TypeName); ok {
		ctx := types.NewContext()
		targs := []types.Type{vtype.Type}
		it, err := types.Instantiate(ctx, tpe.Type(), targs, false)
		if err == nil {
			return metafp.GetTypeInfo(it)
		}
	}

	panic("can't find named type for " + name)
	//  else {
	// 	fppk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

	// 	return fmt.Sprintf("%s.RuntimeNamed", fppk)

	// }

}

func namedOrRuntimeStringExpr(w genfp.ImportSet, working genfp.WorkingPackage, typePkg *types.Package, structName string, name string, labelledGen bool, valueType string) string {

	if labelledGen {
		ret := publicName(name)
		if ret == name {
			ret = fmt.Sprintf("PubNamed%sOf%s", ret, structName)
		} else {
			ret = fmt.Sprintf("Named%sOf%s", ret, structName)
		}

		if isSamePkg(working, genfp.FromTypesPackage(typePkg)) {
			return ret
		} else {
			return fmt.Sprintf("%s.%s", w.GetImportedName(genfp.FromTypesPackage(typePkg)), ret)
		}
	} else {
		fppk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

		return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, valueType)

	}

}

var implicitTypeInference = option.Of(runtime.Version()).Filter(func(v string) bool { return v >= "1.21.0" }).IsDefined()

func (r *TypeClassSummonContext) summonLabelledGenericRepr(ctx CurrentContext, tc metafp.TypeClass, sf structFunctions) fp.Option[GenericRepr] {

	type fieldName = fp.Tuple2[string, string]
	fields := sf.fields
	names := seq.Map(fields, func(v metafp.StructField) fieldName {
		return as.Tuple(v.Name, v.Tag)
	})

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfoExpr {
		return v.TypeInfoExpr(ctx.working)
	})

	result := r.lookupTypeClassFunc(ctx, tc, fmt.Sprintf("Labelled%d", typeArgs.Size()))

	return option.Map(result, func(tm metafp.TypeClassInstance) GenericRepr {
		return GenericRepr{
			Kind: fp.GenericKindStruct,
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr:   sf.asLabelled,
			FromReprExpr: sf.fromLabelled,
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassMemberLabelled(ctx, tc, tm, sf.pack, sf.name, seq.Map(names, xtr.Head), typeArgs, sf.namedGenerated)
			},
		}
	}).Or(func() fp.Option[GenericRepr] {
		return option.Map(r.lookupTypeClassFunc(ctx, tc, "HConsLabelled"), func(hcons metafp.TypeClassInstance) GenericRepr {
			return GenericRepr{
				Kind: fp.GenericKindStruct,
				// ReprType: func() string {
				// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
				// },
				ToReprExpr: func() string {

					if typeArgs.Size() == 0 {
						hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

						return fmt.Sprintf(`func (%s) %s.Nil {
							return %s.Empty()
						}`, sf.typeStr(ctx.working), hlistpk, hlistpk)
					} else if typeArgs.Size() < max.Product {

						arity := fp.Min(typeArgs.Size(), max.Product-1)
						//arity := typeArgs.Size()

						fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
						aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						if r.implicitTypeInference {
							return fmt.Sprintf(`%s.Compose(
								%s,
								%s.HList%dLabelled,
							)`, fppk,
								sf.asLabelled(),
								aspk, arity,
							)
						} else {
							tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
								return namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, f.I1.I1, sf.namedGenerated, f.I2.TypeName(r.w, ctx.working))
							}).Take(arity).MakeString(",")

							return fmt.Sprintf(`%s.Compose(
								%s,
								%s.HList%dLabelled[%s],
							)`, fppk,
								sf.asLabelled(),
								aspk, arity, tp,
							)
						}

					} else {
						hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						hlisttp := seq.Fold(namedTypeArgs.Reverse(), hlistpk+".Nil", func(b string, f fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
							name, a := f.Unapply()
							return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, name.I1, sf.namedGenerated, a.TypeName(r.w, ctx.working)), b)
						})

						varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
							return fmt.Sprintf("i%d", v)
						}).MakeString(",")

						hlistExpr := option.Map(option.Of(sf.namedGenerated).Filter(eq.GivenValue(true)), func(bool) string {
							return seq.Fold(seq.ZipWithIndex(namedTypeArgs).Reverse(), hlistpk+".Empty()", func(expr string, t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
								idx, t2 := t3.Unapply()
								name, tp := t2.Unapply()
								return fmt.Sprintf(`%s.Concat(%s{i%d}, 
									%s,
								)`, hlistpk, namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, name.I1, sf.namedGenerated, tp.TypeName(r.w, ctx.working)), idx, expr)
							})
						}).OrElseGet(func() string {
							aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

							return seq.Fold(seq.ZipWithIndex(namedTypeArgs).Reverse(), hlistpk+".Empty()", func(expr string, t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
								idx, t2 := t3.Unapply()
								name, _ := t2.Unapply()
								return fmt.Sprintf(`%s.Concat(%s.NamedWithTag("%s", i%d, %s), 
									%s,
								)`, hlistpk, aspk, name.I1, idx, "`"+name.I2+"`", expr)
							})
						})
						//hlistExpr :=

						unapplyexpr := sf.unapply("v")
						return fmt.Sprintf(`func(v %s) %s {
							%s := %s
							return %s
						}`, sf.typeStr(ctx.working), hlisttp,
							varlist, unapplyexpr,
							hlistExpr)
					}

				},
				FromReprExpr: func() string {
					if typeArgs.Size() == 0 {
						hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
						valuereceiver := sf.typeStr(ctx.working)
						return fmt.Sprintf(`func (%s.Nil) %s{
							return %s{}
						}`, hlistpk, valuereceiver, valuereceiver)
					} else if typeArgs.Size() < max.Product {
						arity := fp.Min(typeArgs.Size(), max.Product-1)
						//arity := typeArgs.Size()

						fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
						productpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/product", "product"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						hlistToTuple := func() string {
							if r.implicitTypeInference {
								return fmt.Sprintf(`%s.LabelledFromHList%d`,
									productpk,
									arity,
								)
							} else {
								tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
									return namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, f.I1.I1, sf.namedGenerated, f.I2.TypeName(r.w, ctx.working))
								}).Take(arity).MakeString(",")

								return fmt.Sprintf(`%s.LabelledFromHList%d[%s]`,
									productpk,
									arity, tp,
								)
							}
						}()

						tupleToStruct := sf.fromLabelled()
						return fmt.Sprintf(`
						%s.Compose(
							%s, 
							%s ,
						)`, fppk, hlistToTuple, tupleToStruct)
					} else {
						hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						hlisttp := seq.Fold(namedTypeArgs.Reverse(), hlistpk+".Nil", func(b string, t2 fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
							name, a := t2.Unapply()
							return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, namedOrRuntimeStringExpr(r.w, ctx.working, sf.pack, sf.name, name.I1, sf.namedGenerated, a.TypeName(r.w, ctx.working)), b)
						})

						expr := seq.Map(iterator.Range(0, typeArgs.Size()).ToSeq(), func(idx int) string {
							if idx == typeArgs.Size()-1 {
								return fmt.Sprintf(`i%d := hl%d.Head()`, idx, idx)
							}
							return fmt.Sprintf(`i%d , hl%d := %s.Unapply(hl%d)`, idx, idx+1, hlistpk, idx)
						}).MakeString("\n")

						arglist := seq.Map(iterator.Range(0, typeArgs.Size()).ToSeq(), func(idx int) string {
							return fmt.Sprintf("i%d.Value()", idx)
						})
						return fmt.Sprintf(`func(hl0 %s) %s {
								%s
								return %s
							}`, hlisttp, sf.typeStr(ctx.working),
							expr,
							sf.apply(arglist))
					}
				},
				ReprExpr: func() SummonExpr {
					//arity := fp.Min(typeArgs.Size(), max.Product-1)
					arity := typeArgs.Size()

					hnil := r.lookupHNilMust(ctx, tc)
					namedTypeArgs := seq.Zip(names, typeArgs)
					hlist := seq.Fold(namedTypeArgs.Take(arity).Reverse(), newSummonExpr(hnil.PackagedName(r.w, ctx.working)), func(tail SummonExpr, ti fp.Tuple2[fieldName, metafp.TypeInfoExpr]) SummonExpr {
						instance := r.summonFpNamed(ctx, tc, sf.pack, sf.name, ti.I1.I1, ti.I2, sf.namedGenerated)
						return newSummonExpr(fmt.Sprintf(`%s(
							%s,
						%s,
						)`, hcons.PackagedName(r.w, ctx.working), instance, tail), instance.paramInstance, tail.paramInstance)
					})

					return hlist
				},
			}
		})
	})
}

func (r *TypeClassSummonContext) namedStructFuncs(ctx CurrentContext, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) structFunctions {
	hasUnapply := named.Info.Method.Contains("Unapply")

	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	typeStr := func(pk genfp.WorkingPackage) string {
		return fmt.Sprintf("%s%s", named.PackagedName(r.w, ctx.working), valuetp)
	}

	builderTypeStr := func(pk genfp.WorkingPackage) string {
		return fmt.Sprintf("%sBuilder%s", named.PackagedName(r.w, ctx.working), valuetp)
	}

	if !hasUnapply && !isSamePkg(ctx.working, genfp.FromTypesPackage(named.Package)) {
		fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	}

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfoExpr {
		return v.TypeInfoExpr(ctx.working)
	})

	type fieldName = fp.Tuple2[string, string]
	names := seq.Map(fields, func(v metafp.StructField) fieldName {
		return as.Tuple(v.Name, v.Tag)
	})

	tupleFuncExpr := func() string {
		return fmt.Sprintf("%s.AsTuple", typeStr(ctx.working))
	}
	applyFuncExpr := func() string {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		builderreceiver := builderTypeStr(ctx.working)
		return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromTuple)(%s{}),
					%s.Build,
					)`,
			fppk,
			aspk, builderreceiver, builderreceiver, builderreceiver,
		)
	}

	asLabelledFuncExpr := func() string {
		return fmt.Sprintf("%s.AsLabelled", typeStr(ctx.working))
	}

	fromLabelledFuncExpr := func() string {
		builderreceiver := builderTypeStr(ctx.working)
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromLabelled)(%s{}),
					%s.Build,
					)`,
			fppk,
			aspk, builderreceiver, builderreceiver, builderreceiver,
		)
	}

	unapplyFunc := func(structIns string) string {
		return fmt.Sprintf("%s.Unapply()", structIns)
	}

	applyFunc := func(fieldValues []string) string {
		return fmt.Sprintf(`%s{}.Apply(%s).Build()`,
			builderTypeStr(ctx.working), as.Seq(fieldValues).MakeString(","))
	}
	if !hasUnapply {

		tupleFuncExpr = func() string {
			p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
				return f.TypeName(r.w, ctx.working)
			}).MakeString(",")

			fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			return fmt.Sprintf(`func( v %s) %s.Tuple%d[%s] {
			return %s.Tuple%d(%s)
		}`, typeStr(ctx.working), fppk, fields.Size(), p,
				aspk, fields.Size(), seq.Map(names, func(v fieldName) string { return "v." + v.I1 }).MakeString(","),
			)
		}

		applyFuncExpr = func() string {
			p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
				return f.TypeName(r.w, ctx.working)
			}).MakeString(",")

			fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			//aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			assign := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
				return fmt.Sprintf("%s : t.I%d", v.I2.I1, v.I1+1)
			}).MakeString(",\n")
			valuereceiver := typeStr(ctx.working)
			return fmt.Sprintf(`func(t %s.Tuple%d[%s]) %s {
					return %s{
						%s,
					}
				}`, fppk, fields.Size(), p, valuereceiver,
				valuereceiver,
				assign,
			)
		}

		unapplyFunc = func(structIns string) string {
			return seq.Map(names, func(v fieldName) string { return fmt.Sprintf("%s.%s", structIns, v.I1) }).MakeString(",")

		}

		applyFunc = func(fieldValues []string) string {
			argslist := seq.Map(seq.Zip(names, fieldValues), func(v fp.Tuple2[fieldName, string]) string {
				return fmt.Sprintf("%s: %s", v.I1.I1, v.I2)
			}).MakeString(",")

			return fmt.Sprintf(`%s{%s}`, typeStr(ctx.working), argslist)

		}

	}

	hasAsLabelled := named.Info.Method.Contains("AsLabelled")
	if !hasAsLabelled {
		asLabelledFuncExpr = func() string {
			fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			namedTypeArgs := seq.Zip(names, typeArgs)

			labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
				return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
			}).MakeString(",")

			varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
				return fmt.Sprintf("i%d", v)
			}).MakeString(",")

			hlistExpr := seq.Map(seq.ZipWithIndex(namedTypeArgs), func(t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
				idx, t2 := t3.Unapply()
				name, _ := t2.Unapply()
				return fmt.Sprintf(`%s.NamedWithTag("%s", i%d , %s)`, aspk, name.I1, idx, "`"+name.I2+"`")

			}).MakeString(",")

			return fmt.Sprintf(`func(v %s) %s.Labelled%d[%s] {
							%s := %s
							return %s.Labelled%d(%s)
						}`, typeStr(ctx.working), fppk, fields.Size(), labelledtp,
				varlist, unapplyFunc("v"),
				aspk, fields.Size(), hlistExpr)
		}

		fromLabelledFuncExpr = func() string {
			fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			namedTypeArgs := seq.Zip(names, typeArgs)

			labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
				return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
			}).MakeString(",")

			args := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
				return fmt.Sprintf("t.I%d.Value()", v.I1+1)
			})

			return fmt.Sprintf(`func( t %s.Labelled%d[%s] ) %s {
				return %s
			}`, fppk, fields.Size(), labelledtp, typeStr(ctx.working),
				r.structApplyExpr(ctx, option.Some(named), fields, args...),
			)
		}
	}

	sf := structFunctions{
		pack:           named.Package,
		name:           named.Name,
		tpe:            named.Info,
		fields:         fields,
		typeArgs:       typeArgs,
		namedGenerated: hasAsLabelled,
		asTuple:        tupleFuncExpr,
		fromTuple:      applyFuncExpr,
		asLabelled:     asLabelledFuncExpr,
		fromLabelled:   fromLabelledFuncExpr,
		typeStr:        typeStr,
		unapply:        unapplyFunc,
		apply:          applyFunc,
	}
	return sf
}

func (r *TypeClassSummonContext) untypedStructFuncs(ctx CurrentContext, tpe metafp.TypeInfo, fields fp.Seq[metafp.StructField]) structFunctions {

	typeStr := func(pk genfp.WorkingPackage) string {
		valuereceiver := "struct { " + seq.Map(fields, func(v metafp.StructField) string {
			if v.Embedded {
				return v.TypeName(r.w, ctx.working)

			}
			return fmt.Sprintf("%s %s",
				v.Name,
				v.TypeName(r.w, ctx.working),
			)
		}).MakeString("\n") + "}"

		return valuereceiver
	}

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfoExpr {
		return v.TypeInfoExpr(ctx.working)
	})

	type fieldName = fp.Tuple2[string, string]
	names := seq.Map(fields, func(v metafp.StructField) fieldName {
		return as.Tuple2(v.Name, v.Tag)
	})

	tupleFuncExpr := func() string {
		p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
			return f.TypeName(r.w, ctx.working)
		}).MakeString(",")

		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		return fmt.Sprintf(`func( v %s) %s.Tuple%d[%s] {
			return %s.Tuple%d(%s)
		}`, typeStr(ctx.working), fppk, fields.Size(), p,
			aspk, fields.Size(), seq.Map(names, func(v fieldName) string { return "v." + v.I1 }).MakeString(","),
		)
	}

	applyFuncExpr := func() string {
		p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
			return f.TypeName(r.w, ctx.working)
		}).MakeString(",")

		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		//aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		assign := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
			return fmt.Sprintf("%s : t.I%d", v.I2.I1, v.I1+1)
		}).MakeString(",\n")
		return fmt.Sprintf(`func(t %s.Tuple%d[%s]) %s {
					return %s{
						%s,
					}
				}`, fppk, fields.Size(), p, typeStr(ctx.working),
			typeStr(ctx.working),
			assign,
		)
	}

	asLabelledFuncExpr := func() string {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		namedTypeArgs := seq.Zip(names, typeArgs)

		labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
			return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
		}).MakeString(",")

		varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
			return fmt.Sprintf("i%d", v)
		}).MakeString(",")

		hlistExpr := seq.Map(seq.ZipWithIndex(namedTypeArgs), func(t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
			idx, t2 := t3.Unapply()
			name, _ := t2.Unapply()
			return fmt.Sprintf(`%s.NamedWithTag("%s",  i%d, %s)`, aspk, name.I1, idx, "`"+name.I2+"`")

		}).MakeString(",")

		unapplyexpr := r.structUnapplyExpr(ctx, option.None[metafp.NamedTypeInfo](), fields, "v")
		return fmt.Sprintf(`func(v %s) %s.Labelled%d[%s] {
							%s := %s
							return %s.Labelled%d(%s)
						}`, typeStr(ctx.working), fppk, fields.Size(), labelledtp,
			varlist, unapplyexpr,
			aspk, fields.Size(), hlistExpr)
	}

	fromLabelledFuncExpr := func() string {
		fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		namedTypeArgs := seq.Zip(names, typeArgs)

		labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
			return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
		}).MakeString(",")

		args := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
			return fmt.Sprintf("t.I%d.Value()", v.I1+1)
		})

		return fmt.Sprintf(`func( t %s.Labelled%d[%s] ) %s {
				return %s
			}`, fppk, fields.Size(), labelledtp, typeStr(ctx.working),
			r.structApplyExpr(ctx, option.None[metafp.NamedTypeInfo](), fields, args...),
		)
	}

	sf := structFunctions{
		pack:         ctx.working.Package(),
		tpe:          tpe,
		fields:       fields,
		typeArgs:     typeArgs,
		asTuple:      tupleFuncExpr,
		fromTuple:    applyFuncExpr,
		asLabelled:   asLabelledFuncExpr,
		fromLabelled: fromLabelledFuncExpr,
		typeStr:      typeStr,
		unapply: func(structIns string) string {
			return seq.Map(names, func(v fieldName) string { return fmt.Sprintf("%s.%s", structIns, v.I1) }).MakeString(",")
		},
		apply: func(fieldValues []string) string {
			argslist := seq.Map(seq.Zip(names, fieldValues), func(v fp.Tuple2[fieldName, string]) string {
				return fmt.Sprintf("%s: %s", v.I1.I1, v.I2)
			}).MakeString(",")

			return fmt.Sprintf(`%s{%s}`, typeStr(ctx.working), argslist)
		},
	}
	return sf
}

type structFunctions struct {
	pack           *types.Package
	name           string
	tpe            metafp.TypeInfo
	fields         fp.Seq[metafp.StructField]
	typeArgs       fp.Seq[metafp.TypeInfoExpr]
	namedGenerated bool

	typeStr func(pk genfp.WorkingPackage) string

	// func( v struct{} ) fp.Tuple2[A,B] {}
	asTuple func() string

	asLabelled func() string

	fromLabelled func() string

	// func(v fp.Tuple2[A,B]) struct{}
	fromTuple func() string

	// v.A , v.B
	unapply func(structIns string) string

	// struct{ A : x , B : y }
	apply func(fieldValues []string) string
}

//	Tuple%d  instance 가 있는 경우
//	   --> Value generaed 된 type 인 경우
//	   --> 그냥 struct 인 경우
//	   --> untyped struct 인 경우
//
// HCons instance 가 있는 경우
//
//	--> Value generaed 된 type 인 경우
//	--> 그냥 struct 인 경우
//	--> untyped struct 인 경우
func (r *TypeClassSummonContext) summonStructGenericRepr(ctx CurrentContext, tc metafp.TypeClass, sf structFunctions) GenericRepr {
	fields := sf.fields
	result := r.lookupTupleTypeClassFunc(ctx, tc, fmt.Sprintf("Tuple%d", fields.Size()), sf.typeArgs)

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfoExpr {
		return v.TypeInfoExpr(ctx.working)
	})

	if result.IsDefined() {

		tci := result.Get()
		tci.RequiredInstance = seq.Map(tci.RequiredInstance, func(v metafp.RequiredInstance) metafp.RequiredInstance {
			v.FieldOf = option.Some(sf.tpe)
			return v
		})
		// tp := iterator.Map(typeArgs.Iterator(), func(v metafp.TypeInfo) string {
		// 	return r.w.TypeName(ctx.working, v.Type)
		// }).MakeString(",")
		return GenericRepr{
			Kind: fp.GenericKindStruct,
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr:   sf.asTuple,
			FromReprExpr: sf.fromTuple,
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassInstance(ctx, tci)
				//return r.exprTypeClassMember(ctx, tc, result.Get(), typeArgs, option.Some(sf.tpe))
			},
		}
	}

	tupleGeneric := r.summonTupleGenericRepr(ctx, tc, typeArgs, option.Some(sf.tpe), false)

	return GenericRepr{
		Kind: fp.GenericKindStruct,

		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {

			if typeArgs.Size() >= max.Product {
				hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

				hlisttp := seq.Fold(typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfoExpr) string {
					return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, a.TypeName(r.w, ctx.working), b)
				})

				varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
					return fmt.Sprintf("i%d", v)
				}).MakeString(",")

				hlistExpr := seq.Fold(as.Seq(iterator.Range(0, typeArgs.Size()).ToSeq()).Reverse(), hlistpk+".Empty()", func(expr string, v int) string {
					return fmt.Sprintf(`%s.Concat(i%d, 
						%s,
					)`, hlistpk, v, expr)
				})
				return fmt.Sprintf(`func(v %s) %s {
					%s := %s
					return %s
				}`, sf.typeStr(ctx.working), hlisttp,
					varlist, sf.unapply("v"),
					hlistExpr)
			} else if typeArgs.Size() > 0 {
				fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

				return fmt.Sprintf(`%s.Compose(
				%s,
				%s, 
			)`, fppk,
					sf.asTuple(),
					tupleGeneric.ToReprExpr(),
				)
			} else {
				hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
				return fmt.Sprintf(`func(%s) %s.Nil {
					return %s.Empty()
				}`, sf.typeStr(ctx.working), hlistpk, hlistpk)
			}

		},
		FromReprExpr: func() string {
			if typeArgs.Size() >= max.Product {
				hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

				hlisttp := seq.Fold(typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfoExpr) string {
					return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, a.TypeName(r.w, ctx.working), b)
				})

				expr := seq.Map(iterator.Range(0, typeArgs.Size()).ToSeq(), func(idx int) string {
					if idx == typeArgs.Size()-1 {
						return fmt.Sprintf(`i%d := hl%d.Head()`, idx, idx)
					}
					return fmt.Sprintf(`i%d , hl%d := %s.Unapply(hl%d)`, idx, idx+1, hlistpk, idx)
				}).MakeString("\n")

				arglist := seq.Map(iterator.Range(0, typeArgs.Size()).ToSeq(), func(idx int) string {
					return fmt.Sprintf("i%d", idx)
				})
				return fmt.Sprintf(`func(hl0 %s) %s {
					%s
					return %s
				}`, hlisttp, sf.typeStr(ctx.working),
					expr,
					sf.apply(arglist))
			} else if typeArgs.Size() > 0 {

				fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
				//aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

				tupleToStruct := sf.fromTuple()
				return fmt.Sprintf(`
				%s.Compose(
					%s, 
					%s ,
				)`, fppk, tupleGeneric.FromReprExpr(), tupleToStruct)
			} else {
				hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))

				return fmt.Sprintf(`func(%s.Nil) %s {
					return %s{}
				}`, hlistpk, sf.typeStr(ctx.working), sf.typeStr(ctx.working))
			}
		},
		ReprExpr: func() SummonExpr {
			return option.Map(r.lookupTypeClassFunc(ctx, tc, "StructHCons"), func(hcons metafp.TypeClassInstance) SummonExpr {
				arity := typeArgs.Size()
				hnil := r.lookupTypeClassFunc(ctx, tc, "StructHNil").OrElseGet(as.Supplier2(r.lookupHNilMust, ctx, tc))
				hlist := seq.Fold(typeArgs.Take(arity).Reverse(), newSummonExpr(hnil.PackagedName(r.w, ctx.working)), func(tail SummonExpr, ti metafp.TypeInfoExpr) SummonExpr {
					instance := r.summonRequired(ctx, metafp.RequiredInstance{
						TypeClass: ctx.tc.TypeClass,
						Type:      ti.Type,
						FieldOf:   option.Some(sf.tpe),
					})
					return newSummonExpr(fmt.Sprintf(`%s(
					%s,
					%s,
				)`, hcons.PackagedName(r.w, ctx.working), instance, tail), instance.paramInstance, tail.paramInstance)
				})
				return hlist

			}).OrElseGet(tupleGeneric.ReprExpr)
		},
	}

}

// func (r *TypeClassSummonContext) summonNamedGenericRepr(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) GenericRepr {
// 	sf := r.namedStructFuncs(ctx, named, fields)
// 	return r.summonStructGenericRepr(ctx, tc, sf)
// }

func (r *TypeClassSummonContext) summonTupleGenericRepr(ctx CurrentContext, tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfoExpr], fieldOf fp.Option[metafp.TypeInfo], explicit bool) GenericRepr {
	return GenericRepr{
		Kind: fp.GenericKindTuple,
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {
			aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			arity := fp.Min(typeArgs.Size(), max.Product-1)
			//arity := typeArgs.Size()

			if r.implicitTypeInference {
				return fmt.Sprintf(`%s.HList%d`,
					aspk, arity,
				)
			}
			tp := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
				return f.TypeName(r.w, ctx.working)
			}).Take(arity).MakeString(",")

			return fmt.Sprintf(`%s.HList%d[%s]`,
				aspk, arity, tp,
			)

		},
		FromReprExpr: func() string {
			productpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/product", "product"))

			arity := fp.Min(typeArgs.Size(), max.Product-1)
			//arity := typeArgs.Size()

			hlistToTuple := func() string {
				if r.implicitTypeInference && !explicit {
					return fmt.Sprintf(`%s.TupleFromHList%d`,
						productpk, arity,
					)
				} else {

					tp := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
						return f.TypeName(r.w, ctx.working)
					}).Take(arity).MakeString(",")

					return fmt.Sprintf(`%s.TupleFromHList%d[%s]`,
						productpk, arity, tp,
					)
				}
			}()

			// hlistToTuple := fmt.Sprintf(`%s.Func2(
			// 		%s.Case%d[%s,%s.Nil,fp.Tuple%d[%s]],
			// 	).ApplyLast(
			// 		%s.Tuple%d[%s] ,
			// 	)`,
			// 	aspk, hlistpk, typeArgs.Size(), tp, hlistpk, typeArgs.Size(), tp, aspk, typeArgs.Size(), tp,
			// )

			return hlistToTuple
		},
		ReprExpr: func() SummonExpr {
			//arity := fp.Min(typeArgs.Size(), max.Product-1)
			arity := typeArgs.Size()

			hcons := r.lookupTypeClassFunc(ctx, tc, "TupleHCons").OrElseGet(as.Supplier2(r.lookupHConsMust, ctx, tc))

			hnil := r.lookupTypeClassFunc(ctx, tc, "TupleHNil").OrElseGet(as.Supplier2(r.lookupHNilMust, ctx, tc))

			hlist := seq.Fold(typeArgs.Take(arity).Reverse(), newSummonExpr(hnil.PackagedName(r.w, ctx.working)), func(tail SummonExpr, ti metafp.TypeInfoExpr) SummonExpr {
				instance := r.summonRequired(ctx, metafp.RequiredInstance{
					TypeClass: ctx.tc.TypeClass,
					Type:      ti.Type,
					FieldOf:   fieldOf,
				})
				return newSummonExpr(fmt.Sprintf(`%s(
					%s,
					%s,
				)`, hcons.PackagedName(r.w, ctx.working), instance, tail), instance.paramInstance, tail.paramInstance)
			})
			return hlist
		},
	}
}

func (r *TypeClassSummonContext) summonTuple(ctx CurrentContext, tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfoExpr]) SummonExpr {

	result := r.lookupTypeClassFunc(ctx, tc, fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {
		return r.exprTypeClassMember(ctx, tc, result.Get(), typeArgs, fp.Option[metafp.TypeInfo]{})
	}

	tupleGeneric := r.summonTupleGenericRepr(ctx, tc, typeArgs, fp.Option[metafp.TypeInfo]{}, true)
	return r.summonVariant(ctx, tc, fmt.Sprintf("fp.Tuple%d", typeArgs.Size()), tupleGeneric)

}

func (r *TypeClassSummonContext) summonFpNamed(ctx CurrentContext, tc metafp.TypeClass, typePkg *types.Package, structName string, name string, t metafp.TypeInfoExpr, genLabelled bool) SummonExpr {

	rtt := namedOrRuntimeType(r.fpPkg.Get(), ctx.working, typePkg, structName, name, t.Type, genLabelled)
	//fmt.Printf("tc = %s, type = %s, rtt = %s, %T\n", tc, t.Type, rtt, rtt)
	named := r.namedLookup(ctx, metafp.RequiredInstance{
		TypeClass: tc,
		Type:      rtt,
	}, "Named")
	// named := r.lookupTypeClassFunc(ctx, tc, "Named")
	if named.IsDefined() {
		//fmt.Printf("find named\n")
		return r.exprTypeClassInstance(ctx, named.Get())
	}
	instance := r.lookupTypeClassFuncMust(ctx, tc, "Named")

	expr := r.summonRequired(ctx, metafp.RequiredInstance{
		TypeClass: tc,
		Type:      t.Type,
	})

	if instance.RequiredInstance.Size() == 0 {
		return newSummonExpr(
			fmt.Sprintf("%s[%s]()",
				instance.PackagedName(r.w, ctx.working),
				namedOrRuntimeStringExpr(r.w, ctx.working, typePkg, structName, name, genLabelled, t.TypeName(r.w, ctx.working)),
			), expr.paramInstance)
	}
	return newSummonExpr(
		fmt.Sprintf("%s[%s](%s)",
			instance.PackagedName(r.w, ctx.working),
			namedOrRuntimeStringExpr(r.w, ctx.working, typePkg, structName, name, genLabelled, t.TypeName(r.w, ctx.working)),
			expr.expr,
		), expr.paramInstance)

	// pk := r.w.GetImportedName(ctx.working)
	// return fmt.Sprintf("%s.Named(%s)", pk, r.summon(t))
}

func (r *TypeClassSummonContext) SummonExpression(tc metafp.TypeClassDerive) SummonExpr {

	ctx := CurrentContext{
		workingScope: r.tcCache.GetLocal(tc.Package.Package(), tc.TypeClass),
		primScope:    r.tcCache.Get(tc.PrimitiveInstancePkg, tc.TypeClass),
		tc:           tc,
		working:      tc.Package,
		recursiveGen: option.FlatMap(tc.Tags.Get("@fp.Derive"),
			fp.Compose2(metafp.Annotation.Params, as.Func2(fp.Map[string, string].Get).ApplyLast("recursive"))).Exists(eq.GivenValue("true")),
	}

	return r.summonRequired(ctx, metafp.RequiredInstance{
		TypeClass: tc.TypeClass,
		Type:      tc.DeriveFor.Info,
	})

}
func (r *TypeClassSummonContext) summonRequired(ctx CurrentContext, req metafp.RequiredInstance) SummonExpr {

	t := req.Type

	// if req.TypeClass.IsLazy() {
	// 	expr := r.summon(req.Type.TypeArgs.Head().Get())
	// }

	if t.IsTuple() {
		return r.summonTuple(ctx, req.TypeClass, seq.Map(t.TypeArgs, func(v metafp.TypeInfo) metafp.TypeInfoExpr {
			return metafp.TypeInfoExpr{
				Type: v,
			}
		}))
	}

	result := r.lookupTypeClassInstance(ctx, req)

	if result.target.IsRight() {
		return r.exprLookupTarget(ctx, result)
	}

	if ctx.recursiveGen && req.TypeClass.Id() == ctx.tc.TypeClass.Id() {
		must := result.target.Left()
		named := must.instanceOf.AsNamed()
		if named.IsDefined() {
			deriveFor := named.Get().Info

			vt := metafp.LookupStruct(deriveFor.Pkg, deriveFor.Name().Get())

			tc := metafp.TypeClassDerive{
				Package:              ctx.tc.Package,
				PrimitiveInstancePkg: ctx.tc.PrimitiveInstancePkg,
				TypeClass:            ctx.tc.TypeClass,
				TypeClassType:        ctx.tc.TypeClassType,
				DeriveFor:            named.Get(),
				StructInfo:           vt,
				Tags:                 ctx.tc.Tags,
			}

			if !r.tcCache.IsWillGenerated(tc) {
				r.recursiveGen = append(r.recursiveGen, tc)
				r.tcCache.WillGenerated(tc)
			}
			return r.exprLookupTarget(ctx, result)
		}
	}
	return r.exprLookupTarget(ctx, result)

}

func (r *TypeClassSummonContext) summonStruct(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) SummonExpr {

	sf := r.namedStructFuncs(ctx, named, fields)
	labelledExpr := r.summonLabelledGenericRepr(ctx, tc, sf)
	summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
		return r.summonStructGenericRepr(ctx, tc, sf)
	})

	return r.summonVariant(ctx, tc, named.GenericName(), summonExpr)
}

func (r *TypeClassSummonContext) summonUntypedStruct(ctx CurrentContext, tc metafp.TypeClass, tpe metafp.TypeInfo, fields fp.Seq[metafp.StructField]) SummonExpr {

	sf := r.untypedStructFuncs(ctx, tpe, fields)

	labelledExpr := r.summonLabelledGenericRepr(ctx, tc, sf)
	summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
		return r.summonStructGenericRepr(ctx, tc, sf)
	})

	return r.summonVariant(ctx, tc, "struct", summonExpr)
}

func (r *TypeClassSummonContext) summonVariant(ctx CurrentContext, tc metafp.TypeClass, genericName string, genericRepr GenericRepr) SummonExpr {
	mapExpr := option.Map(r.lookupTypeClassFunc(ctx, tc, "Generic"), func(generic metafp.TypeClassInstance) SummonExpr {

		aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))
		repr := genericRepr.ReprExpr()
		return newSummonExpr(fmt.Sprintf(`%s(
					%s.Generic(
							"%s",
							"%s",
							%s,
							%s,
						), 
						%s, 
					)`, generic.PackagedName(r.w, ctx.working), aspk,
			genericName,
			genericRepr.Kind,
			genericRepr.ToReprExpr(),
			genericRepr.FromReprExpr(),
			repr), repr.paramInstance)

	}).Or(func() fp.Option[SummonExpr] {
		return option.Map(r.lookupTypeClassFunc(ctx, tc, "IMap"), func(imapfunc metafp.TypeClassInstance) SummonExpr {
			repr := genericRepr.ReprExpr()

			return newSummonExpr(fmt.Sprintf(`%s( 
						%s, 
						%s , 
						%s,
						)`,
				imapfunc.PackagedName(r.w, ctx.working), repr, genericRepr.FromReprExpr(), genericRepr.ToReprExpr()), repr.paramInstance)
		})
	}).Or(func() fp.Option[SummonExpr] {
		functor := r.lookupTypeClassFunc(ctx, tc, "Map")
		return option.Map(functor, func(v metafp.TypeClassInstance) SummonExpr {
			repr := genericRepr.ReprExpr()

			return newSummonExpr(fmt.Sprintf(`%s( 
						%s, 
						%s,
						)`,
				v.PackagedName(r.w, ctx.working), repr, genericRepr.FromReprExpr(),
			), repr.paramInstance)
		})

	}).OrElseGet(func() SummonExpr {
		contrmap := r.lookupTypeClassFuncMust(ctx, tc, "ContraMap")
		repr := genericRepr.ReprExpr()

		return newSummonExpr(fmt.Sprintf(`%s( 
					%s , 
					%s,
					)`,
			contrmap.PackagedName(r.w, ctx.working), repr, genericRepr.ToReprExpr(),
		), repr.paramInstance)
	})
	return mapExpr

}

func (r *TypeClassSummonContext) summonNamed(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo) SummonExpr {

	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	nameWithTp := named.PackagedName(r.w, ctx.working) + valuetp

	summonExpr := GenericRepr{
		Kind: fp.GenericKindNewType,
		ReprExpr: func() SummonExpr {
			return r.summonRequired(ctx, metafp.RequiredInstance{
				TypeClass: tc,
				Type:      named.Underlying,
			})
		},
		ToReprExpr: func() string {
			return fmt.Sprintf(`func(v %s) %s {
					return %s(v)
				}`, nameWithTp, r.w.TypeName(ctx.working, named.Underlying.Type), r.w.TypeName(ctx.working, named.Underlying.Type))
		},
		FromReprExpr: func() string {
			return fmt.Sprintf(`func(v %s) %s {
					return %s(v)
				}`, r.w.TypeName(ctx.working, named.Underlying.Type), nameWithTp, nameWithTp)
		},
	}

	return r.summonVariant(ctx, tc, named.GenericName(), summonExpr)
}

func (r *TypeClassSummonContext) _summonVar(tc metafp.TypeClassDerive) SummonExpr {
	workingPackage := tc.Package

	ctx := CurrentContext{
		workingScope: r.tcCache.GetLocal(tc.Package.Package(), tc.TypeClass),
		primScope:    r.tcCache.Get(tc.PrimitiveInstancePkg, tc.TypeClass),
		tc:           tc,
		working:      workingPackage,
		recursiveGen: option.FlatMap(tc.Tags.Get("@fp.Derive"),
			fp.Compose2(metafp.Annotation.Params, as.Func2(fp.Map[string, string].Get).ApplyLast("recursive"))).Exists(eq.GivenValue("true")),
	}

	valuetpdec := ""
	valuetp := ""
	if tc.DeriveFor.Info.TypeParam.Size() > 0 {
		valuetpdec = "[" + iterator.Map(seq.Iterator(tc.DeriveFor.Info.TypeParam), func(v metafp.TypeParam) string {
			tn := r.w.TypeName(workingPackage, v.Constraint)
			return fmt.Sprintf("%s %s", v.Name, tn)
		}).MakeString(",") + "]"

		valuetp = "[" + iterator.Map(seq.Iterator(tc.DeriveFor.Info.TypeParam), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	mapExpr := option.Map(tc.StructInfo, func(s metafp.TaggedStruct) SummonExpr {
		fields := s.Fields
		//privateFields := fields.FilterNot(metafp.StructField.Public)
		allFields := fields.FilterNot(func(v metafp.StructField) bool {
			return strings.HasPrefix(v.Name, "_")
		})

		return r.summonStruct(ctx, tc.TypeClass, tc.DeriveFor, allFields)
	}).OrElseGet(func() SummonExpr {
		return r.summonNamed(ctx, tc.TypeClass, tc.DeriveFor)
	})

	if tc.DeriveFor.Info.TypeParam.Size() > 0 {

		tcname := tc.TypeClass.PackagedName(r.w, workingPackage)
		// fargs := seq.Map(v.DeriveFor.Info.TypeParam, func(p metafp.TypeParam) string {
		// 	return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
		// }).MakeString(",")

		fargs := seq.Map(mapExpr.paramInstance, as.Func3(ParamInstance.Expr).ApplyLast2(r.w, ctx.tc.Package)).MakeString(",")

		return newSummonExpr(fmt.Sprintf(`
						func %s%s( %s ) %s[%s%s] {
							return %s
						}
					`, tc.GeneratedInstanceName(), valuetpdec, fargs, tcname, tc.DeriveFor.PackagedName(r.w, workingPackage), valuetp,
			mapExpr), mapExpr.paramInstance)

	} else {
		tcname := tc.TypeClass.PackagedName(r.w, workingPackage)

		return newSummonExpr(fmt.Sprintf(`
						func %s() %s[%s] {
							return %s
						}
					`, tc.GeneratedInstanceName(), tcname, tc.DeriveFor.PackagedName(r.w, workingPackage),
			mapExpr), mapExpr.paramInstance)
	}
}

func (r *TypeClassSummonContext) summonVar(tc metafp.TypeClassDerive) fp.Option[SummonExpr] {

	retOpt := r.summoned.Get(tc.GeneratedInstanceName())
	if retOpt.IsDefined() {
		return option.Some(retOpt.Get().Expr)
	}

	if r.loopCheck.Contains(tc.GeneratedInstanceName()) {
		// fmt.Printf("cycle detected\n")
		return option.None[SummonExpr]()
	}

	r.loopCheck = r.loopCheck.Incl(tc.GeneratedInstanceName())

	ret := r._summonVar(tc)
	r.summoned = r.summoned.Updated(tc.GeneratedInstanceName(), TypeClassInstanceGenerated{
		Derive: tc,
		Expr:   ret,
	})
	return option.Some(ret)
}

// func (r *TypeClassSummonContext) summonRequired(working *types.Package, tc metafp.RequiredInstance) fp.Option[SummonExpr] {
// 	ctx := CurrentContext{
// 		workingScope: r.tcCache.GetLocal(working, tc.TypeClass),
// 		tc:           tc.TypeClass,
// 	}
// }

func FindFpPackage(pk *types.Package) fp.Option[*types.Package] {
	ret := as.Seq(pk.Imports()).Find(func(v *types.Package) bool {
		return v.Path() == "github.com/csgura/fp"
	})
	if ret.IsDefined() {
		return ret
	}

	for _, p := range pk.Imports() {
		ret := FindFpPackage(p)
		if ret.IsDefined() {
			return ret
		}
	}

	return option.None[*types.Package]()
}

func NewTypeClassSummonContext(pkgs []*packages.Package, importSet genfp.ImportSet) *TypeClassSummonContext {

	fpPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindFpPackage(v.Types)
	}).NextOption()

	derives := metafp.FindTypeClassDerive(pkgs)
	tccache := metafp.TypeClassInstanceCache{}

	metafp.FindTypeClassImport(pkgs).Foreach(func(v metafp.TypeClassDirective) {
		//fmt.Printf("Import %s from %s into %s\n", v.TypeClass.Name, v.PrimitiveInstancePkg.Path(), v.Package.Path())
		tccache.Load(v.PrimitiveInstancePkg, v.TypeClass)
	})

	seq.Iterator(derives).Foreach(func(v metafp.TypeClassDerive) {
		tccache.WillGenerated(v)
	})

	moduleInf := option.FlatMap(seq.Head(pkgs),
		option.Compose3(
			func(p *packages.Package) fp.Option[packages.Module] { return option.Ptr(p.Module) },
			option.Pure1(func(m packages.Module) string { return m.GoVersion }),
			option.Pure1(func(v string) bool { return v >= "1.21" }),
		),
	).OrElse(true)

	return &TypeClassSummonContext{
		w:                     importSet,
		fpPkg:                 fpPkg,
		tcCache:               &tccache,
		recursiveGen:          derives,
		implicitTypeInference: implicitTypeInference && moduleInf,
	}
}

func genDerive() {

	pack := os.Getenv("GOPACKAGE")

	genfp.Generate(pack, derive_generated_file_name(pack), func(w genfp.Writer) {

		cwd, _ := os.Getwd()

		cfg := &packages.Config{
			Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
		}

		pkgs, err := packages.Load(cfg, cwd)
		if err != nil {
			fmt.Println(err)
			return
		}

		// fmtalias := w.GetImportedName(genfp.NewImportPackage("fmt", "fmt"))
		// asalias := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		summonCtx := NewTypeClassSummonContext(pkgs, w)
		if summonCtx.recursiveGen.Size() == 0 {
			return
		}

		for len(summonCtx.recursiveGen) > 0 {
			d := summonCtx.recursiveGen
			summonCtx.recursiveGen = nil
			d.Foreach(func(v metafp.TypeClassDerive) {
				summonCtx.summonVar(v).Foreach(func(v SummonExpr) {
					fmt.Fprintf(w, "%s\n", v.expr)
				})
			})
		}

	})
}
