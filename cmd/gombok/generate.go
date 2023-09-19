package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/seq"
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

					selfarg := ""
					if gad.Self {
						selfarg = "self " + w.TypeName(gad.Package.Types, gad.Interface) + ","
					}

					argStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
						return fmt.Sprintf("%s", t.Name())
					}).MakeString(",")

					argTypeStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
						return fmt.Sprintf("%s %s", t.Name(), w.TypeName(gad.Package.Types, t.Type()))
					}).MakeString(",")

					resstr := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
						return w.TypeName(gad.Package.Types, t.Type())
					}).MakeString(",")

					cb := fmt.Sprintf("%s%s func(%s%s) (%s)", opt.Prefix, t.Name(), selfarg, argTypeStr, resstr)

					retExpr := "return "
					voidExpr := ""
					if sig.Results().Len() == 0 {
						retExpr = ""
						voidExpr = "return"
					}

					callcb := func() string {
						if gad.Self == false {
							return fmt.Sprintf(`
								%s r.%s%s(%s)
								%s
							`, retExpr, opt.Prefix, t.Name(), argStr,
								voidExpr,
							)

						}
						return fmt.Sprintf(`
							%s r.%s%s(self,%s)
							%s
							`, retExpr, opt.Prefix, t.Name(), argStr,
							voidExpr)
					}()

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
								superexpr = fmt.Sprintf(`
									type impl interface {
										%s(%s) (%s)
									}
									if super, ok := r.Extends.(impl); ok {
										%s super.%s(self, %s)
										%s
									}
								`, implName, implArgs, resstr,
									retExpr, implName, argStr,
									voidExpr,
								)
							}
							return fmt.Sprintf(`
								if r.Extends != nil {
									%s
									%s r.Extends.%s(%s)
									%s
								}
							`,
								superexpr,
								retExpr, t.Name(), argStr,
								voidExpr,
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
							if gad.Self {

								if opt.DefaultImplSignature != nil {
									os := opt.DefaultImplSignature
									if os.Params().Len() != sig.Params().Len() && os.Params().Len() > 0 {

										if os.Params().At(0).Type() == gad.Interface {
											return fmt.Sprintf(`%s %s(self,%s)
										%s`, retExpr, types.ExprString(opt.DefaultImplExpr), argStr, voidExpr)
										}
									}
								}

							}

							e := types.ExprString(opt.DefaultImplExpr)
							return fmt.Sprintf(`%s %s(%s)
							%s`, retExpr, e, argStr, voidExpr)
						}
						return fmt.Sprintf(`panic("not implemented")`)
					}()

					impl := fmt.Sprintf(`
						func (r *%s) %s(%s) (%s) {
							if r.%s%s != nil {
								%s
							}
							%s
							%s
						}
					`, adaptorTypeName, implName, implArgs, resstr,
						opt.Prefix, t.Name(),
						callcb,
						extendscb,
						defaultcb,
					)

					if gad.Self {
						impl = fmt.Sprintf(`
						func (r *%s) %s(%s) (%s) {
							%s r.%sImpl(r,%s)
						}
					`, adaptorTypeName, t.Name(), argTypeStr, resstr,
							retExpr, t.Name(), argStr,
						) + impl

					}

					return as.Tuple(cb, impl)

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
