package generator

import (
	"bytes"
	"fmt"
	"go/types"
	"slices"
	"strings"
	"text/template"

	"github.com/csgura/fp/genfp"
)

type Writer = genfp.Writer

func FixedParams(w Writer, pk genfp.WorkingPackage, realtp GenericType, p *types.TypeParam) []string {
	if realtp.TypeArgs() != nil {
		args := []string{}
		for i := 0; i < realtp.TypeArgs().Len(); i++ {
			ta := realtp.TypeArgs().At(i)
			if tp, ok := ta.(*types.TypeParam); ok {
				if tp.Obj().Name() == p.Obj().Name() {
				} else {
					args = append(args, w.TypeName(pk, ta))
				}
			} else {
				args = append(args, w.TypeName(pk, ta))
			}
		}

		return args
	}
	return nil
}

func VariableParams(w Writer, pk genfp.WorkingPackage, realtp GenericType, fixedParam []string) []string {
	if realtp.TypeArgs() != nil {
		args := []string{}
		for i := 0; i < realtp.TypeArgs().Len(); i++ {
			ta := realtp.TypeArgs().At(i)
			if tp, ok := ta.(*types.TypeParam); ok {
				if !slices.Contains(fixedParam, tp.Obj().Name()) {
					args = append(args, w.TypeName(pk, ta))
				}
			}
		}
		return args
	}
	return nil
}

func TypeParamReplaced(w Writer, pk genfp.WorkingPackage, realtp GenericType, p *types.TypeParam) func(string, ...any) string {
	return func(newname string, fmtargs ...any) string {
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

			return argsstr
		}
		return ""
	}
}

type GenericType interface {
	TypeArgs() *types.TypeList
	Obj() *types.TypeName
}

