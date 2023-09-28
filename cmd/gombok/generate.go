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

		for _, d := range ret.Delegate {

			if hasMethod(d.TypeOf.Type, m.Name()) {
				opt.Delegate = &genfp.DelegateDirective{
					TypeOf: d.TypeOf,
					Field:  d.Field,
				}

				if opt.Method == "" {
					opt.Private = true
				}
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

func typeDecl(pk *types.Package, w genfp.Writer, t genfp.TypeReference) string {
	if t.Type != nil {
		return w.TypeName(pk, t.Type)
	}
	return t.StringExpr
}

func generateAdaptor(w genfp.Writer, gad genfp.GenerateAdaptorDirective) {

	adaptorTypeName := gad.Name
	if adaptorTypeName == "" {
		adaptorTypeName = gad.Interface.Obj().Name() + "Adaptor"
	}

	fieldSet := fp.Map[string, genfp.TypeReference]{}
	var fieldList []string
	methodSet := fp.Set[string]{}

	for _, i := range gad.EmbeddingInterface {
		efield := seq.Last(strings.Split(i.StringExpr, "."))
		if efield.IsDefined() {
			if !fieldSet.Contains(efield.Get()) {
				fieldList = append(fieldList, fmt.Sprintf("%s", typeDecl(gad.Package.Types, w, i)))
			}
			fieldSet = fieldSet.Updated(efield.Get(), i)

			contains := slices.ContainsFunc(gad.Delegate, func(d genfp.DelegateDirective) bool {
				return d.Field == efield.Get()
			})
			if !contains {
				gad.Delegate = append(gad.Delegate, genfp.DelegateDirective{
					TypeOf: i,
					Field:  efield.Get(),
				})
			}
		}
	}

	for _, e := range gad.Embedding {
		efield := seq.Last(strings.Split(e.StringExpr, "."))
		if efield.IsDefined() {
			if !fieldSet.Contains(efield.Get()) {
				fieldList = append(fieldList, fmt.Sprintf("%s", typeDecl(gad.Package.Types, w, e)))
			}

			fieldSet = fieldSet.Updated(efield.Get(), e)
		}
	}

	for k, e := range gad.ExtendsWith {
		if !fieldSet.Contains(k) {
			fieldList = append(fieldList, fmt.Sprintf("%s %s", k, typeDecl(gad.Package.Types, w, e)))
		}
		fieldSet = fieldSet.Updated(k, e)
	}

	i1 := iterator.Map(iterator.FromSeq(gad.Delegate), func(v genfp.DelegateDirective) string {
		return v.Field
	})

	for _, d := range gad.Delegate {
		exists := as.Seq(gad.ImplementsWith).Exists(func(v genfp.TypeReference) bool {
			return v.StringExpr == d.TypeOf.StringExpr
		})
		if !exists {
			gad.ImplementsWith = append(gad.ImplementsWith, d.TypeOf)
		}
	}

	delegateFields := iterator.Sort(i1, ord.Given[string]()).FilterNot(fieldSet.Contains)

	for _, fn := range delegateFields {

		d := as.Seq(gad.Delegate).Find(func(v genfp.DelegateDirective) bool {
			return v.Field == fn
		}).Get()

		fieldSet = fieldSet.Updated(fn, d.TypeOf)
		fieldList = append(fieldList, fmt.Sprintf("%s %s", fn, typeDecl(gad.Package.Types, w, d.TypeOf)))

	}

	if gad.ExtendsByEmbedding {
		gad.Extends = true
	}

	superField := ""
	if gad.Extends {
		superField = "Extends"
	}

	intf := gad.Interface.Underlying().(*types.Interface)

	gad, _ = fillOption(gad, intf)
	fields, methodSet := fieldAndImplOfInterfaceImpl2(w, gad, gad.Interface, adaptorTypeName, superField, fieldSet, methodSet)

	for _, i := range gad.ImplementsWith {
		if i.Type != nil {
			if intf, ok := i.Type.Underlying().(*types.Interface); ok {
				gad, _ = fillOption(gad, intf)
				af, ms := fieldAndImplOfInterfaceImpl2(w, gad, i.Type, adaptorTypeName, "", fieldSet, methodSet)
				fields = fields.Concat(af)
				methodSet = ms
			}
		}
	}

	extends := ""
	if gad.Extends && gad.ExtendsByEmbedding == false {
		extends = "Extends " + w.TypeName(gad.Package.Types, gad.Interface)
	}

	fieldDecl := seq.Of(extends)

	fieldDecl = fieldDecl.Concat(fieldList)

	fieldDecl = fieldDecl.Concat(seq.Map(fields, fp.Tuple2[string, string].Head).FilterNot(eq.GivenValue("")))

	fmt.Fprintf(w, `type %s struct {
					%s
				}
				`, adaptorTypeName, fieldDecl.MakeString("\n"))

	fmt.Fprintf(w, "%s", seq.Map(fields, fp.Tuple2[string, string].Tail).MakeString("\n"))
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
				generateAdaptor(w, gad)
			}
		})
	}

}

