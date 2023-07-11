package main

import (
	"fmt"
	"go/types"
	"os"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type TypeClassInstanceGenerated struct {
	Derive metafp.TypeClassDerive
	Expr   SummonExpr
}

type TypeClassSummonContext struct {
	w         genfp.Writer
	tcCache   *metafp.TypeClassInstanceCache
	summoned  fp.Map[string, TypeClassInstanceGenerated]
	loopCheck fp.Set[string]
}

type CurrentContext struct {
	working      *types.Package
	tc           metafp.TypeClassDerive
	primScope    metafp.TypeClassScope
	workingScope metafp.TypeClassScope
}

type GenericRepr struct {
	//	ReprType     func() string
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

func (r ParamInstance) Expr(importSet genfp.ImportSet, working *types.Package) string {
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

func collectSummonExpr(list fp.Seq[SummonExpr]) SummonExpr {
	expr := seq.Map(list, SummonExpr.Expr).MakeString(",")
	paramList := seq.Reduce(seq.Map(list, SummonExpr.ParamInstance), MergeSeqDistinct(EqParamInstance))
	return SummonExpr{
		expr:          expr,
		paramInstance: paramList,
	}
}

type lookupTarget struct {
	instanceOf metafp.TypeInfo
	pk         *types.Package
	name       string
	required   fp.Seq[metafp.RequiredInstance]
	typeParam  fp.Option[metafp.TypeClass]
	instance   fp.Option[metafp.TypeClassInstance]
	// tc       *TypeClass

}

func (r lookupTarget) isFunc() bool {
	if r.instance.IsDefined() {
		return !r.instance.Get().Static
	}
	return false
}

func (r lookupTarget) instanceExpr(w genfp.Writer, workingPkg *types.Package) SummonExpr {

	param := option.Map(r.typeParam, func(v metafp.TypeClass) ParamInstance {
		return ParamInstance{
			ArgName:   r.name,
			TypeClass: v,
			ParamName: r.instanceOf.Name().Get(),
		}
	}).ToSeq()

	if r.pk == nil || r.pk.Path() == workingPkg.Path() {
		return SummonExpr{
			expr:          r.name,
			paramInstance: param,
		}
	}

	pk := w.GetImportedName(r.pk)

	return SummonExpr{
		expr:          fmt.Sprintf("%s.%s", pk, r.name),
		paramInstance: param,
	}

}

func (r *TypeClassSummonContext) typeclassInstanceMust(ctx CurrentContext, req metafp.RequiredInstance, name string) lookupTarget {

	f := req.Type
	return lookupTarget{
		instanceOf: f,
		pk:         ctx.working,
		name:       req.TypeClass.Name + publicName(name),
		required: seq.Map(f.TypeArgs, func(v metafp.TypeInfo) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: req.TypeClass,
				Type:      v,
			}
		}),
	}
}

// f 는 Eq 쌓이지 않은 타입
// Eq[T] 같은거 아님
func (r *TypeClassSummonContext) lookupTypeClassInstanceLocalDeclared(ctx CurrentContext, req metafp.RequiredInstance, name ...string) fp.Option[lookupTarget] {

	f := req.Type

	scope := ctx.workingScope
	if req.TypeClass.Id() != ctx.tc.TypeClass.Id() {
		scope = r.tcCache.GetLocal(ctx.working, req.TypeClass)
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
		return option.Iterator(scope.FindByName(v, f))
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
	ins = ins.Filter(func(tci metafp.TypeClassInstance) bool {

		return r.checkRequired(ctx, tci.RequiredInstance)
	})

	if f.TypeArgs.Size() > 0 {
		ins = seq.Iterator(scope.Find(f)).Concat(ins)
	} else {
		ins = ins.Concat(seq.Iterator(scope.Find(f)))
	}

	return iterator.Map(ins, func(v metafp.TypeClassInstance) lookupTarget {
		return lookupTarget{
			instanceOf: f,
			pk:         v.Package,
			name:       v.Name,
			instance:   option.Some(v),

			// 함수의 아규먼트는 Eq 가 포함 되어 있음.
			required: v.RequiredInstance,
		}

	}).NextOption()

}

func (r *TypeClassSummonContext) lookupHNilMust(ctx CurrentContext, tc metafp.TypeClass) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(ctx, tc, "HNil")
	if ret.IsDefined() {
		return ret.Get()
	}

	ret = r.lookupTypeClassFunc(ctx, tc, "HlistNil")
	if ret.IsDefined() {
		return ret.Get()
	}
	nameWithTc := tc.Name + "HNil"

	return metafp.TypeClassInstance{
		Package: ctx.working,
		Name:    nameWithTc,
		Static:  true,
	}
}

func (r *TypeClassSummonContext) lookupTypeClassFunc(ctx CurrentContext, tc metafp.TypeClass, name string) fp.Option[metafp.TypeClassInstance] {
	nameWithTc := tc.Name + name

	workingScope := ctx.workingScope
	primScope := ctx.primScope
	if ctx.tc.TypeClass.Id() != tc.Id() {
		primScope = r.tcCache.GetImported(tc)
		workingScope = r.tcCache.GetLocal(ctx.working, tc)
	}

	ins := workingScope.FindFunc(nameWithTc)
	if ins.IsDefined() {
		return ins
	}

	ins = primScope.FindFunc(nameWithTc)
	if ins.IsDefined() {
		return ins
	}

	ins = primScope.FindFunc(name)
	return ins
}

func (r *TypeClassSummonContext) lookupTypeClassFuncMust(ctx CurrentContext, tc metafp.TypeClass, name string) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(ctx, tc, name)
	if ret.IsDefined() {
		return ret.Get()
	}

	nameWithTc := tc.Name + name

	return metafp.TypeClassInstance{
		Package: ctx.working,
		Name:    nameWithTc,
		Static:  true,
	}
}

