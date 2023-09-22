package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"golang.org/x/tools/go/packages"
)

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
				adaptorTypeName := gad.Name
				if adaptorTypeName == "" {
					adaptorTypeName = gad.Interface.Obj().Name() + "Adaptor"
				}

				intf := gad.Interface.Underlying().(*types.Interface)

				fields := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) fp.Tuple2[string, string] {
					opt := gad.Methods[t.Name()]
					sig := opt.Signature
					valName := fmt.Sprintf("Default%s", t.Name())
					cbName := fmt.Sprintf("%s%s", opt.Prefix, t.Name())
					if opt.Name != "" {
						cbName = opt.Name
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

					withReturn := func(fmtstr string, args ...any) string {
						e := fmt.Sprintf(fmtstr, args...)
						if sig.Results().Len() == 0 {
							return fmt.Sprintf("%s\nreturn", e)
						}
						return "return " + e
					}

					callcb := func() string {
						if gad.Self == false {
							return withReturn(`r.%s(%s)`, cbName, argStr)

						}
						return withReturn(`r.%s(self,%s)`, cbName, argStr)
					}()

					valoverride := opt.ValOverride && sig.Params().Len() == 0 && sig.Results().Len() == 1
					defaultValExpr := ""
					cbExpr := fmt.Sprintf(`if r.%s != nil {
								%s
							}`, cbName,
						callcb)

					cbfield := fmt.Sprintf("%s func(%s%s) %s", cbName, selfarg, argTypeStr, resstr)

					valOverrideOnly := false
					if valoverride {

						zeroVal := w.ZeroExpr(gad.Package.Types, sig.Results().At(0).Type())
						if zeroVal == "nil" {
							defaultValExpr = fmt.Sprintf(`if r.%s != %s {
									return r.%s
								}
						
						`, valName, zeroVal,
								valName)
						} else if zeroVal == "0" || zeroVal == `""` && (opt.OmitGetterIfValOverride == false || gad.Extends == true) {
							defaultValExpr = fmt.Sprintf(`if r.%s != %s {
									return r.%s
								}
						
						`, valName, zeroVal,
								valName)
						} else if types.Comparable(sig.Results().At(0).Type()) && (opt.OmitGetterIfValOverride == false || gad.Extends == true) {
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

					implName := func() string {
						if gad.Self {
							return t.Name() + "Impl"
						}
						return t.Name()
					}()

					implArgs := argTypeStr
					if gad.Self {
						implArgs = "self " + w.TypeName(gad.Package.Types, gad.Interface) + "," + argTypeStr
					}

					extendscb := func() string {
						if gad.Extends {
							superexpr := ""
							if gad.Self {
								superexpr = fmt.Sprintf(`type impl interface {
										%s(%s) %s
									}

									if super, ok := r.Extends.(impl); ok {
										%s 
									}
									`, implName, implArgs, resstr,
									withReturn("super.%s(self, %s)", implName, argStr),
								)
							}
							return fmt.Sprintf(`
								if r.Extends != nil {
									%s%s 
								}
							`,
								superexpr,
								withReturn("r.Extends.%s(%s)", t.Name(), argStr),
							)
						}
						return ""

					}()

					defaultcb := func() string {
						if matchSelExpr(opt.DefaultImplExpr, "genfp", "ZeroReturn") {
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
										return withReturn("%s(%s)", buf.String(), args.Get().args.MakeString(","))
									} else {

										return withReturn(`%s(%s)`, types.ExprString(opt.DefaultImplExpr), args.Get().args.MakeString(","))

									}
								} else {
									fmt.Printf("err : %s\n", args.Failed().Get())
								}
								if gad.Self {
									e := types.ExprString(opt.DefaultImplExpr)
									return withReturn(`%s(self, %s)`, e, argStr)
								} else {
									e := types.ExprString(opt.DefaultImplExpr)
									return withReturn(`%s(r, %s)`, e, argStr)
								}
							}

							fs := token.NewFileSet()

							buf := &bytes.Buffer{}

							printer.Fprint(buf, fs, opt.DefaultImplExpr)
							return "return " + buf.String()

						}

						return fmt.Sprintf(`panic("not implemented")`)
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

					if valOverrideOnly {
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
							withReturn("r.%sImpl(r,%s)", t.Name(), argStr),
						) + impl

					}

					return as.Tuple(cbfield, impl)

				})

				extends := ""
				if gad.Extends {
					extends = "Extends " + w.TypeName(gad.Package.Types, gad.Interface)
				}

				fmt.Fprintf(w, `type %s struct {
					%s
					%s
				}
				`, adaptorTypeName, extends, seq.Map(fields, fp.Tuple2[string, string].Head).MakeString("\n"))

				fmt.Fprintf(w, "%s", seq.Map(fields, fp.Tuple2[string, string].Tail).MakeString("\n"))
			}
		})
	}

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
