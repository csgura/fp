package genfp

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strings"
)

func InstiateTransfomer(w Writer, pk *types.Package, realtp *types.Named, monadType *types.Named, p *types.TypeParam) func(string, ...any) string {

	rettype := NameParamReplaced(w, pk, monadType, p)
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
						args = append(args, rettype(newname, fmtargs...))
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

func exprString(expr ast.Expr) string {
	fs := token.NewFileSet()

	buf := &bytes.Buffer{}

	printer.Fprint(buf, fs, expr)
	return buf.String()
}

func removeTypeParams(s string) string {
	start := strings.Index(s, "[")
	return s[:start]
}

func CallFunc(w Writer, tr TypeReference) string {
	for _, i := range tr.Imports {
		w.AddImport(types.NewPackage(i.Package, i.Name))
	}

	if _, ok := tr.Type.(*types.Signature); ok {
		if fl, ok := tr.Expr.(*ast.FuncLit); ok {
			fs := token.NewFileSet()

			buf := &bytes.Buffer{}

			printer.Fprint(buf, fs, fl)
			return removeTypeParams(buf.String())
		} else {
			return removeTypeParams(exprString(tr.Expr))
		}
	}
	panic("tr is not types.Signature")

}

func WriteMonadTransformers(w Writer, md GenerateMonadTransformerDirective) {

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

	tpargs1 := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A1 %s", w.TypeName(md.Package.Types, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package.Types, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	monadtype := NameParamReplaced(w, md.Package.Types, md.MonadType, md.TypeParm)

	combinedtype := InstiateTransfomer(w, md.Package.Types, md.TargetType, md.MonadType, md.TypeParm)

	// srctype := rettype("A")
	// rettp := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
	// 	if tp, ok := t.(*types.TypeParam); ok {
	// 		if tp.Obj().Name() == md.TypeParm.Obj().Name() {
	// 			return "R"

	// 		} else {
	// 			return tp.Obj().Name()
	// 		}
	// 	}
	// 	return ""
	// }), func(v string) bool { return v != "" }), ",")
	w.AddImport(types.NewPackage("github.com/csgura/fp", "fp"))

	//typeparams := TypeParamReplaced(w, md.Package.Types, md.TargetType, md.TypeParm)
	fixedParams := FixedParams(w, md.Package.Types, md.TargetType, md.TypeParm)

	pureins := CallFunc(w, md.Pure)
	puref := func(v string, args ...any) string {
		return fmt.Sprintf("%s(%s)", pureins, fmt.Sprintf("%s", args...))
	}
	flatmapf := CallFunc(w, md.FlatMap)

	funcs := map[string]any{
		"puret": func(v string, tpe string) string {
			if fixedParams != "" {
				return fmt.Sprintf("Pure[%s](%s)", fixedParams, puref("%s", v))
			}
			return fmt.Sprintf("Pure(%s)", puref("%s", v))
		},
		"pure": func(of string) string {
			return pureins
		},
		"flatmap": func(from, to string) string {
			return flatmapf
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

		"combined": combinedtype,
		"monad":    monadtype,

		"monadIns": func(start, until int) string {
			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "ins%d %s", j, combinedtype("A%d", j))
			}
			return f.String()
		},
		"monadTypes": func(start, until int) string {
			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "%s", combinedtype("A%d", j))
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
					fmt.Fprintf(f, "f%d %s[A%d,%s]", j, ft, j, combinedtype("R"))
				} else {
					fmt.Fprintf(f, "f%d %s[A%d,%s]", j, ft, j, combinedtype("A%d", j+1))
				}
			}
			return f.String()
		},
	}
	param := map[string]any{
		"tpargs":  tpargs,
		"tpargs1": tpargs1,

		"tp":   md.TypeParm.String(),
		"name": md.Name,
	}

	w.Render(`
		func Pure{{.name}}[{{.tpargs}}](a A) {{combined "A"}} {
			return {{puret "a" "A"}}
		}

		func Map{{.name}}[{{.tpargs}},B any](t {{combined "A"}}, f func(A) B) {{combined "B"}} {
			return Map(t, func( ma {{monad "A"}} )  {{monad "B"}} {
				return {{flatmap "A" "B"}}(ma, func(a A) {{monad "B"}} {
					return {{pure "B"}}(f(a))
				})
			})
		}

		func SubFlatMap{{.name}}[{{.tpargs}},B any](t {{combined "A"}}, f func(A) {{monad "B"}}) {{combined "B"}} {
			return Map(t, func( ma {{monad "A"}} )  {{monad "B"}} {
				return {{flatmap "A" "B"}}(ma, func(a A) {{monad "B"}} {
					return f(a)
				})
			})
		}
	`, funcs, param)

}
