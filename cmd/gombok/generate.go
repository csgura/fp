package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"strings"

	"slices"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"golang.org/x/tools/go/packages"
)

func hasMethod(v types.Type, name string) bool {

	if intf, ok := v.Underlying().(*types.Interface); ok {
		f := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) bool {
			return t.Name() == name
		})
		return f.Exists(fp.Id)
	}

	switch t := v.(type) {
	case *types.Named:
		f := iterate(t.NumMethods(), t.Method, func(i int, t *types.Func) bool {
			return t.Name() == name
		})
		return f.Exists(fp.Id)
	case *types.Interface:
		f := iterate(t.NumMethods(), t.Method, func(i int, t *types.Func) bool {
			return t.Name() == name
		})
		return f.Exists(fp.Id)
	}
	return false
}

func fillOption(ret genfp.GenerateAdaptorDirective, intf *types.Interface) (genfp.GenerateAdaptorDirective, error) {
	for i := 0; i < intf.NumMethods(); i++ {
		m := intf.Method(i)

		opt := ret.Methods[m.Name()]
		opt.Type = m

		for df, d := range ret.Delegate {

			if hasMethod(d.Type, m.Name()) {
				opt.DelegateField = df

			}
		}

		sig, ok := m.Type().(*types.Signature)
		if !ok {
			return ret, fmt.Errorf("type is not signature")
		}
		opt.Signature = sig
		if opt.Prefix == "" {
			if slices.Contains(ret.Getter, m.Name()) {
				opt.Prefix = "Get"

				if sig.Results().Len() == 1 {
					res := sig.Results().At(0)
					if res.Type().String() == "bool" && !strings.HasPrefix(m.Name(), "Is") {
						opt.Prefix = "Is"
					}
				}

			} else if slices.Contains(ret.EventHandler, m.Name()) {
				opt.Prefix = "On"
			} else {
				opt.Prefix = "Do"
			}
		}

		if opt.ValOverride == false {
			opt.ValOverride = slices.Contains(ret.ValOverride, m.Name())
			if opt.ValOverride && !slices.Contains(ret.Getter, m.Name()) {
				opt.OmitGetterIfValOverride = true
			}
		}

		if opt.DefaultImplExpr == nil && slices.Contains(ret.ZeroReturn, m.Name()) {
			opt.DefaultImplExpr = &ast.SelectorExpr{X: ast.NewIdent("genfp"), Sel: ast.NewIdent("ZeroReturn")}
		}

		ret.Methods[m.Name()] = opt
	}
	return ret, nil
}

func genGenerate() {
	pack := os.Getenv("GOPACKAGE")

	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	gentemplate := genfp.FindGenerateFromUntil(pkgs, "@fp.Generate")
	genadaptor := genfp.FindGenerateAdaptor(pkgs, "@fp.Generate")

	filelist := iterator.ToGoSet(mutable.MapOf(gentemplate).Keys().Concat(mutable.MapOf(genadaptor).Keys()))

	for file := range filelist {
		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range gentemplate[file] {
				for _, im := range gfu.Imports {
					w.GetImportedName(types.NewPackage(im.Package, im.Name))
				}

				w.Iteration(gfu.From, gfu.Until).Write(gfu.Template, map[string]any{})
			}

			for _, gad := range genadaptor[file] {
				if gad.ExtendsByEmbedding {
					gad.Extends = true
				}

				adaptorTypeName := gad.Name
				if adaptorTypeName == "" {
					adaptorTypeName = gad.Interface.Obj().Name() + "Adaptor"
				}

				superField := ""
				if gad.Extends {
					superField = "Extends"
				}
				intf := gad.Interface.Underlying().(*types.Interface)

				gad, _ = fillOption(gad, intf)
				fields := fieldAndImplOfInterfaceImpl(w, gad, gad.Interface, adaptorTypeName, superField)

				for _, i := range gad.ImplementsWith {
					if i.Type != nil {
						if intf, ok := i.Type.Underlying().(*types.Interface); ok {
							gad, _ = fillOption(gad, intf)
							fields = fields.Concat(fieldAndImplOfInterfaceImpl(w, gad, i.Type, adaptorTypeName, ""))
						}
					}
				}

				extends := ""
				if gad.Extends && gad.ExtendsByEmbedding == false {
					extends = "Extends " + w.TypeName(gad.Package.Types, gad.Interface)
				}

				fieldDecl := seq.Of(extends)

				delegateFields := iterator.Sort(iterator.FromMapKey(gad.Delegate), ord.Given[string]()).FilterNot(func(v string) bool {
					return isEmbeddingField(gad, v)
				})

				fieldDecl = fieldDecl.Concat(seq.Map(delegateFields, func(k string) string {
					tpe := gad.Delegate[k]
					return fmt.Sprintf("%s %s", k, w.TypeName(gad.Package.Types, tpe.Type))
				}))

				fieldDecl = fieldDecl.Concat(seq.Map(gad.Embedding, func(v genfp.TypeReference) string {
					if v.Type != nil {
						return w.TypeName(gad.Package.Types, v.Type)
					}
					return v.StringExpr
				}))

				fieldDecl = fieldDecl.Concat(seq.Map(fields, fp.Tuple2[string, string].Head).FilterNot(eq.GivenValue("")))

				fmt.Fprintf(w, `type %s struct {
					%s
				}
				`, adaptorTypeName, fieldDecl.MakeString("\n"))

				fmt.Fprintf(w, "%s", seq.Map(fields, fp.Tuple2[string, string].Tail).MakeString("\n"))
			}
		})
	}

}

