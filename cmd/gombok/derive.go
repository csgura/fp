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
	"github.com/csgura/fp/genfp/generator"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
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
	hlistPkg              fp.Option[*types.Package]
	tcCache               *metafp.TypeClassInstanceCache
	derived               fp.Map[string, TypeClassInstanceGenerated]
	loopCheck             fp.Set[string]
	recursiveGen          fp.Seq[metafp.TypeClassDerive]
	initVarSet            mutable.Set[string]
	initVars              fp.Seq[generator.TaggedVar]
	implicitTypeInference bool
}

type DeriveContext struct {
	working      genfp.WorkingPackage
	tc           metafp.TypeClassDerive
	primScope    metafp.TypeClassScope
	workingScope metafp.TypeClassScope
	recursiveGen bool
	noinline     bool
}

type SummonContext struct {
	working       genfp.WorkingPackage
	typeClass     metafp.TypeClass
	summonFor     string
	deriveContext fp.Option[DeriveContext]
}

func (r SummonContext) primScope(tcCache *metafp.TypeClassInstanceCache, req metafp.TypeClass) metafp.TypeClassScope {
	if r.deriveContext.IsDefined() {
		ctx := r.deriveContext.Get()
		scope := ctx.primScope
		if req.Id() != ctx.tc.TypeClass.Id() {
			scope = tcCache.GetImported(req)
		}
		return scope
	}
	return tcCache.GetImported(req)
}

func (r SummonContext) workingScope(tcCache *metafp.TypeClassInstanceCache, req metafp.TypeClass) metafp.TypeClassScope {
	if r.deriveContext.IsDefined() {
		ctx := r.deriveContext.Get()
		scope := ctx.workingScope
		if req.Id() != ctx.tc.TypeClass.Id() {
			scope = tcCache.GetLocal(ctx.working.Package(), req)
		}
		return scope
	}
	return tcCache.GetLocal(r.working.Package(), req)
}

func (r SummonContext) recursiveGen() bool {
	return option.Map(r.deriveContext, func(v DeriveContext) bool {
		return v.recursiveGen
	}).OrZero()
}

func (r SummonContext) recursiveDerive(required metafp.RequiredInstance, notDefined NotDefinedInstance) fp.Option[metafp.TypeClassDerive] {
	return option.FlatMap(r.deriveContext, func(ctx DeriveContext) fp.Option[metafp.TypeClassDerive] {
		if ctx.recursiveGen && required.TypeClass.Id() == ctx.tc.TypeClass.Id() {
			named := notDefined.instanceOf.AsNamed()
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

				return option.Some(tc)
			}
		}
		return option.None[metafp.TypeClassDerive]()
	})
}

type GenericRepr struct {
	//	ReprType     func() string
	Kind         string
	Type         func() string
	ReprType     func() string
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
	expr          func() string
	paramInstance fp.Seq[ParamInstance]
}

func (r SummonExpr) Expr() string {
	return try.Of(r.expr).Get()
}
func (r SummonExpr) String() string {
	return try.Of(r.expr).Get()
}

func (r SummonExpr) ParamInstance() fp.Seq[ParamInstance] {
	return r.paramInstance
}

func collectSummonExpr(list fp.Seq[SummonExpr]) SummonExpr {
	expr := func() string {
		return seq.Map(list, SummonExpr.Expr).MakeString(",")
	}
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
		expr:          as.Supplier(r.name),
		paramInstance: param,
	}
}

func (r ArgumentInstance) Required() fp.Seq[metafp.RequiredInstance] {
	return nil
}

type DefinedInstance struct {
	instance   metafp.TypeClassInstance
	searchName string
	checked    bool
}

func (r DefinedInstance) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	if r.instance.Package == nil || r.instance.Package.Path() == workingPkg.Path() {
		return SummonExpr{
			expr: as.Supplier(r.instance.Name),
		}
	}

	return SummonExpr{
		expr: func() string {
			pk := w.GetImportedName(genfp.FromTypesPackage(r.instance.Package))
			return fmt.Sprintf("%s.%s", pk, r.instance.Name)
		},
	}
}

func (r DefinedInstance) Instance() metafp.TypeClassInstance {
	return r.instance
}

func (r DefinedInstance) Required() fp.Seq[metafp.RequiredInstance] {
	return r.instance.RequiredInstance
}

// func (r DefinedInstance) IsLocal() bool {
// 	return r.local
// }

type NotDefinedInstance struct {
	instanceOf metafp.TypeInfo
	name       string
	required   fp.Seq[metafp.RequiredInstance]
}