func NameParamReplaced(w Writer, pk genfp.WorkingPackage, realtp GenericType, p *types.TypeParam) func(string, ...any) string {
	tpname := realtp.Obj().Name()
	nameWithPkg := tpname
	if realtp.Obj().Pkg() != nil && realtp.Obj().Pkg().Path() != pk.Path() {
		alias := w.GetImportedName(genfp.FromTypesPackage(realtp.Obj().Pkg()))

		nameWithPkg = fmt.Sprintf("%s.%s", alias, tpname)
	}

	return func(newname string, fmtargs ...any) string {

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

type genFuncContext struct {
	w               Writer
	definedFunction map[string]bool
	funcs           map[string]any
	param           map[string]any
}

func (r genFuncContext) defineFunc(name string, template string) {
	if !r.definedFunction[name] {
		r.param["funcname"] = name
		r.w.Render(template, r.funcs, r.param)
		r.definedFunction[name] = true
	}
}

func (r genFuncContext) defineFuncs(start, end int, nametempalte string, bodytemplate string) {
	t, err := template.New("defineFuncs").Parse(nametempalte)
	if err == nil {
		for i := start; i < end; i++ {
			buf := &bytes.Buffer{}
			r.param["N"] = i
			err := t.Execute(buf, r.param)
			if err != nil {
				fmt.Printf("template = %s\n", nametempalte)
				panic(err)
			}

			r.defineFunc(buf.String(), bodytemplate)
		}
	} else {
		fmt.Printf("template = %s\n", nametempalte)
		panic(err)
	}
}

func WriteMonadFunctions(w Writer, md GenerateMonadFunctionsDirective, definedFunction map[string]bool) {

	tp := md.TargetType.TypeArgs()
	tpargs := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A %s", w.TypeName(md.Package, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	tpargs1 := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A1 %s", w.TypeName(md.Package, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	rettype := NameParamReplaced(w, md.Package, md.TargetType, md.TypeParm)

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
	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

	//typeparams := TypeParamReplaced(w, md.Package.Types, md.TargetType, md.TypeParm)
	fixedParams := strings.Join(FixedParams(w, md.Package, md.TargetType, md.TypeParm), ",")

	funcs := map[string]any{
		"pure": func(v string, tpe string) string {
			if fixedParams != "" {
				return fmt.Sprintf("Pure[%s](%s)", fixedParams, v)
			}
			return fmt.Sprintf("Pure(%s)", v)
		},

		"infer": func(extra ...string) string {
			if fixedParams == "" {
				if len(extra) > 0 {
					return "[" + strings.Join(extra, ",") + "]"
				}
				return ""
			}
			if len(extra) > 0 {
				return fmt.Sprintf("[%s,%s]", fixedParams, strings.Join(extra, ","))
			}
			return fmt.Sprintf("[%s]", fixedParams)
		},

		"monad": rettype,
		"monadIns": func(start, until int) string {
			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "ins%d %s", j, rettype("A%d", j))
			}
			return f.String()
		},
		"monadTypes": func(start, until int) string {
			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "%s", rettype("A%d", j))
			}
			return f.String()
		},
		"funcChain": func(start, until int, funcType ...string) string {
			ft := "fp.Func1"
			if len(funcType) > 0 {
				ft = funcType[0]
			}

			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				if j == until {
					fmt.Fprintf(f, "f%d %s[A%d,%s]", j, ft, j, "R")
				} else {
					fmt.Fprintf(f, "f%d %s[A%d,%s]", j, ft, j, fmt.Sprintf("A%d", j+1))
				}
			}
			return f.String()
		},
		"monadFuncChain": func(start, until int, funcType ...string) string {
			ft := "fp.Func1"
			if len(funcType) > 0 {
				ft = funcType[0]
			}

			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				if j == until {
					fmt.Fprintf(f, "f%d %s[A%d,%s]", j, ft, j, rettype("R"))
				} else {
					fmt.Fprintf(f, "f%d %s[A%d,%s]", j, ft, j, rettype("A%d", j+1))
				}
			}
			return f.String()
		},
	}
	param := map[string]any{
		"tpargs":  tpargs,
		"tpargs1": tpargs1,

		"tp": md.TypeParm.String(),
	}

	ctx := genFuncContext{
		w:               w,
		definedFunction: definedFunction,
		funcs:           funcs,
		param:           param,
	}

	ctx.defineFunc("Flatten", `
		func {{.funcname}}[{{.tpargs}}](tta {{monad (monad "A")}}) {{monad "A"}} {
			return FlatMap(tta, func(v {{monad "A"}}) {{monad "A"}} {
				return v
			})
		}
	`)

	ctx.defineFunc("Map", fmt.Sprintf(`
			func {{.funcname}}[%s, R any](m %s,  f func(A) R) %s {
				return FlatMap(m, fp.Compose2(f, Pure[%s]))
			}
		`, tpargs, srctype, rettype("R"), rettp),
	)
	ctx.defineFunc("Replace", fmt.Sprintf(`
		// haskell 의 <$
		// map . const 와 같은 함수
		func {{.funcname}}[%s, R any](s %s, b R) %s {
			return Map(s, fp.Const[A](b))
		}
	`, tpargs, srctype, rettype("R")))

	ctx.defineFunc("Map2", fmt.Sprintf(`
		func {{.funcname}}[%s, B, R any](first %s, second %s, fab func(A, B) R) %s {
			return FlatMap(first, func(a A) %s {
				return Map(second, func(b B) R {
					return fab(a, b)
				})
			})
		}
	`, tpargs, srctype, rettype("B"), rettype("R"), rettype("R"),
	))

	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/product", "product"))
	ctx.defineFunc("Zip", fmt.Sprintf(`
		func {{.funcname}}[%s, B any](first %s, second %s) %s {
			return Map2(first, second, product.Tuple2)
		}

	`, tpargs, srctype, rettype("B"), rettype("fp.Tuple2[A,B]"),
	))

	ctx.defineFunc("Ap", fmt.Sprintf(`
		func {{.funcname}}[%s, B any](tfab %s, ta %s) %s {
			return FlatMap(tfab, func(fab fp.Func1[A, B]) %s {
				return Map(ta, fab)
			})
		}

	`, tpargs, rettype("fp.Func1[A,B]"), rettype("A"), rettype("B"),
		rettype("B"),
	))

	ctx.defineFunc("Compose", fmt.Sprintf(`
		func {{.funcname}}[%s, B, C any](f1 func(A) %s, f2 func(B) %s) func(A) %s {
			return func(a A) %s {
				return FlatMap(f1(a), f2)
			}
		}

	`, tpargs, rettype("B"), rettype("C"), rettype("C"),
		rettype("C"),
	))

	ctx.defineFunc("Compose2", fmt.Sprintf(`
		func {{.funcname}}[%s, B, C any](f1 func(A) %s, f2 func(B) %s) func(A) %s {
			return func(a A) %s {
				return FlatMap(f1(a), f2)
			}
		}

	`, tpargs, rettype("B"), rettype("C"), rettype("C"),
		rettype("C"),
	))

	ctx.defineFunc("ApFunc", fmt.Sprintf(`
		func {{.funcname}}[%s, B any](tfab %s, ta func() %s) %s {
			return FlatMap(tfab, func(fab fp.Func1[A, B]) %s {
				return Map(ta(), fab)
			})
		}


	`, tpargs, rettype("fp.Func1[A, B]"), rettype("A"), rettype("B"),
		rettype("B"),
	))

	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/iterator", "iterator"))

	ctx.defineFunc("MapSeqLift", fmt.Sprintf(`
			// Map(ta , seq.Lift(f)) 와 동일
			func {{.funcname}}[%s, B any](ta %s, f func(v A) B) %s {

				return Map(ta, func(a fp.Seq[A]) fp.Seq[B] {
					return iterator.Map(iterator.FromSeq(a), f).ToSeq()
				})
			}


	`, tpargs, rettype("fp.Seq[A]"), rettype("fp.Seq[B]"),
	))

	ctx.defineFunc("MapSliceLift", fmt.Sprintf(`
		// Map(ta , seq.Lift(f)) 와 동일
		func {{.funcname}}[%s, B any](ta %s, f func(v A) B) %s {

			return Map(ta, func(a []A) []B {
				return iterator.Map(iterator.FromSeq(a), f).ToSeq()
			})
		}
	`, tpargs, rettype("[]A"), rettype("[]B"),
	))

	ctx.defineFunc("Lift", `
	func {{.funcname}}[{{.tpargs}}, R any](fa func(v A) R) func({{monad "A"}}) {{monad "R"}} {
		return func(ta {{monad "A"}}) {{monad "R"}} {
			return Map(ta, fa)
		}
	}
	`)

	ctx.defineFunc("LiftA2", `
		func {{.funcname}}[{{.tpargs}}, B, R any](fab func(A, B) R) func({{monad "A"}}, {{monad "B"}}) {{monad "R"}} {
			return func(a {{monad "A"}}, b {{monad "B"}}) {{monad "R"}} {
				return Map2(a, b, fab)
			}
		}
	`)

	ctx.defineFunc("LiftM", `
		func {{.funcname}}[{{.tpargs}}, R any](fa func(v A) {{monad "R"}}) func({{monad "A"}}) {{monad "R"}} {
			return func(ta {{monad "A"}}) {{monad "R"}} {
				return Flatten(Map(ta, fa))
			}
		}
	`)

	ctx.defineFunc("LiftM2", `
		// (a -> b -> m r) -> m a -> m b -> m r
		// 하스켈에서는  liftM2 와 liftA2 는 같은 함수이고
		// 위와 같은 함수는 존재하지 않음.
		// hoogle 에서 검색해 보면 , liftJoin2 , bindM2 등의 이름으로 정의된 것이 있음.
		// 하지만 ,  fp 패키지에서도   LiftA2 와 LiftM2 를 동일하게 하는 것은 낭비이고
		// M 은 Monad 라는 뜻인데, Monad는 Flatten, FlatMap 의 의미가 있으니까
		// LiftM2 를 다음과 같이 정의함.
		func {{.funcname}}[{{.tpargs}}, B, R any](fab func(A, B) {{monad "R"}}) func({{monad "A"}}, {{monad "B"}}) {{monad "R"}} {
			return func(a {{monad "A"}}, b {{monad "B"}}) {{monad "R"}} {
				return Flatten(Map2(a, b, fab))
			}
		}
	`)
	ctx.defineFunc("FlatMap2", `
		func {{.funcname}}[{{.tpargs}}, B, R any](first {{monad "A"}}, second {{monad "B"}}, fab func(A, B) {{monad "R"}}) {{monad "R"}} {
			return LiftM2(fab)(first, second)
		}
	`)

	ctx.defineFunc("Flap", `
		// 하스켈 : m( a -> r ) -> a -> m r
		// 스칼라 : M[ A => r ] => A => M[R]
		// 하스켈이나 스칼라의 기본 패키지에는 이런 기능을 하는 함수가 없는데,
		// hoogle 에서 검색해 보면
		// https://hoogle.haskell.org/?hoogle=m%20(%20a%20-%3E%20b)%20-%3E%20a%20-%3E%20m%20b
		// ?? 혹은 flap 이라는 이름으로 정의된 함수가 있음
		func {{.funcname}}[{{.tpargs}}, R any](tfa {{monad "fp.Func1[A,R]"}}) func(A) {{monad "R"}} {
			return func(a A) {{monad "R"}} {
				return Ap(tfa, {{pure "a" "A"}})
			}
		}
	`)

	ctx.defineFunc("Flap2", `
		// 하스켈 : m( a -> b -> r ) -> a -> b -> m r
		func {{.funcname}}[{{.tpargs}}, B, R any](tfab {{monad "fp.Func1[A, fp.Func1[B, R]]"}}) fp.Func1[A, fp.Func1[B, {{monad "R"}}]] {
			return func(a A) fp.Func1[B, {{monad "R"}}] {
				return Flap(Ap(tfab, {{pure "a" "A"}}))
			}
		}
	`)

	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/curried", "curried"))

	ctx.defineFunc("FlapMap", `
		// (a -> b -> r) -> m a -> b -> m r
		// Map 호출 후에 Flap 을 호출 한 것
		//
		// https://hoogle.haskell.org/?hoogle=%28+a+-%3E+b+-%3E++r+%29+-%3E+m+a+-%3E++b+-%3E+m+r+&scope=set%3Astackage
		// liftOp 라는 이름으로 정의된 것이 있음
		func {{.funcname}}[{{.tpargs}}, B, R any](tfab func(A, B) R, a {{monad "A"}}) func(B) {{monad "R"}} {
			return Flap(Map(a, curried.Func2(tfab)))
		}
	`)

	ctx.defineFunc("FlatFlapMap", `
		// ( a -> b -> m r) -> m a -> b -> m r
		//
		//	Flatten . FlapMap
		//
		// https://hoogle.haskell.org/?hoogle=(%20a%20-%3E%20b%20-%3E%20m%20r%20)%20-%3E%20m%20a%20-%3E%20%20b%20-%3E%20m%20r%20
		// om , ==<<  이름으로 정의된 것이 있음
		func {{.funcname}}[{{.tpargs}}, B, R any](fab func(A, B) {{monad "R"}}, ta {{monad "A"}}) func(B) {{monad "R"}} {
			return fp.Compose(FlapMap(fab, ta), Flatten{{infer "R"}})
		}
	`)

	ctx.defineFunc("Method1", `
		// FlatMap 과는 아규먼트 순서가 다른 함수로
		// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
		// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
		// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
		func {{.funcname}}[{{.tpargs}}, B, R any](ta {{monad "A"}}, fab func(a A, b B) R) func(B) {{monad "R"}} {
			return FlapMap(fab, ta)
		}
	`)
	ctx.defineFunc("FlatMethod1", `
		func {{.funcname}}[{{.tpargs}}, B, R any](ta {{monad "A"}}, fab func(a A, b B) {{monad "R"}}) func(B) {{monad "R"}} {
			return FlatFlapMap(fab, ta)
		}
	`)
	ctx.defineFunc("Method2", `
		func {{.funcname}}[{{.tpargs}}, B, C, R any](ta {{monad "A"}}, fabc func(a A, b B, c C) R) func(B, C) {{monad "R"}} {

			return curried.Revert2(Flap2(Map(ta, curried.Func3(fabc))))
			// return func(b B, c C) {{monad "R"}} {
			// 	return Map(a, func(a A) R {
			// 		return cf(a, b, c)
			// 	})
			// }
		}
	`)
	ctx.defineFunc("FlatMethod2", `
		func {{.funcname}}[{{.tpargs}}, B, C, R any](ta {{monad "A"}}, fabc func(a A, b B, c C) {{monad "R"}}) func(B, C) {{monad "R"}} {

			return curried.Revert2(curried.Compose2(Flap2(Map(ta, curried.Func3(fabc))), Flatten{{infer "R"}}))

			// return func(b B, c C) {{monad "R"}} {
			// 	return FlatMap(ta, func(a A) {{monad "R"}} {
			// 		return cf(a, b, c)
			// 	})
			// }
		}

	`)

	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/xtr", "xtr"))
	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/product", "product"))

	ctx.defineFunc("UnZip", `
		func {{.funcname}}[{{.tpargs}}, B any](t {{monad "fp.Tuple2[A, B]"}}) ({{monad "A"}}, {{monad "B"}}) {
			return Map(t, xtr.Head), Map(t, xtr.Tail)
		}
	`)
	ctx.defineFunc("Zip3", `
		func {{.funcname}}[{{.tpargs}}, B, C any](ta {{monad "A"}}, tb {{monad "B"}}, tc {{monad "C"}}) {{monad "fp.Tuple3[A, B, C]"}} {
			return LiftA3{{infer}}(product.Tuple3[A, B, C])(ta, tb, tc)
		}
	`)
	ctx.defineFunc("With", `
		// fp.With 의 try 버젼
		// fp.With 가 Flip 과 사실상 같은 것처럼
		// FlapMap 의 Flip 버젼과 동일
		// var b fp.Try[B]
		// a := try.Sucesss(A{})
		// a.FlatMap( try.With(A.WithB, b))
		// 형태로 코딩 가능
		func {{.funcname}}[{{.tpargs}}, B any](withf func(A, B) A, v {{monad "B"}}) func(A) {{monad "A"}} {
			return Flap(Map(v, fp.Flip2(withf)))
		}
	`)

	ctx.defineFuncs(3, genfp.MaxFunc, "LiftA{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](f func({{DeclArgs 1 .N}}) R) func({{monadTypes 1 .N}}) {{monad "R"}} {
			return func({{monadIns 1 .N}}) {{monad "R"}} {

				return FlatMap(ins1, func(a1 A1) {{monad "R"}} {
					return LiftA{{dec .N}}{{infer}}(func({{DeclArgs 2 .N}}) R {
						return f({{CallArgs 1 .N}})
					})({{CallArgs 2 .N "ins"}})
				})
			}
		}
			`)
	ctx.defineFuncs(3, genfp.MaxFunc, "Map{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any]({{monadIns 1 .N}}, f func({{DeclArgs 1 .N}}) R) {{monad "R"}} {
			return LiftA{{.N}}{{infer}}(f)({{CallArgs 1 .N "ins"}})
		}
	`)
	ctx.defineFuncs(3, genfp.MaxFunc, "LiftM{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](f func({{DeclArgs 1 .N}}) {{monad "R"}}) func({{monadTypes 1 .N}}) {{monad "R"}} {
			return func({{monadIns 1 .N}}) {{monad "R"}} {

				return FlatMap(ins1, func(a1 A1) {{monad "R"}} {
					return LiftM{{dec .N}}(func({{DeclArgs 2 .N}}) {{monad "R"}} {
						return f({{CallArgs 1 .N}})
					})({{CallArgs 2 .N "ins"}})
				})
			}
		}
	`)
	ctx.defineFuncs(3, genfp.MaxFunc, "FlatMap{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any]({{monadIns 1 .N}}, f func({{DeclArgs 1 .N}}) {{monad "R"}}) {{monad "R"}} {
			return LiftM{{.N}}(f)({{CallArgs 1 .N "ins"}})
		}
	`)
	ctx.defineFuncs(3, genfp.MaxFunc, "Flap{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](tf {{monad (CurriedFunc 1 .N "R")}}) {{CurriedFunc 1 .N (monad "R")}} {
			return func(a1 A1) {{CurriedFunc 2 .N (monad "R")}} {
				return Flap{{dec .N}}(Ap(tf, {{pure "a1" "A1"}}))
			}
		}
	`)
	ctx.defineFuncs(3, genfp.MaxFunc, "Method{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](ta1 {{monad "A1"}}, fa1 func({{DeclArgs 1 .N}}) R) func({{TypeArgs 2 .N}}) {{monad "R"}} {
			return func({{DeclArgs 2 .N}}) {{monad "R"}} {
				return Map(ta1, func(a1 A1) R {
					return fa1({{CallArgs 1 .N}})
				})
			}
		}
	`)

	ctx.defineFuncs(3, genfp.MaxFunc, "FlatMethod{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](ta1 {{monad "A1"}}, fa1 func({{DeclArgs 1 .N}}) {{monad "R"}}) func({{TypeArgs 2 .N}}) {{monad "R"}} {
			return func({{DeclArgs 2 .N}}) {{monad "R"}} {
				return FlatMap(ta1, func(a1 A1) {{monad "R"}} {
					return fa1({{CallArgs 1 .N}})
				})
			}
		}
	`)

	ctx.defineFuncs(3, genfp.MaxCompose, "Compose{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any]({{monadFuncChain 1 .N}}) fp.Func1[A1, {{monad "R"}}] {
			return Compose2(f1, Compose{{dec .N}}({{CallArgs 2 .N "f"}}))
		}
	`)

	ctx.defineFuncs(2, genfp.MaxCompose, "MapCompose{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](m {{monad "A1"}}, {{funcChain 1 .N}}) {{monad "R"}} {
			return Map(m , fp.Compose{{.N}}({{CallArgs 1 .N "f"}}))
		}
	`)

	ctx.defineFuncs(2, genfp.MaxCompose, "FlatMapCompose{{.N}}", `
		func {{.funcname}}[{{.tpargs1}}, {{TypeArgs 2 .N}}, R any](m {{monad "A1"}}, {{monadFuncChain 1 .N}}) {{monad "R"}} {
			return FlatMap(m , Compose{{.N}}({{CallArgs 1 .N "f"}}))
		}
	`)
}
