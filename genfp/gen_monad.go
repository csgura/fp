package genfp

import (
	"fmt"
	"go/types"
	"strings"
)

func asPtr[T any](v T) *T {
	return &v
}

func NameParamReplaced(w Writer, pk *types.Package, realtp *types.Named, p *types.TypeParam) func(string, ...any) string {
	return func(newname string, fmtargs ...any) string {
		tpname := realtp.Origin().Obj().Name()
		nameWithPkg := tpname
		if realtp.Obj().Pkg() != nil && realtp.Obj().Pkg().Path() != pk.Path() {
			alias := w.GetImportedName(realtp.Obj().Pkg())

			nameWithPkg = fmt.Sprintf("%s.%s", alias, tpname)
		}

		if realtp.TypeArgs() != nil {
			args := []string{}
			for i := 0; i < realtp.TypeArgs().Len(); i++ {
				ta := realtp.TypeArgs().At(i)
				if tp, ok := ta.(*types.TypeParam); ok {
					if tp.Obj().Name() == p.Obj().Name() {
						args = append(args, fmt.Sprintf(newname, fmtargs...))
					} else {
						args = append(args, w.TypeName(pk, ta))
					}

				} else {
					args = append(args, w.TypeName(pk, ta))
				}
			}

			argsstr := strings.Join(args, ",")

			return fmt.Sprintf("%s[%s]", nameWithPkg, argsstr)
		} else {

			return nameWithPkg

		}
	}

}