func (r NotDefinedInstance) instanceExpr(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	return SummonExpr{
		expr: as.Supplier(r.name),
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

func constTrue[T any](T) bool {
	return true
}
func (r lookupTarget) checked() bool {
	return either.Fold(
		r.target,
		constTrue,
		as.Func3(either.Fold[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance], bool]).ApplyLast2(
			constTrue,
			as.Func3(either.Fold[ArgumentInstance, DefinedInstance, bool]).ApplyLast2(
				constTrue,
				func(di DefinedInstance) bool {
					return di.checked
				},
			),
		),
	)
}

func (r lookupTarget) isGivenAny() bool {
	return option.Map(r.instance(), metafp.TypeClassInstance.IsGivenAny).OrElse(false)
}

// func (r lookupTarget) isLocal() bool {
// 	return either.Fold(
// 		r.target,
// 		fp.Const[NotDefinedInstance](false),
// 		as.Func3(either.Fold[SummonExprInstance, fp.Either[ArgumentInstance, DefinedInstance], bool]).ApplyLast2(
// 			fp.Const[SummonExprInstance](false),
// 			as.Func3(either.Fold[ArgumentInstance, DefinedInstance, bool]).ApplyLast2(
// 				fp.Const[ArgumentInstance](false),
// 				DefinedInstance.IsLocal,
// 			),
// 		),
// 	)
// }

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

func (r *TypeClassSummonContext) typeclassInstanceMust(ctx SummonContext, req metafp.RequiredInstance, name string) lookupTarget {

	lt := r.namedLookup(ctx, req, false, name)
	if lt.IsDefined() {
		return TypeClassInstanceToLookupTarget(lt.Get())
	}

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
func (r *TypeClassSummonContext) lookupTypeClassInstanceLocalDeclared(ctx SummonContext, req metafp.RequiredInstance, strict bool, name ...string) fp.Option[DefinedInstance] {

	f := req.Type

	scope := ctx.workingScope(r.tcCache, req.TypeClass)
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
			expr := r.deriveFuncExpr(tci.WillGeneratedBy.Get())
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

	filtered := iterator.FilterMap(ins, func(tci metafp.TypeClassInstance) fp.Option[DefinedInstance] {
		none := option.None[DefinedInstance]()

		if r.initVarSet.Contains(tci.Name) {
			return none
		}

		if tci.IsGivenAny() && ctx.recursiveGen() && isRecursiveDerivable(req) {
			return none
			//fmt.Printf("%s is recursive derivable\n", req.Type)
		}

		check := r.checkRequired(ctx, tci, tci.RequiredInstance)
		if strict && !check {
			return none
		}
		return option.Some(DefinedInstance{
			instance:   tci,
			checked:    check,
			searchName: tci.Name,
		})
	})

	// instance 가 있는 경우 , instance 가 Some
	ret := filtered.NextOption()

	return ret
}

func (r *TypeClassSummonContext) lookupHConsMust(ctx SummonContext, tc metafp.TypeClass) metafp.TypeClassInstance {
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

func (r *TypeClassSummonContext) lookupHNilMust(ctx SummonContext, tc metafp.TypeClass) metafp.TypeClassInstance {
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
func (r *TypeClassSummonContext) lookupTypeClassFunc(ctx SummonContext, tc metafp.TypeClass, name string) fp.Option[metafp.TypeClassInstance] {
	nameWithTc := tc.Name + name

	workingScope := ctx.workingScope(r.tcCache, tc)
	primScope := ctx.primScope(r.tcCache, tc)

	ins := workingScope.FindFunc(nameWithTc)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindFunc(nameWithTc)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindFunc(name)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	return option.None[metafp.TypeClassInstance]()
}

func (r *TypeClassSummonContext) lookupTypeClassFuncCheckType(ctx SummonContext, tc metafp.TypeClass, name string, argType metafp.TypeInfo) fp.Option[metafp.TypeClassInstance] {

	workingScope := ctx.workingScope(r.tcCache, tc)
	primScope := ctx.primScope(r.tcCache, tc)

	ins := workingScope.FindByNamePrefix(name, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	ins = primScope.FindByNamePrefix(name, argType)
	if ins.IsDefined() && r.checkRequired(ctx, ins.Get(), ins.Get().RequiredInstance) {
		return ins
	}

	return option.None[metafp.TypeClassInstance]()
}

func (r *TypeClassSummonContext) lookupTupleTypeClassFunc(ctx SummonContext, tc metafp.TypeClass, name string, tupleArgs fp.Seq[metafp.TypeInfoExpr]) fp.Option[metafp.TypeClassInstance] {
	none := option.None[metafp.TypeClassInstance]()
	if r.fpPkg.IsDefined() {
		tupleDef := r.fpPkg.Get().Scope().Lookup(fmt.Sprintf("Tuple%d", tupleArgs.Size()))
		if tupleDef == nil {
			return none
		}
		tupleType := metafp.GetTypeInfo(tupleDef.Type())
		targs := seq.Map(tupleArgs, func(v metafp.TypeInfoExpr) metafp.TypeInfo {
			return v.Type
		})

		it, err := tupleType.Instantiate(targs).Unapply()
		if err != nil {
			return none
		}

		return r.lookupTypeClassFuncCheckType(ctx, tc, name, it)

	}

	return none
}

func (r *TypeClassSummonContext) lookupTypeClassFuncMust(ctx SummonContext, tc metafp.TypeClass, name string) metafp.TypeClassInstance {
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

func (r *TypeClassSummonContext) lookupTypeClassInstancePrimitivePkgLazy(ctx SummonContext, req metafp.RequiredInstance, strict bool, name ...string) func() fp.Option[DefinedInstance] {
	return func() fp.Option[DefinedInstance] {
		return r.lookupTypeClassInstancePrimitivePkg(ctx, req, strict, name...)
	}
}

func (r *TypeClassSummonContext) checkRequired(ctx SummonContext, tci metafp.TypeClassInstance, required fp.Seq[metafp.RequiredInstance]) bool {
	verbose("check required for %s", tci.Name)
	for _, v := range required {
		//fmt.Printf("check %s required of %s\n", v.String(), tci.String())
		if v.Name {
			continue
		}
		if v.Type.IsTuple() {
			req := seq.Map(v.Type.TypeArgs, func(t metafp.TypeInfo) metafp.RequiredInstance {
				return metafp.RequiredInstance{
					TypeClass: v.TypeClass,
					Type:      t,
				}
			})
			res := r.checkRequired(ctx, tci, req)
			if !res {
				return false
			}

		} else {
			// TODO: summonArgs에서 다시  lookup 하는 코드 있음.
			verbose("lookup type class %s[%s], for check required for %s ", v.TypeClass.Name, v.Type, tci.Name)

			res := r.lookupTypeClassInstance(ctx, v)
			verbose("lookup type class result = %s(%s[%s]), for check required for %s ", res.target, v.TypeClass.Name, v.Type, tci.Name)
			if res.target.IsLeft() {
				tc, rgen := ctx.recursiveDerive(v, res.target.Left()).Unapply()
				if rgen {
					if !r.tcCache.IsWillGenerated(tc) {
						r.recursiveGen = append(r.recursiveGen, tc)
						r.tcCache.WillGenerated(tc)
					}
					continue
				}
				return false
			} else if res.checked() == false {
				tc, rgen := ctx.recursiveDerive(v, NotDefinedInstance{
					instanceOf: v.Type,
					name:       v.TypeClass.Name + publicName(v.Type.TypeName),
					required: seq.Map(v.Type.TypeArgs, func(ti metafp.TypeInfo) metafp.RequiredInstance {
						return metafp.RequiredInstance{
							TypeClass: v.TypeClass,
							Type:      ti,
						}
					}),
				}).Unapply()
				if rgen {
					if !r.tcCache.IsWillGenerated(tc) {
						r.recursiveGen = append(r.recursiveGen, tc)
						r.tcCache.WillGenerated(tc)
					}
					continue
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
	if req.Type.IsNamed() && !req.Type.IsAlias() {
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

func (r *TypeClassSummonContext) lookupTypeClassInstancePrimitivePkg(ctx SummonContext, req metafp.RequiredInstance, strict bool, name ...string) fp.Option[DefinedInstance] {

	f := req.Type

	verbose("lokkup tc instance %s[%s] in primitive pkg, name first = %s, f.TypeArgs.Size = %d", req.TypeClass.Name, req.Type, as.Seq(name).MakeString(","), f.TypeArgs.Size())
	scope := ctx.primScope(r.tcCache, req.TypeClass)

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

	filtered := iterator.FilterMap(ins, func(tci metafp.TypeClassInstance) fp.Option[DefinedInstance] {

		none := option.None[DefinedInstance]()
		if isSamePkg(ctx.working, genfp.FromTypesPackage(tci.Package)) {
			if r.initVarSet.Contains(tci.Name) {
				return none
			}
		}

		//fmt.Printf("result for %s[%s] is %s, is given %t\n", req.TypeClass.Name, req.Type, tci.Name, tci.IsGivenAny())

		if tci.IsGivenAny() && ctx.recursiveGen() && isRecursiveDerivable(req) {
			return none
			//fmt.Printf("%s is recursive derivable\n", req.Type)
		}
		//fmt.Printf("checkRequired for %s[%s] is %s, \n", req.TypeClass.Name, req.Type, tci.Name)

		check := r.checkRequired(ctx, tci, tci.RequiredInstance)
		if strict && !check {
			return none
		}
		return option.Some(DefinedInstance{
			instance:   tci,
			checked:    check,
			searchName: tci.Name,
		})
		//fmt.Printf("result for %s[%s] is %s -> %t, \n", req.TypeClass.Name, req.Type, tci.Name, ret)

	})

	// instance 가 있는 경우 , instance 가 Some
	return filtered.NextOption()

}

func (r *TypeClassSummonContext) lookupTypeClassInstanceTypePkg(ctx SummonContext, req metafp.RequiredInstance, strict bool, name string) fp.Option[DefinedInstance] {

	f := req.Type
	if f.Pkg != nil && f.Pkg.Path() != ctx.working.Path() {

		name := req.TypeClass.Name + publicName(name)
		obj := f.Pkg.Scope().Lookup(name)

		if obj != nil {

			ti := metafp.GetTypeInfo(obj.Type())
			rhsType := ti.ResultType()
			if rhsType.IsInstanceOf(ctx.typeClass) {
				return option.Map(metafp.AsTypeClassInstance(req.TypeClass, obj), func(a metafp.TypeClassInstance) DefinedInstance {
					return DefinedInstance{
						instance:   a,
						checked:    true,
						searchName: name,
					}
				})

			}

		}
	}

	return option.None[DefinedInstance]()
}

func (r *TypeClassSummonContext) namedLookup(ctx SummonContext, req metafp.RequiredInstance, strict bool, name string) fp.Option[DefinedInstance] {

	verbose("named lookup name = %s, type = %s[%s], strict %t", name, req.TypeClass.Name, req.Type, strict)
	localInsOpt := r.lookupTypeClassInstanceLocalDeclared(ctx, req, strict, name)
	if localInsOpt.IsDefined() {

		localIns := localInsOpt.Get()
		if !localIns.instance.IsGivenAny() {
			verbose("named lookup name = %s, type = %s[%s]. found = %s", name, req.TypeClass.Name, req.Type, localIns.instance.Name)

			return localInsOpt

		}

	}

	ret := r.lookupTypeClassInstanceTypePkg(ctx, req, strict, name).
		Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, strict, name))

	if localInsOpt.IsDefined() && ret.IsDefined() {
		retIns := ret.Get()
		if retIns.instance.IsGivenAny() {
			verbose("named lookup name = %s, type = %s[%s]. found = %s", name, req.TypeClass.Name, req.Type, localInsOpt.Get().instance.Name)

			return localInsOpt
		}

	}

	result := ret.OrOption(localInsOpt)
	verbose("named lookup name = %s, type = %s[%s]. found = %t", name, req.TypeClass.Name, req.Type, result.IsDefined())

	return result

}

func TypeClassInstanceToLookupTarget(a DefinedInstance) lookupTarget {
	return lookupTarget{
		target: either.Right[NotDefinedInstance](either.Right[SummonExprInstance](either.Right[ArgumentInstance](a))),
	}
}
func (r *TypeClassSummonContext) orMust(ctx SummonContext, req metafp.RequiredInstance, name string, ins fp.Option[DefinedInstance]) lookupTarget {
	lt := option.Map(ins, TypeClassInstanceToLookupTarget)
	return lt.OrElse(r.typeclassInstanceMust(ctx, req, name))

}
func (r *TypeClassSummonContext) namedLookupMust(ctx SummonContext, req metafp.RequiredInstance, name string) lookupTarget {
	return r.orMust(ctx, req, name, r.namedLookup(ctx, req, true, name))
}

// func (r *TypeClassSummonContext) lookupPrimitiveTypeClassInstance(ctx CurrentContext, req metafp.RequiredInstance, name ...string) lookupTarget {
// 	ret := r.lookupTypeClassInstanceLocalDeclared(ctx, req, name...).Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, name...))

// 	return ret.OrElse(r.typeclassInstanceMust(ctx, req, name[0]))

// }

// 타입 추론이 가능한지 따지는 함수
func (r *TypeClassSummonContext) typeParamStringOfLookupTarget(ctx SummonContext, lt lookupTarget) fp.Option[string] {

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

func (r *TypeClassSummonContext) typeParamString(ctx SummonContext, ins metafp.TypeClassInstance, explicit bool) fp.Option[string] {

	if explicit {
		ret := seq.Map(ins.TypeParam, func(v metafp.TypeParam) string {
			return option.Map(ins.ParamMapping.Get(v.Name), func(v metafp.TypeInfo) string {
				return r.w.TypeName(ctx.working, v.Type)
			}).OrElse(v.Name)
		}).MakeString(",")
		return option.Some(ret)
	}

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

func (r *TypeClassSummonContext) summonArgs(ctx SummonContext, args fp.Seq[metafp.RequiredInstance]) SummonExpr {
	list := seq.Map(args, func(t metafp.RequiredInstance) SummonExpr {
		// TODO: checkRequired 에서  lookup 하는 코드 있음. checkRequired 에서 한번 했으면 안하게 할 필요 있음.
		ret := r.summonRequired(ctx, t)
		if t.Lazy {

			return newSummonExpr(func() string {
				lazypk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/lazy", "lazy"))
				return fmt.Sprintf(`%s.Call( func() %s[%s] {
				return %s
			})`, lazypk, t.TypeClass.PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type.Type), ret)
			}, ret.paramInstance)
		}
		return ret
	})

	return collectSummonExpr(list)
}

func (r *TypeClassSummonContext) summonArgsWithRequiredFound(ctx SummonContext, args fp.Seq[fp.Tuple2[metafp.RequiredInstance, metafp.TypeClassInstance]], explicit bool) SummonExpr {
	list := seq.Map(args, func(t fp.Tuple2[metafp.RequiredInstance, metafp.TypeClassInstance]) SummonExpr {
		// TODO: checkRequired 에서  lookup 하는 코드 있음. checkRequired 에서 한번 했으면 안하게 할 필요 있음.
		ret := r.exprTypeClassInstance(ctx, t.I2, explicit)
		if t.I1.Lazy {

			return newSummonExpr(func() string {
				lazypk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/lazy", "lazy"))
				return fmt.Sprintf(`%s.Call( func() %s[%s] {
				return %s
			})`, lazypk, t.I1.TypeClass.PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.I1.Type.Type), ret)
			}, ret.paramInstance)
		}
		return ret
	})

	return collectSummonExpr(list)
}

func newSummonExpr(expr func() string, params ...fp.Seq[ParamInstance]) SummonExpr {
	return SummonExpr{
		expr:          expr,
		paramInstance: seq.Reduce(as.Seq(params), MergeSeqDistinct(EqParamInstance)),
	}
}

func instanceExprOfTypeClassInstance(r metafp.TypeClassInstance, w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
	if r.Package == nil || r.Package.Path() == workingPkg.Path() {
		return SummonExpr{
			expr: as.Supplier(r.Name),
		}
	}

	return SummonExpr{
		expr: func() string {
			pk := w.GetImportedName(genfp.FromTypesPackage(r.Package))
			return fmt.Sprintf("%s.%s", pk, r.Name)
		},
	}
}

func (r *TypeClassSummonContext) exprTypeClassInstance(ctx SummonContext, lt metafp.TypeClassInstance, explicit bool) SummonExpr {
	//fmt.Printf("lt : %s, %v\n", lt.instance(), lt.required())

	if len(lt.RequiredInstance) > 0 {
		list := r.summonArgs(ctx, lt.RequiredInstance)

		instanceExpr := instanceExprOfTypeClassInstance(lt, r.w, ctx.working)
		retExpr := func() string {
			tpstr := r.typeParamString(ctx, lt, explicit)

			if tpstr.IsDefined() {
				return fmt.Sprintf("%s[%s](%s)", instanceExpr, tpstr.Get(), list)
			} else {
				return fmt.Sprintf("%s(%s)", instanceExpr, list)

			}
		}

		return newSummonExpr(
			retExpr, instanceExpr.paramInstance, list.paramInstance)

	}

	if !lt.Static && len(lt.RequiredInstance) == 0 {
		instanceExpr := instanceExprOfTypeClassInstance(lt, r.w, ctx.working)

		retExpr := func() string {
			tpstr := r.typeParamString(ctx, lt, false)
			if tpstr.IsDefined() {
				return fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get())
			} else {
				return fmt.Sprintf("%s()", instanceExpr)
			}
		}

		return newSummonExpr(retExpr, instanceExpr.paramInstance)

	}

	return instanceExprOfTypeClassInstance(lt, r.w, ctx.working)

}

func (r *TypeClassSummonContext) exprTypeClassInstanceWithRequiredFound(ctx SummonContext, lt metafp.TypeClassInstance, explicit bool, requiredIns fp.Seq[metafp.TypeClassInstance]) SummonExpr {
	//fmt.Printf("lt : %s, %v\n", lt.instance(), lt.required())

	if len(lt.RequiredInstance) > 0 {
		list := r.summonArgsWithRequiredFound(ctx, seq.Zip(lt.RequiredInstance, requiredIns), explicit)

		instanceExpr := instanceExprOfTypeClassInstance(lt, r.w, ctx.working)
		retExpr := func() string {
			tpstr := r.typeParamString(ctx, lt, explicit)

			if tpstr.IsDefined() {
				return fmt.Sprintf("%s[%s](%s)", instanceExpr, tpstr.Get(), list)
			} else {
				return fmt.Sprintf("%s(%s)", instanceExpr, list)

			}
		}

		return newSummonExpr(
			retExpr, instanceExpr.paramInstance, list.paramInstance)

	}

	if !lt.Static && len(lt.RequiredInstance) == 0 {
		instanceExpr := instanceExprOfTypeClassInstance(lt, r.w, ctx.working)

		retExpr := func() string {
			tpstr := r.typeParamString(ctx, lt, false)
			if tpstr.IsDefined() {
				return fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get())
			} else {
				return fmt.Sprintf("%s()", instanceExpr)
			}
		}

		return newSummonExpr(retExpr, instanceExpr.paramInstance)

	}

	return instanceExprOfTypeClassInstance(lt, r.w, ctx.working)

}

func (r *TypeClassSummonContext) exprLookupTarget(ctx SummonContext, lt lookupTarget) SummonExpr {
	//fmt.Printf("lt : %s, %v\n", lt.instance(), lt.required())

	if len(lt.required()) > 0 {
		list := r.summonArgs(ctx, lt.required())

		instanceExpr := lt.instanceExpr(r.w, ctx.working)
		retExpr := func() string {
			expr := instanceExpr.String()
			tpstr := r.typeParamStringOfLookupTarget(ctx, lt)

			if tpstr.IsDefined() {
				return fmt.Sprintf("%s[%s](%s)", expr, tpstr.Get(), list)
			} else {
				return fmt.Sprintf("%s(%s)", expr, list)
			}
		}

		//fmt.Printf("%s param infer not possible = %s \n", lt.name, lt.instance.Get().ParamMapping)

		return newSummonExpr(retExpr, instanceExpr.paramInstance, list.paramInstance)

	}

	if lt.isFunc() && len(lt.required()) == 0 {
		instanceExpr := lt.instanceExpr(r.w, ctx.working)

		retExpr := func() string {
			tpstr := r.typeParamStringOfLookupTarget(ctx, lt)

			if tpstr.IsDefined() {
				return fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get())
			} else {
				return fmt.Sprintf("%s()", instanceExpr)
			}
		}

		return newSummonExpr(retExpr, instanceExpr.paramInstance)

	}

	return lt.instanceExpr(r.w, ctx.working)

}

func (r *TypeClassSummonContext) exprTypeClassMember(ctx SummonContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typeArgs fp.Seq[metafp.TypeInfoExpr], fieldOf fp.Option[metafp.TypeInfo]) SummonExpr {
	if len(typeArgs) > 0 {

		list := r.summonArgs(ctx, seq.Map(typeArgs, func(t metafp.TypeInfoExpr) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: tc,
				Type:      t.Type,
			}
		}))

		retExpr := func() string {
			return fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, ctx.working), list)
		}
		return newSummonExpr(retExpr, list.paramInstance)
	}

	return newSummonExpr(func() string {
		return lt.PackagedName(r.w, ctx.working)
	})

}

func (r *TypeClassSummonContext) exprTypeClassMemberLabelled(ctx SummonContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typePkg *types.Package, structName string, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfoExpr], genLabelled bool) SummonExpr {
	if len(typeArgs) > 0 {
		list := collectSummonExpr(seq.Map(seq.Zip(typeArgs, names), func(t fp.Tuple2[metafp.TypeInfoExpr, string]) SummonExpr {
			return r.summonFpNamed(ctx, tc, typePkg, structName, t.I2, t.I1, genLabelled)
		}))

		retExpr := func() string {
			return fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, ctx.working), list)
		}

		return newSummonExpr(retExpr, list.paramInstance)
	}

	return newSummonExpr(func() string {
		return lt.PackagedName(r.w, ctx.working)
	})
}