func isEmbeddingField(gad genfp.GenerateAdaptorDirective, field string) bool {
	is := func(e []genfp.TypeReference) bool {
		for _, e := range e {
			if seq.Last(strings.Split(e.StringExpr, ".")) == option.Some(field) {
				return true
			}
		}
		return false
	}

	return is(gad.Embedding) || is(gad.EmbeddingInterface)
}

type implContext struct {
	w               genfp.Writer
	gad             genfp.GenerateAdaptorDirective
	namedInterface  types.Type
	adaptorTypeName string
	superField      string
	t               *types.Func
	opt             genfp.ImplOptionDirective
	valName         string
	cbName          string
	selfarg         string
	argTypes        fp.Seq[fp.Tuple2[string, types.Type]]
	argStr          string
	argTypeStr      string
	resstr          string
	implArgs        string
	fieldMap        fp.Map[string, genfp.TypeReference]
}

func (r *implContext) withReturn(lastPos bool, fmtstr string, args ...any) string {
	e := fmt.Sprintf(fmtstr, args...)
	sig := r.opt.Signature
	if sig.Results().Len() == 0 {
		if lastPos {
			return e
		}
		return fmt.Sprintf("%s\nreturn", e)
	}
	return "return " + e
}

func (r *implContext) callCb() fp.Option[string] {
	if r.opt.Private {
		return option.None[string]()
	}

	if r.isValOverride() && r.opt.OmitGetterIfValOverride {
		return option.None[string]()
	}

	gad := r.gad

	callcb := func() string {

		if gad.Self {
			return r.withReturn(false, `r.%s(self,%s)`, r.cbName, r.argStr)
		}
		return r.withReturn(false, `r.%s(%s)`, r.cbName, r.argStr)

	}()

	return option.Some(
		fmt.Sprintf(`if r.%s != nil {
						%s
					}`,
			r.cbName,
			callcb))

}

func (r *implContext) callSuperImpl(field string) string {
	gad := r.gad

	fieldTpe := r.fieldMap.Get(r.superField)
	if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
		if !types.IsInterface(fieldTpe.Get().Type.Underlying()) {
			return ""
		}
	} else {
		fmt.Printf("field %s not defined at fieldMap\n", field)
	}

	implName := func() string {
		if gad.Self {
			return r.t.Name() + "Impl"
		}
		return r.t.Name()
	}()

	if gad.Self {
		return fmt.Sprintf(`type impl interface {
										%s(%s) %s
									}

									if super, ok := r.%s.(impl); ok {
										%s 
									}
									`, implName, r.implArgs, r.resstr,
			field,
			r.withReturn(false, "super.%s(self, %s)", implName, r.argStr),
		)
	}
	return ""

}