func (r *TypeClassSummonContext) lookupTypeClassInstancePrimitivePkgLazy(ctx CurrentContext, req metafp.RequiredInstance, name ...string) func() fp.Option[lookupTarget] {
	return func() fp.Option[lookupTarget] {
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
			res := r.lookupTypeClassInstance(ctx, v)
			if res.available.IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (r *TypeClassSummonContext) lookupTypeClassInstancePrimitivePkg(ctx CurrentContext, req metafp.RequiredInstance, name ...string) fp.Option[lookupTarget] {

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
	}).Concat(seq.Iterator(scope.Find(f)))

	if f.TypeArgs.Size() > 0 {
		ins = seq.Iterator(scope.Find(f)).Concat(ins)
	} else {
		ins = ins.Concat(seq.Iterator(scope.Find(f)))
	}

	ins = ins.Filter(func(tci metafp.TypeClassInstance) bool {
		return r.checkRequired(ctx, tci.RequiredInstance)
	})

	return iterator.Map(ins, func(v metafp.TypeClassInstance) lookupTarget {
		return lookupTarget{
			instanceOf: f,
			pk:         v.Package,
			name:       v.Name,
			required:   v.RequiredInstance,
			instance:   option.Some(v),
		}
	}).NextOption()

}

func (r *TypeClassSummonContext) lookupTypeClassInstanceTypePkg(ctx CurrentContext, req metafp.RequiredInstance, name string) fp.Option[lookupTarget] {

	f := req.Type
	if f.Pkg != nil && f.Pkg.Path() != ctx.working.Path() {

		name := req.TypeClass.Name + publicName(name)
		obj := f.Pkg.Scope().Lookup(name)

		if obj != nil {
			ret := lookupTarget{
				instanceOf: f,
				pk:         f.Pkg,
				name:       name,
				required: seq.Map(f.TypeArgs, func(v metafp.TypeInfo) metafp.RequiredInstance {
					return metafp.RequiredInstance{
						TypeClass: req.TypeClass,
						Type:      v,
					}
				}),
				instance: metafp.AsTypeClassInstance(req.TypeClass, obj),
			}

			return option.Some(ret)

		}
	}

	return option.None[lookupTarget]()
}

func (r *TypeClassSummonContext) namedLookup(ctx CurrentContext, req metafp.RequiredInstance, name string) typeClassInstance {
	ret := r.lookupTypeClassInstanceLocalDeclared(ctx, req, name).Or(lazy.Func3(r.lookupTypeClassInstanceTypePkg)(ctx, req, name).Get).Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, name))

	return typeClassInstance{
		ret,
		r.typeclassInstanceMust(ctx, req, name),
	}
}

func (r *TypeClassSummonContext) lookupPrimitiveTypeClassInstance(ctx CurrentContext, req metafp.RequiredInstance, name ...string) typeClassInstance {
	ret := r.lookupTypeClassInstanceLocalDeclared(ctx, req, name...).Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, name...))

	return typeClassInstance{
		ret,
		r.typeclassInstanceMust(ctx, req, name[0]),
	}
}