func WriteMonadFunctions(w Writer, md GenerateMonadFunctionsDirective) {
	fmt.Printf("generateo for %s\n", md.TargetType)

	tp := md.TargetType.TypeArgs()
	tpargs := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A %s", w.TypeName(md.Package.Types, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package.Types, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	rettype := NameParamReplaced(w, md.Package.Types, md.TargetType, md.TypeParm)

	srctype := rettype("A")
	rettp := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return "R"

			} else {
				return tp.Obj().Name()
			}
		}
		return ""
	}), func(v string) bool { return v != "" }), ",")
	w.AddImport(types.NewPackage("github.com/csgura/fp", "fp"))

	funcs := map[string]any{
		"rettype": rettype,
	}
	param := map[string]any{
		"tpargs": tpargs,
		"tp":     md.TypeParm.String(),
	}

	w.Render(`
		func Flatten[{{.tpargs}}](tta {{rettype (rettype .tp)}}) {{rettype .tp}} {
			return FlatMap(tta, func(v {{rettype .tp}}) {{rettype .tp}} {
				return v
			})
		}
	`, funcs, param)

	fmt.Fprintf(w, `
		func Map[%s, R any](m %s,  f func(%s) R) %s {
			return FlatMap(m, fp.Compose2(f, Pure[%s]))
		}
	`, tpargs, srctype, md.TypeParm, rettype("R"), rettp)

	fmt.Fprintf(w, `
		// haskell 의 <$
		// map . const 와 같은 함수
		func Replace[%s, R any](s %s, b R) %s {
			return Map(s, fp.Const[%s](b))
		}
	`, tpargs, srctype, rettype("R"), md.TypeParm)

	fmt.Fprintf(w, `
		func Map2[%s, B, R any](first %s, second %s, fab func(%s, B) R) %s {
			return FlatMap(first, func(a %s) %s {
				return Map(second, func(b B) R {
					return fab(a, b)
				})
			})
		}
	`, tpargs, srctype, rettype("B"), md.TypeParm, rettype("R"),
		md.TypeParm, rettype("R"),
	)

	w.AddImport(types.NewPackage("github.com/csgura/fp/product", "product"))
	fmt.Fprintf(w, `
		func Zip[%s, B any](first %s, second %s) %s {
			return Map2(first, second, product.Tuple2)
		}

	`, tpargs, srctype, rettype("B"), rettype(fmt.Sprintf("fp.Tuple2[%s,B]", md.TypeParm)),
	)

	fmt.Fprintf(w, `
		func Ap[%s, B any](tfab %s, ta %s) %s {
			return FlatMap(tfab, func(fab fp.Func1[A, B]) %s {
				return Map(ta, fab)
			})
		}

	`, tpargs, rettype("fp.Func1[%s,B]", md.TypeParm), rettype("%s", md.TypeParm), rettype("B"),
		rettype("B"),
	)

	fmt.Fprintf(w, `
		func Compose[%s, B, C any](f1 func(A) %s, f2 func(B) %s) func(A) %s {
			return func(a A) %s {
				return FlatMap(f1(a), f2)
			}
		}

	`, tpargs, rettype("B"), rettype("C"), rettype("C"),
		rettype("C"),
	)

	fmt.Fprintf(w, `
		func Compose2[%s, B, C any](f1 func(A) %s, f2 func(B) %s) func(A) %s {
			return func(a A) %s {
				return FlatMap(f1(a), f2)
			}
		}

	`, tpargs, rettype("B"), rettype("C"), rettype("C"),
		rettype("C"),
	)

	fmt.Fprintf(w, `
		func ApFunc[%s, B any](tfab %s, ta func() %s) %s {
			return FlatMap(tfab, func(fab fp.Func1[A, B]) %s {
				return Map(ta(), fab)
			})
		}


	`, tpargs, rettype("fp.Func1[%s, B]", md.TypeParm), rettype("%s", md.TypeParm), rettype("B"),
		rettype("B"),
	)

	w.AddImport(types.NewPackage("github.com/csgura/fp/iterator", "iterator"))

	fmt.Fprintf(w, `
			// Map(ta , seq.Lift(f)) 와 동일
			func MapSeqLift[%s, B any](ta %s, f func(v %s) B) %s {

				return Map(ta, func(a fp.Seq[%s]) fp.Seq[B] {
					return iterator.Map(iterator.FromSeq(a), f).ToSeq()
				})
			}


	`, tpargs, rettype("fp.Seq[%s]", md.TypeParm), md.TypeParm, rettype("fp.Seq[B]"),
		md.TypeParm,
	)

	fmt.Fprintf(w, `
		// Map(ta , seq.Lift(f)) 와 동일
		func MapSliceLift[%s, B any](ta %s, f func(v %s) B) %s {

			return Map(ta, func(a []%s) []B {
				return iterator.Map(iterator.FromSeq(a), f).ToSeq()
			})
		}
	`, tpargs, rettype("[]%s", md.TypeParm), md.TypeParm, rettype("[]B"),
		md.TypeParm,
	)

	w.Render(`
	func Lift[{{.tpargs}}, R any](fa func(v {{.tp}}) R) func({{rettype .tp}}) {{rettype "R"}} {
		return func(ta {{rettype .tp}}) {{rettype "R"}} {
			return Map(ta, fa)
		}
	}
	`, funcs, param)

	w.Render(`
		func LiftA2[{{.tpargs}}, B, R any](fab func({{.tp}}, B) R) func({{rettype .tp}}, {{rettype "B"}}) {{rettype "R"}} {
			return func(a {{rettype .tp}}, b {{rettype "B"}}) {{rettype "R"}} {
				return Map2(a, b, fab)
			}
		}
	`, funcs, param)

	w.Render(`
		func LiftM[{{.tpargs}}, R any](fa func(v {{.tp}}) {{rettype "R"}}) func({{rettype .tp}}) {{rettype "R"}} {
			return func(ta {{rettype .tp}}) {{rettype "R"}} {
				return Flatten(Map(ta, fa))
			}
		}
	`, funcs, param)

	w.Render(`
		// (a -> b -> m r) -> m a -> m b -> m r
		// 하스켈에서는  liftM2 와 liftA2 는 같은 함수이고
		// 위와 같은 함수는 존재하지 않음.
		// hoogle 에서 검색해 보면 , liftJoin2 , bindM2 등의 이름으로 정의된 것이 있음.
		// 하지만 ,  fp 패키지에서도   LiftA2 와 LiftM2 를 동일하게 하는 것은 낭비이고
		// M 은 Monad 라는 뜻인데, Monad는 Flatten, FlatMap 의 의미가 있으니까
		// LiftM2 를 다음과 같이 정의함.
		func LiftM2[{{.tpargs}}, B, R any](fab func({{.tp}}, B) {{rettype "R"}}) func({{rettype .tp}}, {{rettype "B"}}) {{rettype "R"}} {
			return func(a {{rettype .tp}}, b {{rettype "B"}}) {{rettype "R"}} {
				return Flatten(Map2(a, b, fab))
			}
		}
	`, funcs, param)

	w.Render(`
		// 하스켈 : m( a -> r ) -> a -> m r
		// 스칼라 : M[ A => r ] => A => M[R]
		// 하스켈이나 스칼라의 기본 패키지에는 이런 기능을 하는 함수가 없는데,
		// hoogle 에서 검색해 보면
		// https://hoogle.haskell.org/?hoogle=m%20(%20a%20-%3E%20b)%20-%3E%20a%20-%3E%20m%20b
		// ?? 혹은 flap 이라는 이름으로 정의된 함수가 있음
		func Flap[{{.tpargs}}, R any](tfa {{rettype "fp.Func1[%s,R]" .tp}}) func({{.tp}}) {{rettype "R"}} {
			return func(a {{.tp}}) {{rettype "R"}} {
				return Ap(tfa, Success(a))
			}
		}
	`, funcs, param)

	w.Render(`
		// 하스켈 : m( a -> b -> r ) -> a -> b -> m r
		func Flap2[{{.tpargs}}, B, R any](tfab {{rettype "fp.Func1[%s, fp.Func1[B, R]]" .tp}}) fp.Func1[{{.tp}}, fp.Func1[B, fp.Try[R]]] {
			return func(a {{.tp}}) fp.Func1[B, {{rettype "R"}}] {
				return Flap(Ap(tfab, Success(a)))
			}
		}
	`, funcs, param)

	w.AddImport(types.NewPackage("github.com/csgura/fp/curried", "curried"))

	w.Render(`
		// (a -> b -> r) -> m a -> b -> m r
		// Map 호출 후에 Flap 을 호출 한 것
		//
		// https://hoogle.haskell.org/?hoogle=%28+a+-%3E+b+-%3E++r+%29+-%3E+m+a+-%3E++b+-%3E+m+r+&scope=set%3Astackage
		// liftOp 라는 이름으로 정의된 것이 있음
		func FlapMap[{{.tpargs}}, B, R any](tfab func({{.tp}}, B) R, a {{rettype .tp}}) func(B) {{rettype "R"}} {
			return Flap(Map(a, curried.Func2(tfab)))
		}
	`, funcs, param)

	w.Render(`
		// ( a -> b -> m r) -> m a -> b -> m r
		//
		//	Flatten . FlapMap
		//
		// https://hoogle.haskell.org/?hoogle=(%20a%20-%3E%20b%20-%3E%20m%20r%20)%20-%3E%20m%20a%20-%3E%20%20b%20-%3E%20m%20r%20
		// om , ==<<  이름으로 정의된 것이 있음
		func FlatFlapMap[{{.tpargs}}, B, R any](fab func({{.tp}}, B) {{rettype "R"}}, ta {{rettype .tp}}) func(B) {{rettype "R"}} {
			return fp.Compose(FlapMap(fab, ta), Flatten)
		}
	`, funcs, param)

	w.Render(`
		// FlatMap 과는 아규먼트 순서가 다른 함수로
		// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
		// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
		// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
		func Method1[{{.tpargs}}, B, R any](ta {{rettype .tp}}, fab func(a {{.tp}}, b B) R) func(B) {{rettype "R"}} {
			return FlapMap(fab, ta)
		}

		func FlatMethod1[{{.tpargs}}, B, R any](ta {{rettype .tp}}, fab func(a {{.tp}}, b B) {{rettype "R"}}) func(B) {{rettype "R"}} {
			return FlatFlapMap(fab, ta)
		}

		func Method2[{{.tpargs}}, B, C, R any](ta {{rettype .tp}}, fabc func(a {{.tp}}, b B, c C) R) func(B, C) {{rettype "R"}} {

			return curried.Revert2(Flap2(Map(ta, curried.Func3(fabc))))
			// return func(b B, c C) {{rettype "R"}} {
			// 	return Map(a, func(a {{.tp}}) R {
			// 		return cf(a, b, c)
			// 	})
			// }
		}

		func FlatMethod2[{{.tpargs}}, B, C, R any](ta {{rettype .tp}}, fabc func(a {{.tp}}, b B, c C) {{rettype "R"}}) func(B, C) {{rettype "R"}} {

			return curried.Revert2(curried.Compose2(Flap2(Map(ta, curried.Func3(fabc))), Flatten))

			// return func(b B, c C) {{rettype "R"}} {
			// 	return FlatMap(ta, func(a A) {{rettype "R"}} {
			// 		return cf(a, b, c)
			// 	})
			// }
		}

	`, funcs, param)

	w.AddImport(types.NewPackage("github.com/csgura/fp/xtr", "xtr"))
	w.AddImport(types.NewPackage("github.com/csgura/fp/product", "product"))

	w.Render(`
	func UnZip[{{.tpargs}}, B any](t {{rettype "fp.Tuple2[%s, B]" .tp}}) ({{rettype .tp}}, {{rettype "B"}}) {
		return Map(t, xtr.Head), Map(t, xtr.Tail)
	}

	func Zip3[{{.tpargs}}, B, C any](ta {{rettype .tp}}, tb {{rettype "B"}}, tc {{rettype "C"}}) {{rettype "fp.Tuple3[%s, B, C]" .tp}} {
		return LiftA3(product.Tuple3[{{.tp}}, B, C])(ta, tb, tc)
	}

	// fp.With 의 try 버젼
	// fp.With 가 Flip 과 사실상 같은 것처럼
	// FlapMap 의 Flip 버젼과 동일
	// var b fp.Try[B]
	// a := try.Sucesss(A{})
	// a.FlatMap( try.With(A.WithB, b))
	// 형태로 코딩 가능
	func With[{{.tpargs}}, B any](withf func({{.tp}}, B) {{.tp}}, v {{rettype "B"}}) func({{.tp}}) {{rettype .tp}} {
		return Flap(Map(v, fp.Flip2(withf)))
	}
	`, funcs, param)
}