func (r *implContext) callExtends() (fp.Option[string], bool) {
	gad := r.gad

	cbNilCheck := true

	fieldTpe := r.fieldMap.Get(r.superField)
	if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
		zeroval := r.w.ZeroExpr(gad.Package.Types, fieldTpe.Get().Type)
		if zeroval != "nil" {
			cbNilCheck = false
		}
	}

	callSupercheck := func(field string) string {
		if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
			if !types.IsInterface(fieldTpe.Get().Type) {
				return ""
			}
		}
		return fmt.Sprintf(`
					%sif super , ok := r.%s.(%s); ok {
						%s
					}
				`, r.callSuperImpl(field), field, r.w.TypeName(gad.Package.Types, r.namedInterface),
			r.withReturn(false, "super.%s(%s)", r.t.Name(), r.argStr),
		)
	}

	if r.superField != "" && gad.ExtendsByEmbedding == false {

		if cbNilCheck {
			return option.Some(fmt.Sprintf(`
					if r.%s != nil {
						%s%s
					}
				`, r.superField, r.callSuperImpl(r.superField), r.withReturn(false, "r.%s.%s(%s)", r.superField, r.t.Name(), r.argStr),
			)), false
		}
		return option.Some(fmt.Sprintf(`
						%s%s
				`, r.callSuperImpl(r.superField), r.withReturn(false, "r.%s.%s(%s)", r.superField, r.t.Name(), r.argStr),
		)), true
	} else if gad.Extends {
		return option.Some(fmt.Sprintf(`
					if r.Extends != nil {
						%s
					}
				`, callSupercheck("Extends"))), false
	}
	return option.None[string](), false
}

func (r *implContext) adaptorFields() (fp.Option[string], fp.Option[string]) {
	opt := r.opt

	gad := r.gad
	sig := r.opt.Signature

	cbfield := option.Some(fmt.Sprintf("%s func(%s%s) %s", r.cbName, r.selfarg, r.argTypeStr, r.resstr))

	cbfield = cbfield.FilterNot(func(v string) bool { return opt.Delegate != nil })

	cbfield = cbfield.FilterNot(func(v string) bool {
		return r.isValOverride() && opt.OmitGetterIfValOverride
	})

	cbfield = cbfield.FilterNot(fp.Const[string](opt.Private))

	defaultField := option.Map(option.Of(r.isValOverride()).Filter(fp.Id), func(v bool) string {
		return fmt.Sprintf("%s %s", r.valName, r.w.TypeName(gad.Package.Types, sig.Results().At(0).Type()))
	})

	return defaultField, cbfield
}

func (r *implContext) isValOverride() bool {
	opt := r.opt
	sig := r.opt.Signature
	return opt.ValOverride && sig.Params().Len() == 0 && sig.Results().Len() == 1
}
func (r *implContext) valOverride(defaultImpl bool) (fp.Option[string], bool) {
	sig := r.opt.Signature
	w := r.w
	gad := r.gad
	valName := r.valName
	valoverride := r.isValOverride()
	if valoverride {
		zeroVal := w.ZeroExpr(gad.Package.Types, sig.Results().At(0).Type())
		if zeroVal == "nil" {
			ret := fmt.Sprintf(`if r.%s != %s {
								return r.%s
							}
					
					`, valName, zeroVal,
				valName)
			return option.Some(ret), false
		} else if (zeroVal == "0" || zeroVal == `""`) && defaultImpl {
			ret := fmt.Sprintf(`if r.%s != %s {
								return r.%s
							}
					
					`, valName, zeroVal,
				valName)
			return option.Some(ret), false

		} else if types.Comparable(sig.Results().At(0).Type()) && defaultImpl {
			ret := fmt.Sprintf(`
						var _zero %s
						if r.%s != _zero {
							return r.%s
						}
					
					`, w.TypeName(gad.Package.Types, sig.Results().At(0).Type()),
				valName,
				valName)
			return option.Some(ret), false

		} else {
			ret := fmt.Sprintf("return r.%s", valName)
			// opt.OmitGetterIfValOverride = true
			return option.Some(ret), true
		}

		// if opt.OmitGetterIfValOverride {
		// 	cbfield = fmt.Sprintf("%s %s", valName, w.TypeName(gad.Package.Types, sig.Results().At(0).Type()))
		// 	cbExpr = ""

		// } else {
		// 	cbfield = fmt.Sprintf("%s %s\n%s", valName, w.TypeName(gad.Package.Types, sig.Results().At(0).Type()), cbfield)
		// }
	}
	return option.None[string](), false

}