func (r *TypeClassSummonContext) lookupTypeClassInstance(ctx SummonContext, req metafp.RequiredInstance) lookupTarget {
	verbose("lookup tc instance %s[%s]", req.TypeClass.Name, req.Type)

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
	case *types.Alias:
		ret := r.namedLookup(ctx, req, true, at.Obj().Name())
		if ret.IsDefined() {
			return TypeClassInstanceToLookupTarget(ret.Get())
		}
		req.Type = req.Type.Rhs()
		return r.lookupTypeClassInstance(ctx, req)
	case *types.Named:
		if at.Obj().Pkg() != nil && at.Obj().Pkg().Path() == "github.com/csgura/fp/hlist" {
			//fmt.Printf("lookup named hlist %s\n", req.Type)

			if at.Obj().Name() == "Nil" {
				return option.Map(r.lookupTypeClassInstanceLocalDeclared(ctx, req, true, "HNil", "HListNil").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, true, "HNil", "HListNil")), TypeClassInstanceToLookupTarget).OrElse(r.typeclassInstanceMust(ctx, req, "HNil"))

			} else if at.Obj().Name() == "Cons" {
				return option.Map(r.lookupTypeClassInstanceLocalDeclared(ctx, req, true, "HCons", "HListCons").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, true, "HCons", "HListCons")), TypeClassInstanceToLookupTarget).OrElse(r.typeclassInstanceMust(ctx, req, "HCons"))
			}
		}
		return r.namedLookupMust(ctx, req, at.Obj().Name())
	case *types.Array:
		//panic(fmt.Sprintf("can't summon array type, while deriving %s[%s]", req.TypeClass.Name, ctx.summonFor))
		//return r.namedLookup(f, "Array")

		ret := func(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
			sreq := req
			sreq.Type = metafp.GetTypeInfo(types.NewSlice(at.Elem()))

			return r.summonGeneric(ctx, ctx.typeClass, r.w.TypeName(ctx.working, at), GenericRepr{
				Kind: fp.GenericKindConversion,
				Type: func() string {
					return r.w.TypeName(ctx.working, at)
				},
				ReprType: func() string {
					return r.w.TypeName(ctx.working, sreq.Type.Type)
				},
				ToReprExpr: func() string {
					return fmt.Sprintf(`func(v %s) %s {
						return v[:]
					}`, r.w.TypeName(ctx.working, at), r.w.TypeName(ctx.working, sreq.Type.Type))
				},
				FromReprExpr: func() string {
					return fmt.Sprintf(`func(v %s) %s {
						return %s(v)
					}`, r.w.TypeName(ctx.working, sreq.Type.Type), r.w.TypeName(ctx.working, at), r.w.TypeName(ctx.working, at))
				},
				ReprExpr: func() SummonExpr {
					return r.summonRequired(ctx, sreq)
				},
			})
		}
		return lookupTarget{
			target: either.Right[NotDefinedInstance](either.NotRight[fp.Either[ArgumentInstance, DefinedInstance]](SummonExprInstance{ret})),
		}

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
		samePkg := func() bool {
			if at.NumFields() > 0 {
				return isSamePkg(ctx.working, genfp.FromTypesPackage(at.Field(0).Pkg()))
			}
			return true
		}()

		if fields.ForAll(metafp.StructField.Public) || samePkg {
			ret := func(w genfp.ImportSet, workingPkg genfp.WorkingPackage) SummonExpr {
				return r.summonUntypedStruct(ctx, req.TypeClass, f, fields)
			}
			return lookupTarget{
				target: either.Right[NotDefinedInstance](either.NotRight[fp.Either[ArgumentInstance, DefinedInstance]](SummonExprInstance{ret})),
			}
		} else {
			//fmt.Printf("fieldOf = %v\n", req.FieldOf)
			panic(fmt.Sprintf("can't summon unnamed struct type %v containing private field, while deriving %s[%s]", f.Type, ctx.typeClass.Name, ctx.summonFor))
		}

	case *types.Interface:
		if f.IsAny() {
			return r.namedLookupMust(ctx, req, "Given")
		}
		panic(fmt.Sprintf("can't summon unnamed interface type %v, while deriving %s[%s]", f.Type, ctx.typeClass.Name, ctx.summonFor))
	case *types.Chan:
		panic(fmt.Sprintf("can't summon unnamed chan type, while deriving %s[%s]", ctx.typeClass.Name, ctx.summonFor))

	}
	// f.Type.String() 은 com/uangel/ 이런식으로 이상하게 출력됨.
	return r.namedLookupMust(ctx, req, f.Type.String())
}

