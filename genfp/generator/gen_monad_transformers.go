package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strings"

	"github.com/csgura/fp/genfp"
)

func InstiateTransfomer(w genfp.Writer, pk genfp.WorkingPackage, aliastp *types.Alias, realtp *types.Named, monadType GenericType, p *types.TypeParam) func(string, ...any) string {

	if aliastp != nil {
		return NameParamReplaced(w, pk, aliastp, p)
	}

	rettype := NameParamReplaced(w, pk, monadType, p)
	return func(newname string, fmtargs ...any) string {
		tpname := realtp.Origin().Obj().Name()
		nameWithPkg := tpname
		if realtp.Obj().Pkg() != nil && realtp.Obj().Pkg().Path() != pk.Path() {
			alias := w.GetImportedName(genfp.FromTypesPackage(realtp.Obj().Pkg()))

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

func exprString(expr ast.Node) string {
	fs := token.NewFileSet()

	buf := &bytes.Buffer{}

	printer.Fprint(buf, fs, expr)
	return buf.String()
}

type visotor struct {
	replace map[string]string
}

func (r *visotor) Visit(node ast.Node) (w ast.Visitor) {
	switch e := node.(type) {
	case *ast.Ident:
		if tobe, ok := r.replace[e.Name]; ok {
			e.Name = tobe
		}
	}

	return r
}

func reverseMap(m map[string]string) map[string]string {
	ret := map[string]string{}
	for k, v := range m {
		ret[v] = k
	}
	return ret
}

func replaceTypeParam(expr ast.Node, replace map[string]string) string {

	if len(replace) == 0 {
		return exprString(expr)
	}

	ast.Walk(&visotor{replace}, expr)
	ret := exprString(expr)
	ast.Walk(&visotor{reverseMap(replace)}, expr)
	return ret
}

func removeTypeParams(s string) string {
	start := strings.Index(s, "[")
	return s[:start]
}

func CallFunc(w genfp.Writer, tr TypeReference) func(replace map[string]string) string {
	return func(replace map[string]string) string {
		for _, i := range tr.Imports {
			w.AddImport(genfp.NewImportPackage(i.Package, i.Name))
		}

		if _, ok := tr.Type.(*types.Signature); ok {
			return replaceTypeParam(tr.Expr, replace)
		} else if _, ok := tr.Expr.(*ast.FuncLit); ok {
			return replaceTypeParam(tr.Expr, replace)
		}
		panic("tr is not types.Signature")
	}
}

func FlatMapRetType(w genfp.Writer, pk genfp.WorkingPackage, tr TypeReference, fixed []string) string {
	if sig, ok := tr.Type.(*types.Signature); ok {
		if sig.Results().Len() == 1 {
			rettp := sig.Results().At(0).Type()
			if named, ok := rettp.(GenericType); ok {
				vp := VariableParams(w, pk, named, fixed)
				if len(vp) == 1 {
					return vp[0]
				}
			}
		}
	}
	return ""
}

type replaceParam map[string]string

func WriteMonadTransformers(w genfp.Writer, md GenerateMonadTransformerDirective, definedFunc map[string]bool) {

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

	outertype := NameParamReplaced(w, md.Package, md.TargetType, md.TypeParm)
	innertype := NameParamReplaced(w, md.Package, md.ExposureMonadType, md.TypeParm)

	combinedtype := InstiateTransfomer(w, md.Package, md.TargetAlias, md.TargetType, md.ExposureMonadType, md.TypeParm)

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
	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

	//typeparams := TypeParamReplaced(w, md.Package.Types, md.TargetType, md.TypeParm)
	fixedParams := FixedParams(w, md.Package, md.TargetType, md.TypeParm)
	fixedStr := strings.Join(fixedParams, ",")

	givenPkg := ""
	if md.GivenMonad.Pure.Type != nil && len(md.GivenMonad.Pure.Imports) > 0 {
		if md.GivenMonad.Pure.Imports[0].Package != md.Package.Path() {
			w.AddImport(md.GivenMonad.Pure.Imports[0])
			givenPkg = md.GivenMonad.Pure.Imports[0].Name
		}
	}

	givenPure := func(replace map[string]string) string {
		if givenPkg == "" {
			return "Pure"
		}
		givenPureIns := CallFunc(w, md.GivenMonad.Pure)

		return givenPureIns(replace)
	}

	// givenPuref := func(v string, args ...any) string {
	// 	return fmt.Sprintf("%s(%s)", givenPure(replaceParam{
	// 		md.TypeParm.String(): "A",
	// 	}), fmt.Sprintf("%s", args...))
	// }

	givenMap := func() string {
		if givenPkg == "" {
			return "Map"
		}
		return givenPkg + ".Map"
	}()

	givenFlatMap := func() string {
		if givenPkg == "" {
			return "FlatMap"
		}
		return givenPkg + ".FlatMap"
	}()

	pureins := CallFunc(w, md.ExposureMonad.Pure)
	puref := func(args ...any) string {
		return fmt.Sprintf("%s(%s)", pureins(replaceParam{
			md.TypeParm.String(): "A",
		}), fmt.Sprintf("%s", args...))
	}

	flatmapf := CallFunc(w, md.ExposureMonad.FlatMap)

	flatmapRet := FlatMapRetType(w, md.Package, md.ExposureMonad.FlatMap, fixedParams)
	// fmt.Printf("md.TypeParam.String()= %s, flatmapRet = %s\n", md.TypeParm.String(), flatmapRet)

	funcs := map[string]any{
		"puret": func(v string, tpe string) string {
			if fixedStr != "" {
				return fmt.Sprintf("%s[%s](%s)", givenPure(replaceParam{
					md.TypeParm.String(): innertype("A"),
				}), fixedStr, puref(v))
			}
			return fmt.Sprintf("%s(%s)", givenPure(replaceParam{
				md.TypeParm.String(): innertype("A"),
			}), puref(v))
		},
		"pure": func(of string) string {
			return pureins(replaceParam{
				md.TypeParm.String(): of,
			})
		},
		"sequence": func(of string) string {
			sf := CallFunc(w, md.Sequence)
			return sf(replaceParam{md.TypeParm.String(): of})
		},
		"flatmap": func(from, to string) string {
			return flatmapf(replaceParam{
				md.TypeParm.String(): from,
				flatmapRet:           to,
			})
		},

		"infer": func(extra ...string) string {
			if fixedStr == "" {
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
		"inner":    innertype,
		"outer":    outertype,

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

	suffixName := md.Name

	if givenPkg != "" {
		suffixName = ""
	}
	param := map[string]any{
		"tpargs":  tpargs,
		"tpargs1": tpargs1,

		"tp":           md.TypeParm.String(),
		"name":         suffixName,
		"givenMap":     givenMap,
		"givenFlatMap": givenFlatMap,
	}

	ctx := genFuncContext{
		w:               w,
		definedFunction: definedFunc,
		funcs:           funcs,
		param:           param,
	}

	ctx.defineFunc("Pure"+suffixName, `
		func {{.funcname}}[{{.tpargs}}](a A) {{combined "A"}} {
			return {{puret "a" "A"}}
		}
	`)

	targetMonadSingleCh := func() string {
		if strings.HasPrefix(md.Name, md.ExposureMonadType.Obj().Name()) {
			return md.Name[len(md.ExposureMonadType.Obj().Name()):]
		}
		return md.TargetType.Obj().Name()[:1]
	}()

	if suffixName == "" {
		ctx.defineFunc("Lift"+targetMonadSingleCh, `

			func {{.funcname}}[{{.tpargs}}](a {{outer "A"}}) {{combined "A"}} {
				return {{.givenMap}}(a, {{pure "A"}})
			}
		`)
	} else {
		ctx.defineFunc("Lift"+suffixName, `

			func {{.funcname}}[{{.tpargs}}](a {{outer "A"}}) {{combined "A"}} {
				return {{.givenMap}}(a, {{pure "A"}})
			}
		`)
	}
	ctx.defineFunc("Map"+suffixName, `
		func {{.funcname}}[{{.tpargs}},B any](t {{combined "A"}}, f func(A) B) {{combined "B"}} {
			return {{.givenMap}}(t, func( ma {{inner "A"}} )  {{inner "B"}} {
				return {{flatmap "A" "B"}}(ma, func(a A) {{inner "B"}} {
					return {{pure "B"}}(f(a))
				})
			})
		}
	`)
	ctx.defineFunc("SubFlatMap"+suffixName, `
		func {{.funcname}}[{{.tpargs}},B any](t {{combined "A"}}, f func(A) {{inner "B"}}) {{combined "B"}} {
			return {{.givenMap}}(t, func( ma {{inner "A"}} )  {{inner "B"}} {
				return {{flatmap "A" "B"}}(ma, func(a A) {{inner "B"}} {
					return f(a)
				})
			})
		}
	`)

	if md.Sequence.Expr != nil {

		maptname := "Map" + targetMonadSingleCh
		if suffixName != "" {
			maptname = "Traverse" + suffixName
		}

		ctx.defineFunc(maptname, `
			func {{.funcname}}[{{.tpargs}}, B any](t {{combined "A"}}, f func(A) {{outer "B"}}) {{combined "B"}} {
				sequencef := {{sequence "B"}}
				return {{.givenFlatMap}}(Map{{.name}}(t,f), sequencef)
			}
		`)

		ctx.param["maptfunc"] = maptname

		ctx.defineFunc("FlatMap"+suffixName, `
			func {{.funcname}}[{{.tpargs}}, B any](t {{combined "A"}}, f func(A) {{combined "B"}}) {{combined "B"}} {

				flatten := func(v {{inner (inner "B")}}) {{inner "B"}} {
					return {{flatmap (inner "B") "B"}}(v , fp.Id)
				}

				return {{.givenMap}}({{.maptfunc}}(t, f), flatten)

			}
		`)
	}

	targName := privateName(md.Name)
	for _, t := range md.Transform {

		if sig, ok := t.TypeReference.Type.(*types.Signature); ok {

			argTypes := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
				var tpe string
				if sig.Variadic() && i == sig.Params().Len()-1 {
					st := t.Type().(*types.Slice)
					tpe = w.TypeName(md.Package, st.Elem())
				} else {
					tpe = w.TypeName(md.Package, t.Type())
				}
				return tpe
			})

			targetParam := md.TypeParm.String()
			targIdx, ok := seqFirst(seqFilter(iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) int {
				tpe := argTypes[i]
				if tpe == innertype(md.TypeParm.String()) {
					return i
				}
				return -1
			}), func(v int) bool { return v >= 0 }))

			if !ok {
				for i := 0; i < sig.Params().Len(); i++ {
					tpe := sig.Params().At(i)
					if gt, gtok := tpe.Type().(GenericType); gtok {
						//fmt.Printf("compare name %s %t\n", gt.Obj().Name(), gt.Obj().Name() == md.ExposureMonadType.Obj().Name())
						if gt.Obj() == md.ExposureMonadType.Obj() {
							targIdx = i
							targetParam = seqMakeString(iterate(gt.TypeArgs().Len(), gt.TypeArgs().At, func(i int, t types.Type) string {
								return w.TypeName(md.Package, t)
							}), ",")
							ok = true
							break
						}
					}
				}
			}

			if ok {

				argTypeStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
					tpe := argTypes[i]

					if i == targIdx {
						tpe = combinedtype(targetParam)
						return fmt.Sprintf("%s %s", targName, tpe)

					}

					return fmt.Sprintf("%s %s", argName(i, t), tpe)
				})

				callArgs := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {

					if i == targIdx {
						return "insideValue"
					}

					return argName(i, t)
				})

				retType := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
					return w.TypeName(md.Package, t.Type())
				})

				tp := append(fixedParams, seqMap(t.TypeParams, func(v TypeReference) string {
					if p, ok := v.Type.(*types.TypeParam); ok {
						return fmt.Sprintf("%s %s", p.String(), w.TypeName(md.Package, p.Constraint()))
					}
					return ""
				})...)

				param["trans"] = t.Name
				param["args"] = seqMakeString(argTypeStr, ",")
				param["callArgs"] = seqMakeString(callArgs, ",")

				param["targName"] = targName
				param["transExpr"] = exprString(t.TypeReference.Expr)
				param["tparams"] = seqMakeString(tp, ",")

				param["targ"] = targetParam
				ctx := genFuncContext{
					w:               w,
					definedFunction: definedFunc,
					funcs:           funcs,
					param:           param,
				}
				if len(retType) > 0 {
					if len(retType) == 1 {
						param["retType"] = retType[0]

						if gt, gtok := sig.Results().At(0).Type().(GenericType); gtok && gt.Obj() == md.TargetType.Obj() {
							ctx.defineFunc(t.Name+suffixName, `
								func {{.trans}}{{.name}}[{{.tparams}}]({{.args}}) {{.retType}} {
									return {{.givenFlatMap}}({{.targName}}, func(insideValue {{inner (.targ)}}) {{.retType}} {
										return {{.transExpr}}({{.callArgs}})
									} )
								}
							`)
						} else {
							ctx.defineFunc(t.Name+suffixName, `
								func {{.trans}}{{.name}}[{{.tparams}}]({{.args}}) {{outer (.retType)}} {
									return {{.givenMap}}({{.targName}}, func(insideValue {{inner (.targ)}}) {{.retType}} {
										return {{.transExpr}}({{.callArgs}})
									} )
								}
							`)
						}
					} else {
						param["retType"] = fmt.Sprintf("fp.Tuple%d[%s]", len(retType), seqMakeString(retType, ","))
						param["asTuple"] = fmt.Sprintf("as.Tuple%d", len(retType))
						w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/as", "as"))

						ctx.defineFunc(t.Name+suffixName, `
							func {{.trans}}{{.name}}[{{.tparams}}]({{.args}}) {{outer (.retType)}} {
								return {{.givenMap}}({{.targName}}, func(insideValue {{inner (.targ)}}) {{.retType}} {
									return {{.asTuple}}({{.transExpr}}({{.callArgs}}))
								} )
							}
						`)
					}
				} else {
					ctx.defineFunc(t.Name+suffixName, `
						func {{.trans}}{{.name}}[{{.tparams}}]({{.args}}) {
							{{.givenMap}}({{.targName}}, func(insideValue {{inner (.targ)}}) error {
								{{.transExpr}}({{.callArgs}})
								return nil
							} )
						}
					`)
				}
			} else {
				fmt.Printf("not generate %s\n", t.Name)
			}
		}

	}
}

func privateName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func argName(i int, t *types.Var) string {
	var name = t.Name()
	if name == "" {
		name = fmt.Sprintf("a%d", i+1)
	}
	return name
}