func (r *implContext) defaultImpl() fp.Option[string] {
	opt := r.opt
	sig := opt.Signature
	gad := r.gad
	w := r.w

	if matchSelExpr(opt.DefaultImplExpr, "genfp", "ZeroReturn") {
		if sig.Results().Len() == 0 {
			return option.Some[string]("")
		}
		zeroval := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
			return w.ZeroExpr(gad.Package.Types, t.Type())
		}).MakeString(",")
		if zeroval != "" {
			return option.Some(fmt.Sprintf(`return %s`, zeroval))
		}
		return option.Some("return")
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
					return seq.Concat(as.Tuple[string, types.Type]("self", gad.Interface), r.argTypes)
				}
				return seq.Concat(as.Tuple[string, types.Type]("r", gad.Interface), r.argTypes)
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
					return option.Some(r.withReturn(true, "%s(%s)", buf.String(), args.Get().args.MakeString(",")))
				} else {

					return option.Some(r.withReturn(true, `%s(%s)`, types.ExprString(opt.DefaultImplExpr), args.Get().args.MakeString(",")))

				}
			} else {
				fmt.Printf("err : %s\n", args.Failed().Get())
			}
			if gad.Self {
				e := types.ExprString(opt.DefaultImplExpr)
				return option.Some(r.withReturn(true, `%s(self, %s)`, e, r.argStr))
			} else {
				e := types.ExprString(opt.DefaultImplExpr)
				return option.Some(r.withReturn(true, `%s(r, %s)`, e, r.argStr))
			}
		}

		fs := token.NewFileSet()

		buf := &bytes.Buffer{}

		printer.Fprint(buf, fs, opt.DefaultImplExpr)
		if sig.Results().Len() > 1 {
			zeroval := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
				return w.ZeroExpr(gad.Package.Types, t.Type())
			}).Drop(1).MakeString(",")
			return option.Some("return " + buf.String() + "," + zeroval)
		} else {
			return option.Some("return " + buf.String())
		}

	}
	return option.None[string]()

	//return fmt.Sprintf(`panic("%s.%s not implemented")`, r.adaptorTypeName, t.Name())
}