// 타입 추론이 가능한지 따지는 함수
func (r *TypeClassSummonContext) typeParamString(ctx CurrentContext, lt lookupTarget) fp.Option[string] {

	if lt.instance.IsDefined() {
		ins := lt.instance.Get()

		// 타입 추론이 가능하려면,  모든 타입 파라미터가, 아규먼트에서 사용되어야 한다.
		possible := ins.TypeParam.ForAll(func(v metafp.TypeParam) bool {
			return ins.UsedParam.Contains(v.Name)
		})

		// 전부 사용되지 않아 타입 추론이 불가능하다면
		// 타입을 명시한다.
		if !possible {
			ret := seq.Map(ins.TypeParam, func(v metafp.TypeParam) string {
				return option.Map(ins.ParamMapping.Get(v.Name), func(v metafp.TypeInfo) string {
					return r.w.TypeName(ctx.working, v.Type)
				}).OrElse(v.Name)
			}).MakeString(",")
			return option.Some(ret)
		}

	}

	return option.None[string]()
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
		ret := r.summon(ctx, t)
		if t.Lazy {
			lazypk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/lazy", "lazy"))
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

func (r *TypeClassSummonContext) exprTypeClassInstance(ctx CurrentContext, lt lookupTarget) SummonExpr {
	if len(lt.required) > 0 {
		list := r.summonArgs(ctx, lt.required)

		instanceExpr := lt.instanceExpr(r.w, ctx.working)
		tpstr := r.typeParamString(ctx, lt)
		if tpstr.IsDefined() {
			//fmt.Printf("%s param infer not possible = %s \n", lt.name, lt.instance.Get().ParamMapping)

			return newSummonExpr(fmt.Sprintf("%s[%s](%s)", instanceExpr.expr, tpstr.Get(), list.expr), instanceExpr.paramInstance, list.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s(%s)", instanceExpr.expr, list.expr), instanceExpr.paramInstance, list.paramInstance)

		}
	}

	if lt.isFunc() && len(lt.required) == 0 {
		instanceExpr := lt.instanceExpr(r.w, ctx.working)

		tpstr := r.typeParamString(ctx, lt)
		if tpstr.IsDefined() {
			return newSummonExpr(fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get()), instanceExpr.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s()", instanceExpr), instanceExpr.paramInstance)
		}

	}

	return lt.instanceExpr(r.w, ctx.working)

}

func (r *TypeClassSummonContext) exprTypeClassMember(ctx CurrentContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typeArgs fp.Seq[metafp.TypeInfo]) SummonExpr {
	if len(typeArgs) > 0 {
		list := r.summonArgs(ctx, seq.Map(typeArgs, func(t metafp.TypeInfo) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: tc,
				Type:      t,
			}
		}))

		return newSummonExpr(fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, ctx.working), list), list.paramInstance)
	}

	return newSummonExpr(lt.PackagedName(r.w, ctx.working))

}

func (r *TypeClassSummonContext) exprTypeClassMemberLabelled(ctx CurrentContext, tc metafp.TypeClass, lt metafp.TypeClassInstance, typePkg *types.Package, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfo], genLabelled bool) SummonExpr {
	if len(typeArgs) > 0 {
		list := collectSummonExpr(seq.Map(seq.Zip(typeArgs, names), func(t fp.Tuple2[metafp.TypeInfo, string]) SummonExpr {
			return r.summonFpNamed(ctx, tc, typePkg, t.I2, t.I1, genLabelled)
		}))

		return newSummonExpr(fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, ctx.working), list), list.paramInstance)
	}

	return newSummonExpr(lt.PackagedName(r.w, ctx.working))

}

func (r *TypeClassSummonContext) lookupTypeClassInstance(ctx CurrentContext, req metafp.RequiredInstance) typeClassInstance {
	f := req.Type

	switch at := f.Type.(type) {
	case *types.TypeParam:
		return newTypeClassInstance(lookupTarget{
			instanceOf: f,
			name:       privateName(req.TypeClass.Name) + at.Obj().Name(),
			typeParam:  option.Some(req.TypeClass),
		})
	case *types.Named:
		if at.Obj().Pkg().Path() == "github.com/csgura/fp/hlist" {
			if at.Obj().Name() == "Nil" {
				return typeClassInstance{r.lookupTypeClassInstanceLocalDeclared(ctx, req, "HNil", "HListNil").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, "HNil", "HListNil")),
					r.typeclassInstanceMust(ctx, req, "HNil"),
				}

			} else if at.Obj().Name() == "Cons" {
				return typeClassInstance{
					r.lookupTypeClassInstanceLocalDeclared(ctx, req, "HCons", "HListCons").
						Or(r.lookupTypeClassInstancePrimitivePkgLazy(ctx, req, "HCons", "HListCons")),

					r.typeclassInstanceMust(ctx, req, "HCons"),
				}
			}
		}
		return r.namedLookup(ctx, req, at.Obj().Name())
	case *types.Array:
		panic(fmt.Sprintf("can't summon array type, while deriving %s[%s]", req.TypeClass.Name, ctx.tc.DeriveFor.Name))
		//return r.namedLookup(f, "Array")
	case *types.Slice:
		if at.Elem().String() == "byte" {
			bytesInstance := r.namedLookup(ctx,
				metafp.RequiredInstance{
					TypeClass: req.TypeClass,
					Type: metafp.TypeInfo{
						Pkg:      f.Pkg,
						Type:     f.Type,
						TypeArgs: nil,
					}}, "Bytes")

			if bytesInstance.available.IsDefined() {
				return bytesInstance
			}
			return r.namedLookup(ctx, req, "Slice")
		}
		return r.namedLookup(ctx, req, "Slice")
	case *types.Map:
		return r.namedLookup(ctx, req, "GoMap")
	case *types.Pointer:
		return r.namedLookup(ctx, req, "Ptr")
	case *types.Basic:
		return r.namedLookup(ctx, req, at.Name())
	case *types.Struct:
		panic(fmt.Sprintf("can't summon unnamed struct type, while deriving %s[%s]", ctx.tc.TypeClass.Name, ctx.tc.DeriveFor.Name))
	case *types.Interface:
		if f.IsAny() {
			return r.namedLookup(ctx, req, "Given")
		}
		panic(fmt.Sprintf("can't summon unnamed interface type %v, while deriving %s[%s]", f.Type, ctx.tc.TypeClass.Name, ctx.tc.DeriveFor.Name))
	case *types.Chan:
		panic(fmt.Sprintf("can't summon unnamed chan type, while deriving %s[%s]", ctx.tc.TypeClass.Name, ctx.tc.DeriveFor.Name))

	}
	return r.namedLookup(ctx, req, f.Type.String())
}