// v.A , v.B
func (r *TypeClassSummonContext) structUnapplyExpr(ctx SummonContext, named fp.Option[metafp.NamedTypeInfo], fields fp.Seq[metafp.StructField], varexpr string) string {
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
func (r *TypeClassSummonContext) structApplyExpr(ctx SummonContext, named fp.Option[metafp.NamedTypeInfo], fields fp.Seq[metafp.StructField], args ...string) string {
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
	argslist := seq.Map(seq.Zip(names, args), func(v fp.Entry[string]) string {
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

func (r *TypeClassSummonContext) namedOrRuntimeType(working genfp.WorkingPackage, typePkg *types.Package, structName string, name string, vtype metafp.TypeInfoExpr, genLabelled bool) metafp.TypeInfo {

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

	rtobj := r.fpPkg.Get().Scope().Lookup("RuntimeNamed")
	if tpe, ok := rtobj.(*types.TypeName); ok {
		ctx := types.NewContext()
		targs := []types.Type{vtype.Type.Type}
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

func (r *TypeClassSummonContext) namedOrRuntimeStringExpr(w genfp.ImportSet, working genfp.WorkingPackage, typePkg *types.Package, structName string, name string, labelledGen bool, vtype metafp.TypeInfoExpr) string {

	ret := r.namedOrRuntimeType(working, typePkg, structName, name, vtype, labelledGen)
	return r.w.TypeName(working, ret.Type)
	// if labelledGen {
	// 	ret := publicName(name)
	// 	if ret == name {
	// 		ret = fmt.Sprintf("PubNamed%sOf%s", ret, structName)
	// 	} else {
	// 		ret = fmt.Sprintf("Named%sOf%s", ret, structName)
	// 	}

	// 	if isSamePkg(working, genfp.FromTypesPackage(typePkg)) {
	// 		return ret
	// 	} else {
	// 		return fmt.Sprintf("%s.%s", w.GetImportedName(genfp.FromTypesPackage(typePkg)), ret)
	// 	}
	// } else {
	// 	fppk := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

	// 	return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, valueType)

	// }

}

var implicitTypeInference = option.Of(runtime.Version()).Filter(func(v string) bool { return v >= "1.21.0" }).IsDefined()

func (r *TypeClassSummonContext) summonLabelledGenericRepr(ctx SummonContext, tc metafp.TypeClass, sf structFunctions) fp.Option[GenericRepr] {

	type fieldName = fp.Entry[string]
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
			Kind:         fp.GenericKindStruct,
			Type:         as.Supplier1(sf.typeStr, ctx.working),
			ReprType:     r.labelledTupleReprType(ctx, sf),
			ToReprExpr:   r.intoLabelledTupleRepr(ctx, sf),
			FromReprExpr: r.fromLabelledTupleRepr(ctx, sf),
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassMemberLabelled(ctx, tc, tm, sf.pack, sf.name, seq.Map(names, xtr.Head), typeArgs, sf.namedGenerated)
			},
		}
	}).Or(func() fp.Option[GenericRepr] {
		return option.Map(r.lookupTypeClassFunc(ctx, tc, "HConsLabelled"), func(hcons metafp.TypeClassInstance) GenericRepr {
			return GenericRepr{
				Kind:         fp.GenericKindStruct,
				Type:         as.Supplier1(sf.typeStr, ctx.working),
				ReprType:     r.labelledHlistReprType(ctx, sf),
				ToReprExpr:   r.toLabelledHlistRepr(ctx, sf, hcons.Result.TypeArgs.Head()),
				FromReprExpr: r.fromLabelledHlistRepr(ctx, sf, hcons.Result.TypeArgs.Head()),
				ReprExpr: func() SummonExpr {
					//arity := fp.Min(typeArgs.Size(), max.Product-1)
					arity := typeArgs.Size()

					hnil := r.lookupHNilMust(ctx, tc)
					namedTypeArgs := seq.Zip(names, typeArgs)
					hlist := seq.Fold(namedTypeArgs.Take(arity).Reverse(), newSummonExpr(func() string { return hnil.PackagedName(r.w, ctx.working) }), func(tail SummonExpr, ti fp.Tuple2[fieldName, metafp.TypeInfoExpr]) SummonExpr {
						instance := r.summonFpNamed(ctx, tc, sf.pack, sf.name, ti.I1.I1, ti.I2, sf.namedGenerated)
						return newSummonExpr(func() string {

							return fmt.Sprintf(`%s(
									%s,
									%s,
								)`, hcons.PackagedName(r.w, ctx.working), instance, tail,
							)
						}, instance.paramInstance, tail.paramInstance)
					})

					return hlist
				},
			}
		})
	})
}

func (r *TypeClassSummonContext) namedStructFuncs(ctx SummonContext, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) structFunctions {
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

	type fieldName = fp.Entry[string]
	names := seq.Map(fields, func(v metafp.StructField) fieldName {
		return as.Tuple(v.Name, v.Tag)
	})

	// tupleFuncExpr := func() string {
	// 	return fmt.Sprintf("%s.AsTuple", typeStr(ctx.working))
	// }

	// applyFuncExpr := func() string {
	// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
	// 	aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

	// 	builderreceiver := builderTypeStr(ctx.working)
	// 	return fmt.Sprintf(`%s.Compose(
	// 				%s.Curried2(%s.FromTuple)(%s{}),
	// 				%s.Build,
	// 				)`,
	// 		fppk,
	// 		aspk, builderreceiver, builderreceiver, builderreceiver,
	// 	)
	// }

	unapplyFunc := func(structIns string) string {
		return fmt.Sprintf("%s.Unapply()", structIns)
	}

	applyFunc := func(fieldValues []string) string {
		return fmt.Sprintf(`%s{}.Apply(%s).Build()`,
			builderTypeStr(ctx.working), as.Seq(fieldValues).MakeString(","))
	}
	if !hasUnapply {

		// tupleFuncExpr = func() string {
		// 	p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
		// 		return f.TypeName(r.w, ctx.working)
		// 	}).MakeString(",")

		// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		// 	aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		// 	return fmt.Sprintf(`func( v %s) %s.Tuple%d[%s] {
		// 	return %s.Tuple%d(%s)
		// }`, typeStr(ctx.working), fppk, fields.Size(), p,
		// 		aspk, fields.Size(), seq.Map(names, func(v fieldName) string { return "v." + v.I1 }).MakeString(","),
		// 	)
		// }

		// applyFuncExpr = func() string {
		// 	p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
		// 		return f.TypeName(r.w, ctx.working)
		// 	}).MakeString(",")

		// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
		// 	//aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		// 	assign := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
		// 		return fmt.Sprintf("%s : t.I%d", v.I2.I1, v.I1+1)
		// 	}).MakeString(",\n")
		// 	valuereceiver := typeStr(ctx.working)
		// 	return fmt.Sprintf(`func(t %s.Tuple%d[%s]) %s {
		// 			return %s{
		// 				%s,
		// 			}
		// 		}`, fppk, fields.Size(), p, valuereceiver,
		// 		valuereceiver,
		// 		assign,
		// 	)
		// }

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

	sf := structFunctions{
		pack:           named.Package,
		name:           named.Name,
		tpe:            named.Info,
		fields:         fields,
		names:          names,
		typeArgs:       typeArgs,
		valueGenerated: hasUnapply,
		builderTypeStr: builderTypeStr,
		namedGenerated: hasAsLabelled,
		typeStr:        typeStr,
		unapply:        unapplyFunc,
		apply:          applyFunc,
	}
	return sf
}

func (r *TypeClassSummonContext) untypedStructFuncs(ctx SummonContext, tpe metafp.TypeInfo, fields fp.Seq[metafp.StructField]) structFunctions {

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

	type fieldName = fp.Entry[string]
	names := seq.Map(fields, func(v metafp.StructField) fieldName {
		return as.Tuple2(v.Name, v.Tag)
	})

	// tupleFuncExpr := func() string {
	// 	p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
	// 		return f.TypeName(r.w, ctx.working)
	// 	}).MakeString(",")

	// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
	// 	aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

	// 	return fmt.Sprintf(`func( v %s) %s.Tuple%d[%s] {
	// 		return %s.Tuple%d(%s)
	// 	}`, typeStr(ctx.working), fppk, fields.Size(), p,
	// 		aspk, fields.Size(), seq.Map(names, func(v fieldName) string { return "v." + v.I1 }).MakeString(","),
	// 	)
	// }

	// applyFuncExpr := func() string {
	// 	p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
	// 		return f.TypeName(r.w, ctx.working)
	// 	}).MakeString(",")

	// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

	// 	assign := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
	// 		return fmt.Sprintf("%s : t.I%d", v.I2.I1, v.I1+1)
	// 	}).MakeString(",\n")
	// 	return fmt.Sprintf(`func(t %s.Tuple%d[%s]) %s {
	// 				return %s{
	// 					%s,
	// 				}
	// 			}`, fppk, fields.Size(), p, typeStr(ctx.working),
	// 		typeStr(ctx.working),
	// 		assign,
	// 	)
	// }

	// asLabelledFuncExpr := func() string {
	// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
	// 	aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

	// 	namedTypeArgs := seq.Zip(names, typeArgs)

	// 	labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
	// 		return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
	// 	}).MakeString(",")

	// 	varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
	// 		return fmt.Sprintf("i%d", v)
	// 	}).MakeString(",")

	// 	hlistExpr := seq.Map(seq.ZipWithIndex(namedTypeArgs), func(t3 fp.Tuple2[int, fp.Tuple2[fieldName, metafp.TypeInfoExpr]]) string {
	// 		idx, t2 := t3.Unapply()
	// 		name, _ := t2.Unapply()
	// 		return fmt.Sprintf(`%s.NamedWithTag("%s",  i%d, %s)`, aspk, name.I1, idx, "`"+name.I2+"`")

	// 	}).MakeString(",")

	// 	unapplyexpr := r.structUnapplyExpr(ctx, option.None[metafp.NamedTypeInfo](), fields, "v")
	// 	return fmt.Sprintf(`func(v %s) %s.Labelled%d[%s] {
	// 						%s := %s
	// 						return %s.Labelled%d(%s)
	// 					}`, typeStr(ctx.working), fppk, fields.Size(), labelledtp,
	// 		varlist, unapplyexpr,
	// 		aspk, fields.Size(), hlistExpr)
	// }

	// fromLabelledFuncExpr := func() string {
	// 	fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
	// 	namedTypeArgs := seq.Zip(names, typeArgs)

	// 	labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[fieldName, metafp.TypeInfoExpr]) string {
	// 		return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, tp.I2.TypeName(r.w, ctx.working))
	// 	}).MakeString(",")

	// 	args := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, fieldName]) string {
	// 		return fmt.Sprintf("t.I%d.Value()", v.I1+1)
	// 	})

	// 	return fmt.Sprintf(`func( t %s.Labelled%d[%s] ) %s {
	// 			return %s
	// 		}`, fppk, fields.Size(), labelledtp, typeStr(ctx.working),
	// 		r.structApplyExpr(ctx, option.None[metafp.NamedTypeInfo](), fields, args...),
	// 	)
	// }

	sf := structFunctions{
		pack:     ctx.working.Package(),
		tpe:      tpe,
		fields:   fields,
		names:    names,
		typeArgs: typeArgs,
		// asTuple:      tupleFuncExpr,
		// fromTuple:    applyFuncExpr,
		typeStr: typeStr,
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
	valueGenerated bool
	namedGenerated bool

	names          fp.Seq[fp.Entry[string]]
	typeStr        func(pk genfp.WorkingPackage) string
	builderTypeStr func(pk genfp.WorkingPackage) string

	// asLabelled   func() string
	// fromLabelled func() string

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
func (r *TypeClassSummonContext) summonStructGenericRepr(ctx SummonContext, tc metafp.TypeClass, sf structFunctions) GenericRepr {
	fields := sf.fields
	result := r.lookupTupleTypeClassFunc(ctx, tc, fmt.Sprintf("Tuple%d", fields.Size()), sf.typeArgs)

	if result.IsDefined() {

		tci := result.Get()
		// tci.RequiredInstance = seq.Map(tci.RequiredInstance, func(v metafp.RequiredInstance) metafp.RequiredInstance {
		// 	v.FieldOf = option.Some(sf.tpe)
		// 	return v
		// })
		// tp := iterator.Map(typeArgs.Iterator(), func(v metafp.TypeInfo) string {
		// 	return r.w.TypeName(ctx.working, v.Type)
		// }).MakeString(",")
		return GenericRepr{
			Kind:         fp.GenericKindStruct,
			Type:         as.Supplier1(sf.typeStr, ctx.working),
			ReprType:     r.tupleReprType(ctx, sf, tci.Result.TypeArgs.Head()),
			ToReprExpr:   r.intoTupleRepr(ctx, sf, tci.Result.TypeArgs.Head()),
			FromReprExpr: r.fromTupleRepr(ctx, sf, tci.Result.TypeArgs.Head()),
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassInstance(ctx, tci, false)
				//return r.exprTypeClassMember(ctx, tc, result.Get(), typeArgs, option.Some(sf.tpe))
			},
		}

	}

	ret := r.summonStructHlistGenericRepr(ctx, tc, sf, "StructHCons", "StructHNil", false).
		Or(func() fp.Option[GenericRepr] {
			return r.summonStructHlistGenericRepr(ctx, tc, sf, "TupleHCons", "TupleHNil", false)
		}).
		Or(func() fp.Option[GenericRepr] {
			return r.summonStructHlistGenericRepr(ctx, tc, sf, "HCons", "HNil", true)
		})

	return ret.Get()

}

// func (r *TypeClassSummonContext) summonNamedGenericRepr(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) GenericRepr {
// 	sf := r.namedStructFuncs(ctx, named, fields)
// 	return r.summonStructGenericRepr(ctx, tc, sf)
// }

func (r *TypeClassSummonContext) summonTupleGenericRepr(ctx SummonContext, tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfoExpr], fieldOf fp.Option[metafp.TypeInfo], explicit bool) GenericRepr {
	return GenericRepr{
		Kind: fp.GenericKindTuple,
		Type: func() string {
			tuplepk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
			p := seq.Map(typeArgs, func(f metafp.TypeInfoExpr) string {
				return f.TypeName(r.w, ctx.working)
			}).MakeString(",")

			if typeArgs.Size() == 0 {
				return fmt.Sprintf(`%s.Unit`, tuplepk)
			}

			return fmt.Sprintf("%s.Tuple%d[%s]", tuplepk, typeArgs.Size(), p)
		},
		ReprType: func() string {
			hlistpk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/hlist", "hlist"))
			minimalpk := hlistpk

			if typeArgs.Size() == 0 {
				return fmt.Sprintf(`%s.Nil`, hlistpk)
			}

			hlisttp := seq.Fold(typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfoExpr) string {
				return fmt.Sprintf("%s.Cons[%s,%s]", minimalpk, a.TypeName(r.w, ctx.working), b)
			})

			return hlisttp
		},
		ToReprExpr: func() string {
			aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

			arity := fp.Min(typeArgs.Size(), max.Product-1)
			//arity := typeArgs.Size()

			// if r.implicitTypeInference {
			// 	return fmt.Sprintf(`%s.HList%d`,
			// 		aspk, arity,
			// 	)
			// }
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

			return hlistToTuple
		},
		ReprExpr: func() SummonExpr {
			//arity := fp.Min(typeArgs.Size(), max.Product-1)
			arity := typeArgs.Size()

			hcons := r.lookupTypeClassFunc(ctx, tc, "TupleHCons").OrElseGet(as.Supplier2(r.lookupHConsMust, ctx, tc))

			hnil := r.lookupTypeClassFunc(ctx, tc, "TupleHNil").OrElseGet(as.Supplier2(r.lookupHNilMust, ctx, tc))

			hlist := seq.Fold(typeArgs.Take(arity).Reverse(), newSummonExpr(func() string { return hnil.PackagedName(r.w, ctx.working) }), func(tail SummonExpr, ti metafp.TypeInfoExpr) SummonExpr {
				instance := r.summonRequired(ctx, metafp.RequiredInstance{
					TypeClass: ctx.typeClass,
					Type:      ti.Type,
				})
				return newSummonExpr(func() string {
					return fmt.Sprintf(`%s(
							%s,
							%s,
						)`, hcons.PackagedName(r.w, ctx.working), instance, tail,
					)
				}, instance.paramInstance, tail.paramInstance)
			})
			return hlist
		},
	}
}

func (r *TypeClassSummonContext) summonTuple(ctx SummonContext, tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfoExpr]) SummonExpr {

	result := r.lookupTupleTypeClassFunc(ctx, tc, fmt.Sprintf("Tuple%d", typeArgs.Size()), typeArgs)

	if result.IsDefined() {
		return r.exprTypeClassMember(ctx, tc, result.Get(), typeArgs, fp.Option[metafp.TypeInfo]{})
	}

	tupleGeneric := r.summonTupleGenericRepr(ctx, tc, typeArgs, fp.Option[metafp.TypeInfo]{}, true)
	return r.summonGeneric(ctx, tc, fmt.Sprintf("fp.Tuple%d", typeArgs.Size()), tupleGeneric)

}

func (r *TypeClassSummonContext) summonFpNamed(ctx SummonContext, tc metafp.TypeClass, typePkg *types.Package, structName string, name string, t metafp.TypeInfoExpr, genLabelled bool) SummonExpr {

	if r.fpPkg.IsDefined() {
		rtt := r.namedOrRuntimeType(ctx.working, typePkg, structName, name, t, genLabelled)
		//fmt.Printf("tc = %s, type = %s, rtt = %s, %T\n", tc, t.Type, rtt, rtt)
		named := r.namedLookup(ctx, metafp.RequiredInstance{
			TypeClass: tc,
			Type:      rtt,
		}, true, "Named")
		// named := r.lookupTypeClassFunc(ctx, tc, "Named")
		if named.IsDefined() {
			//fmt.Printf("find named\n")
			return r.exprTypeClassInstance(ctx, named.Get().instance, false)
		}
	}

	instance := r.lookupTypeClassFuncMust(ctx, tc, "Named")

	expr := r.summonRequired(ctx, metafp.RequiredInstance{
		TypeClass: tc,
		Type:      t.Type,
	})

	retExpr := func() string {
		if instance.RequiredInstance.Size() == 0 {
			return fmt.Sprintf("%s[%s]()",
				instance.PackagedName(r.w, ctx.working),
				r.namedOrRuntimeStringExpr(r.w, ctx.working, typePkg, structName, name, genLabelled, t),
			)
		}
		return fmt.Sprintf("%s[%s](%s)",
			instance.PackagedName(r.w, ctx.working),
			r.namedOrRuntimeStringExpr(r.w, ctx.working, typePkg, structName, name, genLabelled, t),
			expr,
		)
	}

	return newSummonExpr(retExpr, expr.paramInstance)

	// pk := r.w.GetImportedName(ctx.working)
	// return fmt.Sprintf("%s.Named(%s)", pk, r.summon(t))
}

func (r *TypeClassSummonContext) SummonExpression(tc metafp.TypeClassDerive) SummonExpr {

	ctx := DeriveContext{
		workingScope: r.tcCache.GetLocal(tc.Package.Package(), tc.TypeClass),
		primScope:    r.tcCache.Get(tc.PrimitiveInstancePkg, tc.TypeClass),
		tc:           tc,
		working:      tc.Package,
		recursiveGen: option.FlatMap(tc.Tags.Get("@fp.Derive"),
			fp.Compose2(metafp.Annotation.Params, as.Func2(fp.Map[string, string].Get).ApplyLast("recursive"))).Exists(eq.GivenValue("true")),
		noinline: option.FlatMap(tc.Tags.Get("@fp.Derive"),
			fp.Compose2(metafp.Annotation.Params, as.Func2(fp.Map[string, string].Get).ApplyLast("noinline"))).Exists(eq.GivenValue("true")),
	}

	return r.summonRequired(asSummonContext(ctx), metafp.RequiredInstance{
		TypeClass: tc.TypeClass,
		Type:      tc.DeriveFor.Info,
	})

}

func asSummonContext(ctx DeriveContext) SummonContext {

	return SummonContext{
		working:       ctx.working,
		typeClass:     ctx.tc.TypeClass,
		summonFor:     ctx.tc.DeriveFor.Name,
		deriveContext: option.Some(ctx),
	}
}

func (r *TypeClassSummonContext) summon(ctx SummonContext, req metafp.RequiredInstance) fp.Option[SummonExpr] {
	if req.Name && req.NameTag.IsDefined() {
		if req.NameTag.Get().IsRight() {
			return option.Some(newSummonExpr(func() string {
				name := req.NameTag.Get().Get()
				aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

				return fmt.Sprintf("%s.NameTag(`%s`,`%s`)", aspk, name.I1, name.I2)
			}, nil))
		} else {
			return option.Some(newSummonExpr(func() string {
				aspk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))
				fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

				names := seq.Map(req.NameTag.Get().Left(), func(v fp.NameTag) string {
					return fmt.Sprintf("%s.NameTag(`%s`,`%s`)", aspk, v.I1, v.I2)
				}).MakeString(",")
				return fmt.Sprintf("[]%s.Named{%s}", fppk, names)
			}, nil))
		}
	}
	t := req.Type

	if t.IsTuple() {
		// TODO: alias 된 타입에 대해, 별도의 구현이 있는 경우??
		ut := t.Unalias()
		return option.Some(r.summonTuple(ctx, req.TypeClass, seq.Map(ut.TypeArgs, func(v metafp.TypeInfo) metafp.TypeInfoExpr {
			return metafp.TypeInfoExpr{
				Type: v,
			}
		})))
	}

	result := r.lookupTypeClassInstance(ctx, req)

	if result.target.IsRight() {
		return option.Some(r.exprLookupTarget(ctx, result))
	}
	return option.None[SummonExpr]()
}

