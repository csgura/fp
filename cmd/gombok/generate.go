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
	"github.com/csgura/fp/genfp/generator"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
	"github.com/csgura/fp/xtr"
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

func fillOption(ret generator.GenerateAdaptorDirective, intf *types.Interface) (generator.GenerateAdaptorDirective, error) {
	for i := 0; i < intf.NumMethods(); i++ {
		m := intf.Method(i)

		opt := ret.Methods[m.Name()]
		opt.Type = m

		for _, d := range ret.Delegate {

			if hasMethod(d.TypeOf.Type, m.Name()) {
				opt.Delegate = &generator.DelegateDirective{
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

		if !opt.ValOverride {
			opt.ValOverride = slices.Contains(ret.ValOverride, m.Name())
			if opt.ValOverride && !slices.Contains(ret.Getter, m.Name()) {
				opt.OmitGetterIfValOverride = true
			}
		}

		if !opt.ValOverrideUsingPtr {
			opt.ValOverrideUsingPtr = slices.Contains(ret.ValOverrideUsingPtr, m.Name())
			if opt.ValOverrideUsingPtr && !slices.Contains(ret.Getter, m.Name()) {
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

func typeDecl(pk genfp.WorkingPackage, w genfp.Writer, t generator.TypeReference) string {
	if t.Type != nil {
		return w.TypeName(pk, t.Type)
	}
	return t.StringExpr
}

func generateAdaptor(w genfp.Writer, gad generator.GenerateAdaptorDirective) {

	adaptorTypeName := gad.Name
	if adaptorTypeName == "" {
		adaptorTypeName = gad.Interface.Obj().Name() + "Adaptor"
	}

	fieldSet := fp.Map[string, generator.TypeReference]{}
	var fieldList []string
	methodSet := fp.Set[string]{}

	for _, i := range gad.EmbeddingInterface {
		efield := seq.Last(strings.Split(i.StringExpr, "."))
		if efield.IsDefined() {
			if !fieldSet.Contains(efield.Get()) {
				fieldList = append(fieldList, typeDecl(gad.Package, w, i))
			}
			fieldSet = fieldSet.Updated(efield.Get(), i)

			contains := slices.ContainsFunc(gad.Delegate, func(d generator.DelegateDirective) bool {
				return d.Field == efield.Get()
			})
			if !contains {
				gad.Delegate = append(gad.Delegate, generator.DelegateDirective{
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
				fieldList = append(fieldList, typeDecl(gad.Package, w, e))
			}

			fieldSet = fieldSet.Updated(efield.Get(), e)
		}
	}

	for k, e := range gad.ExtendsWith {
		if !fieldSet.Contains(k) {
			fieldList = append(fieldList, fmt.Sprintf("%s %s", k, typeDecl(gad.Package, w, e)))
		}
		fieldSet = fieldSet.Updated(k, e)
	}

	i1 := iterator.Map(iterator.FromSeq(gad.Delegate), func(v generator.DelegateDirective) string {
		return v.Field
	})

	for _, d := range gad.Delegate {
		exists := as.Seq(gad.ImplementsWith).Exists(func(v generator.TypeReference) bool {
			return v.StringExpr == d.TypeOf.StringExpr
		})
		if !exists {
			gad.ImplementsWith = append(gad.ImplementsWith, d.TypeOf)
		}
	}

	delegateFields := iterator.Sort(i1, ord.Given[string]()).FilterNot(fieldSet.Contains)

	for _, fn := range delegateFields {

		d := as.Seq(gad.Delegate).Find(func(v generator.DelegateDirective) bool {
			return v.Field == fn
		}).Get()

		fieldSet = fieldSet.Updated(fn, d.TypeOf)
		fieldList = append(fieldList, fmt.Sprintf("%s %s", fn, typeDecl(gad.Package, w, d.TypeOf)))

	}

	if gad.ExtendsByEmbedding {
		gad.Extends = true
	}

	superField := ""
	if gad.Extends {
		superField = "Extends"
	}

	intf := gad.Interface.Underlying().(*types.Interface)

	ti := metafp.GetTypeInfo(gad.Interface)
	decltp := ti.TypeParamDecl(w, gad.Package)
	valuetp := ti.TypeParamIns(w, gad.Package)

	gad, _ = fillOption(gad, intf)
	fields, methodSet := fieldAndImplOfInterfaceImpl2(w, gad, gad.Interface, adaptorTypeName+valuetp, superField, fieldSet, methodSet)

	for _, i := range gad.ImplementsWith {
		if i.Type != nil {
			if intf, ok := i.Type.Underlying().(*types.Interface); ok {
				gad, _ = fillOption(gad, intf)
				af, ms := fieldAndImplOfInterfaceImpl2(w, gad, i.Type, adaptorTypeName+valuetp, "", fieldSet, methodSet)
				fields = fields.Concat(af)
				methodSet = ms
			}
		}
	}

	for k, opt := range gad.Methods {
		if !methodSet.Contains(k) && opt.Delegate != nil {
			// gad.Method 는 사용자가 지정하지 않아도
			// 지정한 interface 와 embedding interface 의 method가 모두 추가 된다.
			// 그중에서  embedding interface 는  위에서 구현을 만들지 않지만
			// fillOption 에서 자동으로 추가된다.
			// 사용자가 명시적으로  지정한 경우에만 opt.ReceiverType.Type 이 nil 이 아님.
			if opt.ReceiverType.Type != nil {
				if intf, ok := opt.ReceiverType.Type.Underlying().(*types.Interface); ok {

					of := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) *types.Func {
						return t
					}).Find(func(v *types.Func) bool {
						return v.Name() == opt.Method
					})

					if of.IsDefined() && fp.IsInstanceOf[*types.Signature](of.Get().Type()) {

						gad, _ = fillOption(gad, intf)
						opt = gad.Methods[opt.Method]

						af, ms := generateImpl(opt, gad, of.Get(), w, intf, "", adaptorTypeName, fieldSet, methodSet)
						fields = fields.Append(af)
						methodSet = ms
					}

				}
			}

		}
	}

	extends := ""
	if gad.Extends && !gad.ExtendsByEmbedding {
		extends = "Extends " + w.TypeName(gad.Package, gad.Interface)
	}

	fieldDecl := seq.Of(extends)

	fieldDecl = fieldDecl.Concat(fieldList)

	fieldDecl = fieldDecl.Concat(seq.Map(fields, fp.Entry[string].Head).FilterNot(eq.GivenValue("")))

	fmt.Fprintf(w, `type %s%s struct {
					%s
				}
				`, adaptorTypeName, decltp, fieldDecl.MakeString("\n"))

	fmt.Fprintf(w, "%s", seq.Map(fields, fp.Entry[string].Tail).MakeString("\n"))

}

func removePkgPrefix(v string) string {
	dotidx := strings.IndexByte(v, '.')
	if dotidx > 0 {
		return v[dotidx+1:]
	}
	return v
}

func toTypeName(w genfp.ImportSet, workingPkg genfp.WorkingPackage, v metafp.TypeInfoExpr) genfp.TypeName {
	return genfp.TypeName{
		Package:   genfp.FromTypesPackage(v.Type.Pkg),
		Complete:  v.TypeName(w, workingPkg),
		Name:      v.Type.TypeName,
		IsPtr:     v.Type.IsPtr(),
		IsStruct:  v.Type.Underlying().IsStruct(),
		IsNilable: v.Type.IsNilable(),
		ZeroExpr:  w.ZeroExpr(workingPkg, v.Type.Type),
	}
}

func scanStructTypes(list []metafp.TypeInfo, ti metafp.TypeInfo, uniqCheck map[string]bool, recursive bool) []metafp.TypeInfo {

	if !ti.IsOption() && ti.Underlying().IsStruct() {
		if exists := uniqCheck[ti.ID]; !exists {
			uniqCheck[ti.ID] = true
			list = append(list, ti)
			if recursive {
				for _, f := range ti.Fields() {
					list = scanStructTypes(list, f.FieldType, uniqCheck, recursive)
				}
			}
		}
	}
	if recursive {
		if elem, ok := ti.ElemType().Unapply(); ok {
			list = scanStructTypes(list, elem, uniqCheck, recursive)
		}

		list = seq.Fold(ti.TypeArgs, list, func(l []metafp.TypeInfo, t metafp.TypeInfo) []metafp.TypeInfo {
			return scanStructTypes(l, t, uniqCheck, recursive)
		})
	}
	return list
}

func genGenerate() {
	pack := os.Getenv("GOPACKAGE")

	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	if err != nil {
		fmt.Printf("package load error : %s\n", err)
		return
	}

	workingPkg := genfp.NewWorkingPackage(pkgs[0].Types, pkgs[0].Fset, pkgs[0].Syntax)
	gentemplate := generator.FindGenerateFromUntil(pkgs, "@fp.Generate")
	genlist := generator.FindGenerateFromList(pkgs, "@fp.Generate")
	genstruct := generator.FindGenerateFromStructs(pkgs, "@fp.Generate")

	genadaptor := generator.FindGenerateAdaptor(pkgs, "@fp.Generate")
	monadf := generator.FindGenerateMonadFunctions(pkgs, "@fp.Generate")
	traversef := generator.FindGenerateTraverseFunctions(pkgs, "@fp.Generate")
	monadt := generator.FindGenerateMonadTransfomers(pkgs, "@fp.Generate")

	filelist := iterator.ToGoSet(
		mutable.MapOf(gentemplate).Keys().
			Concat(mutable.MapOf(genlist).Keys()).
			Concat(mutable.MapOf(genstruct).Keys()).
			Concat(mutable.MapOf(genadaptor).Keys()).
			Concat(mutable.MapOf(monadf).Keys()).
			Concat(mutable.MapOf(traversef).Keys()).
			Concat(mutable.MapOf(monadt).Keys()),
	)

	fileSet := map[string]bool{}
	for file := range filelist {
		fullpath := cwd + "/" + file
		fileSet[fullpath] = true
	}

	funcList := map[string]bool{}
	for _, p := range pkgs {
		s := p.Types.Scope()
		for _, n := range s.Names() {
			o := s.Lookup(n)
			if _, ok := o.Type().(*types.Signature); ok {
				file := p.Fset.Position(o.Pos()).Filename
				if !fileSet[file] {
					funcList[o.Name()] = true
				}

			}
		}
	}
	for file := range filelist {
		genfp.Generate(pack, file, func(w genfp.Writer) {
			for _, gfu := range genlist[file] {
				for _, im := range gfu.Imports {
					w.GetImportedName(genfp.NewImportPackage(im.Package, im.Name))
				}

				for _, v := range gfu.List {
					w.Render(gfu.Template, map[string]any{}, map[string]any{
						"N": v,
					})
				}
			}

			for _, gfu := range genstruct[file] {
				for _, im := range gfu.Imports {
					w.GetImportedName(genfp.NewImportPackage(im.Package, im.Name))
				}

				uniqCheck := map[string]bool{}
				all := seq.Fold(gfu.List, []metafp.TypeInfo{}, func(list []metafp.TypeInfo, tr generator.TypeReference) []metafp.TypeInfo {
					return scanStructTypes(list, metafp.GetTypeInfo(tr.Type), uniqCheck, gfu.Recursive)
				})

				for _, ti := range all {

					name := ti.Name()
					if name.IsDefined() {
						is := genfp.NewImportSet()
						fields := seq.Map(ti.Fields(), func(f metafp.StructField) genfp.StructFieldDef {
							tpe := toTypeName(is, workingPkg, f.TypeInfoExpr(workingPkg))
							indtpe := tpe
							if f.FieldType.IsPtr() {
								indtpe = toTypeName(is, workingPkg, metafp.TypeInfoExpr{
									Type: f.FieldType.ElemType().Get(),
								})
							}
							return genfp.StructFieldDef{
								Name:         f.Name,
								Type:         tpe,
								IndirectType: indtpe,
								Tag:          f.Tag,
								IsPublic:     f.Public(),
								IsVisible:    f.Public() || ti.IsSamePkg(workingPkg),
								ElemType: option.Map(f.FieldType.ElemType(), func(v metafp.TypeInfo) genfp.TypeName {
									return toTypeName(is, workingPkg, metafp.TypeInfoExpr{
										Type: v,
									})
								}).OrZero(),
							}
						})

						visible := fields.Filter(func(v genfp.StructFieldDef) bool {
							return v.IsVisible
						})
						st := genfp.StructDef{
							Name:      name.Get(),
							AllFields: fields,
							Fields:    visible,
							Type: toTypeName(is, workingPkg, metafp.TypeInfoExpr{
								Type: ti,
							}),
						}

						w.Render(gfu.Template, map[string]any{}, map[string]any{
							"N": st,
						})
					}
				}
			}

			for _, gfu := range gentemplate[file] {
				for _, im := range gfu.Imports {
					w.GetImportedName(genfp.NewImportPackage(im.Package, im.Name))
				}

				w.Iteration(gfu.From, gfu.Until).Write(gfu.Template, map[string]any{})
			}

			for _, gad := range genadaptor[file] {
				generateAdaptor(w, gad)
			}

			for _, md := range monadf[file] {
				generator.WriteMonadFunctions(w, md, funcList)
			}

			for _, md := range traversef[file] {
				generator.WriteTraverseFunctions(w, md, funcList)
			}

			for _, md := range monadt[file] {
				generator.WriteMonadTransformers(w, md, funcList)
			}
		})
	}

}

func isEmbeddingField(gad generator.GenerateAdaptorDirective, field string) bool {
	is := func(e []generator.TypeReference) bool {
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
	gad             generator.GenerateAdaptorDirective
	namedInterface  types.Type
	adaptorTypeName string
	superField      string
	t               *types.Func
	opt             generator.ImplOptionDirective
	valName         string
	cbName          string
	selfarg         string
	argTypes        fp.Seq[fp.Entry[typeExpr]]
	argStr          string
	argTypeStr      string
	resstr          string
	implArgs        string
	fieldMap        fp.Map[string, generator.TypeReference]
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

func (r *implContext) callSuperImpl(field string) fp.Option[string] {
	gad := r.gad

	fieldTpe := r.fieldMap.Get(field)
	if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
		if !types.IsInterface(fieldTpe.Get().Type.Underlying()) {
			return option.None[string]()
		}
	}

	args := r.matchSuperMethodArgs(field)
	argstr := option.Map(args, CallArgs.ArgList).OrElse(r.argStr)

	implArgs := option.Map(args, as.Func3(CallArgs.ArgTypeList).ApplyLast2(r.w, r.gad.Package)).Map(func(s string) string {
		if gad.ExtendsSelfCheck {
			return "self " + r.w.TypeName(gad.Package, gad.Interface) + "," + s
		}
		return s
	}).OrElse(r.implArgs)

	implName := func() string {
		if gad.ExtendsSelfCheck {
			return r.t.Name() + "Impl"
		}
		return r.t.Name()
	}()

	if gad.ExtendsSelfCheck {
		return option.Some(fmt.Sprintf(`type impl interface {
										%s(%s) %s
									}

									if super, ok := r.%s.(impl); ok {
										%s 
									}
									`, implName, implArgs, r.resstr,
			field,
			r.withReturn(false, "super.%s(self, %s)", implName, argstr),
		))
	}
	return option.None[string]()

}

type GeneratedExpr struct {
	expr             string
	unreachableAfter bool
}

func (r GeneratedExpr) Expr() string {
	return r.expr
}

func (r GeneratedExpr) UnreachableAfter() bool {
	return r.unreachableAfter
}

func someExpr(expr string) fp.Option[GeneratedExpr] {
	return option.Some(GeneratedExpr{expr: expr})
}

func finalExpr(expr string) fp.Option[GeneratedExpr] {
	return option.Some(GeneratedExpr{expr: expr, unreachableAfter: true})
}

type CallArgs struct {
	variadic bool
	avail    fp.Seq[fp.Entry[typeExpr]]
	args     fp.Seq[fp.Entry[typeExpr]]
}

func (r CallArgs) ArgList() string {
	return iterator.Map(iterator.FromSlice(r.args), xtr.Head).MakeString(",")
}

func (r CallArgs) ArgTypeList(w genfp.Writer, pk genfp.WorkingPackage) string {
	return iterator.Map(iterator.ZipWithIndex(iterator.FromSlice(r.args)), as.Tupled2(func(i int, v fp.Entry[typeExpr]) string {
		if r.variadic && i == r.args.Size()-1 {
			return fmt.Sprintf("%s... %s", v.I1, v.I2.TypeName(w, pk))
		}
		return fmt.Sprintf("%s %s", v.I1, v.I2.TypeName(w, pk))
	})).MakeString(",")
}

func (r *implContext) matchSuperMethodArgs(superField string) fp.Option[CallArgs] {
	fieldTpe := r.fieldMap.Get(superField)
	if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
		tpe := metafp.GetTypeInfo(fieldTpe.Get().Type)
		method := tpe.Method.Get(r.t.Name())
		if method.IsDefined() {

			os := method.Get().Type().(*types.Signature)

			return option.FromTry(r.matchFuncArgs(os))

		}
	}
	return option.None[CallArgs]()
}

type typeExpr struct {
	Type types.Type
	Expr fp.Option[ast.Expr]
}

func (r typeExpr) String() string {
	return r.Type.String()
}

func (r typeExpr) TypeName(w genfp.ImportSet, wp genfp.WorkingPackage) string {

	if expr, ok := r.Expr.Unapply(); ok {
		_, iset := wp.EvalTypeExpr(expr)
		for _, v := range iset {
			w.AddImport(v)
		}
		return types.ExprString(expr)
	}

	return w.TypeName(wp, r.Type)
}

func namedTypeExpr(tp *types.Named) typeExpr {
	return typeExpr{
		Type: tp,
	}
}

func (r *implContext) matchFuncArgs(ms *types.Signature) fp.Try[CallArgs] {
	gad := r.gad

	defImplArgs := iterate(ms.Params().Len(), ms.Params().At, func(i int, t *types.Var) typeExpr {
		return varTypeExpr(gad.Package, t)
	})

	availableArgs := func() fp.Seq[fp.Entry[typeExpr]] {
		if gad.ExtendsSelfCheck {
			return seq.Concat(as.Tuple[string, typeExpr]("self", namedTypeExpr(gad.Interface)), r.argTypes)
		}
		return seq.Concat(as.Tuple[string, typeExpr]("r", namedTypeExpr(gad.Interface)), r.argTypes)
	}()

	args := seq.FoldTry(defImplArgs, CallArgs{avail: availableArgs}, func(args CallArgs, tp typeExpr) fp.Try[CallArgs] {
		init, tail := iterator.Span(iterator.FromSeq(args.avail), func(t fp.Entry[typeExpr]) bool {
			return t.I2.String() != tp.String()
		})

		arg := tail.NextOption()
		if arg.IsDefined() {
			return try.Success(CallArgs{
				variadic: ms.Variadic(),
				avail:    init.Concat(tail).ToSeq(),
				args:     append(args.args, arg.Get()),
			})
		}
		return try.Failure[CallArgs](fp.Error(400, "can't find proper args for type %s", tp.String()))
	})
	return args
}

func (r *implContext) callExtends(superField string) fp.Option[GeneratedExpr] {
	gad := r.gad

	cbNilCheck := true

	fieldTpe := r.fieldMap.Get(superField)
	if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
		zeroval := r.w.ZeroExpr(gad.Package, fieldTpe.Get().Type)
		if zeroval != "nil" {
			cbNilCheck = false
		}
	}

	argstr := option.Map(r.matchSuperMethodArgs(superField), CallArgs.ArgList).OrElse(r.argStr)

	if superField != "" && !gad.ExtendsByEmbedding {
		si := r.callSuperImpl(superField)
		if cbNilCheck {

			sc := r.withReturn(false, "r.%s.%s(%s)", superField, r.t.Name(), argstr)

			return someExpr(fmt.Sprintf(`
					if r.%s != nil {
						%s
					}
				`, superField, as.Seq(si.ToSeq()).Add(sc).MakeString("\n"),
			))
		}

		sc := r.withReturn(false, "r.%s.%s(%s)", superField, r.t.Name(), argstr)
		return finalExpr(fmt.Sprintf(`
						%s
				`, as.Seq(si.ToSeq()).Add(sc).MakeString("\n"),
		))
	} else if gad.Extends {
		callSupercheck := func(field string) fp.Option[string] {
			if fieldTpe.IsDefined() && fieldTpe.Get().Type != nil {
				if !types.IsInterface(fieldTpe.Get().Type) {
					return option.None[string]()
				}
			}

			si := r.callSuperImpl(field)
			sc := fmt.Sprintf(`
					if super , ok := r.%s.(%s); ok {
						%s
					}
				`, field, r.w.TypeName(gad.Package, r.namedInterface),
				r.withReturn(false, "super.%s(%s)", r.t.Name(), r.argStr),
			)

			return option.Some(as.Seq(si.ToSeq()).Add(sc).MakeString("\n"))
		}

		return option.Map(callSupercheck("Extends"), func(s string) GeneratedExpr {
			return GeneratedExpr{expr: fmt.Sprintf(`
				if r.Extends != nil {
					%s
				}
			`, s)}
		})
	}
	return option.None[GeneratedExpr]()
}

func (r *implContext) adaptorFields() (fp.Option[string], fp.Option[string]) {
	opt := r.opt

	gad := r.gad
	sig := r.opt.Signature

	cbfield := option.Some(fmt.Sprintf("%s func(%s%s) %s", r.cbName, r.selfarg, r.argTypeStr, r.resstr))

	cbfield = cbfield.FilterNot(func(v string) bool { return opt.Delegate != nil && opt.Private })

	cbfield = cbfield.FilterNot(func(v string) bool {
		return r.isValOverride() && opt.OmitGetterIfValOverride
	})

	cbfield = cbfield.FilterNot(fp.Const[string](opt.Private))

	defaultField := option.Map(option.Of(r.isValOverride()).Filter(fp.Id), func(v bool) string {
		if opt.ValOverrideUsingPtr {
			return fmt.Sprintf("%s *%s", r.valName, varTypeName(r.w, gad.Package, sig.Results().At(0)))
		}
		return fmt.Sprintf("%s %s", r.valName, varTypeName(r.w, gad.Package, sig.Results().At(0)))
	})

	return defaultField, cbfield
}

func (r *implContext) isValOverride() bool {
	opt := r.opt
	sig := r.opt.Signature
	return (opt.ValOverride || opt.ValOverrideUsingPtr) && sig.Params().Len() == 0 && sig.Results().Len() == 1
}

func hasNil(t types.Type) bool {
	switch u := t.Underlying().(type) {
	case *types.Basic:
		return u.Kind() == types.UnsafePointer
	case *types.Slice, *types.Pointer, *types.Signature, *types.Map, *types.Chan:
		return true
	case *types.Interface:
		return true
	}
	return false
}

func (r *implContext) valOverride(defaultImpl bool) (fp.Option[string], bool) {
	sig := r.opt.Signature
	w := r.w
	gad := r.gad
	valName := r.valName
	valoverride := r.isValOverride()
	if valoverride {
		if r.opt.ValOverrideUsingPtr {
			ret := fmt.Sprintf(`if r.%s != nil {
								return *r.%s
							}
					
					`, valName,
				valName)
			return option.Some(ret), false
		}

		zeroVal := w.ZeroExpr(gad.Package, sig.Results().At(0).Type())
		if hasNil(sig.Results().At(0).Type()) {
			ret := fmt.Sprintf(`if r.%s != nil {
								return r.%s
							}
					
					`, valName,
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
					
					`, varTypeName(r.w, gad.Package, sig.Results().At(0)),
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

func exprString(expr ast.Expr) string {
	fs := token.NewFileSet()

	buf := &bytes.Buffer{}

	printer.Fprint(buf, fs, expr)
	return buf.String()
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
			return w.ZeroExpr(gad.Package, t.Type())
		}).MakeString(",")
		if zeroval != "" {
			return option.Some(fmt.Sprintf(`return %s`, zeroval))
		}
		return option.Some("return")
	} else if opt.DefaultImplExpr != nil {

		for _, i := range opt.DefaultImplImports {
			w.AddImport(genfp.NewImportPackage(i.Package, i.Name))
		}

		if opt.DefaultImplSignature != nil {
			os := opt.DefaultImplSignature

			args := r.matchFuncArgs(os)

			if args.IsSuccess() {
				if fl, ok := opt.DefaultImplExpr.(*ast.FuncLit); ok {
					fs := token.NewFileSet()

					buf := &bytes.Buffer{}

					printer.Fprint(buf, fs, fl)
					return option.Some(r.withReturn(true, "%s(%s)", buf.String(), args.Get().ArgList()))
				} else {

					return option.Some(r.withReturn(true, `%s(%s)`, exprString(opt.DefaultImplExpr), args.Get().ArgList()))

				}
			} else {
				fmt.Printf("err : %s\n", args.Failed().Get())
			}
			if gad.Self {
				e := exprString(opt.DefaultImplExpr)
				return option.Some(r.withReturn(true, `%s(self, %s)`, e, r.argStr))
			} else {
				e := exprString(opt.DefaultImplExpr)
				return option.Some(r.withReturn(true, `%s(r, %s)`, e, r.argStr))
			}
		}

		fs := token.NewFileSet()

		buf := &bytes.Buffer{}

		printer.Fprint(buf, fs, opt.DefaultImplExpr)
		if sig.Results().Len() > 1 {
			zeroval := iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
				return w.ZeroExpr(gad.Package, t.Type())
			}).Drop(1).MakeString(",")
			return option.Some("return " + buf.String() + "," + zeroval)
		} else {
			return option.Some("return " + buf.String())
		}

	}
	return option.None[string]()

	//return fmt.Sprintf(`panic("%s.%s not implemented")`, r.adaptorTypeName, t.Name())
}

func argName(i int, t *types.Var) string {
	var name = t.Name()
	if name == "" {
		name = fmt.Sprintf("a%d", i+1)
	}
	return name
}

func fieldAndImplOfInterfaceImpl2(w genfp.Writer, gad generator.GenerateAdaptorDirective, namedInterface types.Type, adaptorTypeName string, superField string, fieldMap fp.Map[string, generator.TypeReference], methodSet fp.Set[string]) (fp.Seq[fp.Entry[string]], fp.Set[string]) {
	//fmt.Printf("generate impl %s of %s\n", namedInterface.String(), adaptorTypeName)
	intf := namedInterface.Underlying().(*types.Interface)

	fields := iterate(intf.NumMethods(), intf.Method, func(i int, t *types.Func) fp.Entry[string] {
		if methodSet.Contains(t.Name()) {
			return as.Tuple("", "")
		}

		opt := gad.Methods[t.Name()]

		// 아무것도 리턴 안할 때는   TypeName 호출하면 안됨.
		// 아무것도 안하는 리턴이기 때문에 ,  위에 중간에 리턴하는 코드에서 이미 처리함.
		// 여긴 절대로 실행안됨.
		//fmt.Printf("generate method %s (super:%s) impl %s of %s \n", t.Name(), ctx.superField, namedInterface.String(), adaptorTypeName)
		ret, ms := generateImpl(opt, gad, t, w, namedInterface, superField, adaptorTypeName, fieldMap, methodSet)
		methodSet = ms
		return ret

	})

	return fields, methodSet
}

func varTypeExpr(wp genfp.WorkingPackage, t *types.Var) typeExpr {
	retexpr := wp.FindNode(t.Pos())
	tpExpr := func() ast.Expr {
		if fl, ok := retexpr.(*ast.FieldList); ok {
			return fl.List[0].Type
		} else if f, ok := retexpr.(*ast.Field); ok {
			return f.Type
		}
		return nil

	}()
	return typeExpr{
		Type: t.Type(),
		Expr: option.NonZero(tpExpr),
	}
}

func varTypeName(w genfp.ImportSet, wp genfp.WorkingPackage, t *types.Var) string {
	retexpr := wp.FindNode(t.Pos())
	tpExpr := func() ast.Expr {
		if fl, ok := retexpr.(*ast.FieldList); ok {
			return fl.List[0].Type
		} else if f, ok := retexpr.(*ast.Field); ok {
			return f.Type
		}
		return nil

	}()

	if tpExpr != nil {
		_, iset := wp.EvalTypeExpr(tpExpr)
		for _, v := range iset {
			w.AddImport(v)
		}
		return types.ExprString(tpExpr)
	}
	return w.TypeName(wp, t.Type())
}

func elemTypeExpr(w genfp.ImportSet, wp genfp.WorkingPackage, t *types.Var) string {
	st := t.Type().(*types.Slice)

	retexpr := wp.FindNode(t.Pos())
	tpExpr := func() ast.Expr {
		if fl, ok := retexpr.(*ast.FieldList); ok {
			return fl.List[0].Type
		} else if f, ok := retexpr.(*ast.Field); ok {
			return f.Type
		}
		return nil

	}()

	if elt, ok := tpExpr.(*ast.Ellipsis); ok {
		tpExpr = elt.Elt
	}

	if tpExpr != nil {
		_, iset := wp.EvalTypeExpr(tpExpr)
		for _, v := range iset {
			w.AddImport(v)
		}
		return types.ExprString(tpExpr)
	}
	return w.TypeName(wp, st.Elem())
}

func generateImpl(opt generator.ImplOptionDirective, gad generator.GenerateAdaptorDirective, t *types.Func, w genfp.Writer, namedInterface types.Type, superField string, adaptorTypeName string, fieldMap fp.Map[string, generator.TypeReference], methodSet fp.Set[string]) (fp.Entry[string], fp.Set[string]) {
	if opt.Delegate != nil && opt.Private {
		if isEmbeddingField(gad, opt.Delegate.Field) {
			if !gad.Self || !gad.ExtendsByEmbedding {

				return as.Tuple("", ""), methodSet
			}
		}
	}

	sig := opt.Signature
	valName := fmt.Sprintf("Default%s", t.Name())
	cbName := fmt.Sprintf("%s%s", opt.Prefix, t.Name())
	if opt.Name != "" {
		cbName = fmt.Sprintf("%s%s", opt.Prefix, opt.Name)
		valName = fmt.Sprintf("Default%s", opt.Name)
	}

	selfarg := ""
	if gad.Self {
		selfarg = "self " + w.TypeName(gad.Package, gad.Interface) + ","
	}

	argTypes := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) fp.Entry[typeExpr] {
		return as.Tuple2(argName(i, t), varTypeExpr(gad.Package, t))
	})

	argStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
		if sig.Variadic() && i == sig.Params().Len()-1 {
			return fmt.Sprintf("%s...", argName(i, t))
		} else {

			return argName(i, t)
		}
	}).MakeString(",")

	argTypeStr := iterate(sig.Params().Len(), sig.Params().At, func(i int, t *types.Var) string {
		if sig.Variadic() && i == sig.Params().Len()-1 {
			return fmt.Sprintf("%s ...%s", argName(i, t), elemTypeExpr(w, gad.Package, t))
		}
		return fmt.Sprintf("%s %s", argName(i, t), varTypeName(w, gad.Package, t))
	}).MakeString(",")

	resstr := ""
	if sig.Results().Len() > 0 {
		resstr = "(" + iterate(sig.Results().Len(), sig.Results().At, func(i int, t *types.Var) string {
			return varTypeName(w, gad.Package, t)
		}).MakeString(",") + ")"
	}

	implArgs := argTypeStr
	if gad.ExtendsSelfCheck {
		implArgs = "self " + w.TypeName(gad.Package, gad.Interface) + "," + argTypeStr
	}

	implName := func() string {
		if gad.ExtendsSelfCheck {
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

	if opt.Delegate != nil && opt.Private {
		if isEmbeddingField(gad, opt.Delegate.Field) {
			if gad.Self && gad.ExtendsByEmbedding {
				methodSet = methodSet.Incl(t.Name())
				return as.Tuple("", fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
						`, adaptorTypeName, t.Name(), argTypeStr, resstr,
					ctx.withReturn(true, `r.%sImpl(r, %s)`, t.Name(), argStr),
				)), methodSet
			}

			return as.Tuple("", ""), methodSet
		}
	}

	defaultField, cbField := ctx.adaptorFields()
	defaultExpr := ctx.defaultImpl()

	delegateExpr := option.FlatMap(option.Ptr(opt.Delegate), func(v generator.DelegateDirective) fp.Option[GeneratedExpr] {
		return ctx.callExtends(v.Field)
	})

	unreachable := delegateExpr.IsDefined() && delegateExpr.Get().UnreachableAfter()

	callExtendsExpr := ctx.callExtends(ctx.superField).FilterNot(fp.Const[GeneratedExpr](unreachable))
	unreachable = unreachable || callExtendsExpr.IsDefined() && callExtendsExpr.Get().unreachableAfter

	cbExpr := option.FlatMap(cbField, func(v string) fp.Option[string] { return ctx.callCb() })
	panicExpr := option.Some(fmt.Sprintf(`panic("%s.%s not implemented")`, adaptorTypeName, t.Name()))
	valExpr, end := ctx.valOverride(defaultExpr.IsDefined() || callExtendsExpr.IsDefined() || cbExpr.IsDefined())
	cbExpr = cbExpr.FilterNot(func(v string) bool { return end })
	callExtendsExpr = callExtendsExpr.FilterNot(func(v GeneratedExpr) bool { return end })
	defaultExpr = defaultExpr.FilterNot(func(v string) bool { return end })

	panicExpr = panicExpr.FilterNot(func(v string) bool { return end || unreachable }).FilterNot(func(v string) bool {
		return defaultExpr.IsDefined()
	})

	impl := fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
						`, adaptorTypeName, implName, implArgs, resstr,
		fp.Seq[string]{}.
			Add(valExpr.OrElse("")).
			Concat(cbExpr.ToSeq()).
			Concat(option.Map(delegateExpr, GeneratedExpr.Expr).ToSeq()).
			Concat(option.Map(callExtendsExpr, GeneratedExpr.Expr).ToSeq()).
			Concat(defaultExpr.ToSeq()).
			Concat(panicExpr.ToSeq()).
			MakeString("\n\n"),
	)
	if gad.ExtendsSelfCheck {
		impl = fmt.Sprintf(`
						func (r *%s) %s(%s) %s {
							%s
						}
					`, adaptorTypeName, t.Name(), argTypeStr, resstr,
			ctx.withReturn(true, "r.%sImpl(r,%s)", t.Name(), argStr),
		) + impl

	}
	methodSet = methodSet.Incl(t.Name())
	return as.Tuple(fp.Seq[string]{}.Concat(defaultField.ToSeq()).Concat(cbField.ToSeq()).MakeString("\n"), impl), methodSet
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