func (r *TypeClassSummonContext) structUnapplyExpr(ctx CurrentContext, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField], varexpr string) (string, fp.Seq[metafp.StructField]) {
	hasUnapply := named.Info.Method.Contains("Unapply")
	if hasUnapply {
		return fmt.Sprintf("%s.Unapply()", varexpr), fields
	}

	fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})

	return fmt.Sprintf(`%s`, seq.Map(names, func(v string) string { return fmt.Sprintf("%s.%s", varexpr, v) }).MakeString(",")), fields
}

func (r *TypeClassSummonContext) structApplyExpr(ctx CurrentContext, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField], args ...string) string {
	hasApply := named.Info.Method.Contains("Builder")
	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	if hasApply {

		builderreceiver := fmt.Sprintf("%sBuilder%s", named.PackagedName(r.w, ctx.working), valuetp)

		return fmt.Sprintf(`%s{}.Apply(%s).Build()`,
			builderreceiver, as.Seq(args).MakeString(","))
	}

	fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})
	argslist := seq.Map(seq.Zip(names, args), func(v fp.Tuple2[string, string]) string {
		return fmt.Sprintf("%s: %s", v.I1, v.I2)
	}).MakeString(",")

	valuereceiver := fmt.Sprintf("%s%s", named.PackagedName(r.w, ctx.working), valuetp)

	return fmt.Sprintf(`%s{%s}`, valuereceiver, argslist)
}

func namedOrRuntime(w genfp.Writer, working *types.Package, typePkg *types.Package, name string, labelledGen bool) string {

	if labelledGen {
		ret := publicName(name)
		if ret == name {
			ret = fmt.Sprintf("PubNamed%s", ret)
		} else {
			ret = fmt.Sprintf("Named%s", ret)
		}

		if isSamePkg(working, typePkg) {
			return ret
		} else {
			return fmt.Sprintf("%s.%s", w.GetImportedName(typePkg), ret)
		}
	} else {
		fppk := w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

		return fmt.Sprintf("%s.RuntimeNamed", fppk)

	}

}