func (r *TypeClassSummonContext) summonRequired(ctx SummonContext, req metafp.RequiredInstance) SummonExpr {

	verbose("summon required %s[%s]", req.TypeClass.Name, req.Type)
	ret := r.summon(ctx, req)
	if ret.IsDefined() {
		verbose("summoned required %s[%s]", req.TypeClass.Name, req.Type)
		return ret.Get()
	}

	result := r.lookupTypeClassInstance(ctx, req)

	tc, rgen := ctx.recursiveDerive(req, result.target.Left()).Unapply()
	if rgen {
		if !r.tcCache.IsWillGenerated(tc) {
			r.recursiveGen = append(r.recursiveGen, tc)
			r.tcCache.WillGenerated(tc)
		}
		return r.exprLookupTarget(ctx, result)
	}
	return r.exprLookupTarget(ctx, result)
}

func (r *TypeClassSummonContext) summonStruct(ctx SummonContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) SummonExpr {

	sf := r.namedStructFuncs(ctx, named, fields)
	labelledExpr := r.summonTupleWithNameGenericRepr(ctx, tc, sf).
		Or(func() fp.Option[GenericRepr] {
			return r.summonLabelledGenericRepr(ctx, tc, sf)
		})
	summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
		return r.summonStructGenericRepr(ctx, tc, sf)
	})

	return r.summonGeneric(ctx, tc, named.GenericName(), summonExpr)
}

