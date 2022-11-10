package main

import (
	"fmt"
	"go/types"
	"os"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

type TypeClassSummonContext struct {
	w            genfp.Writer
	tc           metafp.TypeClassDerive
	genSet       mutable.Set[string]
	tcCache      *metafp.TypeClassInstanceCache
	primScope    metafp.TypeClassScope
	workingScope metafp.TypeClassScope
}

type GenericRepr struct {
	//	ReprType     func() string
	ToReprExpr   func() string
	FromReprExpr func() string
	ReprExpr     func() SummonExpr
}

type SummonExpr struct {
	expr          string
	paramInstance fp.Seq[string]
}

func (r SummonExpr) Expr() string {
	return r.expr
}
func (r SummonExpr) String() string {
	return r.expr
}

func (r SummonExpr) ParamInstance() fp.Seq[string] {
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
	paramList := seq.Map(list, SummonExpr.ParamInstance).Reduce(MergeSeqDistinct(eq.String))
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

	param := option.Map(r.typeParam, func(v metafp.TypeClass) string {
		return fmt.Sprintf("%s %s[%s]", r.name, v.PackagedName(w, workingPkg), r.instanceOf.Name().Get())
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

func (r TypeClassSummonContext) typeclassInstanceMust(req metafp.RequiredInstance, name string) lookupTarget {

	f := req.Type
	return lookupTarget{
		instanceOf: f,
		pk:         r.tc.Package,
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
func (r TypeClassSummonContext) lookupTypeClassInstanceLocalDeclared(req metafp.RequiredInstance, name ...string) fp.Option[lookupTarget] {

	f := req.Type

	scope := r.workingScope
	if req.TypeClass.Id() != r.tc.TypeClass.Id() {
		scope = r.tcCache.GetLocal(r.tc.Package, req.TypeClass)
	}
	itr := seq.FlatMap(name, func(v string) fp.Seq[string] {
		if f.Pkg != nil && r.tc.Package.Path() != f.Pkg.Path() {
			return []string{
				req.TypeClass.Name + publicName(f.Pkg.Name()) + publicName(v),
				req.TypeClass.Name + publicName(v),
			}

		}
		return []string{req.TypeClass.Name + publicName(v)}
	}).Iterator()

	ins := iterator.FlatMap(itr, func(v string) fp.Iterator[metafp.TypeClassInstance] {
		return option.Iterator(scope.FindByName(v, f))
	})

	ins = ins.Filter(func(tci metafp.TypeClassInstance) bool {
		return r.checkRequired(tci.RequiredInstance)
	})

	if f.TypeArgs.Size() > 0 {
		ins = scope.Find(f).Iterator().Concat(ins)
	} else {
		ins = ins.Concat(scope.Find(f).Iterator())
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

	}).Head()

}

func (r TypeClassSummonContext) lookupHNilMust(tc metafp.TypeClass) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(tc, "HNil")
	if ret.IsDefined() {
		return ret.Get()
	}

	ret = r.lookupTypeClassFunc(tc, "HlistNil")
	if ret.IsDefined() {
		return ret.Get()
	}
	nameWithTc := r.tc.TypeClass.Name + "HNil"

	return metafp.TypeClassInstance{
		Package: r.tc.Package,
		Name:    nameWithTc,
		Static:  true,
	}
}

func (r TypeClassSummonContext) lookupTypeClassFunc(tc metafp.TypeClass, name string) fp.Option[metafp.TypeClassInstance] {
	nameWithTc := tc.Name + name

	workingScope := r.workingScope
	primScope := r.primScope
	if r.tc.TypeClass.Id() != tc.Id() {
		primScope = r.tcCache.GetImported(tc)
		workingScope = r.tcCache.GetLocal(r.tc.Package, tc)
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

func (r TypeClassSummonContext) lookupTypeClassFuncMust(tc metafp.TypeClass, name string) metafp.TypeClassInstance {
	ret := r.lookupTypeClassFunc(tc, name)
	if ret.IsDefined() {
		return ret.Get()
	}

	nameWithTc := r.tc.TypeClass.Name + name

	return metafp.TypeClassInstance{
		Package: r.tc.Package,
		Name:    nameWithTc,
		Static:  true,
	}
}

func (r TypeClassSummonContext) lookupTypeClassInstancePrimitivePkgLazy(req metafp.RequiredInstance, name ...string) func() fp.Option[lookupTarget] {
	return func() fp.Option[lookupTarget] {
		return r.lookupTypeClassInstancePrimitivePkg(req, name...)
	}
}

func (r TypeClassSummonContext) checkRequired(required fp.Seq[metafp.RequiredInstance]) bool {
	for _, v := range required {
		if v.Type.IsTuple() {
			req := seq.Map(v.Type.TypeArgs, func(t metafp.TypeInfo) metafp.RequiredInstance {
				return metafp.RequiredInstance{
					TypeClass: v.TypeClass,
					Type:      t,
				}
			})
			res := r.checkRequired(req)
			if res == false {
				return false
			}

		} else {
			res := r.lookupTypeClassInstance(v)
			if res.available.IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (r TypeClassSummonContext) lookupTypeClassInstancePrimitivePkg(req metafp.RequiredInstance, name ...string) fp.Option[lookupTarget] {

	scope := r.primScope
	if r.tc.TypeClass.Id() != req.TypeClass.Id() {
		scope = r.tcCache.GetImported(req.TypeClass)
	}
	f := req.Type
	itr := seq.FlatMap(name, func(v string) fp.Seq[string] {
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
	}).Iterator()

	ins := iterator.FlatMap(itr, func(v string) fp.Iterator[metafp.TypeClassInstance] {
		return option.Iterator(scope.FindByName(v, f))
	}).Concat(scope.Find(f).Iterator())

	if f.TypeArgs.Size() > 0 {
		ins = scope.Find(f).Iterator().Concat(ins)
	} else {
		ins = ins.Concat(scope.Find(f).Iterator())
	}

	ins = ins.Filter(func(tci metafp.TypeClassInstance) bool {
		return r.checkRequired(tci.RequiredInstance)
	})

	return iterator.Map(ins, func(v metafp.TypeClassInstance) lookupTarget {
		return lookupTarget{
			instanceOf: f,
			pk:         v.Package,
			name:       v.Name,
			required:   v.RequiredInstance,
			instance:   option.Some(v),
		}
	}).Head()

}

func (r TypeClassSummonContext) lookupTypeClassInstanceTypePkg(req metafp.RequiredInstance, name string) fp.Option[lookupTarget] {

	f := req.Type
	if f.Pkg != nil && f.Pkg.Path() != r.tc.Package.Path() {

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
			}

			return option.Some(ret)

		}
	}

	return option.None[lookupTarget]()
}

func (r TypeClassSummonContext) namedLookup(req metafp.RequiredInstance, name string) typeClassInstance {
	ret := r.lookupTypeClassInstanceLocalDeclared(req, name).Or(lazy.Func2(r.lookupTypeClassInstanceTypePkg)(req, name)).Or(r.lookupTypeClassInstancePrimitivePkgLazy(req, name))

	return typeClassInstance{
		ret,
		r.typeclassInstanceMust(req, name),
	}
}

func (r TypeClassSummonContext) lookupPrimitiveTypeClassInstance(req metafp.RequiredInstance, name ...string) typeClassInstance {
	ret := r.lookupTypeClassInstanceLocalDeclared(req, name...).Or(r.lookupTypeClassInstancePrimitivePkgLazy(req, name...))

	return typeClassInstance{
		ret,
		r.typeclassInstanceMust(req, name[0]),
	}
}

func (r TypeClassSummonContext) typeParamString(lt lookupTarget) fp.Option[string] {

	if lt.instance.IsDefined() {
		ins := lt.instance.Get()

		possible := ins.TypeParam.ForAll(func(v metafp.TypeParam) bool {
			return ins.UsedParam.Contains(v.Name)
		})

		if !possible {
			ret := seq.Map(ins.TypeParam, func(v metafp.TypeParam) string {
				return option.Map(ins.ParamMapping.Get(v.Name), func(v metafp.TypeInfo) string {
					return r.w.TypeName(r.tc.Package, v.Type)
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

func (r TypeClassSummonContext) summArgs(args fp.Seq[metafp.RequiredInstance]) SummonExpr {
	list := seq.Map(args, func(t metafp.RequiredInstance) SummonExpr {
		return r.summon(t)
	})

	return collectSummonExpr(list)
}

func newSummonExpr(expr string, params ...fp.Seq[string]) SummonExpr {
	return SummonExpr{
		expr:          expr,
		paramInstance: as.Seq(params).Reduce(MergeSeqDistinct(eq.String)),
	}
}

func (r TypeClassSummonContext) exprTypeClassInstance(lt lookupTarget) SummonExpr {
	if len(lt.required) > 0 {
		list := r.summArgs(lt.required)

		instanceExpr := lt.instanceExpr(r.w, r.tc.Package)
		tpstr := r.typeParamString(lt)
		if tpstr.IsDefined() {
			//fmt.Printf("%s param infer not possible = %s \n", lt.name, lt.instance.Get().ParamMapping)

			return newSummonExpr(fmt.Sprintf("%s[%s](%s)", instanceExpr.expr, tpstr.Get(), list.expr), instanceExpr.paramInstance, list.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s(%s)", instanceExpr.expr, list.expr), instanceExpr.paramInstance, list.paramInstance)

		}
	}

	if lt.isFunc() && len(lt.required) == 0 {
		instanceExpr := lt.instanceExpr(r.w, r.tc.Package)

		tpstr := r.typeParamString(lt)
		if tpstr.IsDefined() {
			return newSummonExpr(fmt.Sprintf("%s[%s]()", instanceExpr, tpstr.Get()), instanceExpr.paramInstance)

		} else {
			return newSummonExpr(fmt.Sprintf("%s()", instanceExpr), instanceExpr.paramInstance)
		}

	}

	return lt.instanceExpr(r.w, r.tc.Package)

}

func (r TypeClassSummonContext) exprTypeClassMember(tc metafp.TypeClass, lt metafp.TypeClassInstance, typeArgs fp.Seq[metafp.TypeInfo]) SummonExpr {
	if len(typeArgs) > 0 {
		list := r.summArgs(seq.Map(typeArgs, func(t metafp.TypeInfo) metafp.RequiredInstance {
			return metafp.RequiredInstance{
				TypeClass: tc,
				Type:      t,
			}
		}))

		return newSummonExpr(fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, r.tc.Package), list), list.paramInstance)
	}

	return newSummonExpr(lt.PackagedName(r.w, r.tc.Package))

}

func (r TypeClassSummonContext) exprTypeClassMemberLabelled(tc metafp.TypeClass, lt metafp.TypeClassInstance, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfo]) SummonExpr {
	if len(typeArgs) > 0 {
		list := collectSummonExpr(seq.Map(seq.Zip(typeArgs, names), func(t fp.Tuple2[metafp.TypeInfo, string]) SummonExpr {
			return r.summonFpNamed(tc, t.I2, t.I1)
		}))

		return newSummonExpr(fmt.Sprintf("%s(%s)", lt.PackagedName(r.w, r.tc.Package), list), list.paramInstance)
	}

	return newSummonExpr(lt.PackagedName(r.w, r.tc.Package))

}

func (r TypeClassSummonContext) lookupTypeClassInstance(req metafp.RequiredInstance) typeClassInstance {
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
				return typeClassInstance{r.lookupTypeClassInstanceLocalDeclared(req, "HNil", "HListNil").
					Or(r.lookupTypeClassInstancePrimitivePkgLazy(req, "HNil", "HListNil")),
					r.typeclassInstanceMust(req, "HNil"),
				}

			} else if at.Obj().Name() == "Cons" {
				return typeClassInstance{
					r.lookupTypeClassInstanceLocalDeclared(req, "HCons", "HListCons").
						Or(r.lookupTypeClassInstancePrimitivePkgLazy(req, "HCons", "HListCons")),

					r.typeclassInstanceMust(req, "HCons"),
				}
			}
		}
		return r.namedLookup(req, at.Obj().Name())
	case *types.Array:
		panic(fmt.Sprintf("can't summon array type, while deriving %s[%s]", req.TypeClass.Name, r.tc.DeriveFor.Name))
		//return r.namedLookup(f, "Array")
	case *types.Slice:
		if at.Elem().String() == "byte" {
			bytesInstance := r.namedLookup(
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
			return r.namedLookup(req, "Slice")
		}
		return r.namedLookup(req, "Slice")
	case *types.Map:
		return r.namedLookup(req, "GoMap")
	case *types.Pointer:
		return r.namedLookup(req, "Ptr")
	case *types.Basic:
		return r.namedLookup(req, at.Name())
	case *types.Struct:
		panic(fmt.Sprintf("can't summon unnamed struct type, while deriving %s[%s]", r.tc.TypeClass.Name, r.tc.DeriveFor.Name))
	case *types.Interface:
		if f.IsAny() {
			return r.namedLookup(req, "Given")
		}
		panic(fmt.Sprintf("can't summon unnamed interface type %v, while deriving %s[%s]", f.Type, r.tc.TypeClass.Name, r.tc.DeriveFor.Name))
	case *types.Chan:
		panic(fmt.Sprintf("can't summon unnamed chan type, while deriving %s[%s]", r.tc.TypeClass.Name, r.tc.DeriveFor.Name))

	}
	return r.namedLookup(req, f.Type.String())
}

func (r TypeClassSummonContext) summonLabelledGenericRepr(tc metafp.TypeClass, receiver string, receiverType string, builderreceiver string, names fp.Seq[string], typeArgs fp.Seq[metafp.TypeInfo]) fp.Option[GenericRepr] {
	result := r.lookupTypeClassFunc(tc, fmt.Sprintf("Labelled%d", typeArgs.Size()))

	return option.Map(result, func(tm metafp.TypeClassInstance) GenericRepr {
		return GenericRepr{
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr: func() string {
				return fmt.Sprintf("%s.AsLabelled", receiver)
			},
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
				return r.exprTypeClassMemberLabelled(tc, tm, names, typeArgs)
			},
		}
	}).Or(func() fp.Option[GenericRepr] {
		return option.Map(r.lookupTypeClassFunc(tc, "HConsLabelled"), func(hcons metafp.TypeClassInstance) GenericRepr {
			return GenericRepr{
				// ReprType: func() string {
				// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
				// },
				ToReprExpr: func() string {

					if typeArgs.Size() == 0 {
						hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))

						return fmt.Sprintf(`func (%s%s) %s.Nil {
							return %s.Empty()
						}`, receiver, receiverType, hlistpk, hlistpk)
					} else {
						fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
						aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

						namedTypeArgs := seq.Zip(names, typeArgs)

						tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
							return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
						}).MakeString(",")

						return fmt.Sprintf(`%s.Compose(
						%s.AsLabelled,
						%s.HList%dLabelled[%s],
					)`, fppk,
							receiver,
							aspk, typeArgs.Size(), tp,
						)
					}

				},
				FromReprExpr: func() string {
					fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
					aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
					//hlistpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/hlist", "hlist"))
					productpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/product", "product"))

					namedTypeArgs := seq.Zip(names, typeArgs)

					tp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
						return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
					}).MakeString(",")

					// hlisttp := seq.Map(namedTypeArgs, func(f fp.Tuple2[string, metafp.TypeInfo]) string {
					// 	return fmt.Sprintf("NameIs%s[%s]", publicName(f.I1), r.w.TypeName(r.tc.Package, f.I2.Type))
					// }).MakeString(",")

					hlistToTuple := fmt.Sprintf(`%s.LabelledFromHList%d[%s]`,
						productpk,
						typeArgs.Size(), tp,
					)

					// hlistToTuple := fmt.Sprintf(`%s.Func2(
					// 	%s.Case%d[%s,%s.Nil,fp.Labelled%d[%s]],
					// ).ApplyLast(
					// 	%s.Labelled%d[%s] ,
					// )`,
					// 	aspk,
					// 	hlistpk, typeArgs.Size(), hlisttp, hlistpk, typeArgs.Size(), tp,
					// 	aspk, typeArgs.Size(), tp,
					// )

					tupleToStruct := fmt.Sprintf(`%s.Compose(
						%s.Curried2(%s.FromLabelled)(%s{}),
						%s.Build,
						)`,
						fppk,
						aspk, builderreceiver, builderreceiver, builderreceiver,
					)
					return fmt.Sprintf(`
						fp.Compose(
							%s, 
							%s ,
						)`, hlistToTuple, tupleToStruct)
				},
				ReprExpr: func() SummonExpr {
					hnil := r.lookupHNilMust(tc)
					namedTypeArgs := seq.Zip(names, typeArgs)
					hlist := seq.Fold(namedTypeArgs.Reverse(), newSummonExpr(hnil.PackagedName(r.w, r.tc.Package)), func(tail SummonExpr, ti fp.Tuple2[string, metafp.TypeInfo]) SummonExpr {
						instance := r.summonFpNamed(tc, ti.I1, ti.I2)
						return newSummonExpr(fmt.Sprintf(`%s(
							%s,
						%s,
						)`, hcons.PackagedName(r.w, r.tc.Package), instance, tail), instance.paramInstance, tail.paramInstance)
					})

					return hlist
				},
			}
		})
	})
}

func (r TypeClassSummonContext) summonGenericRepr(tc metafp.TypeClass, receiver string, receiverType string, builderreceiver string, typeArgs fp.Seq[metafp.TypeInfo]) GenericRepr {
	result := r.lookupTypeClassFunc(tc, fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {

		// tp := iterator.Map(typeArgs.Iterator(), func(v metafp.TypeInfo) string {
		// 	return r.w.TypeName(r.tc.Package, v.Type)
		// }).MakeString(",")
		return GenericRepr{
			// ReprType: func() string {
			// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
			// },
			ToReprExpr: func() string {
				return fmt.Sprintf("%s.AsTuple", receiver)
			},
			FromReprExpr: func() string {
				fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
				aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

				return fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromTuple)(%s{}),
					%s.Build,
					)`,
					fppk,
					aspk, builderreceiver, builderreceiver, builderreceiver,
				)
			},
			ReprExpr: func() SummonExpr {
				return r.exprTypeClassMember(tc, result.Get(), typeArgs)
			},
		}
	}

	tupleGeneric := r.summonTupleGenericRepr(tc, typeArgs)

	return GenericRepr{
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {
			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))

			return fmt.Sprintf(`%s.Compose(
				%s.AsTuple,
				%s, 
			)`, fppk,
				receiver,
				tupleGeneric.ToReprExpr(),
			)

		},
		FromReprExpr: func() string {
			fppk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp", "fp"))
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			tupleToStruct := fmt.Sprintf(`%s.Compose(
					%s.Curried2(%s.FromTuple)(%s{}),
					%s.Build,
					)`,
				fppk,
				aspk, builderreceiver, builderreceiver, builderreceiver,
			)
			return fmt.Sprintf(`
				fp.Compose(
					%s, 
					%s ,
				)`, tupleGeneric.FromReprExpr(), tupleToStruct)
		},
		ReprExpr: func() SummonExpr {
			return tupleGeneric.ReprExpr()
		},
	}
}

func (r TypeClassSummonContext) summonTupleGenericRepr(tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfo]) GenericRepr {
	return GenericRepr{
		// ReprType: func() string {
		// 	return fmt.Sprintf("Tuple%d[%s]", typeArgs.Size(), tp)
		// },
		ToReprExpr: func() string {
			aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))

			tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(r.tc.Package, f.Type)
			}).MakeString(",")

			return fmt.Sprintf(`%s.HList%d[%s]`,
				aspk, typeArgs.Size(), tp,
			)

		},
		FromReprExpr: func() string {
			productpk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/product", "product"))

			tp := seq.Map(typeArgs, func(f metafp.TypeInfo) string {
				return r.w.TypeName(r.tc.Package, f.Type)
			}).MakeString(",")

			hlistToTuple := fmt.Sprintf(`%s.TupleFromHList%d[%s]`,
				productpk, typeArgs.Size(), tp,
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
			hcons := r.lookupTypeClassFuncMust(tc, "HCons")

			hnil := r.lookupHNilMust(tc)

			hlist := seq.Fold(typeArgs.Reverse(), newSummonExpr(hnil.PackagedName(r.w, r.tc.Package)), func(tail SummonExpr, ti metafp.TypeInfo) SummonExpr {
				instance := r.summon(metafp.RequiredInstance{
					TypeClass: r.tc.TypeClass,
					Type:      ti,
				})
				return newSummonExpr(fmt.Sprintf(`%s(
					%s,
					%s,
				)`, hcons.PackagedName(r.w, r.tc.Package), instance, tail), instance.paramInstance, tail.paramInstance)
			})
			return hlist
		},
	}
}

func (r TypeClassSummonContext) summonTuple(tc metafp.TypeClass, typeArgs fp.Seq[metafp.TypeInfo]) SummonExpr {

	result := r.lookupTypeClassFunc(tc, fmt.Sprintf("Tuple%d", typeArgs.Size()))

	if result.IsDefined() {
		return r.exprTypeClassMember(tc, result.Get(), typeArgs)
	}

	tupleGeneric := r.summonTupleGenericRepr(tc, typeArgs)
	return r.summonVariant(tc, "", tupleGeneric)

}

func (r TypeClassSummonContext) summonFpNamed(tc metafp.TypeClass, name string, t metafp.TypeInfo) SummonExpr {

	instance := r.lookupTypeClassFuncMust(tc, "Named")

	expr := r.summon(metafp.RequiredInstance{
		TypeClass: r.tc.TypeClass,
		Type:      t,
	})

	return newSummonExpr(fmt.Sprintf("%s[NameIs%s[%s]](%s)", instance.PackagedName(r.w, r.tc.Package), publicName(name),
		r.w.TypeName(r.tc.Package, t.Type), expr.expr), expr.paramInstance)

	// pk := r.w.GetImportedName(r.tc.Package)
	// return fmt.Sprintf("%s.Named(%s)", pk, r.summon(t))
}

func (r TypeClassSummonContext) summon(req metafp.RequiredInstance) SummonExpr {

	t := req.Type

	if t.IsTuple() {
		return r.summonTuple(req.TypeClass, t.TypeArgs)
	}

	result := r.lookupTypeClassInstance(req)

	if result.available.IsDefined() {
		return r.exprTypeClassInstance(result.available.Get())
	}

	// instance := r.lookupTypeClassMember("UInt")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Int")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Float")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Number")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
	// 		}
	// 	}
	// }

	// instance = r.lookupTypeClassMember("Given")
	// if instance.IsDefined() {
	// 	if _, ok := instance.Get().Type().(*types.Signature); ok {
	// 		ctx := types.NewContext()
	// 		_, err := types.Instantiate(ctx, instance.Get().Type(), []types.Type{t.Type}, true)
	// 		if err == nil {
	// 			return fmt.Sprintf("%s[%s]()", instance.Get().PackagedName(r.w, r.tc.Package), r.w.TypeName(r.tc.Package, t.Type))
	// 		}
	// 	}
	// }

	return r.exprTypeClassInstance(result.must)

}

func (r TypeClassSummonContext) summonStruct(tc metafp.TypeClass, named metafp.NamedTypeInfo, fields fp.Seq[metafp.StructField]) SummonExpr {

	// fmt.Printf("lookup %s.Option = %v\n", v.Generator.Name(), l)
	//fmt.Printf("derive %v for %v\n", v.TypeClass, v.DeriveFor)
	//privateFields := fields.FilterNot(metafp.StructField.Public)

	names := seq.Map(fields, func(v metafp.StructField) string {
		return v.Name
	})

	typeArgs := seq.Map(fields, func(v metafp.StructField) metafp.TypeInfo {
		return v.Type
	})

	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(named.Info.TypeParam.Iterator(), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	builderreceiver := fmt.Sprintf("%sBuilder%s", named.PackagedName(r.w, r.tc.Package), valuetp)
	valuereceiver := fmt.Sprintf("%s%s", named.PackagedName(r.w, r.tc.Package), valuetp)

	labelledExpr := r.summonLabelledGenericRepr(tc, valuereceiver, valuetp, builderreceiver, names, typeArgs)
	summonExpr := labelledExpr.OrElseGet(func() GenericRepr {
		return r.summonGenericRepr(tc, valuereceiver, valuetp, builderreceiver, typeArgs)
	})

	return r.summonVariant(tc, named.GenericName(), summonExpr)
}

func (r TypeClassSummonContext) summonVariant(tc metafp.TypeClass, genericName string, genericRepr GenericRepr) SummonExpr {
	mapExpr := option.Map(r.lookupTypeClassFunc(tc, "Generic"), func(generic metafp.TypeClassInstance) SummonExpr {

		aspk := r.w.GetImportedName(types.NewPackage("github.com/csgura/fp/as", "as"))
		repr := genericRepr.ReprExpr()
		return newSummonExpr(fmt.Sprintf(`%s(
					%s.Generic(
							"%s",
							%s,
							%s,
						), 
						%s, 
					)`, generic.PackagedName(r.w, r.tc.Package), aspk,
			genericName,
			genericRepr.ToReprExpr(),
			genericRepr.FromReprExpr(),
			repr), repr.paramInstance)

	}).Or(func() fp.Option[SummonExpr] {
		return option.Map(r.lookupTypeClassFunc(tc, "IMap"), func(imapfunc metafp.TypeClassInstance) SummonExpr {
			repr := genericRepr.ReprExpr()

			return newSummonExpr(fmt.Sprintf(`%s( 
						%s, 
						%s , 
						%s,
						)`,
				imapfunc.PackagedName(r.w, r.tc.Package), repr, genericRepr.FromReprExpr(), genericRepr.ToReprExpr()), repr.paramInstance)
		})
	}).Or(func() fp.Option[SummonExpr] {
		functor := r.lookupTypeClassFunc(tc, "Map")
		return option.Map(functor, func(v metafp.TypeClassInstance) SummonExpr {
			repr := genericRepr.ReprExpr()

			return newSummonExpr(fmt.Sprintf(`%s( 
						%s, 
						%s,
						)`,
				v.PackagedName(r.w, r.tc.Package), repr, genericRepr.FromReprExpr(),
			), repr.paramInstance)
		})

	}).OrElseGet(func() SummonExpr {
		contrmap := r.lookupTypeClassFuncMust(tc, "ContraMap")
		repr := genericRepr.ReprExpr()

		return newSummonExpr(fmt.Sprintf(`%s( 
					%s , 
					%s,
					)`,
			contrmap.PackagedName(r.w, r.tc.Package), repr, genericRepr.ToReprExpr(),
		), repr.paramInstance)
	})
	return mapExpr

}

func (r TypeClassSummonContext) summonNamed(tc metafp.TypeClass, named metafp.NamedTypeInfo) SummonExpr {

	valuetp := ""
	if named.Info.TypeParam.Size() > 0 {
		valuetp = "[" + iterator.Map(named.Info.TypeParam.Iterator(), func(v metafp.TypeParam) string {
			return v.Name
		}).MakeString(",") + "]"
	}

	nameWithTp := named.Name + valuetp
	summonExpr := GenericRepr{
		ReprExpr: func() SummonExpr {
			return r.summon(metafp.RequiredInstance{
				TypeClass: tc,
				Type:      named.Underlying,
			})
		},
		ToReprExpr: func() string {
			return fmt.Sprintf(`func(v %s) %s {
					return %s(v)
				}`, nameWithTp, r.w.TypeName(r.tc.Package, named.Underlying.Type), r.w.TypeName(r.tc.Package, named.Underlying.Type))
		},
		FromReprExpr: func() string {
			return fmt.Sprintf(`func(v %s) %s {
					return %s(v)
				}`, r.w.TypeName(r.tc.Package, named.Underlying.Type), nameWithTp, nameWithTp)
		},
	}

	return r.summonVariant(tc, named.GenericName(), summonExpr)
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

		genSet := iterator.ToGoSet(iterator.Map(d.Iterator(), func(v metafp.TypeClassDerive) string {
			tccache.WillGenerated(v)
			return fmt.Sprintf("%s", v.GeneratedInstanceName())
		}))

		d.Foreach(func(v metafp.TypeClassDerive) {

			workingPackage := v.Package

			summonCtx := TypeClassSummonContext{
				w:            w,
				tc:           v,
				genSet:       genSet,
				tcCache:      &tccache,
				workingScope: tccache.GetLocal(v.Package, v.TypeClass),
				primScope:    tccache.Get(v.PrimitiveInstancePkg, v.TypeClass),
			}

			valuetpdec := ""
			valuetp := ""
			if v.DeriveFor.Info.TypeParam.Size() > 0 {
				valuetpdec = "[" + iterator.Map(v.DeriveFor.Info.TypeParam.Iterator(), func(v metafp.TypeParam) string {
					tn := w.TypeName(workingPackage, v.Constraint)
					return fmt.Sprintf("%s %s", v.Name, tn)
				}).MakeString(",") + "]"

				valuetp = "[" + iterator.Map(v.DeriveFor.Info.TypeParam.Iterator(), func(v metafp.TypeParam) string {
					return v.Name
				}).MakeString(",") + "]"
			}

			mapExpr := option.Map(v.StructInfo, func(s metafp.TaggedStruct) SummonExpr {
				fields := s.Fields
				privateFields := fields.FilterNot(metafp.StructField.Public)

				return summonCtx.summonStruct(v.TypeClass, v.DeriveFor, privateFields)
			}).OrElseGet(func() SummonExpr {
				return summonCtx.summonNamed(v.TypeClass, v.DeriveFor)
			})

			if v.DeriveFor.Info.TypeParam.Size() > 0 {

				tcname := v.TypeClass.PackagedName(w, workingPackage)
				// fargs := seq.Map(v.DeriveFor.Info.TypeParam, func(p metafp.TypeParam) string {
				// 	return fmt.Sprintf("%s%s %s[%s] ", privateName(v.TypeClass.Name), p.Name, tcname, p.Name)
				// }).MakeString(",")

				fargs := mapExpr.paramInstance.MakeString(",")

				fmt.Fprintf(w, `
						func %s%s( %s ) %s[%s%s] {
							return %s
						}
					`, v.GeneratedInstanceName(), valuetpdec, fargs, tcname, v.DeriveFor.PackagedName(w, workingPackage), valuetp,
					mapExpr)

			} else if v.IsRecursive() {
				tcname := v.TypeClass.PackagedName(w, workingPackage)

				fmt.Fprintf(w, `
						func %s() %s[%s] {
							return %s
						}
					`, v.GeneratedInstanceName(), tcname, v.DeriveFor.PackagedName(w, workingPackage),
					mapExpr)
			} else {
				fmt.Fprintf(w, `
						var %s = %s
					`, v.GeneratedInstanceName(), mapExpr)
			}

		})
	})
}
