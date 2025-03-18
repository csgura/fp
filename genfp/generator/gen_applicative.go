package generator

import (
	"fmt"
	"go/types"

	"github.com/csgura/fp/genfp"
)

func WriteApplicativeFunctions(w Writer, md GenerateApplicative, definedFunc map[string]bool) {

	name := md.TargetType.Obj().Name()

	rettype := NameParamReplaced(w, md.Package, md.TargetType, md.TypeParam)
	if md.TargetAlias != nil {
		name = md.TargetAlias.Obj().Name()
		rettype = NameParamReplaced(w, md.Package, md.TargetAlias, md.TypeParam)
	}

	funcs := map[string]any{
		"monad": rettype,
		"name": func() string {
			return name
		},
	}
	param := map[string]any{}
	ctx := genFuncContext{
		w:               w,
		definedFunction: definedFunc,
		funcs:           funcs,
		param:           param,
	}

	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp", "fp"))
	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/curried", "curried"))

	tmpl := seqMakeString(seqMap(md.Mapper, func(v Mapping) string {

		if sig, ok := v.Mapper.TypeReference.Type.(*types.Signature); ok {
			in := sig.Params()
			ret := sig.Results()
			transExpr := CallFunc(w, v.Mapper.TypeReference)
			if in.Len() == 1 && ret.Len() == 1 {
				inType := in.At(0).Type()
				if gt, ok := inType.(GenericType); ok {
					mtype := NameParamReplaced(w, md.Package, gt, md.TypeParam)

					return fmt.Sprintf(`
						{{template "Receiver" .}} %s(a %s) {{template "Next" .}} {
							return r.Ap(%s(a))
						}
						`, v.Prefix, mtype("A1"), transExpr(replaceParam{
						md.TypeParam.String(): "A1",
					})) + fmt.Sprintf(`
					{{template "Receiver" .}} Lazy%s(a func() %s) {{template "Next" .}} {
							return r.ApFunc(func() {{monad "A1"}} { 
							return %s(a()) 
							})
						}
					`, v.Prefix, mtype("A1"), transExpr(replaceParam{
						md.TypeParam.String(): "A1",
					}))
				} else {
					return fmt.Sprintf(`
						{{template "Receiver" .}} %s(a A1) {{template "Next" .}} {
							return r.Ap(%s(a))
						}
					`, v.Prefix, transExpr(replaceParam{
						md.TypeParam.String(): "A1",
					})) + fmt.Sprintf(`
						{{template "Receiver" .}} Lazy%s(a func() A1) {{template "Next" .}} {
							return r.ApFunc(func() {{monad "A1"}} {
								return %s(a())
							})
						}
					`, v.Prefix, transExpr(replaceParam{
						md.TypeParam.String(): "A1",
					}))
				}
			}
		}
		return ""
	}), "\n")

	ctx.defineFunc(`Applicative1`, `
		{{define "Receiver"}}func (r ApplicativeFunctor1[A1, R]){{end}}
		{{define "Next"}}{{monad "R"}}{{end}}

		type ApplicativeFunctor1[A1, R any] struct {
			fn {{monad "fp.Func1[A1,R]"}}
		}

		{{template "Receiver" .}} Ap(a {{monad "A1"}}) {{template "Next" .}} {
			return Ap(r.fn, a)
		}

		{{template "Receiver" .}} ApFunc(a func() {{monad "A1"}}) {{template "Next" .}} {
			return ApFunc(r.fn, a)
		}

		{{template "Receiver" .}} {{name}}(a {{monad "A1"}}) {{template "Next" .}} {
			return Ap(r.fn, a)
		}

		{{template "Receiver" .}} Lazy{{name}}(a func() {{monad "A1"}}) {{template "Next" .}} {
			return ApFunc(r.fn, a)
		}

		func {{.funcname}}[A1, R any](fn func(A1) R) ApplicativeFunctor1[A1, R]  {
			return ApplicativeFunctor1[A1, R]{Pure(fn)}
		}
	`+tmpl)

	ctx.defineFuncs(2, genfp.MaxFunc, `Applicative{{.N}}`, `
		{{define "Receiver"}}func (r ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]){{end}}
		{{define "Next"}}ApplicativeFunctor{{dec .N}}[{{TypeArgs 2 .N}}, R]{{end}}

		type ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R any] struct {
			fn {{monad (CurriedFunc 1 .N "R")}}
		}

		{{template "Receiver" .}} Ap(a {{monad "A1"}}) {{template "Next" .}} {
			return {{template "Next" .}}{Ap(r.fn, a)}
		}

		{{template "Receiver" .}} ApFunc(a func() {{monad "A1"}}) {{template "Next" .}} {
			return {{template "Next" .}}{ApFunc(r.fn, a)}
		}

		{{template "Receiver" .}} {{name}}(a {{monad "A1"}}) {{template "Next" .}} {
			return {{template "Next" .}}{Ap(r.fn, a)}
		}

		{{template "Receiver" .}} Lazy{{name}}(a func() {{monad "A1"}}) {{template "Next" .}} {
			return {{template "Next" .}}{ApFunc(r.fn, a)}
		}

		func {{.funcname}}[{{TypeArgs 1 .N}}, R any](fn func({{TypeArgs 1 .N}}) R) ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]  {
			return ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]{Pure(curried.Func{{.N}}(fn))}
		}
	`+tmpl)

}