func (r *TypeClassSummonContext) summonUntypedStruct(ctx SummonContext, tc metafp.TypeClass, tpe metafp.TypeInfo, fields fp.Seq[metafp.StructField]) SummonExpr {

	sf := r.untypedStructFuncs(ctx, tpe, fields)

	labelledExpr := r.summonTupleWithNameGenericRepr(ctx, tc, sf).
		Or(func() fp.Option[GenericRepr] {
			return r.summonLabelledGenericRepr(ctx, tc, sf)
		})
	summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
		return r.summonStructGenericRepr(ctx, tc, sf)
	})

	return r.summonGeneric(ctx, tc, "struct", summonExpr)
}

func (r *TypeClassSummonContext) summonGeneric(ctx SummonContext, tc metafp.TypeClass, genericName string, genericRepr GenericRepr) SummonExpr {
	mapExpr := option.Map(r.lookupTypeClassFunc(ctx, tc, "ContraGeneric"), func(generic metafp.TypeClassInstance) SummonExpr {
		repr := genericRepr.ReprExpr()

		retExpr := func() string {
			return fmt.Sprintf(`%s(
						"%s",
						"%s",
						%s,
						%s,
					)`, generic.PackagedName(r.w, ctx.working),
				genericName,
				genericRepr.Kind,
				repr,
				genericRepr.ToReprExpr(),
			)
		}

		return newSummonExpr(retExpr, repr.paramInstance)

	}).Or(func() fp.Option[SummonExpr] {
		return option.Map(r.lookupTypeClassFunc(ctx, tc, "Generic"), func(generic metafp.TypeClassInstance) SummonExpr {
			repr := genericRepr.ReprExpr()

			retExpr := func() string {
				fppk := r.w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
				return fmt.Sprintf(`%s(
					%s.Generic[%s,%s]{
							Type: "%s",
							Kind: "%s",
							To: %s,
							From: %s,
						}, 
						%s, 
					)`, generic.PackagedName(r.w, ctx.working),
					fppk, genericRepr.Type(), genericRepr.ReprType(),
					genericName,
					genericRepr.Kind,
					genericRepr.ToReprExpr(),
					genericRepr.FromReprExpr(),
					repr,
				)
			}

			return newSummonExpr(retExpr, repr.paramInstance)
		})
	}).Or(func() fp.Option[SummonExpr] {
		return option.Map(r.lookupTypeClassFunc(ctx, tc, "IMap"), func(imapfunc metafp.TypeClassInstance) SummonExpr {

			repr := genericRepr.ReprExpr()
			retExpr := func() string {
				return fmt.Sprintf(`%s( 
						%s, 
						%s , 
						%s,
						)`,
					imapfunc.PackagedName(r.w, ctx.working), repr, genericRepr.FromReprExpr(), genericRepr.ToReprExpr())
			}

			return newSummonExpr(retExpr, repr.paramInstance)
		})
	}).Or(func() fp.Option[SummonExpr] {
		functor := r.lookupTypeClassFunc(ctx, tc, "Map")
		return option.Map(functor, func(v metafp.TypeClassInstance) SummonExpr {
			repr := genericRepr.ReprExpr()

			retExpr := func() string {
				return fmt.Sprintf(`%s( 
						%s, 
						%s,
						)`,
					v.PackagedName(r.w, ctx.working), repr, genericRepr.FromReprExpr(),
				)
			}
			return newSummonExpr(retExpr, repr.paramInstance)
		})

	}).OrElseGet(func() SummonExpr {
		contrmap := r.lookupTypeClassFuncMust(ctx, tc, "ContraMap")
		repr := genericRepr.ReprExpr()

		retExpr := func() string {
			return fmt.Sprintf(`%s( 
					%s , 
					%s,
					)`,
				contrmap.PackagedName(r.w, ctx.working), repr, genericRepr.ToReprExpr(),
			)
		}

		return newSummonExpr(retExpr, repr.paramInstance)
	})
	return mapExpr

}