func isEmbeddingField(gad genfp.GenerateAdaptorDirective, field string) bool {
	for _, e := range gad.Embedding {
		return seq.Last(strings.Split(e.StringExpr, ".")) == option.Some(field)
	}
	return false
}
func fieldAndImplOfInterfaceImpl(w genfp.Writer, gad genfp.GenerateAdaptorDirective, namedInterface types.Type, adaptorTypeName string, superField string) fp.Seq[fp.Tuple2[string, string]] {
	intf := namedInterface.Underlying().(*types.Interface)

	fields := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) fp.Tuple2[string, string] {
		opt := gad.Methods[t.Name()]
		sig := opt.Signature
		valName := fmt.Sprintf("Default%s", t.Name())
		cbName := fmt.Sprintf("%s%s", opt.Prefix, t.Name())
		if opt.Name != "" {
			cbName = fmt.Sprintf("%s%s", opt.Prefix, opt.Name)
			valName = fmt.Sprintf("Default%s", opt.Name)
		}

		selfarg := ""
		if gad.Self {
			selfarg = "self " + w.TypeName(gad.Package.Types, gad.Interface) + ","
		}

		argTypes := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) fp.Tuple2[string, types.Type] {
			return as.Tuple2(t.Name(), t.Type())
		})

		argStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
			if sig.Variadic() && i == sig.Params().Len()-1 {
				return fmt.Sprintf("%s...", t.Name())
			} else {
				return fmt.Sprintf("%s", t.Name())
			}
		}).MakeString(",")

		argTypeStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
			if sig.Variadic() && i == sig.Params().Len()-1 {
				st := t.Type().(*types.Slice)
				return fmt.Sprintf("%s ...%s", t.Name(), w.TypeName(gad.Package.Types, st.Elem()))
			}
			return fmt.Sprintf("%s %s", t.Name(), w.TypeName(gad.Package.Types, t.Type()))
		}).MakeString(",")

		resstr := ""
		if sig.Results().Len() > 0 {
			resstr = "(" + iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
				return w.TypeName(gad.Package.Types, t.Type())
			}).MakeString(",") + ")"
		}

		implArgs := argTypeStr
		if gad.Self {
			implArgs = "self " + w.TypeName(gad.Package.Types, gad.Interface) + "," + argTypeStr
		}

		implName := func() string {
			if gad.Self {
				return t.Name() + "Impl"
			}
			return t.Name()
		}()

		withReturn := func(lastPos bool, fmtstr string, args ...any) string {
			e := fmt.Sprintf(fmtstr, args...)
			if sig.Results().Len() == 0 {
				if lastPos {
					return e
				}
				return fmt.Sprintf("%s\nreturn", e)
			}
			return "return " + e
		}

		if opt.DelegateField != "" {
			cbName = opt.DelegateField
			if isEmbeddingField(gad, opt.DelegateField) {
				if gad.Self && gad.ExtendsByEmbedding {
					return as.Tuple("", fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
						`, adaptorTypeName, t.Name(), argTypeStr, resstr,
						withReturn(true, `r.%sImpl(r, %s)`, t.Name(), argStr),
					))
				}
				return as.Tuple("", "")
			}
		}

		callSuperImpl := func(field string) string {

			if gad.Self {
				return fmt.Sprintf(`type impl interface {
										%s(%s) %s
									}

									if super, ok := r.%s.(impl); ok {
										%s 
									}
									`, implName, implArgs, resstr,
					field,
					withReturn(false, "super.%s(self, %s)", implName, argStr),
				)
			}
			return ""

		}

		callSupercheck := func(field string) string {
			return fmt.Sprintf(`
					%sif super , ok := r.%s.(%s); ok {
						%s
					}
				`, callSuperImpl(field), field, w.TypeName(gad.Package.Types, namedInterface),
				withReturn(false, "super.%s(%s)", t.Name(), argStr),
			)
		}

		callcb := func() string {
			if opt.DelegateField != "" {

				return fmt.Sprintf("%s%s", callSuperImpl(opt.DelegateField), withReturn(false, `r.%s.%s(%s)`, opt.DelegateField, t.Name(), argStr))
			}

			if gad.Self {
				return withReturn(false, `r.%s(self,%s)`, cbName, argStr)
			}
			return withReturn(false, `r.%s(%s)`, cbName, argStr)

		}()

		valoverride := opt.ValOverride && sig.Params().Len() == 0 && sig.Results().Len() == 1
		defaultValExpr := ""
		cbExpr := fmt.Sprintf(`if r.%s != nil {
								%s
							}`, cbName,
			callcb)

		cbfield := fmt.Sprintf("%s func(%s%s) %s", cbName, selfarg, argTypeStr, resstr)
		if opt.DelegateField != "" {
			cbfield = ""
		}

		extendscb := func() string {
			if superField != "" && gad.ExtendsByEmbedding == false {
				return fmt.Sprintf(`
					if r.%s != nil {
						%s%s
					}
				`, superField, callSuperImpl(superField), withReturn(false, "r.%s.%s(%s)", superField, t.Name(), argStr),
				)
			} else if gad.Extends {
				return fmt.Sprintf(`
					if r.Extends != nil {
						%s
					}
				`, callSupercheck("Extends"))
			}
			return ""

		}()

		valOverrideOnly := false
		if valoverride {

			zeroVal := w.ZeroExpr(gad.Package.Types, sig.Results().At(0).Type())
			if zeroVal == "nil" {
				defaultValExpr = fmt.Sprintf(`if r.%s != %s {
									return r.%s
								}
						
						`, valName, zeroVal,
					valName)
			} else if (zeroVal == "0" || zeroVal == `""`) && (opt.OmitGetterIfValOverride == false || extendscb != "") {
				defaultValExpr = fmt.Sprintf(`if r.%s != %s {
									return r.%s
								}
						
						`, valName, zeroVal,
					valName)
			} else if types.Comparable(sig.Results().At(0).Type()) && (opt.OmitGetterIfValOverride == false || extendscb != "") {
				defaultValExpr = fmt.Sprintf(`
							var _zero %s
							if r.%s != _zero {
								return r.%s
							}
						
						`, w.TypeName(gad.Package.Types, sig.Results().At(0).Type()),
					valName,
					valName)
			} else {
				defaultValExpr = fmt.Sprintf("return r.%s", valName)
				opt.OmitGetterIfValOverride = true
				valOverrideOnly = true
			}

			if opt.OmitGetterIfValOverride {
				cbfield = fmt.Sprintf("%s %s", valName, w.TypeName(gad.Package.Types, sig.Results().At(0).Type()))
				cbExpr = ""

			} else {
				cbfield = fmt.Sprintf("%s %s\n%s", valName, w.TypeName(gad.Package.Types, sig.Results().At(0).Type()), cbfield)
			}

		}

		if opt.Private {
			cbfield = ""
		}

		defaultcb := func() string {
			if matchSelExpr(opt.DefaultImplExpr, "genfp", "ZeroReturn") {
				if sig.Results().Len() == 0 {
					return ""
				}
				zeroval := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
					return w.ZeroExpr(gad.Package.Types, t.Type())
				}).MakeString(",")
				if zeroval != "" {
					return fmt.Sprintf(`return %s`, zeroval)
				}
				return "return"
			} else if opt.DefaultImplExpr != nil {

				for _, i := range opt.DefaultImplImports {
					w.AddImport(types.NewPackage(i.Package, i.Name))
				}

				if opt.DefaultImplSignature != nil {
					os := opt.DefaultImplSignature

					defImplArgs := iterate(os.Params().Len(), os.Params().At, func(i int, t *types.Var) types.Type {
						return t.Type()
					})

					availableArgs := func() fp.Seq[fp.Tuple2[string, types.Type]] {
						if gad.Self {
							return seq.Concat(as.Tuple[string, types.Type]("self", gad.Interface), argTypes)
						}
						return seq.Concat(as.Tuple[string, types.Type]("r", gad.Interface), argTypes)
					}()

					type CallArgs struct {
						avail fp.Seq[fp.Tuple2[string, types.Type]]
						args  fp.Seq[string]
					}

					args := seq.FoldTry(defImplArgs, CallArgs{avail: availableArgs}, func(args CallArgs, tp types.Type) fp.Try[CallArgs] {
						init, tail := iterator.Span(iterator.FromSeq(args.avail), func(t fp.Tuple2[string, types.Type]) bool {
							return t.I2.String() != tp.String()
						})

						arg := tail.NextOption()
						if arg.IsDefined() {
							return try.Success(CallArgs{init.Concat(tail).ToSeq(), append(args.args, arg.Get().I1)})
						}
						return try.Failure[CallArgs](fp.Error(400, "can't find proper args for type %s", tp.String()))
					})

					if args.IsSuccess() {
						if fl, ok := opt.DefaultImplExpr.(*ast.FuncLit); ok {
							fs := token.NewFileSet()

							buf := &bytes.Buffer{}

							printer.Fprint(buf, fs, fl)
							return withReturn(true, "%s(%s)", buf.String(), args.Get().args.MakeString(","))
						} else {

							return withReturn(true, `%s(%s)`, types.ExprString(opt.DefaultImplExpr), args.Get().args.MakeString(","))

						}
					} else {
						fmt.Printf("err : %s\n", args.Failed().Get())
					}
					if gad.Self {
						e := types.ExprString(opt.DefaultImplExpr)
						return withReturn(true, `%s(self, %s)`, e, argStr)
					} else {
						e := types.ExprString(opt.DefaultImplExpr)
						return withReturn(true, `%s(r, %s)`, e, argStr)
					}
				}

				fs := token.NewFileSet()

				buf := &bytes.Buffer{}

				printer.Fprint(buf, fs, opt.DefaultImplExpr)
				if sig.Results().Len() > 1 {
					zeroval := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
						return w.ZeroExpr(gad.Package.Types, t.Type())
					}).Drop(1).MakeString(",")
					return "return " + buf.String() + "," + zeroval
				} else {
					return "return " + buf.String()
				}

			}

			return fmt.Sprintf(`panic("%s.%s not implemented")`, adaptorTypeName, t.Name())
		}()

		impl := fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
							%s
							%s
							%s
						}
						`, adaptorTypeName, implName, implArgs, resstr,
			defaultValExpr,
			cbExpr,
			extendscb,
			defaultcb,
		)

		if opt.Private {
			impl = fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
						`, adaptorTypeName, implName, implArgs, resstr,
				defaultcb)
		} else if valOverrideOnly {
			impl = fmt.Sprintf(`
							func (r *%s) %s(%s) %s {
								%s
							}
						`, adaptorTypeName, implName, implArgs, resstr,
				defaultValExpr,
			)
		}

		if gad.Self {
			impl = fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
					`, adaptorTypeName, t.Name(), argTypeStr, resstr,
				withReturn(true, "r.%sImpl(r,%s)", t.Name(), argStr),
			) + impl

		}

		return as.Tuple(cbfield, impl)

	})

	return fields
}

func matchSelExpr(expr ast.Expr, exp ...string) bool {
	if len(exp) == 0 {
		return false
	}
	init := exp[0 : len(exp)-1]
	last := exp[len(exp)-1]
	if sel, ok := expr.(*ast.SelectorExpr); ok {
		if sel.Sel.Name == last {
			switch t := sel.X.(type) {
			case *ast.SelectorExpr:
				return matchSelExpr(t, init...)
			case *ast.Ident:
				if len(init) == 1 {
					return t.Name == init[0]
				}
			}
		}
	}
	return false
}