func fieldAndImplOfInterfaceImpl2(w genfp.Writer, gad genfp.GenerateAdaptorDirective, namedInterface types.Type, adaptorTypeName string, superField string, fieldMap fp.Map[string, genfp.TypeReference], methodSet fp.Set[string]) (fp.Seq[fp.Tuple2[string, string]], fp.Set[string]) {
	fmt.Printf("generate impl %s of %s\n", namedInterface.String(), adaptorTypeName)
	intf := namedInterface.Underlying().(*types.Interface)

	fields := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) fp.Tuple2[string, string] {
		if methodSet.Contains(t.Name()) {
			return as.Tuple("", "")
		}

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

		ctx := &implContext{
			namedInterface:  namedInterface,
			superField:      superField,
			adaptorTypeName: adaptorTypeName,
			w:               w,
			t:               t,
			gad:             gad,
			opt:             opt,
			valName:         valName,
			cbName:          cbName,
			selfarg:         selfarg,
			argTypes:        argTypes,
			argStr:          argStr,
			argTypeStr:      argTypeStr,
			resstr:          resstr,
			implArgs:        implArgs,
			fieldMap:        fieldMap,
		}

		if opt.Delegate != nil {
			ctx.superField = opt.Delegate.Field
		}

		defaultField, cbField := ctx.adaptorFields()
		defaultExpr := ctx.defaultImpl()
		callExtendsExpr, ceEnd := ctx.callExtends()
		cbExpr := option.FlatMap(cbField, func(v string) fp.Option[string] { return ctx.callCb() })
		panicExpr := option.Some(fmt.Sprintf(`panic("%s.%s not implemented")`, adaptorTypeName, t.Name()))
		valExpr, end := ctx.valOverride(defaultExpr.IsDefined() || callExtendsExpr.IsDefined() || cbExpr.IsDefined())
		panicExpr = panicExpr.FilterNot(func(v string) bool { return end || ceEnd }).FilterNot(func(v string) bool {
			return defaultExpr.IsDefined()
		})

		if opt.Delegate != nil {
			cbName = opt.Delegate.Field
			if isEmbeddingField(gad, opt.Delegate.Field) {
				if gad.Self && gad.ExtendsByEmbedding {
					methodSet = methodSet.Incl(t.Name())
					return as.Tuple("", fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
						`, adaptorTypeName, t.Name(), argTypeStr, resstr,
						ctx.withReturn(true, `r.%sImpl(r, %s)`, t.Name(), argStr),
					))
				}
				return as.Tuple("", "")
			}
		}
		impl := fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
						`, adaptorTypeName, implName, implArgs, resstr,
			fp.Seq[string]{}.
				Add(valExpr.OrElse("")).
				Concat(cbExpr.ToSeq()).
				Concat(callExtendsExpr.ToSeq()).
				Concat(defaultExpr.ToSeq()).
				Concat(panicExpr.ToSeq()).
				MakeString("\n\n"),
		)
		if gad.Self {
			impl = fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
					`, adaptorTypeName, t.Name(), argTypeStr, resstr,
				ctx.withReturn(true, "r.%sImpl(r,%s)", t.Name(), argStr),
			) + impl

		}
		methodSet = methodSet.Incl(t.Name())
		return as.Tuple(fp.Seq[string]{}.Concat(defaultField.ToSeq()).Concat(cbField.ToSeq()).MakeString("\n"), impl)

	})

	return fields, methodSet
}

func fieldAndImplOfInterfaceImpl(w genfp.Writer, gad genfp.GenerateAdaptorDirective, namedInterface types.Type, adaptorTypeName string, superField string, fieldMap fp.Map[string, genfp.TypeReference], methodSet fp.Set[string]) (fp.Seq[fp.Tuple2[string, string]], fp.Set[string]) {
	//fmt.Printf("generate impl %s of %s\n", namedInterface.String(), adaptorTypeName)
	intf := namedInterface.Underlying().(*types.Interface)

	fields := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) fp.Tuple2[string, string] {
		if methodSet.Contains(t.Name()) {
			return as.Tuple("", "")
		}

		opt := gad.Methods[t.Name()]
		sig := opt.Signature
		valName := fmt.Sprintf("Default%s", t.Name())
		cbName := fmt.Sprintf("%s%s", opt.Prefix, t.Name())
		cbNilCheck := true
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

		if opt.Delegate != nil {
			cbName = opt.Delegate.Field
			if isEmbeddingField(gad, opt.Delegate.Field) {
				if gad.Self && gad.ExtendsByEmbedding {
					methodSet = methodSet.Incl(t.Name())
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

			fieldTpe := fieldMap.Get(opt.Delegate.Field)
			if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
				zeroval := w.ZeroExpr(gad.Package.Types, fieldTpe.Get().Type)
				if zeroval != "nil" {
					cbNilCheck = false
				}

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
			if opt.Delegate != nil {
				fieldTpe := fieldMap.Get(opt.Delegate.Field)
				if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
					if types.IsInterface(fieldTpe.Get().Type) {
						return fmt.Sprintf("%s%s", callSuperImpl(opt.Delegate.Field), withReturn(false, `r.%s.%s(%s)`, opt.Delegate.Field, t.Name(), argStr))
					}
				}
				return withReturn(false, `r.%s.%s(%s)`, opt.Delegate.Field, t.Name(), argStr)
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
		if cbNilCheck == false {
			cbExpr = callcb
		}

		cbfield := fmt.Sprintf("%s func(%s%s) %s", cbName, selfarg, argTypeStr, resstr)
		if opt.Delegate != nil {
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
			if cbNilCheck == false {
				return ""
			}
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

		methodSet = methodSet.Incl(t.Name())
		return as.Tuple(cbfield, impl)

	})

	return fields, methodSet
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