func (r *TypeClassSummonContext) summonNamed(ctx SummonContext, tc metafp.TypeClass, named metafp.NamedTypeInfo) SummonExpr {

	summonExpr := GenericRepr{
		Kind: fp.GenericKindNewType,
		Type: func() string {
			valuetp := ""
			if named.Info.TypeParam.Size() > 0 {
				valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			nameWithTp := named.PackagedName(r.w, ctx.working) + valuetp

			return nameWithTp
		},
		ReprType: func() string {
			return r.w.TypeName(ctx.working, named.Underlying.Type)
		},
		ReprExpr: func() SummonExpr {
			return r.summonRequired(ctx, metafp.RequiredInstance{
				TypeClass: tc,
				Type:      named.Underlying,
			})
		},
		ToReprExpr: func() string {
			valuetp := ""
			if named.Info.TypeParam.Size() > 0 {
				valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			nameWithTp := named.PackagedName(r.w, ctx.working) + valuetp

			return fmt.Sprintf(`func(v %s) %s {
					return %s(v)
				}`, nameWithTp, r.w.TypeName(ctx.working, named.Underlying.Type), r.w.TypeName(ctx.working, named.Underlying.Type))
		},
		FromReprExpr: func() string {
			valuetp := ""
			if named.Info.TypeParam.Size() > 0 {
				valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			nameWithTp := named.PackagedName(r.w, ctx.working) + valuetp

			return fmt.Sprintf(`func(v %s) %s {
					return %s(v)
				}`, r.w.TypeName(ctx.working, named.Underlying.Type), nameWithTp, nameWithTp)
		},
	}

	return r.summonGeneric(ctx, tc, named.GenericName(), summonExpr)
}

func (r *TypeClassSummonContext) _deriveFuncExpr(tc metafp.TypeClassDerive) SummonExpr {

	verbose("derive type class func for %s", tc.DeriveFor.Name)
	workingPackage := tc.Package

	ctx := DeriveContext{
		workingScope: r.tcCache.GetLocal(tc.Package.Package(), tc.TypeClass),
		primScope:    r.tcCache.Get(tc.PrimitiveInstancePkg, tc.TypeClass),
		tc:           tc,
		working:      workingPackage,
		recursiveGen: option.FlatMap(tc.Tags.Get("@fp.Derive"),
			fp.Compose2(metafp.Annotation.Params, as.Func2(fp.Map[string, string].Get).ApplyLast("recursive"))).Exists(eq.GivenValue("true")),
		noinline: option.FlatMap(tc.Tags.Get("@fp.Derive"),
			fp.Compose2(metafp.Annotation.Params, as.Func2(fp.Map[string, string].Get).ApplyLast("noinline"))).Exists(eq.GivenValue("true")),
	}

	funcDirective := ""
	if ctx.noinline {
		funcDirective = "//go:noinline\n"
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

		return r.summonStruct(asSummonContext(ctx), tc.TypeClass, tc.DeriveFor, allFields)
	}).OrElseGet(func() SummonExpr {
		return r.summonNamed(asSummonContext(ctx), tc.TypeClass, tc.DeriveFor)
	})

	retExpr := func() string {
		if tc.DeriveFor.Info.TypeParam.Size() > 0 {
			tcname := tc.TypeClass.PackagedName(r.w, workingPackage)
			// fargs := seq.Map(v.DeriveFor.Info.TypeParam, func(p metafp.TypeParam) string {
			// 	return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
			// }).MakeString(",")

			fargs := seq.Map(mapExpr.paramInstance, as.Func3(ParamInstance.Expr).ApplyLast2(r.w, ctx.tc.Package)).MakeString(",")
			return fmt.Sprintf(`
						%sfunc %s%s( %s ) %s[%s%s] {
							return %s
						}
					`, funcDirective, tc.GeneratedInstanceName(), valuetpdec, fargs, tcname, tc.DeriveFor.PackagedName(r.w, workingPackage), valuetp,
				mapExpr)
		}
		tcname := tc.TypeClass.PackagedName(r.w, workingPackage)
		return fmt.Sprintf(`
						%sfunc %s() %s[%s] {
							return %s
						}
					`, funcDirective, tc.GeneratedInstanceName(), tcname, tc.DeriveFor.PackagedName(r.w, workingPackage),
			mapExpr)
	}

	return newSummonExpr(retExpr, mapExpr.paramInstance)

}

func (r *TypeClassSummonContext) deriveFuncExpr(tc metafp.TypeClassDerive) fp.Option[SummonExpr] {

	retOpt := r.derived.Get(tc.GeneratedInstanceName())
	if retOpt.IsDefined() {
		return option.Some(retOpt.Get().Expr)
	}

	if r.loopCheck.Contains(tc.GeneratedInstanceName()) {
		// fmt.Printf("cycle detected\n")
		return option.None[SummonExpr]()
	}

	r.loopCheck = r.loopCheck.Incl(tc.GeneratedInstanceName())

	ret := r._deriveFuncExpr(tc)
	r.derived = r.derived.Updated(tc.GeneratedInstanceName(), TypeClassInstanceGenerated{
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
		return v.Path() == ""
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

func NewTypeClassSummonContext(pkgs []*packages.Package, importSet genfp.ImportSet, fpPkg, hlistPkg *types.Package) *TypeClassSummonContext {

	// fpPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp")
	// hlistPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/hlist")

	derives := metafp.FindTypeClassDerive(pkgs)
	summons := generator.FindTaggedNotInitalizedVariable(pkgs, "@fp.Summon")

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

	initSet := seq.ToGoSet(seq.Map(summons, func(v generator.TaggedVar) string {
		return v.Name
	}))
	return &TypeClassSummonContext{
		w:                     importSet,
		fpPkg:                 option.Some(fpPkg),
		hlistPkg:              option.Some(hlistPkg),
		tcCache:               &tccache,
		recursiveGen:          derives,
		implicitTypeInference: implicitTypeInference && moduleInf,
		initVars:              summons,
		initVarSet:            initSet,
	}
}

const MaxFileSize = 2000000

//const MaxFileSize = 10

func genDerive() {

	pack := os.Getenv("GOPACKAGE")

	genfp.Generate(pack, derive_generated_file_name(pack), func(w genfp.Writer) {

		cwd, _ := os.Getwd()

		cfg := &packages.Config{
			Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
		}

		pkgs, err := packages.Load(cfg, cwd, "github.com/csgura/fp", "github.com/csgura/fp/hlist")
		if err != nil {
			fmt.Printf("package load error : %s\n", err)
			return
		}

		fpPkgs, remains := seq.Partition(pkgs, func(v *packages.Package) bool {
			return v.Types.Path() == "github.com/csgura/fp" || v.Types.Path() == "github.com/csgura/fp/hlist"
		})

		fpPkg := fpPkgs.Filter(func(v *packages.Package) bool {
			return v.Types.Path() == "github.com/csgura/fp"
		})

		hlistPkg := fpPkgs.Filter(func(v *packages.Package) bool {
			return v.Types.Path() == "github.com/csgura/fp/hlist"
		})

		// fmtalias := w.GetImportedName(genfp.NewImportPackage("fmt", "fmt"))
		// asalias := w.GetImportedName(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

		summonCtx := NewTypeClassSummonContext(remains, w, fpPkg.Head().Get().Types, hlistPkg.Head().Get().Types)

		for len(summonCtx.recursiveGen) > 0 {
			d := summonCtx.recursiveGen
			summonCtx.recursiveGen = nil
			d.Foreach(func(v metafp.TypeClassDerive) {
				summonCtx.deriveFuncExpr(v).Foreach(func(v SummonExpr) {
					fmt.Fprintf(w, "%s\n", v)
					w.CheckMaxFileSize(MaxFileSize)
				})
			})
		}

		if len(summonCtx.initVars) > 0 {
			fmt.Fprint(w, `func init() {
			`)
			for _, v := range summonCtx.initVars {
				rq, ok := metafp.AsRequiredInstance(metafp.GetTypeInfo(v.Type)).Unapply()
				if ok {

					ctx := SummonContext{
						working:   v.Package,
						typeClass: rq.TypeClass,
						summonFor: rq.Type.String(),
					}

					expr := summonCtx.summonRequired(ctx, rq)
					fmt.Fprintf(w, "%s = %s\n", v.Name, expr)
				}
			}
			fmt.Fprint(w, `}`)
		}

	})
}