func (r *TypeClassSummonContext) summonLabelledGenericRepr(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) fp.Option[GenericRepr] {

	hasUnapply := named.Info.Method.Contains("Unapply")
	hasAsLabelled := named.Info.Method.Contains("AsLabelled")

	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	builderreceiver := fmt.Sprintf("%sBuilder%s", named.PackagedName(r.w, ctx.working), valuetp)
	valuereceiver := fmt.Sprintf("%s%s", named.PackagedName(r.w, ctx.working), valuetp)

	if !hasUnapply && !hasAsLabelled {
		fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	}

	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfo {
		return v.Type
	})

	asLabelledExpr := func() string {
		return fmt.Sprintf("%s.AsLabelled", valuereceiver)
	}

	if !hasAsLabelled {
		asLabelledExpr = func() string {
			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			namedTypeArgs := seq.Zip(names, typeArgs)

			labelledtp := seq.Map(namedTypeArgs, func(tp fp.Tuple2[string, metafp.TypeInfo]) string {
				return fmt.Sprintf("%s.RuntimeNamed[%s]", fppk, r.w.TypeName(ctx.working, tp.I2.Type))
			}).MakeString(",")

			varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
				return fmt.Sprintf("i%d", v)
			}).MakeString(",")

			hlistExpr := seq.Map(seq.ZipWithIndex(namedTypeArgs), func(t3 fp.Tuple2[int, fp.Tuple2[string, metafp.TypeInfo]]) string {
				idx, t2 := t3.Unapply()
				name, tp := t2.Unapply()
				return fmt.Sprintf(`%s.RuntimeNamed[%s]{ I1 : "%s", I2 : i%d }`, fppk, r.w.TypeName(ctx.working, tp.Type), name, idx)

			}).MakeString(",")

			unapplyexpr, _ := r.structUnapplyExpr(ctx, named, fields, "v")
			return fmt.Sprintf(`func(v %s) %s.Labelled%d[%s] {
							%s := %s
							return %s.Labelled%d(%s)
						}`, valuereceiver, fppk, fields.Size(), labelledtp,
				varlist, unapplyexpr,
				aspk, fields.Size(), hlistExpr)
		}

	}

	result := r.lookupTypeClassFunc(ctx, tc, fmt.Sprintf("Labelled%d", typeArgs.Size()))

	return option.Map(result, func(tm metafp.TypeClassInstance) GenericRepr {
		return GenericRepr{
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr: asLabelledExpr,
			FromReprExpr: func() string {
				fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

				return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromLabelled)(%s{}),
					%s.Build,
					)`,
					fppk,
					aspk, builderreceiver, builderreceiver, builderreceiver,
				)
			},
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassMemberLabelled(ctx, tc, tm, named.Package, names, typeArgs, hasAsLabelled)
			},
		}
	}).Or(func() fp.Option[GenericRepr] {
		return option.Map(r.lookupTypeClassFunc(ctx, tc, "HConsLabelled"), func(hcons metafp.TypeClassInstance) GenericRepr {
			return GenericRepr{
				// ReprType: func() string {
				// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
				// },
				ToReprExpr: func() string {

					if typeArgs.Size() == 0 {
						hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

						return fmt.Sprintf(`func (%s) %s.Nil {
							return %s.Empty()
						}`, valuereceiver, hlistpk, hlistpk)
					} else if typeArgs.Size() < max.Product {

						arity := fp.Min(typeArgs.Size(), max.Product-1)
						//arity := typeArgs.Size()

						fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
						aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
							return fmt.Sprintf("%s[%s]", namedOrRuntime(r.w, ctx.working, named.Package, f.I1, hasAsLabelled), r.w.TypeName(ctx.working, f.I2.Type))
						}).Take(arity).MakeString(",")

						return fmt.Sprintf(`%s.Compose(
						%s,
						%s.HList%dLabelled[%s],
					)`, fppk,
							asLabelledExpr(),
							aspk, arity, tp,
						)
					} else {
						hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						hlisttp := seq.Fold(namedTypeArgs.Reverse(), hlistpk+".Nil", func(b string, f fp.Tuple2[string, metafp.TypeInfo]) string {
							name, a := f.Unapply()
							return fmt.Sprintf("%s.Cons[%s[%s],%s]", hlistpk, namedOrRuntime(r.w, ctx.working, named.Package, name, hasAsLabelled), r.w.TypeName(ctx.working, a.Type), b)
						})

						varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
							return fmt.Sprintf("i%d", v)
						}).MakeString(",")

						hlistExpr := option.Map(option.Of(hasAsLabelled).Filter(eq.GivenValue(true)), func(bool) string {
							return seq.Fold(seq.ZipWithIndex(namedTypeArgs).Reverse(), hlistpk+".Empty()", func(expr string, t3 fp.Tuple2[int, fp.Tuple2[string, metafp.TypeInfo]]) string {
								idx, t2 := t3.Unapply()
								name, tp := t2.Unapply()
								return fmt.Sprintf(`%s.Concat(%s[%s]{i%d}, 
						%s,
					)`, hlistpk, namedOrRuntime(r.w, ctx.working, named.Package, name, hasAsLabelled), r.w.TypeName(ctx.working, tp.Type), idx, expr)
							})
						}).OrElseGet(func() string {
							return seq.Fold(seq.ZipWithIndex(namedTypeArgs).Reverse(), hlistpk+".Empty()", func(expr string, t3 fp.Tuple2[int, fp.Tuple2[string, metafp.TypeInfo]]) string {
								idx, t2 := t3.Unapply()
								name, tp := t2.Unapply()
								return fmt.Sprintf(`%s.Concat(%s[%s]{I1: "%s", I2 : i%d}, 
						%s,
					)`, hlistpk, namedOrRuntime(r.w, ctx.working, named.Package, name, hasAsLabelled), r.w.TypeName(ctx.working, tp.Type), name, idx, expr)
							})
						})
						//hlistExpr :=

						unapplyexpr, _ := r.structUnapplyExpr(ctx, named, fields, "v")
						return fmt.Sprintf(`func(v %s) %s {
							%s := %s
							return %s
						}`, valuereceiver, hlisttp,
							varlist, unapplyexpr,
							hlistExpr)
					}

				},
				FromReprExpr: func() string {
					if typeArgs.Size() == 0 {
						hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

						return fmt.Sprintf(`func (%s.Nil) %s{
							return %s{}
						}`, hlistpk, valuereceiver, valuereceiver)
					} else if typeArgs.Size() < max.Product {
						arity := fp.Min(typeArgs.Size(), max.Product-1)
						//arity := typeArgs.Size()

						fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
						aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
						productpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/product", "product"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
							return fmt.Sprintf("%s[%s]", namedOrRuntime(r.w, ctx.working, named.Package, f.I1, hasAsLabelled), r.w.TypeName(ctx.working, f.I2.Type))
						}).Take(arity).MakeString(",")

						hlistToTuple := fmt.Sprintf(`%s.LabelledFromHList%d[%s]`,
							productpk,
							arity, tp,
						)

						tupleToStruct := fmt.Sprintf(`%s.Compose(
						%s.Curried2(%s.FromLabelled)(%s{}),
						%s.Build,
						)`,
							fppk,
							aspk, builderreceiver, builderreceiver, builderreceiver,
						)
						return fmt.Sprintf(`
						%s.Compose(
							%s, 
							%s ,
						)`, fppk, hlistToTuple, tupleToStruct)
					} else {
						hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						hlisttp := seq.Fold(namedTypeArgs.Reverse(), hlistpk+".Nil", func(b string, t2 fp.Tuple2[string, metafp.TypeInfo]) string {
							name, a := t2.Unapply()
							return fmt.Sprintf("%s.Cons[%s[%s],%s]", hlistpk, namedOrRuntime(r.w, ctx.working, named.Package, name, hasAsLabelled), r.w.TypeName(ctx.working, a.Type), b)
						})

						expr := seq.Map(iterator.Range(0, typeArgs.Size()).ToSeq(), func(idx int) string {
							if idx == typeArgs.Size()-1 {
								return fmt.Sprintf(`i%d := hl%d.Head()`, idx, idx)
							}
							return fmt.Sprintf(`i%d , hl%d := %s.Unapply(hl%d)`, idx, idx+1, hlistpk, idx)
						}).MakeString("\n")

						arglist := seq.Map(iterator.Range(0, typeArgs.Size()).ToSeq(), func(idx int) string {
							return fmt.Sprintf("i%d.Value()", idx)
						}).MakeString(",")
						return fmt.Sprintf(`func(hl0 %s) %s {
					%s
					return %s{}.Apply(%s).Build()
				}`, hlisttp, valuereceiver,
							expr,
							builderreceiver, arglist)
					}
				},
				ReprExpr: func() SummonExpr {
					//arity := fp.Min(typeArgs.Size(), max.Product-1)
					arity := typeArgs.Size()

					hnil := r.lookupHNilMust(ctx, tc)
					namedTypeArgs := seq.Zip(names, typeArgs)
					hlist := seq.Fold(namedTypeArgs.Take(arity).Reverse(), newSummonExpr(hnil.PackagedName(r.w, ctx.working)), func(tail SummonExpr, ti fp.Tuple2[string, metafp.TypeInfo]) SummonExpr {
						instance := r.summonFpNamed(ctx, tc, named.Package, ti.I1, ti.I2, hasAsLabelled)
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

func (r *TypeClassSummonContext) summonGenericRepr(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) GenericRepr {
	hasUnapply := named.Info.Method.Contains("Unapply")

	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(seq.Iterator(named.Info.TypeParam), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	builderreceiver := fmt.Sprintf("%sBuilder%s", named.PackagedName(r.w, ctx.working), valuetp)
	valuereceiver := fmt.Sprintf("%s%s", named.PackagedName(r.w, ctx.working), valuetp)

	if !hasUnapply {
		fields = fields.Filter(func(v metafp.StructField) bool { return v.Public() })
	}

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfo {
		return v.Type
	})

	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})

	tupleExpr := func() string {
		return fmt.Sprintf("%s.AsTuple", valuereceiver)
	}
	applyExpr := func() string {
		fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
		aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

		return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromTuple)(%s{}),
					%s.Build,
					)`,
			fppk,
			aspk, builderreceiver, builderreceiver, builderreceiver,
		)
	}
	if !hasUnapply {

		tupleExpr = func() string {
			p := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(ctx.working, f.Type)
			}).MakeString(",")

			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			return fmt.Sprintf(`func( v %s) %s.Tuple%d[%s] {
			return %s.Tuple%d(%s)
		}`, valuereceiver, fppk, fields.Size(), p,
				aspk, fields.Size(), seq.Map(names, func(v string) string { return "v." + v }).MakeString(","),
			)
		}

		applyExpr = func() string {
			p := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(ctx.working, f.Type)
			}).MakeString(",")

			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
			//aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			assign := seq.Map(seq.ZipWithIndex(names), func(v fp.Tuple2[int, string]) string {
				return fmt.Sprintf("%s : t.I%d", v.I2, v.I1+1)
			}).MakeString(",\n")
			return fmt.Sprintf(`func(t %s.Tuple%d[%s]) %s {
					return %s{
						%s,
					}
				}`, fppk, fields.Size(), p, valuereceiver,
				valuereceiver,
				assign,
			)
		}

	}

	result := r.lookupTypeClassFunc(ctx, tc, fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {

		// tp := iterator.Map(typeArgs.Iterator(), func(v metafp.TypeInfo) string {
		// 	return r.w.TypeName(ctx.working, v.Type)
		// }).MakeString(",")
		return GenericRepr{
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr:   tupleExpr,
			FromReprExpr: applyExpr,
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassMember(ctx, tc, result.Get(), typeArgs)
			},
		}
	}

	tupleGeneric := r.summonTupleGenericRepr(ctx, tc, typeArgs)

	return GenericRepr{
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {

			if typeArgs.Size() >= max.Product {
				hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

				hlisttp := seq.Fold(typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfo) string {
					return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, r.w.TypeName(ctx.working, a.Type), b)
				})

				varlist := iterator.Map(iterator.Range(0, typeArgs.Size()), func(v int) string {
					return fmt.Sprintf("i%d", v)
				}).MakeString(",")

				hlistExpr := seq.Fold(as.Seq(iterator.Range(0, typeArgs.Size()).ToSeq()).Reverse(), hlistpk+".Empty()", func(expr string, v int) string {
					return fmt.Sprintf(`%s.Concat(i%d, 
						%s,
					)`, hlistpk, v, expr)
				})
				unapplyExpr, _ := r.structUnapplyExpr(ctx, named, fields, "v")
				return fmt.Sprintf(`func(v %s) %s {
					%s := %s
					return %s
				}`, valuereceiver, hlisttp,
					varlist, unapplyExpr,
					hlistExpr)
			} else if typeArgs.Size() > 0 {
				fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

				return fmt.Sprintf(`%s.Compose(
				%s,
				%s, 
			)`, fppk,
					tupleExpr(),
					tupleGeneric.ToReprExpr(),
				)
			} else {
				hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))
				return fmt.Sprintf(`func(%s) %s.Nil {
					return %s.Empty()
				}`, valuereceiver, hlistpk, hlistpk)
			}

		},
		FromReprExpr: func() string {
			if typeArgs.Size() >= max.Product {
				hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

				hlisttp := seq.Fold(typeArgs.Reverse(), hlistpk+".Nil", func(b string, a metafp.TypeInfo) string {
					return fmt.Sprintf("%s.Cons[%s,%s]", hlistpk, r.w.TypeName(ctx.working, a.Type), b)
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
				}`, hlisttp, valuereceiver,
					expr,
					r.structApplyExpr(ctx, named, fields, arglist...))
			} else if typeArgs.Size() > 0 {

				fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				//aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

				tupleToStruct := applyExpr()
				return fmt.Sprintf(`
				%s.Compose(
					%s, 
					%s ,
				)`, fppk, tupleGeneric.FromReprExpr(), tupleToStruct)
			} else {
				hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

				return fmt.Sprintf(`func(%s.Nil) %s {
					return %s{}
				}`, hlistpk, valuereceiver, valuereceiver)
			}
		},
		ReprExpr: func() SummonExpr {
			return tupleGeneric.ReprExpr()
		},
	}
}

func (r *TypeClassSummonContext) summonTupleGenericRepr(ctx CurrentContext, tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfo]) GenericRepr {
	return GenericRepr{
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			arity := fp.Min(typeArgs.Size(), max.Product-1)
			//arity := typeArgs.Size()

			tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(ctx.working, f.Type)
			}).Take(arity).MakeString(",")

			return fmt.Sprintf(`%s.HList%d[%s]`,
				aspk, arity, tp,
			)

		},
		FromReprExpr: func() string {
			productpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/product", "product"))

			arity := fp.Min(typeArgs.Size(), max.Product-1)
			//arity := typeArgs.Size()

			tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(ctx.working, f.Type)
			}).Take(arity).MakeString(",")

			hlistToTuple := fmt.Sprintf(`%s.TupleFromHList%d[%s]`,
				productpk, arity, tp,
			)

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

			hcons := r.lookupTypeClassFuncMust(ctx, tc, "HCons")

			hnil := r.lookupHNilMust(ctx, tc)

			hlist := seq.Fold(typeArgs.Take(arity).Reverse(), newSummonExpr(hnil.PackagedName(r.w, ctx.working)), func(tail SummonExpr, ti metafp.TypeInfo) SummonExpr {
				instance := r.summon(ctx, metafp.RequiredInstance{
					TypeClass: ctx.tc.TypeClass,
					Type:      ti,
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

func (r *TypeClassSummonContext) summonTuple(ctx CurrentContext, tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfo]) SummonExpr {

	result := r.lookupTypeClassFunc(ctx, tc, fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {
		return r.exprTypeClassMember(ctx, tc, result.Get(), typeArgs)
	}

	tupleGeneric := r.summonTupleGenericRepr(ctx, tc, typeArgs)
	return r.summonVariant(ctx, tc, "", tupleGeneric)

}

func (r *TypeClassSummonContext) summonFpNamed(ctx CurrentContext, tc metafp.TypeClass, typePkg *types.Package, name string, t metafp.TypeInfo, genLabelled bool) SummonExpr {

	instance := r.lookupTypeClassFuncMust(ctx, tc, "Named")

	expr := r.summon(ctx, metafp.RequiredInstance{
		TypeClass: tc,
		Type:      t,
	})

	return newSummonExpr(fmt.Sprintf("%s[%s[%s]](%s)", instance.PackagedName(r.w, ctx.working), namedOrRuntime(r.w, ctx.working, typePkg, name, genLabelled),
		r.w.TypeName(ctx.working, t.Type), expr.expr), expr.paramInstance)

	// pk := r.w.GetImportedName(ctx.working)
	// return fmt.Sprintf("%s.Named(%s)", pk, r.summon(t))
}

func (r *TypeClassSummonContext) summon(ctx CurrentContext, req metafp.RequiredInstance) SummonExpr {

	t := req.Type

	// if req.TypeClass.IsLazy() {
	// 	expr := r.summon(req.Type.TypeArgs.Head().Get())
	// }

	if t.IsTuple() {
		return r.summonTuple(ctx, req.TypeClass, t.TypeArgs)
	}

	result := r.lookupTypeClassInstance(ctx, req)

	if result.available.IsDefined() {
		return r.exprTypeClassInstance(ctx, result.available.Get())
	}

	// instance := r.lookupTypeClassMember("UInt")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Int")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Float")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Number")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Given")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, ctx.working), r.w.TypeName(ctx.working, t.Type))
	// 		}
	// 	}
	// }

	return r.exprTypeClassInstance(ctx, result.must)

}

func (r *TypeClassSummonContext) summonStruct(ctx CurrentContext, tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) SummonExpr {

	//fmt.Printf("named %s's method list = %v\n", named.Name, named.Info.Method.Keys().MakeString(","))

	// fmt.Printf("lookup %s.Option = %v\n", v.Generator.Name(), l)
	//fmt.Printf("derive %v for %v\n", v.TypeClass, v.DeriveFor)
	//privateFields := fields.FilterNot(metafp.StructField.Public)

	labelledExpr := r.summonLabelledGenericRepr(ctx, tc, named, fields)
	summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
		return r.summonGenericRepr(ctx, tc, named, fields)
	})

	return r.summonVariant(ctx, tc, named.GenericName(), summonExpr)
}

func (r *TypeClassSummonContext) summonVariant(ctx CurrentContext, tc metafp.TypeClass, genericName string, genericRepr GenericRepr) SummonExpr {
	mapExpr := option.Map(r.lookupTypeClassFunc(ctx, tc, "Generic"), func(generic metafp.TypeClassInstance) SummonExpr {

		aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
		repr := genericRepr.ReprExpr()
		return newSummonExpr(fmt.Sprintf(`%s(
					%s.Generic(
							"%s",
							%s,
							%s,
						), 
						%s, 
					)`, generic.PackagedName(r.w, ctx.working), aspk,
			genericName,
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

	nameWithTp := named.Name + valuetp
	summonExpr := GenericRepr{
		ReprExpr: func() SummonExpr {
			return r.summon(ctx, metafp.RequiredInstance{
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
		workingScope: r.tcCache.GetLocal(tc.Package, tc.TypeClass),
		primScope:    r.tcCache.Get(tc.PrimitiveInstancePkg, tc.TypeClass),
		tc:           tc,
		working:      workingPackage,
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

	} else if tc.IsRecursive() {
		tcname := tc.TypeClass.PackagedName(r.w, workingPackage)

		return newSummonExpr(fmt.Sprintf(`
						func %s() %s[%s] {
							return %s
						}
					`, tc.GeneratedInstanceName(), tcname, tc.DeriveFor.PackagedName(r.w, workingPackage),
			mapExpr), mapExpr.paramInstance)
	} else {
		return newSummonExpr(fmt.Sprintf(`
						var %s = %s
					`, tc.GeneratedInstanceName(), mapExpr), mapExpr.paramInstance)
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

func genDerive() {
	pack := os.Getenv("GOPACKAGE")

	genfp.Generate(pack, pack+"_derive_generated.go", func(w genfp.Writer) {

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

		d := metafp.FindTypeClassDerive(pkgs)

		if d.Size() == 0 {
			return
		}

		tccache := metafp.TypeClassInstanceCache{}

		metafp.FindTypeClassImport(pkgs).Foreach(func(v metafp.TypeClassDirective) {
			fmt.Printf("Import %s from %s\n", v.TypeClass.Name, v.Package.Path())
			tccache.Load(v.PrimitiveInstancePkg, v.TypeClass)
		})

		seq.Iterator(d).Foreach(func(v metafp.TypeClassDerive) {
			tccache.WillGenerated(v)
		})

		summonCtx := TypeClassSummonContext{
			w:       w,
			tcCache: &tccache,
		}

		d.Foreach(func(v metafp.TypeClassDerive) {
			summonCtx.summonVar(v).Foreach(func(v SummonExpr) {
				fmt.Fprintf(w, "%s\n", v.expr)
			})
		})
	})
}
